package main

import (
	"fmt"
	"time"
)

// 使用多线程打印数字
func printNums(i int) {
	for a := 0; a < i; a++ {
		go fmt.Println(a)
	}
}

func main() {
	printNums(100)
	time.Sleep(5 * time.Second)
}
