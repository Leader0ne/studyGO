package main

import (
	"fmt"

	_ "Leader0ne/studygolang/_package/math_pkg"
)

// init()示例
var today = "Sunday"

const Week = 7

type student struct {
	Name string
	Age  int
}

// 包被导入时会自动执行(多用来做初始化操作)
func init() {
	fmt.Println("包被初始化了。")
	fmt.Println(Week)
}

func main() {
	fmt.Println("这是main函数。")
}
