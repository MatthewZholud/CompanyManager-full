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
	//todo: mzh: Why use deprecated method?
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		//todo: mzh: hardcode of timeout
		BatchTimeout: 10 * time.Millisecond,
	})
}

func (k *kafkaClient) KafkaSend(str []byte, topic string) ([]byte, error) {
	writer := getKafkaWriter(os.Getenv("KAFKA_BROKERS"), topic)
	defer writer.Close()
	//todo: mzh: would be nice to include what message is being sent
	logger.Log.Infof("Ready to send message to kafka")
	currentUUID := uuid.New()
	//todo: mzh: UUID is already is a byte array, why convert if to string and back to byte array??
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
