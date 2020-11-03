package interService

import "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka"

type interService struct {
	kafka kafka.KafkaRep
}

func Initialize(kafka kafka.KafkaRep) *interService {
	return &interService{
		kafka: kafka,
	}
}