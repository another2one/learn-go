package utils

type dog struct {
	Name string
}

var singleDog *dog

func init() {
	singleDog = new(dog)
}

func NewDog() *dog {
	return singleDog
}
