package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数

func main() {
	// 创建一个标志位参数
	name := flag.String("name", "晓樱", "请输入名字")
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "输入是否结婚")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")

	//var name string
	//flag.StringVar(&name,"name","Quin","请输入名字")

	// 使用flag
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数)  //返回命令行参数后的其他参数，以[]string类型
}
