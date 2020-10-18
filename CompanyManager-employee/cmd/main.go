package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/driver/repository"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/usecase"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/producers"

	"os"

	_ "github.com/lib/pq"
)

func main() {
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "company_manager")

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	conn := repository.NewPostgresRepository(db)
	service := usecase.NewService(conn)


	//consumers.ExampleConsumerGroupParallelReaders()
	//topics := []string{"getCompany", "getEmployee"}

	msg1 := make(chan []byte)
	msg2 := make(chan []byte)

	go consumers.KafkaConsumer("EmployeeGETRequest", msg1)
	go consumers.KafkaConsumer("EmployeePOSTRequest", msg2)

	for {
		select {
		case message := <-msg2:
			service.CreateEmployee(message)
		case message := <-msg1:
			service.GetEmployee(message)
		}
	}

}
