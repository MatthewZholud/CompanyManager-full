package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"os"
)

const (
	CompanyGETRequest     = "CompanyGETRequest"
	CompanyPOSTRequest    = "CompanyPOSTRequest"
	CompanyPUTRequest     = "CompanyPUTRequest"
	CompanyDeleteRequest  = "CompanyDeleteRequest"
	CompanyGETResponse    = "CompanyGETResponse"
	CompanyPOSTResponse   = "CompanyPOSTResponse"
	CompanyPUTResponse    = "CompanyPUTResponse"
	CompanyDeleteResponse = "CompanyDeleteResponse"
)

func StartKafkaCommunication(service *Service) {

	CompanyGETRequestChan := make(chan []byte)
	CompanyPOSTRequestChan := make(chan []byte)
	CompanyPUTRequestChan := make(chan []byte)
	CompanyDeleteRequestChan := make(chan []byte)

	broker := os.Getenv("KAFKA_BROKERS")

	go consumers.KafkaConsumer(CompanyGETRequest, broker, CompanyGETRequestChan)
	go consumers.KafkaConsumer(CompanyPOSTRequest, broker, CompanyPOSTRequestChan)
	go consumers.KafkaConsumer(CompanyPUTRequest, broker, CompanyPUTRequestChan)
	go consumers.KafkaConsumer(CompanyDeleteRequest, broker, CompanyDeleteRequestChan)

	for {
		select {
		case message := <-CompanyGETRequestChan:
			response, err := service.GetCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			producers.KafkaSend(response, CompanyGETResponse)

		case message := <-CompanyPOSTRequestChan:
			response, err := service.CreateCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			producers.KafkaSend(response, CompanyPOSTResponse)

		case message := <-CompanyPUTRequestChan:
			response, err := service.UpdateCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			producers.KafkaSend(response, CompanyPUTResponse)

		case message := <-CompanyDeleteRequestChan:
			response, err := service.DeleteCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			producers.KafkaSend(response, CompanyDeleteResponse)
		}
	}
}
