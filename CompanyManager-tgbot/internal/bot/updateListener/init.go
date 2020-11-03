package updateListener

import (
	botDir "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Updates struct {
	Ch     tgbotapi.UpdatesChannel
	Active map[int]*botDir.Channels
}

func NewUpdateChan(botAPI *tgbotapi.BotAPI) Updates {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch, err := botAPI.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}
	StateChecker := botDir.New()
	logger.Log.Info("New Update channel")
	return Updates{
		Ch:     ch,
		Active: StateChecker,
	}
}
