package main

import "fmt"

func foo() (string, int) {
	return "alax", 9000
}

func main() {
	var name string
	var age int
	// 声明变量必须要使用
	fmt.Println(name)
	fmt.Println(age)

	var (
		a string //""
		b int    //0
		c bool   //false
		d string //""
	)
	a = "摸鱼"
	b = 100
	c = true
	d = "100"
	fmt.Println(a, b, c, d)
	//声明变量并初始化
	var x string = "quin"
	fmt.Println(x)
	fmt.Printf("%s摸了%d", x, b)
	//类型推导(根据变量初始值推导)
	var y = 200
	var z = true
	fmt.Println(y)
	fmt.Println(z)
	//简短变量声明(只能在函数内部使用)
	friend := "韦天" // var friend string = "韦天"
	fmt.Println(friend)

	//调用foo函数
	//用于接收不需要的值
	aa, _ := foo()
	fmt.Println(aa)
}
