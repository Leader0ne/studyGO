package main

import "fmt"

type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// 自己实现一个构造函数
func newStudent(n string, age int, g string, h []string) *student {
	return &student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
}

func main() {
	hobbySlice := []string{"重庆开车", "大老爹拿球", "黑楼黑旗黑暗剑"}
	quin := newStudent("小秦", 22, "男", hobbySlice)
	fmt.Println(quin.name)
}
