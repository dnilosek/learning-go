package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	// Happy path
	cenAvg := CenteredAvg([]int{0, 1, 2, 3, 100})
	if cenAvg != 2 {
		t.Error("Expected 2 got", cenAvg)
	}

	// Zero path
	cenAvg = CenteredAvg([]int{})
	if cenAvg != 0 {
		t.Error("Expected 0 got", cenAvg)
	}
}

func ExampleCenteredAve() {
	cenAvg := CenteredAvg([]int{0, 1, 2, 3, 100})
	fmt.Println(cenAvg)
	//Output: 2
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{0, 1, 2, 3, 100})
	}
}
