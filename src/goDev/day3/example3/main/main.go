// 计算一个程序运行的时长 精确到微秒

package main

import (
	"fmt"
	"time"
)

func testedFunc() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	start := time.Now().UnixNano()
	testedFunc()
	end := time.Now().UnixNano()

	fmt.Printf("总共执行时间是 %d 微秒\n", (end-start)/1000)
}
