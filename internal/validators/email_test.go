package validators

import "testing"

func TestValidateEmail(t *testing.T) {
	if err := ValidateEmail("user@example.com", []string{"example.com"}); err != nil {
		t.Fatalf("unexpected email error: %v", err)
	}
	if err := ValidateEmail("user@other.com", []string{"example.com"}); err == nil {
		t.Fatal("expected allow-domain failure")
	}
	if err := ValidateEmail("John <user@example.com>", []string{"example.com"}); err == nil {
		t.Fatal("expected display-name rejection")
	}
	if err := ValidateEmail("not-an-email", nil); err == nil {
		t.Fatal("expected malformed email rejection")
	}
	if err := ValidateEmail("@example.com", nil); err == nil {
		t.Fatal("expected missing local-part rejection")
	}
}
