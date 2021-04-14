package main

import "fmt"

// interface

// Dog 结构体
type Dog struct{}

// Say 狗会叫
func (d Dog) Say() string { return "汪汪汪" }

// Cat 结构体
type Cat struct{}

// Say 猫会叫
func (c Cat) Say() string { return "喵喵喵" }

// Sayer 接口
type animal interface {
	Say() string
}

func main() {
	var animalList []animal
	c := Cat{}
	d := Dog{}
	animalList = append(animalList, c, d)
	fmt.Println(animalList)
}
