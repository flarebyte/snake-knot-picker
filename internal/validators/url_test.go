package validators

import "testing"

func TestValidateURL(t *testing.T) {
	opts := URLOptions{
		Scheme:       "https",
		Secure:       true,
		AllowQuery:   true,
		AllowDomains: []string{"example.com"},
	}
	if err := ValidateURL("https://api.example.com/path?q=1", opts); err != nil {
		t.Fatalf("unexpected valid url error: %v", err)
	}
	if err := ValidateURL("http://api.example.com/path", opts); err == nil {
		t.Fatal("expected secure/scheme failure")
	}
	if err := ValidateURL("https://example.com.evil.org/path", opts); err == nil {
		t.Fatal("expected domain boundary failure")
	}
	if err := ValidateURL("https://user:pw@example.com/path", opts); err == nil {
		t.Fatal("expected user-info failure")
	}
	if err := ValidateURL("https://example.com:8443/path", opts); err == nil {
		t.Fatal("expected port failure")
	}
	if err := ValidateURL("https://127.0.0.1/path", opts); err == nil {
		t.Fatal("expected ip failure")
	}
	if err := ValidateURL("https://example.com/path#frag", opts); err == nil {
		t.Fatal("expected fragment failure")
	}
	if err := ValidateURL("mailto:x@example.com", opts); err == nil {
		t.Fatal("expected non-http scheme failure")
	}
}
