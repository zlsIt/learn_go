package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	var topic = "test_msg"
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("create consumer failed. err:", err)
		return
	}
	defer consumer.Close()

	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition.AsyncClose()

	for {
		select {
		case msg := <-partition.Messages():
			fmt.Printf("Partition:%d offset:%d key:%v, value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		case err := <-partition.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

	// 根据topic获取所有partitions
	//partitions, err := consumer.Partitions(topic)
	//if err != nil {
	//	fmt.Println("get topic partitions failed. err:", err)
	//	return
	//}
	//fmt.Println(partitions)
	//done := make(chan bool)
	//for _, partition := range partitions {
	//	pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	//	if err != nil {
	//		fmt.Printf("failed to start consumer for partition %d, err%d\n", pc, err)
	//		return
	//	}
	//	defer pc.AsyncClose()
	//	go func(sarama.PartitionConsumer) {
	//		for msg := range pc.Messages() {
	//			fmt.Printf("Partition:%d offset:%d key:%v, value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	//			done <- true
	//		}
	//	}(pc)
	//}
	//<-done
}
