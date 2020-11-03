package kafka

import (
	"context"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"os"

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

func (k *kafkaClient) KafkaGet(topic string, byteUUID []byte) ([]byte, error) {
	reader := getKafkaReader(os.Getenv("KAFKA_BROKERS"), topic)

	reader.SetOffset(kafka.LastOffset)

	logger.Log.Info("Start consuming", topic, "... !!")

	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logger.Log.Debugf("Error receiving message from kafka, topic: %s", topic)
			return nil, err
		} else {
			logger.Log.Infof("Got message from kafka, topic: %s", topic)
		}
		if string(byteUUID) != string(msg.Key) {
			logger.Log.Errorf("The key of the sent message does not match the key of the received one, please try again")
			return nil, nil
		} else {
			logger.Log.Info("Received message with correct key")
		}
		return msg.Value, nil
	}
}
