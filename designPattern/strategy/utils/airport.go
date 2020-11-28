package utils

import "fmt"

type Airport struct{}

func (airport *Airport) Travel(destination string) {
	fmt.Println("airport travel to ", destination)
}
