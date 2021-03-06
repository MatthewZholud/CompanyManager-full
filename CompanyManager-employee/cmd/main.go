package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/profiling"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/usecase"

	"os"

	_ "github.com/lib/pq"
)

const postgres = "postgres"

func main() {
	logger.InitLog()
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "company_manager")

	db, err := sql.Open(postgres, PsqlInfo)
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


	go profiling.ProfilingServer()

	kafka := kafka.Initialize()
	usecase.InitService(db, kafka)
}
