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
