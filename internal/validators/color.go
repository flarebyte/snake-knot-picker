// purpose: Implement color-string validation with configurable format expectations.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"regexp"

	"github.com/flarebyte/snake-knot-picker"
)

var (
	hex6 = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	hex8 = regexp.MustCompile(`^#[0-9a-fA-F]{8}$`)
)

// ValidateColor validates a color string according to format and alpha settings.
func ValidateColor(value, format string, allowAlpha bool) error {
	if format != "hex" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "color"})
	}
	if hex6.MatchString(value) {
		return nil
	}
	if allowAlpha && hex8.MatchString(value) {
		return nil
	}
	return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "color"})
}
