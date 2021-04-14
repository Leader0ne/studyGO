package main

import "fmt"

// 结构体
// 创造新的类型要使用type关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var quin = student{
		name:   "小秦",
		age:    27,
		gender: "男",
		hobby:  []string{"ps4", "pc", "switch"},
	}

	// 结构体支持.访问属性
	fmt.Println(quin)
	fmt.Println(quin.name)
	fmt.Println(quin.age)
	fmt.Println(quin.gender)
	fmt.Println(quin.hobby)

	// 实例化方法1
	// struct是值类型的
	// 如果初始化时没有给属性（字段）设置对应初始值，那个对应属性就是其类型的默认值
	var WeiTian = student{}
	fmt.Println(WeiTian.name)
	fmt.Println(WeiTian.age)
	fmt.Println(WeiTian.gender)
	fmt.Println(WeiTian.hobby)

	// 实例化方法2 new(T) T:表示类型或结构体
	var ShuangGe = new(student)
	fmt.Println(ShuangGe)
	// (*ShuangGe).name
	ShuangGe.name = "爽哥"
	ShuangGe.age = 18
	fmt.Println(ShuangGe.name, ShuangGe.age)
	// 实例化方法3:取结构体的地址实例化
	var stan = &student{}
	fmt.Println(stan)
	stan.name = "屎蛋"
	fmt.Println(stan.name)

	// 结构体初始化
	// 只填值初始化
	var stu1 = student{
		"Trump",
		74,
		"男",
		[]string{"手风琴", "发推特"},
	}
	fmt.Println(stu1.name, stu1.age)
	// 键值对初始化
	var stu2 = &student{
		name:   "Trump",
		gender: "男",
	}
	fmt.Printf("%T\n", stu1)
	fmt.Printf("%T\n", stu2)
	fmt.Println(stu2.name, stu2.age, stu2.gender)
}
