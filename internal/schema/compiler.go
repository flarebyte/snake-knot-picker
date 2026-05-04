package schema

import (
	"sort"

	"github.com/flarebyte/snake-knot-picker"
)

type CompiledSpec struct {
	Head      string
	Operator  string
	Flags     map[string][]string
	Raw       []string
	TupleSlot *int
}

type Compiler struct {
	registry picker.Registry
}

func NewCompiler(registry picker.Registry) *Compiler {
	return &Compiler{registry: registry}
}

func (c *Compiler) Compile(ast *CommandAST) (*CompiledSpec, error) {
	if ast == nil {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": "ast"})
	}
	if c.registry == nil {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": "registry"})
	}
	if _, ok := c.registry.Lookup(ast.Operator); !ok {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaUnknownOperator, map[string]string{"operator": ast.Operator})
	}

	allowed := allowedFlags(ast.Operator, ast.Head)
	flags := make(map[string][]string, len(ast.Flags))
	for _, f := range ast.Flags {
		if _, ok := allowed[f.Name]; !ok {
			return nil, picker.NewSchemaError(
				picker.ErrorIDSchemaUnknownFlag,
				map[string]string{"operator": ast.Operator, "flag": f.Name},
			)
		}
		flags[f.Name] = append(flags[f.Name], f.Values...)
	}

	// Minimal combination checks as scaffolding until full validator-specific compiler.
	if hasFlag(flags, "--secure") && firstFlagValue(flags, "--scheme") != "https" {
		return nil, picker.NewSchemaError(
			picker.ErrorIDSchemaInvalidCombination,
			map[string]string{"flag": "--secure", "requires": "--scheme https"},
		)
	}
	if hasFlag(flags, "--enum") {
		for _, group := range flags["--enum"] {
			for _, candidate := range splitEnum(group, firstFlagValueDefault(flags, "--enum-separator", ",")) {
				if candidate == "" {
					return nil, picker.NewSchemaError(picker.ErrorIDSchemaEnumEmpty, nil)
				}
				if candidate != trimSpaces(candidate) {
					return nil, picker.NewSchemaError(picker.ErrorIDSchemaEnumWhitespace, map[string]string{"value": candidate})
				}
			}
		}
	}

	out := &CompiledSpec{
		Head:      ast.Head,
		Operator:  ast.Operator,
		Flags:     flags,
		Raw:       append([]string(nil), ast.Raw...),
		TupleSlot: ast.TupleIndex,
	}
	return out, nil
}

func allowedFlags(operator, head string) map[string]struct{} {
	if head == "custom" {
		return map[string]struct{}{
			"--required": {},
			"--country":  {},
		}
	}
	base := map[string]struct{}{
		"--required":            {},
		"--tuple":               {},
		"--enum":                {},
		"--enum-separator":      {},
		"--size":                {},
		"--int":                 {},
		"--min-length":          {},
		"--max-length":          {},
		"--scheme":              {},
		"--secure":              {},
		"--allow-query":         {},
		"--allow-domains":       {},
		"--color":               {},
		"--format":              {},
		"--allow-alpha":         {},
		"--arn":                 {},
		"--allow-partition":     {},
		"--allow-service":       {},
		"--allow-region":        {},
		"--allow-account-id":    {},
		"--allow-resource":      {},
		"--whitespace":          {},
		"--alphabetic":          {},
		"--lowercase":           {},
		"--uppercase":           {},
		"--punctuation":         {},
		"--blank":               {},
		"--unicode-letter":      {},
		"--unicode-number":      {},
		"--unicode-punctuation": {},
		"--unicode-symbol":      {},
		"--unicode-separator":   {},
		"--latin":               {},
		"--han":                 {},
		"--devanagari":          {},
		"--arabic":              {},
		"--hiragana":            {},
		"--katakana":            {},
		"--hangul":              {},
		"--tamil":               {},
		"--gujarati":            {},
		"--ethiopic":            {},
		"--email":               {},
		"--datetime":            {},
		"--layout":              {},
		"--allow-timezone":      {},
		"--location":            {},
		"--duration":            {},
		"--min-duration":        {},
		"--max-duration":        {},
		"--base64":              {},
		"--starts-with":         {},
		"--codepoint-range":     {},
		"--hexa":                {},
		"--date":                {},
		"--boolean":             {},
		"--time":                {},
		"--uri":                 {},
	}
	_ = operator
	return base
}

func hasFlag(flags map[string][]string, name string) bool {
	_, ok := flags[name]
	return ok
}

func firstFlagValue(flags map[string][]string, name string) string {
	values := flags[name]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func firstFlagValueDefault(flags map[string][]string, name, fallback string) string {
	v := firstFlagValue(flags, name)
	if v == "" {
		return fallback
	}
	return v
}

func splitEnum(values, sep string) []string {
	if sep == "" {
		sep = ","
	}
	out := []string{}
	start := 0
	for i := 0; i <= len(values); i++ {
		if i == len(values) || string(values[i]) == sep {
			out = append(out, values[start:i])
			start = i + len(sep)
		}
	}
	if sep != "," {
		parts := []string{}
		cursor := 0
		for {
			idx := -1
			for i := cursor; i+len(sep) <= len(values); i++ {
				if values[i:i+len(sep)] == sep {
					idx = i
					break
				}
			}
			if idx < 0 {
				parts = append(parts, values[cursor:])
				break
			}
			parts = append(parts, values[cursor:idx])
			cursor = idx + len(sep)
		}
		return parts
	}
	return out
}

func trimSpaces(s string) string {
	r := []rune(s)
	start := 0
	for start < len(r) && (r[start] == ' ' || r[start] == '\t') {
		start++
	}
	end := len(r)
	for end > start && (r[end-1] == ' ' || r[end-1] == '\t') {
		end--
	}
	return string(r[start:end])
}

func (s *CompiledSpec) Clone() *CompiledSpec {
	if s == nil {
		return nil
	}
	out := &CompiledSpec{
		Head:     s.Head,
		Operator: s.Operator,
		Raw:      append([]string(nil), s.Raw...),
	}
	if s.TupleSlot != nil {
		v := *s.TupleSlot
		out.TupleSlot = &v
	}
	out.Flags = make(map[string][]string, len(s.Flags))
	keys := make([]string, 0, len(s.Flags))
	for k := range s.Flags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out.Flags[k] = append([]string(nil), s.Flags[k]...)
	}
	return out
}
