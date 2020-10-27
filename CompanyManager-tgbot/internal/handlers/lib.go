package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	"net/http"
	"strconv"
)


func JsonToEmployee(msg []byte) (*presenter.Employee, error) {
	employee := presenter.Employee{}
	if err := json.Unmarshal(msg, &employee); err != nil {
		logger.Log.Debug("Can't convert Json to employee struct")

		return nil, err
	}
	return &employee, nil
}

func JsonToEmployeeArr(msg []byte) ([]presenter.Employee, error) {
	employee := []presenter.Employee{}
	if err := json.Unmarshal(msg, &employee); err != nil {
		logger.Log.Debug("Can't convert Json to array of employee struct")
		return nil, err
	}

	return employee, nil
}

func JsonToCompanyArr(msg []byte) ([]presenter.Company, error) {
	companies := []presenter.Company{}
	if err := json.Unmarshal(msg, &companies); err != nil {
		logger.Log.Debug("Can't convert Json to array of company struct")
		return nil, err
	}

	return companies, nil
}

func JsonToCompany(msg []byte) (*presenter.Company, error) {
	company := presenter.Company{}
	if err := json.Unmarshal(msg, &company); err != nil {
		logger.Log.Debug("Can't convert Json to company struct")

		return nil, err
	}
	return &company, nil
}

func ByteToInt64(msg []byte) (int64, error) {
	str := string(msg)
	id, err := strconv.Atoi(str)
	if err != nil {
		logger.Log.Debug("Can't convert byte to int64")
		return 0, err
	}
	return int64(id), nil
}

func IsNumericAndPositive(s string) bool {
	i, err := strconv.ParseFloat(s, 64)
	if err == nil && i >= 0 {
		return true
	} else {
		return false
	}
}


func respondWithError(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errorMessage))
}






