package user

type User struct {
	Id int
	Name string
	Sex string
	Age int
	Telephone string
	Email string
}

func NewUser() User {
	return User{}
}