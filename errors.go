package picker

import "fmt"

const (
	ErrorIDSchemaUnknownOperator      = "schema.unknown_operator"
	ErrorIDSchemaUnknownFlag          = "schema.unknown_flag"
	ErrorIDSchemaMissingValue         = "schema.missing_value"
	ErrorIDSchemaInvalidValue         = "schema.invalid_value"
	ErrorIDSchemaInvalidCombination   = "schema.invalid_combination"
	ErrorIDSchemaDuplicateRegistration = "schema.duplicate_registration"
	ErrorIDSchemaEnumWhitespace       = "schema.enum_whitespace"
	ErrorIDSchemaEnumEmpty            = "schema.enum_empty"
	ErrorIDSchemaTupleMissingIndex    = "schema.tuple_missing_index"
	ErrorIDSchemaTupleIndexOutOfRange = "schema.tuple_index_out_of_range"
	ErrorIDSchemaTupleDuplicateSlot   = "schema.tuple_duplicate_slot"
	ErrorIDSchemaTupleMissingSlot     = "schema.tuple_missing_slot"
	ErrorIDValidationRequired         = "validation.required"
	ErrorIDValidationUnexpectedFlag   = "validation.unexpected_flag"
	ErrorIDValidationSchemaCommandForbidden = "validation.schema_command_forbidden"
	ErrorIDValidationInvalidType      = "validation.invalid_type"
	ErrorIDValidationString           = "validation.string"
	ErrorIDValidationNumber           = "validation.number"
	ErrorIDValidationTuple            = "validation.tuple"
	ErrorIDValidationList             = "validation.list"
	ErrorIDValidationFormat           = "validation.format"
	ErrorIDValidationRange            = "validation.range"
)

const (
	ErrorKindSchema     = "schema"
	ErrorKindValidation = "validation"
)

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

type ValidationError struct {
	Details []ErrorDetail
}

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

