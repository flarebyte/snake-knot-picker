// purpose: Implement URL validation with option-based restrictions such as scheme and host allowlists.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import (
	"net"
	"net/url"
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

type URLOptions struct {
	Scheme       string
	Secure       bool
	AllowQuery   bool
	AllowUser    bool
	AllowPort    bool
	AllowFrag    bool
	AllowDomains []string
	AllowIPs     bool
}

func ValidateURL(value string, opt URLOptions) error {
	u, err := url.Parse(value)
	if err != nil || u.Scheme == "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "url"})
	}
	if u.Opaque != "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "opaque"})
	}
	if u.Host == "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "url"})
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "url"})
	}
	if opt.Secure && u.Scheme != "https" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "secure"})
	}
	if opt.Scheme != "" && u.Scheme != opt.Scheme {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "scheme"})
	}
	if !opt.AllowQuery && u.RawQuery != "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "query"})
	}
	if !opt.AllowFrag && u.Fragment != "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "fragment"})
	}
	if !opt.AllowUser && u.User != nil {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "userinfo"})
	}
	host := u.Hostname()
	if !opt.AllowPort && u.Port() != "" {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "port"})
	}
	ip := net.ParseIP(host)
	if ip != nil && !opt.AllowIPs {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"reason": "ip"})
	}
	if len(opt.AllowDomains) > 0 && ip == nil {
		ok := false
		for _, allowed := range opt.AllowDomains {
			allowed = strings.ToLower(strings.TrimSpace(allowed))
			h := strings.ToLower(host)
			if h == allowed || strings.HasSuffix(h, "."+allowed) {
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
