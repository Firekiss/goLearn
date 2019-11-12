package main

import (
	"fmt"
)

func swap(a *string, b *string) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := "a"
	b := "b"

	swap(&a, &b)
	fmt.Printf("a=%s\n", a)
	fmt.Printf("b=%s\n", b)
}
