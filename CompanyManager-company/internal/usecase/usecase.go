package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"

	"encoding/json"
	"strconv"
)

func StringToInt64(message string) (int64, error) {
	id, err := strconv.Atoi(message)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func JsonToCompany(msg []byte) (*company.Company, error) {
	company := company.Company{}
	err := json.Unmarshal(msg, &company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}