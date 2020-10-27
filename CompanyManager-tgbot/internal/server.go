package internal

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		logger.Log.Fatal("Can't connect to telegram bot: ", err)
	} else {
		logger.Log.Infof("Authorized on account %s", bot.Self.UserName)
	}
	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		logger.Log.Errorf("Check GetUpdatesChan Run for %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /sayhi or /status."
			case "getCompanies":
				responce := handlers.GetCompanies()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(responce))
				bot.Send(msg)

				//resp := ListenId(updates, bot)

				//fmt.Printf("%T", resp)
				//handlers.GetCompany(resp)

			default:
				msg.Text = "I don't know that command"
			}
			//bot.Send(msg)
		}
	}
}

func Listen (updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) string  {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		return update.Message.Text
	}
	return ""
}