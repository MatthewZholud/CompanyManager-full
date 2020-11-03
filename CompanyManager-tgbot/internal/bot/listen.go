package bot

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/interService"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Updates struct {
	Bot    *tgbotapi.BotAPI
	Ch     tgbotapi.UpdatesChannel
	Redis  redis.RedisRep
	interService interService.InterServiceRep
	Active map[int] *Ch
}

func NewUpdateChan(bot *tgbotapi.BotAPI, rep redis.RedisRep, usecase interService.InterServiceRep) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}
	Active := New()
	logger.Log.Info("New Update channel")
	return Updates{
		Ch:     ch,
		Bot:    bot,
		Redis:  rep,
		interService: usecase,
		Active: Active,
	}
}




func (u Updates) Listen() {
	u.Ch.Clear()
	var str *Ch
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			str = checkIfInTheActive(u.Active, update.CallbackQuery.From.ID)
		} else {
			str = checkIfInTheActive(u.Active, update.Message.From.ID)
		}
		if str != nil {
			if str.ButtonInput != nil {
				if update.CallbackQuery != nil {
					str.ButtonInput <- *update.CallbackQuery
					continue
				}
				if update.Message.IsCommand() {
					u.Active[update.Message.From.ID] = nil
					go u.switchCommand(update.Message)
					continue
				} else {
					continue
				}
			} else if str.SimpleInput != nil {
				if update.CallbackQuery != nil {
					continue
				}
				if update.Message.IsCommand() {
					u.Active[update.Message.From.ID] = nil
					go u.switchCommand(update.Message)
				} else {
					str.SimpleInput <- *update.Message
					continue
				}
			}
		} else {
			go func() {
				if update.Message == nil {
					return
				}
				if update.CallbackQuery != nil {
					return
				}
				if update.Message.IsCommand() {
					go u.switchCommand(update.Message)
					return
				}
			}()
			continue
		}
		continue
	}
}



func checkIfInTheActive(Active map[int]*Ch, id int) *Ch {
	if val, ok := Active[id]; ok {
		return val
	} else {
		return nil
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

func (u Updates) NotifyAll(text string) {
	//sl, err := u.Redis.Get()
	//if err != nil {
	//	panic(err)
	//}
	//for i := range sl {
	//	c := int64(sl[i])
	//	msg := tgbotapi.NewMessage(c, text)
	//	u.Bot.Send(msg)
	//}
	return
}

