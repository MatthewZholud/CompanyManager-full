package main

import (
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
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	routes.RegisterCompanyRoutes(r)
	routes.RegisterProfilingRoutes(r)

	//profiling.RegisterCompanyRoutes(r)
	err := http.ListenAndServe(apiGatewayPort, r)
	if err != nil {
		logger.Log.Fatal("Can't connect to server: ", err)
	}
}
