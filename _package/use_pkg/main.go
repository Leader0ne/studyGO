package main

import (
	"fmt"

	"Leader0ne/studygolang/_package/math_pkg"
)

func main() {
	math_pkg.Add(1, 2)
	stu := math_pkg.Student{Name: "quin", Age: 18}
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}
