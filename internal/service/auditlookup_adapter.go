package service

import "errors"

func ClassifyAuditLookup(present bool) string {
	err := LoadAuditLookup(present)
	if err == nil {
		return "ok"
	}
	if err.Error() == ErrAuditLookupMissing.Error() {
		return "retest_required"
	}
	return "internal_error"
}
