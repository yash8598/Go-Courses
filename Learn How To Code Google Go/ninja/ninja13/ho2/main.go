package main

import (
	"fmt"

	"github.com/yash15122001/gotesting2/quote"
	"github.com/yash15122001/gotesting2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
