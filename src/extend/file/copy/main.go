package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	_ "io/ioutil"
	"log"
	"time"
	_ "strconv"
	"strings"
)

var (
	readfile = "E:/360下载/图片/8aef83dfb5b87eced98c4e39af35fa96.jpg"
	// writefile = "E:/360下载/图片/8aef83dfb5b87eced98c4e39af35fa96_" + strconv.Itoa(int(time.Now().Unix())) + ".jpg"
)

func getNewFilename(name string) string {

	lastIndex := strings.LastIndex(name, ".")
	if lastIndex == -1 {
		log.Fatal("文件中不含.")
	}
	// return strings.Join([]string{name[:lastIndex], time.Now().Format("20060102150405"), name[lastIndex:]}, "_")
	return name[:lastIndex] + "_" + time.Now().Format("20060102150405") + name[lastIndex:]
}

// 复制文件
func copyFile(readfile string) (int64, error) {

	// 获取reader
	rfile, err := os.Open(readfile)
	if err != nil {
		log.Fatal("open readfile err :", err)
	}
	defer rfile.Close()
	reader := bufio.NewReader(rfile)

	// 获取writer
	wfile, err := os.OpenFile(getNewFilename(readfile), os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("open writefile err :", err)
	}
	defer wfile.Close()
	writer := bufio.NewWriter(wfile)

	return io.Copy(writer, reader)
}

// 获取目录下文件
func getDirfile(dir string)  (names []string, err error) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal("open readfile err :", err)
	}
	defer file.Close()
	return file.Readdirnames(20)
}

func main() {

	num, err := copyFile(readfile)
	if err != nil {
		log.Fatal("open readfile err :", err)
	}
	fmt.Println(num)

	fileInfo, err := getDirfile("../../..")
	if err != nil {
		log.Fatal("open readfile err :", err)
	}
	for _, value := range fileInfo {
		fmt.Printf("%+v\n", value)
	}
	fmt.Printf("%+v\n", fileInfo)
}