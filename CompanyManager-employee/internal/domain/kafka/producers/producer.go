package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/entity/employee"
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

func SendFromApiEmployee(employee *employee.Employee){
	// get kafka reader using environment variables.
	e, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	topic := "sendEmployee"

	writer := getKafkaWriter("kafka:9092", topic)
	defer writer.Close()

	writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(e) })

}