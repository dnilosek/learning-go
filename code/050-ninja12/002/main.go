package main

import (
	"fmt"
	"github.com/dnilosek/learning-go/code/050-ninja12/002/quote"
	"github.com/dnilosek/learning-go/code/050-ninja12/002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
