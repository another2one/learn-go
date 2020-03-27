package hash

import (
	"errors"
	"fmt"
)

// 关键字索引实现
// 注意:
// 1. 数据量大，更新查询频繁，而且不是很重要的数据
// 2. 对比 map 及 hashtable 速度 1000w

var (
	ErrorNotFound = errors.New("not found any user")
)

type HTable struct {
	MaxSize int
	Arr     []*User
}

type User struct {
	Id     int
	Status byte
	Next   *User
}

func GetHTable(maxSize int) *HTable {
	ht := &HTable{
		MaxSize: maxSize,
		Arr:     make([]*User, maxSize),
	}
	for i, _ := range ht.Arr {
		ht.Arr[i] = &User{}
	}
	return ht
}

func (ht *HTable) getId(id int) int {
	return id % ht.MaxSize
}

func (ht *HTable) Add(user *User) {
	ht.Arr[ht.getId(user.Id)].Add(user)
}

func (head *User) Add(user *User) {
	temp := head
	for {
		if temp.Next == nil {
			temp.Next = user
			break
		}
		if temp.Next.Id > user.Id {
			user.Next = temp.Next
			temp.Next = user
			break
		}
		temp = temp.Next
	}
}

func (ht *HTable) Delete(id int) {
	ht.Arr[ht.getId(id)].Delete(id)
}

func (head *User) Delete(id int) {
	temp := head
	for {
		if temp.Next == nil {
			return
		}
		if temp.Next.Id == id {
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
	}
}

func (ht *HTable) Search(id int) (user *User, err error) {
	return ht.Arr[ht.getId(id)].Search(id)
}

func (head *User) Search(id int) (user *User, err error) {
	temp := head
	for {
		if temp.Next == nil {
			return nil, ErrorNotFound
		}
		if temp.Next.Id == id {
			return temp.Next, nil
		}
		temp = temp.Next
	}
}

func (ht *HTable) Update(id int, status byte) (err error) {
	return ht.Arr[ht.getId(id)].Update(id, status)
}

func (head *User) Update(id int, status byte) (err error) {
	temp := head
	for {
		if temp.Next == nil {
			return ErrorNotFound
		}
		if temp.Next.Id == id {
			temp.Next.Status = status
			return
		}
		temp = temp.Next
	}
}

func (ht *HTable) Show() {
	for i, val := range ht.Arr {
		fmt.Printf("第 %d 散列内容: \n", i)
		val.Show()
	}
}

func (head *User) Show() {
	temp := head
	for {
		if temp.Next == nil {
			return
		}
		fmt.Printf("%d: %d \n", temp.Next.Id, temp.Next.Status)
		temp = temp.Next
	}
}
