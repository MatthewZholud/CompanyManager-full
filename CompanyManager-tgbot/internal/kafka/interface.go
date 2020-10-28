package kafka

type Exchanger interface {
	KafkaSend(str []byte, topic string) ([]byte, error)
	KafkaGet(topic string, byteUUID []byte) ([]byte, error)
}
