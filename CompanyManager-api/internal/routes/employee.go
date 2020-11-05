package routes

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"

	"github.com/gorilla/mux"
	"net/http"
)

//todo: mzh: why return the parameter that you received in funtion call?
//todo: mzh: same to other registering functions
func RegisterEmployeeRoutes(r *mux.Router, empl handlers.EmployeeRep) *mux.Router {
	logger.Log.Infof("Ready to process employee requests ")
	r.HandleFunc("/employee", empl.CreateEmployee()).Methods(http.MethodPost)
	r.HandleFunc("/employee", empl.UpdateEmployee()).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empl.GetEmployee()).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empl.FormUpdateEmployee()).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empl.DeleteEmployee()).Methods(http.MethodDelete)
	return r
}
