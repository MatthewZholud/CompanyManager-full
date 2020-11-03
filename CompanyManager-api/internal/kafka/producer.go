package kafka

import (
	"context"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
	"time"
)

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	})
}

func (k *kafkaClient) KafkaSend(str []byte, topic string) ([]byte, error) {
	writer := getKafkaWriter(os.Getenv("KAFKA_BROKERS"), topic)
	defer writer.Close()
	logger.Log.Infof("Ready to send message to kafka")
	currentUUID := uuid.New()
	byteUUID := []byte(fmt.Sprintf("%s", currentUUID))
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   byteUUID,
			Value: str,
		})
	if err != nil {
		logger.Log.Debugf("Error sending message to kafka, topic: %v", topic)
		return nil, err
	} else {
		logger.Log.Infof("Sent message to kafka, topic: %v", topic)
		return byteUUID, nil
	}
}
