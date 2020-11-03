package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"net/http"

	_ "net/http/pprof"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/routes"

	"github.com/gorilla/mux"
)

const (
	apiGatewayPort   = ":8005"
)

func main() {
	logger.InitLog()
	kafka := kafka.Initialize()
	company := handlers.InitializeCompany(kafka)
	employee := handlers.InitializeEmployee(kafka)
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r, employee)
	routes.RegisterCompanyRoutes(r, company)
	routes.RegisterProfilingRoutes(r)

	//profiling.RegisterCompanyRoutes(r)
	err := http.ListenAndServe(apiGatewayPort, r)
	if err != nil {
		logger.Log.Fatal("Can't connect to botServer: ", err)
	}
}
