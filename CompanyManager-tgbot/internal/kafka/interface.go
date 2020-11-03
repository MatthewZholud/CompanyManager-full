package kafka

type KafkaRep interface {
	send
	get
}

type send interface {
	KafkaSend(str []byte, topic string) ([]byte, error)
}

type get interface {
	KafkaGet(topic string, byteUUID []byte) ([]byte, error)
}
