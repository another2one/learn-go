package main

import (
	"os"
	_ "fmt"
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
	return strings.Join([]string{name[:lastIndex], time.Now().Format("20060102150405"), name[lastIndex:]}, "_")
}

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

func main() {

	copyFile(readfile)
}