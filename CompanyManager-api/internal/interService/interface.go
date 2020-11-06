package interService

type BrokerRep interface {
	BrokerSend(str []byte, topic string) ([]byte, error)
	BrokerGet(topic string, byteUUID []byte) ([]byte, error)
}