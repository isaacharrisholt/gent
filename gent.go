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
	Type     string                  `json:"type"`
	Named    bool                    `json:"named"`
	Fields   map[string]nodeChildren `json:"fields"`
	Children nodeChildren            `json:"children"`
	Subtypes []nodeChildType         `json:"subtypes"`
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

type structFieldDef struct {
	name     string
	typeName string
	pointer  bool
	array    bool
}

type structDef struct {
	name string
	// Tree-sitter node kind
	tsKind string
	fields []structFieldDef
}

// Inner map keys are all Tree-sitter node names
type nodeMap struct {
	// Top-level node definitions, supertypes.
	// Values are struct names.
	namedExported map[string]string

	// Symbols, etc. that have `"named": false` but still need exporting.
	// Values are struct names.
	unnamedExported map[string]string

	// Supertypes. Values are `unionType` structs.
	supertypes map[string]unionType

	// Types made from a combination of other types, e.g. the multiple types possible in
	// `fields.types`.
	// Values are `unionType` structs.
	unionTypes map[string]unionType

	// Types that are present in `fields` or `children` but are not defined at the top
	// level. We always keep references to these and create an empty, private struct for
	// them.
	// Values are struct names.
	unknown map[string]string
}

func newNodeMap() nodeMap {
	return nodeMap{
		namedExported:   make(map[string]string),
		unnamedExported: make(map[string]string),
		supertypes:      make(map[string]unionType),
		unionTypes:      make(map[string]unionType),
		unknown:         make(map[string]string),
	}
}

func (nm *nodeMap) registerNodeType(nodeType nodeType) {
	structName := generateStructName(nodeType.Type, nodeType.Named)
	if nodeType.Named {
		nm.namedExported[nodeType.Type] = structName
		return
	}

	nm.unnamedExported[nodeType.Type] = structName
}

func (nm *nodeMap) getStructName(typeName string, named bool) (string, bool) {
	// Check relevant named maps, then check unknown types
	if named {
		if structName, ok := nm.namedExported[typeName]; ok {
			return structName, true
		}
		if structName, ok := nm.supertypes[typeName]; ok {
			return structName.name, true
		}
	}

	if !named {
		if structName, ok := nm.unnamedExported[typeName]; ok {
			return structName, true
		}
	}

	structName, ok := nm.unknown[typeName]
	return structName, ok
}

func (nm *nodeMap) registerSupertype(typeName string, members []nodeChildType) {
	structName := generateStructName(typeName, true)
	nm.supertypes[typeName] = unionType{name: structName, members: members}
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

	nm.unionTypes[tsNodeTypeName] = unionType{
		name:    structTypeName,
		members: members,
	}

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

	ut, ok := nm.unionTypes[tsNodeTypeName]
	if !ok {
		return unionType{}, false
	}

	return ut, true
}

func (nm *nodeMap) registerUnknownType(typeName string) {
	structName := "Unknown__" + createPrivateName(typeName)
	nm.unknown[typeName] = structName
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
			for name, children := range nodeType.Fields {
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
		for _, field := range nodeType.Fields {
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
	for tsKindName, structName := range nm.namedExported {
		addPublicType(structName, tsKindName)
	}
	if g.options.Debug {
		publicTypes = append(publicTypes, file.Comment("Unnamed types"))
	}
	for tsKindName, structName := range nm.unnamedExported {
		addPublicType(structName, tsKindName)
	}
	if g.options.Debug {
		publicTypes = append(publicTypes, file.Comment("Supertypes"))
	}
	for tsKindName, unionType := range nm.supertypes {
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
	for _, supertype := range nm.supertypes {
		err := addUnionType(file, supertype, &nm)
		if err != nil {
			return "", fmt.Errorf("Failed to add supertype %s: %w", supertype.name, err)
		}
	}

	if g.options.Debug {
		file.Comment("\nUNION TYPES\n")
	}
	for _, unionType := range nm.unionTypes {
		err := addUnionType(file, unionType, &nm)
		if err != nil {
			return "", fmt.Errorf("Failed to add union type %s: %w", unionType.name, err)
		}
	}

	// Add empty structs for the unknown types. They can be private.
	if g.options.Debug {
		file.Comment("\nUNKNOWN TYPES\n")
	}
	for tsKind, unknownType := range nm.unknown {
		writeStruct(file, structDef{
			name:   unknownType,
			tsKind: tsKind,
			fields: []structFieldDef{},
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
	fieldDefs := []structFieldDef{}

	for name, field := range nodeType.Fields {
		typeName := ""

		if len(field.Types) == 1 {
			fieldType := field.Types[0]
			structName, ok := nm.getStructName(fieldType.Type, fieldType.Named)
			if !ok {
				return fmt.Errorf("Failed to find TS node %s in map", field.Types[0].Type)
			}
			typeName = structName
		} else {
			unionType, ok := nm.getUnionType(field.Types)
			if !ok {
				return fmt.Errorf("Failed to find union type for %s.%s types", nodeType.Type, name)
			}
			typeName = unionType.name
		}

		fieldDefs = append(fieldDefs, structFieldDef{
			name:     createPrivateName(name),
			typeName: typeName,
			pointer:  !field.Required,
			array:    field.Multiple,
		})
	}

	structName, ok := nm.getStructName(nodeType.Type, nodeType.Named)
	if !ok {
		return fmt.Errorf("Failed to find struct name for %s", nodeType.Type)
	}

	writeStruct(file, structDef{
		name:   structName,
		tsKind: nodeType.Type,
		fields: fieldDefs,
	})

	return nil
}

func addUnionType(file *jen.File, unionType unionType, nm *nodeMap) error {
	// A union type is a struct containing fields for each of the types in the
	// union. These are all pointers to indicate that any of them could be nil.
	// The types of the fields should always be exported.
	fieldDefs := []structFieldDef{}
	for _, member := range unionType.members {
		typeName, ok := nm.getStructName(member.Type, member.Named)
		if !ok {
			return fmt.Errorf("Failed to find struct name for %s.%s", unionType.name, member.Type)
		}
		fieldDefs = append(fieldDefs, structFieldDef{
			name:     createPrivateName(member.Type),
			typeName: typeName,
			pointer:  true,
			array:    false,
		})
	}

	writeStruct(file, structDef{
		name:   unionType.name,
		fields: fieldDefs,
		// TODO: TS kind? Or separate validator logic into new functions, most likely
	})

	return nil
}

func writeStruct(file *jen.File, def structDef) {
	embedField := jen.Qual("github.com/tree-sitter/go-tree-sitter", "Node")
	structFields := []jen.Code{embedField}

	for _, fieldDef := range def.fields {
		stmt := jen.Id(fieldDef.name)
		if fieldDef.pointer {
			stmt = stmt.Op("*")
		}
		if fieldDef.array {
			stmt = stmt.Index()
		}
		stmt.Id(fieldDef.typeName)
		structFields = append(structFields, stmt)
	}

	file.Type().Id(def.name).Struct(structFields...)

	structMethodIdentifier := strings.ToLower(string(def.name[0]))
	for _, fieldDef := range def.fields {
		funcName := strings.ToUpper(string(fieldDef.name[0])) + fieldDef.name[1:]

		// Need to make sure we don't override any of the reserved node methods,
		// so prefix with 'Get' until the name is unique.
		for slices.Contains(reservedNodeMethods, funcName) {
			funcName = "Get" + funcName
		}

		stmt := jen.Func().
			Parens(
				jen.Id(structMethodIdentifier).Op("*").Id(def.name),
			).
			Id(funcName).
			Parens(jen.Null())

		if fieldDef.pointer {
			stmt = stmt.Op("*")
		}
		if fieldDef.array {
			stmt = stmt.Index()
		}

		stmt = stmt.
			Id(fieldDef.typeName).
			Block(
				jen.Return(jen.Id(structMethodIdentifier).Dot(fieldDef.name)),
			)

		file.Add(stmt)
	}

	file.Func().
		Parens(
			jen.Id(structMethodIdentifier).Op("*").Id(def.name),
		).
		Id("Validate").
		Parens(jen.Id("node").Qual("github.com/tree-sitter/go-tree-sitter", "Node")).
		Parens(
			jen.List(
				jen.Id(def.name),
				jen.Error(),
			),
		).
		Block(
			// TODO: add validation logic
			jen.
				If(
					jen.Id("node").Dot("Kind").Call().Op("!=").Id(def.tsKind),
				).
				Block(
					jen.Return(
						jen.Id(def.name).Values(),
						jen.Qual("fmt", "Errorf").Call(
							jen.Lit("Expected node of kind "+strings.ReplaceAll(def.tsKind, "%", "%%")+", got %s"),
							jen.Id("node").Dot("Kind").Call(),
						),
					),
				),
			jen.Return(
				jen.Id(def.name).Values(jen.Dict{
					jen.Id("Node"): jen.Id("node"),
				}),
				jen.Nil(),
			),
		)
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
