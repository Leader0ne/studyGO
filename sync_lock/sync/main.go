package main

import (
	"fmt"
	"sync"
)

var x int64 // 定义全局变量x
var wg sync.WaitGroup

// 定义一个互斥锁
var lock sync.Mutex

// 定义一个函数 对全局的变量x做累加的操作
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 打开锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
