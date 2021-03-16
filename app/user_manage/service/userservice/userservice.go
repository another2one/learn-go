package userservice

import (
	"learn-go/app/user_manage/model/user"
)

type UserService struct {
	UserSlice []user.User
	UserNum   int // 累计用户数量
}

func NewUserService() *UserService {
	userService := &UserService{}
	userService.UserNum = 1
	return userService
}

// 获取用户列表
func (userService *UserService) GetList() []user.User {
	return userService.UserSlice
}

// 增加用户
func (userService *UserService) Add(user user.User) {
	user.Id = userService.UserNum
	userService.UserNum++
	userService.UserSlice = append(userService.UserSlice, user)

}

// 删除用户
func (userService *UserService) Delete(index int) {
	userService.UserSlice = append(userService.UserSlice[:index], userService.UserSlice[index+1:]...)
}

// 根据id查找用户index
func (userService *UserService) Search(id int) (serachIndex int) {
	serachIndex = -1
	for index, value := range userService.UserSlice {
		if value.Id == id {
			serachIndex = index
		}
	}
	return serachIndex
}
