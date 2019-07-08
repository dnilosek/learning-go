package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	s := "I am a dog"
	countS := Count(s)
	if countS != 4 {
		t.Error("Expected 4, got", countS)
	}

	s = ""
	countS = Count(s)
	if countS != 0 {
		t.Error("Expected 0, got", countS)
	}
}

func ExampleCount() {
	s := "I am a dog"
	countS := Count(s)

	fmt.Println(countS)
	//Output: 4
}

func BenchmarkCount(b *testing.B) {
	s := "I am a dog"
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}
