package main

import (
	"fmt"
	"reflect"
)

// 反射的TypeOf()
func reflectType(x interface{}) {
	t := reflect.TypeOf(x) // 拿到x的动态类型信息
	// fmt.Printf("type:%v\n", v)
	// fmt.Printf("%T\n", x) // 原理就是用反射 代码补全也是用反射
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	// 指针类型的t.Name()是空
}

type cat struct {
	name string
}

type person struct {
	name string
	age  int
}

func main() {
	// reflectType(100)
	// reflectType(false)
	// reflectType("透明无声系主播ywwuyi")
	// reflectType([3]int{1, 2, 3})
	// reflectType(map[string]int{})

	// 测试自定义结构体的类型
	var c1 = cat{
		name: "喵酱",
	}
	var p1 = person{
		name: "Quin",
		age:  16,
	}
	reflectType(c1)
	reflectType(p1)

	var i int32 = 100
	var f float32 = 12.34
	reflectType(&i)
	reflectType(&f)

	reflectType([]int{1, 2, 3})
	reflectType(map[string]int{})
}
