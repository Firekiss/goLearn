package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	nowTime := time.Now().Unix()
	if nowTime%Female == 0 {
		fmt.Println("Female")
	} else {
		fmt.Println("Man")
	}
}
