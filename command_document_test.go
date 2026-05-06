package picker

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCommandDocumentCompileFromFixture(t *testing.T) {
	raw, err := os.ReadFile("doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	doc, err := ParseCommandDocumentJSON(raw)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		t.Fatalf("compile fixture: %v", err)
	}
	if len(compiled.Flags) != len(doc.Flags) {
		t.Fatalf("compiled flag count mismatch: got=%d want=%d", len(compiled.Flags), len(doc.Flags))
	}
}

func TestCommandDocumentJSONRoundTrip(t *testing.T) {
	doc := NewCommandBuilder("wash", "start").
		SetAdminOnly(true).
		AddFlag(CommandFlagDef{
			Kind:   "string",
			Name:   "mode",
			Schema: []string{"schema", "string", "--enum", "normal,delicate"},
		}).
		Build()

	b, err := doc.ToJSON()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	parsed, err := ParseCommandDocumentJSON(b)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	left, _ := json.Marshal(doc)
	right, _ := json.Marshal(parsed)
	if string(left) != string(right) {
		t.Fatalf("round-trip mismatch:\nleft=%s\nright=%s", left, right)
	}
}

func TestCommandDocumentCompileErrors(t *testing.T) {
	cases := []struct {
		name string
		doc  CommandDocument
		id   string
	}{
		{
			name: "missing-version",
			doc: CommandDocument{
				CommandPath: []string{"wash"},
				Flags:       []CommandFlagDef{{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}}},
			},
			id: ErrorIDSchemaInvalidValue,
		},
		{
			name: "missing-path",
			doc: CommandDocument{
				Version: "1",
				Flags:   []CommandFlagDef{{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}}},
			},
			id: ErrorIDSchemaInvalidValue,
		},
		{
			name: "duplicate-name",
			doc: CommandDocument{
				Version:     "1",
				CommandPath: []string{"wash"},
				Flags: []CommandFlagDef{
					{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}},
					{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}},
				},
			},
			id: ErrorIDSchemaInvalidCombination,
		},
		{
			name: "invalid-kind",
			doc: CommandDocument{
				Version:     "1",
				CommandPath: []string{"wash"},
				Flags:       []CommandFlagDef{{Kind: "map", Name: "mode", Schema: []string{"schema", "string"}}},
			},
			id: ErrorIDSchemaInvalidValue,
		},
		{
			name: "tuple-missing-slot",
			doc: func() CommandDocument {
				doc := makeInvalidTupleMissingSlotDoc()
				doc.CommandPath = []string{"wash"}
				return doc
			}(),
			id: ErrorIDSchemaTupleMissingSlot,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := CompileCommandDocument(tc.doc)
			assertDocErrID(t, err, tc.id)
		})
	}
}

func TestCommandDocumentCustomValidatorAndAdminOnly(t *testing.T) {
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash"},
		AdminOnly:   true,
		Flags: []CommandFlagDef{
			{
				Kind:   "string",
				Name:   "postal-code",
				Schema: []string{"custom", "postal-code", "--country", "US", "--required"},
			},
		},
	}
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		t.Fatalf("compile custom doc: %v", err)
	}
	if !compiled.AdminOnly {
		t.Fatal("adminOnly should be preserved")
	}
}

func assertDocErrID(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error %s", want)
	}
	verr, ok := err.(*ValidationError)
	if !ok || len(verr.Details) == 0 {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if verr.Details[0].ID != want {
		t.Fatalf("unexpected id: got=%s want=%s", verr.Details[0].ID, want)
	}
}
