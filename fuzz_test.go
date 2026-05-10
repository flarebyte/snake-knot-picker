// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import (
	"strings"
	"testing"
)

func FuzzCompileCommandDocument(f *testing.F) {
	f.Add(`{"version":"1","commandPath":["wash","start"],"flags":[{"kind":"string","name":"mode","schema":["schema","string","--required"]}]}`)
	f.Add(`{"version":"1","commandPath":["wash","start"],"flags":[{"kind":"tuple","name":"range","schema":["schema","tuple","--size","2"],"schemas":[["schema","number","--tuple","0","--int"],["schema","number","--tuple","1","--int"]]}]}`)
	f.Add(`{"version":"1","commandPath":["wash"],"flags":[{"kind":"tuple","name":"range","schema":["schema","tuple","--size","2"],"schemas":[["schema","number","--tuple","999999999","--int"]]}]}`)
	f.Fuzz(func(t *testing.T, raw string) {
		doc, err := ParseCommandDocumentJSON([]byte(raw))
		if err != nil {
			return
		}
		_, _ = CompileCommandDocument(doc)
	})
}

func FuzzValidateRuntimeArgv(f *testing.F) {
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}},
			{Kind: "number", Name: "spin", Schema: []string{"schema", "number"}},
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
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		f.Fatalf("compile seed command failed: %v", err)
	}
	f.Add("wash start --mode normal --spin 1200 --range 10,20")
	f.Add("wash start schema string --required")
	f.Add("wash start --unknown x")
	f.Fuzz(func(t *testing.T, argvRaw string) {
		argv := strings.Fields(argvRaw)
		_, _ = Validate(compiled, argv)
	})
}
