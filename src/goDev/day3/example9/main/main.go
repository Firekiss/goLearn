package main

import "fmt"

func add(a int, args ...int) int {
	if len(args) > 0 {
		for _, v := range args {
			a += v
		}
	}
	return a
}

func concat(str string, args ...string) string {
	if len(args) > 0 {
		for _, v := range args {
			str += v
		}
	}
	return str
}

func main() {
	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(concat("my", " name", " alex"))
}
