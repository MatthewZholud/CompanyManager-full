package usecase

import "github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity/employee"

type Reader interface {
	Get(id int64) (*employee.Employee, error)
	GetEmployeesByCompany(id int64) (*[]employee.Employee, error)
}

type Writer interface {
	Create(e *employee.Employee) (string, error)
	Update(e *employee.Employee) (string, error)
	Delete(id int64) (string, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetEmployee(message []byte) error
	CreateEmployee(message []byte) error
	UpdateEmployee(message []byte) error
	DeleteEmployee(message []byte) error
	GetEmployeeByCompany(message []byte) error
}