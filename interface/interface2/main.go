package main

import (
	"fmt"
	"time"
)

// 空接口

// 空接口作为函数参数
func showType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {
	// var x interface{}
	// x = 100
	// fmt.Println(x)
	// x = "南昌"
	// fmt.Println(x)
	// x = false
	// fmt.Println(x)
	// x = struct{ name string }{name: "quin"}
	// fmt.Println(x)
	// showType(x)

	// 定义一个值为空接口的map
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["韦天"] = 100
	stuInfo["爽哥"] = true
	stuInfo["小秦"] = 22.22
	stuInfo["机核王"] = "播了"
	stuInfo["本尼"] = time.Now()
	fmt.Println(stuInfo)
}
