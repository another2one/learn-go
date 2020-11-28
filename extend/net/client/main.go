package main

// 服务器端：
// 1. 监听端口
// 2. 结收并建立tcp链接
// 3. 创建 gorouting 处理请求

// 客户端
// 1. 建立与服务端tcp链接
// 2. 发送请求接收返回（可以长连接）
// 3. 关闭连接

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func reciveMsg(conn net.Conn){
	for{
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Printf("read error: %v \n", err)
			break
		}
		fmt.Printf("%v", string(b[:n]))
	}
}

func main() {

	// 建立连接
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("connect error: ", err)
	}
	defer conn.Close()

	go reciveMsg(conn)

	for {
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("write error: ", err)
			break
		}
		if strings.Trim(str, " \r\n") == "exit" {
			fmt.Println("byebye ... ")
			break
		}
		fmt.Printf("send : %q, len is %d \n", str, len(str) )
		conn.Write([]byte(str))
	}	
}