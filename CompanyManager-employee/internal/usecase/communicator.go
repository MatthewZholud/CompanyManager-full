package usecase

import (
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

	EmployeeGETRequestChan := make(chan []byte)
	EmployeePOSTRequestChan := make(chan []byte)
	EmployeePUTRequestChan := make(chan []byte)
	EmployeeDeleteRequestChan := make(chan []byte)
	EmployeeByCompanyGETRequestChan := make(chan []byte)

	broker := os.Getenv("KAFKA_BROKERS")

	go consumers.KafkaConsumer(EmployeeGETRequest, broker, EmployeeGETRequestChan)
	go consumers.KafkaConsumer(EmployeePOSTRequest, broker, EmployeePOSTRequestChan)
	go consumers.KafkaConsumer(EmployeePUTRequest, broker, EmployeePUTRequestChan)
	go consumers.KafkaConsumer(EmployeeDeleteRequest, broker, EmployeeDeleteRequestChan)
	go consumers.KafkaConsumer(EmployeeByCompanyGETRequest, broker, EmployeeByCompanyGETRequestChan)


	for {
		select {
		case message := <-EmployeeGETRequestChan:
			response, err := service.GetEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't get company", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			producers.KafkaSend(response, EmployeeGETResponse)

		case message := <-EmployeePOSTRequestChan:
			response, err := service.CreateEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't create company:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			producers.KafkaSend(response, EmployeePOSTResponse)

		case message := <-EmployeePUTRequestChan:
			response, err := service.UpdateEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't update company:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			producers.KafkaSend(response, EmployeePUTResponse)

		case message := <-EmployeeDeleteRequestChan:
			response, err := service.DeleteEmployee(message)
			if err != nil {
				logger.Log.Fatal("Can't delete company:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			producers.KafkaSend(response, EmployeeDeleteResponse)
		case message := <-EmployeeByCompanyGETRequestChan:
			response, err := service.GetEmployeeByCompany(message)
			if err != nil {
				logger.Log.Fatal("Can't get employee by company:", err)
			} else {
				logger.Log.Info("Get(employee by company) request completed")
			}
			producers.KafkaSend(response, EmployeeByCompanyGETResponse)
		}
	}
}
