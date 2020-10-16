package main

import (
	"log"
	"net/http"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/routes"

	"github.com/gorilla/mux"
)

const (
	apiGatewayPort   = ":8005"
)

type someInterface interface {}

func main() {
	r := mux.NewRouter()

	routes.RegisterEmployeeRoutes(r)

	//routes.RegisterCompanyRoutes(r)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))


}
