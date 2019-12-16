package main

import (
	"fmt"
	"sort"
)

func main() {
	list := [5]int{3, 6, 7, 8, 1}
	// sort 模块用来排序
	// 对切片进行排序可以改变原来的数组
	sort.Ints(list[:])
	fmt.Println(list)
}
