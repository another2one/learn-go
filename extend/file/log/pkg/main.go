package main

import (
	"log"
	"os"
)

func main() {
	// flags
	// const (
	//	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	//	Ltime                         // the time in the local time zone: 01:23:23
	//	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	//	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	//	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	//	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	//	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	//	LstdFlags     = Ldate | Ltime // initial values for the standard logger
	//)
	// 准备日志文件
	logFile, _ := os.Create("demo.log")
	defer func() { _ = logFile.Close() }()

	// 初始化日志对象
	logger := log.New(logFile, "[Debug] - ", log.Lshortfile|log.Lmsgprefix)
	logger.Print("Print\n")
	logger.Println("Println")

	// 修改日志配置
	logger.SetOutput(os.Stdout)
	logger.SetPrefix("[Info] - ")
	logger.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	logger.Print("Print\n")
	logger.Println("Println")
}
