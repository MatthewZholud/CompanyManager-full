package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetEmployee(message []byte) ([]byte, error) {
	id, err := StringToInt64(string(message))
	if err != nil {
		return nil, err
	}
	employee, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by GET request was successful")
	}
	e, _ := json.Marshal(employee)
	return e, nil
}

func (s *Service) CreateEmployee(message []byte) ([]byte, error){
	empl, err := JsonToCompany(message)
	if err != nil {
		return nil, err
	}
	//newEmpl := employee.NewEmployee(empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.HireDate, empl.Position, empl.CompanyID)
	id, err := s.repo.Create(empl)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by CREATE request was successful")
	}
	return []byte(id), nil
}


func (s *Service) UpdateEmployee(message []byte) ([]byte, error) {
	empl, err := JsonToCompany(message)
	if err != nil {
		return nil, err
	}
	response, err := s.repo.Update(empl)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by UPDATE request was successful")
	}
	return []byte(response), nil
}


func (s *Service) DeleteEmployee(message []byte) ([]byte, error) {
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

func (s *Service) GetEmployeeByCompany(message []byte) ([]byte, error) {
	id, err := StringToInt64(string(message))
	if err != nil {
		return nil, err
	}
	employees, err := s.repo.GetEmployeesByCompany(id)
	if err != nil {
		return nil, err
	} else {
		logger.Log.Info("Working with database by GET(employees by company) request was successful")
	}
	e, _ := json.Marshal(employees)
	return e, nil
}
