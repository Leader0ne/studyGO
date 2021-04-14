package main

import "fmt"

// 从终端获取用户的输入内容
func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println(name, age, married)
	// fmt.Scan(&name, &age, &married)
	// fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Println(name, age, married)
}
