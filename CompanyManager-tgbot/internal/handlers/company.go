package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	companyHandler "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/usecase/company"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	CompanyNotFound = "Company not found"
	Success = "Successful update"
)

func (u Updates) GetCompaniesCommand(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig{
	response := companyHandler.GetCompanies()
	msg.Text = FormatCompanyArr(response)
	return msg
}


func (u Updates) UpdateCompanyCommand(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig{

	msg1 := u.simpleListen()
	if !IsNumericAndPositive(msg1.Text){
		logger.Log.Errorf("Data is not numeric and positive: %v")
		msg.Text = "Please, try again\nInput is not correct"
		return msg
	}
	msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)

	company, response := companyHandler.GetCompany(msg.Text)
	if response == CompanyNotFound {
		msg.Text = "Company not found"
		logger.Log.Info("Company not found")
		return msg
	}

	oldCompany := presenter.Company{
		ID: company.ID,
		Name: company.Name,
		Legalform: company.Legalform,
	}

	company = u.ButtonListenCompany(msg, company)

	if company == nil {
		msg.Text = "Break"
		return msg
	}
	if oldCompany.ID == company.ID && oldCompany.Name == company.Name && oldCompany.Legalform == company.Legalform {
		msg.Text = "You didn't change anything:"
		return msg
	}

	response = companyHandler.UpdateCompany(company)
	if response != Success {
		msg.Text = "Updating failed"
		logger.Log.Errorf("Updating failed: ")
		return msg
	} else {
		msg.Text = fmt.Sprintf("Successful update\n\nNew Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s",
			company.ID, company.Name, company.Legalform)
		logger.Log.Infof("Successful update")
		return msg
	}
}

