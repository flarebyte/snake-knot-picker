// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import "testing"

func TestParseFlagValueBranches(t *testing.T) {
	// boolean: default true (len(rawValues)==0 branch)
	bv, err := parseFlagValue(CompiledFlag{Kind: "boolean", Name: "enabled"}, nil)
	if err != nil || bv.Bool == nil || *bv.Bool != true {
		t.Fatalf("unexpected boolean default parse: value=%#v err=%v", bv, err)
	}

	// boolean parse error
	_, err = parseFlagValue(CompiledFlag{Kind: "boolean", Name: "enabled"}, []string{"not-bool"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)

	// number parse error
	_, err = parseFlagValue(CompiledFlag{Kind: "number", Name: "spin"}, []string{"NaNx"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)

	// tuple invalid flag-like segment
	_, err = parseFlagValue(CompiledFlag{Kind: "tuple", Name: "range", TupleSize: 2}, []string{"1,--x"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)

	// tuple arity mismatch
	_, err = parseFlagValue(CompiledFlag{Kind: "tuple", Name: "range", TupleSize: 2}, []string{"1"})
	assertErrorDetail(t, err, ErrorIDValidationTuple, ErrorKindValidation)

	// repeatable invalid flag-like segment
	_, err = parseFlagValue(CompiledFlag{Kind: "string", Name: "add", Repeatable: true}, []string{"soap,--x"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)

	// top-level flag-like value guard
	_, err = parseFlagValue(CompiledFlag{Kind: "string", Name: "mode"}, []string{" --x"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
}

func TestFlattenListBranches(t *testing.T) {
	existing := []Value{{String: strPtr("a")}, {String: strPtr("b")}}
	gotExisting := flattenList(Value{List: existing})
	if len(gotExisting) != 2 || gotExisting[0].String == nil || *gotExisting[0].String != "a" {
		t.Fatalf("unexpected flatten existing list: %#v", gotExisting)
	}

	gotSingle := flattenList(Value{String: strPtr("x")})
	if len(gotSingle) != 1 || gotSingle[0].String == nil || *gotSingle[0].String != "x" {
		t.Fatalf("unexpected flatten single fallback: %#v", gotSingle)
	}
}

func strPtr(s string) *string { return &s }
