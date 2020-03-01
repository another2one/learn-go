package main

import (
	"fmt"
	"os"
	"flag"
)

// 命令行参数

func main() {

	// 第一种  os.Args
	fmt.Printf("%+v \n", os.Args)

	// 第二种
	var name,pwd string
	var port int
	flag.StringVar(&name, "name", "", "请输入用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "请输入密码，默认为空")
	flag.IntVar(&port, "port", 3306, "请输入端口，默认为3306")
	flag.Parse()
	fmt.Printf("name = %v, pwd = %v, port = %v \n", name, pwd, port)
}