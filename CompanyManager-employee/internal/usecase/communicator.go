package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
	"os"
)

const (
	EmployeeGETRequest     = "EmployeeGETRequest"
	EmployeeGETAllRequest = "EmployeeGETAllRequest"
	EmployeePOSTRequest    = "EmployeePOSTRequest"
	EmployeePUTRequest     = "EmployeePUTRequest"
	EmployeeDeleteRequest  = "EmployeeDeleteRequest"
	EmployeeByCompanyGETRequest = "EmployeeByCompanyGETRequest"
	EmployeeGETResponse    = "EmployeeGETResponse"
	EmployeeGETAllResponse = "EmployeeGETAllResponse"
	EmployeePOSTResponse   = "EmployeePOSTResponse"
	EmployeePUTResponse    = "EmployeePUTResponse"
	EmployeeDeleteResponse = "EmployeeDeleteResponse"
	EmployeeByCompanyGETResponse = "EmployeeByCompanyGETResponse"
)

func StartKafkaCommunication(service *Service, kafka KafkaRep) {

	EmployeeGETRequestChan := make(chan entity.Message)
	EmployeeGETAllRequestChan := make(chan entity.Message)
	EmployeePOSTRequestChan := make(chan entity.Message)
	EmployeePUTRequestChan := make(chan entity.Message)
	EmployeeDeleteRequestChan := make(chan entity.Message)
	EmployeeByCompanyGETRequestChan := make(chan entity.Message)


	broker := os.Getenv("KAFKA_BROKERS")

	go kafka.KafkaConsumer(EmployeeGETRequest, broker, EmployeeGETRequestChan)
	go kafka.KafkaConsumer(EmployeeGETAllRequest, broker, EmployeeGETAllRequestChan)
	go kafka.KafkaConsumer(EmployeePOSTRequest, broker, EmployeePOSTRequestChan)
	go kafka.KafkaConsumer(EmployeePUTRequest, broker, EmployeePUTRequestChan)
	go kafka.KafkaConsumer(EmployeeDeleteRequest, broker, EmployeeDeleteRequestChan)
	go kafka.KafkaConsumer(EmployeeByCompanyGETRequest, broker, EmployeeByCompanyGETRequestChan)


	for {
		select {
		case message := <-EmployeeGETRequestChan:
			response, err := service.GetEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get employee", err)
			} else {
				logger.Log.Info("Get request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeeGETResponse)

		case message := <-EmployeeGETAllRequestChan:
			response, err := service.GetAllEmployee()
			if err != nil {
				logger.Log.Errorf("Can't get all employees", err)
			} else {
				logger.Log.Info("Get all request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeeGETAllResponse)

		case message := <-EmployeePOSTRequestChan:
			response, err := service.CreateEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't create employee:", err)
			} else {
				logger.Log.Info("Create request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeePOSTResponse)

		case message := <-EmployeePUTRequestChan:
			response, err := service.UpdateEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't update employee:", err)
			} else {
				logger.Log.Info("Update request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeePUTResponse)

		case message := <-EmployeeDeleteRequestChan:
			response, err := service.DeleteEmployee(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't delete employee:", err)
			} else {
				logger.Log.Info("Delete request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeeDeleteResponse)
		case message := <-EmployeeByCompanyGETRequestChan:
			response, err := service.GetEmployeeByCompany(message.Value)
			if err != nil {
				logger.Log.Errorf("Can't get employee by company:", err)
			} else {
				logger.Log.Info("Get(employee by company) request completed")
			}
			kafka.KafkaSend(response, message.Key, EmployeeByCompanyGETResponse)
		}
	}
}
