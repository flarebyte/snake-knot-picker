// purpose: Provide assertion helpers for tests to validate structured error behavior consistently.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package testutil

import (
	"github.com/flarebyte/snake-knot-picker"
)

type fatalHelper interface {
	Helper()
	Fatalf(format string, args ...any)
	Fatal(args ...any)
}

// MustNoError fails the test helper when err is non-nil.
func MustNoError(t fatalHelper, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// MustValidationErrorWithID fails unless err is a ValidationError whose first detail matches id.
func MustValidationErrorWithID(t fatalHelper, err error, id string) *picker.ValidationError {
	t.Helper()
	verr, ok := err.(*picker.ValidationError)
	if !ok {
		t.Fatalf("expected *picker.ValidationError, got %T", err)
	}
	if len(verr.Details) == 0 {
		t.Fatal("expected at least one error detail")
	}
	if verr.Details[0].ID != id {
		t.Fatalf("expected first error id %q, got %q", id, verr.Details[0].ID)
	}
	return verr
}

// MustEqualPath fails unless two string slices match exactly in length and values.
func MustEqualPath(t fatalHelper, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("path length mismatch: got=%v want=%v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("path mismatch: got=%v want=%v", got, want)
		}
	}
}
