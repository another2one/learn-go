package linkedList

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Number int
	Now    int
	Next   *Person
}

func NewCircleLink(n int) *Person {

	person1 := &Person{
		Number: 1,
		Next:   nil,
	}

	temp := person1

	if n < 2 {
		person1.Next = person1
		return person1
	}
	for i := 2; i < n+1; i++ {
		person := &Person{
			Number: i,
			Next:   nil,
		}
		temp.Next = person
		temp = temp.Next
	}
	temp.Next = person1
	return person1
}

func (p *Person) Show() {
	temp := p.Next
	for {
		fmt.Println(temp)
		if temp == p {
			return
		}
		temp = temp.Next
	}
}

func (p *Person) Search(k int) *Person {
	if p.Number == k {
		return p
	}
	temp := p.Next
	for {
		if temp == p {
			return nil
		}
		if temp.Number == k {
			return temp
		}
		temp = temp.Next
	}
}

func (p *Person) Report(start, m int) {

	temp := p

	for {

		fmt.Printf("%d 号报数：%d \n", temp.Number, start)
		if start%m == 0 {
			temp.DeleteSelf()
			fmt.Printf("%d 号出列 \n", temp.Number)
			start = 0
		}
		if temp.Next == nil {
			fmt.Println("the last one: ", temp)
			return
		}
		start++
		temp = temp.Next
	}
}

func (p *Person) GetPre() *Person {

	if p.Next == nil {
		return nil
	}

	temp := p.Next

	for {
		if temp.Next == p {
			return temp
		}
		if temp == p {
			return nil
		}
		temp = temp.Next
	}
}

func (p *Person) DeleteSelf() {
	pre := p.GetPre()
	if pre == nil {
		// 只剩一个链表
		return
	} else if pre == p.Next {
		pre.Next = nil
	} else {
		pre.Next = p.Next
	}
}

func RunJosephu() {
	n := 10
	head := NewCircleLink(n) // 创建n个小朋友座一圈
	//p.Show() // 查看所有小朋友
	rand.Seed(time.Now().Unix())
	pk := head.Search(rand.Intn(n) + 1) // 1-n 随机取一个小朋友
	m := 5
	start := 1
	pk.Report(start, m) // 从 start 开始报数，到 m 出列
}
