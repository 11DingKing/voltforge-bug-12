package service

import (
	"errors"
	"testing"
)

func TestVoltForge12(t *testing.T) {
	if got := ClassifyAuditLookup(false); got != "retest_required" {
		t.Fatalf("got %s", got)
	}
	if !errors.Is(LoadAuditLookup(false), ErrAuditLookupMissing) {
		t.Fatal("missing identity was not preserved")
	}
}
