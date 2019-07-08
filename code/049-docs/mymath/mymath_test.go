package mymath

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	// Sum happy case
	sum := Sum(1, 2, 3)
	if sum != 6 {
		t.Errorf("Expected value 6, got %v", sum)
	}

	// Sum empty case
	sum = Sum()
	if sum != 0 {
		t.Errorf("Expected value 0, got %v", sum)
	}

	type testData struct {
		data   []int
		answer int
	}

	// With table structure
	tests := []testData{
		testData{[]int{}, 0},
		testData{[]int{1, 1}, 2},
		testData{[]int{-1, 1}, 0},
		testData{[]int{10, 5}, 15},
		testData{[]int{21, 21}, 42},
	}

	for _, test := range tests {
		sum = Sum(test.data...)
		if sum != test.answer {
			t.Errorf("Expected value %v, got %v", test.answer, sum)
		}
	}
}

func ExampleSum() {

	// Pass in values to sum
	sum := Sum(1, 2, 3)

	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
