package schema

import (
	"strconv"
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

var knownFlagArity = map[string]int{
	"--required":            0,
	"--int":                 0,
	"--secure":              0,
	"--allow-query":         0,
	"--allow-alpha":         0,
	"--alphabetic":          0,
	"--hexa":                0,
	"--whitespace":          0,
	"--lowercase":           0,
	"--uppercase":           0,
	"--punctuation":         0,
	"--blank":               0,
	"--unicode-letter":      0,
	"--unicode-number":      0,
	"--unicode-punctuation": 0,
	"--unicode-symbol":      0,
	"--unicode-separator":   0,
	"--latin":               0,
	"--han":                 0,
	"--devanagari":          0,
	"--arabic":              0,
	"--hiragana":            0,
	"--katakana":            0,
	"--hangul":              0,
	"--tamil":               0,
	"--gujarati":            0,
	"--ethiopic":            0,
	"--email":               0,
	"--datetime":            0,
	"--duration":            0,
	"--base64":              0,
	"--boolean":             0,
	"--color":               0,
	"--date":                0,
	"--time":                0,
	"--uri":                 0,
	"--arn":                 0,
	"--enum":                1,
	"--enum-separator":      1,
	"--size":                1,
	"--tuple":               1,
	"--min-length":          1,
	"--max-length":          1,
	"--scheme":              1,
	"--allow-domains":       1,
	"--allow-partition":     1,
	"--allow-service":       1,
	"--allow-region":        1,
	"--allow-account-id":    1,
	"--allow-resource":      1,
	"--country":             1,
	"--layout":              1,
	"--min-duration":        1,
	"--max-duration":        1,
	"--starts-with":         1,
	"--format":              1,
	"--codepoint-range":     2,
}

func ParseTokens(tokens []string) (*CommandAST, error) {
	if len(tokens) < 2 {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, nil)
	}
	head := tokens[0]
	if head != "schema" && head != "custom" {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"head": head})
	}

	ast := &CommandAST{
		Head:     head,
		Operator: tokens[1],
		Raw:      append([]string(nil), tokens...),
	}
	if ast.Operator == "" {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"operator": ast.Operator})
	}

	flagsByName := map[string]int{}
	for i := 2; i < len(tokens); {
		token := tokens[i]
		if !isLongFlag(token) {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"token": token})
		}
		values, next, err := parseFlagValues(tokens, i, token)
		if err != nil {
			return nil, err
		}
		idx, ok := flagsByName[token]
		if !ok {
			ast.Flags = append(ast.Flags, ParsedSchemaFlag{Name: token})
			idx = len(ast.Flags) - 1
			flagsByName[token] = idx
		}
		ast.Flags[idx].Values = append(ast.Flags[idx].Values, values...)
		i = next
	}

	if tupleValues := findFlagValues(ast.Flags, "--tuple"); len(tupleValues) > 0 {
		if len(tupleValues) != 1 {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"flag": "--tuple"})
		}
		n, err := strconv.Atoi(tupleValues[0])
		if err != nil || n < 0 {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"flag": "--tuple", "value": tupleValues[0]})
		}
		ast.TupleIndex = &n
	}

	return ast, nil
}

func parseFlagValues(tokens []string, i int, flag string) ([]string, int, error) {
	if arity, ok := knownFlagArity[flag]; ok {
		if arity == 0 {
			return nil, i + 1, nil
		}
		if i+arity >= len(tokens) {
			return nil, 0, picker.NewSchemaError(picker.ErrorIDSchemaMissingValue, map[string]string{"flag": flag})
		}
		values := make([]string, 0, arity)
		for j := 1; j <= arity; j++ {
			value := tokens[i+j]
			if isLongFlag(value) {
				return nil, 0, picker.NewSchemaError(picker.ErrorIDSchemaMissingValue, map[string]string{"flag": flag})
			}
			values = append(values, value)
		}
		return values, i + 1 + arity, nil
	}

	j := i + 1
	for j < len(tokens) && !isLongFlag(tokens[j]) {
		j++
	}
	if j == i+1 {
		return nil, i + 1, nil
	}
	return append([]string(nil), tokens[i+1:j]...), j, nil
}

func isLongFlag(token string) bool {
	return strings.HasPrefix(token, "--")
}

func findFlagValues(flags []ParsedSchemaFlag, name string) []string {
	for _, f := range flags {
		if f.Name == name {
			return f.Values
		}
	}
	return nil
}

