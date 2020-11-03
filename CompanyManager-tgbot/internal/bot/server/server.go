package server

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
}

func Init(token string) Bot {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		logger.Log.Fatal("Can't connect to telegram bot: ", err)
	} else {
		logger.Log.Infof("Authorized on account %s", bot.Self.UserName)
	}
	return Bot{
		BotAPI: bot,
	}
}

func StartBot() Bot {
	bot := Init(os.Getenv("TG_TOKEN"))
	return bot
}
