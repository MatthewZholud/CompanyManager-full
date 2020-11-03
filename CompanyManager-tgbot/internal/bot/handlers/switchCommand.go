package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (u Handlers) SwitchCommand(update *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(update.Chat.ID, "")
	switch update.Command() {
	case "start":
		u.Redis.Set(update.From.ID)
		msg.Text = fmt.Sprintf("Hello, %s", update.From.UserName)
		u.Redis.Get()
		u.Bot.Send(msg)
		return
	case "help":
		msg.Text = "type: \n/getCompanies  \n/updateCompany  \n/getEmployees  \n/updateEmployee"
		u.Bot.Send(msg)
		return

	case "getCompanies":
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		go u.GetCompaniesCommand(msg, msgChan)
		msg = <-msgChan
		u.Bot.Send(msg)
		return

	case "updateCompany":
		msgChan := make(chan tgbotapi.MessageConfig, 1)

		msg.Text = "Please, enter company id:"
		u.Bot.Send(msg)
		go u.UpdateCompanyCommand(msg, msgChan)
		msg = <-msgChan
		u.Active[int(msg.ChatID)] = nil

		if msg.Text == "continue" {
			break
		}
		u.Bot.Send(msg)
		return

	case "getEmployees":
		msgChan := make(chan tgbotapi.MessageConfig, 1)

		go u.GetEmployeesCommand(msg, msgChan)
		msg = <- msgChan
		u.Bot.Send(msg)
		return

	case "updateEmployee":
		msgChan := make(chan tgbotapi.MessageConfig, 1)
		msg.Text = "Please, enter employee id:"
		u.Bot.Send(msg)
		go u.UpdateEmployeeCommand(msg, msgChan)
		msg = <-msgChan
		u.Active[int(msg.ChatID)] = nil
		if msg.Text == "continue" {
			break
		}
		u.Bot.Send(msg)
		return

	default:
		msg.Text = "Unknown command"
		u.Bot.Send(msg)
		return

	}
}

//func (u Updates) NotifyAll(text string) {
	//sl, err := u.Redis.Get()
	//if err != nil {
	//	panic(err)
	//}
	//for i := range sl {
	//	c := int64(sl[i])
	//	msg := tgbotapi.NewMessage(c, text)
	//	u.Bot.Send(msg)
	//}
//	return
//}

