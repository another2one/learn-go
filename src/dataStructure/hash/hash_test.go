package hash

import (
	"math/rand"
	"testing"
)

var (
	num = 10000
	ht  = GetHTable(num / 100)
)

func TestHTable_Add(t *testing.T) {

	for i := 0; i < num; i++ {
		user := &User{
			Id:     i,
			Status: byte(i % 2),
			Next:   nil,
		}
		ht.Add(user)
	}
}

func TestHTable_Delete(t *testing.T) {

	rand.Seed(1458702589)
	deleteId := 0
	for i := 0; i < 10; i++ {
		deleteId = rand.Intn(num)
		t.Logf("delete %d \n", deleteId)
		ht.Delete(deleteId)
	}
}

func TestHTable_Update(t *testing.T) {

	rand.Seed(145870252412)
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

func TestHTable_Search(t *testing.T) {

	rand.Seed(51489523144)
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
