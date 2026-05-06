package picker

import "testing"

func mustValidationError(t *testing.T, err error) *ValidationError {
	t.Helper()
	verr, ok := err.(*ValidationError)
	if !ok {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if len(verr.Details) == 0 {
		t.Fatal("expected at least one error detail")
	}
	return verr
}

func assertErrorDetail(t *testing.T, err error, wantID, wantKind string) {
	t.Helper()
	verr := mustValidationError(t, err)
	if verr.Details[0].ID != wantID {
		t.Fatalf("unexpected error id: got=%s want=%s", verr.Details[0].ID, wantID)
	}
	if verr.Details[0].Kind != wantKind {
		t.Fatalf("unexpected error kind: got=%s want=%s", verr.Details[0].Kind, wantKind)
	}
	if verr.Details[0].Message == "" {
		t.Fatalf("expected non-empty error message for id=%s", wantID)
	}
}

func assertStringValue(t *testing.T, got map[string]Value, key, want string) {
	t.Helper()
	v, ok := got[key]
	if !ok || v.String == nil {
		t.Fatalf("expected string value for %q, got=%#v", key, v)
	}
	if *v.String != want {
		t.Fatalf("unexpected string value for %q: got=%q want=%q", key, *v.String, want)
	}
}

func assertNumberValue(t *testing.T, got map[string]Value, key string, want float64) {
	t.Helper()
	v, ok := got[key]
	if !ok || v.Number == nil {
		t.Fatalf("expected number value for %q, got=%#v", key, v)
	}
	if *v.Number != want {
		t.Fatalf("unexpected number value for %q: got=%v want=%v", key, *v.Number, want)
	}
}

func assertBoolValue(t *testing.T, got map[string]Value, key string, want bool) {
	t.Helper()
	v, ok := got[key]
	if !ok || v.Bool == nil {
		t.Fatalf("expected bool value for %q, got=%#v", key, v)
	}
	if *v.Bool != want {
		t.Fatalf("unexpected bool value for %q: got=%v want=%v", key, *v.Bool, want)
	}
}

func assertTupleStringValues(t *testing.T, got map[string]Value, key string, want ...string) {
	t.Helper()
	v, ok := got[key]
	if !ok {
		t.Fatalf("expected tuple value for %q", key)
	}
	if len(v.Tuple) != len(want) {
		t.Fatalf("unexpected tuple len for %q: got=%d want=%d", key, len(v.Tuple), len(want))
	}
	for i, expected := range want {
		if v.Tuple[i].String == nil {
			t.Fatalf("tuple[%d] for %q is not string: %#v", i, key, v.Tuple[i])
		}
		if *v.Tuple[i].String != expected {
			t.Fatalf("unexpected tuple[%d] for %q: got=%q want=%q", i, key, *v.Tuple[i].String, expected)
		}
	}
}
