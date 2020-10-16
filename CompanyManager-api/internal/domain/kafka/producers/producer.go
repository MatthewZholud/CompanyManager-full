package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func KafkaSendId(id, topic string, partition int) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(id)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	//if err := conn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}
}
