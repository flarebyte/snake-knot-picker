package picker

import (
	"os"
	"testing"
)

func TestEndToEndValidateWithDocumentJSON(t *testing.T) {
	raw, err := os.ReadFile("doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	cases := []struct {
		name      string
		argv      []string
		wantErrID string
		check     func(t *testing.T, got *ParseResult)
	}{
		{
			name: "valid-basic",
			argv: []string{"wash", "start", "--mode", "normal", "--spin=1200", "--extra-rinse", "--range", "10,20"},
			check: func(t *testing.T, got *ParseResult) {
				t.Helper()
				if got == nil {
					t.Fatal("expected parse result")
				}
				if got.Values["mode"].String == nil || *got.Values["mode"].String != "normal" {
					t.Fatalf("unexpected mode: %#v", got.Values["mode"])
				}
				if got.Values["spin"].Number == nil || *got.Values["spin"].Number != 1200 {
					t.Fatalf("unexpected spin: %#v", got.Values["spin"])
				}
				if got.Values["extra-rinse"].Bool == nil || !*got.Values["extra-rinse"].Bool {
					t.Fatalf("unexpected extra-rinse: %#v", got.Values["extra-rinse"])
				}
				if len(got.Values["range"].Tuple) != 2 {
					t.Fatalf("unexpected range tuple: %#v", got.Values["range"])
				}
			},
		},
		{
			name:      "runtime-unknown-flag",
			argv:      []string{"wash", "start", "--not-a-flag", "x"},
			wantErrID: ErrorIDValidationUnexpectedFlag,
		},
		{
			name:      "runtime-schema-command-forbidden",
			argv:      []string{"wash", "start", "schema", "string", "--required"},
			wantErrID: ErrorIDValidationSchemaCommandForbidden,
		},
		{
			name:      "runtime-invalid-type",
			argv:      []string{"wash", "start", "--spin", "abc"},
			wantErrID: ErrorIDValidationInvalidType,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := ValidateWithDocumentJSON(raw, tc.argv)
			if tc.wantErrID != "" {
				assertDocErrID(t, err, tc.wantErrID)
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
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{
				Kind:   "tuple",
				Name:   "range",
				Schema: []string{"schema", "tuple", "--size", "2"},
				Schemas: [][]string{
					{"schema", "number", "--tuple", "0", "--int"},
				},
			},
		},
	}
	_, err := ValidateWithDocument(doc, []string{"wash", "start", "--range", "10,20"})
	assertDocErrID(t, err, ErrorIDSchemaTupleMissingSlot)
}
