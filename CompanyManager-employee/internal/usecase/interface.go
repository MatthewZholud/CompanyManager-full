package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity"
)

type Reader interface {
	Get(id int64) (*entity.Employee, error)
	GetAll() (*[]entity.Employee, error)
	GetEmployeesByCompany(id int64) (*[]entity.Employee, error)
}

type Writer interface {
	Create(e *entity.Employee) (string, error)
	Update(e *entity.Employee) (string, error)
	Delete(id int64) (string, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetEmployee(message []byte) ([]byte, error)
	GetAllCompany() ([]byte, error)
	CreateEmployee(message []byte) ([]byte, error)
	UpdateEmployee(message []byte) ([]byte, error)
	DeleteEmployee(message []byte) ([]byte, error)
	GetEmployeeByCompany(message []byte) ([]byte, error)
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
