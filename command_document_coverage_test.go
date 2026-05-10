package picker

import (
	"testing"
)

func TestParseCommandDocumentJSONInvalid(t *testing.T) {
	_, err := ParseCommandDocumentJSON([]byte(`{"version":`))
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func TestValidateWithDocumentJSONInvalidDocument(t *testing.T) {
	_, err := ValidateWithDocumentJSON([]byte(`{"version":`), []string{"wash", "start"})
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func TestCompileCommandDocumentWithOptionsDefaultsWhenNilValidator(t *testing.T) {
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string", "--required"}},
		},
	}
	compiled, err := CompileCommandDocumentWithOptions(doc, CompileOptions{})
	if err != nil {
		t.Fatalf("unexpected compile error: %v", err)
	}
	if len(compiled.Flags) != 1 || compiled.Flags[0].Name != "mode" {
		t.Fatalf("unexpected compiled flags: %#v", compiled.Flags)
	}
}

func TestCompileCommandDocumentWithOptionsCustomValidatorSuccess(t *testing.T) {
	doc := CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string", "--required"}},
		},
	}
	custom := func(name string) error {
		if name == "mode" {
			return nil
		}
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": fieldFlagName, "name": name})
	}
	compiled, err := CompileCommandDocumentWithOptions(doc, CompileOptions{FlagNameValidator: custom})
	if err != nil {
		t.Fatalf("unexpected compile error: %v", err)
	}
	if len(compiled.Flags) != 1 || compiled.Flags[0].Name != "mode" {
		t.Fatalf("unexpected compiled result: %#v", compiled)
	}
}

func TestCompileFlagRepeatableDetection(t *testing.T) {
	flag := CommandFlagDef{
		Kind:   "string",
		Name:   "add",
		Schema: []string{"schema", "string"},
		Schemas: [][]string{
			{"schema", "repeatable", "--min-length", "1", "--max-length", "3"},
		},
	}
	cf, err := compileFlag(flag)
	if err != nil {
		t.Fatalf("unexpected compile flag error: %v", err)
	}
	if !cf.Repeatable {
		t.Fatal("expected repeatable=true")
	}
}

func TestCompileFlagErrorPath(t *testing.T) {
	_, err := compileFlag(CommandFlagDef{
		Kind:   "tuple",
		Name:   "range",
		Schema: []string{"schema", "tuple"},
	})
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func TestParsePrimarySchemaTupleEdgeCases(t *testing.T) {
	_, errMissing := parsePrimarySchema([]string{"schema", "tuple", "--size"})
	assertDocErrID(t, errMissing, ErrorIDSchemaMissingValue)

	_, errInvalid := parsePrimarySchema([]string{"schema", "tuple", "--size", "-1"})
	assertDocErrID(t, errInvalid, ErrorIDSchemaInvalidValue)

	_, errNoSize := parsePrimarySchema([]string{"schema", "tuple"})
	assertDocErrID(t, errNoSize, ErrorIDSchemaInvalidValue)
}

func TestValidateFlagStructureCustomAndSchemaHeadRules(t *testing.T) {
	seen := map[string]struct{}{}
	_, _, errCustomMissingOp := validateFlagStructure(
		CommandFlagDef{
			Kind:   "string",
			Name:   "postal",
			Schema: []string{"custom", ""},
		},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errCustomMissingOp, ErrorIDSchemaInvalidValue)

	seen = map[string]struct{}{}
	_, _, errBadHead := validateFlagStructure(
		CommandFlagDef{
			Kind:   "string",
			Name:   "mode",
			Schema: []string{"xhead", "string"},
		},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errBadHead, ErrorIDSchemaInvalidValue)
}

func TestValidateFlagStructureExtraBranches(t *testing.T) {
	seen := map[string]struct{}{}
	_, _, errNameEmpty := validateFlagStructure(
		CommandFlagDef{Kind: "string", Name: "", Schema: []string{"schema", "string"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errNameEmpty, ErrorIDSchemaInvalidValue)

	seen = map[string]struct{}{}
	_, _, errShortSchema := validateFlagStructure(
		CommandFlagDef{Kind: "string", Name: "mode", Schema: []string{"schema"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errShortSchema, ErrorIDSchemaInvalidValue)

	seen = map[string]struct{}{"mode": {}}
	_, _, errDuplicate := validateFlagStructure(
		CommandFlagDef{Kind: "string", Name: "mode", Schema: []string{"schema", "string"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errDuplicate, ErrorIDSchemaInvalidCombination)

	seen = map[string]struct{}{}
	_, _, errInvalidKind := validateFlagStructure(
		CommandFlagDef{Kind: "map", Name: "mode", Schema: []string{"schema", "string"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errInvalidKind, ErrorIDSchemaInvalidValue)

	seen = map[string]struct{}{}
	ctx, shape, errCustomOk := validateFlagStructure(
		CommandFlagDef{Kind: "string", Name: "postal_code", Schema: []string{"custom", "postal-code", "--country", "US"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	if errCustomOk != nil {
		t.Fatalf("unexpected custom flag structure error: %v", errCustomOk)
	}
	if ctx.name != "postal_code" || shape.Head != "custom" || shape.Operator != "postal-code" {
		t.Fatalf("unexpected custom structure output: ctx=%#v shape=%#v", ctx, shape)
	}

	seen = map[string]struct{}{}
	_, _, errPrimaryShape := validateFlagStructure(
		CommandFlagDef{Kind: "tuple", Name: "range", Schema: []string{"schema", "tuple"}},
		seen,
		DefaultManualFlagNameValidator(),
	)
	assertDocErrID(t, errPrimaryShape, ErrorIDSchemaInvalidValue)
}

func TestValidateNonTupleSchemasInvalidChild(t *testing.T) {
	err := validateNonTupleSchemas([][]string{{"schema", "string"}}, "mode")
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func TestValidateNonTupleSchemasValidAndParseFailure(t *testing.T) {
	if err := validateNonTupleSchemas([][]string{{"schema", "repeatable", "--min-length", "1"}}, "add"); err != nil {
		t.Fatalf("unexpected valid repeatable error: %v", err)
	}

	err := validateNonTupleSchemas([][]string{{"schema"}}, "add")
	assertDocErrID(t, err, ErrorIDSchemaInvalidValue)
}

func TestValidateTupleSchemasBranches(t *testing.T) {
	errBadHead := validateTupleSchemas(2, [][]string{{"custom", "number", "--tuple", "0", "--int"}}, "range")
	assertDocErrID(t, errBadHead, ErrorIDSchemaInvalidValue)

	errMissingIndex := validateTupleSchemas(2, [][]string{{"schema", "number", "--int"}}, "range")
	assertDocErrID(t, errMissingIndex, ErrorIDSchemaTupleMissingIndex)

	errOutOfRange := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple", "0", "--int"},
		{"schema", "number", "--tuple", "3", "--int"},
	}, "range")
	assertDocErrID(t, errOutOfRange, ErrorIDSchemaTupleIndexOutOfRange)

	errDup := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple", "0", "--int"},
		{"schema", "number", "--tuple", "0", "--int"},
	}, "range")
	assertDocErrID(t, errDup, ErrorIDSchemaTupleDuplicateSlot)

	errMissingSlot := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple", "0", "--int"},
	}, "range")
	assertDocErrID(t, errMissingSlot, ErrorIDSchemaTupleMissingSlot)

	errMissingTupleValue := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple"},
	}, "range")
	assertDocErrID(t, errMissingTupleValue, ErrorIDSchemaMissingValue)

	errTupleInvalidValue := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple", "x"},
	}, "range")
	assertDocErrID(t, errTupleInvalidValue, ErrorIDSchemaInvalidValue)

	if err := validateTupleSchemas(2, [][]string{
		{"schema", "number", "--tuple", "0", "--int"},
		{"schema", "number", "--tuple", "1", "--int"},
		{"schema", "repeatable", "--min-length", "1"},
	}, "range"); err != nil {
		t.Fatalf("unexpected tuple schema error: %v", err)
	}
}

func TestParseChildSchemaAndTupleIndexBranches(t *testing.T) {
	_, errShort := parseChildSchema([]string{"schema"})
	assertDocErrID(t, errShort, ErrorIDSchemaInvalidValue)

	shape, err := parseChildSchema([]string{"schema", "number", "--tuple", "1", "--int"})
	if err != nil {
		t.Fatalf("unexpected parseChildSchema error: %v", err)
	}
	if !shape.hasSlot || shape.slot != 1 {
		t.Fatalf("unexpected child shape: %#v", shape)
	}

	_, _, errTupleMissing := parseTupleIndex([]string{"schema", "number", "--tuple"})
	assertDocErrID(t, errTupleMissing, ErrorIDSchemaMissingValue)

	_, _, errTupleInvalid := parseTupleIndex([]string{"schema", "number", "--tuple", "x"})
	assertDocErrID(t, errTupleInvalid, ErrorIDSchemaInvalidValue)

	slot, has, errNone := parseTupleIndex([]string{"schema", "number", "--int"})
	if errNone != nil {
		t.Fatalf("unexpected parseTupleIndex no-tuple error: %v", errNone)
	}
	if has || slot != 0 {
		t.Fatalf("unexpected no-tuple parse result: slot=%d has=%v", slot, has)
	}

	_, errChildMissing := parseChildSchema([]string{"schema", "number", "--tuple"})
	assertDocErrID(t, errChildMissing, ErrorIDSchemaMissingValue)

	_, errChildInvalid := parseChildSchema([]string{"schema", "number", "--tuple", "x"})
	assertDocErrID(t, errChildInvalid, ErrorIDSchemaInvalidValue)
}

func TestDefaultRegexFlagNameValidator(t *testing.T) {
	v := DefaultRegexFlagNameValidator()
	if err := v("mode_1.alpha-beta"); err != nil {
		t.Fatalf("expected valid name, got %v", err)
	}
	if err := v("mode!"); err == nil {
		t.Fatal("expected invalid name")
	}
}
