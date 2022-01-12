package funcs

import (
	"io/ioutil"
	"os"
)

//判断文件夹是否存在
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

// readFile 读取文件
func ReadFile(path string) string {
	str, err := ioutil.ReadFile(path)
	if err == nil {
		return string(str)
	}
	return ""
}
