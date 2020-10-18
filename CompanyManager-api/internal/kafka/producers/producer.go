package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		BatchTimeout: 1 * time.Millisecond,
	})
}

func KafkaSend(str []byte, topic string) {
	writer := getKafkaWriter("kafka:9092", topic)
	defer writer.Close()

	writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: str })
}


