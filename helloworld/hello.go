package main

import (
	"fmt"
	"os"
	"time"
)

type Person struct {
	name string
	id   string
	age  int
}

//birthYear 计算出生年份
func (p *Person) birthYear() int {
	return time.Now().Year() - p.age
}

func main() {
	fmt.Printf("现在的时间:%s\n", time.Now())
	fmt.Println("命令行参数回显")
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = "|"
	}
	fmt.Println(s)

	fmt.Println("系统环境变量")
	for i, env := range os.Environ() {
		fmt.Printf("[%d]:%s\n", i, env)
	}

	//验证能否指定基础类型的新方法
	p := Person{
		"",
		"",
		28,
	}
	fmt.Printf("birth year=%d", p.birthYear())
	html.
}
