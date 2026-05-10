// purpose: Enforce structural and semantic constraints on command documents before compilation.
// responsibilities: Validate document fields, check flag-name policy, enforce kind/schema invariants, and verify tuple/repeatable child schema semantics.
// architecture notes: Tuple child schemas are validated by explicit slot indexing and repeatable children are handled as tuple modifiers, not slot validators.
package picker

import "strconv"

const (
	fieldVersion            = "version"
	fieldCommandPath        = "commandPath"
	fieldFlagName           = "flag.name"
	fieldFlagKind           = "flag.kind"
	fieldFlagSchema         = "flag.schema"
	fieldFlagSchemaHead     = "flag.schema.head"
	fieldFlagSchemaOperator = "flag.schema.operator"
	fieldFlagSchemas        = "flag.schemas"
)

type flagCtx struct {
	name string
	kind string
}

type childShape struct {
	head     string
	operator string
	slot     int
	hasSlot  bool
}

func validateCommandDocument(doc CommandDocument, flagNameValidator FlagNameValidator) error {
	if doc.Version == "" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldVersion})
	}
	if len(doc.CommandPath) == 0 {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldCommandPath})
	}

	seenNames := map[string]struct{}{}
	for _, flag := range doc.Flags {
		ctx, shape, err := validateFlagStructure(flag, seenNames, flagNameValidator)
		if err != nil {
			return err
		}
		if err := validateFlagSemantics(ctx, shape, flag.Schemas); err != nil {
			return err
		}
	}
	return nil
}

func validateFlagStructure(flag CommandFlagDef, seenNames map[string]struct{}, flagNameValidator FlagNameValidator) (flagCtx, primarySchemaShape, error) {
	ctx := flagCtx{name: flag.Name, kind: flag.Kind}
	if flag.Name == "" {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName})
	}
	if _, ok := seenNames[flag.Name]; ok {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidCombination, map[string]string{"field": fieldFlagName, "name": flag.Name})
	}
	if err := flagNameValidator(flag.Name); err != nil {
		return ctx, primarySchemaShape{}, err
	}
	seenNames[flag.Name] = struct{}{}

	if !isSupportedKind(flag.Kind) {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagKind, "kind": flag.Kind})
	}
	if len(flag.Schema) < 2 {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagSchema, "name": flag.Name})
	}

	shape, err := parsePrimarySchema(flag.Schema)
	if err != nil {
		return ctx, primarySchemaShape{}, err
	}
	if shape.Head != "schema" && shape.Head != "custom" {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagSchemaHead})
	}
	if shape.Head == "custom" && shape.Operator == "" {
		return ctx, primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagSchemaOperator})
	}
	return ctx, shape, nil
}

func validateFlagSemantics(ctx flagCtx, shape primarySchemaShape, schemas [][]string) error {
	if shape.Operator == "tuple" {
		return validateTupleSchemas(shape.TupleSize, schemas, ctx.name)
	}
	return validateNonTupleSchemas(schemas, ctx.name)
}

func validateNonTupleSchemas(schemas [][]string, fieldName string) error {
	for _, child := range schemas {
		shape, err := parseChildSchema(child)
		if err != nil {
			return err
		}
		if shape.head != "schema" || shape.operator != "repeatable" {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagSchemas, "name": fieldName})
		}
	}
	return nil
}

func validateTupleSchemas(tupleSize int, schemas [][]string, fieldName string) error {
	seenSlots := map[int]struct{}{}
	for _, child := range schemas {
		shape, err := parseChildSchema(child)
		if err != nil {
			return err
		}
		if shape.head != "schema" {
			return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldName})
		}
		if shape.operator == "repeatable" {
			continue
		}
		if !shape.hasSlot {
			return NewSchemaError(ErrorIDSchemaTupleMissingIndex, map[string]string{"field": fieldName})
		}
		if shape.slot < 0 || shape.slot >= tupleSize {
			return NewSchemaError(ErrorIDSchemaTupleIndexOutOfRange, map[string]string{"field": fieldName})
		}
		if _, exists := seenSlots[shape.slot]; exists {
			return NewSchemaError(ErrorIDSchemaTupleDuplicateSlot, map[string]string{"field": fieldName})
		}
		seenSlots[shape.slot] = struct{}{}
	}
	for i := 0; i < tupleSize; i++ {
		if _, ok := seenSlots[i]; !ok {
			return NewSchemaError(ErrorIDSchemaTupleMissingSlot, map[string]string{"field": fieldName})
		}
	}
	return nil
}

func parseChildSchema(tokens []string) (childShape, error) {
	if len(tokens) < 2 {
		return childShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagSchemas})
	}
	slot, hasSlot, err := parseTupleIndex(tokens)
	if err != nil {
		return childShape{}, err
	}
	return childShape{
		head:     tokens[0],
		operator: tokens[1],
		slot:     slot,
		hasSlot:  hasSlot,
	}, nil
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
