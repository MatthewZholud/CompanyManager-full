package consumers

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func ExampleConsumerGroupParallelReaders() {
	group, err := kafka.NewConsumerGroup(kafka.ConsumerGroupConfig{
		ID:      "my-group",
		Brokers: []string{"kafka:9092"},
		Topics:  []string{"getCompany", "getEmployee"},
	})
	if err != nil {
		fmt.Printf("error creating consumer group: %+v\n", err)
		os.Exit(1)
	}
	defer group.Close()
	ctx := context.Background()


	for {
		gen, err := group.Next(ctx)
		fmt.Println(gen)
		if err != nil {
			break
		}
		fmt.Println()
			gen.Start(func(ctx context.Context) {
				reader := kafka.NewReader(kafka.ReaderConfig{
					Brokers:   []string{"kafka:9092"},
					Topic:     "getEmployee",
					Partition: 0,
				})
				reader.SetOffset(kafka.LastOffset)
				defer reader.Close()


				reader1 := kafka.NewReader(kafka.ReaderConfig{
					Brokers:   []string{"kafka:9092"},
					Topic:     "getCompany",
					Partition: 0,
				})
				reader1.SetOffset(kafka.LastOffset)
				defer reader1.Close()

				for {
					msg, err := reader.ReadMessage(ctx)
					switch err {
					case kafka.ErrGenerationEnded:
						return
					case nil:
						fmt.Printf("received message %s/%d/%d : %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
					default:
						fmt.Printf("error reading message: %+v\n", err)
					}
					break

				}
			})

	}
}