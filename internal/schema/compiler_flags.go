// purpose: Normalize parsed flags into queryable sets used by compiler rules and mapping.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import "github.com/flarebyte/snake-knot-picker"

type FlagSet struct {
	raw map[string][]string
}

func collectFlags(ast *CommandAST) (FlagSet, error) {
	allowed := allowedFlagsForHead(ast.Head)
	flags := make(map[string][]string, len(ast.Flags))
	for _, f := range ast.Flags {
		if _, ok := allowed[f.Name]; !ok {
			return FlagSet{}, picker.NewSchemaError(
				picker.ErrorIDSchemaUnknownFlag,
				map[string]string{"operator": ast.Operator, "flag": f.Name},
			)
		}
		flags[f.Name] = append(flags[f.Name], f.Values...)
	}
	return FlagSet{raw: flags}, nil
}

func (f FlagSet) Has(name string) bool {
	_, ok := f.raw[name]
	return ok
}

func (f FlagSet) First(name string) string {
	values := f.raw[name]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (f FlagSet) FirstOr(name, fallback string) string {
	v := f.First(name)
	if v == "" {
		return fallback
	}
	return v
}

func (f FlagSet) All(name string) []string {
	return append([]string(nil), f.raw[name]...)
}

func (f FlagSet) CloneRaw() map[string][]string {
	out := make(map[string][]string, len(f.raw))
	for k, v := range f.raw {
		out[k] = append([]string(nil), v...)
	}
	return out
}

// Compatibility helpers used by existing schema mapping/tuple files.
func hasFlag(flags map[string][]string, name string) bool {
	return FlagSet{raw: flags}.Has(name)
}

func firstFlagValue(flags map[string][]string, name string) string {
	return FlagSet{raw: flags}.First(name)
}

func firstFlagValueDefault(flags map[string][]string, name, fallback string) string {
	return FlagSet{raw: flags}.FirstOr(name, fallback)
}
