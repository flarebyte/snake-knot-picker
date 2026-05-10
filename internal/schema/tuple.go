// purpose: Compile tuple primary and child schema specs into validated tuple configuration.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import (
	"strconv"

	"github.com/flarebyte/snake-knot-picker"
)

type TupleSpec struct {
	Size       int
	Required   bool
	Slots      map[int]*CompiledSpec
	Repeatable *CompiledSpec
}

func CompileTupleSpec(primary *CompiledSpec, children []*CompiledSpec, field string) (*TupleSpec, error) {
	if primary == nil || primary.Operator != "tuple" {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": field, "operator": "tuple"})
	}
	sizeRaw := firstFlagValue(primary.Flags, "--size")
	if sizeRaw == "" {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": field, "flag": "--size"})
	}
	size, err := strconv.Atoi(sizeRaw)
	if err != nil || size < 0 {
		return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": field, "flag": "--size", "value": sizeRaw})
	}

	out := &TupleSpec{
		Size:     size,
		Required: hasFlag(primary.Flags, "--required"),
		Slots:    map[int]*CompiledSpec{},
	}

	for _, child := range children {
		if child == nil {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": field})
		}
		if child.Operator == "repeatable" {
			if out.Repeatable != nil {
				return nil, picker.NewSchemaError(picker.ErrorIDSchemaInvalidCombination, map[string]string{"field": field, "operator": "repeatable"})
			}
			out.Repeatable = child.Clone()
			continue
		}
		if child.TupleSlot == nil {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaTupleMissingIndex, map[string]string{"field": field, "operator": child.Operator})
		}
		idx := *child.TupleSlot
		if idx < 0 || idx >= size {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaTupleIndexOutOfRange, map[string]string{"field": field, "index": strconv.Itoa(idx), "size": strconv.Itoa(size)})
		}
		if _, exists := out.Slots[idx]; exists {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaTupleDuplicateSlot, map[string]string{"field": field, "index": strconv.Itoa(idx)})
		}
		out.Slots[idx] = child.Clone()
	}

	for i := 0; i < size; i++ {
		if _, ok := out.Slots[i]; !ok {
			return nil, picker.NewSchemaError(picker.ErrorIDSchemaTupleMissingSlot, map[string]string{"field": field, "index": strconv.Itoa(i)})
		}
	}
	return out, nil
}
