package hash

import (
	"math/rand"
	"testing"
	"time"
)

var (
	num        = 10000
	ht         = GetHTable(num / 100)
	seedNumber = time.Now().Unix()
)

func TestAdd(t *testing.T) {
	for i := 0; i < num; i++ {
		user := &User{
			Id:     i,
			Status: byte(i % 2),
			Next:   nil,
		}
		ht.Add(user)
	}
	if ht.Len == num {
		t.Logf("add succeed \n")
	} else {
		t.Errorf("add error \n")
	}
}

func TestDelete(t *testing.T) {

	rand.Seed(seedNumber - 1)
	deleteId := 0
	for i := 0; i < 10; i++ {
		deleteId = rand.Intn(num)
		t.Logf("delete %d \n", deleteId)
		ht.Delete(deleteId)
		user, err := ht.Search(deleteId)
		if user == nil {
			user = &User{
				Id:     deleteId,
				Status: byte(deleteId % 2),
				Next:   nil,
			}
			ht.Add(user)
		} else {
			t.Errorf("delete %d error: %v \n", deleteId, err)
		}
	}
	if ht.Len == num {
		t.Logf("delete succeed \n")
	} else {
		t.Errorf("delete error \n")
	}
}

func TestUpdate(t *testing.T) {

	rand.Seed(seedNumber - 2)
	updateId := 0
	status := 0
	for i := 0; i < 10; i++ {
		status = rand.Intn(4) + 2
		updateId = rand.Intn(num)
		err := ht.Update(updateId, byte(status))
		if err != nil {
			t.Errorf("update %d error: %v \n", updateId, err)
		} else {
			t.Logf("update %d status to %d success \n", updateId, status)
		}
	}
}

func TestSearch(t *testing.T) {

	rand.Seed(seedNumber - 3)
	searchId := 0
	for i := 0; i < 10; i++ {
		searchId = rand.Intn(num)
		user, err := ht.Search(searchId)
		if err != nil {
			t.Errorf("search %d error: %v \n", searchId, err)
		} else {
			t.Logf("search %d correct: %v  \n", searchId, user)
		}
	}
	//ht.Show()
}
