package main

import (
	"fmt"
	"sync"
)

// 启动goroutine
var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Quin")
}

func main() {
	wg.Add(1)
	go hello() // 1. 创建一个goroutine 2. 在新的goroutine中执行函数
	fmt.Println("Hello main func.")
	// time.Sleep(time.Second)
	// 等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait() // 阻塞，一直等待所有的goroutine结束
}
