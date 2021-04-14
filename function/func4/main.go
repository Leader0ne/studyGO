package main

import "fmt"

// 匿名函数和闭包
// 定义一个函数它的返回值是一个函数
func a(name string) func() {
	return func() {
		fmt.Println("Hello", name)
	}
}

func main() {

	// 执行方式1:赋值
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	// 执行方式2:定义完后直接加括号
	func() {
		fmt.Println("匿名函数")
	}()

	// 闭包 = 函数 + 外层变量的引用
	r := a("爽哥") // r此时就是一个闭包
	r()          // 相当于执行了a函数内部的匿名函数
}
