package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/driver/repository"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase"
	"log"
	"os"

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

	conn := repository.NewPostgresRepository(db)
	service := usecase.NewService(conn)



	msg1 := make(chan []byte)
	msg2 := make(chan []byte)

	go consumers.KafkaConsumer("CompanyGETRequest", msg1)
	go consumers.KafkaConsumer("CompanyPOSTRequest", msg2)


	for {
		select {
		case message := <-msg2:
			service.CreateCompany(message)
		case message := <-msg1:
			service.GetCompany(message)
		}
	}

}
