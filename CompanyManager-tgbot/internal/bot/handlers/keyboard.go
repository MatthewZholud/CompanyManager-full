package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
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

func (u Handlers) CompanyKeyboard(comp *presenter.Company, msg tgbotapi.MessageConfig, ch chan *presenter.Company) {
	oldCompany := presenter.Company{
		ID:        comp.ID,
		Name:      comp.Name,
		LegalForm: comp.LegalForm,
	}
	msg.Text = fmt.Sprintf("New Company Info:\nCompany ID: %v\nCompany Name: %s\nCompany Legal form: %s\nSelect what parameter do you whant to change?",
		oldCompany.ID, oldCompany.Name, oldCompany.LegalForm)
	msg.ReplyMarkup = companyNumericKeyboard
	u.Bot.Send(msg)
	msg.ReplyMarkup = nil
	msgChan := make(chan tgbotapi.CallbackQuery, 1)

	u.Active[int(msg.ChatID)] = &bot.Ch{
		SimpleInput: nil,
		ButtonInput: msgChan,
	}
	msg1 := <-msgChan
	u.Active[int(msg.ChatID)] = &bot.Ch{
		SimpleInput: nil,
		ButtonInput: nil,
	}

	switch msg1.Data {
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
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldCompany.Name = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.CompanyKeyboard(&oldCompany, msg, ch)
		return

	case "CompanyLegalForm":
		msg.Text = "Enter new Legal form:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldCompany.LegalForm = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.CompanyKeyboard(&oldCompany, msg, ch)
		return

	default:
		return
	}
}

func (u Handlers) EmployeeKeyboard(empl *presenter.Employee, msg tgbotapi.MessageConfig, ch chan *presenter.Employee) {
	oldEmployee := presenter.Employee{
		ID:         empl.ID,
		Name:       empl.Name,
		SecondName: empl.SecondName,
		Surname:    empl.Surname,
		Position:   empl.Position,
		PhotoUrl:   empl.PhotoUrl,
		HireDate:   empl.HireDate,
		CompanyID:  empl.CompanyID,
	}
	msg.Text = fmt.Sprintf("New Employee Info:\nEmployee ID: %v\nEmployee Name: %s\nEmployee Second "+
		"Name: %s\nEmployee Surname: %s\nEmployee PhotoUrl: %s\nEmployee Position: %s\n"+
		"Employee CompanyID: %v\nSelect what parameter do you whant to change?",
		empl.ID, empl.Name, empl.SecondName, empl.Surname, empl.PhotoUrl, empl.Position, empl.CompanyID)
	msg.ReplyMarkup = employeeNumericKeyboard
	u.Bot.Send(msg)
	msg.ReplyMarkup = nil

	msgChan := make(chan tgbotapi.CallbackQuery, 1)

	u.Active[int(msg.ChatID)] = &bot.Ch{
		SimpleInput: nil,
		ButtonInput: msgChan,
	}
	msg1 := <-msgChan
	u.Active[int(msg.ChatID)] = &bot.Ch{
		SimpleInput: nil,
		ButtonInput: nil,
	}

	switch msg1.Data {
	case "Break":
		empl = nil
		ch <- empl
		return

	case "Save":
		ch <- &oldEmployee
		return

	case "EmployeeName":
		msg.Text = "Enter new Employee Name:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.Name = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "EmployeeSecondName":
		msg.Text = "Enter new Employee Second Name:"
		u.Bot.Send(msg)

		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.SecondName = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "Surname":
		msg.Text = "Enter new Employee Surname:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.Surname = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "PhotoUrl":
		msg.Text = "Enter new Photo Url:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.PhotoUrl = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "HireDate":
		msg.Text = "Enter new Hire Date:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.HireDate = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "Position":
		msg.Text = "Enter new Position:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1
		oldEmployee.Position = msg1.Text
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	case "CompanyID":
		msg.Text = "Enter new Company ID:"
		u.Bot.Send(msg)
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: mshChan1,
			ButtonInput: nil,
		}
		msg1 := <-mshChan1

		if !IsNumericAndPositive(msg1.Text) {
			logger.Log.Debug("Data is not numeric and positive: %v")
			msg.Text = "Please, try again\nInput is not correct"
			u.Bot.Send(msg)
			u.Active[int(msg.ChatID)] = &bot.Ch{
				SimpleInput: nil,
				ButtonInput: nil,
			}
			u.EmployeeKeyboard(&oldEmployee, msg, ch)
			return
		}

		_, response := u.interService.GetCompany(msg1.Text)
		fmt.Println(response)

		if response == CompanyNotFound {
			msg.Text = "Please, try again\nCompany with such ID not found"
			u.Bot.Send(msg)
			logger.Log.Info("Company not found")
			u.Active[int(msg.ChatID)] = &bot.Ch{
				SimpleInput: nil,
				ButtonInput: nil,
			}
			u.EmployeeKeyboard(&oldEmployee, msg, ch)
			return
		}

		id, _ := strconv.Atoi(msg1.Text)

		oldEmployee.CompanyID = int64(id)
		u.Active[int(msg.ChatID)] = &bot.Ch{
			SimpleInput: nil,
			ButtonInput: nil,
		}
		u.EmployeeKeyboard(&oldEmployee, msg, ch)
		return

	default:
		return
	}

}
