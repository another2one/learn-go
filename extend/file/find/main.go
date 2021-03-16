package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 目录下文件查找

func main() {

	str := "main"
	findPath := "D:/go/learn/extend/file"

	// 递归查找文件
	filepath.Walk(findPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		MatchFile(info, str)
		return nil
	})

	// 查找当前目录
	f, err := os.OpenFile(findPath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	files, _ := f.Readdir(-1)
	for _, v := range files {
		MatchFile(v, str)
	}
}

func MatchFile (info os.FileInfo, str string) {
	if !info.IsDir() && strings.Contains(info.Name(), str) {
		fmt.Println(info.Name())
	}
}