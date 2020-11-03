package interService

type interService struct {
	kafka KafkaRep
}

func Initialize(kafka KafkaRep) *interService {
	return &interService{
		kafka: kafka,
	}
}