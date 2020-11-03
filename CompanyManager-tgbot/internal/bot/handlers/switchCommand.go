package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (u Handlers) SwitchCommand(update *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(update.Chat.ID, "")
	switch update.Command() {
	case "start":
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		u.Redis.Set(update.From.ID)
		msg.Text = fmt.Sprintf("Hello, %s!!!\nWrite /help to see all commands.", update.From.UserName)
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	case "help":
		msg.Text = "type: \n/getCompanies  \n/updateCompany  \n/getEmployees  \n/updateEmployee"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	case "getCompanies":
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		go u.getCompaniesCommand(msg, msgChan)
		msg = <-msgChan
		logger.Log.Debugf("Received message from channel")

		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	case "updateCompany":
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		msg.Text = "Please, enter company id:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		go u.updateCompanyCommand(msg, msgChan)
		msg = <-msgChan
		logger.Log.Debugf("Received message from channel")

		u.Active[int(msg.ChatID)] = nil
		if msg.Text == "continue" {
			break
		}
		_, err = u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	case "getEmployees":
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		go u.getEmployeesCommand(msg, msgChan)
		msg = <-msgChan
		logger.Log.Debugf("Received message from channel: %v", msg.Text)

		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	case "updateEmployee":
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		msg.Text = "Please, enter employee id:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		go u.updateEmployeeCommand(msg, msgChan)
		msg = <-msgChan
		logger.Log.Debugf("Received message from channel: %v", msg.Text)

		u.Active[int(msg.ChatID)] = nil
		if msg.Text == "continue" {
			break
		}
		_, err = u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return

	default:
		logger.Log.Debugf("Command %v from user with %v", update.Text, update.From.UserName)
		msg.Text = "Unknown command"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		return
	}
}
