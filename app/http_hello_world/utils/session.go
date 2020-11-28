package utils

// 以文件和数据库方式可以防止丢失

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// SessionType struct
type SessionType struct {
	ExpireTime int
	ClearTime  int
	session    map[string]map[string]string
}

// Session 初始化session
var Session *SessionType

func init() {
	temp := make(map[string]map[string]string)
	Session = &SessionType{
		session:    temp,
		ExpireTime: 900,
		ClearTime:  10,
	}
}

// GetSessionID 获取sessionID
func (s *SessionType) GetSessionID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// Get 获取session
func (s *SessionType) Get(sid, key string) (res string, err error) {
	res, ok := s.session[sid][key]
	if !ok {
		return "", errors.New("not found")
	}
	return res, nil
}

// Set 设置 session
func (s *SessionType) Set(sid, key, value string) {
	rand.Seed(time.Now().Unix())
	s.session[sid][key] = value
	s.session[sid]["time"] = strconv.Itoa(int(time.Now().Unix()) + s.ExpireTime)
	if rand.Intn(s.ClearTime) == 0 {
		s.Clear()
	}
}

// Clear 清除 session
func (s *SessionType) Clear() {
	for i, _ := range s.session {
		expire, _ := strconv.Atoi(s.session[i]["time"])
		if expire < int(time.Now().Unix()) {
			delete(s.session, i)
		}
	}
}
