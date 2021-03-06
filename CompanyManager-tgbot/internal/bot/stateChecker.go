package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

type Channels struct {
	SimpleInput chan tgbotapi.Message
	ButtonInput chan tgbotapi.CallbackQuery
}

type ActiveUsers map[int] *Channels

var lock = &sync.Mutex{}

var state ActiveUsers

func New() ActiveUsers {
	if state == nil {
		lock.Lock()
		defer lock.Unlock()
		if state == nil {
			state = make(ActiveUsers)
		}
	}
	return state
}

