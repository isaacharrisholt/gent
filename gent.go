package gent

import (
	"cmp"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)
var reservedNodeMethods = getNodeMethodNames()

type Generator struct {
	options GeneratorOptions
}

type GeneratorOptions struct {
	PackageName string
	// Run the generator in debug mode, adding extra comments to the generated file
	// and printing debug logs to stderr.
	Debug bool
	// TODO: Add more options
}

func NewGenerator(options GeneratorOptions) *Generator {
	return &Generator{
		options: options,
	}
}

type nodeTypes []nodeType

type nodeType struct {
	Type     string                                      `json:"type"`
	Named    bool                                        `json:"named"`
	Fields   orderedmap.OrderedMap[string, nodeChildren] `json:"fields"`
	Children nodeChildren                                `json:"children"`
	Subtypes []nodeChildType                             `json:"subtypes"`
}

type nodeChildren struct {
	Multiple bool            `json:"multiple"`
	Required bool            `json:"required"`
	Types    []nodeChildType `json:"types"`
}

type nodeChildType struct {
	Type  string `json:"type"`
	Named bool   `json:"named"`
}

// Node types that are made of other types and require a union type
// to be created.
type unionType struct {
	name    string
	members []nodeChildType
}

type methodDef struct {
	name       string
	returnType string
	array      bool
	tsKinds    []string
}

type structDef struct {
	name string
	// Tree-sitter node kind
	tsKind            string
	methods           []methodDef
	isUnionType       bool
	childrenMethodDef *methodDef
}

// Inner map keys are all Tree-sitter node names
type nodeMap struct {
	// Top-level node definitions, supertypes.
	// Values are struct names.
	namedExported *orderedmap.OrderedMap[string, string]

	// Symbols, etc. that have `"named": false` but still need exporting.
	// Values are struct names.
	unnamedExported *orderedmap.OrderedMap[string, string]

	// Supertypes. Values are `unionType` structs.
	supertypes *orderedmap.OrderedMap[string, unionType]

	// Types made from a combination of other types, e.g. the multiple types possible in
	// `fields.types`.
	// Values are `unionType` structs.
	unionTypes *orderedmap.OrderedMap[string, unionType]

	// Types that are present in `fields` or `children` but are not defined at the top
	// level. We always keep references to these and create an empty, private struct for
	// them.
	// Values are struct names.
	unknown *orderedmap.OrderedMap[string, string]
}

func newNodeMap() nodeMap {
	return nodeMap{
		namedExported:   orderedmap.New[string, string](),
		unnamedExported: orderedmap.New[string, string](),
		supertypes:      orderedmap.New[string, unionType](),
		unionTypes:      orderedmap.New[string, unionType](),
		unknown:         orderedmap.New[string, string](),
	}
}

func (nm *nodeMap) registerNodeType(nodeType nodeType) {
	structName := generateStructName(nodeType.Type, nodeType.Named)
	if nodeType.Named {
		nm.namedExported.Set(nodeType.Type, structName)
		return
	}

	nm.unnamedExported.Set(nodeType.Type, structName)
}

func (nm *nodeMap) getStructName(typeName string, named bool) (string, bool) {
	// Check relevant named maps, then check unknown types
	if named {
		if structName, ok := nm.namedExported.Get(typeName); ok {
			return structName, true
		}
		if structName, ok := nm.supertypes.Get(typeName); ok {
			return structName.name, true
		}
	}

	if !named {
		if structName, ok := nm.unnamedExported.Get(typeName); ok {
			return structName, true
		}
	}

	structName, ok := nm.unknown.Get(typeName)
	return structName, ok
}

func (nm *nodeMap) registerSupertype(typeName string, members []nodeChildType) {
	structName := generateStructName(typeName, true)
	nm.supertypes.Set(typeName, unionType{name: structName, members: members})
}

func (nm *nodeMap) registerUnionType(members []nodeChildType) error {
	if len(members) < 2 {
		return fmt.Errorf("Cannot create union type with less than 2 members")
	}

	structTypeNames := []string{}
	tsNodeTypeNames := []string{}
	slices.SortFunc(members, func(a, b nodeChildType) int {
		return cmp.Compare(a.Type, b.Type)
	})

	for _, type_ := range members {
		structTypeNames = append(structTypeNames, createPrivateName(type_.Type))
		tsNodeTypeNames = append(tsNodeTypeNames, type_.Type)
	}

	structTypeName := strings.Join(structTypeNames, "_")
	tsNodeTypeName := strings.Join(tsNodeTypeNames, "_")

	nm.unionTypes.Set(tsNodeTypeName, unionType{
		name:    structTypeName,
		members: members,
	})

	return nil
}

func (nm *nodeMap) getUnionType(types []nodeChildType) (unionType, bool) {
	if len(types) < 2 {
		return unionType{}, false
	}

	tsNodeTypeNames := []string{}
	slices.SortFunc(types, func(a, b nodeChildType) int {
		return cmp.Compare(a.Type, b.Type)
	})

	for _, type_ := range types {
		tsNodeTypeNames = append(tsNodeTypeNames, type_.Type)
	}

	tsNodeTypeName := strings.Join(tsNodeTypeNames, "_")

	ut, ok := nm.unionTypes.Get(tsNodeTypeName)
	if !ok {
		return unionType{}, false
	}

	return ut, true
}

func (nm *nodeMap) registerUnknownType(typeName string) {
	structName := "Unknown__" + createPrivateName(typeName)
	nm.unknown.Set(typeName, structName)
}

// Method definitions describe what methods are available on the generated
// struct. They also store all the possible Tree-sitter node kinds that
// a particular field can represent.
// They have to do this recursively, as a supertype might reference another
// supertype, and so on.
func (nm *nodeMap) getTSRecursiveTSKinds(typeName string) []string {
	typesToCheck := []string{typeName}
	tsKinds := []string{}

	appendUnionTypeMembers := func(ut unionType) {
		for _, member := range ut.members {
			typesToCheck = append(typesToCheck, member.Type)
		}
	}

	for len(typesToCheck) > 0 {
		typeName := typesToCheck[0]
		typesToCheck = typesToCheck[1:]

		if superType, ok := nm.supertypes.Get(typeName); ok {
			appendUnionTypeMembers(superType)
			// Can't be a supertype and a non-supertype union
			continue
		}

		if unionType, ok := nm.unionTypes.Get(typeName); ok {
			appendUnionTypeMembers(unionType)
			continue
		}

		// Otherwise, it's a base TS kind
		tsKinds = append(tsKinds, typeName)
	}

	return tsKinds
}

func (g *Generator) Generate(data []byte) (string, error) {
	var nodeTypes nodeTypes
	err := json.Unmarshal(data, &nodeTypes)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}

	nodeTypeMap := make(map[string]nodeType)
	for _, nodeType := range nodeTypes {
		nodeTypeMap[nodeType.Type] = nodeType
	}

	packageName := "node_types"
	if g.options.PackageName != "" {
		packageName = g.options.PackageName
	}

	file := jen.NewFile(packageName)

	file.ImportName("github.com/tree-sitter/go-tree-sitter", "tree_sitter")

	nm := newNodeMap()

	// Get all the node types available in the file
	for _, nodeType := range nodeTypes {
		// Collect all the union types in the file and register them in the slice
		// All supertypes are considered exported union types, and we create private
		// union types from `fields` and `children` of other types.
		if nodeType.Subtypes != nil {
			// This is a supertype, so we always export it and use the type name
			// as the name of the union type.
			nm.registerSupertype(nodeType.Type, nodeType.Subtypes)
		} else {
			// Top level names go straight into the map
			nm.registerNodeType(nodeType)

			// Check the fields and children of the type
			for name, children := range nodeType.Fields.FromOldest() {
				if len(children.Types) > 1 {
					// This is a union type, so we create a private union type
					// from the types in the field.
					err = nm.registerUnionType(children.Types)
					if err != nil {
						return "", fmt.Errorf("Failed to register a union type for %s.%s: %w", nodeType.Type, name, err)
					}
				}
			}

			if len(nodeType.Children.Types) > 1 {
				err = nm.registerUnionType(nodeType.Children.Types)
				if err != nil {
					return "", fmt.Errorf("Failed to register a union type for %s children: %w", nodeType.Type, err)
				}
			}
		}
	}

	// Check through again to see if there are any types used in fields/children that
	// aren't defined. At time of writing, Python node-types.json uses `as_pattern_target`
	// which is never declared.
	//
	// If we find any, we'll create private type names for them and mark them as unknown.
	for _, nodeType := range nodeTypes {
		for _, field := range nodeType.Fields.FromOldest() {
			for _, fieldType := range field.Types {
				_, ok := nm.getStructName(fieldType.Type, fieldType.Named)
				if !ok {
					nm.registerUnknownType(fieldType.Type)
				}
			}
		}

		for _, childType := range nodeType.Children.Types {
			_, ok := nm.getStructName(childType.Type, childType.Named)
			if !ok {
				nm.registerUnknownType(childType.Type)
			}
		}
	}

	if g.options.Debug {
		fmt.Println(nm)
	}

	// Create an enum for the public node types
	file.Type().Id("SyntaxKind").Op("=").String()
	var publicTypes []jen.Code

	addPublicType := func(name string, tsKindName string) {
		publicTypes = append(publicTypes, jen.Id("SyntaxKind_"+name).Id("SyntaxKind").Op("=").Lit(tsKindName))
	}

	if g.options.Debug {
		publicTypes = append(publicTypes, file.Comment("Named types"))
	}
	for tsKindName, structName := range nm.namedExported.FromOldest() {
		addPublicType(structName, tsKindName)
	}
	if g.options.Debug {
		publicTypes = append(publicTypes, file.Comment("Unnamed types"))
	}
	for tsKindName, structName := range nm.unnamedExported.FromOldest() {
		addPublicType(structName, tsKindName)
	}
	if g.options.Debug {
		publicTypes = append(publicTypes, file.Comment("Supertypes"))
	}
	for tsKindName, unionType := range nm.supertypes.FromOldest() {
		addPublicType(unionType.name, tsKindName)
	}

	file.Var().Defs(publicTypes...)

	if g.options.Debug {
		file.Comment("\nGENERAL NODES\n")
	}
	for _, nodeType := range nodeTypes {
		if nodeType.Subtypes != nil {
			continue
		}

		if g.options.Debug {
			fmt.Printf("Adding node type %s\n", nodeType.Type)
		}
		err := addNodeType(file, nodeType, &nm)
		if err != nil {
			return "", fmt.Errorf("Failed to add node type %s: %w", nodeType.Type, err)
		}
	}

	if g.options.Debug {
		file.Comment("\nSUPERTYPES\n")
	}
	for _, supertype := range nm.supertypes.FromOldest() {
		err := addUnionType(file, supertype, &nm)
		if err != nil {
			return "", fmt.Errorf("Failed to add supertype %s: %w", supertype.name, err)
		}
	}

	if g.options.Debug {
		file.Comment("\nUNION TYPES\n")
	}
	for _, unionType := range nm.unionTypes.FromOldest() {
		err := addUnionType(file, unionType, &nm)
		if err != nil {
			return "", fmt.Errorf("Failed to add union type %s: %w", unionType.name, err)
		}
	}

	// Add empty structs for the unknown types. They can be private.
	if g.options.Debug {
		file.Comment("\nUNKNOWN TYPES\n")
	}
	for tsKind, unknownType := range nm.unknown.FromOldest() {
		writeStruct(file, structDef{
			name:    unknownType,
			tsKind:  tsKind,
			methods: []methodDef{},
		})
	}

	outputBuilder := &strings.Builder{}
	err = file.Render(outputBuilder)
	if err != nil {
		return "", fmt.Errorf("Failed to render file: %w", err)
	}

	return outputBuilder.String(), nil
}

// generateStructName generates a private or exported struct name based on the given
// nodeType. It will also add an `Unnamed_` prefix if the node is not a named node.
func generateStructName(typeName string, named bool) string {
	structName := createExportedName(typeName)
	if !named {
		structName = "Unnamed_" + structName
	}
	return structName
}

func addNodeType(file *jen.File, nodeType nodeType, nm *nodeMap) error {
	methodDefs := []methodDef{}

	for name, field := range nodeType.Fields.FromOldest() {
		typeName := ""
		tsKinds := []string{}

		if len(field.Types) == 1 {
			fieldType := field.Types[0]
			structName, ok := nm.getStructName(fieldType.Type, fieldType.Named)
			if !ok {
				return fmt.Errorf("Failed to find TS node %s in map", field.Types[0].Type)
			}
			typeName = structName
			tsKinds = append(tsKinds, nm.getTSRecursiveTSKinds(fieldType.Type)...)
		} else {
			unionType, ok := nm.getUnionType(field.Types)
			if !ok {
				return fmt.Errorf("Failed to find union type for %s.%s types", nodeType.Type, name)
			}
			typeName = unionType.name
			tsKinds = append(tsKinds, nm.getTSRecursiveTSKinds(unionType.name)...)
		}

		methodDefs = append(methodDefs, methodDef{
			name:       createPrivateName(name),
			returnType: typeName,
			array:      field.Multiple,
			tsKinds:    tsKinds,
		})
	}

	structName, ok := nm.getStructName(nodeType.Type, nodeType.Named)
	if !ok {
		return fmt.Errorf("Failed to find struct name for %s", nodeType.Type)
	}

	// Add `Children` method if required
	var childrenMethodDef *methodDef
	if len(nodeType.Children.Types) > 0 {
		methodName := "TypedChild"
		if nodeType.Children.Multiple {
			methodName = "TypedChildren"
		}

		if len(nodeType.Children.Types) == 1 {
			child := nodeType.Children.Types[0]
			childNodeName, ok := nm.getStructName(child.Type, child.Named)
			if !ok {
				return fmt.Errorf("Failed to find struct name for child %s of %s", child.Type, nodeType.Type)
			}
			childrenMethodDef = &methodDef{
				name:       methodName,
				returnType: childNodeName,
				array:      nodeType.Children.Multiple,
				tsKinds:    nm.getTSRecursiveTSKinds(child.Type),
			}
		} else {
			unionType, ok := nm.getUnionType(nodeType.Children.Types)
			if !ok {
				return fmt.Errorf("Failed to find union type name for children of %s", nodeType.Type)
			}
			tsKinds := nm.getTSRecursiveTSKinds(unionType.name)
			childrenMethodDef = &methodDef{
				name:       methodName,
				returnType: unionType.name,
				array:      nodeType.Children.Multiple,
				tsKinds:    tsKinds,
			}
		}
	}

	writeStruct(file, structDef{
		name:              structName,
		tsKind:            nodeType.Type,
		methods:           methodDefs,
		isUnionType:       false,
		childrenMethodDef: childrenMethodDef,
	})

	return nil
}

func addUnionType(file *jen.File, unionType unionType, nm *nodeMap) error {
	// A union type is a struct containing fields for each of the types in the
	// union. These are all pointers to indicate that any of them could be nil.
	// The types of the fields should always be exported.
	methodDefs := []methodDef{}
	for _, member := range unionType.members {
		typeName, ok := nm.getStructName(member.Type, member.Named)
		if !ok {
			return fmt.Errorf("Failed to find struct name for %s.%s", unionType.name, member.Type)
		}
		methodDefs = append(methodDefs, methodDef{
			name:       createPrivateName(member.Type),
			returnType: typeName,
			array:      false,
			tsKinds:    nm.getTSRecursiveTSKinds(member.Type),
		})
	}

	writeStruct(file, structDef{
		name:        unionType.name,
		methods:     methodDefs,
		isUnionType: true,
	})

	return nil
}

func writeStruct(file *jen.File, stDef structDef) {
	embedField := jen.Qual("github.com/tree-sitter/go-tree-sitter", "Node")
	structFields := []jen.Code{embedField}

	// for _, fieldDef := range def.fields {
	// 	stmt := jen.Id(fieldDef.name)
	// 	if fieldDef.pointer {
	// 		stmt = stmt.Op("*")
	// 	}
	// 	if fieldDef.array {
	// 		stmt = stmt.Index()
	// 	}
	// 	stmt.Id(fieldDef.typeName)
	// 	structFields = append(structFields, stmt)
	// }

	file.Type().Id(stDef.name).Struct(structFields...)

	structMethodIdentifier := strings.ToLower(string(stDef.name[0]))
	for _, fieldDef := range stDef.methods {
		funcName := strings.ToUpper(string(fieldDef.name[0])) + fieldDef.name[1:]

		// Need to make sure we don't override any of the reserved node methods,
		// so prefix with 'Get' until the name is unique.
		for slices.Contains(reservedNodeMethods, funcName) {
			funcName = "Get" + funcName
		}

		returnTypeStmt := jen.Null()
		if fieldDef.array {
			returnTypeStmt = returnTypeStmt.Index()
		}
		returnTypeStmt = returnTypeStmt.Id(fieldDef.returnType)

		// Default return type includes an error, but for arrays, we just return an empty
		// array
		if !fieldDef.array {
			returnTypeStmt = jen.Parens(jen.List(
				jen.List(
					returnTypeStmt,
					jen.Error(),
				),
			))
		}

		functionParams := jen.Null()
		var functionBody []jen.Code

		if stDef.isUnionType {
			tsKindsVarName := "tsKinds"
			tsKindsArray := []jen.Code{}
			for _, tsKind := range fieldDef.tsKinds {
				tsKindsArray = append(tsKindsArray, jen.Lit(tsKind))
			}
			tsKindsArrayLit := jen.Index().String().Values(tsKindsArray...)
			functionBody = []jen.Code{
				jen.Id(tsKindsVarName).Op(":=").Add(tsKindsArrayLit),
				jen.If(
					jen.
						Qual("slices", "Contains").
						Call(
							jen.Id(tsKindsVarName),
							jen.
								Id(structMethodIdentifier).
								Dot("Node").
								Dot("Kind").
								Call(),
						).
						Block(
							jen.Return(
								jen.Id(fieldDef.returnType).Values(jen.Dict{
									jen.Id("Node"): jen.Id(structMethodIdentifier).Dot("Node"),
								}),
								jen.Nil(),
							),
						),
				),
				jen.Return(
					jen.Id(fieldDef.returnType).Values(),
					jen.Qual("fmt", "Errorf").Call(
						jen.Lit("Node is a %s, not in %v"),
						jen.Id(structMethodIdentifier).Dot("Node").Dot("Kind").Call(),
						jen.Id(tsKindsVarName),
					),
				),
			}
		} else if fieldDef.array {
			// Take in a tree cursor for iterating
			cursorVarName := "cursor"
			functionParams = jen.Id(cursorVarName).Op("*").Qual("github.com/tree-sitter/go-tree-sitter", "TreeCursor")

			singularVarName := "child"
			pluralVarName := "children"
			outputVarName := "output"

			functionBody = []jen.Code{
				jen.Id(pluralVarName).
					Op(":=").
					Id(structMethodIdentifier).
					Dot("Node").
					Dot("ChildrenByFieldName").
					Call(
						jen.Lit(fieldDef.name),
						jen.Id(cursorVarName),
					),
				jen.Id(outputVarName).Op(":=").Index().Id(fieldDef.returnType).Values(),
				jen.For(
					jen.List(jen.Id("_"), jen.Id(singularVarName)).
						Op(":=").
						Range().
						Id(pluralVarName),
				).
					Block(
						jen.Id(outputVarName).Op("=").Append(
							jen.Id(outputVarName),
							jen.Id(fieldDef.returnType).Values(jen.Dict{
								jen.Id("Node"): jen.Id(singularVarName),
							}),
						),
					),
				jen.Return(jen.Id(outputVarName)),
			}
		} else {
			varName := "child"

			functionBody = []jen.Code{
				jen.Id(varName).
					Op(":=").
					Id(structMethodIdentifier).
					Dot("Node").
					Dot("ChildByFieldName").
					Call(jen.Lit(fieldDef.name)),
				jen.
					If(
						jen.Id(varName).Op("==").Nil(),
					).
					Block(
						jen.Return(
							jen.Id(fieldDef.returnType).Values(),
							jen.Qual("fmt", "Errorf").Call(
								jen.Lit("Node of kind %s has no "+varName+" of name %s"),
								jen.Lit(stDef.tsKind),
								jen.Lit(fieldDef.name),
							),
						),
					),
				jen.Return(
					jen.Id(fieldDef.returnType).Values(jen.Dict{
						jen.Id("Node"): jen.Op("*").Id("child"),
					}),
					jen.Nil(),
				),
			}
		}

		stmt := jen.Func().
			Parens(
				jen.Id(structMethodIdentifier).Op("*").Id(stDef.name),
			).
			Id(funcName).
			Parens(functionParams).
			Add(returnTypeStmt).
			Block(functionBody...)

		file.Add(stmt)
	}

	if stDef.childrenMethodDef == nil {
		return
	}

	returnTypeStmt := jen.Null()
	if stDef.childrenMethodDef.array {
		returnTypeStmt = returnTypeStmt.Index()
	}
	returnTypeStmt = returnTypeStmt.Id(stDef.childrenMethodDef.returnType)

	// Default return type includes an error, but for arrays, we just return an empty
	// array
	if !stDef.childrenMethodDef.array {
		returnTypeStmt = jen.Parens(jen.List(
			jen.List(
				returnTypeStmt,
				jen.Error(),
			),
		))
	}

	functionParams := jen.Null()
	var functionBody []jen.Code

	// Take in a tree cursor for iterating
	cursorVarName := "cursor"
	functionParams = jen.Id(cursorVarName).Op("*").Qual("github.com/tree-sitter/go-tree-sitter", "TreeCursor")

	singularVarName := "child"
	pluralVarName := "children"
	outputVarName := "output"

	functionBody = []jen.Code{
		jen.Id(pluralVarName).
			Op(":=").
			Id(structMethodIdentifier).
			Dot("Node").
			Dot("Children").
			Call(jen.Id(cursorVarName)),
		jen.Id(outputVarName).Op(":=").Index().Id(stDef.childrenMethodDef.returnType).Values(),
		jen.For(
			jen.List(jen.Id("_"), jen.Id(singularVarName)).
				Op(":=").
				Range().
				Id(pluralVarName),
		).
			Block(
				jen.If(jen.Id(singularVarName).Dot("IsNamed").Call()).Block(
					jen.Id(outputVarName).Op("=").Append(
						jen.Id(outputVarName),
						jen.Id(stDef.childrenMethodDef.returnType).Values(jen.Dict{
							jen.Id("Node"): jen.Id(singularVarName),
						}),
					),
				),
			),
	}

	if stDef.childrenMethodDef.array {
		functionBody = append(functionBody, jen.Return(jen.Id(outputVarName)))
	} else {
		functionBody = append(functionBody, []jen.Code{
			jen.If(jen.Len(jen.Id(outputVarName)).Op("==").Lit(0)).
				Block(
					jen.Return(
						jen.Id(stDef.childrenMethodDef.returnType).Values(),
						jen.Qual("fmt", "Errorf").Call(
							jen.Lit("No children found on node of kind %s"),
							jen.Lit(stDef.tsKind),
						),
					),
				),
			jen.Return(jen.Id(outputVarName).Index(jen.Lit(0)), jen.Nil()),
		}...)
	}

	file.Func().
		Parens(
			jen.Id(structMethodIdentifier).Op("*").Id(stDef.name),
		).
		Id(stDef.childrenMethodDef.name).
		Parens(functionParams).
		Add(returnTypeStmt).
		Block(functionBody...)
}

// createExportedName creates a name that is safe to use as an exported symbol name
// from a snake_case name. it also replaces sybols with word representations of
// those symbols.
func createExportedName(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Handle special case of the name '_'
	if s == "_" {
		return "Underscore"
	}

	s = strings.ToLower(s)
	s = strings.Trim(s, "_")
	words := strings.Split(s, "_")

	output := ""
	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		title := caser.String(word)

		for _, r := range title {
			if unicode.IsLetter(r) {
				output += string(r)
				continue
			}
			output += symbolToWord(r)
		}
	}

	return output
}

// createPrivateName creates a name that is safe to use as a private symbol name
// from a snake_case name. it also replaces sybols with word representations of
// those symbols.
func createPrivateName(s string) string {
	s = createExportedName(s)
	s = strings.ToLower(string(s[0])) + s[1:]
	if isReservedKeyword(s) {
		s += "_"
	}
	return s
}

// Thanks to https://github.com/Jakobeha/type-sitter/tree/303900dbb44af327e8731031f1781ad89d34f29c,
// which was MIT/Apache 2.0 at the time of writing on commit 303900d.
func symbolToWord(r rune) string {
	switch r {
	case '&':
		return "Ampersand"
	case '|':
		return "Bar"
	case '!':
		return "Not"
	case '=':
		return "Eq"
	case '<':
		return "Lt"
	case '>':
		return "Gt"
	case '+':
		return "Add"
	case '-':
		return "Sub"
	case '*':
		return "Mul"
	case '/':
		return "Div"
	case '~':
		return "BitNot"
	case '%':
		return "Mod"
	case '^':
		return "BitXor"
	case '?':
		return "Question"
	case ':':
		return "Colon"
	case '.':
		return "Dot"
	case ',':
		return "Comma"
	case ';':
		return "Semicolon"
	case '(':
		return "LParen"
	case ')':
		return "RParen"
	case '[':
		return "LBracket"
	case ']':
		return "RBracket"
	case '{':
		return "LBrace"
	case '}':
		return "RBrace"
	case '\\':
		return "Backslash"
	case '\'':
		return "Quote"
	case '"':
		return "DoubleQuote"
	case '#':
		return "Hash"
	case '@':
		return "At"
	case '$':
		return "Dollar"
	case '`':
		return "Backtick"
	case ' ':
		return "Space"
	case '\t':
		return "Tab"
	case '\n':
		return "Newline"
	case '\r':
		return "CarriageReturn"
	default:
		return fmt.Sprintf("U%X", r)
	}
}

func isReservedKeyword(s string) bool {
	switch s {
	case "break", "case", "chan", "const", "continue", "default",
		"defer", "else", "fallthrough", "for", "func", "go", "goto",
		"if", "import", "interface", "map", "package", "range", "return",
		"select", "struct", "switch", "type", "var":
		return true
	}
	return false
}

func getNodeMethodNames() []string {
	var node *tree_sitter.Node
	t := reflect.TypeOf(node)

	var names []string
	for i := 0; i < t.NumMethod(); i++ {
		names = append(names, t.Method(i).Name)
	}
	return names
}
