package main

import (
	"bufio"
	"fmt"
	"learn-go/common/tool"
	model2 "learn-go/extend/file/excel/mysqlToExcel/model"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// 生成柱状图标
func main() {
	// 获取数据
	var err error
	file, err := os.OpenFile("./sql.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file error: %s", err)
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

	db := tool.MustGetDb(newLogger)

	// 遍历
	pageSize := 2000
	page := 0
	for {
		results := []map[string]any{}
		db.Model(&model2.LewaimaiOrderHistory{}).Offset(page * pageSize).Limit(pageSize).Find(&results)
		fmt.Sprintf("results == %v \n", results)
		if len(results) < pageSize {
			break
		}
		page++
	}

}
