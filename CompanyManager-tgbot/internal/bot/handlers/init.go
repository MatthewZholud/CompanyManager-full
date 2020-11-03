package handlers

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handlers struct {
	Bot          *tgbotapi.BotAPI
	Redis        RedisRep
	interService InterServiceRep
	Active       map[int] *bot.Channels
}

func NewHandlerService(botAPI *tgbotapi.BotAPI, redis RedisRep, interService InterServiceRep) Handlers {
	Active := bot.New()
	return Handlers{
		Bot:    botAPI,
		Redis:  redis,
		interService: interService,
		Active: Active,
	}
}