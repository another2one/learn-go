package model

import (
	"github.com/gomodule/redigo/redis"
	"learn-go/app/chat/conf"
	"learn-go/app/chat/utils"
	"strconv"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (userDao *UserDao) GetUserById(i int) (user *User, err error) {

	conn := utils.GetConn()
	defer conn.Close()
	key := "user" + strconv.Itoa(i)
	res, err := redis.Strings(conn.Do("hmget", key, "name", "password"))
	// TODO: err是nil只能说明命令执行了， 必须数据存在才能说明有数据
	if err != nil {
		err = conf.ErrorRedisGetData
		return
	}
	if res[0] == "" {
		err = conf.ErrorUserNotExists
		return
	}
	user = &User{}
	user.Id = i
	user.Name = res[0]
	user.SetPassWord(res[1])
	return
}

func (userDao *UserDao) GetUserListByIds(list []int) (users []*User, err error) {
	users = make([]*User, 0, len(list))
	for _, i := range list {
		user, err := userDao.GetUserById(i)
		if err == nil {
			users = append(users, user)
		}
	}
	return
}

func (userDao *UserDao) AddUser(registerData conf.Register) (user *User, err error) {

	conn := utils.GetConn()
	defer conn.Close()
	id, err := redis.Int(conn.Do("incr", "userlen"))
	if err != nil {
		err = conf.ErrorRedisGetData
	}
	key := "user" + strconv.Itoa(id)
	_, err = conn.Do("hmset", key, "name", registerData.Name, "password", registerData.PassWord)
	if err != nil {
		err = conf.ErrorRedisGetData
		return
	}
	user = &User{}
	user.Id = id
	user.Name = registerData.Name
	user.SetPassWord(registerData.PassWord)
	return
}
