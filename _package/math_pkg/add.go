package math_pkg

import "fmt"

// Student 学生
type Student struct {
	Name string
	Age  int
}

func init() {
	fmt.Println("这个包导入并初始化")
}

// Add 加法
func Add(x, y int) int {
	return x + y
}
