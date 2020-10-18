package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/producers"
	"log"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetEmployee(message []byte) {
	id := StringToInt64(string(message))
	employee, err := s.repo.Get(id)
	if err != nil {
		log.Fatal(entity.ErrNotFound)
	}
	if err != nil {
		log.Fatal(err)
	}
	e, err := json.Marshal(employee)
	if err != nil {
		log.Fatal(err)
	}
	producers.KafkaSend(e, "EmployeeGETResponse")
}

func (s *Service) CreateEmployee(message []byte) {
	empl := JsonToEmployee(message)
	//newEmpl := employee.NewEmployee(empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.HireDate, empl.Position, empl.CompanyID)
	id, err := s.repo.Create(&empl)
	if err != nil {
		log.Fatal(entity.ErrCannotBeCreated)
	}
	producers.KafkaSend([]byte(id), "EmployeePOSTResponse")
}


func (s *Service) UpdateEmployee(message []byte) {
	empl := JsonToEmployee(message)
	response, err := s.repo.Create(&empl)
	if err != nil {
		log.Fatal(entity.ErrCannotBeCreated)
	}
	producers.KafkaSend([]byte(response), "EmployeePUTResponse")
}


func (s *Service) DeleteEmployee(message []byte) {
	id := StringToInt64(string(message))
	response, err := s.repo.Delete(id)
	if err != nil {
		log.Fatal(entity.ErrNotFound)
	}
	producers.KafkaSend([]byte(response), "EmployeeGETResponse")
}

