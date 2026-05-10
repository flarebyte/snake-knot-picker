// purpose: Define the canonical structured error model and stable IDs used across schema compilation and runtime validation.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import "fmt"

// Schema and validation error identifiers.
const (
	ErrorIDSchemaUnknownOperator            = "schema.unknown_operator"
	ErrorIDSchemaUnknownFlag                = "schema.unknown_flag"
	ErrorIDSchemaMissingValue               = "schema.missing_value"
	ErrorIDSchemaInvalidValue               = "schema.invalid_value"
	ErrorIDSchemaInvalidCombination         = "schema.invalid_combination"
	ErrorIDSchemaDuplicateRegistration      = "schema.duplicate_registration"
	ErrorIDSchemaEnumWhitespace             = "schema.enum_whitespace"
	ErrorIDSchemaEnumEmpty                  = "schema.enum_empty"
	ErrorIDSchemaTupleMissingIndex          = "schema.tuple_missing_index"
	ErrorIDSchemaTupleIndexOutOfRange       = "schema.tuple_index_out_of_range"
	ErrorIDSchemaTupleDuplicateSlot         = "schema.tuple_duplicate_slot"
	ErrorIDSchemaTupleMissingSlot           = "schema.tuple_missing_slot"
	ErrorIDValidationRequired               = "validation.required"
	ErrorIDValidationUnexpectedFlag         = "validation.unexpected_flag"
	ErrorIDValidationSchemaCommandForbidden = "validation.schema_command_forbidden"
	ErrorIDValidationInvalidType            = "validation.invalid_type"
	ErrorIDValidationString                 = "validation.string"
	ErrorIDValidationNumber                 = "validation.number"
	ErrorIDValidationTuple                  = "validation.tuple"
	ErrorIDValidationList                   = "validation.list"
	ErrorIDValidationFormat                 = "validation.format"
	ErrorIDValidationRange                  = "validation.range"
)

// Error kinds used in structured error details.
const (
	ErrorKindSchema     = "schema"
	ErrorKindValidation = "validation"
)

var errorMessageTemplates = map[string]string{
	ErrorIDSchemaUnknownOperator:            "Unknown schema operator",
	ErrorIDSchemaUnknownFlag:                "Unknown schema flag for the selected operator",
	ErrorIDSchemaMissingValue:               "Schema flag requires a following value",
	ErrorIDSchemaInvalidValue:               "Schema flag value is malformed",
	ErrorIDSchemaInvalidCombination:         "Schema flags cannot be used together",
	ErrorIDSchemaDuplicateRegistration:      "Validation operator is already registered",
	ErrorIDSchemaEnumWhitespace:             "Enum value has leading or trailing whitespace",
	ErrorIDSchemaEnumEmpty:                  "Enum value is empty after trimming",
	ErrorIDSchemaTupleMissingIndex:          "Tuple slot schema must include --tuple index",
	ErrorIDSchemaTupleIndexOutOfRange:       "Tuple slot index is outside tuple size",
	ErrorIDSchemaTupleDuplicateSlot:         "Tuple slot index is defined more than once",
	ErrorIDSchemaTupleMissingSlot:           "Tuple schema has at least one slot without validation",
	ErrorIDValidationRequired:               "Required value is missing",
	ErrorIDValidationUnexpectedFlag:         "User argv contains an unknown flag",
	ErrorIDValidationSchemaCommandForbidden: "User argv must not contain schema authoring commands",
	ErrorIDValidationInvalidType:            "User argv value cannot be parsed as the expected type",
	ErrorIDValidationString:                 "Value failed string validation",
	ErrorIDValidationNumber:                 "Value failed number validation",
	ErrorIDValidationTuple:                  "Tuple value failed validation",
	ErrorIDValidationList:                   "List value failed validation",
	ErrorIDValidationFormat:                 "Value failed format validation",
	ErrorIDValidationRange:                  "Value is outside the allowed range",
}

type ErrorDetail struct {
	ID         string
	Kind       string
	Message    string
	Path       []string
	Field      string
	Flag       string
	Operator   string
	TupleIndex *int
	Params     map[string]string
}

// ValidationError aggregates one or more structured validation details.
type ValidationError struct {
	Details []ErrorDetail
}

// NewErrorDetail creates one structured error detail with a rendered message.
func NewErrorDetail(id, kind string, params map[string]string) ErrorDetail {
	detail := ErrorDetail{
		ID:     id,
		Kind:   kind,
		Params: cloneParams(params),
	}
	detail.Message = RenderMessage(id, detail.Params)
	return detail
}

// NewSchemaError creates a schema-kind validation error containing one detail.
func NewSchemaError(id string, params map[string]string) *ValidationError {
	return &ValidationError{
		Details: []ErrorDetail{NewErrorDetail(id, ErrorKindSchema, params)},
	}
}

// NewValidationError creates a validation-kind error containing one detail.
func NewValidationError(id string, params map[string]string) *ValidationError {
	return &ValidationError{
		Details: []ErrorDetail{NewErrorDetail(id, ErrorKindValidation, params)},
	}
}

// Add appends one detail and returns the receiver for chaining.
func (e *ValidationError) Add(detail ErrorDetail) *ValidationError {
	if e == nil {
		return &ValidationError{Details: []ErrorDetail{detail}}
	}
	e.Details = append(e.Details, detail)
	return e
}

// MessageTemplate returns the message template for an error ID.
func MessageTemplate(id string) string {
	return errorMessageTemplates[id]
}

// RenderMessage renders a message for an error ID and params.
func RenderMessage(id string, params map[string]string) string {
	tmpl := MessageTemplate(id)
	if tmpl == "" {
		return "validation failed"
	}
	return tmpl
}

func cloneParams(in map[string]string) map[string]string {
	if len(in) == 0 {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// Error implements the standard error interface.
func (e *ValidationError) Error() string {
	if e == nil || len(e.Details) == 0 {
		return "validation failed"
	}
	d := e.Details[0]
	if d.Message != "" {
		return d.Message
	}
	if d.ID != "" {
		return fmt.Sprintf("%s: validation failed", d.ID)
	}
	return "validation failed"
}
