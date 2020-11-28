package main

import (
	_ "fmt"
	"learn-go/app/account_manage/untils/account"
)

func main() {
	account.NewAccount().Manage()
}
