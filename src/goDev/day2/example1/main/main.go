package main

import "fmt"

func printAddSons(n int) {
	for f := 0; f <= n; f++ {
		fmt.Printf("%d + %d = %d\n", f, n-f, n)
	}
}

func main() {
	printAddSons(5)
	printAddSons(10)
}
