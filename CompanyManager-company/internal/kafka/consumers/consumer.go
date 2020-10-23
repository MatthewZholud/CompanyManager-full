package consumers

import (
	"context"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		StartOffset: kafka.LastOffset,
		MaxWait:     10 * time.Millisecond,
	})
}

func KafkaConsumer(topic, brokers string, ch chan []byte) []byte {

	reader := getKafkaReader(brokers, topic)
	reader.SetOffset(kafka.LastOffset)
	defer reader.Close()

	logger.Log.Info("start consuming", topic, "... !!")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logger.Log.Fatalf("Can't read messages from Kafka with topic %v: %v", topic, err)
		} else {
			logger.Log.Infof("Got message from kafka, topic: %v", topic)
		}
		ch <- m.Value
	}
	return nil
}
