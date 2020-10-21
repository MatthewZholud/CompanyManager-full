package repository

//import (
//	"database/sql"
//	"fmt"
//	"os"
//	_ "github.com/lib/pq"
//
//)
//
//func InitDb() (*sql.DB, error) {
//	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
//		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
//	db, err := sql.Open("postgres", PsqlInfo)
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}
