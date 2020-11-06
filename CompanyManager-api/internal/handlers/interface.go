package handlers

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
)

type InterServiceEmployee interface {
	GetEmployee(id string) (*presenter.Employee, error)
	UpdateEmployee(employee *presenter.Employee) error
	CreateEmployee(company *presenter.Employee) (int64, error)
	DeleteEmployee(id string) error
}

type InterServiceCompany interface {
	GetCompany(id string) (*presenter.Company, error)
	UpdateCompany(company *presenter.Company) error
	CreateCompany(company *presenter.Company) (int64, error)
	DeleteCompany(id string) error
	GetEmployeesByCompany(id string) ([]presenter.Employee, error)
}




