package main

import "fmt"

// select 练习
func main() {
	var ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case ret := <-ch:
			fmt.Println(ret)
		}
	}
}
