package consumers

import (
	"context"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		StartOffset: kafka.LastOffset,
		MaxWait: 10 * time.Millisecond,

	})
}

func KafkaGetStruct(topic string) ([]byte, error) {
	reader := getKafkaReader("kafka:9092", topic)

	reader.SetOffset(kafka.LastOffset)

	logger.Log.Info("Start consuming", topic, "... !!")


	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logger.Log.Debugf("Error receiving message from kafka, topic: %v", topic)
			return nil, err
		} else {
			logger.Log.Infof("Got message from kafka, topic: %v", topic)
		}
		return m.Value, nil
	}
}
