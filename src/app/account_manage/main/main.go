package main

import (
	_ "fmt"
	"app/account_manage/untils/account"
)

func main() {
	account.NewAccount().Manage()
}