package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/driver/repository"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/profiling"

	"os"

	_ "github.com/lib/pq"

)

func main() {
	logger.InitLog()
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		logger.Log.Fatal("Can't create connection with Db:", err)
	} else {
		logger.Log.Info("Database connection successfully established")
	}
	defer db.Close()

	conn := repository.NewPostgresRepository(db)
	service := usecase.NewService(conn)


	go profiling.ProfilingServer()

	msg1 := make(chan []byte)
	msg2 := make(chan []byte)
	msg3 := make(chan []byte)
	msg4 := make(chan []byte)

	go consumers.KafkaConsumer("CompanyGETRequest", "kafka:9092", msg1)
	go consumers.KafkaConsumer("CompanyPOSTRequest","kafka:9092", msg2)
	go consumers.KafkaConsumer("CompanyPUTRequest","kafka:9092", msg3)
	go consumers.KafkaConsumer("CompanyDeleteRequest","kafka:9092", msg4)

	for {
		select {
		case message := <-msg2:
			err := service.CreateCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
		case message := <-msg1:
			err := service.GetCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
		case message := <-msg3:
			err := service.UpdateCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
		case message := <-msg4:
			err := service.DeleteCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
		}
	}
}
