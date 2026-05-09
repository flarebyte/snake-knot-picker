package picker

import (
	"strconv"
	"strings"
)

type Runtime struct {
	command CompiledCommand
}

func NewRuntime(command CompiledCommand) (*Runtime, error) {
	return &Runtime{command: command}, nil
}

func Parse(command CompiledCommand, argv []string) (*ParseResult, error) {
	values := map[string]Value{}
	flags := map[string]CompiledFlag{}
	for _, f := range command.Flags {
		flags["--"+f.Name] = f
	}

	tokens := stripCommandPath(command.CommandPath, argv)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "schema" || token == "custom" || strings.HasPrefix(token, "schema ") || strings.HasPrefix(token, "custom ") {
			return nil, NewValidationError(ErrorIDValidationSchemaCommandForbidden, nil)
		}
		if !strings.HasPrefix(token, "--") {
			return nil, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"token": token})
		}

		name, inline, hasInline := splitFlag(token)
		def, ok := flags[name]
		if !ok {
			return nil, NewValidationError(ErrorIDValidationUnexpectedFlag, map[string]string{"flag": name})
		}

		var rawValues []string
		if hasInline {
			_ = inline
			return nil, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": name})
		} else if def.Kind == "boolean" {
			rawValues = append(rawValues, "true")
		} else {
			if i+1 >= len(tokens) || strings.HasPrefix(tokens[i+1], "--") {
				return nil, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": name})
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

	return &ParseResult{
		CommandPath: append([]string(nil), command.CommandPath...),
		Values:      values,
	}, nil
}

func Validate(command CompiledCommand, argv []string) (*ParseResult, error) {
	return Parse(command, argv)
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

func parseFlagValue(def CompiledFlag, rawValues []string) (Value, error) {
	if hasFlagLikeValue(rawValues) {
		return Value{}, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
	}

	switch def.Kind {
	case "boolean":
		v := true
		if len(rawValues) > 0 {
			b, err := strconv.ParseBool(rawValues[0])
			if err != nil {
				return Value{}, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
			}
			v = b
		}
		return Value{Bool: &v}, nil
	case "number":
		n, err := strconv.ParseFloat(rawValues[0], 64)
		if err != nil {
			return Value{}, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
		}
		return Value{Number: &n}, nil
	case "tuple":
		parts := splitCSV(rawValues[0])
		if hasFlagLikeValue(parts) {
			return Value{}, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
		}
		if def.TupleSize > 0 && len(parts) != def.TupleSize {
			return Value{}, NewValidationError(ErrorIDValidationTuple, map[string]string{"flag": def.Name})
		}
		tuple := make([]Value, 0, len(parts))
		for _, p := range parts {
			cp := p
			tuple = append(tuple, Value{String: &cp})
		}
		return Value{Tuple: tuple}, nil
	default:
		s := rawValues[0]
		if def.Repeatable {
			parts := splitCSV(s)
			if hasFlagLikeValue(parts) {
				return Value{}, NewValidationError(ErrorIDValidationInvalidType, map[string]string{"flag": def.Name, "kind": def.Kind})
			}
			list := make([]Value, 0, len(parts))
			for _, p := range parts {
				cp := p
				list = append(list, Value{String: &cp})
			}
			return Value{List: list}, nil
		}
		return Value{String: &s}, nil
	}
}

func hasFlagLikeValue(values []string) bool {
	for _, v := range values {
		if strings.HasPrefix(strings.TrimSpace(v), "--") {
			return true
		}
	}
	return false
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

func flattenList(v Value) []Value {
	if len(v.List) > 0 {
		return v.List
	}
	return []Value{v}
}
