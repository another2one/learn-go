package main

import "fmt"

type Person struct {
	name string
}

func main() {
	persons := []Person{{"zhangsan"}, {"lisi"}}
	newPersons := []*Person{}

	// for range 其实会生成 k person 临时变量去接收值
	for k, person := range persons {
		fmt.Printf("add=%p \n", &person)
		fmt.Printf("add=%p \n", &k)
		fmt.Println("name=" + person.name)
		newPersons = append(newPersons, &person)
	}

	// 正确写法
	// for _, person := range persons {
	// 	p1 := person
	// 	newPersons = append(newPersons, &p1)
	// }

	fmt.Println(newPersons[0].name)
	fmt.Println(newPersons[1].name)

}
