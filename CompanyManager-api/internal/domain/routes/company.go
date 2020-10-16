package routes


import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/handlers"

	"github.com/gorilla/mux"
	"net/http"
)
//
func RegisterCompanyRoutes(r *mux.Router) *mux.Router {
//
//	//r.HandleFunc("/company/", handlers.CreateCompany).Methods(http.MethodPost)
//	//r.HandleFunc("/company/", handlers.UpdateCompany).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyId}", handlers.GetCompany()).Methods(http.MethodGet)
//	//r.HandleFunc("/company/{companyId}", handlers.FormUpdateCompany).Methods(http.MethodPost)
//	//r.HandleFunc("/company/{companyId}", handlers.DeleteCompany).Methods(http.MethodDelete)
//	//r.HandleFunc("/company/{companyId}/employees", handlers.GetEmployeesByCompany).Methods(http.MethodGet)
//
	return r
}
