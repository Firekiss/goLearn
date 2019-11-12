// 打印 100 - 999 之中所有的水仙花数
package main

import (
	"fmt"
	"math"
)

// 判断传进来的数是否为水仙花数
func isDaffNum(num int) bool {
	hundred := num / 100
	decade := (num - hundred*100) / 10
	unit := (num - hundred*100 - decade*10)

	if math.Pow(float64(hundred), 3)+math.Pow(float64(decade), 3)+math.Pow(float64(unit), 3) == float64(num) {
		return true
	}
	return false
}

func main() {
	start := 100
	end := 999
	var daffNums []int

	for i := start; i <= end; i++ {
		if isDaffNum(i) {
			daffNums = append(daffNums, i)
		}
	}

	fmt.Printf("%d - %d 之间的水仙花数有 %d 个，它们是: %v", start, end, len(daffNums), daffNums)
}
