package main

import "fmt"

type Human interface {
	Talk()
}

func SpeakTo(h Human) {
	h.Talk()
}

type Person struct{ Name string }
type Citizen struct {
	Person
	Country string
}

func (p *Person) Talk() {
	fmt.Println("my name is ", p.Name)
}

func (c *Citizen) Talk() {
	fmt.Println("my name is ", c.Name, " from ", c.Country)
}

// 多态实现
func main() {
	p := Person{Name: "Dave"}
	c := Citizen{Person: Person{Name: "Steve"}, Country: "America"}

	SpeakTo(&p)
	SpeakTo(&c)

	h := Human(&p)
	h.Talk()
	h = &c
	h.Talk()

}
