package conf

const (
	LoginType            = "login"
	RegisterType         = "register"
	ResultType           = "result"
	OnlineListType       = "onlineList"
	UserStatusNotifyType = "userStatusNotify"
	SendGroupType        = "sendGroup"
	SendUserType         = "sendUser"
)

const (
	LoginStatusOffline = iota
	LoginStatusOnline
)

// 消息传输类型
type Msg struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 登入
type Login struct {
	UserId   int    `json:"user_id"`
	PassWord string `json:"pass_word"`
}

// 注册
type Register struct {
	Name     string `json:"name"`
	PassWord string `json:"pass_word"`
}

// 在线用户信息
type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 在线列表
type OnlineList struct {
	UserList []UserInfo `json:"user_list"`
}

// 返回
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`  // 信息
	Data string `json:"data"` // 数据
}

// 登入信息推送
type UserStatusNotify struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
}

// 群聊
type SendGroup struct {
	Content string `json:"content"`
	FormId  int    `json:"form_id"`
	Time    int    `json:"time"`
}

// 私聊
type SendUser struct {
	Content string `json:"content"`
	ToId    int    `json:"to_id"`
	FormId  int    `json:"form_id"`
	Time    int    `json:"time"`
}

// 登入返回
type LoginReturn struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 注册返回
type GetOnlineListReturn struct {
	UserList []UserInfo `json:"user_list"`
}
