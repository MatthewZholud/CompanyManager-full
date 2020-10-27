package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/consumers"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
)
const (
	CompanyGETRequest = "CompanyGETRequest"
	CompanyGETAllRequest = "CompanyGETAllRequest"
	CompanyPUTRequest = "CompanyPUTRequest"
	CompanyGETResponse = "CompanyGETResponse"
	CompanyGETAllResponse = "CompanyGETAllResponse"
	CompanyPUTResponse = "CompanyPUTResponse"
)


func GetCompanies () []byte {
	//var companies []presenter.Company
	byteUUID, err := producers.KafkaSend([]byte("Get all Request"), CompanyGETAllRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	msg, err := consumers.KafkaGetStruct(CompanyGETAllResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return nil
	}
	return msg
	//companies, err = JsonToCompanyArr(msg)
	//if err != nil {
	//	logger.Log.Errorf("Can't convert json to employee array: %v", err)
	//	return
	//}
}

func GetCompany(id string) *presenter.Company {
	var company *presenter.Company

	byteUUID, err := producers.KafkaSend([]byte(id), CompanyGETRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil
	}
	msg, err := consumers.KafkaGetStruct(CompanyGETResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to kafka: %v", err)
		return nil
	}
	company, err = JsonToCompany(msg)
	if err != nil {
		logger.Log.Errorf("Can't convert json to company struct: %v", err)
		return nil
	}
	return company
}

func UpdateCompany(company *presenter.Company) string {
	comp, err := json.Marshal(company)
	if err != nil {
		logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
		return ""

	}
	byteUUID, err := producers.KafkaSend(comp, CompanyPUTRequest)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return ""
	}
	msg, err := consumers.KafkaGetStruct(CompanyPUTResponse, byteUUID)
	if err != nil {
		logger.Log.Errorf("Error sending message to env: %v", err)
		return ""
	}
	return string(msg)
}
