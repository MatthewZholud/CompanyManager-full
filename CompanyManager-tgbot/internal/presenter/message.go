package presenter


type Message struct {
	Key   []byte`json:"key"   valid:"-"`
	Value []byte `json:"value" valid:"-"`
}

type MessagePack struct {
	Topic   string  `json:"topic"   valid:"required"`
	Message Message `json:"message" valid:"required"`
}
