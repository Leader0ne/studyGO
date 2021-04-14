package main

import "fmt"

// 结构体内嵌模拟“继承”
type animal struct {
	name string
}

// 定义一个动物会动的方法
func (a *animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type dog struct {
	feet int
	*animal
}

// 定义一个狗的方法
func (d *dog) bark() {
	fmt.Printf("%s 在叫：汪！\n", d.name)
}

func main() {
	var a = dog{
		feet: 4,
		animal: &animal{
			name: "柯基",
		},
	}
	a.move()
	a.bark()
}
