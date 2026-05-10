// purpose: Bridge schema-layer errors into standardized picker validation detail structures.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import "github.com/flarebyte/snake-knot-picker"

// ErrorContext carries schema-layer context fields for structured error details.
type ErrorContext struct {
	Path       []string
	Field      string
	Flag       string
	Operator   string
	TupleIndex *int
}

// NewSchemaDetail builds a schema-kind error detail with context attached.
func NewSchemaDetail(id string, ctx ErrorContext, params map[string]string) picker.ErrorDetail {
	detail := picker.NewErrorDetail(id, picker.ErrorKindSchema, params)
	detail.Path = append([]string(nil), ctx.Path...)
	detail.Field = ctx.Field
	detail.Flag = ctx.Flag
	detail.Operator = ctx.Operator
	detail.TupleIndex = ctx.TupleIndex
	return detail
}
