package process

import (
	"app/chat/conf"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Process struct {
	Conn net.Conn
	id   int
}

func NewProcess(conn net.Conn) *Process {
	return &Process{Conn: conn}
}

func (process *Process) Handle() {

	defer process.Close()

	for {

		// 获取消息 （阻塞）
		msg, err := NewSmsProcess(process.Conn).ReadMsg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client close the connection ...")
				return
			} else {
				fmt.Println("read error: ", err)
				return
			}
		}

		// 处理消息
		err = process.dealMsg(msg)
		if err != nil {
			fmt.Println("deal message error: ", err)
		}
	}
}

// 结束process
func (process *Process) Close() {
	// 关闭连接
	process.Conn.Close()
	// 移除在线列表
	_, err := userManage.getUserProcess(process.id)
	if err == nil {
		userManage.RemoveOnline(process.id)
	}
}

// 处理消息体
func (process *Process) dealMsg(msg conf.Msg) (err error) {

	switch msg.Type {

	case conf.LoginType: // 登入

		// 获取用户输入
		var loginData conf.Login
		err = json.Unmarshal([]byte(msg.Data), &loginData)
		if err != nil {
			err = conf.ErrorUnmarshal
			return
		}

		// 登入
		userProcess := NewUserProcess(process.Conn)
		user, resErr := userProcess.Login(loginData)
		if resErr == nil {
			process.id = user.Id
		}
		loginRes := conf.LoginReturn{
			Id:   user.Id,
			Name: user.Name,
		}

		dataBS, _ := json.Marshal(loginRes)
		err = NewSmsProcess(process.Conn).SendMsg(resErr, string(dataBS))

	case conf.RegisterType: // 注册

		// 获取用户输入
		var registerData conf.Register
		err = json.Unmarshal([]byte(msg.Data), &registerData)
		if err != nil {
			err = conf.ErrorUnmarshal
			return
		}

		// 注册
		userProcess := NewUserProcess(process.Conn)
		user, resErr := userProcess.Register(registerData)
		dataBS, _ := json.Marshal(user)
		err = NewSmsProcess(process.Conn).SendMsg(resErr, string(dataBS))

	case conf.OnlineListType: // 获取在线列表

		onlineList, err := NewUserProcess(process.Conn).GetOnlineList()
		dataBS, _ := json.Marshal(onlineList)
		err = NewSmsProcess(process.Conn).SendMsg(err, string(dataBS))

	case conf.SendUserType: // 私聊

		var sendUserMsg conf.SendUser
		err = json.Unmarshal([]byte(msg.Data), &sendUserMsg)

		if err != nil {
			err = conf.ErrorUnmarshal
			return
		}
		userProcess, err := userManage.getUserProcess(sendUserMsg.ToId)
		if err != nil {
			return err
		}
		err = NewSmsProcess(userProcess.Conn).SendUser(msg.Data)

	case conf.SendGroupType: // 群聊

		var sendGroupMsg conf.SendGroup
		err = json.Unmarshal([]byte(msg.Data), &sendGroupMsg)

		if err != nil {
			err = conf.ErrorUnmarshal
			return
		}
		userProcess, err := userManage.getUserProcess(sendGroupMsg.FormId)
		if err != nil {
			return conf.ErrorUserNotOnline
		}
		msgBS, _ := json.Marshal(msg)
		err = userProcess.PushGroupMsgToOtherOnline(msgBS)

	default:
		fmt.Printf("%+v \n", msg)
		err = conf.ErrorMsgType
	}

	return
}
