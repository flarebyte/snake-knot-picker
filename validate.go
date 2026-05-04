package picker

import "strings"

type Runtime struct {
	command CompiledCommand
}

func NewRuntime(command CompiledCommand) (*Runtime, error) {
	return &Runtime{command: command}, nil
}

func Parse(command CompiledCommand, argv []string) (*ParseResult, error) {
	for _, token := range argv {
		if token == "schema" || token == "custom" || strings.HasPrefix(token, "schema ") || strings.HasPrefix(token, "custom ") {
			return nil, &ValidationError{
				Details: []ErrorDetail{
					{
						ID:      ErrorIDValidationSchemaCommandForbidden,
						Kind:    ErrorKindValidation,
						Message: "User argv must not contain schema authoring commands",
					},
				},
			}
		}
	}
	return &ParseResult{
		CommandPath: command.CommandPath,
		Values:      map[string]Value{},
	}, nil
}

func Validate(command CompiledCommand, argv []string) (*ParseResult, error) {
	return Parse(command, argv)
}

