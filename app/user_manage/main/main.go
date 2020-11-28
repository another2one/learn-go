package main

import (
	_ "fmt"
	"learn-go/app/user_manage/view/userview"
)

func main() {
	userview.NewUserview().MainView()
}