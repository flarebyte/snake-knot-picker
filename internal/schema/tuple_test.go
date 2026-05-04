package schema

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestCompileTupleSpecValidOutOfOrderSlotsWithRepeatable(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	primary := mustCompile(t, c, []string{"schema", "tuple", "--size", "2", "--required"})
	slot1 := mustCompile(t, c, []string{"schema", "number", "--tuple", "1", "--int"})
	slot0 := mustCompile(t, c, []string{"schema", "string", "--tuple", "0", "--alphabetic"})
	repeatable := mustCompile(t, c, []string{"schema", "repeatable", "--min-length", "1", "--max-length", "5"})

	spec, err := CompileTupleSpec(primary, []*CompiledSpec{slot1, slot0, repeatable}, "range")
	if err != nil {
		t.Fatalf("unexpected tuple compile error: %v", err)
	}
	if spec.Size != 2 || !spec.Required {
		t.Fatalf("unexpected tuple metadata: %#v", spec)
	}
	if spec.Repeatable == nil {
		t.Fatal("expected repeatable modifier")
	}
	if spec.Slots[0] == nil || spec.Slots[1] == nil {
		t.Fatalf("expected two slots: %#v", spec.Slots)
	}
}

func TestCompileTupleSpecErrors(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	primary := mustCompile(t, c, []string{"schema", "tuple", "--size", "2"})
	slot0 := mustCompile(t, c, []string{"schema", "number", "--tuple", "0", "--int"})
	slot1 := mustCompile(t, c, []string{"schema", "number", "--tuple", "1", "--int"})
	slot2 := mustCompile(t, c, []string{"schema", "number", "--tuple", "2", "--int"})
	slotNoTuple := mustCompile(t, c, []string{"schema", "number", "--int"})
	repeatableA := mustCompile(t, c, []string{"schema", "repeatable", "--min-length", "1", "--max-length", "2"})
	repeatableB := mustCompile(t, c, []string{"schema", "repeatable", "--min-length", "1", "--max-length", "3"})

	cases := []struct {
		name     string
		children []*CompiledSpec
		id       string
	}{
		{name: "missing-index", children: []*CompiledSpec{slotNoTuple, slot1}, id: picker.ErrorIDSchemaTupleMissingIndex},
		{name: "out-of-range", children: []*CompiledSpec{slot0, slot2}, id: picker.ErrorIDSchemaTupleIndexOutOfRange},
		{name: "duplicate-slot", children: []*CompiledSpec{slot0, slot0.Clone()}, id: picker.ErrorIDSchemaTupleDuplicateSlot},
		{name: "missing-slot", children: []*CompiledSpec{slot0}, id: picker.ErrorIDSchemaTupleMissingSlot},
		{name: "duplicate-repeatable", children: []*CompiledSpec{slot0, slot1, repeatableA, repeatableB}, id: picker.ErrorIDSchemaInvalidCombination},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := CompileTupleSpec(primary, tc.children, "range")
			assertErrorID(t, err, tc.id)
		})
	}
}

func TestOldSingleCommandMultiTupleShapeRejected(t *testing.T) {
	_, err := ParseTokens([]string{"schema", "number", "--tuple", "0", "--tuple", "1", "--int"})
	assertErrorID(t, err, picker.ErrorIDSchemaInvalidValue)
}

func mustCompile(t *testing.T, c *Compiler, tokens []string) *CompiledSpec {
	t.Helper()
	ast, err := ParseTokens(tokens)
	if err != nil {
		t.Fatalf("parse failed %v: %v", tokens, err)
	}
	spec, err := c.Compile(ast)
	if err != nil {
		t.Fatalf("compile failed %v: %v", tokens, err)
	}
	return spec
}

