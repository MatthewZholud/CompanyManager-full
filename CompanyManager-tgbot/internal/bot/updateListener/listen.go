package updateListener

import (
	botDir "github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot/handlers"
)

func (u Updates) Listen(botHandlerService handlers.HandlersRep) {
	u.Ch.Clear()
	var str *botDir.Channels
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
					go botHandlerService.SwitchCommand(update.Message)
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
					go botHandlerService.SwitchCommand(update.Message)
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
					go botHandlerService.SwitchCommand(update.Message)
					return
				}
			}()
			continue
		}
		continue
	}
}



func checkIfInTheActive(Active map[int]*botDir.Channels, id int) *botDir.Channels {
	if val, ok := Active[id]; ok {
		return val
	} else {
		return nil
	}
}




