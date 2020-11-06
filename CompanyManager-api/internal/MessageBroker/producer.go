package MessageBroker

import (
	"context"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"os"
)


const tcp = "tcp"

func (k *kafkaClient) BrokerSend(str []byte, topic string) ([]byte, error) {
	partition := getEnvAsInt("PARTITION", 0)
	kafkaURL := os.Getenv("KAFKA_BROKERS")
	conn, err := kafka.DialLeader(context.Background(), tcp, kafkaURL, topic, partition)
	if err != nil {
		logger.Log.Debugf("Can't deal with MessageBroker Leader: %v", err)
		return nil, err
	}

	logger.Log.Infof(`Ready to send message "%v" to MessageBroker`, string(str))
	currentUUID := uuid.New()
	byteUUID, err := currentUUID.MarshalBinary()
	if err != nil {
		logger.Log.Debug("Error bringing UUID to []byte: %v", err)
		return nil, err
	}
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   byteUUID,
			Value: str,
		})
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker, topic: %v", topic)
		return nil, err
	} else {
		logger.Log.Infof("Sent message to MessageBroker, topic: %v", topic)
		return byteUUID, nil
	}
}
