package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// bufio读数据
func readByLine() {
	file, err := os.Open("../open_read_close/xx.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	// 利用缓冲区从文件读数据
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 字符
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		fmt.Print(str)
	}
}

// ioutil读取文件
func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	readByLine()
	fmt.Println()
	// ioutil
	readFile("../open_read_close/xx.txt")
}
