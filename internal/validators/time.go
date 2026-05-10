// purpose: Implement date, datetime, time, and duration validation helpers.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"time"

	"github.com/flarebyte/snake-knot-picker"
)

func ValidateDate(value, layout string) error {
	if layout == "" || layout == "ISO8601" {
		layout = "2006-01-02"
	}
	if _, err := time.Parse(layout, value); err != nil {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "date"})
	}
	return nil
}

func ValidateDateTime(value, layout string) error {
	switch layout {
	case "", "RFC3339":
		_, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "datetime"})
		}
	case "RFC1123Z":
		_, err := time.Parse(time.RFC1123Z, value)
		if err != nil {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "datetime"})
		}
	case "Unix":
		_, err := ParseNumberString(value)
		if err != nil {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "datetime"})
		}
	default:
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "datetime"})
	}
	return nil
}

func ValidateClockTime(value, layout string) error {
	switch layout {
	case "HHMMSS":
		if _, err := time.Parse("150405", value); err != nil {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "time"})
		}
	case "HHMM":
		if _, err := time.Parse("1504", value); err != nil {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "time"})
		}
	default:
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "time"})
	}
	return nil
}

func ValidateDuration(value string, min, max *time.Duration) error {
	d, err := time.ParseDuration(value)
	if err != nil {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "duration"})
	}
	if min != nil && d < *min {
		return picker.NewValidationError(picker.ErrorIDValidationRange, map[string]string{"bound": "min_duration"})
	}
	if max != nil && d > *max {
		return picker.NewValidationError(picker.ErrorIDValidationRange, map[string]string{"bound": "max_duration"})
	}
	return nil
}
