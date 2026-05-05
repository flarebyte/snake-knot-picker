package validators

import (
	"math"
	"strconv"

	"github.com/flarebyte/snake-knot-picker"
)

type NumberOptions struct {
	Int        bool
	Min        *float64
	Max        *float64
	MultipleOf *float64
}

func ParseNumberString(value string) (float64, error) {
	n, err := strconv.ParseFloat(value, 64)
	if err != nil || math.IsNaN(n) || math.IsInf(n, 0) {
		return 0, picker.NewValidationError(picker.ErrorIDValidationInvalidType, map[string]string{"kind": "number"})
	}
	return n, nil
}

func ValidateNumber(value float64, options NumberOptions) error {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return picker.NewValidationError(picker.ErrorIDValidationNumber, nil)
	}
	if options.Int && value != math.Trunc(value) {
		return picker.NewValidationError(picker.ErrorIDValidationNumber, map[string]string{"reason": "int"})
	}
	if options.Min != nil && value < *options.Min {
		return picker.NewValidationError(picker.ErrorIDValidationRange, map[string]string{"bound": "min"})
	}
	if options.Max != nil && value > *options.Max {
		return picker.NewValidationError(picker.ErrorIDValidationRange, map[string]string{"bound": "max"})
	}
	if options.MultipleOf != nil && *options.MultipleOf != 0 {
		quotient := value / *options.MultipleOf
		if math.Abs(quotient-math.Round(quotient)) > 1e-9 {
			return picker.NewValidationError(picker.ErrorIDValidationNumber, map[string]string{"reason": "multiple_of"})
		}
	}
	return nil
}
