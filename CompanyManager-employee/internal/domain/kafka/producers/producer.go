package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/entity/employee"

)

func SendFromApiEmployee(employee *employee.Employee){
	// get kafka reader using environment variables.
	e, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}

	topic := "sendEmployee"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(e)},

	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	//if err := conn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}
}