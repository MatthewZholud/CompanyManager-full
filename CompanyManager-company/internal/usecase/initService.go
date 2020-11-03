package usecase

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/driver/repository"
)


func InitService(db *sql.DB, kafka KafkaRep)  {
	conn := repository.NewPostgresRepository(db)
	service := NewService(conn)
	StartKafkaCommunication(service, kafka)
}