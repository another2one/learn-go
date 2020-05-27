package utils

import "fmt"

type Bus struct{}

func (bus *Bus) Travel(destination string) {
	fmt.Println("bus travel to ", destination)
}
