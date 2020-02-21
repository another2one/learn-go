package main

import (
	_ "fmt"
	"app/user_manage/view/userview"
)

func main() {
	userview.NewUserview().MainView()
}