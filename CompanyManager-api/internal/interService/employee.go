package interService

import (
	"encoding/json"
	"errors"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
)

const (
	EmployeeGETRequest     = "EmployeeGETRequest"
	EmployeePOSTRequest    = "EmployeePOSTRequest"
	EmployeePUTRequest     = "EmployeePUTRequest"
	EmployeeDeleteRequest  = "EmployeeDeleteRequest"
	EmployeeGETResponse    = "EmployeeGETResponse"
	EmployeePOSTResponse   = "EmployeePOSTResponse"
	EmployeePUTResponse    = "EmployeePUTResponse"
	EmployeeDeleteResponse = "EmployeeDeleteResponse"

	EmployeeNotFound = "Employee not found"
)

func (i *interService) GetEmployee(id string) (*presenter.Employee, error) {
	var employee *presenter.Employee

	byteUUID, err := i.kafka.BrokerSend([]byte(id), EmployeeGETRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	}
	msg, err := i.kafka.BrokerGet(EmployeeGETResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	} else if string(msg) == EmployeeNotFound {
		return nil, errors.New("employee not found")
	}
	employee, err = JsonToEmployee(msg)
	if err != nil {
		logger.Log.Debugf("Can't convert json to employee struct: %v", err)
		return nil, err
	}
	return employee, nil
}

func (i *interService) UpdateEmployee(employee *presenter.Employee) error {
	comp, err := json.Marshal(employee)
	if err != nil {
		logger.Log.Debugf("Can't prepare employee struct for sending to MessageBroker: %v", err)
		return err
	}
	byteUUID, err := i.kafka.BrokerSend(comp, EmployeePUTRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	}
	msg, err := i.kafka.BrokerGet(EmployeePUTResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	} else if string(msg) != SuccessUdt {
		return errors.New("update was not successful")
	}
	return nil
}

func (i *interService) CreateEmployee(employee *presenter.Employee) (int64, error) {
	empl, err := json.Marshal(employee)
	if err != nil {
		logger.Log.Debugf("Can't prepare employee struct for sending to MessageBroker: %v", err)
		return 0, err
	}
	byteUUID, err := i.kafka.BrokerSend(empl, EmployeePOSTRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return 0, err
	}
	msg, err := i.kafka.BrokerGet(EmployeePOSTResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return 0, err
	}
	id, err := ByteToInt64(msg)
	if err != nil {
		logger.Log.Debugf("Can't convert json to int: %v", err)
		return 0, err
	}
	return id, nil
}

func (i *interService) DeleteEmployee(id string) error {
	byteUUID, err := i.kafka.BrokerSend([]byte(id), EmployeeDeleteRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	}
	msg, err := i.kafka.BrokerGet(EmployeeDeleteResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	} else if string(msg) == EmployeeNotFound {
		logger.Log.Debug("employee with id %v not found", err)
		return errors.New("employee not found")
	}
	return nil
}
