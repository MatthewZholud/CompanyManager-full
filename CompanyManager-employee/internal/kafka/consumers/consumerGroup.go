package consumers

//import (
//	"context"
//	"fmt"
//	"github.com/segmentio/env-go"
//	"os"
//)
//
//func ExampleConsumerGroupParallelReaders() {
//	group, err := env.NewConsumerGroup(env.ConsumerGroupConfig{
//		ID:      "my-group",
//		Brokers: []string{"env:9092"},
//		Topics:  []string{"getCompany", "getEmployee"},
//	})
//	if err != nil {
//		fmt.Printf("error creating consumer group: %+v\n", err)
//		os.Exit(1)
//	}
//	defer group.Close()
//	ctx := context.Background()
//
//
//	for {
//		gen, err := group.Next(ctx)
//		fmt.Println(gen)
//		if err != nil {
//			break
//		}
//		fmt.Println()
//			gen.Start(func(ctx context.Context) {
//				reader := env.NewReader(env.ReaderConfig{
//					Brokers:   []string{"env:9092"},
//					Topic:     "getEmployee",
//					Partition: 0,
//				})
//				reader.SetOffset(env.LastOffset)
//				defer reader.Close()
//
//
//				reader1 := env.NewReader(env.ReaderConfig{
//					Brokers:   []string{"env:9092"},
//					Topic:     "getCompany",
//					Partition: 0,
//				})
//				reader1.SetOffset(env.LastOffset)
//				defer reader1.Close()
//
//				for {
//					msg, err := reader.ReadMessage(ctx)
//					switch err {
//					case env.ErrGenerationEnded:
//						return
//					case nil:
//						fmt.Printf("received message %s/%d/%d : %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
//					default:
//						fmt.Printf("error reading message: %+v\n", err)
//					}
//					break
//
//				}
//			})
//
//	}
//}