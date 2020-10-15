package routes

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/handlers"

	"github.com/gorilla/mux"
	"net/http"
)

func RegisterEmployeeRoutes(r *mux.Router) *mux.Router {
	//r.HandleFunc("/employee", handlers.CreateEmployee).Methods(http.MethodPost)
	//r.HandleFunc("/employee", handlers.UpdateEmployee).Methods(http.MethodPut)
	fmt.Println(4)
	r.HandleFunc("/employee/{id}", handlers.GetEmployee()).Methods(http.MethodGet)
	//r.HandleFunc("/employee/{id}", handlers.FormUpdateEmployee).Methods(http.MethodPost)
	//r.HandleFunc("/employee/{id}", handlers.DeleteEmployee).Methods(http.MethodDelete)
	return r
}
