package main

import "fmt"

// 结构体的嵌套
type address struct {
	province string
	city     string
}

type email struct {
	province string
}

type student struct {
	name    string
	age     int
	address //嵌套了别的结构体
	email   //嵌入匿名结构体
}

func main() {
	var stu1 = student{
		name: "quin",
		age:  16,
		address: address{
			province: "Guangdong",
			city:     "Guangzhou",
		},
	}
	fmt.Println(stu1)
	fmt.Println(stu1.name)
	// fmt.Println(stu1.province) // 匿名字段支持直接访问
	fmt.Println(stu1.address.province) // 当匿名字段有冲突时必须显式调用
	fmt.Println(stu1.email.province)   // 当匿名字段有冲突时必须显式调用
}
