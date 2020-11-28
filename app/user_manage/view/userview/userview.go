package userview

import (
	"fmt"
	"learn-go/app/user_manage/service/userservice"
	"learn-go/app/user_manage/model/user"
)

type UserView struct {
	UserService *userservice.UserService
	firstStep int // 第一次操作选择
	sureExit string // 是否确认离开
}

func NewUserview() *UserView {
	return &UserView{userservice.NewUserService(), 0, "n"}
}

// 输出首页
func (userView *UserView) firstPage(){
	fmt.Println("------------ 客户信息管理软件 -------------")
	fmt.Println("\t 1.添加客户")
	fmt.Println("\t 2.修改客户")
	fmt.Println("\t 3.删除客户")
	fmt.Println("\t 4.客户列表")
	fmt.Println("\t 5.退    出")
	fmt.Printf("请选择 1~5：")
	fmt.Scanln(&userView.firstStep)
}

// 打印用户列表
func (userView UserView) PrintUser(){

	fmt.Println("ID\t 姓名\t 性别\t 年龄\t 手机\t 邮箱")
	for _, value := range userView.UserService.GetList() {
		fmt.Printf("%d\t %v\t %v\t %v\t %v \n", value.Id, value.Name, value.Sex, value.Telephone, value.Email)
	}
}

// 增加用户
func (userView *UserView) AddUser(){

	user := user.NewUser()
	fmt.Printf("输入姓名：")
	fmt.Scanln(&user.Name)
	fmt.Printf("输入性别：")
	fmt.Scanln(&user.Sex)
	fmt.Printf("输入年龄：")
	fmt.Scanln(&user.Age)
	fmt.Printf("输入手机：")
	fmt.Scanln(&user.Telephone)
	fmt.Printf("输入邮箱：")
	fmt.Scanln(&user.Email)
	userView.UserService.Add(user)
	fmt.Println("添加成功")
}

// 增加用户
func (userView *UserView) EditUser(){

	editId := 0
	fmt.Printf("输入修改人Id：")
	fmt.Scanln(&editId)

	index := userView.UserService.Search(editId)
	if index == -1 {
		fmt.Println("用户不存在")
	} else {
		if userView.Confirm() {
			user := &userView.UserService.UserSlice[index]
			var name string
			fmt.Printf("输入修改后名称：")
			fmt.Scanln(&name)
			if name != "" {
				user.Name = name
			}
			fmt.Println("修改成功")
		}
	}
}


// 删除用户
func (userView *UserView) DeleteUser(){

	deleteId := 0
	fmt.Printf("输入删除人Id：")
	fmt.Scanln(&deleteId)

	index := userView.UserService.Search(deleteId)
	if index == -1 {
		fmt.Println("用户不存在")
	} else {
		if userView.Confirm() {
			userView.UserService.Delete(index)
			fmt.Println("删除成功")
		}
	}
}

// 确认操作
func (userView UserView) Confirm() bool{

	fmt.Printf("是否确定操作？ 输入 y 继续：")
	fmt.Scanln(&userView.sureExit)
	if userView.sureExit == "y" {
		return true
	}else{
		return false
	}
}



func (userView *UserView) MainView() {

	main:
	for{

		userView.firstPage()

		if 1 > userView.firstStep || 5 < userView.firstStep {

			fmt.Println("输入错误，重新输入 ...")
			continue
		}

		switch userView.firstStep {
			case 1:
				userView.AddUser()
			case 2:
				userView.EditUser()
			case 3:
				userView.DeleteUser()
			case 4:
				userView.PrintUser()
			case 5:
				if userView.Confirm() {
					break main
				}
		}
	} 
}