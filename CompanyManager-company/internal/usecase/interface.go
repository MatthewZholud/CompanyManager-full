package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
)

type Reader interface {
	Get(id int64) (*entity.Company, error)
	GetAll() (*[]entity.Company, error)
}

type Writer interface {
	Create(e *entity.Company) (string, error)
	Update(e *entity.Company) (string, error)
	Delete(id int64) (string, error)
}

//repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetCompany(message []byte) ([]byte, error)
	GetAllCompany() ([]byte, error)


	CreateCompany(message []byte) ([]byte, error)
	UpdateCompany(message []byte) ([]byte, error)
	DeleteCompany(message []byte) ([]byte, error)
}



type KafkaRep interface {
	send
	get
}

type send interface {
	KafkaConsumer(topic, brokers string, ch chan entity.Message) []byte
}

type get interface {
	KafkaSend(str, Key []byte, topic string)
}


