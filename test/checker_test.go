package test

import (
	"testing"

	"github.com/fathimasithara01/realimage-challenge-2016/internal/permissions"
)

func TestChecker_BasicEvaluation(t *testing.T) {
	svc := permissions.NewService()

	svc.AddRule("D1", permissions.Include, "UNITEDSTATES")
	svc.AddRule("D1", permissions.Include, "INDIA")
	svc.AddRule("D1", permissions.Exclude, "CHENNAI-TAMILNADU-INDIA")

	ch := permissions.NewChecker(svc)

	if !ch.Evaluate(permissions.Line{"D1", "CHICAGO-ILLINOIS-UNITEDSTATES"}) {
		t.Fatalf("expected CHICAGO allowed")
	}

	if ch.Evaluate(permissions.Line{"D1", "CHENNAI-TAMILNADU-INDIA"}) {
		t.Fatalf("expected CHENNAI denied")
	}
}

func TestChecker_MissingDistributor(t *testing.T) {
	svc := permissions.NewService()
	ch := permissions.NewChecker(svc)

	if ch.Evaluate(permissions.Line{"XYZ", "INDIA"}) {
		t.Fatalf("missing distributor must be denied")
	}
}
