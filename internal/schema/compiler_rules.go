package schema

import (
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

func validateCompilerRules(flags FlagSet) error {
	rules := []func(FlagSet) error{
		validateSecureSchemeRule,
		validateEnumDefinitionRule,
	}
	for _, rule := range rules {
		if err := rule(flags); err != nil {
			return err
		}
	}
	return nil
}

func validateSecureSchemeRule(flags FlagSet) error {
	if flags.Has("--secure") && flags.First("--scheme") != "https" {
		return picker.NewSchemaError(
			picker.ErrorIDSchemaInvalidCombination,
			map[string]string{"flag": "--secure", "requires": "--scheme https"},
		)
	}
	return nil
}

func validateEnumDefinitionRule(flags FlagSet) error {
	if !flags.Has("--enum") {
		return nil
	}
	sep := flags.FirstOr("--enum-separator", ",")
	for _, group := range flags.All("--enum") {
		for _, candidate := range splitEnum(group, sep) {
			if candidate == "" {
				return picker.NewSchemaError(picker.ErrorIDSchemaEnumEmpty, nil)
			}
			if candidate != strings.TrimSpace(candidate) {
				return picker.NewSchemaError(picker.ErrorIDSchemaEnumWhitespace, map[string]string{"value": candidate})
			}
		}
	}
	return nil
}

func splitEnum(values, sep string) []string {
	if sep == "" {
		sep = ","
	}
	return strings.Split(values, sep)
}
