package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// template 示例
type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {
	// 打开一个模板文件
	htmlByte, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 添加自定义的方法要在parse模板文件之前添加

	// 1. 自定义一个函数
	// 自定义一个夸人的模板函数
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 2. 把自定义的函数告诉模板系统
	//template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}) // 追加自定义函数

	// 链式操作
	// 原理：每一次执行完方法之后返回操作对象本身
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("parse html file failed, err:", err)
		return
	}
	// 用数据去渲染模板
	// data := "《我的世界》"
	// t.Execute(w, data)

	//user := User{
	//	UserName: "Quin",
	//	Password: "1234",
	//	Age:      16,
	//}
	//
	//t.Execute(w, user)
	userMap := map[int]User{
		1: {"Quin", "1234", 16},
		2: {"韦天", "12341", 28},
		3: {"ywwuyi", "12345", 32},
	}
	//userSlice := []User{
	//	{"Quin", "1234", 16},
	//	{"韦天", "12341", 28},
	//	{"ywwuyi", "12345", 32},
	//}
	t.Execute(w, userMap)
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
