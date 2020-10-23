package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/producers"
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

func (s *Service) GetEmployee(message []byte) error {
	id, err := StringToInt64(string(message))
	if err != nil {
		return err
	}
	employee, err := s.repo.Get(id)
	if err != nil {
		return err
	} else {
		logger.Log.Info("Working with database by GET request was successful")
	}
	e, _ := json.Marshal(employee)
	producers.KafkaSend(e, "EmployeeGETResponse")
	return nil
}

func (s *Service) CreateEmployee(message []byte) error{
	empl, err := JsonToCompany(message)
	if err != nil {
		return err
	}
	//newEmpl := employee.NewEmployee(empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.HireDate, empl.Position, empl.CompanyID)
	id, err := s.repo.Create(empl)
	if err != nil {
		return err
	} else {
		logger.Log.Info("Working with database by CREATE request was successful")
	}
	producers.KafkaSend([]byte(id), "EmployeePOSTResponse")
	return nil
}


func (s *Service) UpdateEmployee(message []byte) error {
	empl, err := JsonToCompany(message)
	if err != nil {
		return err
	}
	response, err := s.repo.Update(empl)
	if err != nil {
		return err
	} else {
		logger.Log.Info("Working with database by UPDATE request was successful")
	}
	producers.KafkaSend([]byte(response), "EmployeePUTResponse")
	return nil
}


func (s *Service) DeleteEmployee(message []byte) error {
	id, err := StringToInt64(string(message))
	if err != nil {
		return err
	}
	response, err := s.repo.Delete(id)
	if err != nil {
		return err
	} else {
		logger.Log.Info("Working with database by DELETE request was successful")
	}
	producers.KafkaSend([]byte(response), "EmployeeGETResponse")
	return nil
}

func (s *Service) GetEmployeeByCompany(message []byte) error {
	id, err := StringToInt64(string(message))
	if err != nil {
		return err
	}
	employees, err := s.repo.GetEmployeesByCompany(id)
	if err != nil {
		return err
	} else {
		logger.Log.Info("Working with database by GET(employees by company) request was successful")
	}
	e, _ := json.Marshal(employees)
	producers.KafkaSend(e, "EmployeeByCompanyGETResponse")
	return nil
}
