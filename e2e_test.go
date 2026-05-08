package picker

import (
	"os"
	"testing"
)

func TestEndToEndCaseBListArgvIsPrimary(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	argv := []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--extra-rinse", "--range", "10,20"}

	got, err := ValidateWithDocumentJSON(raw, argv)
	if err != nil {
		t.Fatalf("unexpected error for list argv: %v", err)
	}
	if got == nil {
		t.Fatal("expected parse result")
	}
	assertStringValue(t, got.Values, "mode", "normal")
	assertNumberValue(t, got.Values, "spin", 1200)
	assertBoolValue(t, got.Values, "extra-rinse", true)
	assertTupleStringValues(t, got.Values, "range", "10", "20")
}

func TestEndToEndCaseAStringArgvIsConvenienceOnly(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	line := "wash start --mode normal --spin 1200 --extra-rinse --range 10,20"

	got, err := ValidateWithDocumentJSONArgvString(raw, line)
	if err != nil {
		t.Fatalf("unexpected error for argv string convenience: %v", err)
	}
	if got == nil {
		t.Fatal("expected parse result")
	}
	assertStringValue(t, got.Values, "mode", "normal")
	assertNumberValue(t, got.Values, "spin", 1200)
	assertBoolValue(t, got.Values, "extra-rinse", true)
	assertTupleStringValues(t, got.Values, "range", "10", "20")
}

func TestEndToEndCaseAStringArgvQuotedValueLimitation(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)

	// Primary API (case B): caller controls tokenization, so spaces in values are preserved.
	listArgv := []string{"wash", "start", "--mode", "normal with space", "--spin", "1200", "--range", "10,20"}
	listGot, listErr := ValidateWithDocumentJSON(raw, listArgv)
	if listErr != nil {
		t.Fatalf("unexpected case B error: %v", listErr)
	}
	assertStringValue(t, listGot.Values, "mode", "normal with space")

	// Convenience API (case A): strings.Fields does not preserve shell-style quoting.
	line := `wash start --mode "normal with space" --spin 1200 --range 10,20`
	_, err := ValidateWithDocumentJSONArgvString(raw, line)
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
}

func TestEndToEndValidateWithDocumentJSON(t *testing.T) {
	raw, err := os.ReadFile("doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	cases := []struct {
		name      string
		argv      []string
		wantErrID string
		wantKind  string
		check     func(t *testing.T, got *ParseResult)
	}{
		{
			name: "valid-basic",
			argv: []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--extra-rinse", "--range", "10,20"},
			check: func(t *testing.T, got *ParseResult) {
				t.Helper()
				if got == nil {
					t.Fatal("expected parse result")
				}
				if len(got.CommandPath) != 2 || got.CommandPath[0] != "wash" || got.CommandPath[1] != "start" {
					t.Fatalf("unexpected command path: %#v", got.CommandPath)
				}
				assertStringValue(t, got.Values, "mode", "normal")
				assertNumberValue(t, got.Values, "spin", 1200)
				assertBoolValue(t, got.Values, "extra-rinse", true)
				assertTupleStringValues(t, got.Values, "range", "10", "20")
				if _, exists := got.Values["not-a-flag"]; exists {
					t.Fatalf("unexpected value key present: not-a-flag")
				}
			},
		},
		{
			name:      "runtime-unknown-flag",
			argv:      []string{"wash", "start", "--not-a-flag", "x"},
			wantErrID: ErrorIDValidationUnexpectedFlag,
			wantKind:  ErrorKindValidation,
		},
		{
			name:      "runtime-schema-command-forbidden",
			argv:      []string{"wash", "start", "schema", "string", "--required"},
			wantErrID: ErrorIDValidationSchemaCommandForbidden,
			wantKind:  ErrorKindValidation,
		},
		{
			name:      "runtime-invalid-type",
			argv:      []string{"wash", "start", "--spin", "abc"},
			wantErrID: ErrorIDValidationInvalidType,
			wantKind:  ErrorKindValidation,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := ValidateWithDocumentJSON(raw, tc.argv)
			if tc.wantErrID != "" {
				assertErrorDetail(t, err, tc.wantErrID, tc.wantKind)
				if got != nil {
					t.Fatalf("expected nil parse result on error, got=%#v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.check != nil {
				tc.check(t, got)
			}
		})
	}
}

func TestEndToEndSchemaErrorHappensBeforeRegistration(t *testing.T) {
	doc := makeInvalidTupleMissingSlotDoc()
	_, err := ValidateWithDocument(doc, []string{"wash", "start", "--range", "10,20"})
	assertErrorDetail(t, err, ErrorIDSchemaTupleMissingSlot, ErrorKindSchema)
}
