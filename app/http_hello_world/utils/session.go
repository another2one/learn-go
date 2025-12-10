package utils

// 以文件和数据库方式可以防止丢失

import (
	"encoding/json"
	"fmt"
	"learn-go/common/tool"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	uuid "github.com/satori/go.uuid"
)

// SessionType struct
type SessionType struct {
	ExpireSeconds int // 过期时间
	ClearTime     int
	session       map[string]map[string]any
	File          string // 存储文件
}

// Session 初始化session
var Session *SessionType

func InitSession() {
	temp := make(map[string]map[string]any)
	Session = &SessionType{
		ExpireSeconds: 900,
		session:       temp,
		ClearTime:     10,
		File:          tool.ProjectPath + "app/http_hello_world/http/session/session.json",
	}
	Session.Open()
}

func GetAll(sid string) map[string]any {
	s, ok := Session.session[sid]
	if !ok {
		return nil
	}
	return s
}

// GetSessionID 获取sessionID
func (s *SessionType) GetSessionID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// Get 获取session
func (s *SessionType) Get(sid, key string) (res any) {
	if sid == "" || key == "" {
		return ""
	}
	res, ok := s.session[sid][key]
	if !ok {
		return ""
	}
	return res
}

// Set 设置 session
func (s *SessionType) Set(sid, key string, value any) {
	_, ok := s.session[sid]
	if !ok {
		s.session[sid] = make(map[string]any)
	}
	s.session[sid][key] = value
	s.session[sid]["expire_at"] = time.Now().Add(time.Duration(s.ExpireSeconds) * time.Second)
	if rand.Intn(s.ClearTime) == 0 {
		s.Clear()
	}
}

// Clear 清除 session
func (s *SessionType) Clear() {
	for i, _ := range s.session {
		if time.Now().After(s.session[i]["expire_at"].(time.Time)) {
			fmt.Println("过期了 。。。")
			delete(s.session, i)
		}
	}
}

// Open 启动 session
func (s *SessionType) Open() {
	err := os.MkdirAll(filepath.Dir(s.File), 0755)
	if err != nil {
		panic("session file stat err: " + err.Error())
	}
	file, err := os.OpenFile(s.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("open session file err: " + err.Error())
	}
	defer file.Close()

	content, err := os.ReadFile(s.File)
	if err != nil {
		panic("read session file err: " + err.Error())
	}
	if len(content) > 0 {
		err := json.Unmarshal(content, &s.session)
		if err != nil {
			panic("unmarshal session file err: " + err.Error())
		}
	}
}

// Close 关闭 session
func (s *SessionType) Close() {
	content, err := json.Marshal(&s.session)
	if err != nil {
		panic("write session file err: " + err.Error())
	}
	err = os.WriteFile(s.File, content, 0666)
	if err != nil {
		panic("write session file err: " + err.Error())
	}
}
