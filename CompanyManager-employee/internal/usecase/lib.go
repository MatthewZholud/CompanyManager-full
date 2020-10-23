package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity/employee"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
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

func JsonToCompany(msg []byte) (*employee.Employee, error) {
	employee := employee.Employee{}
	err := json.Unmarshal(msg, &employee)
	if err != nil {
		logger.Log.Debug("Can't convert Json to employee struct")
		return nil, err
	}
	return &employee, nil
}