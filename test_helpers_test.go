package picker

import (
	"os"
	"testing"
)

func makeInvalidTupleMissingSlotDoc() CommandDocument {
	return CommandDocument{
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
}

func mustLoadArgsCommandFixture(t testing.TB) []byte {
	t.Helper()
	raw, err := os.ReadFile("doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read args-command fixture: %v", err)
	}
	return raw
}

func mustCompileArgsCommandFixture(t testing.TB) CompiledCommand {
	t.Helper()
	raw := mustLoadArgsCommandFixture(t)
	doc, err := ParseCommandDocumentJSON(raw)
	if err != nil {
		t.Fatalf("parse args-command fixture: %v", err)
	}
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		t.Fatalf("compile args-command fixture: %v", err)
	}
	return compiled
}
