package picker

import "testing"

func TestErrorIDsAreUnique(t *testing.T) {
	ids := []string{
		ErrorIDSchemaUnknownOperator,
		ErrorIDSchemaUnknownFlag,
		ErrorIDSchemaMissingValue,
		ErrorIDSchemaInvalidValue,
		ErrorIDSchemaInvalidCombination,
		ErrorIDSchemaDuplicateRegistration,
		ErrorIDSchemaEnumWhitespace,
		ErrorIDSchemaEnumEmpty,
		ErrorIDSchemaTupleMissingIndex,
		ErrorIDSchemaTupleIndexOutOfRange,
		ErrorIDSchemaTupleDuplicateSlot,
		ErrorIDSchemaTupleMissingSlot,
		ErrorIDValidationRequired,
		ErrorIDValidationUnexpectedFlag,
		ErrorIDValidationSchemaCommandForbidden,
		ErrorIDValidationInvalidType,
		ErrorIDValidationString,
		ErrorIDValidationNumber,
		ErrorIDValidationTuple,
		ErrorIDValidationList,
		ErrorIDValidationFormat,
		ErrorIDValidationRange,
	}
	seen := map[string]struct{}{}
	for _, id := range ids {
		if id == "" {
			t.Fatal("error id must not be empty")
		}
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate error id: %s", id)
		}
		seen[id] = struct{}{}
	}
}

func TestPublicContractsCompile(t *testing.T) {
	builder := NewCommandBuilder("wash", "start").SetAdminOnly(true).AddFlag(
		CommandFlagDef{
			Kind:   "string",
			Name:   "mode",
			Schema: []string{"schema", "string", "--required"},
		},
	)
	doc := builder.Build()
	if len(doc.CommandPath) != 2 {
		t.Fatalf("unexpected command path: %#v", doc.CommandPath)
	}

	cmd := CompiledCommand{
		CommandPath: []string{"wash", "start"},
		AdminOnly:   true,
		Flags: []CompiledFlag{
			{Kind: "string", Name: "mode"},
		},
	}
	if _, err := NewRuntime(cmd); err != nil {
		t.Fatalf("unexpected runtime constructor error: %v", err)
	}
}

func TestParseRejectsSchemaCommandsAtRuntime(t *testing.T) {
	cmd := CompiledCommand{CommandPath: []string{"wash", "start"}}

	result, err := Parse(cmd, []string{"schema", "string", "--required"})
	if result != nil {
		t.Fatalf("expected nil result, got %#v", result)
	}
	if err == nil {
		t.Fatal("expected error")
	}

	verr, ok := err.(*ValidationError)
	if !ok {
		t.Fatalf("expected ValidationError, got %T", err)
	}
	if len(verr.Details) != 1 {
		t.Fatalf("expected one detail, got %d", len(verr.Details))
	}
	if verr.Details[0].ID != ErrorIDValidationSchemaCommandForbidden {
		t.Fatalf("unexpected error id: %s", verr.Details[0].ID)
	}
}

