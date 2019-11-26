package main

import (
	"fmt"
)

func chgVal(pointer *int) {
	*pointer = 12
}

func main() {
	var intVal int = 10
	var strVal string = "alex"

	fmt.Printf("intVal pointer address is: %d\n", &intVal)
	fmt.Printf("strVal pointer address is: %d\n", &strVal)

	chgVal(&intVal)
	fmt.Printf("intVal curVal is: %d\n", intVal)
}
