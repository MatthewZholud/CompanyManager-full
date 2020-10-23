package producers

import (
	"context"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
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
		BatchTimeout: 10 * time.Millisecond,
	})
}

func KafkaSend(str []byte, topic string) {
	writer := getKafkaWriter("kafka:9092", topic)
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: str })
	if err != nil {
		logger.Log.Fatalf("Can't read messages from Kafka with topic %v: %v", topic, err)
	} else {
		logger.Log.Infof("Sent message to kafka, topic: %v", topic)
	}
}