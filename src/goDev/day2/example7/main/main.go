package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
