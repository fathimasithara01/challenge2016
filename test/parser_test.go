package test

import (
	"testing"

	"github.com/fathimasithara01/realimage-challenge-2016/internal/geo"
)

func TestParseRegion_Full(t *testing.T) {
	r := geo.ParseRegion("BANGALORE-KARNATAKA-INDIA")

	if r.City != "BANGALORE" {
		t.Fatalf("expected City=BANGALORE, got %q", r.City)
	}
	if r.State != "KARNATAKA" {
		t.Fatalf("expected State=KARNATAKA, got %q", r.State)
	}
	if r.Country != "INDIA" {
		t.Fatalf("expected Country=INDIA, got %q", r.Country)
	}
}

func TestParseRegion_StateCountry(t *testing.T) {
	r := geo.ParseRegion("KARNATAKA-INDIA")

	if r.City != "" {
		t.Fatalf("expected empty City, got %q", r.City)
	}
	if r.State != "KARNATAKA" {
		t.Fatalf("expected State=KARNATAKA, got %q", r.State)
	}
	if r.Country != "INDIA" {
		t.Fatalf("expected Country=INDIA, got %q", r.Country)
	}
}

func TestParseRegion_Country(t *testing.T) {
	r := geo.ParseRegion("INDIA")

	if r.City != "" || r.State != "" || r.Country != "INDIA" {
		t.Fatalf("expected only Country=INDIA, got %+v", r)
	}
}
