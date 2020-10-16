package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/entity/employee"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/usecase"
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

	conn := employee.NewPostgresRepository(db)

	//consumers.ExampleConsumerGroupParallelReaders()
	//topics := []string{"getCompany", "getEmployee"}

	msg1 := make(chan string)
	msg2 := make(chan string)

	go consumers.GetFromApiEmployee("getEmployee", msg1)
	go consumers.GetFromApiEmployee("getCompany", msg2)

	for {
		select {
		case message := <-msg2:
			fmt.Println("main", message)

		case message := <-msg1:
			id := usecase.MessageService(message)
			empl, err := conn.GetEmployee(id)
			if err != nil {
				panic(err)
			}
			producers.SendFromApiEmployee(empl)
		}
	}

}
