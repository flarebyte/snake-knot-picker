package picker

import (
	"regexp"
	"strings"
	"testing"
)

func TestDefaultManualFlagNameValidator(t *testing.T) {
	v := DefaultManualFlagNameValidator()
	valid := []string{"mode", "mode_1", "mode.name", "mode-name", "A1"}
	for _, name := range valid {
		if err := v(name); err != nil {
			t.Fatalf("expected valid name %q, got err=%v", name, err)
		}
	}
	invalid := []string{"", " mode", "mode!", "mode/name", strings.Repeat("a", 65)}
	for _, name := range invalid {
		if err := v(name); err == nil {
			t.Fatalf("expected invalid name %q", name)
		}
	}
}

func TestRegexFlagNameValidatorFactory(t *testing.T) {
	v := NewRegexFlagNameValidator(regexp.MustCompile(`^[A-Za-z0-9._-]+$`), 2, 5)
	if err := v("ab"); err != nil {
		t.Fatalf("expected valid name: %v", err)
	}
	if err := v("a"); err == nil {
		t.Fatal("expected min-length rejection")
	}
	if err := v("abcdef"); err == nil {
		t.Fatal("expected max-length rejection")
	}
	if err := v("ab!"); err == nil {
		t.Fatal("expected regex charset rejection")
	}
}

func TestCompileCommandDocumentWithInjectedFlagNameValidator(t *testing.T) {
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}},
		},
	}
	custom := func(name string) error {
		if !strings.HasPrefix(name, "x-") {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
		}
		return nil
	}
	_, err := CompileCommandDocumentWithOptions(doc, CompileOptions{FlagNameValidator: custom})
	if err == nil {
		t.Fatal("expected custom validator rejection")
	}
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func BenchmarkFlagNameValidatorRegex(b *testing.B) {
	v := DefaultRegexFlagNameValidator()
	name := "mode_123.alpha-beta"
	for i := 0; i < b.N; i++ {
		_ = v(name)
	}
}

func BenchmarkFlagNameValidatorManual(b *testing.B) {
	v := DefaultManualFlagNameValidator()
	name := "mode_123.alpha-beta"
	for i := 0; i < b.N; i++ {
		_ = v(name)
	}
}
