package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


type Updates struct {
	Bot *tgbotapi.BotAPI
	Ch  tgbotapi.UpdatesChannel
}

func  NewUpdateChan(bot *tgbotapi.BotAPI) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return Updates{
		Ch: ch,
		Bot: bot,
	}
}

func (u Updates) Listen() {

	for update := range u.Ch {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /getCompanies  /updateCompany  /getEmployees  /updateEmployee"
				u.Bot.Send(msg)
			case "getCompanies":
				msg = u.GetCompaniesCommand(msg)
				u.Bot.Send(msg)
			case "updateCompany":
				msg.Text = "Please, enter company id:"
				u.Bot.Send(msg)
				msg = u.UpdateCompanyCommand(msg)
				u.Bot.Send(msg)
			case "getEmployees":
				msg = u.GetEmployeesCommand(msg)
				u.Bot.Send(msg)
			case "updateEmployee":
				msg.Text = "Please, enter employee id:"
				u.Bot.Send(msg)
				msg = u.UpdateEmployeeCommand(msg)
				u.Bot.Send(msg)
			default:
				msg.Text = "I don't know this command"
				u.Bot.Send(msg)
			}
		}
	}
}

func (u Updates) simpleListen() *tgbotapi.Message {
	for update := range u.Ch {
		if update.Message.IsCommand() {
			return nil
		} else {
			return update.Message
		}
	}
	return nil
}