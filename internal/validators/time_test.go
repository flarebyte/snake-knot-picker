// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"testing"
	"time"
)

func TestValidateDateTimeAndDuration(t *testing.T) {
	if err := ValidateDate("2026-05-05", ""); err != nil {
		t.Fatalf("unexpected default-layout date error: %v", err)
	}
	if err := ValidateDate("2026-05-05", "ISO8601"); err != nil {
		t.Fatalf("unexpected date error: %v", err)
	}
	if err := ValidateDate("05/05/2026", "ISO8601"); err == nil {
		t.Fatal("expected invalid date rejection")
	}
	if err := ValidateDateTime("2026-05-05T12:30:00Z", "RFC3339"); err != nil {
		t.Fatalf("unexpected datetime error: %v", err)
	}
	if err := ValidateDateTime("2026-05-05T12:30:00Z", ""); err != nil {
		t.Fatalf("unexpected default datetime error: %v", err)
	}
	if err := ValidateDateTime("Tue, 05 May 2026 12:30:00 +0000", "RFC1123Z"); err != nil {
		t.Fatalf("unexpected RFC1123Z datetime error: %v", err)
	}
	if err := ValidateDateTime("bad-rfc1123z", "RFC1123Z"); err == nil {
		t.Fatal("expected invalid RFC1123Z datetime rejection")
	}
	if err := ValidateDateTime("invalid", "RFC3339"); err == nil {
		t.Fatal("expected invalid datetime rejection")
	}
	if err := ValidateDateTime("1746448200", "Unix"); err != nil {
		t.Fatalf("unexpected Unix datetime error: %v", err)
	}
	if err := ValidateDateTime("x", "Unix"); err == nil {
		t.Fatal("expected invalid Unix datetime rejection")
	}
	if err := ValidateDateTime("2026-05-05T12:30:00Z", "CUSTOM"); err == nil {
		t.Fatal("expected unsupported datetime layout rejection")
	}
	if err := ValidateClockTime("235959", "HHMMSS"); err != nil {
		t.Fatalf("unexpected HHMMSS error: %v", err)
	}
	if err := ValidateClockTime("2359", "HHMM"); err != nil {
		t.Fatalf("unexpected HHMM error: %v", err)
	}
	if err := ValidateClockTime("236060", "HHMMSS"); err == nil {
		t.Fatal("expected invalid clock time rejection")
	}
	if err := ValidateClockTime("2500", "HHMM"); err == nil {
		t.Fatal("expected invalid HHMM clock time rejection")
	}
	if err := ValidateClockTime("235959", "ISO"); err == nil {
		t.Fatal("expected unsupported clock layout rejection")
	}
	min := 5 * time.Minute
	max := 2 * time.Hour
	if err := ValidateDuration("30m", &min, &max); err != nil {
		t.Fatalf("unexpected duration error: %v", err)
	}
	if err := ValidateDuration("2m", &min, &max); err == nil {
		t.Fatal("expected min duration failure")
	}
	if err := ValidateDuration("3h", &min, &max); err == nil {
		t.Fatal("expected max duration failure")
	}
	if err := ValidateDuration("not-duration", &min, &max); err == nil {
		t.Fatal("expected malformed duration rejection")
	}
	if err := ValidateDuration("30m", nil, nil); err != nil {
		t.Fatalf("unexpected nil-bounds duration error: %v", err)
	}
}
