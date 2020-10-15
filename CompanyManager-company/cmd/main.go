package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/entity/company"
	_ "github.com/lib/pq"
)

const (
	companyServiceAddr = ":4443"
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

	_ = company.NewPostgresRepository(db)

}
