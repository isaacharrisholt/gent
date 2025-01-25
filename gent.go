package gent

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

type Generator struct {
	options GeneratorOptions
}

type GeneratorOptions struct {
	PackageName string
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
	unionTypes := make(map[string]unionType)

	// Collect all the union types in the file and register them in the slice
	// All supertypes are considered exported union types, and we create private
	// union types from `fields` and `children` of other types.
	for _, nodeType := range nodeTypes {
		if nodeType.Subtypes != nil {
			// This is a supertype, so we always export it and use the type name
			// as the name of the union type.
			unionType := unionType{
				name:    createExportedName(nodeType.Type),
				members: nodeType.Subtypes,
			}
			unionTypes[unionType.name] = unionType
		} else {
			// Check the fields and children of the type
			for name, children := range nodeType.Fields {
				if len(children.Types) > 1 {
					// This is a union type, so we create a private union type
					// from the types in the field.
					unionTypeName, err := generateUnionTypeName(children.Types)
					if err != nil {
						return "", fmt.Errorf("Failed to generate union type name for %s.%s: %w", nodeType.Type, name, err)
					}
					unionType := unionType{
						name:    unionTypeName,
						members: children.Types,
					}
					unionTypes[unionType.name] = unionType
				}
			}

			if len(nodeType.Children.Types) > 1 {
				unionTypeName, err := generateUnionTypeName(nodeType.Children.Types)
				if err != nil {
					return "", fmt.Errorf("Failed to generate union type name for %s: %w", nodeType.Type, err)
				}
				unionType := unionType{
					name:    unionTypeName,
					members: nodeType.Children.Types,
				}
				unionTypes[unionType.name] = unionType
			}
		}
	}

	fmt.Println(unionTypes)

	for _, nodeType := range nodeTypes {
		if nodeType.Subtypes == nil {
			err := addNodeType(file, nodeType)
			if err != nil {
				return "", fmt.Errorf("Failed to add node type %s: %w", nodeType.Type, err)
			}
		}
	}

	for _, unionType := range unionTypes {
		err := addUnionType(file, unionType)
		if err != nil {
			return "", fmt.Errorf("Failed to add union type %s: %w", unionType.name, err)
		}
	}

	outputBuilder := &strings.Builder{}
	err = file.Render(outputBuilder)
	if err != nil {
		return "", fmt.Errorf("Failed to render file: %w", err)
	}

	return outputBuilder.String(), nil
}

// generateUnionTypeName creates a combined name for a union type. This is always
// unexported.
func generateUnionTypeName(types []nodeChildType) (string, error) {
	if len(types) < 2 {
		return "", fmt.Errorf("Cannot create union type with less than 2 members")
	}

	typeNames := []string{}
	slices.SortFunc(types, func(a, b nodeChildType) int {
		return cmp.Compare(a.Type, b.Type)
	})

	for _, type_ := range types {
		typeNames = append(typeNames, createPrivateName(type_.Type))
	}
	return strings.Join(typeNames, "_"), nil
}

func addNodeType(file *jen.File, nodeType nodeType) error {
	fieldDefs := []structFieldDef{}

	for name, field := range nodeType.Fields {
		typeName := ""

		if len(field.Types) == 1 {
			typeName = createExportedName(field.Types[0].Type)
		} else {
			unionTypeName, err := generateUnionTypeName(field.Types)
			if err != nil {
				return fmt.Errorf("Failed to generate union type name for %s.%s: %w", nodeType.Type, name, err)
			}
			typeName = unionTypeName
		}

		if !field.Required {
			typeName = "*" + typeName
		}

		fieldDefs = append(fieldDefs, structFieldDef{
			name:     name,
			typeName: typeName,
		})
	}

	structFields := []jen.Code{}

	for _, fieldDef := range fieldDefs {
		stmt := jen.Id(createPrivateName(fieldDef.name)).Id(fieldDef.typeName)
		structFields = append(structFields, stmt)
	}

	structName := createExportedName(nodeType.Type)
	structMethodIdentifier := strings.ToLower(string(structName[0]))
	file.Type().Id(createExportedName(nodeType.Type)).Struct(structFields...)

	// TODO: pull struct writing out into a function
	for _, fieldDef := range fieldDefs {
		file.Func().
			Parens(
				jen.Id(structMethodIdentifier).Op("*").Id(structName),
			).
			Id(createExportedName(fieldDef.name)).
			Parens(jen.Null()).
			Id(fieldDef.typeName).
			Block(
				jen.Return(jen.Id(structMethodIdentifier).Dot(createPrivateName(fieldDef.name))),
			)
	}

	return nil
}

func addUnionType(file *jen.File, unionType unionType) error {
	// A union type is a struct containing fields for each of the types in the
	// union. These are all pointers to indicate that any of them could be nil.
	// The types of the fields should always be exported.
	fieldDefs := []structFieldDef{}
	for _, type_ := range unionType.members {
		fieldDefs = append(fieldDefs, structFieldDef{
			name:     createPrivateName(type_.Type),
			typeName: createExportedName(type_.Type),
		})
	}

	structFields := []jen.Code{}
	for _, fieldDef := range fieldDefs {
		stmt := jen.Id(fieldDef.name).Op("*").Id(fieldDef.typeName)
		structFields = append(structFields, stmt)
	}

	file.Type().Id(unionType.name).Struct(structFields...)

	structMethodIdentifier := strings.ToLower(string(unionType.name[0]))
	for _, fieldDef := range fieldDefs {
		file.Func().
			Parens(
				jen.Id(structMethodIdentifier).Op("*").Id(unionType.name),
			).
			Id(fieldDef.typeName).
			Parens(jen.Null()).
			Op("*").
			Id(fieldDef.typeName).
			Block(
				jen.Return(jen.Id(structMethodIdentifier).Dot(fieldDef.name)),
			)
	}

	return nil
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
		return "And"
	case '|':
		return "Or"
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
