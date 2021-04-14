package main

import "fmt"

func main() {
	// 切片的扩容
	var a []int // 此时它并没有申请内存
	// for i := 0; i < 10; i++ {
	// 	a = append(a, i)
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a)
	// }
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	b := []int{12, 13, 14, 15}
	a = append(a, b...)
	fmt.Println(a)

	// 切片删除元素
	city := []string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen"}
	city = append(city[0:2], city[3:]...)
	fmt.Println(city)

	// append(a[:index], a[index+1:]...)

}
