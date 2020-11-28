package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	dir = "test/test.txt"
)

func PathOrFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {

	// 第一种
	file, err := os.OpenFile(dir, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err : ", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello Garden\r\n")
	}
	writer.Flush()

	//for i := 0; i < 5; i++ {
	//	file.WriteString("hello Garden\n")
	//}

	// 第二种
	str := ""
	for i := 0; i < 5; i++ {
		str += "hello Dog\r\n"
	}
	ioutil.WriteFile(dir, []byte(str), 0666)
}
