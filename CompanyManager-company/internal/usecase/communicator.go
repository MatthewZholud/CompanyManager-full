package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"os"
)

const (
	CompanyGETRequest     = "CompanyGETRequest"
	CompanyGETAllRequest = "CompanyGETAllRequest"
	CompanyPOSTRequest    = "CompanyPOSTRequest"
	CompanyPUTRequest     = "CompanyPUTRequest"
	CompanyDeleteRequest  = "CompanyDeleteRequest"
	CompanyGETResponse    = "CompanyGETResponse"
	CompanyGETAllResponse = "CompanyGETAllResponse"
	CompanyPOSTResponse   = "CompanyPOSTResponse"
	CompanyPUTResponse    = "CompanyPUTResponse"
	CompanyDeleteResponse = "CompanyDeleteResponse"
)



func StartKafkaCommunication(service *Service) {

	CompanyGETRequestChan := make(chan entity.Message)
	CompanyGETAllRequestChan := make(chan entity.Message)
	CompanyPOSTRequestChan := make(chan entity.Message)
	CompanyPUTRequestChan := make(chan entity.Message)
	CompanyDeleteRequestChan := make(chan entity.Message)

	broker := os.Getenv("KAFKA_BROKERS")

	go consumers.KafkaConsumer(CompanyGETRequest, broker, CompanyGETRequestChan)
	go consumers.KafkaConsumer(CompanyGETAllRequest, broker, CompanyGETAllRequestChan)
	go consumers.KafkaConsumer(CompanyPOSTRequest, broker, CompanyPOSTRequestChan)
	go consumers.KafkaConsumer(CompanyPUTRequest, broker, CompanyPUTRequestChan)
	go consumers.KafkaConsumer(CompanyDeleteRequest, broker, CompanyDeleteRequestChan)


	for {
		select {
		case message := <-CompanyGETRequestChan:
			response, err := service.GetCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			producers.KafkaSend(response, message.Key, CompanyGETResponse)

		case message := <-CompanyGETAllRequestChan:
			response, err := service.GetAllCompany()
			if err != nil {
				logger.Log.Errorf("Can't get all companies", err)
			} else {
				logger.Log.Info("Get all request completed")
			}
			producers.KafkaSend(response, message.Key, CompanyGETAllResponse)

		case message := <-CompanyPOSTRequestChan:
			response, err := service.CreateCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			producers.KafkaSend(response, message.Key, CompanyPOSTResponse)

		case message := <-CompanyPUTRequestChan:
			response, err := service.UpdateCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			producers.KafkaSend(response, message.Key, CompanyPUTResponse)

		case message := <-CompanyDeleteRequestChan:
			response, err := service.DeleteCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			producers.KafkaSend(response, message.Key, CompanyDeleteResponse)
		}
	}
}
