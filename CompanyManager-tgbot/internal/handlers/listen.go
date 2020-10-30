package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


type Updates struct {
	Bot *tgbotapi.BotAPI
	Ch  tgbotapi.UpdatesChannel
	Redis service.RedisRep
}



func  NewUpdateChan(bot *tgbotapi.BotAPI, rep service.RedisRep) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return Updates{
		Ch: ch,
		Bot: bot,
		Redis: rep,
	}
}

func (u Updates) Listen() {
	for update := range u.Ch {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			u.switchCommand(update.Message)
		}
	}
}

func (u Updates) simpleListen() *tgbotapi.Message {
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			continue
		}
			if update.Message.IsCommand() {
			u.switchCommand(update.Message)
		} else {
			return update.Message
		}
	}
	return nil
}

func (u Updates) switchCommand(update *tgbotapi.Message)  {
	msg := tgbotapi.NewMessage(update.Chat.ID, "")
	switch update.Command() {
	case "start":
		u.Redis.Set(update.From.ID)
		msg.Text = fmt.Sprintf("Hello, %s", update.From.UserName)
		u.Redis.Get()
		u.Bot.Send(msg)
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