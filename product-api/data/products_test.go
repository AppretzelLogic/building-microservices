package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nico",
		Price: 1.00,
		SKU:   "abc-abc-abc",
	}

	err := p.date()

	if err != nil {
		t.Fatal(err)
	}
}