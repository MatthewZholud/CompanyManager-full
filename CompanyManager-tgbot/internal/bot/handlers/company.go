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
	Success = "Successful update"
)

func (u Handlers) GetCompaniesCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig){
	response := u.interService.GetCompanies()
	msg.Text = FormatCompanyArr(response)
	ch <- msg
}


func (u Handlers) UpdateCompanyCommand(msg tgbotapi.MessageConfig, ch chan tgbotapi.MessageConfig){
	mshChan1 := make(chan tgbotapi.Message, 1)

	u.Active[int(msg.ChatID)] = &bot.Ch{
		SimpleInput: mshChan1,
		ButtonInput: nil,
	}

	msg1 := <- mshChan1

	if !IsNumericAndPositive(msg1.Text){
		logger.Log.Debug("Data is not numeric and positive: %v")
		msg.Text = "Please, try again\nInput is not correct"
		ch <- msg
		return
	}

	msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)

	company, response := u.interService.GetCompany(msg.Text)


	if response == CompanyNotFound {
		msg.Text = "Company with such ID not found"
		logger.Log.Info("Company not found")
		ch <- msg
		return
	}

	oldCompany := presenter.Company{
		ID: company.ID,
		Name: company.Name,
		LegalForm: company.LegalForm,
	}



	compFromChan := make(chan *presenter.Company, 1)
	go u.CompanyKeyboard(company, msg, compFromChan)
	c := <- compFromChan

	if c == nil {
		msg.Text = "continue"
		ch <- msg
		return
	}

	if oldCompany.ID == c.ID && oldCompany.Name == c.Name && oldCompany.LegalForm == c.LegalForm {
		msg.Text = "You didn't change anything:"
		ch <- msg
		return
	}


	response = u.interService.UpdateCompany(c)
	if response != Success {
		msg.Text = "Updating failed"
		logger.Log.Errorf("Updating failed: ")
		ch <- msg
		return
	} else {
		msg.Text = fmt.Sprintf("Successful update\n\nNew Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s",
			c.ID, c.Name, c.LegalForm)
		//go u.NotifyAll(fmt.Sprintf("Company with ID %v was updated.", c.ID))
		logger.Log.Infof("Successful update")
		ch <- msg
		return
	}
}

