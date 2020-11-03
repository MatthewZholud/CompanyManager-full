package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var companyNumericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Name", "CompanyName"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Legal form", "CompanyLegalForm"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Save", "Save"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Break", "Break"),
	),
)

func (u Handlers) companyKeyboard(comp *presenter.Company, msg tgbotapi.MessageConfig, ch chan *presenter.Company) {
	oldCompany := &presenter.Company{
		ID:        comp.ID,
		Name:      comp.Name,
		LegalForm: comp.LegalForm,
	}
	msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
		oldCompany.ID, oldCompany.Name, oldCompany.LegalForm)
	msg.ReplyMarkup = companyNumericKeyboard
	_, err := u.Bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Can't send message to user: %v", err)
		return
	}
	msg.ReplyMarkup = nil

	msgChan := make(chan tgbotapi.CallbackQuery, 1)
	u.Active[int(msg.ChatID)].ButtonInput = msgChan
	msg1 := <-msgChan
	u.Active[int(msg.ChatID)].ButtonInput = nil

	u.SwitchCallBack(msg1, oldCompany, ch, msg)

}

func (u Handlers) SwitchCallBack(msg1 tgbotapi.CallbackQuery, oldCompany *presenter.Company, ch chan *presenter.Company, msg tgbotapi.MessageConfig) {

	switch msg1.Data {
	case "Break":
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		oldCompany = nil
		ch <- oldCompany
		return

	case "Save":
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		ch <- oldCompany
		return

	case "CompanyName":
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)

		msg.Text = "Enter new Company Name:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldCompany.Name {
			logger.Log.Debugf("Name has not benn changed\n")
		} else {
			oldCompany.Name = msg1.Text
			logger.Log.Debugf("Company name with id %v changed to %v \n", oldCompany.ID, oldCompany.Name)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.companyKeyboard(oldCompany, msg, ch)
		return

	case "CompanyLegalForm":
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)

		msg.Text = "Enter new Legal form:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		oldCompany.LegalForm = msg1.Text
		if msg1.Text == oldCompany.LegalForm {
			logger.Log.Debugf("Name has not benn changed\n")
		} else {
			oldCompany.LegalForm = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldCompany.ID, oldCompany.LegalForm)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.companyKeyboard(oldCompany, msg, ch)
		return

	default:
		logger.Log.Debugf("User %v put unknown button", msg1.From.UserName)
		oldCompany = nil
		ch <- oldCompany
		return
	}
}
