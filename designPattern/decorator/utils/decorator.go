package utils

import "fmt"

type DecoratorInter interface {
	Cache(line *Line)
}

type Decorator struct {
}

func (de *Decorator) Cache(line *Line) {
	if line.Draw() == "直线" {
		fmt.Println("cache")
	}

}
