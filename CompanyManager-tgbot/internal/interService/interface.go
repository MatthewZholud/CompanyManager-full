package interService

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
)

type InterServiceRep interface {
	employee
	company
}

type employee interface {
	GetEmployees () []presenter.Employee
	GetEmployee(id string) (*presenter.Employee, string)
	UpdateEmployee(employee *presenter.Employee) string
}

type company interface {
	GetCompanies () []presenter.Company
	GetCompany(id string) (*presenter.Company, string)
	UpdateCompany(company *presenter.Company) string
}