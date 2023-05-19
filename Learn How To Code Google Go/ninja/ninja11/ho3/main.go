package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffe",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
	fmt.Println("foo ran - ", e.(customErr).info)
}
