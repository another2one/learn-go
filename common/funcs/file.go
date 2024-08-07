package funcs

import (
	"os"
)

const ProjectPath = "D:/app/go/learn/"

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadFile 读取文件
//   - path 文件路径
func ReadFile(path string) string {
	str, err := os.ReadFile(path)
	if err == nil {
		return string(str)
	}
	return ""
}
