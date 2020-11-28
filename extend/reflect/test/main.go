package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func test(i interface{}) {
}

func main() {

	var (
		s = "66"
		i = 6
		t = true
		slice = []int{66, 12, 16}
		m = map[int]int{1:2, 2:3}
		model = &Student{}
		elem reflect.Value
		st reflect.Type
	)

	Canset(&s)
	Canset(&i)
	Canset(&t)
	Canset(&slice)
	Canset(&m)

	fmt.Printf("model type: %T \n", model)
	st = reflect.TypeOf(model) // ptr
	fmt.Println("reflect.TypeOf(model):", st.String())
	st = st.Elem() // struct
	fmt.Println("reflect.TypeOf(model).Elem():", st.String())
	// 根据结构体类型创建一个结构体指针的反射值
	elem = reflect.New(st) // ptr
	fmt.Println("reflect.New(st)", elem.String())
	fmt.Println("reflect.New(st).Elem()", elem.Elem().String())
	// 取到实际的结构体指针
	model1 := elem.Interface().(*Student)
	elem = elem.Elem()
	model.Name = "lizhi"
	fmt.Println("model.Name = lizhi")
	fmt.Println("model: ", model)
	fmt.Println("model1: ", model1)
	fmt.Println("elem:", elem)
	fmt.Println()
	elem.FieldByName("Name").SetString("lipan")
	fmt.Println("elem.FieldByName.SetString = lipan")
	fmt.Println("model: ", model)
	fmt.Println("model1: ", model1)
	fmt.Println("elem:", elem)
	fmt.Println()
	model1.Name = "lizhu"
	fmt.Println("model1.Name = lizhu")
	fmt.Println("model: ", model)
	fmt.Println("model1: ", model1)
	fmt.Println("elem:", elem)
	// 取到结构体实例的值并复制给model2
	model2 := elem.Interface().(Student)
	fmt.Println("model2: ", model2)
}

func Canset(i interface{}) {
	s := reflect.ValueOf(i).Elem()
	if s.CanSet() {
		fmt.Printf("%v can set \n", s.Kind())
	}else{
		fmt.Printf("%v can't set \n", s.Kind())
	}
}