package main

import (
	"fmt"
	"learn-go/app/elk/utils"
	"time"
)

func main() {

	client, err := utils.KafkaProducerInit()
	if err != nil {
		fmt.Println("kafka init fail: ", err)
		return
	}

	tails, err := utils.TailInit("../log/my.log")
	if err != nil {
		fmt.Println("tail file error: ", err)
		return
	}

	for {
		msg, ok := <-tails.Lines
		if !ok {
			fmt.Println("tail lines fail:")
			time.Sleep(time.Second)
			continue
		}
		utils.KafkaSendMessage(client, msg.Text, "test")
	}
}
