package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	CompanyNotFound = "Company not found"
	Success         = "Successful update"
)

func (u Handlers) getCompaniesCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig) {
	response := u.interService.GetCompanies()
	msg.Text = FormatCompanyArr(response)
	ch <- msg
}

func (u Handlers) updateCompanyCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig) {

	mshChan1 := make(chan tgbotapi.Message, 1)

	u.Active[int(msg.ChatID)] = &bot.Channels{
		SimpleInput: mshChan1,
		ButtonInput: nil,
	}
	msg1 := <-mshChan1
	u.Active[int(msg.ChatID)].SimpleInput = nil

	if !IsNumericAndPositive(msg1.Text) {
		logger.Log.Debug("Data is not numeric and positive: %v")
		msg.Text = "Please, try again\nInput is not correct"
		ch <- msg
		return
	}


	company, response := u.interService.GetCompany(msg1.Text)

	if response == CompanyNotFound {
		msg.Text = "Company with such ID not found"
		logger.Log.Debug("Company not found")
		ch <- msg
		return
	}

	oldCompany := presenter.Company{
		ID:        company.ID,
		Name:      company.Name,
		LegalForm: company.LegalForm,
	}

	compFromChan := make(chan *presenter.Company, 1)

	go u.companyKeyboard(company, msg, compFromChan)

	newCompany := <-compFromChan

	if newCompany == nil {
		logger.Log.Debugf("Break updating, user: %v", msg1.From.UserName)
		msg.Text = "continue"
		ch <- msg
		return
	}

	if oldCompany.ID == newCompany.ID && oldCompany.Name == newCompany.Name && oldCompany.LegalForm == newCompany.LegalForm {
		logger.Log.Debugf("User %v didn't change anything: %v", msg1.From.UserName)
		msg.Text = "You didn't change anything:"
		ch <- msg
		return
	}

	response = u.interService.UpdateCompany(newCompany)
	if response != Success {
		msg.Text = "Updating failed"
		logger.Log.Errorf("Updating failed: ")
		ch <- msg
		return
	} else {
		msg.Text = fmt.Sprintf("Successful update\n\nNew Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s",
			newCompany.ID, newCompany.Name, newCompany.LegalForm)
		//go u.NotifyAll(fmt.Sprintf("Company with ID %v was updated.", c.ID))
		logger.Log.Infof("Employee with ID %v was updated.\", e.ID")
		ch <- msg
		return
	}
}
