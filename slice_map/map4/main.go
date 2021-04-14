package main

import "fmt"

func main() {
	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) // 只是完成了切片的初始化
	// [nil, nil, nil, nil, nil, nil, nil, nil]
	fmt.Println(mapSlice[0] == nil)
	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8) //完成了map的初始化

	mapSlice[0]["大木博士"] = 1000
	fmt.Println(mapSlice)

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		// sliceMap中没有中国这个词
		sliceMap["中国"] = make([]int, 8, 8) // 完成了切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}
	// 遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
