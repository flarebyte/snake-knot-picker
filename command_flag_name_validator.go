// purpose: Build configurable flag-name validators used to enforce naming policy during document compilation.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import (
	"regexp"
	"unicode"
)

// FlagNameValidator validates one flag name and returns a schema error when invalid.
type FlagNameValidator func(name string) error

// CompileOptions configures command document compilation behavior.
type CompileOptions struct {
	FlagNameValidator FlagNameValidator
}

// DefaultCompileOptions returns the default compile options used by public compile entry points.
func DefaultCompileOptions() CompileOptions {
	return CompileOptions{
		FlagNameValidator: DefaultManualFlagNameValidator(),
	}
}

// DefaultManualFlagNameValidator returns the default manual rune-check flag name validator.
func DefaultManualFlagNameValidator() FlagNameValidator {
	return NewManualFlagNameValidator(1, 64, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '.'
	})
}

// DefaultRegexFlagNameValidator returns the default regex-based flag name validator.
func DefaultRegexFlagNameValidator() FlagNameValidator {
	return NewRegexFlagNameValidator(regexp.MustCompile(`^[A-Za-z0-9._-]+$`), 1, 64)
}

// NewRegexFlagNameValidator builds a regex-based validator with length bounds.
func NewRegexFlagNameValidator(re *regexp.Regexp, minLen, maxLen int) FlagNameValidator {
	return func(name string) error {
		if len(name) < minLen || len(name) > maxLen {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
		}
		if !re.MatchString(name) {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
		}
		return nil
	}
}

// NewManualFlagNameValidator builds a rune-predicate validator with length bounds.
func NewManualFlagNameValidator(minLen, maxLen int, allowedRune func(r rune) bool) FlagNameValidator {
	return func(name string) error {
		if len(name) < minLen || len(name) > maxLen {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
		}
		for _, r := range name {
			if !allowedRune(r) {
				return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
			}
		}
		return nil
	}
}
