package utils

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"time"
)

func KafkaProducerInit() (sarama.SyncProducer, error) {

	//初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//生产者
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return client, err
	}
	defer client.Close()

	return client, nil

}

func KafkaSendMessage(client sarama.SyncProducer, msgStr, topic string) {
	//创建消息
	msg := &sarama.ProducerMessage{Timestamp: time.Now().Add(time.Second * 3)}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(msgStr)
	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n,", pid, offset)
}

func TailInit(fileName strig) (*Tail, error) {

	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	return tail.TailFile(fileName, config)
}
