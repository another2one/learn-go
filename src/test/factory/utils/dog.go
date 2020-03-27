package utils

import "fmt"

type dog struct{}

func (c *dog) Do() {
	fmt.Println("汪汪汪 ...")
}

func init() {
	Register("dog", func() Class {
		return new(dog)
	})
}
