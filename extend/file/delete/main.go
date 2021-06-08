package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	tmpDir := `C:/Users/Administrator/AppData/Local/Temp/`
	parttern := `^(phpcs|staticcheck|go-build).*`
	reg := regexp.MustCompile(parttern)

	file, err := os.Open(tmpDir)
	if err != nil {
		log.Fatalf("open temp dir error: %s \n", err)
		return
	}
	files, _ := file.ReadDir(-1)
	for _, v := range files {
		if reg.MatchString(v.Name()) {
			if err := os.RemoveAll(tmpDir + v.Name()); err != nil {
				log.Fatalf("del dir %s error: %s \n", v.Name(), err)
			} else {
				fmt.Println("delete ", v.Name())
			}
			// fmt.Println(v.Name())
		}
	}

	fmt.Println("press ctrl+c to exit")
	s := <-c
	fmt.Println("Got signal:", s)
}
