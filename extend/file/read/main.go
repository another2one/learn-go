package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

var (
	filename = "../test.txt"
)

func main() {

	// 文件读取（带缓冲区）
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("open error :", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)


	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Printf("str : %v \n", len(str))
	}

	// 文件读取（一次性）
	str, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(" ioutil read error :", err)
	}
	fmt.Printf("file content : \n%v", string(str))
}