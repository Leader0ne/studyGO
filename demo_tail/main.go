package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf的用法示例

func main() {
	filename := "./my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个地方开始读
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
    var msg *tail.Line
	var ok bool
	for  {
		msg,ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filenam:%s\n",tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:",msg.Text)
	}
}
