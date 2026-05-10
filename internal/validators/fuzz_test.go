// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"testing"
	"time"
)

func FuzzValidateString(f *testing.F) {
	f.Add("warm")
	f.Add(" \t")
	f.Add("aGVsbG8=")
	f.Add("漢字")
	f.Fuzz(func(t *testing.T, s string) {
		_ = ValidateString(s, StringOptions{Alphabetic: true})
		_ = ValidateString(s, StringOptions{Base64: true})
		_ = ValidateString(s, StringOptions{UnicodeLetter: true})
		_ = ValidateString(s, StringOptions{BooleanString: true})
	})
}

func FuzzValidateNumber(f *testing.F) {
	f.Add("42")
	f.Add("3.14")
	f.Add("nan")
	f.Fuzz(func(t *testing.T, s string) {
		n, err := ParseNumberString(s)
		if err != nil {
			return
		}
		min, max, multiple := 0.0, 1000.0, 0.5
		_ = ValidateNumber(n, NumberOptions{Min: &min, Max: &max, MultipleOf: &multiple})
	})
}

func FuzzValidateFormats(f *testing.F) {
	f.Add("https://example.com")
	f.Add("arn:aws:sns:us-east-2:123456789012:topic")
	f.Add("user@example.com")
	f.Add("#A1B2C3")
	f.Add("2026-05-06")
	f.Add("2026-05-06T12:00:00Z")
	f.Add("123000")
	f.Add("30m")
	f.Fuzz(func(t *testing.T, s string) {
		_ = ValidateURL(s, URLOptions{})
		_ = ValidateARN(s, ARNOptions{})
		_ = ValidateEmail(s, nil)
		_ = ValidateColor(s, "hex", true)
		_ = ValidateDate(s, "ISO8601")
		_ = ValidateDateTime(s, "RFC3339")
		_ = ValidateClockTime(s, "HHMMSS")
		min, max := 5*time.Minute, 2*time.Hour
		_ = ValidateDuration(s, &min, &max)
	})
}
