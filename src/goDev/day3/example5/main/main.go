package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inStr string
	fmt.Println("请输入需要转换的值: ")
	fmt.Scanf("%s", &inStr)
	i, err := strconv.Atoi(inStr)
	if err != nil {
		fmt.Printf("can not convert to int %s ", err)
		return
	}

	fmt.Printf("转换之后值为 %d\n", i)
}
