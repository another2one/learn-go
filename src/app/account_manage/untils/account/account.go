package account

import (
	"fmt"
)

var firstStep int // 操作选择
var amount float64 // 输入金额
var desc string // 输入描述
var sureExit string // 是否确认离开
var IncomeType = map[int]string {
	1: "收入",
	2: "支出",
}

// 返回Account结构体字段
func NewAccount() *Account {
	return &Account{}
}

type IncomeLog struct {
	IncomeType int 
	Amount float64
	Balance float64
	Desc string
}

type Account struct {
	Logs []IncomeLog
	Balance float64
}

// 检查使用金额是否合法
func (account Account) checkAmount(amount float64) bool {
	return account.Balance > amount
}

// 输出首页
func (account Account) firstPage(){
	fmt.Println("------------ 家庭收支记账软件 -------------")
	fmt.Println("\t 1.收支明细")
	fmt.Println("\t 2.登记收入")
	fmt.Println("\t 3.登记支出")
	fmt.Println("\t 4.退    出")
	fmt.Printf("请选择 1~4：")
	fmt.Scanln(&firstStep)
}

// 打印收支明细
func (account Account) PrintIncome() {
	fmt.Println("ID\t 收支\t 余额\t 收支金额\t 说明")
	for index, value := range account.Logs {
		fmt.Printf("%d\t %v\t %.2f\t %.2f\t\t %v \n", index+1, IncomeType[value.IncomeType], value.Balance, value.Amount, value.Desc)
	}
}

// 登记收入
func (account *Account) Income(amount float64, desc string) {
	account.Balance+= amount
	account.Logs = append(account.Logs, IncomeLog{
		1,
		amount,
		account.Balance,
		desc,
	})
}

// 登记支出
func (account *Account) Outcome(amount float64, desc string) {
	account.Balance-= amount
	account.Logs = append(account.Logs, IncomeLog{
		2,
		amount,
		account.Balance,
		desc,
	})
}

// 用户管理
func (account *Account) Manage(){

	main:
	for{

		account.firstPage()

		if 1 > firstStep || 4 < firstStep {

			fmt.Println("输入错误，重新输入 ...")
			continue
		}

		switch firstStep {
			case 1:
				if len(account.Logs) > 0 {
					account.PrintIncome()
				} else {
					fmt.Println("没有数据")
				}
			case 2:
				fmt.Printf("请输入收入金额：")
				fmt.Scanln(&amount)
				fmt.Printf("请输入说明：")
				fmt.Scanln(&desc)
				account.Income(amount, desc)
			case 3:
				if account.Balance <= 0 {
					fmt.Println("你个小穷子，你 no money !!!")
					break
				}
				for {
					fmt.Printf("请输入支出金额：")
					fmt.Scanln(&amount)
					if account.checkAmount(amount) {
						break
					}else{
						fmt.Println("你个小穷子，你只有", account.Balance, "元")
					}
				}
				fmt.Printf("请输入说明：")
				fmt.Scanln(&desc)
				account.Outcome(amount, desc)
			case 4:
				fmt.Printf("是否确定退出？ 输入 y 继续退出：")
				fmt.Scanln(&sureExit)
				if sureExit == "y" {
					break main
				}
		}
	} 
}