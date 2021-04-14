package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 0b11111
	var a int = 10
	var b int = 077
	var c int = 0xff
	fmt.Println(x, a, b, c)
	fmt.Printf("%b  \n", x)
	fmt.Printf("%o  \n", b)
	fmt.Printf("%x  \n", c)
	//求c变量的内存地址
	fmt.Printf("%p  \n", &c)

	//浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}
