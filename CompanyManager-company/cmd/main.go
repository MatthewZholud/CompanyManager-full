package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/usecase"
	"log"
	"os"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/entity/company"
	_ "github.com/lib/pq"
)

const (
	companyServiceAddr = ":4443"
)

func main() {

	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"localhost", "5432", "postgres", "mypassword", "company_manager")
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	conn := company.NewPostgresRepository(db)


	msg1 := make(chan string)
	msg2 := make(chan string)

	go consumers.GetFromApiCompany("getCompany", msg1)

	for {
		select {
		case message := <-msg2:
			fmt.Println("main", message)

		case message := <-msg1:
			id := usecase.MessageService(message)
			company, err := conn.GetCompany(id)
			if err != nil {
				panic(err)
			}
			producers.SendFromApiCompany(company)
		}
	}

}
