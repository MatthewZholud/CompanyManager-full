package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	EmployeeNotFound = "Employee not found"
)

func (u Handlers) getEmployeesCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig) {
	response := u.interService.GetEmployees()
	msg.Text = FormatEmployeeArr(response)
	ch <- msg
}

func (u Handlers) updateEmployeeCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig) {
	mshChan1 := make(chan tgbotapi.Message, 1)

	u.Active[int(msg.ChatID)] = &bot.Channels{
		SimpleInput: mshChan1,
		ButtonInput: nil,
	}

	msg1 := <-mshChan1
	if !IsNumericAndPositive(msg1.Text) {
		logger.Log.Debug("Data is not numeric and positive: %v")
		msg.Text = "Please, try again\nInput is not correct"
		ch <- msg
		return
	}

	msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)

	employee, response := u.interService.GetEmployee(msg.Text)
	if response == EmployeeNotFound {
		msg.Text = "Employee not found"
		logger.Log.Info("Employee not found")
		ch <- msg
		return
	}

	oldEmployee := presenter.Employee{
		ID:         employee.ID,
		Name:       employee.Name,
		SecondName: employee.SecondName,
		Surname:    employee.Surname,
		Position:   employee.Position,
		PhotoUrl:   employee.PhotoUrl,
		HireDate:   employee.HireDate,
		CompanyID:  employee.CompanyID,
	}

	employeeFromChan := make(chan *presenter.Employee, 1)
	go u.employeeKeyboard(employee, msg, employeeFromChan)
	newEmployee := <-employeeFromChan

	if newEmployee == nil {
		logger.Log.Debugf("Break updating, user: %v", msg1.From.UserName)
		msg.Text = "continue"
		ch <- msg
		return
	}

	if oldEmployee.ID == newEmployee.ID && oldEmployee.Name == newEmployee.Name && oldEmployee.Surname == newEmployee.Surname &&
			oldEmployee.SecondName == newEmployee.SecondName && oldEmployee.Position == newEmployee.Position &&
			oldEmployee.PhotoUrl == newEmployee.Position && oldEmployee.CompanyID == newEmployee.CompanyID {
		logger.Log.Debugf("User %v didn't change anything: %v", msg1.From.UserName)
		msg.Text = "You didn't change anything:"
		ch <- msg
		return
	}

	response = u.interService.UpdateEmployee(newEmployee)

	if response != Success {
		msg.Text = "Updating failed"
		logger.Log.Errorf("Updating failed: ")
		ch <- msg
		return
	} else {
		msg.Text = fmt.Sprintf("Successful update\n\nNew Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second "+
			"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl:  %s\nEmployee Position: %s\n"+
			"Employee CompanyID: %v",
			newEmployee.ID, newEmployee.Name, newEmployee.SecondName, newEmployee.Surname, newEmployee.PhotoUrl, newEmployee.Position, newEmployee.CompanyID)
		//go u.NotifyAll(fmt.Sprintf("Employee with ID %v was updated.", e.ID))
		logger.Log.Infof("Employee with ID %v was updated.\", e.ID")
		ch <- msg
		return
	}
}
