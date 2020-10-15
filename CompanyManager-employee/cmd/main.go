package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/entity/employee"

	"log"

	_ "github.com/lib/pq"
)

const (
	employeeServicePort = ":3443"
)

func main() {
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "mypassword", "company_manager")
	//PsqlInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
	//	config.GetUser(), config.GetPassword(), config.GetHost(), config.GetDBName(), config.GetPort())
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	_ = employee.NewPostgresRepository(db)

	employee.GetFromApiEmployee()


}
