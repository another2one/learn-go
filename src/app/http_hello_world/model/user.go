package model

import (
	"app/http_hello_world/utils"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

func (u *User) Add() error {
	res, err := utils.Db.Exec("insert into user (name, create_time) values (?, ?)", u.Name, time.Now().Unix())
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	u.CreateTime = int(time.Now().Unix())
	return err
}

func (u *User) Update() error {
	_, err := utils.Db.Exec("update user set name = ?, update_time = ? where id = ?", u.Name, time.Now().Unix(), u.Id)
	return err
}

func (u *User) Delete(id int) error {
	_, err := utils.Db.Exec("delete from user where id = ?", id)
	return err
}

func (u *User) GetById(id int) (user *User, err error) {
	user = &User{}
	err = utils.Db.QueryRow("select id, name, create_time, update_time from user where id = ?", id).Scan(&user.Id, &user.Name, &user.CreateTime, &user.UpdateTime)
	return
}

func (u *User) GetByName(name string) (users []*User, err error) {
	users = make([]*User, 0)
	rows, err := utils.Db.Query("select * from user where name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &User{}
		if err1 := rows.Scan(&user.Id, &user.Name, &user.CreateTime, &user.UpdateTime); err1 == nil {
			users = append(users, user)
		}
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
