package picker

import "strconv"

func validateCommandDocument(doc CommandDocument) error {
	if doc.Version == "" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "version"})
	}
	if len(doc.CommandPath) == 0 {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "commandPath"})
	}

	seenNames := map[string]struct{}{}
	for _, flag := range doc.Flags {
		if err := validateFlagDefinition(flag, seenNames); err != nil {
			return err
		}
	}
	return nil
}

func validateFlagDefinition(flag CommandFlagDef, seenNames map[string]struct{}) error {
	if flag.Name == "" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.name"})
	}
	if _, ok := seenNames[flag.Name]; ok {
		return NewSchemaError(ErrorIDSchemaInvalidCombination, map[string]string{"field": "flag.name", "name": flag.Name})
	}
	seenNames[flag.Name] = struct{}{}

	if !isSupportedKind(flag.Kind) {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.kind", "kind": flag.Kind})
	}
	if len(flag.Schema) < 2 {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema", "name": flag.Name})
	}

	shape, err := parsePrimarySchema(flag.Schema)
	if err != nil {
		return err
	}
	if shape.Head != "schema" && shape.Head != "custom" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema.head"})
	}
	if shape.Head == "custom" && shape.Operator == "" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema.operator"})
	}
	if shape.Operator == "tuple" {
		return validateTupleSchemas(shape.TupleSize, flag.Schemas, flag.Name)
	}
	return validateNonTupleSchemas(flag.Schemas, flag.Name)
}

func validateNonTupleSchemas(schemas [][]string, fieldName string) error {
	for _, child := range schemas {
		if len(child) < 2 || child[0] != "schema" || child[1] != "repeatable" {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schemas", "name": fieldName})
		}
	}
	return nil
}

func validateTupleSchemas(tupleSize int, schemas [][]string, fieldName string) error {
	seenSlots := map[int]struct{}{}
	for _, child := range schemas {
		if len(child) < 2 || child[0] != "schema" {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldName})
		}
		if child[1] == "repeatable" {
			continue
		}
		slot, ok, err := parseTupleIndex(child)
		if err != nil {
			return err
		}
		if !ok {
			return NewSchemaError(ErrorIDSchemaTupleMissingIndex, map[string]string{"field": fieldName})
		}
		if slot < 0 || slot >= tupleSize {
			return NewSchemaError(ErrorIDSchemaTupleIndexOutOfRange, map[string]string{"field": fieldName})
		}
		if _, exists := seenSlots[slot]; exists {
			return NewSchemaError(ErrorIDSchemaTupleDuplicateSlot, map[string]string{"field": fieldName})
		}
		seenSlots[slot] = struct{}{}
	}
	for i := 0; i < tupleSize; i++ {
		if _, ok := seenSlots[i]; !ok {
			return NewSchemaError(ErrorIDSchemaTupleMissingSlot, map[string]string{"field": fieldName})
		}
	}
	return nil
}

func parseTupleIndex(tokens []string) (int, bool, error) {
	for i := 2; i < len(tokens); i++ {
		if tokens[i] == "--tuple" {
			if i+1 >= len(tokens) {
				return 0, false, NewSchemaError(ErrorIDSchemaMissingValue, map[string]string{"flag": "--tuple"})
			}
			v, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"flag": "--tuple"})
			}
			return v, true, nil
		}
	}
	return 0, false, nil
}
