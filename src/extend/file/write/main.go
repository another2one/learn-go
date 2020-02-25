package main

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
)

func PathOrFileExists (name string) (bool, error) {
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
	file , err := os.OpenFile("../test1.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err : ", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	str := ""

	for i := 0; i < 5; i++ {

		writer.WriteString("hello Graden\r\n")
		str += "hello Graden2\r\n"
		
		// file.WriteString("hello Graden\n")	
	}

	writer.Flush()

	// 第二种
	ioutil.WriteFile("../test1.txt", []byte(str), 0666)
}