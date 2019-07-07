package davemath

import "testing"

func TestAverage(t *testing.T) {
	// Expected path
	v := Average(1.0, 2.0)
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}

	// Empty path
	v = Average()
	if v != 0.0 {
		t.Error("Expected 0, got ", v)
	}
}
