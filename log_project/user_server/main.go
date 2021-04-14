package main

import "studygolang/log_project/mylogger"

var logger mylogger.Logger

// 一个使用自定义日志库的用户程序
func main() {
	logger = mylogger.NewFileLogger("Info", "./", "xxx.log")
	// logger = mylogger.NewConsoleLogger("debug")
	defer logger.Close()
	// for {
	yw := "爽哥"
	logger.Debug("%s开麦了", yw)
	logger.Info("Info 这是一条测试日志")
	logger.Error("Error 这是一条测试日志")
	// }
}
