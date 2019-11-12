package main

import (
	"fmt"
	"goDev/day1/pkgExample/calc"
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)

	fmt.Println("sum= ", sum)
	fmt.Println("sub= ", sub)
}
