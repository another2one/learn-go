package utils

// 以文件和数据库方式可以防止丢失

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
)

type SessionType struct {
	ExpireTime int
	ClearTime  int
	session    map[string]map[string]string
}

var Session *SessionType

func init() {
	temp := make(map[string]map[string]string)
	Session = &SessionType{
		session:    temp,
		ExpireTime: 900,
		ClearTime:  10,
	}
}

func (s *SessionType) GetSessionId() string {
	return uuid.Must(uuid.NewV4()).String()
}

func (s *SessionType) Get(sid, key string) (res string, err error) {
	res, ok := s.session[sid][key]
	if !ok {
		return "", errors.New("not found")
	}
	return res, nil
}

func (s *SessionType) Set(sid, key, value string) {
	rand.Seed(time.Now().Unix())
	s.session[sid][key] = value
	s.session[sid]["time"] = strconv.Itoa(int(time.Now().Unix()) + s.ExpireTime)
	if rand.Intn(s.ClearTime) == 0 {
		s.Clear()
	}
}

func (s *SessionType) Clear() {
	for i, _ := range s.session {
		expire, _ := strconv.Atoi(s.session[i]["time"])
		if expire < int(time.Now().Unix()) {
			delete(s.session, i)
		}
	}
}
