package main

import (
	"fmt"
)

// 切片（slice）
func main() {

	// 基于数组的到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// 切片再次切片
	c := b[0:]
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	// make函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	// 通过len()函数获取切片的长度
	fmt.Println(len(d))
	// 通过cap()函数获取切片的容量
	fmt.Println(cap(d))

}
