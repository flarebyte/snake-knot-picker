// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import "testing"

func TestValidateColor(t *testing.T) {
	if err := ValidateColor("#A1B2C3", "hex", false); err != nil {
		t.Fatalf("unexpected hex color error: %v", err)
	}
	if err := ValidateColor("#A1B2C3DD", "hex", true); err != nil {
		t.Fatalf("unexpected alpha hex color error: %v", err)
	}
	if err := ValidateColor("#A1B2C3DD", "hex", false); err == nil {
		t.Fatal("expected alpha disallowed failure")
	}
	if err := ValidateColor("rgb(1,2,3)", "hex", true); err == nil {
		t.Fatal("expected css color rejection")
	}
	if err := ValidateColor("#A1B2C3", "rgb", true); err == nil {
		t.Fatal("expected unsupported format rejection")
	}
}
