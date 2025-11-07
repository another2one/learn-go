package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn-go/common/funcs"
)

type User struct {
	gorm.Model
	Name string
	Role string
	Age  int
}

func main() {

	db, err := gorm.Open(sqlite.Open(funcs.ProjectPath+"extend/tool/gorm/mysql01/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("connect db err: " + err.Error())
	}
	db.AutoMigrate(&User{})
	user := User{Name: "lizhi", Role: "admin", Age: 18}
	// 大部分 CRUD API 都是兼容的
	//db.Create(&user)
	db.First(&user, 1)
	fmt.Printf("user = %v \n", user)
	db.Model(&user).Update("Age", 18)
	db.Model(&user).Omit("Role").Updates(map[string]interface{}{"Name": "jinzhu", "Role": "admin"})
	//db.Delete(&user)
}
