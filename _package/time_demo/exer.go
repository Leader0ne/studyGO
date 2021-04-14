package main

import (
	"fmt"
	"time"
)

// time包练习

// 获取当前时间，格式输出为2017/06/19 20:30:05 格式
// 统计一段代码执行耗时，精确到微秒
func printTime(t time.Time) {
	form := t.Format("2006/01/02 15:04:05")
	fmt.Println(form)
}

func main() {
	s := time.Now()
	start := time.Now().UnixNano()
	now := time.Now()
	printTime(now)
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000)
	fmt.Println(time.Since(s))
}
