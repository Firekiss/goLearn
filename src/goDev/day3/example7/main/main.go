package main

import "fmt"

func drawTree(high int) {
	for i := 1; i <= high; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("A")

			if j == i {
				fmt.Println("")
			}
		}
	}
}

func main() {
	drawTree(15)
}
