package service

import (
	"errors"
	"fmt"
)

var ErrAuditLookupMissing = errors.New("auditlookup evidence missing")

func LoadAuditLookup(present bool) error {
	if present {
		return nil
	}
	return fmt.Errorf("load auditlookup: %v", ErrAuditLookupMissing)
}
