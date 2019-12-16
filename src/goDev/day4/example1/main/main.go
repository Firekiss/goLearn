// 非递归编写斐波那契额数列前100个值
// https://studygolang.com/articles/9425
package main

func printFebs(num int) []uint64 {
	febs := make([]uint64, num)

	febs[0] = 1
	febs[1] = 1

	for i := 2; i < len(febs); i++ {
		febs[i] = febs[i-1] + febs[i-2]
	}

	return febs
}

func main() {
	// fmt.Printf("前1000个斐波那契数列是 %v\n", printFebs(1000))
}
