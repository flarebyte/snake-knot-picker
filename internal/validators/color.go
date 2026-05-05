package validators

import (
	"regexp"

	"github.com/flarebyte/snake-knot-picker"
)

var (
	hex6 = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	hex8 = regexp.MustCompile(`^#[0-9a-fA-F]{8}$`)
)

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
