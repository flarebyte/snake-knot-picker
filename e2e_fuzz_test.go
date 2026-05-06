package picker

import (
	"encoding/json"
	"strings"
	"testing"
)

func FuzzE2EValidateWithDocumentJSON(f *testing.F) {
	fixture := mustLoadArgsCommandFixture(f)

	f.Add(string(fixture), "wash start --mode normal --spin 1200 --extra-rinse --range 10,20")
	f.Add(string(fixture), "wash start --not-a-flag x")
	f.Add(string(fixture), "wash start schema string --required")
	f.Add(string(fixture), "wash start --spin abc")
	f.Add(string(fixture), "")

	f.Fuzz(func(t *testing.T, rawDoc string, argvRaw string) {
		argv := strings.Fields(argvRaw)
		got, err := ValidateWithDocumentJSON([]byte(rawDoc), argv)
		assertE2EResultInvariant(t, got, err)
	})
}

func FuzzE2EValidateCompiledRuntime(f *testing.F) {
	compiled := mustCompileArgsCommandFixture(f)

	f.Add("wash start --mode normal --spin 1200 --extra-rinse --range 10,20")
	f.Add("wash start --mode=normal --spin=1200 --range=10,20")
	f.Add("wash start --range 10")
	f.Add("schema string --required")
	f.Add("--unknown x")
	f.Add("")

	f.Fuzz(func(t *testing.T, argvRaw string) {
		argv := strings.Fields(argvRaw)
		got, err := Validate(compiled, argv)
		assertE2EResultInvariant(t, got, err)
		if err == nil {
			if len(got.CommandPath) != len(compiled.CommandPath) {
				t.Fatalf("command path length mismatch: got=%v want=%v", got.CommandPath, compiled.CommandPath)
			}
			for i := range compiled.CommandPath {
				if got.CommandPath[i] != compiled.CommandPath[i] {
					t.Fatalf("command path mismatch: got=%v want=%v", got.CommandPath, compiled.CommandPath)
				}
			}
		}
	})
}

func FuzzE2ECompileThenValidateFromDocument(f *testing.F) {
	baseDoc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string", "--required"}},
			{Kind: "number", Name: "spin", Schema: []string{"schema", "number", "--int"}},
			{
				Kind:   "tuple",
				Name:   "range",
				Schema: []string{"schema", "tuple", "--size", "2"},
				Schemas: [][]string{
					{"schema", "number", "--tuple", "0", "--int"},
					{"schema", "number", "--tuple", "1", "--int"},
				},
			},
		},
	}
	b, _ := json.Marshal(baseDoc)
	f.Add(string(b), "wash start --mode normal --spin 1200 --range 10,20")
	f.Add(string(b), "wash start --spin abc")
	f.Add(string(b), "wash start --unknown x")
	f.Add(`{"version":"1","commandPath":["wash","start"],"flags":[{"kind":"tuple","name":"range","schema":["schema","tuple","--size","2"],"schemas":[["schema","number","--tuple","0","--int"]]}]}`, "wash start --range 10,20")

	f.Fuzz(func(t *testing.T, rawDoc string, argvRaw string) {
		doc, err := ParseCommandDocumentJSON([]byte(rawDoc))
		if err != nil {
			assertE2EResultInvariant(t, nil, err)
			return
		}
		got, err := ValidateWithDocument(doc, strings.Fields(argvRaw))
		assertE2EResultInvariant(t, got, err)
	})
}

func assertE2EResultInvariant(t *testing.T, got *ParseResult, err error) {
	t.Helper()
	if err == nil {
		if got == nil {
			t.Fatal("expected non-nil parse result when err is nil")
		}
		if got.Values == nil {
			t.Fatal("expected non-nil values map on successful parse")
		}
		return
	}
	if got != nil {
		t.Fatalf("expected nil parse result when err is non-nil, got=%#v", got)
	}
	verr, ok := err.(*ValidationError)
	if !ok {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if len(verr.Details) == 0 {
		t.Fatal("expected at least one validation detail")
	}
	for _, d := range verr.Details {
		if d.ID == "" {
			t.Fatal("expected non-empty error detail ID")
		}
		if d.Kind != ErrorKindSchema && d.Kind != ErrorKindValidation {
			t.Fatalf("unexpected error kind %q", d.Kind)
		}
	}
}
