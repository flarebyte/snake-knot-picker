package argv

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
	"github.com/flarebyte/snake-knot-picker/internal/testutil"
)

func TestParserValidPathsAndFlagForms(t *testing.T) {
	p := NewParser()
	cmd := picker.CompiledCommand{
		CommandPath: []string{"wash", "start"},
		Flags: []picker.CompiledFlag{
			{Kind: "boolean", Name: "extra-rinse"},
			{Kind: "string", Name: "mode"},
			{Kind: "number", Name: "spin"},
			{Kind: "tuple", Name: "range", TupleSize: 2},
			{Kind: "string", Name: "add", Repeatable: true},
		},
	}

	got, err := p.Parse(cmd, []string{
		"wash", "start",
		"--extra-rinse",
		"--mode", "delicate",
		"--spin", "1200",
		"--range", "10,20",
		"--add", "soap",
		"--add", "bleach,rinse",
	})
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if got == nil || len(got.Values) == 0 {
		t.Fatalf("unexpected empty parse result: %#v", got)
	}
	if got.Values["extra-rinse"].Bool == nil || !*got.Values["extra-rinse"].Bool {
		t.Fatalf("expected bool true: %#v", got.Values["extra-rinse"])
	}
	if got.Values["mode"].String == nil || *got.Values["mode"].String != "delicate" {
		t.Fatalf("expected mode delicate: %#v", got.Values["mode"])
	}
	if got.Values["spin"].Number == nil || *got.Values["spin"].Number != 1200 {
		t.Fatalf("expected spin 1200: %#v", got.Values["spin"])
	}
	if len(got.Values["range"].Tuple) != 2 {
		t.Fatalf("expected tuple size 2: %#v", got.Values["range"])
	}
	if len(got.Values["add"].List) != 3 {
		t.Fatalf("expected repeatable list size 3: %#v", got.Values["add"])
	}
}

func TestParserErrors(t *testing.T) {
	p := NewParser()
	cmd := picker.CompiledCommand{
		CommandPath: []string{"wash", "start"},
		Flags: []picker.CompiledFlag{
			{Kind: "string", Name: "mode"},
			{Kind: "tuple", Name: "range", TupleSize: 2},
			{Kind: "string", Name: "add", Repeatable: true},
		},
	}
	cases := []struct {
		name string
		argv []string
		id   string
	}{
		{name: "unexpected-flag", argv: []string{"wash", "start", "--unknown", "x"}, id: picker.ErrorIDValidationUnexpectedFlag},
		{name: "schema-injection", argv: []string{"wash", "start", "schema", "string"}, id: picker.ErrorIDValidationSchemaCommandForbidden},
		{name: "missing-value", argv: []string{"wash", "start", "--mode"}, id: picker.ErrorIDValidationInvalidType},
		{name: "tuple-arity", argv: []string{"wash", "start", "--range", "10"}, id: picker.ErrorIDValidationTuple},
		{name: "inline-mode-forbidden", argv: []string{"wash", "start", "--mode=normal", "--range", "10,20"}, id: picker.ErrorIDValidationInvalidType},
		{name: "repeatable-inline-forbidden", argv: []string{"wash", "start", "--mode", "normal", "--range", "10,20", "--add=x"}, id: picker.ErrorIDValidationInvalidType},
		{name: "non-flag-token", argv: []string{"wash", "start", "oops"}, id: picker.ErrorIDValidationInvalidType},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := p.Parse(cmd, tc.argv)
			assertErrorID(t, err, tc.id)
		})
	}
}

func TestParserNoCommandPathPrefixStillParses(t *testing.T) {
	p := NewParser()
	cmd := picker.CompiledCommand{
		CommandPath: []string{"wash", "start"},
		Flags:       []picker.CompiledFlag{{Kind: "string", Name: "mode"}},
	}
	got, err := p.Parse(cmd, []string{"--mode", "normal"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Values["mode"].String == nil || *got.Values["mode"].String != "normal" {
		t.Fatalf("unexpected mode value: %#v", got.Values["mode"])
	}
}

func assertErrorID(t *testing.T, err error, want string) {
	t.Helper()
	_ = testutil.MustValidationErrorWithID(t, err, want)
}
