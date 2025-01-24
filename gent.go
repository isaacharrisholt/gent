package gent

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

func (g *Generator) Generate(data []byte) (string, error) {
	var nodeTypes nodeTypes
	err := json.Unmarshal(data, &nodeTypes)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}

	nodeTypeMap := make(map[string]nodeType)
	for _, nodeType := range nodeTypes {
		nodeTypeMap[nodeType.Type] = nodeType
		fmt.Println(len(nodeType.Subtypes))
	}

	fmt.Printf("%#v\n", nodeTypeMap)

	packageName := "node_types"
	if g.options.PackageName != "" {
		packageName = g.options.PackageName
	}

	file := jen.NewFile(packageName)

	for _, nodeType := range nodeTypes {
		if nodeType.Subtypes == nil {
			err := addNodeType(file, nodeType)
			if err != nil {
				return "", fmt.Errorf("Failed to add node type %s: %w", nodeType.Type, err)
			}
		}
	}

	outputBuilder := &strings.Builder{}
	err = file.Render(outputBuilder)
	if err != nil {
		return "", fmt.Errorf("Failed to render file: %w", err)
	}

	return outputBuilder.String(), nil
}

func addSupertype(_ *jen.File, nodeType nodeType) error {
	if nodeType.Subtypes == nil {
		return fmt.Errorf("Node %s does not have any subtypes and is not a supertype", nodeType.Type)
	}

	return fmt.Errorf("Supertypes not implemented")
}

func addNodeType(file *jen.File, nodeType nodeType) error {
	file.Type().Id(createExportedName(nodeType.Type)).Struct()

	return nil
}

var caser = cases.Title(language.English)

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
	return strings.ToLower(string(s[0])) + s[1:]
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
