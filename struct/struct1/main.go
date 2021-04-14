package main

import "fmt"

// 自定义类型
type newInt int

// 类型别名:只存在代码编写过程中，代码编译后根本不存在MyInt
// 提高代码可读性
type myInt = int

func main() {
	var a newInt
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	var b myInt
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c byte
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
