package main

import "fmt"

// 结构体是一个值类型
type student struct {
	name string
	age  int8
}

func main() {
	var stu1 = student{
		name: "韦天",
		age:  29,
	}

	stu2 := stu1 // 将结构体stu1完整的复制一份给stu2
	stu2.name = "爽哥"
	fmt.Println(stu1.name)
	fmt.Println(stu2.name)

	stu3 := &stu1 // 将stu1对应的地址赋值给stu3，stu3的类型是一个*student
	fmt.Printf("%T\n", stu3)
	(*stu3).name = "小秦"
	fmt.Println(stu1.name, stu2.name, stu3.name)
}
