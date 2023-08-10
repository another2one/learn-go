package main

import (
	"bufio"
	"fmt"
	"learn-go/extend/tool/gorm/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// https://learnku.com/docs/gorm/v2
func main() {
	var err error
	file, err := os.OpenFile("./sql.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file error:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	newLogger := logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	// db, err = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open(mysql.Open("root:123456@/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("connect db err: ", err)
	}
	fmt.Printf("%v \n", db)
	find("lizhi", "11129996211")

	// add
	animal := model.Credit{Name: "lizhi", Phone: "11129996211"}
	db.Create(&animal)

	find("lizhi", "11129996211")
}

func find(name, phone string) {
	c := new(model.Credit)
	db.Where("name = ? and phone = ?", name, phone).Last(c)
	if c.Id > 0 {
		fmt.Printf("%+v \n", c)
	} else {
		fmt.Println("not found")
	}
}
