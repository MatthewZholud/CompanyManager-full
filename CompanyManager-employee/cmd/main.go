package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/driver/repository"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/profiling"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/usecase"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/domain/kafka/producers"

	"os"

	_ "github.com/lib/pq"
)

func main() {
	logger.InitLog()
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "company_manager")

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		logger.Log.Fatal("Can't create connection with Db:", err)
	} else {
		logger.Log.Info("Database connection successfully established")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal("Can't keep the connection with Db:", err)
	}

	conn := repository.NewPostgresRepository(db)
	service := usecase.NewService(conn)

	go profiling.ProfilingServer()

	msg1 := make(chan []byte)
	msg2 := make(chan []byte)
	msg3 := make(chan []byte)
	msg4 := make(chan []byte)
	msg5 := make(chan []byte)

	go consumers.KafkaConsumer("EmployeeGETRequest", "kafka:9092", msg1)
	go consumers.KafkaConsumer("EmployeePOSTRequest", "kafka:9092", msg2)
	go consumers.KafkaConsumer("EmployeePUTRequest", "kafka:9092", msg3)
	go consumers.KafkaConsumer("EmployeeDeleteRequest", "kafka:9092", msg4)
	go consumers.KafkaConsumer("EmployeeByCompanyGETRequest", "kafka:9092", msg5)

	for {
		select {
		case message := <-msg2:
			err := service.CreateEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't create employee:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
		case message := <-msg1:
			err := service.GetEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't get employee", err)
			} else {
				logger.Log.Info("Get request completed")
			}
		case message := <-msg3:
			err := service.UpdateEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't update employee:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
		case message := <-msg4:
			err := service.DeleteEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't delete employee:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
		case message := <-msg5:
			err := service.GetEmployeeByCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't get employee by company:", err)
			} else {
				logger.Log.Info("Get(employee by company) request completed")
			}
		}
	}
}
