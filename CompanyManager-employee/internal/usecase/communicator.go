package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
	"os"
)

const (
	EmployeeGETRequest     = "EmployeeGETRequest"
	EmployeePOSTRequest    = "EmployeePOSTRequest"
	EmployeePUTRequest     = "EmployeePUTRequest"
	EmployeeDeleteRequest  = "EmployeeDeleteRequest"
	EmployeeByCompanyGETRequest = "EmployeeByCompanyGETRequest"
	EmployeeGETResponse    = "EmployeeGETResponse"
	EmployeePOSTResponse   = "EmployeePOSTResponse"
	EmployeePUTResponse    = "EmployeePUTResponse"
	EmployeeDeleteResponse = "EmployeeDeleteResponse"
	EmployeeByCompanyGETResponse = "EmployeeByCompanyGETResponse"
)

func StartKafkaCommunication(service *Service) {

	EmployeeGETRequestChan := make(chan entity.Message)
	EmployeePOSTRequestChan := make(chan entity.Message)
	EmployeePUTRequestChan := make(chan entity.Message)
	EmployeeDeleteRequestChan := make(chan entity.Message)
	EmployeeByCompanyGETRequestChan := make(chan entity.Message)

	broker := os.Getenv("KAFKA_BROKERS")

	go consumers.KafkaConsumer(EmployeeGETRequest, broker, EmployeeGETRequestChan)
	go consumers.KafkaConsumer(EmployeePOSTRequest, broker, EmployeePOSTRequestChan)
	go consumers.KafkaConsumer(EmployeePUTRequest, broker, EmployeePUTRequestChan)
	go consumers.KafkaConsumer(EmployeeDeleteRequest, broker, EmployeeDeleteRequestChan)
	go consumers.KafkaConsumer(EmployeeByCompanyGETRequest, broker, EmployeeByCompanyGETRequestChan)


	for {
		select {
		case message := <-EmployeeGETRequestChan:
			response, err := service.GetEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			producers.KafkaSend(response, message.Key, EmployeeGETResponse)

		case message := <-EmployeePOSTRequestChan:
			response, err := service.CreateEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			producers.KafkaSend(response, message.Key, EmployeePOSTResponse)

		case message := <-EmployeePUTRequestChan:
			response, err := service.UpdateEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			producers.KafkaSend(response, message.Key, EmployeePUTResponse)

		case message := <-EmployeeDeleteRequestChan:
			response, err := service.DeleteEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			producers.KafkaSend(response, message.Key, EmployeeDeleteResponse)
		case message := <-EmployeeByCompanyGETRequestChan:
			response, err := service.GetEmployeeByCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get employee by company:", err)
			} else {
				logger.Log.Info("Get(employee by company) request completed")
			}
			producers.KafkaSend(response, message.Key, EmployeeByCompanyGETResponse)
		}
	}
}
