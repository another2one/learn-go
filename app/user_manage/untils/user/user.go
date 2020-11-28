package user

import (
	"fmt"
	"learn-go/app/user_manage/model/user"
)

var firstStep int // 第一次操作选择
var sureExit string // 是否确认离开

// 返回User结构体字段
func NewUser() UserSlice {
	return UserSlice{}
}

type UserSlice []*user.User

// 输出首页
func (userSlice UserSlice) firstPage(){
	fmt.Println("------------ 客户信息管理软件 -------------")
	fmt.Println("\t 1.添加客户")
	fmt.Println("\t 2.修改客户")
	fmt.Println("\t 3.删除客户")
	fmt.Println("\t 4.客户列表")
	fmt.Println("\t 5.退    出")
	fmt.Printf("请选择 1~5：")
	fmt.Scanln(&firstStep)
}

// 打印用户列表
func (userSlice UserSlice) PrintIncome() {
	fmt.Println("ID\t 姓名\t 性别\t 年龄\t 手机\t\t 邮箱")
	// for index, value := range userSlice {
	// 	// fmt.Printf("%d\t %v\t %.2f\t %.2f\t\t %v \n", index,)
	// }
}

// 添加用户
func (userSlice UserSlice) AddUser(amount float64, desc string) {
	userSlice = append(userSlice, &User{})
}

// 用户管理
func  (userSlice UserSlice) Manage(){

	main:
	for{

		userSlice.firstPage()

		if 1 > firstStep || 5 < firstStep {

			fmt.Println("输入错误，重新输入 ...")
			continue
		}

		switch firstStep {
			case 1:
				// if len(account.Logs) > 0 {
				// 	account.PrintIncome()
				// } else {
				// 	fmt.Println("没有数据")
				// }
			case 2:
				// fmt.Printf("请输入收入金额：")
				// fmt.Scanln(&amount)
				// fmt.Printf("请输入说明：")
				// fmt.Scanln(&desc)
				// account.Income(amount, desc)
			case 3:
				// if account.Balance <= 0 {
				// 	fmt.Println("你个小穷子，你 no money !!!")
				// 	break
				// }
				// for {
				// 	fmt.Printf("请输入支出金额：")
				// 	fmt.Scanln(&amount)
				// 	if account.checkAmount(amount) {
				// 		break
				// 	}else{
				// 		fmt.Println("你个小穷子，你只有", account.Balance, "元")
				// 	}
				// }
				// fmt.Printf("请输入说明：")
				// fmt.Scanln(&desc)
				// account.Outcome(amount, desc)
			case 4:
				
			case 5:
				fmt.Printf("是否确定退出？ 输入 y 继续退出：")
				fmt.Scanln(&sureExit)
				if sureExit == "y" {
					break main
				}
		}
	} 
}