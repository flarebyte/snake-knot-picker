package validators

import (
	"strings"

	"github.com/flarebyte/snake-knot-picker"
)

type ARNOptions struct {
	AllowPartition []string
	AllowService   []string
	AllowRegion    []string
	AllowAccountID []string
	AllowResource  []string
}

type arnParts struct {
	Partition string
	Service   string
	Region    string
	AccountID string
	Resource  string
}

func ValidateARN(value string, opt ARNOptions) error {
	p, ok := parseARN(value)
	if !ok {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"format": "arn"})
	}
	if !inAllowList(opt.AllowPartition, p.Partition) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"part": "partition"})
	}
	if !inAllowList(opt.AllowService, p.Service) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"part": "service"})
	}
	if !inAllowList(opt.AllowRegion, p.Region) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"part": "region"})
	}
	if !inAllowList(opt.AllowAccountID, p.AccountID) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"part": "account"})
	}
	if !inAllowList(opt.AllowResource, p.Resource) {
		return picker.NewValidationError(picker.ErrorIDValidationFormat, map[string]string{"part": "resource"})
	}
	return nil
}

func parseARN(v string) (arnParts, bool) {
	parts := strings.SplitN(v, ":", 6)
	if len(parts) != 6 || parts[0] != "arn" {
		return arnParts{}, false
	}
	if parts[2] == "" || parts[5] == "" {
		return arnParts{}, false
	}
	return arnParts{
		Partition: parts[1],
		Service:   parts[2],
		Region:    parts[3],
		AccountID: parts[4],
		Resource:  parts[5],
	}, true
}

func inAllowList(allowed []string, value string) bool {
	if len(allowed) == 0 {
		return true
	}
	for _, a := range allowed {
		if a == value {
			return true
		}
	}
	return false
}
