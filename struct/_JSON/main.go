package main

import (
	"encoding/json"
	"fmt"
)

// json序列化

// Student 学生
// 定义元信息
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var stu1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "列宁",
	}
	// 序列化：把编程语言里的数据转换成JSON字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON格式化出错了！")
		fmt.Println(err)
	}
	fmt.Println(v)                 // []byte
	fmt.Printf("%#v\n", string(v)) // 把byte转换成string

	str := string(v)
	// 反序列化：把满足JSON格式的字符串转化成 当前编程语言里的对象
	var stu2 = &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)
	fmt.Printf("%T\n", stu2)
}
