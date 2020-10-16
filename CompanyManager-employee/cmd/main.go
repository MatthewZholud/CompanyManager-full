package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/entity/employee"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/producers"
	"strconv"

	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/producers"


	"os"


	_ "github.com/lib/pq"
)

const (
	employeeServicePort = ":3443"
)

func main() {
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "time_tracker")

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
	for  {
		getEmployee := consumers.GetFromApiEmployee()

		id, _ := strconv.Atoi(getEmployee)

		id64 := int64(id)

		empl, _ := conn.GetEmployee(id64)

		producers.SendFromApiEmployee(empl)
	}





}
