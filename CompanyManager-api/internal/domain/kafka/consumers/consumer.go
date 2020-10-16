package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		StartOffset: kafka.LastOffset,
	})
}

func KafkaGetStruct(topic string) presenter.Employee {

	reader := getKafkaReader("kafka:9092", topic)
	reader.SetOffset(kafka.LastOffset)
	//defer reader.Close()

	fmt.Println("start consuming ... !!")

	employee := presenter.Employee{}

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal([]byte(m.Value), &employee)
		return employee
	}
	return employee
}
