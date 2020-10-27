package handlers

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/producers"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka/consumers"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
)
const (
	CompanyGETAllRequest = "CompanyGETAllRequest"
	CompanyGETAllResponse = "CompanyGETAllResponse"

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
