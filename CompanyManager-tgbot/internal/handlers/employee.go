package handlers

import (
	//"fmt"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	employeeHandler "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/usecase/employee"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	EmployeeNotFound = "Employee not found"
)

func (u Updates) GetEmployeesCommand(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig{
	response := employeeHandler.GetEmployees()
	msg.Text = FormatEmployeeArr(response)
	return msg
}


//func (u Updates) UpdateEmployeeCommand(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig{

	//msg1 := u.simpleListen()
	//if !IsNumericAndPositive(msg1.Text){
	//	logger.Log.Errorf("Data is not numeric and positive: %v")
	//	msg.Text = "Please, try again\nInput is not correct"
	//	return msg
	//}
	//msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)
	//
	//employee, response := employeeHandler.GetEmployee(msg.Text)
	//if response == EmployeeNotFound {
	//	msg.Text = "Employee not found"
	//	logger.Log.Info("Employee not found")
	//	return msg
	//}
	//
	//oldEmployee := presenter.Employee{
	//	ID: employee.ID,
	//	Name: employee.Name,
	//	SecondName: employee.SecondName,
	//	Surname: employee.Surname,
	//	Position: employee.Position,
	//	PhotoUrl: employee.PhotoUrl,
	//	HireDate: employee.HireDate,
	//	CompanyID: employee.CompanyID,
	//}
	//
	//employee = u.ButtonListenEmployee(msg, employee)
	//
	//if employee == nil {
	//	msg.Text = "Break"
	//	return msg
	//}
	//if (oldEmployee.ID == employee.ID && oldEmployee.Name == employee.Name && oldEmployee.Surname == employee.Surname &&
	//	oldEmployee.SecondName == employee.SecondName && oldEmployee.Position == employee.Position &&
	//	oldEmployee.PhotoUrl == employee.Position && oldEmployee.CompanyID == employee.CompanyID && oldEmployee.HireDate == employee.HireDate) {
	//	msg.Text = "You didn't change anything:"
	//	return msg
	//}
	//
	//response = employeeHandler.UpdateCompany(employee)
	//if response != Success {
	//	msg.Text = "Updating failed"
	//	logger.Log.Errorf("Updating failed: ")
	//	return msg
	//} else {
	//	msg.Text = fmt.Sprintf("Successful update\n\nNew Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second " +
	//		"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl: %s\nEmployee HireDate: %s\nEmployee Position: %s\n" +
	//		"Employee CompanyID: %v\nSelect what parameter do you whant to change?",
	//		employee.ID, employee.Name, employee.SecondName, employee.Surname, employee.PhotoUrl, employee.HireDate, employee.Position, employee.CompanyID)
	//	u.NotifyAll(fmt.Sprintf("Employee with ID %v was updated.", employee.ID))
	//	logger.Log.Infof("Successful update")
	//	return msg
	//}
//}
