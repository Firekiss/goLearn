package main

import "fmt"

func main() {
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["name"] = "alex"
	myMap["age"] = "23"

	_, ok := myMap["sex"]
	if ok {
		fmt.Println("性别字段是存在的")
	} else {
		fmt.Println("性别字段不存在")
	}

	fmt.Println(myMap)
}
