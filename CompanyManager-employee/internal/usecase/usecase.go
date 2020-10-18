package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity/employee"
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

func JsonToEmployee(msg []byte) employee.Employee {
	employee := employee.Employee{}
	json.Unmarshal(msg, &employee)
	return employee
}