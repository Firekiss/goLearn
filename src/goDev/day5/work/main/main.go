package main

import(
	"fmt"
)

// 冒泡排序
func bubSort(arr []int){
	
	// 代表当前需要确定第 i + 1 大的值索引
	// j 代表当前两两排序的索引值
	l := len(arr)
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

// 选择排序
func chooseSort(arr []int) {

	l := len(arr)
	for i := 0; i < l - 1; i++ {
		min := arr[i]
		for j := i; j < l; j++ {
			if min > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				min = arr[i]
			}
		}
	}
}



func main() {
	arr := []int{5, 15, 4, 75, 86, 3, 2, 49}
	fmt.Println(arr)
	bubSort(arr)
	fmt.Println(arr)

	arrChoose := []int{7, 45, 63, 15, 865, 43, 81, 2, 21, 19}
	fmt.Println(arrChoose)
	chooseSort(arrChoose)
	fmt.Println(arrChoose)
}