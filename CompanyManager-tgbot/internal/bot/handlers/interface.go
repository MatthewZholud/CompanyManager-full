package handlers

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type HandlersRep interface {
	SwitchCommand(update *tgbotapi.Message)
}


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

type RedisRep interface {
	Set( msg int)
	Get() ([]int, error)
}