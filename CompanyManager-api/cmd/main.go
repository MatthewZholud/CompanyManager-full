package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/routes"



	"github.com/gorilla/mux"

)

const (
	apiGatewayPort           = ":8005"
	employeeMicroServiceAddr = "localhost:3443"
	companyMicroServiceAddr  = "localhost:4443"
)

type someInterface interface {}

func main() {
	fmt.Println(1)
	r := mux.NewRouter()
	fmt.Println(2)

	routes.RegisterEmployeeRoutes(r)
	fmt.Println(3)

	//routes.RegisterCompanyRoutes(r)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))

}
