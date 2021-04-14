package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志

type logData struct {
	Message  string
	LogLevel string
	LineNo   int
	TimeStr  string
	FuncName string
	FileName string
}

// FileLogger 文件日志结构体
type FileLogger struct {
	level       Level
	fileName    string
	filePath    string
	file        *os.File
	errFile     *os.File
	maxSize     int64
	logDataChan chan *logData // 定义一个存放日志信息的通道
}

// NewFileLogger 文件日志结构体构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:       logLevel,
		fileName:    fileName,
		filePath:    filePath,
		maxSize:     10 * 1024 * 1024,
		logDataChan: make(chan *logData, 50000), // 通道初始化
	}
	fl.initFile() // 根据文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应字段
	return fl
}

// 将指定的日志文件打开 赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	// 打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v", logName, err))
	}
	f.file = fileObj
	// 打开错误日志的文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v", errLogName, err))
	}
	f.errFile = errFileObj
	// 开启goroutine去写日志
	go f.writeLogBackground()
}

// 检查是否要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize // 当传进来的日志文件的大小超过maxSize 就返回true
}

// 封装切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	// 切分文件
	fileName := file.Name() // 获得文件的完整路径
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	// 1. 把原来的文件关闭
	file.Close()
	// 2. 备份原来的文件
	os.Rename(fileName, backupName)
	// 3. 新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	// f.file.Write()
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志
	// 日志格式：[时间][文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	logLevelStr := getLevelStr(level)

	// 将日志信息发送到通道中，已备写日志的goroutine接收值
	// 1. 构造日志数据结构体
	logdata := &logData{
		Message:  msg,
		LogLevel: logLevelStr,
		LineNo:   line,
		TimeStr:  nowStr,
		FuncName: funcName,
		FileName: fileName,
	}
	select {
	case f.logDataChan <- logdata:
	default:
		// <-f.logDataChan          // 丢弃最早的一条
		// f.logDataChan <- logdata // 把最新的发送进去
	}

}

// 真正往文件里面写日志的函数
func (f *FileLogger) writeLogBackground() {

	for data := range f.logDataChan {

		level := parseLogLevel(data.LogLevel)

		logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
			data.TimeStr, data.FileName, data.LineNo, data.FuncName, data.LogLevel, data.Message)
		// 往文件里写之前做检查
		if f.checkSplit(f.file) {
			f.file = f.splitLogFile(f.file)
		}
		fmt.Fprintln(f.file, logMsg) //利用fmt包将msg写入f.file文件中
		// 如果是error或者fatal级别的日志还要记录到 f.errFile
		if level >= ErrorLevel {
			if f.checkSplit(f.errFile) {
				f.errFile = f.splitLogFile(f.errFile)
			}
			fmt.Fprintln(f.errFile, logMsg)
		}
	}
}

// Debug debug方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info info方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn warn方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error error方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatal fatal方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Close 关闭日志文件句柄
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
