package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/routes"

	"github.com/gorilla/mux"
)

const (
	apiGatewayPort   = ":8005"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	routes.RegisterCompanyRoutes(r)
	routes.RegisterProfilingRoutes(r)

	//profiling.RegisterCompanyRoutes(r)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))
}
