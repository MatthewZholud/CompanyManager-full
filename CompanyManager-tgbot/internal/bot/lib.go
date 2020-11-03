package bot

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	"sort"
	"strconv"
)




func IsNumericAndPositive(s string) bool {
	i, err := strconv.ParseFloat(s, 64)
	if err == nil && i >= 0 {
		return true
	} else {
		return false
	}
}

func FormatCompanyArr(companies []presenter.Company) string {
	message := "List of Companies:\n\n"
	sort.Slice(companies, func(i, j int) (less bool) {
		return companies[i].ID < companies[j].ID
	})
	for i := range companies{
		msg := fmt.Sprintf("Company ID: %v\nCompany Name: %s\n\n", companies[i].ID, companies[i].Name)
		message = message + msg
	}

	return message
}

func FormatEmployeeArr(employees []presenter.Employee) string {
	message := "List of Employees:\n\n"
	sort.Slice(employees, func(i, j int) (less bool) {
		return employees[i].ID < employees[j].ID
	})
	for i := range employees{
		msg := fmt.Sprintf("Employee ID: %v\nEmployee Name: %s\nEmployee Position: %s\nEmployee CompanyID: %v\n\n",
			employees[i].ID,  employees[i].Name, employees[i].Position, employees[i].CompanyID)
		message = message + msg
	}
	return message
}






