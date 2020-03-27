package model

import (
	"testing"
)

func TestUser(t *testing.T) {

	user := &User{}

	user.Name = "lizhi"
	err := user.Add()
	if err != nil {
		t.Error("add error: ", err)
	}

	users, err := user.GetByName("lizhi")
	if err != nil {
		t.Error("query error: ", err)
	}
	if len(users) == 0 {
		t.Error("no user has found")
	}
	t.Logf("user: %+v \n", users[0])

	users[0].Name = "lizhu"
	err = users[0].Update()
	if err != nil {
		t.Error("update error: ", err)
	}
	users, err = user.GetByName("lizhu")
	if err != nil {
		t.Error("query error: ", err)
	}
	if len(users) == 0 {
		t.Error("no user has found")
	}
	t.Logf("user: %+v \n", users[0])

	user.Delete(users[0].Id)
	users, _ = user.GetByName("lizhu")
	if len(users) == 0 {
		t.Log("delete ok")
	}

}
