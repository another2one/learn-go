package main

import (
	"fmt"
)

type CalcStack struct {
	maxSize, head int
	data          []int
}

func (as *CalcStack) push(i int) error {
	if as.head == as.maxSize-1 {
		return ErrorStackFull
	} else {
		as.head++
		as.data[as.head] = i
		return nil
	}
}

func (as *CalcStack) getHead() int {
	if as.head == -1 {
		return 0
	}
	return as.data[as.head]
}

func (as *CalcStack) pop() int {
	if as.head == -1 {
		return 0
	} else {
		i := as.data[as.head]
		as.head--
		return i
	}
}

func (as *CalcStack) show() {
	fmt.Println("栈数据：")
	i := as.head
	for i > -1 {
		fmt.Printf("%d:%d\n", i, as.data[i])
		i--
	}
	fmt.Println()
}

func main() {
	numStack := CalcStack{
		maxSize: 20,
		head:    -1,
		data:    make([]int, 8),
	}
	symbolStack := CalcStack{
		maxSize: 20,
		head:    -1,
		data:    make([]int, 8),
	}
	str := []byte("3+2*6*1-2+9*3-2/1")
	for _, v := range str {
		if isNumber(v) {
			numStack.push(toNumber(v))
		} else {
			for symbolStack.head >= 0 && !checkAdvance(int(v), symbolStack.getHead()) {
				numStack.push(calc(symbolStack.pop(), numStack.pop(), numStack.pop()))
			}
			symbolStack.push(int(v))
		}
	}
	for symbolStack.getHead() != 0 {
		numStack.push(calc(symbolStack.pop(), numStack.pop(), numStack.pop()))
	}
	res := numStack.pop()
	if res == 3+2*6*1-2+9*3-2/1 {
		fmt.Println("结果正确")
	} else {
		fmt.Printf("结果错误, 应该是 %d, 实际为 %d \n", res, 3+2*6*1-2+9*3-2/1)
	}
}

func calc(symbol, num1, num2 int) int {
	switch symbol {
	case '+':
		return num2 + num1
	case '-':
		return num2 - num1
	case '*':
		return num2 * num1
	case '/':
		return num2 / num1
	}
	return 0
}

func isNumber(v byte) bool {
	if v <= '9' && v >= '0' {
		return true
	} else {
		return false
	}
}

func checkAdvance(now, before int) bool {
	if before == 0 {
		return true
	}
	if (before == '*' || before == '/') && (now == '+' || now == '-') {
		return false
	} else if before == '-' && now == '+' {
		return false
	} else {
		return true
	}
}

func toNumber(v byte) int {
	return int(v - 48)
}
