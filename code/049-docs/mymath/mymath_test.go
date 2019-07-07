package mymath

import (
	"github.eagleview.com/david.nilosek/learning-go/code/049-docs/mymath"
	"testing"
)

func TestingSum(t *testing.T) {
	// Sum happy case
	sum := mymath.Sum(1, 2, 3)
	if sum != 6 {
		t.Errorf("Expected value 6, got %v", sum)
	}

	// Sum empty case
	sum = mymath.Sum()
	if sum != 6 {
		t.Errorf("Expected value 6, got %v", sum)
	}
}
