// 两个大叔相加求和
package main

import (
	"fmt"
	"strconv"
)

func bigIntAdd(a string, b string) string {
	var (
		lenA   = len(a)
		lenB   = len(b)
		maxLen int
		minLen int
		max    string
		min    string
		sum    string
		isOver bool
	)

	const ZERO = 48

	if lenA > lenB {
		max = a
		min = b
		maxLen = lenA
		minLen = lenB
	} else {
		max = b
		min = a
		maxLen = lenB
		minLen = lenA
	}

	for i, _ := range max {
		var singleVal int
		// 没有达到小数的边界
		if i < minLen {
			maxV := int(max[maxLen-i-1]) - ZERO
			minV := int(min[minLen-i-1]) - ZERO
			singleVal = maxV + minV
		} else {
			// 达到小数边界
			singleVal = int(max[maxLen-i-1]) - ZERO
		}

		if isOver {
			singleVal++
			isOver = false
		}

		if singleVal > 10 {
			isOver = true
			singleVal -= 10
		}

		sum = strconv.Itoa(singleVal) + sum
	}

	return sum
}

func main() {
	fmt.Printf("2464654876879876746464698 + 2323215454445 的结果是: %s\n", bigIntAdd("2464654876879876746464698", "2323215454445"))
}
