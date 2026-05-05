package schema

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestCompileArgsCommandFixtures(t *testing.T) {
	raw, err := os.ReadFile("../../doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var doc struct {
		Flags []struct {
			Schema  []string   `json:"schema"`
			Schemas [][]string `json:"schemas"`
		} `json:"flags"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("decode fixture: %v", err)
	}
	c := NewCompiler(picker.NewRegistry())
	for i, f := range doc.Flags {
		ast, err := ParseTokens(f.Schema)
		if err != nil {
			t.Fatalf("parse fixture flag[%d]: %v", i, err)
		}
		if _, err := c.Compile(ast); err != nil {
			t.Fatalf("compile fixture flag[%d]: %v", i, err)
		}
		for j, child := range f.Schemas {
			childAst, err := ParseTokens(child)
			if err != nil {
				t.Fatalf("parse fixture flag[%d].schemas[%d]: %v", i, j, err)
			}
			if _, err := c.Compile(childAst); err != nil {
				t.Fatalf("compile fixture flag[%d].schemas[%d]: %v", i, j, err)
			}
		}
	}
}

func TestCompileUnknownOperator(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "frobnicate", "--required"})
	_, err := c.Compile(ast)
	assertErrorID(t, err, picker.ErrorIDSchemaUnknownOperator)
}

func TestCompileUnknownFlag(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "string", "--bogus"})
	_, err := c.Compile(ast)
	assertErrorID(t, err, picker.ErrorIDSchemaUnknownFlag)
}

func TestCompileInvalidCombination(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "string", "--uri", "--secure", "--scheme", "http"})
	_, err := c.Compile(ast)
	assertErrorID(t, err, picker.ErrorIDSchemaInvalidCombination)
}

func TestCompileEnumInvalidDefinitions(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())

	astWhitespace, _ := ParseTokens([]string{"schema", "string", "--enum", "cold, warm,hot"})
	_, err := c.Compile(astWhitespace)
	assertErrorID(t, err, picker.ErrorIDSchemaEnumWhitespace)

	astEmpty, _ := ParseTokens([]string{"schema", "string", "--enum", "cold,,hot"})
	_, err = c.Compile(astEmpty)
	assertErrorID(t, err, picker.ErrorIDSchemaEnumEmpty)
}

func TestCompileSpecIsMutationSafe(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "string", "--enum", "cold,warm,hot"})
	spec, err := c.Compile(ast)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	cloned := spec.Clone()
	spec.Flags["--enum"][0] = "MUTATED"
	spec.Raw[0] = "MUTATED"

	if cloned.Flags["--enum"][0] != "cold,warm,hot" {
		t.Fatalf("clone should retain original flag values: %#v", cloned.Flags["--enum"])
	}
	if cloned.Raw[0] != "schema" {
		t.Fatalf("clone should retain original raw tokens: %#v", cloned.Raw)
	}
}

func TestRegistryDuplicateRegistration(t *testing.T) {
	reg := picker.NewRegistry()
	err := reg.Register(picker.NewStaticFactory("string"))
	assertErrorID(t, err, picker.ErrorIDSchemaDuplicateRegistration)
}

func assertErrorID(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error %s", want)
	}
	verr, ok := err.(*picker.ValidationError)
	if !ok || len(verr.Details) == 0 {
		t.Fatalf("expected structured error, got %T", err)
	}
	if verr.Details[0].ID != want {
		t.Fatalf("unexpected id: got=%s want=%s", verr.Details[0].ID, want)
	}
}
