package utils

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn-go/app/http_hello_world/model"
	"learn-go/common/tool"
	"log"
	"strconv"
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("mysql connect error: ", err)
	}
	//Db.SetMaxOpenConns(50)
}

func Query() {
	user := model.User{
		Name: "lizhi",
	}
	Db.Create(&user)
	user.Name = "lipan"
	user.ID = 0
	Db.Create(&user)
	user.ID = 0
	user.Name = "lihan"
	Db.Create(&user)

	if err := Db.Where("name = ?", "lizhi").First(&user).Error; err != nil {
		Logger.Errorf("query error: %v", err)
	}
	fmt.Printf("user = %+v \n", user)

	if err := Db.Table("user").Where("id > 0").Update("name", gorm.Expr("concat(name, id)")).Error; err != nil {
		Logger.Errorf("update error: %v", err)
	}

	var users []model.User
	if err := Db.Find(&users).Error; err != nil {
		Logger.Errorf("query error: %v", err)
	}
	var data [][]string
	data = append(data, []string{"id", "name", "created_at", "updated_at"})
	for _, user := range users {
		data = append(data, []string{strconv.Itoa(int(user.ID)), user.Name, user.CreatedAt.Format("2006-01-02 15:04:05"), user.UpdatedAt.Format("2006-01-02 15:04:05")})
	}
	tool.PrettyPrint("users", data)
}
