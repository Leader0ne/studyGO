package main

import "fmt"

func main() {
	// 切片的赋值拷贝
	a := make([]int, 3) //[0, 0, 0]
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	// 切片的遍历
	n := []int{1, 2, 3, 4, 5}
	// 根据索引来遍历
	for i := 0; i < len(n); i++ {
		fmt.Println(i, n[i])
	}
	fmt.Println()
	// for range 遍历
	for index, value := range n {
		fmt.Println(index, value)
	}
}
