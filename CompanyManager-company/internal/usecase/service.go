package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/producers"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetCompany(message []byte) error {
	id, err := StringToInt64(string(message))
	if err != nil {
		return err
	}
	company, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	c, _ := json.Marshal(company)
	producers.KafkaSend(c, "CompanyGETResponse")
	return nil
}

func (s *Service) CreateCompany(message []byte) error {
	comp, err := JsonToCompany(message)
	if err != nil {
		return err
	}
	//newComp := company.NewCompany(comp.Name, comp.Legalform)
	id, err := s.repo.Create(comp)
	if err != nil {
		return err
	}
	producers.KafkaSend([]byte(id), "CompanyPOSTResponse")
	return nil
}

func (s *Service) UpdateCompany(message []byte) error {
	comp, err := JsonToCompany(message)
	if err != nil {
		return err
	}
	response, err := s.repo.Update(comp)
	if err != nil {
		return err
	}
	producers.KafkaSend([]byte(response), "CompanyPUTResponse")
	return nil

}

func (s *Service) DeleteCompany(message []byte) error {
	id, err := StringToInt64(string(message))
	if err != nil {
		return err
	}
	response, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	producers.KafkaSend([]byte(response), "CompanyDeleteResponse")
	return nil

}
