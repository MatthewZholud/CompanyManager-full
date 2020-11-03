package handlers

import "net/http"

type CompanyRep interface {
	CreateCompany() http.HandlerFunc
	GetCompany() http.HandlerFunc
	DeleteCompany() http.HandlerFunc
	UpdateCompany() http.HandlerFunc
	FormUpdateCompany() http.HandlerFunc
	GetEmployeesByCompany() http.HandlerFunc
}

type EmployeeRep interface {
	CreateEmployee() http.HandlerFunc
	GetEmployee() http.HandlerFunc
	DeleteEmployee() http.HandlerFunc
	UpdateEmployee() http.HandlerFunc
	FormUpdateEmployee() http.HandlerFunc
}
