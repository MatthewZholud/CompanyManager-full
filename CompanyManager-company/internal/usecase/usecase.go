package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"

	"encoding/json"
	"log"
	"strconv"
)

func  StringToInt64(message string) int64 {
	id, err := strconv.Atoi(message)
	if err != nil {
		log.Fatal(err)
	}
	return int64(id)
}

func JsonToCompany(msg []byte) company.Company {
	employee := company.Company{}
	json.Unmarshal(msg, &employee)
	return employee
}