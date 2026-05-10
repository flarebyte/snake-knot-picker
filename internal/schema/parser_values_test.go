// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestParseFlagValuesDirectBranches(t *testing.T) {
	values, next, err := parseFlagValues([]string{"schema", "string", "--required"}, 2, "--required")
	if err != nil {
		t.Fatalf("unexpected arity-0 error: %v", err)
	}
	if len(values) != 0 || next != 3 {
		t.Fatalf("unexpected arity-0 parse output: values=%#v next=%d", values, next)
	}

	_, _, errKnownMissing := parseFlagValues([]string{"schema", "string", "--enum", "--required"}, 2, "--enum")
	if errKnownMissing == nil {
		t.Fatal("expected known-arity missing value error")
	}
	verr, ok := errKnownMissing.(*picker.ValidationError)
	if !ok || len(verr.Details) == 0 || verr.Details[0].ID != picker.ErrorIDSchemaMissingValue {
		t.Fatalf("unexpected missing-value error: %#v", errKnownMissing)
	}

	values, next, err = parseFlagValues([]string{"schema", "string", "--unknown", "x", "y", "--required"}, 2, "--unknown")
	if err != nil {
		t.Fatalf("unexpected unknown-flag values error: %v", err)
	}
	if next != 5 || len(values) != 2 || values[0] != "x" || values[1] != "y" {
		t.Fatalf("unexpected unknown-flag parse output: values=%#v next=%d", values, next)
	}

	values, next, err = parseFlagValues([]string{"schema", "string", "--unknown", "--required"}, 2, "--unknown")
	if err != nil {
		t.Fatalf("unexpected unknown-flag empty error: %v", err)
	}
	if next != 3 || len(values) != 0 {
		t.Fatalf("unexpected unknown-flag empty output: values=%#v next=%d", values, next)
	}
}
