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
	if at <= 0 || at+1 >= len(addr.Address) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "email"})
	}
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
