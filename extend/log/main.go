package main

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {

	if fileHandler, err := os.OpenFile(getLogFileName("info"), os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic(err)
	} else {
		InfoLog = log.New(fileHandler, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if fileHandler, err := os.OpenFile(getLogFileName("error"), os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic(err)
	} else {
		ErrorLog = log.New(io.MultiWriter(fileHandler, os.Stdout), "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

// 日期为目录， 错误类型+时间 为文件名
func getLogFileName(logType string) string {

	dir := "log/" + time.Now().Format("2006/01/")
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0666); err != nil {
			log.Fatal("create fir error: ", err)
		}
	}
	return dir + time.Now().Format("02") + "-" + logType + ".txt"
}

func main() {

	InfoLog.Print("666")
	ErrorLog.Print("666")
}
