package routes

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"

	"github.com/gorilla/mux"
	"net/http"
)

func RegisterEmployeeRoutes(r *mux.Router, empl EmployeeRep){
	logger.Log.Infof("Ready to process employee requests ")
	r.HandleFunc("/employee", empl.CreateEmployee()).Methods(http.MethodPost)
	r.HandleFunc("/employee", empl.UpdateEmployee()).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empl.GetEmployee()).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empl.FormUpdateEmployee()).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empl.DeleteEmployee()).Methods(http.MethodDelete)
}
