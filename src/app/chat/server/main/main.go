package main

import (
	"app/chat/server/process"
	"fmt"
	"net"
)

func main() {

	// 创建监听
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen port 8888 error: ", err)
	} else {
		fmt.Println("listening port 8888 ... ")
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connect error: ", err)
		}
		go process.NewProcess(conn).Handle()
	}
}
