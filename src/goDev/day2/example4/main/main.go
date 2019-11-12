package main

import (
	"fmt"
	"os"
)

func main() {
	pathEnv := os.Getenv("PATH")
	fmt.Printf("this mac's PATH env is: %s\n", pathEnv)
}
