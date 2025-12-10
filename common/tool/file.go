package tool

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

const ProjectPath = "/Users/lizhi/Desktop/app/go/learn-go/"
const DateTimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"
const TimeFormat = "15:04:05"
const TimeWithoutSecondFormat = "15:04"
const DateTimeWithoutSecondFormat = "2006-01-02 15:04"
const DateTimeWithMillSecondFormat = "2006-01-02 15:04:05.000"
const DateTimeWithMicroSecondFormat = "2006-01-02 15:04:05.000000"

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

func GetImageDimensions(filePath string) (width, height int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	return width, height, nil
}

// WriteExcel map数组写入excel
// data 格式 [["1", "2", "3", ...] ...]
func WriteExcel[T any](data [][]T, file string) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	f := excelize.NewFile()
	for rowIndex, row := range data {
		for columnIndex, value := range row {
			writeCell(f, rowIndex, columnIndex, value)
		}
	}
	return f.SaveAs(file)
}

func writeCell(f *excelize.File, rowIndex int, columnIndex int, value interface{}) {
	cell := GetCell(columnIndex, rowIndex)
	if err := f.SetCellValue("Sheet1", cell, value); err != nil {
		log.Printf("write excel cell = [%s] valie = [%s] err: %s \n", cell, value, err)
	}
}

func ReadExcel(file string) ([][]string, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, fmt.Errorf("read excel file [%s] err: %s", file, err)
	}

	return f.GetRows("Sheet1")
}
