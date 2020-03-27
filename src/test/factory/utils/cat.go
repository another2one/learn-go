package utils

import "fmt"

type cat struct{}

func (c *cat) Do() {
	fmt.Println("喵喵喵 ...")
}

func init() {
	Register("cat", func() Class {
		return new(cat)
	})
}
