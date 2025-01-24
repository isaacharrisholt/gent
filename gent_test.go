package gent_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/isaacharrisholt/gent"
)

//go:embed testdata/python-node-types.json
var pythonNodeTypes []byte

func TestGenerator_Generate(t *testing.T) {
	gen := gent.NewGenerator(gent.GeneratorOptions{})
	res, err := gen.Generate(pythonNodeTypes)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fmt.Println(res)
}
