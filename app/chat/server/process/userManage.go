package process

import "learn-go/app/chat/conf"

var (
	userManage *UserManage
)

// redis 不好做， 因为每个用户还要一个连接
// 定义全局变量

type UserManage struct {
	onlineList map[int]*UserProcess
}

func init() {
	userManage = &UserManage{onlineList: make(map[int]*UserProcess)}
}

// 加入在线列表
func (userManage *UserManage) AddOnline(userProcess *UserProcess) (err error) {
	userManage.onlineList[userProcess.Id] = userProcess
	return
}

// 移除在线列表
func (userManage *UserManage) RemoveOnline(i int) (err error) {
	delete(userManage.onlineList, i)
	return
}

// 查询是否在线
func (userManage *UserManage) getUserProcess(i int) (*UserProcess, error) {
	_, ok := userManage.onlineList[i]
	if ok {
		return userManage.onlineList[i], nil
	} else {
		return nil, conf.ErrorUserNotOnline
	}
}

// 返回在线列表
func (userManage *UserManage) GetOnlineList() (map[int]*UserProcess, error) {
	return userManage.onlineList, nil
}
