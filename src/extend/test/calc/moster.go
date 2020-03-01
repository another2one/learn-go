package calc

import (
	"encoding/json"
	"fmt"
	_"os"
	_"io"
	"io/ioutil"
	_"bufio"
)

type Monster struct {
	Name string 
	Age int
	Skill []string
}

func NewMonster() *Monster {
	return &Monster{}
}

func (this Monster) Store (filename string) bool {
	// 序列化
	mJson, err := json.Marshal(this)
	if err != nil {
		fmt.Println("json 序列化错误：", err)
		return false
	}
	// 保存文件
	// file, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println("打开文件错误：", err)
	// 	return false
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(file)
	// writer.WriteString(string(mJson))
	// writer.Flush()
	err = ioutil.WriteFile(filename, mJson, 0666)
	if err != nil {
		fmt.Println("json 写入错误", err)
		return false
	}

	return true
}

func (this Monster) ReStore (filename string) Monster {
	// 读取文件
	// file, err := os.Open(filename)
	// if err != nil {
	// 	fmt.Println("打开文件错误：", err)
	// }
	// defer file.Close()
	// reader := bufio.NewReader(file)
	// str := ""
	// for {
	// 	strTemp, err := reader.ReadString('\n')
	// 	str+= strTemp
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	mbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("json 读取错误", err)
	}
	// 反序列化
	var m Monster
	// json.Unmarshal([]byte(str), &m)	
	json.Unmarshal(mbyte, &m)	
	return m
}