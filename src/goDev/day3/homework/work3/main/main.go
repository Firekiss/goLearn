// 判断一个字符串是否为回文 "abcdcba"
package main

import (
	"fmt"
	"math"
)

func isSymmetryStr(str string) bool {
	strLen := len(str)
	strMidLen := int(math.Floor(float64(strLen / 2)))

	for i := 0; i <= strMidLen; i++ {
		if str[i] != str[strLen-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%s 是回文吗 => %v\n", "abcdcba", isSymmetryStr("abcdcba"))
	fmt.Printf("%s 是回文吗 => %v\n", "23344332", isSymmetryStr("23344332"))
	fmt.Printf("%s 是回文吗 => %v\n", "3ertyre3", isSymmetryStr("3ertyre3"))
}
