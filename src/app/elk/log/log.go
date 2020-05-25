package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file error: ", err)
		return
	}

	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail lines fail:")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", msg.Text)
	}
}
