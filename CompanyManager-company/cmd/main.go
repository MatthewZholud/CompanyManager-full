package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/profiling"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase"

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

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal("Can't keep the connection with Db:", err)
	}
	go profiling.ProfilingServer()
	usecase.InitService(db)
}
