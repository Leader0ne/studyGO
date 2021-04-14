package main

import "fmt"

func main() {
	// 切片的copy
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	c := b     // 浅拷贝
	copy(b, a) // 深拷贝
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
