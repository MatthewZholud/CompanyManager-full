package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Updates struct {
	Bot   *tgbotapi.BotAPI
	Ch    tgbotapi.UpdatesChannel
	Redis service.RedisRep
}

func NewUpdateChan(bot *tgbotapi.BotAPI, rep service.RedisRep) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}
	logger.Log.Info("New Update channel")
	return Updates{
		Ch:    ch,
		Bot:   bot,
		Redis: rep,
	}
}

func (u Updates) Listen() {
	//EmployeeGETRequestChan := make(chan tgbotapi.MessageConfig)
	u.Ch.Clear()
	for update := range u.Ch {
		if update.Message == nil {
			return
		}
		if update.CallbackQuery != nil {
			return
		}
		if update.Message.IsCommand() {
			 u.switchCommand(update.Message)
		}
	}
}

func (u Updates) simpleListen(ch chan *tgbotapi.Message, id int64) {
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			continue
		}
		fmt.Println(u)
		fmt.Println(id, update.Message.Chat.ID)
		if update.Message.Chat.ID == id {
			ch <- update.Message
			close(ch)
			return
		} else if update.Message.IsCommand() {
			u.switchCommand(update.Message)
			ch <- update.Message
			return
		}
		return
	}
}





func (u Updates) switchCommand(update *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(update.Chat.ID, "")
	switch update.Command() {
	case "start":
		u.Redis.Set(update.From.ID)
		msg.Text = fmt.Sprintf("Hello, %s", update.From.UserName)
		u.Redis.Get()
		u.Bot.Send(msg)
		return
	case "help":
		msg.Text = "type /getCompanies  /updateCompany  /getEmployees  /updateEmployee"
		u.Bot.Send(msg)
		return

	case "getCompanies":
		msg = u.GetCompaniesCommand(msg)
		u.Bot.Send(msg)
		return

	case "updateCompany":
		mshChan := make(chan tgbotapi.MessageConfig, 1)

		msg.Text = "Please, enter company id:"
		u.Bot.Send(msg)
		go u.UpdateCompanyCommand(msg, mshChan)
		msg = <-mshChan
		if msg.Text == "continue" {
			break
		}
		u.Bot.Send(msg)
		return

	case "getEmployees":
		msg = u.GetEmployeesCommand(msg)
		u.Bot.Send(msg)
		return

	case "updateEmployee":
		msg.Text = "Please, enter employee id:"
		u.Bot.Send(msg)
		//msg = u.UpdateEmployeeCommand(msg)
		u.Bot.Send(msg)
		return

	default:
		msg.Text = "I don't know this command"
		u.Bot.Send(msg)
		return

	}
}

//func (u Updates) NotifyAll(text string) {
//	sl, err := u.Redis.Get()
//	if err != nil {
//		panic(err)
//	}
//	for i := range sl {
//		c := int64(sl[i])
//		msg := tgbotapi.NewMessage(c, text)
//		u.Bot.Send(msg)
//	}
//}

func (u Updates) SendTextToTelegramChat(chatId int, text string) {

}
