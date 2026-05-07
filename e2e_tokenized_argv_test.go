package picker

import "testing"

func TestE2ETokenizedArgvFlagFormParity(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	argvA := []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20"}
	argvB := []string{"wash", "start", "--mode=normal", "--spin=1200", "--range=10,20"}

	gotA, errA := ValidateWithDocumentJSON(raw, argvA)
	if errA != nil {
		t.Fatalf("unexpected error A: %v", errA)
	}
	gotB, errB := ValidateWithDocumentJSON(raw, argvB)
	if errB != nil {
		t.Fatalf("unexpected error B: %v", errB)
	}

	assertStringValue(t, gotA.Values, "mode", "normal")
	assertStringValue(t, gotB.Values, "mode", "normal")
	assertNumberValue(t, gotA.Values, "spin", 1200)
	assertNumberValue(t, gotB.Values, "spin", 1200)
	assertTupleStringValues(t, gotA.Values, "range", "10", "20")
	assertTupleStringValues(t, gotB.Values, "range", "10", "20")
}

func TestE2ETokenizedArgvRepeatableMix(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	argv := []string{"wash", "start", "--add=a,b", "--add", "c", "--mode", "normal", "--spin", "1200", "--range", "10,20"}
	got, err := ValidateWithDocumentJSON(raw, argv)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	v := got.Values["add"]
	if len(v.List) != 3 {
		t.Fatalf("unexpected repeatable add list length: %#v", v)
	}
}

func TestE2ETokenizedArgvTupleFailures(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	cases := []struct {
		name string
		argv []string
	}{
		{name: "too-short", argv: []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10"}},
		{name: "too-long", argv: []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20,30"}},
		{name: "empty", argv: []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", ""}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := ValidateWithDocumentJSON(raw, tc.argv)
			assertErrorDetail(t, err, ErrorIDValidationTuple, ErrorKindValidation)
		})
	}
}

func TestE2ETokenizedArgvBooleanHandling(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)

	gotTrue, errTrue := ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--extra-rinse"})
	if errTrue != nil {
		t.Fatalf("unexpected true boolean error: %v", errTrue)
	}
	assertBoolValue(t, gotTrue.Values, "extra-rinse", true)

	gotFalse, errFalse := ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--extra-rinse=false"})
	if errFalse != nil {
		t.Fatalf("unexpected explicit false boolean error: %v", errFalse)
	}
	assertBoolValue(t, gotFalse.Values, "extra-rinse", false)
}

func TestE2ETokenizedArgvCommandPathVariants(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)

	// With prefix.
	_, err := ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20"})
	if err != nil {
		t.Fatalf("unexpected prefixed argv error: %v", err)
	}

	// Without prefix.
	_, err = ValidateWithDocumentJSON(raw, []string{"--mode", "normal", "--spin", "1200", "--range", "10,20"})
	if err != nil {
		t.Fatalf("unexpected non-prefixed argv error: %v", err)
	}

	// Wrong prefix should be treated as runtime tokens and fail type check.
	_, err = ValidateWithDocumentJSON(raw, []string{"wash", "stop", "--mode", "normal"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
}

func TestE2ETokenizedArgvUnknownFlagAndNumericEdges(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)

	_, err := ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--mod", "x"})
	assertErrorDetail(t, err, ErrorIDValidationUnexpectedFlag, ErrorKindValidation)

	// 0 and -1 parse as numbers in current runtime parser.
	_, err = ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "0", "--range", "10,20"})
	if err != nil {
		t.Fatalf("unexpected spin=0 error: %v", err)
	}
	_, err = ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "-1", "--range", "10,20"})
	if err != nil {
		t.Fatalf("unexpected spin=-1 error: %v", err)
	}
	_, err = ValidateWithDocumentJSON(raw, []string{"wash", "start", "--mode", "normal", "--spin", "abc", "--range", "10,20"})
	assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
}

func TestE2ETokenizedArgvDeterminism(t *testing.T) {
	raw := mustLoadArgsCommandFixture(t)
	argv := []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--add", "a,b", "--add", "c"}

	got1, err1 := ValidateWithDocumentJSON(raw, argv)
	if err1 != nil {
		t.Fatalf("unexpected first run error: %v", err1)
	}
	got2, err2 := ValidateWithDocumentJSON(raw, argv)
	if err2 != nil {
		t.Fatalf("unexpected second run error: %v", err2)
	}

	assertStringValue(t, got1.Values, "mode", "normal")
	assertStringValue(t, got2.Values, "mode", "normal")
	assertNumberValue(t, got1.Values, "spin", 1200)
	assertNumberValue(t, got2.Values, "spin", 1200)
	assertTupleStringValues(t, got1.Values, "range", "10", "20")
	assertTupleStringValues(t, got2.Values, "range", "10", "20")
	if len(got1.Values["add"].List) != len(got2.Values["add"].List) {
		t.Fatalf("non-deterministic repeatable list length: got1=%d got2=%d", len(got1.Values["add"].List), len(got2.Values["add"].List))
	}
}
