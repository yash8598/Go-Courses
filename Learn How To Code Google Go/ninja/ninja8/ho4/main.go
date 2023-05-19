package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 65, 999, 21, 12, 3, 1, 56, 786, 42}
	xs := []string{
		"random", "rainbow", "in", "nimo", "kashin", "parthin", "loki", "logi", "joshi", "bonde",
	}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
