package main

import "fmt"

func main() {

	// nil
	var a []int     // 声明int类型切片
	var b = []int{} //声明并初始化
	c := make([]int, 0)
	if a == nil {
		fmt.Println("a是一个nil")
	}
	fmt.Println(a, len(a), cap(a))
	if b == nil {
		fmt.Println("b是一个nil")
	}
	fmt.Println(b, len(b), cap(b))
	if c == nil {
		fmt.Println("c是一个nil")
	}
	fmt.Println(c, len(c), cap(c))
}
