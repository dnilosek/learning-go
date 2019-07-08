package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	// Simple test
	years := Years(1)
	if years != 7 {
		t.Error("Expected 7 got", years)
	}

	// Simple test
	years = Years(0)
	if years != 0 {
		t.Error("Expected 0 got", years)
	}
}

func ExampleYears() {
	years := Years(1)
	fmt.Println("1 human year is", years, "dog years")
	//Output: 1 human year is 7 dog years
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func TestYearsTwo(t *testing.T) {
	// Simple test
	years := YearsTwo(1)
	if years != 7 {
		t.Error("Expected 7 got", years)
	}

	// Simple test
	years = YearsTwo(0)
	if years != 0 {
		t.Error("Expected 0 got", years)
	}
}

func ExampleYearsTwo() {
	years := YearsTwo(1)
	fmt.Println("1 human year is", years, "dog years")
	//Output: 1 human year is 7 dog years
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
