package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	for range 5 {
		writer.WriteString("hello Garden\r\n")
	}
	writer.Flush()

	//for i := 0; i < 5; i++ {
	//	file.WriteString("hello Garden\n")
	//}

	// 第二种
	var str strings.Builder
	for range 5 {
		str.WriteString("hello Dog\r\n")
	}
	os.WriteFile(dir, []byte(str.String()), 0666)
}
