package consumers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		StartOffset: kafka.LastOffset,
		MaxWait: 1 * time.Millisecond,

	})
}

func KafkaGetStruct(topic string) []byte {
	reader := getKafkaReader("kafka:9092", topic)

	reader.SetOffset(kafka.LastOffset)

	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		return m.Value
	}
	return nil
}
