// https://go.dev/play/p/bPTXaPHm6jD
package main

import (
	"fmt"
)

type order struct {
	order_id string
}

func main() {
	order1 := order{order_id: "123"}
	fmt.Printf("res6: %v \n", map[string]map[string]string{"order": {"order_id": "123"}})
	fmt.Printf("res6: %+v \n", map[string]map[string]string{"order": {"order_id": "123"}})
	fmt.Printf("res6: %v \n", order1)
	fmt.Printf("res6: %+v \n", order1)
	fmt.Printf("res6: %#v \n", order1)
}
