package schema

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestValidateCompileInputsBranches(t *testing.T) {
	c := NewCompiler(picker.NewRegistry())
	if err := c.validateCompileInputs(nil); err == nil {
		t.Fatal("expected nil ast error")
	}

	cNoRegistry := NewCompiler(nil)
	ast, _ := ParseTokens([]string{"schema", "string"})
	if err := cNoRegistry.validateCompileInputs(ast); err == nil {
		t.Fatal("expected nil registry error")
	}
}

func TestSplitEnumEmptySeparator(t *testing.T) {
	got := splitEnum("a,b", "")
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("unexpected split result: %#v", got)
	}
}

func TestCompiledSpecCloneNilReceiver(t *testing.T) {
	var spec *CompiledSpec
	if spec.Clone() != nil {
		t.Fatal("expected nil clone for nil receiver")
	}
}

func TestMappingNilAndParseFailures(t *testing.T) {
	sopts, serr := StringOptionsFromSpec(nil)
	if serr != nil {
		t.Fatalf("unexpected nil spec string error: %v", serr)
	}
	if sopts.Enum != nil {
		t.Fatalf("unexpected enum for nil spec: %#v", sopts)
	}

	nopts, nerr := NumberOptionsFromSpec(nil)
	if nerr != nil {
		t.Fatalf("unexpected nil spec number error: %v", nerr)
	}
	if nopts.Min != nil || nopts.Max != nil || nopts.MultipleOf != nil {
		t.Fatalf("unexpected number opts for nil spec: %#v", nopts)
	}

	_, errEnum := StringOptionsFromSpec(&CompiledSpec{
		Flags: map[string][]string{
			"--enum":           {"a, b"},
			"--enum-separator": {","},
		},
	})
	if errEnum == nil {
		t.Fatal("expected enum parse error")
	}

	_, errMin := NumberOptionsFromSpec(&CompiledSpec{Flags: map[string][]string{"--min": {"x"}}})
	if errMin == nil {
		t.Fatal("expected invalid --min parse error")
	}
	_, errMax := NumberOptionsFromSpec(&CompiledSpec{Flags: map[string][]string{"--max": {"x"}}})
	if errMax == nil {
		t.Fatal("expected invalid --max parse error")
	}
	_, errMultiple := NumberOptionsFromSpec(&CompiledSpec{Flags: map[string][]string{"--multiple-of": {"x"}}})
	if errMultiple == nil {
		t.Fatal("expected invalid --multiple-of parse error")
	}
}

func TestParseTokensEmptyOperatorAndUnknownFlagModes(t *testing.T) {
	_, errEmptyOp := ParseTokens([]string{"schema", ""})
	if errEmptyOp == nil {
		t.Fatal("expected empty operator error")
	}

	ast, errUnknownNoValue := ParseTokens([]string{"schema", "string", "--x"})
	if errUnknownNoValue != nil {
		t.Fatalf("unexpected unknown no-value parse error: %v", errUnknownNoValue)
	}
	if len(ast.Flags) != 1 || ast.Flags[0].Name != "--x" || len(ast.Flags[0].Values) != 0 {
		t.Fatalf("unexpected ast for unknown no-value flag: %#v", ast)
	}

	ast, errUnknownWithValues := ParseTokens([]string{"schema", "string", "--x", "a", "b", "--required"})
	if errUnknownWithValues != nil {
		t.Fatalf("unexpected unknown with-values parse error: %v", errUnknownWithValues)
	}
	if len(findFlagValues(ast.Flags, "--x")) != 2 {
		t.Fatalf("expected two unknown flag values: %#v", ast.Flags)
	}
}

func TestCompileTupleSpecGuardBranches(t *testing.T) {
	_, errPrimaryNil := CompileTupleSpec(nil, nil, "range")
	assertErrorID(t, errPrimaryNil, picker.ErrorIDSchemaInvalidValue)

	_, errWrongPrimary := CompileTupleSpec(&CompiledSpec{Operator: "string"}, nil, "range")
	assertErrorID(t, errWrongPrimary, picker.ErrorIDSchemaInvalidValue)

	_, errMissingSize := CompileTupleSpec(&CompiledSpec{Operator: "tuple", Flags: map[string][]string{}}, nil, "range")
	assertErrorID(t, errMissingSize, picker.ErrorIDSchemaInvalidValue)

	_, errInvalidSize := CompileTupleSpec(&CompiledSpec{
		Operator: "tuple",
		Flags:    map[string][]string{"--size": {"x"}},
	}, nil, "range")
	assertErrorID(t, errInvalidSize, picker.ErrorIDSchemaInvalidValue)

	_, errNilChild := CompileTupleSpec(&CompiledSpec{
		Operator: "tuple",
		Flags:    map[string][]string{"--size": {"1"}},
	}, []*CompiledSpec{nil}, "range")
	assertErrorID(t, errNilChild, picker.ErrorIDSchemaInvalidValue)
}
