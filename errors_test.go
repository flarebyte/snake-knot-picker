package picker

import (
	"encoding/json"
	"os"
	"testing"
)

type errorFixture struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

func TestErrorFixturesCoverEveryKnownID(t *testing.T) {
	raw, err := os.ReadFile("testdata/fixtures/errors/error-catalog.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var fixtures []errorFixture
	if err := json.Unmarshal(raw, &fixtures); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}

	got := map[string]string{}
	for _, f := range fixtures {
		got[f.ID] = f.Kind
	}

	want := map[string]string{
		ErrorIDSchemaUnknownOperator:            ErrorKindSchema,
		ErrorIDSchemaUnknownFlag:                ErrorKindSchema,
		ErrorIDSchemaMissingValue:               ErrorKindSchema,
		ErrorIDSchemaInvalidValue:               ErrorKindSchema,
		ErrorIDSchemaInvalidCombination:         ErrorKindSchema,
		ErrorIDSchemaDuplicateRegistration:      ErrorKindSchema,
		ErrorIDSchemaEnumWhitespace:             ErrorKindSchema,
		ErrorIDSchemaEnumEmpty:                  ErrorKindSchema,
		ErrorIDSchemaTupleMissingIndex:          ErrorKindSchema,
		ErrorIDSchemaTupleIndexOutOfRange:       ErrorKindSchema,
		ErrorIDSchemaTupleDuplicateSlot:         ErrorKindSchema,
		ErrorIDSchemaTupleMissingSlot:           ErrorKindSchema,
		ErrorIDValidationRequired:               ErrorKindValidation,
		ErrorIDValidationUnexpectedFlag:         ErrorKindValidation,
		ErrorIDValidationSchemaCommandForbidden: ErrorKindValidation,
		ErrorIDValidationInvalidType:            ErrorKindValidation,
		ErrorIDValidationString:                 ErrorKindValidation,
		ErrorIDValidationNumber:                 ErrorKindValidation,
		ErrorIDValidationTuple:                  ErrorKindValidation,
		ErrorIDValidationList:                   ErrorKindValidation,
		ErrorIDValidationFormat:                 ErrorKindValidation,
		ErrorIDValidationRange:                  ErrorKindValidation,
	}

	if len(got) != len(want) {
		t.Fatalf("fixture count mismatch: got %d want %d", len(got), len(want))
	}
	for id, kind := range want {
		if got[id] != kind {
			t.Fatalf("fixture mismatch for %s: got %q want %q", id, got[id], kind)
		}
		if MessageTemplate(id) == "" {
			t.Fatalf("missing template for %s", id)
		}
	}
}

func TestValidationErrorCollectMultipleDetails(t *testing.T) {
	err := NewValidationError(ErrorIDValidationRequired, nil)
	err = err.Add(NewErrorDetail(ErrorIDValidationString, ErrorKindValidation, map[string]string{"field": "mode"}))
	if len(err.Details) != 2 {
		t.Fatalf("expected 2 details, got %d", len(err.Details))
	}
}

func TestValidationErrorAddOnNilReceiver(t *testing.T) {
	var err *ValidationError
	err = err.Add(NewErrorDetail(ErrorIDValidationInvalidType, ErrorKindValidation, nil))
	if err == nil || len(err.Details) != 1 {
		t.Fatalf("expected new error with one detail, got %#v", err)
	}
}

func TestRenderMessageUnknownTemplate(t *testing.T) {
	if got := RenderMessage("unknown.id", nil); got != "validation failed" {
		t.Fatalf("unexpected fallback message: %q", got)
	}
}

func TestValidationErrorErrorMethodBranches(t *testing.T) {
	var nilErr *ValidationError
	if got := nilErr.Error(); got != "validation failed" {
		t.Fatalf("unexpected nil receiver message: %q", got)
	}

	empty := &ValidationError{}
	if got := empty.Error(); got != "validation failed" {
		t.Fatalf("unexpected empty details message: %q", got)
	}

	withMessage := &ValidationError{
		Details: []ErrorDetail{{ID: ErrorIDValidationRequired, Message: "custom message"}},
	}
	if got := withMessage.Error(); got != "custom message" {
		t.Fatalf("unexpected message detail return: %q", got)
	}

	withIDNoMessage := &ValidationError{
		Details: []ErrorDetail{{ID: ErrorIDValidationRequired}},
	}
	if got := withIDNoMessage.Error(); got != ErrorIDValidationRequired+": validation failed" {
		t.Fatalf("unexpected id fallback: %q", got)
	}

	noIDNoMessage := &ValidationError{
		Details: []ErrorDetail{{}},
	}
	if got := noIDNoMessage.Error(); got != "validation failed" {
		t.Fatalf("unexpected no-id fallback: %q", got)
	}
}
