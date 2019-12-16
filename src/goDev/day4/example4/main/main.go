package main

import "fmt"



func main() {
	var myMap map[string]map[string]string
	myMap = make(map[string]map[string]string)

	_, ok := myMap["province"]
	if !ok {
		myMap["province"] = make(map[string]string)
		myMap["province"]["city"] = "南京"
	}

	fmt.Println(myMap)
}