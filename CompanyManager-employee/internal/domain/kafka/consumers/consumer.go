package consumers

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		StartOffset: kafka.LastOffset,
	})
}

func GetFromApiEmployee(topic string, ch chan string) string {
	fmt.Println("hi my topic is", topic)

	reader := getKafkaReader("kafka:9092", topic)
	reader.SetOffset(kafka.LastOffset)
	//defer reader.Close()

	fmt.Println("start consuming", topic ,"... !!")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if m.Value != nil{
			ch <- string(m.Value)
		}
	}
	return ""
}
