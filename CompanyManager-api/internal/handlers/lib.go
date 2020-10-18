package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
	"log"
	"strconv"
)


func JsonToEmployee(msg []byte) (*presenter.Employee, error) {
	employee := presenter.Employee{}
	if err := json.Unmarshal(msg, &employee); err != nil {
		return nil, err
	}
	return &employee, nil
}

func JsonToEmployeeArr(msg []byte) ([]presenter.Employee, error) {
	employee := []presenter.Employee{}
	if err := json.Unmarshal(msg, &employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func JsonToCompany(msg []byte) (*presenter.Company, error) {
	company := presenter.Company{}
	if err := json.Unmarshal(msg, &company); err != nil {
		return nil, err
	}
	return &company, nil
}

func ByteToInt64(msg []byte) int64 {
	str := string(msg)
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return int64(id)
}

func IsNumericAndPositive(s string) bool {
	i, err := strconv.ParseFloat(s, 64)
	if err == nil && i >= 0 {
		return true
	} else {
		return false
	}
}







