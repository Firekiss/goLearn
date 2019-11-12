// 找出 101 - 200 之前所有的素数 输出个数和所有的值
package main

import "fmt"

// 判断输入的整数是否是素数
func isPrimeNum(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var primeNums []int
	min := 101
	max := 200

	for i := min; i <= max; i++ {
		if isPrimeNum(i) {
			primeNums = append(primeNums, i)
		}
	}

	fmt.Printf("%d - %d 之前的素数一共有 %d 个\n", min, max, len(primeNums))
	fmt.Printf("他们是 %v \n", primeNums)
}
