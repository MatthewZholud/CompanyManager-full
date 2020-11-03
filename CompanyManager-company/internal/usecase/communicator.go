package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
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



func StartKafkaCommunication(service *Service, kafka KafkaRep) {

	CompanyGETRequestChan := make(chan entity.Message)
	CompanyGETAllRequestChan := make(chan entity.Message)
	CompanyPOSTRequestChan := make(chan entity.Message)
	CompanyPUTRequestChan := make(chan entity.Message)
	CompanyDeleteRequestChan := make(chan entity.Message)

	broker := os.Getenv("KAFKA_BROKERS")

	go kafka.KafkaConsumer(CompanyGETRequest, broker, CompanyGETRequestChan)
	go kafka.KafkaConsumer(CompanyGETAllRequest, broker, CompanyGETAllRequestChan)
	go kafka.KafkaConsumer(CompanyPOSTRequest, broker, CompanyPOSTRequestChan)
	go kafka.KafkaConsumer(CompanyPUTRequest, broker, CompanyPUTRequestChan)
	go kafka.KafkaConsumer(CompanyDeleteRequest, broker, CompanyDeleteRequestChan)


	for {
		select {
		case message := <-CompanyGETRequestChan:
			response, err := service.GetCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			kafka.KafkaSend(response, message.Key, CompanyGETResponse)

		case message := <-CompanyGETAllRequestChan:
			response, err := service.GetAllCompany()
			if err != nil {
				logger.Log.Errorf("Can't get all companies", err)
			} else {
				logger.Log.Info("Get all request completed")
			}
			kafka.KafkaSend(response, message.Key, CompanyGETAllResponse)

		case message := <-CompanyPOSTRequestChan:
			response, err := service.CreateCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			kafka.KafkaSend(response, message.Key, CompanyPOSTResponse)

		case message := <-CompanyPUTRequestChan:
			response, err := service.UpdateCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			kafka.KafkaSend(response, message.Key, CompanyPUTResponse)

		case message := <-CompanyDeleteRequestChan:
			response, err := service.DeleteCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			kafka.KafkaSend(response, message.Key, CompanyDeleteResponse)
		}
	}
}
