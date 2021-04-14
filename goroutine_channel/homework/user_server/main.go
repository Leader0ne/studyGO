package main

import (
	"studygolang/goroutine_channel/homework/mylogger"
)

// 测试mylogger的程序
var logger mylogger.Logger

func main() {
	logger = mylogger.NewFileLogger("debug", "xx.log", "./")
	defer logger.Close()
	for {
		logger.Debug("下雨了")
		logger.Error("打雷了")
	}

}
