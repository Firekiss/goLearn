package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100)
	var inStr string

	for {
		fmt.Println("请输入你猜测的数据: ")
		fmt.Scanf("%s", &inStr)
		guessNum, err := strconv.Atoi(inStr)
		if err != nil {
			fmt.Println("请输入正确的数字!!")
		} else {
			if guessNum == answer {
				fmt.Printf("恭喜你猜对了，答案就是 %d ", answer)
				fmt.Println("=======游戏结束=======")
				return
			} else if guessNum < answer {
				fmt.Printf("猜错了，大于 %d ", guessNum)
			} else if guessNum > answer {
				fmt.Printf("猜错了，小于 %d ", guessNum)
			}
		}
	}
}
