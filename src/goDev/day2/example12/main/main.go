// 对于一个数 n 求 n 的阶乘之和
package main

import "fmt"

// 获取一个数的阶乘
func calcFactorial(num int64) int64 {
	var ret int64 = 1
	var i int64 = 1
	for ; i <= num; i++ {
		ret = ret * i
	}
	return ret
}

// 计算阶乘累加之和
func calcFactorialAsse(n int64) int64 {
	var ret int64
	var i int64 = 1
	for ; i <= n; i++ {
		ret += calcFactorial(i)
	}
	return ret
}

func main() {
	fmt.Printf("10 的阶乘累加结果是 %d\n", calcFactorialAsse(10))
	fmt.Printf("100 的阶乘累加结果是 %d\n", calcFactorialAsse(100))
	fmt.Printf("1000 的阶乘累加结果是 %d\n", calcFactorialAsse(1000))
}
