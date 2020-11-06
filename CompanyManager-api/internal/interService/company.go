package interService

import (
	"encoding/json"
	"errors"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
)

const (
	CompanyGETRequest            = "CompanyGETRequest"
	CompanyPOSTRequest           = "CompanyPOSTRequest"
	CompanyPUTRequest            = "CompanyPUTRequest"
	CompanyDeleteRequest         = "CompanyDeleteRequest"
	CompanyGETResponse           = "CompanyGETResponse"
	CompanyPOSTResponse          = "CompanyPOSTResponse"
	CompanyPUTResponse           = "CompanyPUTResponse"
	CompanyDeleteResponse        = "CompanyDeleteResponse"
	EmployeeByCompanyGETResponse = "EmployeeByCompanyGETResponse"
	EmployeeByCompanyGETRequest  = "EmployeeByCompanyGETRequest"
	CompanyNotFound              = "Company not found"
	SuccessUdt                   = "Successful update"
)

func (i *interService) GetCompany(id string) (*presenter.Company, error) {
	var company *presenter.Company

	byteUUID, err := i.kafka.BrokerSend([]byte(id), CompanyGETRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	}
	msg, err := i.kafka.BrokerGet(CompanyGETResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	} else if string(msg) == CompanyNotFound {
		logger.Log.Debug("Company with id %v not found", err)
		return nil, errors.New("company not found")
	}
	company, err = JsonToCompany(msg)
	if err != nil {
		logger.Log.Debugf("Can't convert json to company struct: %v", err)
		return nil, err
	}
	return company, nil
}

func (i *interService) UpdateCompany(company *presenter.Company) error {
	comp, err := json.Marshal(company)
	if err != nil {
		logger.Log.Debugf("Can't prepare company struct for sending to MessageBroker: %v", err)
		return err
	}
	byteUUID, err := i.kafka.BrokerSend(comp, CompanyPUTRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	}
	msg, err := i.kafka.BrokerGet(CompanyPUTResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	} else if string(msg) != SuccessUdt {
		return errors.New("update was not successful")
	}
	return nil
}

func (i *interService) CreateCompany(company *presenter.Company) (int64, error) {
	comp, err := json.Marshal(company)
	if err != nil {
		logger.Log.Debugf("Can't prepare company struct for sending to MessageBroker: %v", err)
		return 0, err
	}
	byteUUID, err := i.kafka.BrokerSend(comp, CompanyPOSTRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return 0, err
	}
	msg, err := i.kafka.BrokerGet(CompanyPOSTResponse, byteUUID)
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

func (i *interService) DeleteCompany(id string) error {
	byteUUID, err := i.kafka.BrokerSend([]byte(id), CompanyDeleteRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	}
	msg, err := i.kafka.BrokerGet(CompanyDeleteResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return err
	} else if string(msg) == CompanyNotFound {
		logger.Log.Debug("Company with id %v not found", err)
		return errors.New("company not found")
	}
	return nil
}

func (i *interService) GetEmployeesByCompany(id string) ([]presenter.Employee, error) {

	var employee []presenter.Employee

	byteUUID, err := i.kafka.BrokerSend([]byte(id), EmployeeByCompanyGETRequest)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	}
	msg, err := i.kafka.BrokerGet(EmployeeByCompanyGETResponse, byteUUID)
	if err != nil {
		logger.Log.Debugf("Error sending message to MessageBroker: %v", err)
		return nil, err
	} else if string(msg) == CompanyNotFound {
		logger.Log.Debug("Company with id %v not found", err)
		return nil, errors.New("company not found")
	}
	employee, err = JsonToEmployeeArr(msg)
	if err != nil {
		logger.Log.Debug("Can't convert json to employee array: %v", err)
		return nil, err
	}

	return employee, nil
}
