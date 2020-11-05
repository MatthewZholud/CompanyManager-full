package kafka

//todo: mzh: strange name. propose a better one
// interface introduces interaction with message broker, what if we change to rabbitmq?
type KafkaRep interface {
	send
	get
}

//todo: mzh: why used this approach of interface composition?
// why not to define two methods in KafkaRep interface?
type send interface {
	KafkaSend(str []byte, topic string) ([]byte, error)
}

type get interface {
	KafkaGet(topic string, byteUUID []byte) ([]byte, error)
}
