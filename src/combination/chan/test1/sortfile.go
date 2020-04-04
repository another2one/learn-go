package test1

// 并发n个协程，随机写入m个数并放入n个文件
// 并发n个协程，读取上面n个文件并排序再放入文件

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	NUMBER_NUM    = 100000 // 生成的数字个数
	GOROUTING_NUM = 10     // 并发协程个数
	DIR           = "./json/"
	FILE_NAME     = DIR + "test"
	NEW_NAME      = DIR + "testSort"
)

var wg sync.WaitGroup

// 根据编号获取文件名称
func getFileName(name string, i int) string {
	return name + strconv.Itoa(i+1) + ".json"
}

func WriteDataTofile(writechan chan<- int, i int) {

	var intArray [NUMBER_NUM]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < NUMBER_NUM; i++ {
		intArray[i] = rand.Intn(2000)
	}
	strByteArray, err := json.Marshal(intArray)
	if err != nil {
		log.Fatalln("json err: ", err)
	}
	err = ioutil.WriteFile(getFileName(FILE_NAME, i), strByteArray, 0666)
	if err != nil {
		log.Fatalln("write file err: ", err)
	}
	fmt.Printf("第%d个文件写入完成 \n", i)
	writechan <- i
}

func sortData(i int) {

	defer wg.Done()
	strByteArray, err := ioutil.ReadFile(getFileName(FILE_NAME, i))
	if err != nil {
		log.Fatalln("read file err: ", err)
	}
	var intSlice []int
	err = json.Unmarshal(strByteArray, &intSlice)
	if err != nil {
		log.Fatalln("json err: ", err)
	}
	// fmt.Println("before sort: ", intArray)
	sort.Ints(intSlice)
	// fmt.Println("after sort: ", intArray)
	strByteArray, err = json.Marshal(intSlice)
	if err != nil {
		log.Fatalln("json err: ", err)
	}

	file, err := os.OpenFile(getFileName(NEW_NAME, i), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("open file err: ", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(strByteArray))
	writer.Flush()

	// err = ioutil.WriteFile(getFileName(NEW_NAME, i), strByteArray, 0666)
	// if err != nil {
	// 	log.Fatalln("write file err: ", err)
	// }
	// fmt.Printf("第%d个文件排序完成完成 \n", i)
}

// 判断是否为目录
func IsDirExist(f string) bool {
	flleInfo, err := os.Stat(f)
	return (err == nil || os.IsExist(err)) && flleInfo.IsDir()
}

// 清空目录
func clearDir(s string) {
	if IsDirExist(s) {
		os.RemoveAll(s)
	}
	os.MkdirAll(s, 0666)
}

func Sort() {

	clearDir(DIR)

	writeChan := make(chan int, GOROUTING_NUM)

	for i := 0; i < GOROUTING_NUM; i++ {
		go WriteDataTofile(writeChan, i)
	}

	for i := 0; i < GOROUTING_NUM; i++ {
		wg.Add(1)
		go sortData(<-writeChan)
	}

	wg.Wait()
}
