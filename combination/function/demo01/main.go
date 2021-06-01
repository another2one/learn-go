package main

import "fmt"

func double(a *int) {
	*a += *a
	a = nil
	fmt.Printf("func-p address %p\n", &a)
}

func main() {
	a := 1
	p := &a
	double(p)
	fmt.Println(a)
	fmt.Printf("main-p address %p\n", &a)
	p = nil
	fmt.Println(a)
}
