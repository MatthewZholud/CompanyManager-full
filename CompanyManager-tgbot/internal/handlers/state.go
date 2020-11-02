package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

type Ch struct {
	SimplInput chan tgbotapi.Message
	ButtonInput chan tgbotapi.CallbackQuery
}

type ActiveUsers map[int] *Ch

var lock = &sync.Mutex{}

var instance ActiveUsers

func New() ActiveUsers {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = make(ActiveUsers)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}

