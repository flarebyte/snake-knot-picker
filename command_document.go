package picker

import (
	"encoding/json"
	"strconv"
)

func ParseCommandDocumentJSON(data []byte) (CommandDocument, error) {
	var doc CommandDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return CommandDocument{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "json"})
	}
	return doc, nil
}

func (d CommandDocument) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}

func CompileCommandDocument(doc CommandDocument) (CompiledCommand, error) {
	if doc.Version == "" {
		return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "version"})
	}
	if len(doc.CommandPath) == 0 {
		return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "commandPath"})
	}

	seen := map[string]struct{}{}
	out := CompiledCommand{
		CommandPath: append([]string(nil), doc.CommandPath...),
		AdminOnly:   doc.AdminOnly,
		Flags:       make([]CompiledFlag, 0, len(doc.Flags)),
	}

	for _, flag := range doc.Flags {
		if flag.Name == "" {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.name"})
		}
		if _, ok := seen[flag.Name]; ok {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidCombination, map[string]string{"field": "flag.name", "name": flag.Name})
		}
		seen[flag.Name] = struct{}{}

		if !isSupportedKind(flag.Kind) {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.kind", "kind": flag.Kind})
		}
		if len(flag.Schema) < 2 {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema", "name": flag.Name})
		}

		head, operator, tupleSize, err := parsePrimarySchema(flag.Schema)
		if err != nil {
			return CompiledCommand{}, err
		}
		if head != "schema" && head != "custom" {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema.head"})
		}
		if head == "custom" && operator == "" {
			return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema.operator"})
		}

		cf := CompiledFlag{
			Kind:      flag.Kind,
			Name:      flag.Name,
			TupleSize: tupleSize,
		}
		if operator == "tuple" {
			if tupleSize < 0 {
				return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schema.--size"})
			}
			if err := validateTupleSchemas(tupleSize, flag.Schemas, flag.Name); err != nil {
				return CompiledCommand{}, err
			}
			for _, child := range flag.Schemas {
				if len(child) >= 2 && child[0] == "schema" && child[1] == "repeatable" {
					cf.Repeatable = true
				}
			}
		} else {
			for _, child := range flag.Schemas {
				if len(child) < 2 || child[0] != "schema" || child[1] != "repeatable" {
					return CompiledCommand{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "flag.schemas", "name": flag.Name})
				}
				cf.Repeatable = true
			}
		}
		out.Flags = append(out.Flags, cf)
	}

	return out, nil
}

func parsePrimarySchema(tokens []string) (head, operator string, tupleSize int, err error) {
	tupleSize = 0
	head = tokens[0]
	operator = tokens[1]
	if operator == "tuple" {
		sizeRaw := ""
		for i := 2; i < len(tokens); i++ {
			if tokens[i] == "--size" {
				if i+1 >= len(tokens) {
					return "", "", 0, NewSchemaError(ErrorIDSchemaMissingValue, map[string]string{"flag": "--size"})
				}
				sizeRaw = tokens[i+1]
				break
			}
		}
		if sizeRaw == "" {
			return "", "", 0, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "tuple.size"})
		}
		size, parseErr := strconv.Atoi(sizeRaw)
		if parseErr != nil || size < 0 {
			return "", "", 0, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "tuple.size"})
		}
		tupleSize = size
	}
	return head, operator, tupleSize, nil
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

func isSupportedKind(kind string) bool {
	switch kind {
	case "boolean", "string", "number", "tuple":
		return true
	default:
		return false
	}
}

