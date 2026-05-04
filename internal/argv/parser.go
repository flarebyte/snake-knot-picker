package argv

import (
	"strconv"
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

type Parser struct{}

func NewParser() *Parser { return &Parser{} }

func (p *Parser) Parse(command picker.CompiledCommand, argv []string) (*picker.ParseResult, error) {
	values := map[string]picker.Value{}
	flags := map[string]picker.CompiledFlag{}
	for _, f := range command.Flags {
		flags["--"+f.Name] = f
	}

	tokens := stripCommandPath(command.CommandPath, argv)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "schema" || token == "custom" || strings.HasPrefix(token, "schema ") || strings.HasPrefix(token, "custom ") {
			return nil, picker.NewValidationError(picker.ErrorIDValidationSchemaCommandForbidden, nil)
		}
		if !strings.HasPrefix(token, "--") {
			return nil, picker.NewValidationError(picker.ErrorIDValidationInvalidType, map[string]string{"token": token})
		}

		name, inline, hasInline := splitFlag(token)
		def, ok := flags[name]
		if !ok {
			return nil, picker.NewValidationError(picker.ErrorIDValidationUnexpectedFlag, map[string]string{"flag": name})
		}

		var rawValues []string
		if hasInline {
			rawValues = append(rawValues, inline)
		} else if def.Kind == "boolean" {
			rawValues = append(rawValues, "true")
		} else {
			if i+1 >= len(tokens) || strings.HasPrefix(tokens[i+1], "--") {
				return nil, picker.NewValidationError(picker.ErrorIDValidationInvalidType, map[string]string{"flag": name})
			}
			i++
			rawValues = append(rawValues, tokens[i])
		}

		parsed, err := parseFlagValue(def, rawValues)
		if err != nil {
			return nil, err
		}
		if def.Repeatable {
			existing := values[def.Name]
			existing.List = append(existing.List, flattenList(parsed)...)
			values[def.Name] = existing
			continue
		}
		values[def.Name] = parsed
	}

	return &picker.ParseResult{
		CommandPath: append([]string(nil), command.CommandPath...),
		Values:      values,
	}, nil
}

func stripCommandPath(path, argv []string) []string {
	if len(path) == 0 || len(argv) < len(path) {
		return append([]string(nil), argv...)
	}
	for i := range path {
		if argv[i] != path[i] {
			return append([]string(nil), argv...)
		}
	}
	return append([]string(nil), argv[len(path):]...)
}

func splitFlag(token string) (name, inline string, hasInline bool) {
	if idx := strings.Index(token, "="); idx > 0 {
		return token[:idx], token[idx+1:], true
	}
	return token, "", false
}

func parseFlagValue(def picker.CompiledFlag, rawValues []string) (picker.Value, error) {
	switch def.Kind {
	case "boolean":
		v := true
		if len(rawValues) > 0 {
			b, err := strconv.ParseBool(rawValues[0])
			if err != nil {
				return picker.Value{}, picker.NewValidationError(picker.ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
			}
			v = b
		}
		return picker.Value{Bool: &v}, nil
	case "number":
		n, err := strconv.ParseFloat(rawValues[0], 64)
		if err != nil {
			return picker.Value{}, picker.NewValidationError(picker.ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
		}
		return picker.Value{Number: &n}, nil
	case "tuple":
		parts := splitCSV(rawValues[0])
		if def.TupleSize > 0 && len(parts) != def.TupleSize {
			return picker.Value{}, picker.NewValidationError(picker.ErrorIDValidationTuple, map[string]string{"flag": def.Name})
		}
		tuple := make([]picker.Value, 0, len(parts))
		for _, p := range parts {
			cp := p
			tuple = append(tuple, picker.Value{String: &cp})
		}
		return picker.Value{Tuple: tuple}, nil
	default:
		s := rawValues[0]
		if def.Repeatable {
			parts := splitCSV(s)
			list := make([]picker.Value, 0, len(parts))
			for _, p := range parts {
				cp := p
				list = append(list, picker.Value{String: &cp})
			}
			return picker.Value{List: list}, nil
		}
		return picker.Value{String: &s}, nil
	}
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	if len(out) == 0 {
		return []string{value}
	}
	return out
}

func flattenList(v picker.Value) []picker.Value {
	if len(v.List) > 0 {
		return v.List
	}
	return []picker.Value{v}
}

