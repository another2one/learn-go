package process

import (
	"learn-go/app/chat/conf"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type User struct {
	LoginStatus int
	Id          int
	Name        string
}

var MsgChan chan *conf.Msg
var OnlineMap map[int]*User

func init() {
	MsgChan = make(chan *conf.Msg)
	OnlineMap = make(map[int]*User)
}

type Process struct {
	Conn        net.Conn
	buf         [4096]byte
	LoginStatus int
	userProcess *UserProcess
}

func NewProcess(conn net.Conn) *Process {
	return &Process{Conn: conn}
}

func (process *Process) Handel() {
	go process.receiveMsg()
	process.showIndex()
}

// 这里统一接收后如何分发到需要的地方 如登入注册
// 方案：这里循环接收服务器数据，如果是推送就自行处理，否则放入管道等待其他程序使用
func (process *Process) receiveMsg() {

	for {
		msg, err := NewSmsProcess(process.Conn).ReadMsg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("牛逼, 强行离开系统！！！")
			} else {
				fmt.Println(err)
			}
			break
		}
		switch msg.Type {

		case conf.UserStatusNotifyType: // 用户状态通知

			var userStatusNotify conf.UserStatusNotify
			err = json.Unmarshal([]byte(msg.Data), &userStatusNotify)

			if err != nil {
				err = conf.ErrorUnmarshal
				fmt.Println(err)
			}

			OnlineMap[userStatusNotify.Id] = &User{
				LoginStatus: userStatusNotify.Status,
				Id:          userStatusNotify.Id,
				Name:        userStatusNotify.Name,
			}

			fmt.Printf("来自服务端的消息: %d 号用户（%v）上线了 ！！！ \n", userStatusNotify.Id, userStatusNotify.Name)

		case conf.SendUserType: // 私聊消息

			var sendUserMsg conf.SendUser
			err = json.Unmarshal([]byte(msg.Data), &sendUserMsg)

			if err != nil {
				err = conf.ErrorUnmarshal
				fmt.Println(err)
			}

			fmt.Printf("来自 %d 的消息：%v \n", sendUserMsg.FormId, sendUserMsg.Content)

		case conf.SendGroupType: // 群聊消息

			var sendGroupMsg conf.SendGroup
			err = json.Unmarshal([]byte(msg.Data), &sendGroupMsg)

			if err != nil {
				err = conf.ErrorUnmarshal
				fmt.Println(err)
			}

			fmt.Printf("%d 对大家说：%v \n", sendGroupMsg.FormId, sendGroupMsg.Content)

		default:
			MsgChan <- &msg
		}
	}
}

func (process *Process) showIndex() {

	viewProcess := NewViewProcess()
	process.userProcess = NewUserProcess(process.Conn)

	for {

		// 显示首页
		viewProcess.IndexPage()

		switch viewProcess.Key {

		case 1: // 登入
			if process.LoginStatus != conf.LoginStatusOnline {
				err := process.userProcess.Login()
				if err != nil {
					fmt.Println("login error: ", err)
				} else {
					process.LoginStatus = conf.LoginStatusOnline
					fmt.Println("login success !!!")
					process.showServer()
				}
			} else {
				process.showServer()
			}

		case 2: // 注册
			err := process.userProcess.Register()
			fmt.Println("register error: ", err)

		case 3: // 退出
			if viewProcess.SureExit() {
				return
			}

		default:
			fmt.Println("输入有误")
		}
	}
}

// 显示客户服务
func (process *Process) showServer() {

	viewProcess := NewViewProcess()

	for {
		viewProcess.ServerPage(process.userProcess.Name)

		switch viewProcess.Key {

		case 1: // 在线列表
			err := process.userProcess.GetOnlineList()
			if err != nil {
				fmt.Println("getOnlineList error: ", err)
			}
			//process.showOnlineList()

		case 2: //
			process.userProcess.SendMsg()

		case 3: // 退出
			return

		default:
			fmt.Println("输入有误")
		}
	}

}

func (process *Process) showOnlineList() {
	for id, value := range OnlineMap {
		if value.LoginStatus == conf.LoginStatusOnline {
			fmt.Printf("%d: %v \n", id, value.Name)
		}
	}
}
