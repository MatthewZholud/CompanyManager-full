package consumers

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		StartOffset: kafka.LastOffset,
		MaxWait: 1 * time.Millisecond,
	})
}

func GetFromApiCompany(topic string, ch chan string) string {

	reader := getKafkaReader("kafka:9092", topic)
	reader.SetOffset(kafka.LastOffset)
	defer reader.Close()

	fmt.Println("start consuming", topic ,"... !!")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

			ch <- string(m.Value)
	}
	return ""
}
