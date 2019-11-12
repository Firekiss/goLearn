package main

import "fmt"

// 字符串反转方法
func reverse(str string) string {
	var ret string
	var strIdx int
	strLen := len(str)

	for strIdx = 1; strIdx <= strLen; strIdx++ {
		ret = ret + fmt.Sprintf("%c", str[strLen-strIdx])
	}

	return ret
}

func reverseByArr(str string) string {
	var ret []byte
	tmp := []byte(str)
	len := len(tmp)

	for i := 1; i <= len; i++ {
		ret = append(ret, tmp[len-i])
	}
	return string(ret)
}

func main() {
	str := "fuck"
	str = reverse(str)
	fmt.Printf("%s\n", str)
	str = reverseByArr(str)
	fmt.Printf("%s\n", str)
}
