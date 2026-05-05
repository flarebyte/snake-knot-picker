package validators

import "testing"

func TestParseEnumCandidates(t *testing.T) {
	values, err := ParseEnumCandidates("cold,warm,hot", ",")
	if err != nil || len(values) != 3 {
		t.Fatalf("unexpected enum parse: values=%v err=%v", values, err)
	}
	if values[1] != "warm" {
		t.Fatalf("unexpected enum value: %v", values)
	}

	_, err = ParseEnumCandidates("red; green;blue", ";")
	if err == nil {
		t.Fatal("expected whitespace enum definition error")
	}
	_, err = ParseEnumCandidates("cold,,hot", ",")
	if err == nil {
		t.Fatal("expected empty enum definition error")
	}
}

func TestValidateStringEnumNoRuntimeTrim(t *testing.T) {
	err := ValidateString(" warm", StringOptions{Enum: []string{"cold", "warm", "hot"}})
	if err == nil {
		t.Fatal("expected enum mismatch with leading whitespace")
	}
	if err := ValidateString("warm", StringOptions{Enum: []string{"cold", "warm", "hot"}}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateStringClassesAndScripts(t *testing.T) {
	cases := []struct {
		name    string
		value   string
		options StringOptions
		wantErr bool
	}{
		{name: "alphabetic-ok", value: "AbCd", options: StringOptions{Alphabetic: true}},
		{name: "alphabetic-fail", value: "abc1", options: StringOptions{Alphabetic: true}, wantErr: true},
		{name: "hexa-ok", value: "0A1f", options: StringOptions{Hexa: true}},
		{name: "blank-ok", value: " \t\t", options: StringOptions{Blank: true}},
		{name: "unicode-letter-ok", value: "東京", options: StringOptions{UnicodeLetter: true}},
		{name: "unicode-number-ok", value: "١٢٣", options: StringOptions{UnicodeNumber: true}},
		{name: "latin-ok", value: "resume", options: StringOptions{Latin: true}},
		{name: "han-ok", value: "漢字", options: StringOptions{Han: true}},
		{name: "devanagari-ok", value: "नमस्ते", options: StringOptions{Devanagari: true}},
		{name: "arabic-ok", value: "مرحبا", options: StringOptions{Arabic: true}},
		{name: "hiragana-ok", value: "あいう", options: StringOptions{Hiragana: true}},
		{name: "katakana-ok", value: "アイウ", options: StringOptions{Katakana: true}},
		{name: "hangul-ok", value: "한글", options: StringOptions{Hangul: true}},
		{name: "tamil-ok", value: "தமிழ்", options: StringOptions{Tamil: true}},
		{name: "gujarati-ok", value: "ગુજરાતી", options: StringOptions{Gujarati: true}},
		{name: "ethiopic-ok", value: "ሰላም", options: StringOptions{Ethiopic: true}},
	}
	for _, tc := range cases {
		err := ValidateString(tc.value, tc.options)
		if tc.wantErr && err == nil {
			t.Fatalf("%s expected error", tc.name)
		}
		if !tc.wantErr && err != nil {
			t.Fatalf("%s unexpected error: %v", tc.name, err)
		}
	}
}

func TestValidateStringBase64PrefixAndBoolean(t *testing.T) {
	if err := ValidateString("aGVsbG8=", StringOptions{Base64: true}); err != nil {
		t.Fatalf("unexpected base64 error: %v", err)
	}
	if err := ValidateString("not-base64", StringOptions{Base64: true}); err == nil {
		t.Fatal("expected invalid base64 error")
	}
	if err := ValidateString("wash-start", StringOptions{StartsWith: "wash-"}); err != nil {
		t.Fatalf("unexpected prefix error: %v", err)
	}
	if err := ValidateString("start-wash", StringOptions{StartsWith: "wash-"}); err == nil {
		t.Fatal("expected prefix failure")
	}
	if err := ValidateString("true", StringOptions{BooleanString: true}); err != nil {
		t.Fatalf("unexpected boolean string error: %v", err)
	}
	if err := ValidateString("True", StringOptions{BooleanString: true}); err == nil {
		t.Fatal("expected strict boolean string failure")
	}
}
