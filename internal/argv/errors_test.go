// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package argv

import (
	"testing"

	"github.com/flarebyte/snake-knot-picker"
)

func TestNewValidationDetailPreservesContext(t *testing.T) {
	tupleIndex := 0
	path := []string{"wash", "start", "flags", "range", "list", "2"}
	detail := NewValidationDetail(
		picker.ErrorIDValidationList,
		ErrorContext{
			Path:       path,
			Field:      "range",
			Flag:       "--range",
			Operator:   "list",
			TupleIndex: &tupleIndex,
		},
		map[string]string{"index": "2"},
	)
	if detail.Kind != picker.ErrorKindValidation {
		t.Fatalf("unexpected kind: %s", detail.Kind)
	}
	if detail.ID != picker.ErrorIDValidationList {
		t.Fatalf("unexpected id: %s", detail.ID)
	}
	if detail.Field != "range" || detail.Flag != "--range" || detail.Operator != "list" {
		t.Fatalf("unexpected context: %#v", detail)
	}
	if detail.TupleIndex == nil || *detail.TupleIndex != 0 {
		t.Fatalf("unexpected tuple index: %#v", detail.TupleIndex)
	}
	if len(detail.Path) != len(path) {
		t.Fatalf("unexpected path: %#v", detail.Path)
	}
}
