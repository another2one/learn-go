package process

import (
	"app/chat/conf"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type UserProcess struct {
	Conn net.Conn
	Name string
	Id   int
}

func NewUserProcess(conn net.Conn) *UserProcess {
	return &UserProcess{Conn: conn}
}

// 登入
func (userProcess *UserProcess) Login() (err error) {

	// 2. 输入登入数据
	var loginData conf.Login
	fmt.Printf("请输入id：")
	fmt.Scanln(&loginData.UserId)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&loginData.PassWord)

	// 3. 发送消息数据
	smsProcess := NewSmsProcess(userProcess.Conn)
	data, _ := json.Marshal(loginData)
	err = smsProcess.SendMsg(conf.LoginType, string(data))
	if err != nil {
		return
	}

	// 4. 读取返回
	resMsg := <-MsgChan
	resData, err := userProcess.dealMsg(*resMsg)
	if err != nil {
		return
	}

	var loginReturn conf.LoginReturn
	err = json.Unmarshal([]byte(resData), &loginReturn)
	if err != nil {
		return conf.ErrorUnmarshal
	}
	userProcess.Name = loginReturn.Name
	userProcess.Id = loginReturn.Id
	fmt.Println(userProcess)
	return
}

// 注册
func (userProcess *UserProcess) Register() (err error) {

	// 2. 输入注册数据
	var registerData conf.Register
	fmt.Printf("请输入名字：")
	fmt.Scanln(&registerData.Name)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&registerData.PassWord)

	// 3. 发送消息数据
	smsProcess := NewSmsProcess(userProcess.Conn)
	data, _ := json.Marshal(registerData)
	err = smsProcess.SendMsg(conf.RegisterType, string(data))
	if err != nil {
		return
	}

	// 4. 读取返回
	resMsg := <-MsgChan

	_, err = userProcess.dealMsg(*resMsg)

	return
}

// 发送消息
func (userProcess *UserProcess) SendMsg() (err error) {

	// TODO: defer 记录消息

	var content string
	fmt.Println("id::content 表示私聊\t content 表示群发\t exit 离开")

	for {
		fmt.Scanln(&content)
		if strings.Contains(content, "::") { // 私聊
			contentSlice := strings.Split(content, "::")
			id, convertErr := strconv.Atoi(contentSlice[0])
			if convertErr != nil {
				fmt.Println("系统提示：输入不合法")
				continue
			}
			if id == userProcess.Id {
				fmt.Println("系统提示：不能发送给自己")
				continue
			}
			err = userProcess.SendUser(id, contentSlice[1])
			if err != nil {
				fmt.Println("系统发送出错：", err)
				continue
			}

		} else if content == "exit" { // 离开
			return nil
		} else { // 群聊
			err = userProcess.SendGroup(content)
		}
	}

}

// 发送私聊
func (userProcess *UserProcess) SendUser(id int, content string) (err error) {

	userMsg := conf.SendUser{
		Content: content,
		ToId:    id,
		FormId:  userProcess.Id,
		Time:    int(time.Now().Unix()),
	}
	userMsgBS, err := json.Marshal(userMsg)
	if err != nil {
		return
	}
	return NewSmsProcess(userProcess.Conn).SendMsg(conf.SendUserType, string(userMsgBS))
}

// 发送群聊
func (userProcess *UserProcess) SendGroup(content string) (err error) {

	userMsg := conf.SendGroup{
		Content: content,
		FormId:  userProcess.Id,
		Time:    int(time.Now().Unix()),
	}
	userMsgBS, err := json.Marshal(userMsg)
	if err != nil {
		return
	}
	return NewSmsProcess(userProcess.Conn).SendMsg(conf.SendGroupType, string(userMsgBS))
}

// 获取在线列表
func (userProcess *UserProcess) GetOnlineList() (err error) {

	// 发送消息数据
	smsProcess := NewSmsProcess(userProcess.Conn)
	err = smsProcess.SendMsg(conf.OnlineListType, "")
	if err != nil {
		return
	}

	// 读取并处理返回
	resMsg := <-MsgChan

	resData, err := userProcess.dealMsg(*resMsg)
	if err != nil {
		return
	}
	var list conf.GetOnlineListReturn
	err = json.Unmarshal([]byte(resData), &list)
	if err != nil {
		return conf.ErrorUnmarshal
	}
	for _, value := range list.UserList {
		fmt.Printf("%d号用户: %v \n", value.Id, value.Name)
	}
	return
}

// 处理消息体
func (userProcess *UserProcess) dealMsg(msg conf.Msg) (data string, err error) {

	switch msg.Type {

	case conf.ResultType:

		// 获取数据
		var resultData conf.Result
		err = json.Unmarshal([]byte(msg.Data), &resultData)
		if err != nil {
			err = conf.ErrorUnmarshal
			return
		}

		// 处理结果数据
		if resultData.Code == 200 {
			err = nil
			data = resultData.Data
		} else {
			err = errors.New(resultData.Msg)
		}

	default:
		err = conf.ErrorMsgType
	}

	return
}
