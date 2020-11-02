package company

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/usecase"
	"sync"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
)

const (
	CompanyGETRequest = "CompanyGETRequest"
	CompanyGETAllRequest = "CompanyGETAllRequest"
	CompanyPUTRequest = "CompanyPUTRequest"
	CompanyGETResponse = "CompanyGETResponse"
	CompanyGETAllResponse = "CompanyGETAllResponse"
	CompanyPUTResponse = "CompanyPUTResponse"


	CompanyNotFound = "Company not found"
)

var mutex sync.Mutex


func GetCompanies () []presenter.Company {
	mutex.Lock()
	defer mutex.Unlock()
	var companies []presenter.Company
	byteUUID, err := producers.KafkaSend([]byte("Get all Request"), CompanyGETAllRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	msg, err := consumers.KafkaGet(CompanyGETAllResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	companies, err = usecase.JsonToCompanyArr(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to company array: %v", err)
		return nil
	}
	return companies
}

func GetCompany(id string) (*presenter.Company, string) {
	mutex.Lock()
	defer mutex.Unlock()
	var company *presenter.Company

	byteUUID, err := producers.KafkaSend([]byte(id), CompanyGETRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil, "Error"
	}
	msg, err := consumers.KafkaGet(CompanyGETResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil, "Error"
	} else if string(msg) == CompanyNotFound{
		return nil, string(msg)
	}
	company, err = usecase.JsonToCompany(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to company struct: %v", err)
		return nil, "Error"
	}
	return company, "Success"
}

func UpdateCompany(company *presenter.Company) string {
	mutex.Lock()
	defer mutex.Unlock()
	comp, err := json.Marshal(company)
	if err != nil {
		logger.Log.Errorf("Can't prepare company struct for sending to kafka: %v", err)
		return ""

	}
	byteUUID, err := producers.KafkaSend(comp, CompanyPUTRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return ""
	}
	msg, err := consumers.KafkaGet(CompanyPUTResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return ""
	}
	return string(msg)
}
