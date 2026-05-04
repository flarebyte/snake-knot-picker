package schema

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

type commandFixture struct {
	Flags []struct {
		Schema  []string   `json:"schema"`
		Schemas [][]string `json:"schemas"`
	} `json:"flags"`
}

func TestParseTokensArgsCommandFixtures(t *testing.T) {
	raw, err := os.ReadFile("../../doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read args-command fixture: %v", err)
	}
	var doc commandFixture
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("decode args-command fixture: %v", err)
	}

	for i, f := range doc.Flags {
		ast, err := ParseTokens(f.Schema)
		if err != nil {
			t.Fatalf("schema parse failed for flag[%d]: %v", i, err)
		}
		assertRawEqual(t, ast.Raw, f.Schema)
		for j, child := range f.Schemas {
			childAST, err := ParseTokens(child)
			if err != nil {
				t.Fatalf("schema parse failed for flag[%d].schemas[%d]: %v", i, j, err)
			}
			assertRawEqual(t, childAST.Raw, child)
		}
	}
}

func TestParseTokensEdgeCaseCSVTokenizeRows(t *testing.T) {
	file, err := os.Open("../../doc/design-meta/examples/parser-edge-cases.csv")
	if err != nil {
		t.Fatalf("open edge-case csv: %v", err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("read edge-case csv: %v", err)
	}
	for i, row := range rows {
		if i == 0 || len(row) < 5 {
			continue
		}
		stage := row[1]
		input := row[2]
		expectedID := row[4]
		if stage != "schema-tokenize" {
			continue
		}

		tokens := strings.Fields(input)
		_, err := ParseTokens(tokens)
		if expectedID == "" && err != nil {
			t.Fatalf("unexpected parse failure for %q: %v", input, err)
		}
		if expectedID != "" {
			if err == nil {
				t.Fatalf("expected error %s for %q", expectedID, input)
			}
			verr, ok := err.(*picker.ValidationError)
			if !ok || len(verr.Details) == 0 {
				t.Fatalf("expected structured error for %q, got %T", input, err)
			}
			if verr.Details[0].ID != expectedID {
				t.Fatalf("unexpected error id for %q: got=%s want=%s", input, verr.Details[0].ID, expectedID)
			}
		}
	}
}

func TestParseTokensMalformedInput(t *testing.T) {
	cases := []struct {
		name   string
		tokens []string
		id     string
	}{
		{name: "empty", tokens: nil, id: picker.ErrorIDSchemaInvalidValue},
		{name: "missing-operator", tokens: []string{"schema"}, id: picker.ErrorIDSchemaInvalidValue},
		{name: "bad-head", tokens: []string{"user", "string"}, id: picker.ErrorIDSchemaInvalidValue},
		{name: "non-flag-after-operator", tokens: []string{"schema", "string", "oops"}, id: picker.ErrorIDSchemaInvalidValue},
		{name: "missing-value-known-flag", tokens: []string{"schema", "string", "--enum"}, id: picker.ErrorIDSchemaMissingValue},
		{name: "tuple-not-int", tokens: []string{"schema", "string", "--tuple", "x"}, id: picker.ErrorIDSchemaInvalidValue},
		{name: "tuple-negative", tokens: []string{"schema", "string", "--tuple", "-1"}, id: picker.ErrorIDSchemaInvalidValue},
		{name: "codepoint-needs-two", tokens: []string{"schema", "string", "--codepoint-range", "U+3040"}, id: picker.ErrorIDSchemaMissingValue},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseTokens(tc.tokens)
			if err == nil {
				t.Fatalf("expected error %s", tc.id)
			}
			verr, ok := err.(*picker.ValidationError)
			if !ok || len(verr.Details) == 0 {
				t.Fatalf("expected structured validation error, got %T", err)
			}
			if verr.Details[0].ID != tc.id {
				t.Fatalf("unexpected error id: got=%s want=%s", verr.Details[0].ID, tc.id)
			}
		})
	}
}

func TestParseTokensSupportsRepeatedFlags(t *testing.T) {
	tokens := []string{
		"schema", "string",
		"--allow-service", "s3",
		"--allow-service", "sns",
		"--required",
	}
	ast, err := ParseTokens(tokens)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if ast.Head != "schema" || ast.Operator != "string" {
		t.Fatalf("unexpected ast: %#v", ast)
	}
	if len(ast.Flags) < 2 {
		t.Fatalf("expected multiple flags, got %#v", ast.Flags)
	}
	values := findFlagValues(ast.Flags, "--allow-service")
	if len(values) != 2 || values[0] != "s3" || values[1] != "sns" {
		t.Fatalf("unexpected repeated values: %#v", values)
	}
}

func assertRawEqual(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("raw length mismatch: got=%v want=%v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("raw mismatch at %d: got=%q want=%q", i, got[i], want[i])
		}
	}
}

