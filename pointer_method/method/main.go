package main

import "fmt"

// 函数是所有人都可以调用的
// 方法(method) 特定类型才能使用的函数
type people struct {
	name   string
	gender string
}

// 函数指定接受者之后就是方法
// 在go语言中约定成俗不用this也不用self,而是使用后面类型的首字母的小写
func (p *people) dream() {
	p.gender = "男"
	fmt.Printf("%s梦想是睡觉赚钱\n", p.name)
}

func main() {
	var ywwuyi = &people{
		name:   "爽哥",
		gender: "处男",
	}
	ywwuyi.dream()
	fmt.Println(ywwuyi.gender)
}
