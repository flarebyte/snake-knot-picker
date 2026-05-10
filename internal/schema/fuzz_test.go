// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import (
	"strings"
	"testing"
)

func FuzzParseTokens(f *testing.F) {
	seeds := [][]string{
		{"schema", "string", "--enum", "cold,warm,hot", "--required"},
		{"schema", "tuple", "--size", "2"},
		{"schema", "number", "--tuple", "0", "--int"},
		{"custom", "postal-code", "--country", "US"},
		{"schema", "string", "--enum"},
		{"schema", "string", "--tuple", "999999999999"},
	}
	for _, s := range seeds {
		f.Add(strings.Join(s, " "))
	}
	f.Fuzz(func(t *testing.T, input string) {
		tokens := strings.Fields(input)
		_, _ = ParseTokens(tokens)
	})
}
