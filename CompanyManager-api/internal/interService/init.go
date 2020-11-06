package interService

type interService struct {
	kafka BrokerRep
}

func Initialize(kafka BrokerRep) *interService {
	return &interService{
		kafka: kafka,
	}
}