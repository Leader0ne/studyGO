package main

import "fmt"

func main() {
	s1 := "Golang"
	c1 := 'G'
	fmt.Println(s1, c1)
	s2 := "中国"
	c2 := '中'
	fmt.Println(s2, c2)

	s := "hello天哥"
	// 遍历字符串
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	// for range 循环
	for _, r := range s { //rune
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()

	// 强制类型转换
	s5 := "big"
	byteArray := []byte(s5)
	byteArray[0] = 'p'
	fmt.Println(string(byteArray))

}
