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
	"fmt"
	"net"
	"strconv"
	"strings"
)

type User struct {
	Id      int
	Conn    net.Conn
	Name    string
	Status  bool        // 在线状态
	msgChan chan string // 消息
	// CurrentChat *ChatObj // 当前聊天对象
}

// type ChatObj struct {
// 	Type int // 1为人 2为群组
// 	Uid int
// 	Gid int
// }

// type Group struct {
// 	Id int
// 	Name string
// 	num int
// 	UidSlice []int // 包含的用户id
// }

var (
	UserSlice = make([]*User, 0)
)

const (
	LIMIT_MSG = 10 // 连续无处理消息数
)

// 发送消息
func sendMsg(i int) {
	for msg := range UserSlice[i].msgChan {
		UserSlice[i].Conn.Write([]byte(strings.Trim(msg, " \r\n") + "\n"))
	}
	fmt.Printf("%v的通道关闭\n", getNameById(i))
}

// 检查用户合法性
func checkUserStatus(i int) bool {
	if i < len(UserSlice) {
		return UserSlice[i].Status
	}
	return false
}

// 处理链接
func handleConnection(i int) {

	defer closeConn(i)

	conn := UserSlice[i].Conn

	fmt.Println("RemoteAddr: ", conn.RemoteAddr().String())

	go sendMsg(i)

	for {

		// 阻塞等待接收信息
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Printf("%d read error: %v \n", i, err)
			break
		}

		// 获取信息
		msg := string(b[:n])

		// 处理信息
		dealMsg(msg, i)
	}
}

// 关闭连接
func closeConn(i int) {
	name := getNameById(i)
	fmt.Printf("%v: close connect ... \n", name)
	UserSlice[i].Conn.Close()
	fmt.Printf("%v: close connect ... \n", name)
	UserSlice[i].Status = false
}

// 通过id获取姓名
func getNameById(i int) string {
	return UserSlice[i].Name
}

// 处理消息
func dealMsg(msg string, i int) {

	conn := UserSlice[i].Conn
	msgSlice := strings.Split(msg, "::")
	name := getNameById(i)

	if len(msgSlice) == 2 {

		// 指令操作
		sendTo, err := strconv.Atoi(msgSlice[0])
		if err == nil && checkUserStatus(sendTo) {
			if i == sendTo {
				// 发送给自己
				UserSlice[i].msgChan <- "send to yourself: " + msgSlice[1]
			} else {
				// 发送给其他用户
				if len(UserSlice[sendTo].msgChan) >= LIMIT_MSG {
					// 发送数目太多
					conn.Write([]byte("大哥!你能不能别发了\n"))
				} else {
					UserSlice[sendTo].msgChan <- fmt.Sprintf("recive from %v : %v", name, msgSlice[1])
				}
			}
		} else if msgSlice[0] == "name" {
			// 修改名字
			UserSlice[i].Name = strings.Trim(msgSlice[1], " \r\n")
			UserSlice[i].msgChan <- "修改名字成功: " + UserSlice[i].Name
		} else {
			// 发送给服务器
			fmt.Printf("server recive from %v: %v \n", name, msg)
		}
	} else {
		// 发送给服务器
		fmt.Printf("server recive from %v: %v \n", name, msg)
	}
}

func main() {

	// 创建监听
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println("conn error:", err)
	}
	fmt.Println("success, wait connect ...")

	i := 0

	for {

		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Accept error:", err)
			continue
		} else {
			fmt.Printf("\n第%d个连接成功 \n", i)
		}

		// chanTemp := make(chan string, LIMIT_MSG)

		UserSlice = append(UserSlice, &User{i, conn, strconv.Itoa(i), true, make(chan string, LIMIT_MSG)})

		fmt.Printf("UserSlice = %v \n", UserSlice)

		go handleConnection(i)

		conn.Write([]byte("send from server: you number is " + strconv.Itoa(i) + "\n"))

		i++
	}
}
