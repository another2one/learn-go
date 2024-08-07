package main

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"learn-go/common/funcs"
	model2 "learn-go/extend/file/excel/mysqlToExcel/model"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

// 生成柱状图标
func main() {
	// 获取数据
	var err error
	file, err := os.OpenFile("./sql.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file error:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	newLogger := logger.New(
		log.New(writer, "\r\n", log.Ldate), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	db := funcs.MustGetDb(newLogger)

	// 遍历
	pageSize := 2000
	page := 0
	for {
		results := []map[string]interface{}{}
		db.Model(&model2.LewaimaiOrderHistory{}).Offset(page * pageSize).Limit(pageSize).Find(&results)
		err := writeExcel(results, funcs.ProjectPath+"extend/file/excel/mysqlToExcel/lwm_order"+strconv.Itoa(page+1)+".xlsx")
		if err != nil {
			fmt.Println("写入excel出错：", err)
			break
		}
		if len(results) < pageSize {
			break
		}
		page++
	}

}

func writeExcel(results []map[string]interface{}, path string) error {
	line := 1
	f := excelize.NewFile()
	for _, result := range results {
		var start byte = 'A'
		for _, v := range result {
			f.SetCellValue("Sheet1", getCell(start)+strconv.Itoa(line), v)
			//fmt.Println(getCell(start)+strconv.Itoa(line), " ---- ", v)
			start++
		}
		line++
	}

	// 保存文件
	if err := f.SaveAs(path); err != nil {
		return err
	}
	return nil
}

func getCell(num byte) string {
	if num < 'A' {
		panic("参数错误 不能小于A")
	}
	if num < 'Z' {
		return byteToStr(num)
	}
	return byteToStr(byte(math.Floor(float64(num-'Z')/26+float64('A')))) + byteToStr((num-'Z')%26+'A')
}

func byteToStr(num byte) string {
	return string([]byte{num})
}
