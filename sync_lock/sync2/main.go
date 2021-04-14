package main

import (
	"fmt"
	"sync"
	"time"
)

// 读比写多的时候要使用读写锁 能够提高性能

var x int64
var wg sync.WaitGroup
var lock sync.Mutex     // 互斥锁
var rwlock sync.RWMutex // 读写互斥 : 多个goroutine同时读加的是读锁 写的时候加的是写锁

func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.RLock() // 加读锁
	time.Sleep(time.Millisecond * 1)
	rwlock.RUnlock() //释放读锁
	// lock.Unlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock() // 加写锁
	// lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock() // 释放写锁
	// lock.Unlock()

}

func main() {
	start := time.Now()
	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("耗费了%v.\n", end.Sub(start))
}
