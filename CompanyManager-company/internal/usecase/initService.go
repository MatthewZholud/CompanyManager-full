package usecase

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/driver/repository"
)


func InitService(db *sql.DB)  {
	conn := repository.NewPostgresRepository(db)
	service := NewService(conn)
	StartKafkaCommunication(service)
}