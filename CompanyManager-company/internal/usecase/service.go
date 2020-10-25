package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetCompany(message []byte) ([]byte, error) {
	id, err := StringToInt64(string(message))
	if err != nil {
		return nil, err
	}
	company, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by GET request was successful")
	}
	c, _ := json.Marshal(company)
	return c, nil
}

func (s *Service) CreateCompany(message []byte) ([]byte, error) {
	comp, err := JsonToCompany(message)
	if err != nil {
		return nil, err
	}
	//newComp := company.NewCompany(comp.Name, comp.Legalform)
	id, err := s.repo.Create(comp)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by CREATE request was successful")
	}
	return []byte(id), nil
}

func (s *Service) UpdateCompany(message []byte) ([]byte, error) {
	comp, err := JsonToCompany(message)
	if err != nil {
		return nil, err
	}
	response, err := s.repo.Update(comp)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by UPDATE request was successful")
	}
	return []byte(response), nil

}

func (s *Service) DeleteCompany(message []byte) ([]byte, error) {
	id, err := StringToInt64(string(message))
	if err != nil {
		return nil, err
	}
	response, err := s.repo.Delete(id)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by DELETE request was successful")
	}
	return []byte(response), nil
}
