package validators

import (
	"testing"
	"time"
)

func TestValidateDateTimeAndDuration(t *testing.T) {
	if err := ValidateDate("2026-05-05", "ISO8601"); err != nil {
		t.Fatalf("unexpected date error: %v", err)
	}
	if err := ValidateDate("05/05/2026", "ISO8601"); err == nil {
		t.Fatal("expected invalid date rejection")
	}
	if err := ValidateDateTime("2026-05-05T12:30:00Z", "RFC3339"); err != nil {
		t.Fatalf("unexpected datetime error: %v", err)
	}
	if err := ValidateDateTime("Tue, 05 May 2026 12:30:00 +0000", "RFC1123Z"); err != nil {
		t.Fatalf("unexpected RFC1123Z datetime error: %v", err)
	}
	if err := ValidateDateTime("invalid", "RFC3339"); err == nil {
		t.Fatal("expected invalid datetime rejection")
	}
	if err := ValidateClockTime("235959", "HHMMSS"); err != nil {
		t.Fatalf("unexpected HHMMSS error: %v", err)
	}
	if err := ValidateClockTime("236060", "HHMMSS"); err == nil {
		t.Fatal("expected invalid clock time rejection")
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
}
