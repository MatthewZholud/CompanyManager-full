package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var companyNumericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Name", "CompanyName"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Company Legal form", "CompanyLegalForm"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Save", "Save"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Break", "Break"),
	),
)

var employeeNumericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Name", "EmployeeName"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Second Name", "EmployeeSecondName"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Surname", "Surname"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee PhotoUrl", "PhotoUrl"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Hire Date", "HireDate"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Position", "Position"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee CompanyID", "CompanyID"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Save", "Save"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Break", "Break"),
	),
)

func (u Updates) ButtonListenCompany(msg tgbotapi.MessageConfig, comp *presenter.Company, ch chan *presenter.Company) {
	oldCompany := presenter.Company{
		ID:        comp.ID,
		Name:      comp.Name,
		Legalform: comp.Legalform,
	}
	msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
		oldCompany.ID, oldCompany.Name, oldCompany.Legalform)
	msg.ReplyMarkup = companyNumericKeyboard
	u.Bot.Send(msg)
	msg.ReplyMarkup = nil
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			if update.CallbackQuery.From.ID == int(msg.ChatID) {

				switch update.CallbackQuery.Data {
				case "Break":
					comp = nil
					ch <- comp
					return

				case "Save":
					ch <- &oldCompany
					return

				case "CompanyName":
					msg.Text = "Enter new Company Name:"
					u.Bot.Send(msg)
					mshChan1 := make(chan *tgbotapi.Message, 1)
					u.simpleListen(mshChan1, msg.ChatID)
					msg1 := <-mshChan1
					if msg1.IsCommand() {
						u.switchCommand(msg1)
						comp = nil
						ch <- comp
						return
					}
					oldCompany.Name = msg1.Text
					u.ButtonListenCompany(msg, &oldCompany, ch)
					return

				case "CompanyLegalForm":
					msg.Text = "Enter new Legal form:"
					u.Bot.Send(msg)
					mshChan1 := make(chan *tgbotapi.Message, 1)
					u.simpleListen(mshChan1, msg.ChatID)
					msg1 := <-mshChan1
					if msg1.IsCommand() {
						u.switchCommand(msg1)
						comp = nil
						ch <- comp
						return
					}
					oldCompany.Legalform = msg1.Text
					u.ButtonListenCompany(msg, &oldCompany, ch)
					return

				default:
					return
				}
			} else {
				continue
			}
		} else if update.Message.IsCommand() {
			u.switchCommand(update.Message)
			comp = nil
			ch <- comp
			return
		}
		continue
	}
	return
}

//func (u Updates) ButtonListenEmployee(msg tgbotapi.MessageConfig, empl *presenter.Employee) *presenter.Employee {
//	msg.Text = fmt.Sprintf("New Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second " +
//		"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl: %s\nEmployee HireDate: %s\nEmployee Position: %s\n" +
//		"Employee CompanyID: %v\nSelect what parameter do you whant to change?",
//		empl.ID, empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.HireDate, empl.Position, empl.CompanyID)
//	msg.ReplyMarkup = employeeNumericKeyboard
//	u.Bot.Send(msg)
//	msg.ReplyMarkup = nil
//	for update := range u.Ch {
//		if update.CallbackQuery != nil {
//			switch update.CallbackQuery.Data {
//			case "Break":
//				empl = nil
//				return empl
//
//			case "Save":
//				return empl
//
//			case "EmployeeName":
//				msg.Text = "Enter new Employee Name:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.Name = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//
//			case "EmployeeSecondName":
//				msg.Text = "Enter new Legal form:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.SecondName = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//			case "Surname":
//				msg.Text = "Enter new Employee Name:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.Surname = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//
//			case "PhotoUrl":
//				msg.Text = "Enter new Legal form:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.PhotoUrl = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//			case "HireDate":
//				msg.Text = "Enter new Employee Name:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.HireDate = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//
//			case "Position":
//				msg.Text = "Enter new Legal form:"
//				u.Bot.Send(msg)
//				msg1 := u.simpleListen()
//				empl.Position = msg1.Text
//				empl = u.ButtonListenEmployee(msg, empl)
//				return empl
//
//			//case "CompanyID":
//			//	msg.Text = "Enter new Legal form:"
//			//	u.Bot.Send(msg)
//			//	msg1 := u.simpleListen()
//			//	empl.CompanyID = msg1.Text
//			//	empl = u.ButtonListenEmployee(msg, empl)
//			//	return empl
//
//			default:
//				return nil
//			}
//		}
//		if update.Message.IsCommand() {
//			return nil
//		}
//
//	}
//	return nil
//}
