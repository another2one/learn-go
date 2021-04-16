package main

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	maxSize, head int
	data          []interface{}
}

func NewArrayStack(maxSize int) *ArrayStack {
	return &ArrayStack{
		maxSize: maxSize,
		head:    -1,
		data:    make([]interface{}, maxSize),
	}
}

var (
	ErrorStackEmpty = errors.New("stack empty")
	ErrorStackFull  = errors.New("stack full")
)

func (as *ArrayStack) push(i int) error {
	fmt.Println("push: ", i)
	if as.head == as.maxSize-1 {
		return ErrorStackFull
	} else {
		as.head++
		as.data[as.head] = i
		as.show()
		return nil
	}
}

func (as *ArrayStack) pop() (i int, err error) {
	fmt.Println("pop")
	if as.head == -1 {
		return 0, ErrorStackEmpty
	} else {
		i = as.data[as.head]
		as.head--
		as.show()
		return
	}
}

func (as *ArrayStack) show() {
	fmt.Println("栈数据：")
	i := as.head
	for i > -1 {
		fmt.Printf("%d:%d\n", i, as.data[i])
		i--
	}
	fmt.Println()
}
