package model

type User struct {
	BaseModel
	Name string `json:"name" gorm:"size:32"`
}

func (User) TableName() string {
	return "user"
}
