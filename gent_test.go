package gent_test

import (
	_ "embed"
	"testing"

	"github.com/isaacharrisholt/gent"
	python "github.com/isaacharrisholt/gent/testdata"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_python "github.com/tree-sitter/tree-sitter-python/bindings/go"
)

//go:embed testdata/python-node-types.json
var pythonNodeTypes []byte

//go:embed testdata/test_program.py
var testPythonProgram []byte

func TestGenerator_Generate(t *testing.T) {
	gen := gent.NewGenerator(gent.GeneratorOptions{
		PackageName: "python_nodes",
	})
	_, err := gen.Generate(pythonNodeTypes)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func getPythonNodeText(node *tree_sitter.Node) string {
	start, end := node.ByteRange()
	return string(testPythonProgram[start:end])
}

func TestPythonParse(t *testing.T) {
	tsPython := tree_sitter_python.Language()
	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tsPython))

	tree := parser.Parse(testPythonProgram, nil)
	defer tree.Close()

	cursor := tree.Walk()
	defer cursor.Close()

	module, err := python.NewModule(tree.RootNode())
	if err != nil {
		t.Fatalf("Failed to create gent node: %v", err)
	}
	topLevelStatements := module.TypedChildren(cursor)

	expressionStatement0 := topLevelStatements[0]
	if expressionStatement0.Kind() != python.SyntaxKind_ImportStatement {
		t.Fatalf("Expected first child to be an import statement, got %v", expressionStatement0.Kind())
	}
	// Validate that we can get to the import statement in a type-safe way
	importSimpleStatement, err := expressionStatement0.SimpleStatement()
	if err != nil {
		t.Fatalf("Failed to get simple statement: %v", err)
	}

	importStatement, err := importSimpleStatement.ImportStatement()
	if err != nil {
		t.Fatalf("Failed to get import statement: %v", err)
	}

	importStatementNames := importStatement.Name(cursor)
	if len(importStatementNames) != 1 {
		t.Fatalf("Expected 1 import statement name, got %v", len(importStatementNames))
	}

	if name := getPythonNodeText(&importStatementNames[0].Node); name != "sys" {
		t.Fatalf("Expected import statement name to be sys, got %s", name)
	}
}
