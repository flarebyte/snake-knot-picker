package schema

import "github.com/flarebyte/snake-knot-picker"

type ErrorContext struct {
	Path       []string
	Field      string
	Flag       string
	Operator   string
	TupleIndex *int
}

func NewSchemaDetail(id string, ctx ErrorContext, params map[string]string) picker.ErrorDetail {
	detail := picker.NewErrorDetail(id, picker.ErrorKindSchema, params)
	detail.Path = append([]string(nil), ctx.Path...)
	detail.Field = ctx.Field
	detail.Flag = ctx.Flag
	detail.Operator = ctx.Operator
	detail.TupleIndex = ctx.TupleIndex
	return detail
}
