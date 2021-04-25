package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("[猜数字小游戏]（go语言版本） v1.0")
	min, max := 1, 100
	for {
		guess(min, max)
		if !again() {
			break
		}
	}
}

func again() bool {
	fmt.Printf("再玩一次好不好？(y/N):")
	s := input()
	return s == "y" || s == "Y"
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ""
	}
	return scanner.Text()
}

func guess(min int, max int) {
	//产生随机数字
	rand.Seed(time.Now().UnixNano())
	sn := rand.Intn(max-min+1) + min
	fmt.Printf("我已经想好了一个(%d~%d)之间的数字，请你来猜一猜。\n", min, max)

	for times := 1; ; times++ {
		fmt.Printf("请输入(%d~%d之间的数字<回车>退出)：\n", min, max)
		str := input()
		if len(str) < 1 {
			break
		}
		//读数字
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "格式错误: %v\n", err)
			continue
		}
		//判断是否猜中
		ok, text := comp(sn, n)
		fmt.Printf(" -- [第#%d次] 输入%d -> %s\n", times, n, text)
		if ok { //猜对了
			fmt.Printf("...我刚才默默想好的数字是%d，你一共试了%d次。\n", sn, times)
			break
		}
	}
}

func comp(n int, guessed int) (ok bool, text string) {
	switch {
	case n == guessed:
		return true, "=== 恭喜你，猜对了！==="
	case guessed < n:
		return false, "<<< 太小了 <<<"
	case guessed > n:
		return false, ">>> 太大了 >>>"
	default:
		return false, "???"
	}
}
