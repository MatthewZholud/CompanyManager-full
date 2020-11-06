package routes

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"net/http"
)

func RegisterCompanyRoutes(r *mux.Router, comp CompanyRep) *mux.Router {
	logger.Log.Infof("Ready to process company requests ")
	r.HandleFunc("/company/", comp.CreateCompany()).Methods(http.MethodPost)
	r.HandleFunc("/company/", comp.UpdateCompany()).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyId}", comp.GetCompany()).Methods(http.MethodGet)
	r.HandleFunc("/company/{companyId}", comp.FormUpdateCompany()).Methods(http.MethodPost)
	r.HandleFunc("/company/{companyId}", comp.DeleteCompany()).Methods(http.MethodDelete)
	r.HandleFunc("/company/{companyId}/employees", comp.GetEmployeesByCompany()).Methods(http.MethodGet)
	return r
}
