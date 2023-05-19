package word

import (
	"fmt"
	"testing"

	"github.com/yash15122001/gotesting2/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		case "three":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output :
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
