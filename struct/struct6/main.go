package main

import "fmt"

// 匿名字段
type student struct {
	name string
	string
	int
}

func main() {
	var stu1 = student{
		name: "韦天",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.int)
	fmt.Println(stu1.string)
}
