package main

import (
	"fmt"
)

// 1.4 版本里面新添加的一个命令，当运行该命令时，它将扫描与当前包相关的源代码文件，找出所有包含//go:generate的特殊注释，提取并执行该特殊注释后面的命令
//
//  注意：
//  1. 该特殊注释必须在 .go 源码文件中；
//  2. 每个源码文件可以包含多个 generate 特殊注释；
//  3. 运行go generate命令时，才会执行特殊注释后面的命令；
//  4. 当go generate命令执行出错时，将终止程序的运行；
//  5. 特殊注释必须以//go:generate开头，双斜线后面没有空格。
//
//  场景：
//  1. yacc：从 .y 文件生成 .go 文件；
//  2. protobufs：从 protocol buffer 定义文件（.proto）生成 .pb.go 文件；
//  3. Unicode：从 UnicodeData.txt 生成 Unicode 表；
//  4. HTML：将 HTML 文件嵌入到 go 源码；
//  5. bindata：将形如 JPEG 这样的文件转成 go 代码中的字节数组。
//
//  格式： go generate [-run regexp] [-n] [-v] [-x] [command] [build flags] [file.go... | packages]
//  参数说明如下：
//  -run 正则表达式匹配命令行，仅执行匹配的命令；
//  -v 输出被处理的包名和源文件名；
//  -n 显示不执行命令；
//  -x 显示并执行命令；
//  command 可以是在环境变量 PATH 中的任何命令。
//  环境变量：
//  $GOARCH 体系架构（arm、amd64 等）；
//  $GOOS 当前的 OS 环境（linux、windows 等）；
//  $GOFILE 当前处理中的文件名；
//  $GOLINE 当前命令在文件中的行号；
//  $GOPACKAGE 当前处理文件的包名；
//  $DOLLAR 固定的$，不清楚具体用途。
//go:generate go run main.go
//go:generate go version
func main() {
	s := "ssss"
	a := s
	fmt.Println(a)
}
