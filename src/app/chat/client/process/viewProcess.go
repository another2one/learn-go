package process

import "fmt"

type ViewProcess struct {
	MenuLen int
	Key     int
}

var indexMenu = [4]string{"欢迎进入多人聊天室", "登入聊天室", "注册用户", "退出系统"}
var serverMenu = [4]string{"登入系统", "在线人员列表", "聊天", "退出"}

func NewViewProcess() *ViewProcess {
	return &ViewProcess{}
}

// 首页
func (viewProcess *ViewProcess) IndexPage() {
	viewProcess.showPage(indexMenu[:], "")
	fmt.Scanln(&viewProcess.Key)
}

// 服务页
func (viewProcess *ViewProcess) ServerPage(name string) {
	viewProcess.showPage(serverMenu[:], name)
	fmt.Scanln(&viewProcess.Key)
}

func (viewProcess *ViewProcess) showPage(menu []string, name string) {

	viewProcess.MenuLen = len(menu) - 1
	fmt.Println(name)
	for index, value := range menu {
		if index == 0 {
			if len(name) > 0 {
				fmt.Println("恭喜 ", name, value)
			} else {
				fmt.Println(value)
			}
		} else {
			fmt.Printf("%d %v \n", index, value)
		}
	}
	fmt.Printf("请选择 1 - %d : \n", viewProcess.MenuLen)
}

// 确认离开
func (viewProcess *ViewProcess) SureExit() bool {
	var s string
	fmt.Printf("确认离开？输入Y/N : ")
	fmt.Scanln(&s)
	if s == "y" || s == "Y" {
		return true
	}
	return false
}
