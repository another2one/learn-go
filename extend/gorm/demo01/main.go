package main

import (
	"fmt"
	"learn-go/extend/gorm/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// TODO: profile
func main() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect db err: ", err)
	}
	defer db.Close()
	fmt.Printf("%v \n", db)
	find("lizhi", "11129996211")

	// add
	animal := model.Credit{Name: "demo-test", Phone: "11129996211"}
	db.Create(&animal)

	find("lizhi", "11129996211")
}

func find(name, phone string) {
	c := new(model.Credit)
	db.Where("name = ? and phone = ?", name, phone).First(c)
	if c.Id > 0 {
		fmt.Printf("%+v \n", c)
	} else {
		fmt.Println("not found")
	}
}
