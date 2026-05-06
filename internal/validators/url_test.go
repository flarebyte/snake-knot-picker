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
	if err := ValidateURL("https:opaque-data", opts); err == nil {
		t.Fatal("expected opaque URL rejection")
	}
	if err := ValidateURL("http://%zz", opts); err == nil {
		t.Fatal("expected parse error rejection")
	}
	if err := ValidateURL("//example.com/path", opts); err == nil {
		t.Fatal("expected empty-scheme rejection")
	}
	if err := ValidateURL("https:///path", opts); err == nil {
		t.Fatal("expected empty-host rejection")
	}
	if err := ValidateURL("ftp://example.com/path", opts); err == nil {
		t.Fatal("expected scheme whitelist rejection")
	}
	if err := ValidateURL("https://example.com/path?q=1", URLOptions{}); err == nil {
		t.Fatal("expected query rejection when AllowQuery is false")
	}
}

func TestValidateURLAllowBranches(t *testing.T) {
	opts := URLOptions{
		AllowQuery: true,
		AllowUser:  true,
		AllowPort:  true,
		AllowFrag:  true,
		AllowIPs:   true,
	}
	if err := ValidateURL("https://user:pw@127.0.0.1:8443/path?q=1#frag", opts); err != nil {
		t.Fatalf("unexpected allow-branches error: %v", err)
	}
	if err := ValidateURL("https://example.com/path", URLOptions{AllowDomains: []string{"example.com"}}); err != nil {
		t.Fatalf("unexpected exact-domain error: %v", err)
	}
	if err := ValidateURL("https://api.example.com/path", URLOptions{AllowDomains: []string{" other.com ", " EXAMPLE.COM "}}); err != nil {
		t.Fatalf("unexpected normalized multi-domain allow-list error: %v", err)
	}
	if err := ValidateURL("https://127.0.0.1/path", URLOptions{AllowDomains: []string{"example.com"}, AllowIPs: true}); err != nil {
		t.Fatalf("unexpected IP allow with domain-list present error: %v", err)
	}
	if err := ValidateURL("http://example.com/path", URLOptions{Scheme: "https"}); err == nil {
		t.Fatal("expected explicit scheme-option mismatch rejection")
	}
}
