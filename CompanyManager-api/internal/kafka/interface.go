package kafka

type Exchanger interface {
	KafkaSend(str []byte, topic string)
	KafkaGetStruct(topic string) []byte
}