// purpose: Define argv-specific error-detail helpers for parser error reporting in tests and adapters.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package argv

import "github.com/flarebyte/snake-knot-picker"

type ErrorContext struct {
	Path       []string
	Field      string
	Flag       string
	Operator   string
	TupleIndex *int
}

func NewValidationDetail(id string, ctx ErrorContext, params map[string]string) picker.ErrorDetail {
	detail := picker.NewErrorDetail(id, picker.ErrorKindValidation, params)
	detail.Path = append([]string(nil), ctx.Path...)
	detail.Field = ctx.Field
	detail.Flag = ctx.Flag
	detail.Operator = ctx.Operator
	detail.TupleIndex = ctx.TupleIndex
	return detail
}
