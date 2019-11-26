// 获取格式化过的当前时间
package main

import (
	"fmt"
	"time"
)

func getCurTime() string {
	curTime := time.Now()
	return curTime.Format("2006/1/02 15:04:05")
}

func main() {
	fmt.Printf("current time: %s\n", getCurTime())
}
