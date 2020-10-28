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

func (u Updates) ButtonListenCompany(msg tgbotapi.MessageConfig, comp *presenter.Company) *presenter.Company {
	msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
		comp.ID, comp.Name, comp.Legalform)
	msg.ReplyMarkup = companyNumericKeyboard
	u.Bot.Send(msg)
	msg.ReplyMarkup = nil
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "Break":
				comp = nil
				return comp

			case "Save":
				return comp

			case "CompanyName":
				msg.Text = "Enter new Company Name:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				comp.Name = msg1.Text
				comp = u.ButtonListenCompany(msg, comp)
				return comp

			case "CompanyLegalForm":
				msg.Text = "Enter new Legal form:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				comp.Legalform = msg1.Text
				comp = u.ButtonListenCompany(msg, comp)
				return comp

			default:
				return nil
			}
		}
		if update.Message.IsCommand() {
			return nil
		}

	}
	return nil
}


func (u Updates) ButtonListenEmployee(msg tgbotapi.MessageConfig, empl *presenter.Employee) *presenter.Employee {
	msg.Text = fmt.Sprintf("New Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second " +
		"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl: %s\nEmployee HireDate: %s\nEmployee Position: %s\n" +
		"Employee CompanyID: %v\nSelect what parameter do you whant to change?",
		empl.ID, empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.HireDate, empl.Position, empl.CompanyID)
	msg.ReplyMarkup = employeeNumericKeyboard
	u.Bot.Send(msg)
	msg.ReplyMarkup = nil
	for update := range u.Ch {
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "Break":
				empl = nil
				return empl

			case "Save":
				return empl

			case "EmployeeName":
				msg.Text = "Enter new Company Name:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.Name = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl

			case "EmployeeSecondName":
				msg.Text = "Enter new Legal form:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.SecondName = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl
			case "Surname":
				msg.Text = "Enter new Company Name:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.Surname = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl

			case "PhotoUrl":
				msg.Text = "Enter new Legal form:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.PhotoUrl = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl
			case "HireDate":
				msg.Text = "Enter new Company Name:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.HireDate = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl

			case "Position":
				msg.Text = "Enter new Legal form:"
				u.Bot.Send(msg)
				msg1 := u.simpleListen()
				empl.Position = msg1.Text
				empl = u.ButtonListenEmployee(msg, empl)
				return empl

			//case "CompanyID":
			//	msg.Text = "Enter new Legal form:"
			//	u.Bot.Send(msg)
			//	msg1 := u.simpleListen()
			//	empl.CompanyID = msg1.Text
			//	empl = u.ButtonListenEmployee(msg, empl)
			//	return empl

			default:
				return nil
			}
		}
		if update.Message.IsCommand() {
			return nil
		}

	}
	return nil
}