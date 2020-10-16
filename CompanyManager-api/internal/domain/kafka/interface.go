package kafka

type Exchanger interface {
	KafkaSendId(id, topic string, partition int)
	KafkaGetStruct(topic string) []byte
}

type KafkaExchange struct {
	Exchange Exchanger
}