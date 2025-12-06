package test

import (
	"testing"

	"github.com/fathimasithara01/realimage-challenge-2016/internal/permissions"
)

func TestPermissions_IncludeExclude(t *testing.T) {
	svc := permissions.NewService()

	svc.AddRule("D1", permissions.Include, "INDIA")
	svc.AddRule("D1", permissions.Exclude, "KARNATAKA-INDIA")

	if !svc.Check(permissions.Line{"D1", "INDIA"}) {
		t.Fatalf("expected INDIA allowed")
	}
	if svc.Check(permissions.Line{"D1", "KARNATAKA-INDIA"}) {
		t.Fatalf("expected KARNATAKA-INDIA denied")
	}
}

func TestPermissions_LastRuleWins(t *testing.T) {
	svc := permissions.NewService()

	svc.AddRule("D2", permissions.Include, "INDIA")
	svc.AddRule("D2", permissions.Exclude, "KARNATAKA-INDIA")
	svc.AddRule("D2", permissions.Include, "BANGALORE-KARNATAKA-INDIA")

	if !svc.Check(permissions.Line{"D2", "BANGALORE-KARNATAKA-INDIA"}) {
		t.Fatalf("expected city override → allowed")
	}
}

func TestPermissions_UnknownDistributor(t *testing.T) {
	svc := permissions.NewService()

	if svc.Check(permissions.Line{"X", "INDIA"}) {
		t.Fatalf("unknown distributor should be denied")
	}
}
