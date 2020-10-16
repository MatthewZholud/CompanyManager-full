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

func GetFromApiEmployee() string {
	// get kafka reader using environment variables.
	kafkaURL := "kafka:9092"
	topic := "getEmployee"

	reader := getKafkaReader(kafkaURL, topic)
	reader.SetOffset(kafka.LastOffset)
	//defer reader.Close()

	fmt.Println("start consuming getEmployee ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		fmt.Println(string(m.Value))
		return string(m.Value)
	}
}
