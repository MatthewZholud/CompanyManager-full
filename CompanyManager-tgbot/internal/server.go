package internal

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)


var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Name","Company Name"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Legal form","Company Legal form"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Save","Save"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Break","Break"),
	),
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
	for  {
		Listen(updates, bot)
	}
}

func Listen (updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) *tgbotapi.Message  {
	for update := range updates {

		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /getCompanies or /updateCompany."
			case "getCompanies":
				response := handlers.GetCompanies()
				msg.Text = string(response)
				bot.Send(msg)
			case "updateCompany":
				var oldComp presenter.Company
				msg.Text = "Please, enter company id:"
				bot.Send(msg)
				msg1 := Listen(updates, bot)
				msg = tgbotapi.NewMessage(msg1.Chat.ID, msg1.Text)
				compStr := handlers.GetCompany(msg.Text)
				oldComp.Name = compStr.Name
				oldComp.ID = compStr.ID
				oldComp.Legalform = compStr.Legalform

				compStr = ButtonListen(updates, bot, msg, compStr)

				if compStr == nil {
					msg.Text = "Break"
					bot.Send(msg)
					continue
				}
				if oldComp.ID == compStr.ID && oldComp.Name == compStr.Name && oldComp.Legalform == compStr.Legalform{
					msg.Text = "You didn't changed anything:"
					bot.Send(msg)
					continue
				}

				msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
					compStr.ID, compStr.Name, compStr.Legalform)
				bot.Send(msg)
				Successful := handlers.UpdateCompany(compStr)
				if Successful != "Successful update" {
					logger.Log.Errorf("Updating failed: ")
				} else {
					msg.Text = "Successful update"
					bot.Send(msg)
					logger.Log.Infof("Successful update")
				}


			default:
				msg.Text = "I don't know that command"
			}
		} else {
			return update.Message
		}
	}
	return nil
}


func ButtonListen(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, comp *presenter.Company) *presenter.Company {
	msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
		comp.ID, comp.Name, comp.Legalform)
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
	for update := range updates {
		msg.ReplyMarkup = nil
		if update.CallbackQuery != nil{
			switch update.CallbackQuery.Data {
			case "Break":
				comp = nil
				return comp

			case "Save":
				return comp

			case "Company Name":
				msg.Text = "Enter new Company Name:"
				bot.Send(msg)
				msg1 := Listen(updates, bot)
				comp.Name = msg1.Text
				comp = ButtonListen(updates, bot, msg, comp)
				return comp


			case "Company Legal form":
				msg.Text = "Enter new Legal form:"
				bot.Send(msg)
				msg1 := Listen(updates, bot)
				comp.Legalform = msg1.Text
				comp = ButtonListen(updates, bot, msg, comp)
				return comp

			default:
				return nil
			}
		}
	}
	return nil
}
