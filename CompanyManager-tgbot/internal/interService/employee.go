package interService

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	"sync"
)

const (
	EmployeeGETRequest    = "EmployeeGETRequest"
	EmployeeGETAllRequest = "EmployeeGETAllRequest"
	EmployeePUTRequest    = "EmployeePUTRequest"
	EmployeeGETResponse    = "EmployeeGETResponse"
	EmployeeGETAllResponse = "EmployeeGETAllResponse"
	EmployeePUTResponse    = "EmployeePUTResponse"
	EmployeeNotFound       = "Employee not found"
)

var mutex sync.Mutex

func (i *interService) GetEmployees() []presenter.Employee {
	mutex.Lock()
	defer mutex.Unlock()
	var employees []presenter.Employee
	req := "Get all Request"
	byteUUID, err := i.kafka.KafkaSend([]byte(req), EmployeeGETAllRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	msg, err := i.kafka.KafkaGet(EmployeeGETAllResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	employees, err = JsonToEmployeeArr(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to employee array: %v", err)
		return nil
	}
	return employees
}

func (i *interService) GetEmployee(id string) (*presenter.Employee, string) {
	mutex.Lock()
	defer mutex.Unlock()
	var employee *presenter.Employee

	byteUUID, err := i.kafka.KafkaSend([]byte(id), EmployeeGETRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return nil, Error
	}
	msg, err := i.kafka.KafkaGet(EmployeeGETResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return nil, Error
	} else if string(msg) == EmployeeNotFound {
		return nil, string(msg)
	}
	employee, err = JsonToEmployee(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to employee struct: %v", err)
		return nil, Error
	}
	return employee, Success
}

func (i *interService) UpdateEmployee(employee *presenter.Employee) string {
	mutex.Lock()
	defer mutex.Unlock()
	comp, err := json.Marshal(employee)
	if err != nil {
		logger.Log.Errorf("Can't prepare employee struct for sending to MessageBroker: %v", err)
		return Error
	}
	byteUUID, err := i.kafka.KafkaSend(comp, EmployeePUTRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return Error
	}
	msg, err := i.kafka.KafkaGet(EmployeePUTResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return Error
	}
	return string(msg)
}
