package main

import "fmt"

func main() {
	// 判断某个键是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["小赤"] = 100
	scoreMap["小绿"] = 200

	// 判断 大木博士 在不在 scoreMap中
	value, ok := scoreMap["小绿"]
	if ok {
		fmt.Println("小绿在scoreMap中", value)
	} else {
		fmt.Println("查无此人！")
	}

	// for range
	// map是无序的，键值对和添加的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历map中的key
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 只遍历map中的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// 删除 小绿 这个键值对
	delete(scoreMap, "小绿")
	fmt.Println(scoreMap)
}
