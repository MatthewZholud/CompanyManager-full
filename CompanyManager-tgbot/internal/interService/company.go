package interService

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
)

const (
	CompanyGETRequest     = "CompanyGETRequest"
	CompanyGETAllRequest  = "CompanyGETAllRequest"
	CompanyPUTRequest     = "CompanyPUTRequest"
	CompanyGETResponse    = "CompanyGETResponse"
	CompanyGETAllResponse = "CompanyGETAllResponse"
	CompanyPUTResponse    = "CompanyPUTResponse"
	CompanyNotFound       = "Company not found"
	Error = "Error"
	Success = "Success"
)


func (i *interService) GetCompanies() []presenter.Company {
	mutex.Lock()
	defer mutex.Unlock()
	var companies []presenter.Company
	req := "Get all Request"
	byteUUID, err := i.kafka.KafkaSend([]byte(req), CompanyGETAllRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	msg, err := i.kafka.KafkaGet(CompanyGETAllResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	companies, err = JsonToCompanyArr(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to company array: %v", err)
		return nil
	}
	return companies
}

func (i *interService) GetCompany(id string) (*presenter.Company, string) {
	mutex.Lock()
	defer mutex.Unlock()
	var company *presenter.Company

	byteUUID, err := i.kafka.KafkaSend([]byte(id), CompanyGETRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return nil, Error
	}
	msg, err := i.kafka.KafkaGet(CompanyGETResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return nil, Error
	} else if string(msg) == CompanyNotFound {
		return nil, string(msg)
	}
	company, err = JsonToCompany(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to company struct: %v", err)
		return nil, Error
	}
	return company, Success
}

func (i *interService) UpdateCompany(company *presenter.Company) string {
	mutex.Lock()
	defer mutex.Unlock()
	comp, err := json.Marshal(company)
	if err != nil {
		logger.Log.Errorf("Can't prepare company struct for sending to MessageBroker: %v", err)
		return Error
	}
	byteUUID, err := i.kafka.KafkaSend(comp, CompanyPUTRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return Error
	}
	msg, err := i.kafka.KafkaGet(CompanyPUTResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to MessageBroker: %v", err)
		return Error
	}
	return string(msg)
}
