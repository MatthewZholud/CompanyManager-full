package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"

	"log"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetCompany(message []byte)  {
	id := StringToInt64(string(message))
	company, err := s.repo.Get(id)
	if err != nil {
		log.Fatal(entity.ErrNotFound)
	}
	if err != nil {
		log.Fatal(err)
	}
	c, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	producers.KafkaSend(c, "CompanyGETResponse")
}


func (s *Service) CreateCompany(message []byte) {
	comp := JsonToEmployee(message)
	newComp := company.NewCompany(comp.Name, comp.Legalform)
	id, err := s.repo.Create(newComp)
	if err != nil {
		log.Fatal(entity.ErrCannotBeCreated)
	}
	producers.KafkaSend([]byte(id), "CompanyPOSTResponse")
}
