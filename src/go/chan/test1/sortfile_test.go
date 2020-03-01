package test1

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"time"
)

// 检查
// 文件数目
// 每个文件随机抽取10%的数检查顺序是否正确，并列出错误文件
func TestMySort(t *testing.T) {

	Sort()
	
	files, err := ioutil.ReadDir("./json")
	if err != nil {
		t.Fatalf("read dir err: %v", err)
	}
	// 文件数目检查
	if len(files) != GOROUTING_NUM * 2 {
		for _, val := range files {
			t.Log(val.Name())
		}
		t.Fatalf("file num err, would %d, but %d", GOROUTING_NUM * 2, len(files))
	}
	resChan := make(chan string, GOROUTING_NUM)

	// 并发检查每个文件排序
	for i := 0; i < GOROUTING_NUM; i++ {
		go isFileSort(getFileName(NEW_NAME, i), resChan, t)
	}

	worngNumberSlice := make([]string, 0, GOROUTING_NUM)
	for i := 0; i < GOROUTING_NUM; i++ {
		s := <-resChan
		if len(s) > 0 {
			worngNumberSlice = append(worngNumberSlice, s)
		}
	}

	// 错误文件输出
	if len(worngNumberSlice) > 0{
		t.Fatalf("worng file %v", worngNumberSlice)
	}
}

// 检查排列是否正确
func isFileSort(fileName string, resChan chan<- string, t *testing.T){

	rand.Seed(time.Now().Unix())
	var intSlice []int

	// 读取文件
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		resChan<- fileName
		t.Errorf("%v read fail error: %v \n", fileName, err)
		return
	}
	err = json.Unmarshal(byteSlice, &intSlice)
	if err != nil {
		resChan<- fileName
		t.Errorf("%v json unmarshal error: %v \n", fileName, err)
		return
	}

	// 随机检查
	// t.Logf("intSlice len = %v, type is %T\n", len(intSlice), intSlice)
	for s := 0; s < int(NUMBER_NUM/10); s++ {
		i, j := rand.Intn(NUMBER_NUM), rand.Intn(NUMBER_NUM)
		if (i - j) * (intSlice[i] - intSlice[j]) < 0 {
			resChan<- fileName
			t.Errorf("%v sort error: %v \n", fileName, err)
			return
		}
	}
	resChan<- ""
}