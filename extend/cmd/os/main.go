package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"
)

func main() {

	var s int
	flag.IntVar(&s, "s", 5, "inter seconds default 5")

	flag.Parse()

	// 定时器 + 清除
	inter1 := make(chan bool, 1)
	fmt.Println("time inter == ", s)
	go intervalTimeExec(time.Duration(s)*time.Second, inter1, func() {
		fmt.Println("interval")
	})

	// 阻塞等待终止信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	msg := <-c
	fmt.Println("Got signal:", msg)
}

type ByteWriter struct {
	b []byte
}

func (bw *ByteWriter) Write(p []byte) (n int, err error) {
	bw.b = append(bw.b, p...)
	return len(p), nil
}

func intervalTimeExec(intervalTime time.Duration, interval chan bool, f func()) {
loop:
	for {
		select {
		case <-interval:
			break loop
		default:
			cmd := exec.Command("go", "version")
			bytes, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("exec error: %s \n", err)
			} else {
				log.Printf("exec result: %s \n", string(bytes))
			}

			// ps -ef | grep -E "nginx|php"
			if runtime.GOOS == "linux" {
				cmd2 := exec.Command("sh", "-c", "ps -ef | grep -E \"nginx|php\"")
				bytes2, err := cmd2.CombinedOutput()
				if err != nil {
					log.Printf("exec error: %s \n", err)
				} else {
					log.Printf("exec result: %s \n", string(bytes2))
				}
			}

			//c1 := exec.Command("netstat", "-an")
			//c2 := exec.Command("findstr", "9001")
			//c2.Stdin, _ = c1.StdoutPipe()
			//file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND, 0666)
			//c2.Stdout = file
			//_ = c2.Start()
			//_ = c1.Run()
			//_ = c2.Wait()
			//file.Close()

			c1 := exec.Command("netstat", "-an")
			c2 := exec.Command("findstr", "9001")
			c2.Stdin, _ = c1.StdoutPipe()
			var bufBytes ByteWriter
			file := bufio.NewWriter(&bufBytes)
			c2.Stdout = file
			_ = c2.Start()
			_ = c1.Run()
			_ = c2.Wait()
			file.Flush()
			fmt.Printf("%s \n", bufBytes)
			f()
			time.Sleep(intervalTime)
		}
	}
}

func clearInterval(interval chan bool) {
	interval <- true
}
