package schema

import "sort"

type CompiledSpec struct {
	Head      string
	Operator  string
	Flags     map[string][]string
	Raw       []string
	TupleSlot *int
}

func buildCompiledSpec(ast *CommandAST, flags FlagSet) *CompiledSpec {
	return &CompiledSpec{
		Head:      ast.Head,
		Operator:  ast.Operator,
		Flags:     flags.CloneRaw(),
		Raw:       append([]string(nil), ast.Raw...),
		TupleSlot: ast.TupleIndex,
	}
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
