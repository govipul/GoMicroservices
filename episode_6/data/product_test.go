package data

import "testing"

func TestCheckValidattionFail(t *testing.T) {
	p := &Product{}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
func TestCheckValidattionSuccess(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 1.0,
		SKU:   "abc-def-ghik",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
