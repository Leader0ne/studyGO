package main

import "fmt"

const pi = 3.14

//批量声明常量
const (
	a = 100
	b = 1000
	c
	d
)

//iota
const (
	aa = iota //0
	bb = iota //1
	cc = iota //2
	dd = iota //3
)

const (
	n1 = iota //0
	n2 = 100
	n3 = iota //2
	n4        //3
)

const n5 = iota //0

//常量
func main() {
	// pi = 3.1415   //常量不允许修改
	fmt.Println(pi)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(n1, n2, n3, n4, n5)
}

/*
 * iota
 * 0.const中声明如果不写，默认就和上一行一样
 * 1.遇到const初始化为零
 * 2.const中每新增一行变量声明iota就递增
 */
