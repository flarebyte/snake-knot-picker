// purpose: Implement email validation with optional domain allowlist filtering.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"net/mail"
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

func ValidateEmail(value string, allowDomains []string) error {
	addr, err := mail.ParseAddress(value)
	if err != nil {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "email"})
	}
	// Reject display-name style input.
	if strings.TrimSpace(addr.Address) != strings.TrimSpace(value) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "display_name"})
	}
	at := strings.LastIndex(addr.Address, "@")
	domain := strings.ToLower(addr.Address[at+1:])
	if len(allowDomains) > 0 {
		ok := false
		for _, d := range allowDomains {
			d = strings.ToLower(strings.TrimSpace(d))
			if domain == d {
				ok = true
				break
			}
		}
		if !ok {
			return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "domain"})
		}
	}
	return nil
}
