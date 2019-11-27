// 找出 1000 以内的完数 6 = 1 + 2 + 3
package main

import (
	"fmt"
	"math"
)

func getFactors(num int) []int {
	// 获取开平方乡下取整数值
	sqrt := int(math.Floor(math.Sqrt(float64(num))))
	factors := []int{}

	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			factors = append(factors, i)

			if num/i < num {
				factors = append(factors, num/i)
			}
		}
	}

	return factors
}

func isPerfectNum(num int) bool {
	sum := 0
	factors := getFactors(num)
	for _, i := range factors {
		sum = sum + i
	}
	if sum == num {
		return true
	}
	return false
}

func main() {
	perNums := []int{}
	for i := 1; i <= 1000; i++ {
		if isPerfectNum(i) {
			perNums = append(perNums, i)
		}
	}

	fmt.Printf("1000以内所有的完数为：%v\n", perNums)
}
