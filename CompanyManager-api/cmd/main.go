package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/MessageBroker"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/interService"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/routes"

	"github.com/gorilla/mux"
)


func main() {
	logger.InitLog()
	kafka := MessageBroker.Initialize()
	interService := interService.Initialize(kafka)
	company := handlers.InitializeCompany(interService)
	employee := handlers.InitializeEmployee(interService)
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r, employee)
	routes.RegisterCompanyRoutes(r, company)
	routes.RegisterProfilingRoutes(r)

	//profiling.RegisterCompanyRoutes(r)
	port := os.Getenv("API_GATEWAY_PORT")
	err := http.ListenAndServe(port, r)
	if err != nil {
		logger.Log.Fatalf(`Can't listen to port "%v:": %v `, port, err)
	}
}
