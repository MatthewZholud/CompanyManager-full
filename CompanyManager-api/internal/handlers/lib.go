package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
	"log"
	"strconv"
)



func JsonToEmployee(msg []byte) presenter.Employee {
	employee := presenter.Employee{}
	json.Unmarshal(msg, &employee)
	return employee
}

func JsonToCompany(msg []byte) presenter.Company {
	company := presenter.Company{}
	json.Unmarshal(msg, &company)
	return company
}

func ByteToInt64(msg []byte) int64 {
	str := string(msg)
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return int64(id)

}







