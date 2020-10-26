package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"strconv"
)

func StringToInt64(message string) (int64, error) {
	id, err := strconv.Atoi(message)

	if err != nil {
		logger.Log.Debug("Can't convert String to int64")
		return 0, err
	}
	return int64(id), nil
}

func JsonToCompany(msg []byte) (*entity.Company, error) {
	company := entity.Company{}
	err := json.Unmarshal(msg, &company)
	if err != nil {
		logger.Log.Debug("Can't convert Json to company struct")
		return nil, err
	}
	return &company, nil
}
