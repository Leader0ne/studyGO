package main

import "fmt"

// 接收值时判断通道是否关闭

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 不一定要显式关闭
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	// 利用for循环去通道ch1取值
	// for {
	// 	ret, ok := <-ch1 // 使用 value, ok := <-ch1 取值方式，通道关闭时，ok = false
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(ret)
	// }
	for ret := range ch1 {
		fmt.Println(ret)
	}
}
