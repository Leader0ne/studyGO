package main

import (
	"fmt"
	"strings"
)

// 字符串操作
func main() {
	s1 := "quin mole"
	fmt.Println(len(s1))
	// 字符串拼接
	s2 := "halo"
	fmt.Println(s1 + s2)
	fmt.Println(fmt.Sprintf("%s----%s", s1, s2))
	// 分割
	ret := strings.Split(s1, " ")
	fmt.Println(ret)
	// 判断是否包含
	ret2 := strings.Contains(s1, "mo")
	fmt.Println(ret2)
	// 判断前缀和后缀
	ret3 := strings.HasPrefix(s1, "quin")
	ret4 := strings.HasSuffix(s1, "le")
	fmt.Println(ret3, ret4)
	// 求子串位置
	s3 := "abcabdcbd"
	fmt.Println(strings.Index(s3, "bd"))
	fmt.Println(strings.LastIndex(s3, "bd"))
	//Join
	a1 := []string{"黑楼", "黑旗", "黑暗剑"}
	fmt.Println(strings.Join(a1, "——"))
}
