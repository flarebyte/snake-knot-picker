package schema

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestStringOptionsFromSpec(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "string", "--enum", "red;green;blue", "--enum-separator", ";", "--alphabetic", "--starts-with", "re"})
	spec, err := c.Compile(ast)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	opts, err := StringOptionsFromSpec(spec)
	if err != nil {
		t.Fatalf("mapping failed: %v", err)
	}
	if len(opts.Enum) != 3 || opts.Enum[0] != "red" {
		t.Fatalf("unexpected enum options: %#v", opts.Enum)
	}
	if !opts.Alphabetic || opts.StartsWith != "re" {
		t.Fatalf("unexpected mapped options: %#v", opts)
	}
}

func TestNumberOptionsFromSpec(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	ast, _ := ParseTokens([]string{"schema", "number", "--int", "--min", "1", "--max", "10", "--multiple-of", "0.5"})
	spec, err := c.Compile(ast)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	opts, err := NumberOptionsFromSpec(spec)
	if err != nil {
		t.Fatalf("mapping failed: %v", err)
	}
	if !opts.Int || opts.Min == nil || opts.Max == nil || opts.MultipleOf == nil {
		t.Fatalf("unexpected number options: %#v", opts)
	}
}

