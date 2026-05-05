package validators

import (
	"math"
	"testing"
)

func TestParseNumberString(t *testing.T) {
	if n, err := ParseNumberString("42"); err != nil || n != 42 {
		t.Fatalf("unexpected parse int: n=%v err=%v", n, err)
	}
	if n, err := ParseNumberString("3.14"); err != nil || n != 3.14 {
		t.Fatalf("unexpected parse float: n=%v err=%v", n, err)
	}
	if _, err := ParseNumberString("abc"); err == nil {
		t.Fatal("expected invalid numeric string error")
	}
	if _, err := ParseNumberString("1e9999"); err == nil {
		t.Fatal("expected overflow-like parse rejection")
	}
}

func TestValidateNumber(t *testing.T) {
	min := 1.0
	max := 10.0
	multiple := 0.5
	if err := ValidateNumber(4, NumberOptions{Int: true, Min: &min, Max: &max}); err != nil {
		t.Fatalf("unexpected int/range validation error: %v", err)
	}
	if err := ValidateNumber(4.2, NumberOptions{Int: true}); err == nil {
		t.Fatal("expected int-only failure")
	}
	if err := ValidateNumber(0.5, NumberOptions{Min: &min}); err == nil {
		t.Fatal("expected min range failure")
	}
	if err := ValidateNumber(12, NumberOptions{Max: &max}); err == nil {
		t.Fatal("expected max range failure")
	}
	if err := ValidateNumber(1.5, NumberOptions{MultipleOf: &multiple}); err != nil {
		t.Fatalf("unexpected multiple-of success failure: %v", err)
	}
	if err := ValidateNumber(1.3, NumberOptions{MultipleOf: &multiple}); err == nil {
		t.Fatal("expected multiple-of failure")
	}
	if err := ValidateNumber(math.NaN(), NumberOptions{}); err == nil {
		t.Fatal("expected NaN failure")
	}
	if err := ValidateNumber(math.Inf(1), NumberOptions{}); err == nil {
		t.Fatal("expected +Inf failure")
	}
}

