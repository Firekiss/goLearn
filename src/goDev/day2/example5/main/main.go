package main

import (
	"fmt"
)

func swap(a string, b string) (string, string) {
	c := a
	a = b
	b = c
	return a, b
}

func main() {
	a := "a"
	b := "b"
	a, b = swap(a, b)
	fmt.Printf("a = %s\n", a)
	fmt.Printf("b = %s\n", b)
}
