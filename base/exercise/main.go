package main

import (
	"fmt"
)

func main() {
	// exercise
	// "hello" ——> "olleh"
	// 方法1
	// s1 := "hello"
	// byteS1 := []byte(s1)
	// l := len(byteS1)
	// s2 := ""
	// for i := l - 1; i >= 0; i-- {
	// 	s2 = s2 + string(byteS1[i])
	// }
	// fmt.Println(s2)

	// 方法2
	// for i := 0; i < l/2; i++ {
	// 	byteS1[i], byteS1[l-1-i] = byteS1[l-1-i], byteS1[i]
	// }
	// fmt.Println(string(byteS1))

	// 打印200到1000的素数
	// left := 101
	// right := 200
	// for num := left; num <= right; num++ {
	// 	isPrime := true
	// 	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
	// 		if num%i == 0 {
	// 			isPrime = false
	// 			break
	// 		}
	// 	}
	// 	if isPrime == true {
	// 		fmt.Printf("%d ", num)
	// 	}
	// }

	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%d ", j, i, i*j)
			}

		}
	}

}
