package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"

	//"fmt"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	employeeHandler "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/usecase/employee"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	EmployeeNotFound = "Employee not found"
)

func (u Updates) GetEmployeesCommand(msg tgbotapi.MessageConfig,  ch chan tgbotapi.MessageConfig){
	response := employeeHandler.GetEmployees()
	msg.Text = FormatEmployeeArr(response)
	ch <- msg
}


func (u Updates) UpdateEmployeeCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig){
	mshChan1 := make(chan tgbotapi.Message, 1)

	u.Active[int(msg.ChatID)] = &Ch{
		SimplInput: mshChan1,
		ButtonInput: nil,
	}

	msg1 := <- mshChan1
	if !IsNumericAndPositive(msg1.Text){
		logger.Log.Errorf("Data is not numeric and positive: %v")
		msg.Text = "Please, try again\nInput is not correct"
		ch <- msg
		return
	}


	msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)

	employee, response := employeeHandler.GetEmployee(msg.Text)
	if response == EmployeeNotFound {
		msg.Text = "Employee not found"
		logger.Log.Info("Employee not found")
		ch <- msg
		return
	}

	oldEmployee := presenter.Employee{
		ID: employee.ID,
		Name: employee.Name,
		SecondName: employee.SecondName,
		Surname: employee.Surname,
		Position: employee.Position,
		PhotoUrl: employee.PhotoUrl,
		HireDate: employee.HireDate,
		CompanyID: employee.CompanyID,
	}

	emplFromChan := make(chan *presenter.Employee, 1)
	go u.EmployeeKeyboard(employee, msg, emplFromChan)
	e := <- emplFromChan


	if e == nil {
		msg.Text = "continue"
		ch <- msg
		return
	}


	if (oldEmployee.ID == employee.ID && oldEmployee.Name == employee.Name && oldEmployee.Surname == employee.Surname &&
		oldEmployee.SecondName == employee.SecondName && oldEmployee.Position == employee.Position &&
		oldEmployee.PhotoUrl == employee.Position && oldEmployee.CompanyID == employee.CompanyID && oldEmployee.HireDate == employee.HireDate) {
		msg.Text = "You didn't change anything:"
		ch <- msg
		return
	}

	response = employeeHandler.UpdateCompany(e)

	if response != Success {
		msg.Text = "Updating failed"
		logger.Log.Errorf("Updating failed: ")
		ch <- msg
		return
	} else {
		msg.Text = fmt.Sprintf("Successful update\n\nNew Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second " +
			"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl: %s\nEmployee HireDate: %s\nEmployee Position: %s\n" +
			"Employee CompanyID: %v",
			employee.ID, employee.Name, employee.SecondName, employee.Surname, employee.PhotoUrl, employee.HireDate, employee.Position, employee.CompanyID)
		//u.NotifyAll(fmt.Sprintf("Employee with ID %v was updated.", employee.ID))
		logger.Log.Infof("Successful update")
		ch <- msg
		return
	}
}
