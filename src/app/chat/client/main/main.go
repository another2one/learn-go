package main

import (
	"app/chat/client/process"
	"fmt"
	"net"
)

var conn net.Conn

// 发送流程
// 1. 创建 msg 结构体 type 为消息类型, data 为序列化后的消息数据
// 2. 序列化 msg （打包）
// 3. 为防止丢包：1）先发送包长度 再发送包数据 2）msg 再封装一个长度

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect error: ", err)
		return
	}
	defer conn.Close()
	process.NewProcess(conn).Handel()
}
