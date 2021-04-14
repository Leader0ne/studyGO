package main

import "fmt"

func main() {
	// 错误写法
	// var a *int // a是一个int类型的指针
	// var b *string
	// var c *[3]int
	var a = new(int) // 得到一个int类型的指针
	fmt.Println(a)
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	fmt.Println(*c)
}
