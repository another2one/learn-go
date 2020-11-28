package process

import (
	"learn-go/app/chat/conf"
	"learn-go/app/chat/server/model"
	"learn-go/app/chat/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	Name string
	Id   int
}

func NewUserProcess(conn net.Conn) *UserProcess {
	return &UserProcess{Conn: conn}
}

func (userProcess *UserProcess) Login(login conf.Login) (user *model.User, err error) {

	user, err = model.NewUserDao().GetUserById(login.UserId)

	if err != nil {
		return &model.User{}, err
	}

	if user.PassWord() == login.PassWord {
		err = userProcess.afterLogin(user)
		return
	} else {
		return &model.User{}, conf.ErrorUserPasswordFail
	}
}

// 登入成功操作
func (userProcess *UserProcess) afterLogin(user *model.User) (err error) {
	userProcess.Id = user.Id
	userProcess.Name = user.Name
	// 加入在线列表
	//return utils.AddOnline(userProcess)
	err = userManage.AddOnline(userProcess)
	if err != nil {
		return
	}
	// 推送给其他在线用户
	err = userProcess.PushStatusMsgToOtherOnline(fmt.Sprintf("%d 号%v用户上线啊!!!", user.Id, user.Name))
	return
}

// 推送用户状态消息给其他在线用户
func (userProcess *UserProcess) PushStatusMsgToOtherOnline(message string) (err error) {
	userProcesses, err := userManage.GetOnlineList()
	if err != nil {
		return
	}
	serverMsg := conf.UserStatusNotify{Id: userProcess.Id, Status: 1, Name: userProcess.Name}
	serverMsgBS, err := json.Marshal(serverMsg)
	if err != nil {
		return conf.ErrorMarshal
	}
	msgBS, err := utils.GetServerPushData(string(serverMsgBS), conf.UserStatusNotifyType)
	if err != nil {
		return
	}
	for id, process := range userProcesses {
		if id == userProcess.Id {
			continue
		}
		pushErr := NewSmsProcess(process.Conn).ServerPush(msgBS)
		if pushErr != nil {
			fmt.Printf("服务端推送给%d号失败\n", id)
		}
	}
	return
}

// 推送用户状态消息给其他在线用户
func (userProcess *UserProcess) PushGroupMsgToOtherOnline(msgBS []byte) (err error) {

	userProcesses, err := userManage.GetOnlineList()

	if err != nil {
		return
	}

	for id, process := range userProcesses {
		if id == userProcess.Id {
			continue
		}
		pushErr := NewSmsProcess(process.Conn).ServerPush(msgBS)
		if pushErr != nil {
			fmt.Printf("服务端推送给%d号失败\n", id)
		}
	}
	return
}

func (userProcess *UserProcess) Register(registerData conf.Register) (user *model.User, err error) {

	user, err = model.NewUserDao().AddUser(registerData)

	if err != nil {
		return &model.User{}, err
	}

	return
}

func (userProcess *UserProcess) GetOnlineList() (onlineList conf.OnlineList, err error) {

	userProcesses, _ := userManage.GetOnlineList()
	onlineList.UserList = make([]conf.UserInfo, len(userProcesses))
	i := 0
	for _, userProcess := range userProcesses {
		onlineList.UserList[i].Name = userProcess.Name
		onlineList.UserList[i].Id = userProcess.Id
		i++
	}
	return
}
