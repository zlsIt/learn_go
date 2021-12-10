package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"time"
)

func main() {
	config := sarama.NewConfig()
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true
	// 构建一个消息
	//msg := &sarama.ProducerMessage{}
	var testMsgTopic = "test_msg"
	//msg.Topic = testMsgTopic
	//msg.Value = sarama.StringEncoder("this is a test msg")

	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("test_msg topic create producer err :%s\n", err)
		return
	}
	defer producer.AsyncClose()

	msg := &sarama.ProducerMessage{
		Topic: testMsgTopic,
		Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message"
	index := 0
	for {
		index++
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		msg.Value = sarama.ByteEncoder(fmt.Sprintf("%s%d", value, index))
		producer.Input() <- msg
		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %+v\n", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}

	// 连接kafka
	//client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	//if err != nil {
	//	fmt.Println("product closed, err:", err)
	//	return
	//}
	//defer client.Close()
	//// 发送消息
	//pid, offset, err := client.SendMessage(msg)
	//if err != nil {
	//	fmt.Println("send msg failed. err:", err)
	//	return
	//}
	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
