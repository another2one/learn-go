package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// go run main.go -name sss

	var name string
	//name = *flag.String("name", "lizhi111", "user name")
	flag.StringVar(&name, "name", "lizhi", "user name")

	// go run main.go -help
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println("hello ", name)
}
