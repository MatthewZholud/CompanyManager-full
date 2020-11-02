package employee

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/usecase"
	"sync"
)

const (
	EmployeeGETRequest = "EmployeeGETRequest"
	EmployeeGETAllRequest = "EmployeeGETAllRequest"
	EmployeePUTRequest = "EmployeePUTRequest"


	EmployeeGETResponse = "EmployeeGETResponse"
	EmployeeGETAllResponse = "EmployeeGETAllResponse"
	EmployeePUTResponse = "EmployeePUTResponse"
	EmployeeNotFound = "Employee not found"


)

var mutex sync.Mutex



func GetEmployees () []presenter.Employee {
	mutex.Lock()
	defer mutex.Unlock()
	var employees []presenter.Employee
	byteUUID, err := producers.KafkaSend([]byte("Get all Request"), EmployeeGETAllRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	msg, err := consumers.KafkaGet(EmployeeGETAllResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	employees, err = usecase.JsonToEmployeeArr(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to employee array: %v", err)
		return nil
	}
	return employees
}

func GetEmployee(id string) (*presenter.Employee, string) {
	mutex.Lock()
	defer mutex.Unlock()
	var employee *presenter.Employee

	byteUUID, err := producers.KafkaSend([]byte(id), EmployeeGETRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil, "Error"
	}
	msg, err := consumers.KafkaGet(EmployeeGETResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil, "Error"
	} else if string(msg) == EmployeeNotFound{
		return nil, string(msg)
	}
	employee, err = usecase.JsonToEmployee(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to employee struct: %v", err)
		return nil, "Error"
	}
	return employee, "Success"
}

func UpdateCompany(employee *presenter.Employee) string {
	mutex.Lock()
	defer mutex.Unlock()
	comp, err := json.Marshal(employee)
	if err != nil {
		logger.Log.Errorf("Can't prepare employee struct for sending to kafka: %v", err)
		return ""
	}
	byteUUID, err := producers.KafkaSend(comp, EmployeePUTRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return ""
	}
	msg, err := consumers.KafkaGet(EmployeePUTResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return ""
	}
	return string(msg)
}