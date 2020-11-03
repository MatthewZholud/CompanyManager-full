package updateListener

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Updates struct {
	Ch           tgbotapi.UpdatesChannel
	Active       map[int] *bot.Ch
}

func NewUpdateChan(botAPI *tgbotapi.BotAPI) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := botAPI.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}
	Active := bot.New()
	logger.Log.Info("New Update channel")
	return Updates{
		Ch:     ch,
		Active: Active,
	}
}