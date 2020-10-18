package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/producers"
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
	comp := JsonToCompany(message)
	//newComp := company.NewCompany(comp.Name, comp.Legalform)
	id, err := s.repo.Create(&comp)
	if err != nil {
		log.Fatal(entity.ErrCannotBeCreated)
	}
	producers.KafkaSend([]byte(id), "CompanyPOSTResponse")
}

func (s *Service) UpdateCompany(message []byte) {
	comp := JsonToCompany(message)
	response, err := s.repo.Update(&comp)
	if err != nil {
		log.Fatal(entity.ErrCannotBeCreated)
	}
	producers.KafkaSend([]byte(response), "CompanyPUTResponse")
}


func (s *Service) DeleteCompany(message []byte)  {
	id := StringToInt64(string(message))
	response, err := s.repo.Delete(id)
	if err != nil {
		log.Fatal(entity.ErrNotFound)
	}
	producers.KafkaSend([]byte(response), "CompanyDeleteResponse")
}
