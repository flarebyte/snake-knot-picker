// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package testutil

import (
	"errors"
	"fmt"
	"testing"

	picker "github.com/flarebyte/snake-knot-picker"
)

type stubT struct {
	fatalfCount int
	lastFatalf  string
}

func (s *stubT) Helper() {}

func (s *stubT) Fatalf(format string, args ...any) {
	s.fatalfCount++
	s.lastFatalf = fmt.Sprintf(format, args...)
	panic("fatalf-called")
}

func (s *stubT) Fatal(args ...any) {
	s.fatalfCount++
	s.lastFatalf = fmt.Sprint(args...)
	panic("fatal-called")
}

func TestMustNoError(t *testing.T) {
	s := &stubT{}
	MustNoError(s, nil)
	if s.fatalfCount != 0 {
		t.Fatalf("unexpected fatal count for nil error: %d", s.fatalfCount)
	}

	mustPanic(t, func() { MustNoError(s, errors.New("boom")) })
	if s.fatalfCount != 1 {
		t.Fatalf("expected one fatal call for non-nil error, got %d", s.fatalfCount)
	}
}

func TestMustValidationErrorWithID(t *testing.T) {
	s := &stubT{}
	verr := picker.NewValidationError(picker.ErrorIDValidationInvalidType, nil)
	got := MustValidationErrorWithID(s, verr, picker.ErrorIDValidationInvalidType)
	if got == nil || s.fatalfCount != 0 {
		t.Fatalf("expected success path, got=%#v fatalf=%d", got, s.fatalfCount)
	}

	mustPanic(t, func() { _ = MustValidationErrorWithID(s, errors.New("boom"), picker.ErrorIDValidationInvalidType) })
	if s.fatalfCount != 1 {
		t.Fatalf("expected fatal on wrong error type, got %d", s.fatalfCount)
	}

	mustPanic(t, func() {
		_ = MustValidationErrorWithID(s, &picker.ValidationError{}, picker.ErrorIDValidationInvalidType)
	})
	if s.fatalfCount != 2 {
		t.Fatalf("expected fatal on empty details, got %d", s.fatalfCount)
	}

	mustPanic(t, func() {
		_ = MustValidationErrorWithID(s, picker.NewValidationError(picker.ErrorIDValidationTuple, nil), picker.ErrorIDValidationInvalidType)
	})
	if s.fatalfCount != 3 {
		t.Fatalf("expected fatal on wrong detail id, got %d", s.fatalfCount)
	}
}

func TestMustEqualPath(t *testing.T) {
	s := &stubT{}
	MustEqualPath(s, []string{"wash", "start"}, []string{"wash", "start"})
	if s.fatalfCount != 0 {
		t.Fatalf("unexpected fatal count on equal path: %d", s.fatalfCount)
	}

	mustPanic(t, func() { MustEqualPath(s, []string{"wash"}, []string{"wash", "start"}) })
	if s.fatalfCount != 1 {
		t.Fatalf("expected fatal on path length mismatch, got %d", s.fatalfCount)
	}

	mustPanic(t, func() { MustEqualPath(s, []string{"wash", "stop"}, []string{"wash", "start"}) })
	if s.fatalfCount != 2 {
		t.Fatalf("expected fatal on path value mismatch, got %d", s.fatalfCount)
	}
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic from stub fatalf")
		}
	}()
	fn()
}
