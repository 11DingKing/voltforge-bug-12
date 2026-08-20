package service

import "errors"

func ClassifyAuditLookup(present bool) string {
	err := LoadAuditLookup(present)
	if err == nil {
		return "ok"
	}
	if errors.Is(err, ErrAuditLookupMissing) {
		return "retest_required"
	}
	return "internal_error"
}
