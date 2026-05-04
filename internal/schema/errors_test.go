package schema

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestNewSchemaDetailPreservesContext(t *testing.T) {
	tupleIndex := 1
	path := []string{"wash", "start", "flags", "pair", "tuple", "1"}
	detail := NewSchemaDetail(
		picker.ErrorIDSchemaTupleDuplicateSlot,
		ErrorContext{
			Path:       path,
			Field:      "pair",
			Flag:       "--tuple",
			Operator:   "tuple",
			TupleIndex: &tupleIndex,
		},
		map[string]string{"slot": "1"},
	)
	if detail.Kind != picker.ErrorKindSchema {
		t.Fatalf("unexpected kind: %s", detail.Kind)
	}
	if detail.ID != picker.ErrorIDSchemaTupleDuplicateSlot {
		t.Fatalf("unexpected id: %s", detail.ID)
	}
	if detail.Field != "pair" || detail.Flag != "--tuple" || detail.Operator != "tuple" {
		t.Fatalf("unexpected context: %#v", detail)
	}
	if detail.TupleIndex == nil || *detail.TupleIndex != 1 {
		t.Fatalf("unexpected tuple index: %#v", detail.TupleIndex)
	}
	if len(detail.Path) != len(path) {
		t.Fatalf("unexpected path: %#v", detail.Path)
	}
}

