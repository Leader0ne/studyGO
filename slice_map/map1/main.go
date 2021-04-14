package main

import "fmt"

// map(映射)
func main() {
	// 光声明map类型 但是没有初始化 a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil)
	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// map中添加键值对
	a["中山大学"] = 100
	a["复旦大学"] = 150
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)
	// 声明map同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)
}
