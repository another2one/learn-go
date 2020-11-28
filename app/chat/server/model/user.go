package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	passWord string
}

func (user *User) PassWord() string {
	return user.passWord
}

func (user *User) SetPassWord(passWord string) {
	user.passWord = passWord
}
