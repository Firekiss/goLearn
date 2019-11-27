// 输入一行字符，分别统计出其中英文字母，空格，数字和其他字符的个数
package main

import "fmt"

func arrangeChars(str string) (enCharCount, emptyCharCount, numCharCount, otherCharCount int) {
	enCharCount = 0
	emptyCharCount = 0
	numCharCount = 0
	otherCharCount = 0

	for _, char := range str {
		switch {
		case (char >= 65 && char <= 90) || (char >= 97 && char <= 122):
			enCharCount++
		case char == 32:
			emptyCharCount++
		case char >= 48 && char <= 57:
			numCharCount++
		default:
			otherCharCount++
		}
	}

	return
}

func main() {
	enCharCount, emptyCharCount, numCharCount, otherCharCount := arrangeChars("34mb **()da90 ")
	fmt.Printf("'34mb **()da90 ' 英文字符个数是:%d 空格字符个数是:%d 数字字符个数是:%d 其他字符个数是%d\n", enCharCount, emptyCharCount, numCharCount, otherCharCount)
}
