package validators

import (
	"encoding/base64"
	"strings"
	"unicode"

	"github.com/flarebyte/snake-knot-picker"
)

type StringOptions struct {
	Enum                 []string
	Alphabetic           bool
	Whitespace           bool
	Lowercase            bool
	Uppercase            bool
	Punctuation          bool
	Hexa                 bool
	Blank                bool
	UnicodeLetter        bool
	UnicodeNumber        bool
	UnicodePunctuation   bool
	UnicodeSymbol        bool
	UnicodeSeparator     bool
	Latin                bool
	Han                  bool
	Devanagari           bool
	Arabic               bool
	Hiragana             bool
	Katakana             bool
	Hangul               bool
	Tamil                bool
	Gujarati             bool
	Ethiopic             bool
	Base64               bool
	StartsWith           string
	BooleanString        bool
}

func ParseEnumCandidates(raw, separator string) ([]string, error) {
	if separator == "" {
		separator = ","
	}
	parts := strings.Split(raw, separator)
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaEnumEmpty, nil)
		}
		if part != trimmed {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaEnumWhitespace, map[string]string{"value": part})
		}
		out = append(out, trimmed)
	}
	return out, nil
}

func ValidateString(value string, options StringOptions) error {
	if len(options.Enum) > 0 {
		ok := false
		for _, candidate := range options.Enum {
			if value == candidate {
				ok = true
				break
			}
		}
		if !ok {
			return picker.NewValidationError(picker.ErrorIDValidationString, map[string]string{"reason": "enum"})
		}
	}

	checks := []struct {
		on bool
		fn func(string) bool
	}{
		{on: options.Alphabetic, fn: isAlphabetic},
		{on: options.Whitespace, fn: isWhitespace},
		{on: options.Lowercase, fn: isLowercase},
		{on: options.Uppercase, fn: isUppercase},
		{on: options.Punctuation, fn: isPunctuation},
		{on: options.Hexa, fn: isHexa},
		{on: options.Blank, fn: isBlank},
		{on: options.UnicodeLetter, fn: isUnicodeLetter},
		{on: options.UnicodeNumber, fn: isUnicodeNumber},
		{on: options.UnicodePunctuation, fn: isUnicodePunctuation},
		{on: options.UnicodeSymbol, fn: isUnicodeSymbol},
		{on: options.UnicodeSeparator, fn: isUnicodeSeparator},
		{on: options.Latin, fn: isLatin},
		{on: options.Han, fn: isHan},
		{on: options.Devanagari, fn: isDevanagari},
		{on: options.Arabic, fn: isArabic},
		{on: options.Hiragana, fn: isHiragana},
		{on: options.Katakana, fn: isKatakana},
		{on: options.Hangul, fn: isHangul},
		{on: options.Tamil, fn: isTamil},
		{on: options.Gujarati, fn: isGujarati},
		{on: options.Ethiopic, fn: isEthiopic},
	}
	for _, c := range checks {
		if c.on && !c.fn(value) {
			return picker.NewValidationError(picker.ErrorIDValidationString, nil)
		}
	}

	if options.Base64 && !isBase64(value) {
		return picker.NewValidationError(picker.ErrorIDValidationString, map[string]string{"reason": "base64"})
	}
	if options.StartsWith != "" && !strings.HasPrefix(value, options.StartsWith) {
		return picker.NewValidationError(picker.ErrorIDValidationString, map[string]string{"reason": "prefix"})
	}
	if options.BooleanString && value != "true" && value != "false" {
		return picker.NewValidationError(picker.ErrorIDValidationString, map[string]string{"reason": "boolean"})
	}
	return nil
}

func allRunes(s string, pred func(rune) bool) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !pred(r) {
			return false
		}
	}
	return true
}

func isAlphabetic(s string) bool      { return allRunes(s, func(r rune) bool { return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') }) }
func isWhitespace(s string) bool      { return allRunes(s, unicode.IsSpace) }
func isLowercase(s string) bool       { return allRunes(s, unicode.IsLower) }
func isUppercase(s string) bool       { return allRunes(s, unicode.IsUpper) }
func isPunctuation(s string) bool     { return allRunes(s, unicode.IsPunct) }
func isHexa(s string) bool            { return allRunes(s, func(r rune) bool { return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') }) }
func isBlank(s string) bool           { return allRunes(s, func(r rune) bool { return r == ' ' || r == '\t' }) }
func isUnicodeLetter(s string) bool   { return allRunes(s, unicode.IsLetter) }
func isUnicodeNumber(s string) bool   { return allRunes(s, unicode.IsNumber) }
func isUnicodePunctuation(s string) bool {
	return allRunes(s, unicode.IsPunct)
}
func isUnicodeSymbol(s string) bool    { return allRunes(s, unicode.IsSymbol) }
func isUnicodeSeparator(s string) bool { return allRunes(s, unicode.IsSpace) }

func isLatin(s string) bool      { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Latin) }) }
func isHan(s string) bool        { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Han) }) }
func isDevanagari(s string) bool { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Devanagari) }) }
func isArabic(s string) bool     { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Arabic) }) }
func isHiragana(s string) bool   { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Hiragana) }) }
func isKatakana(s string) bool   { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Katakana) }) }
func isHangul(s string) bool     { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Hangul) }) }
func isTamil(s string) bool      { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Tamil) }) }
func isGujarati(s string) bool   { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Gujarati) }) }
func isEthiopic(s string) bool   { return allRunes(s, func(r rune) bool { return unicode.In(r, unicode.Ethiopic) }) }

func isBase64(s string) bool {
	if s == "" {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	return base64.StdEncoding.EncodeToString(b) == s
}
