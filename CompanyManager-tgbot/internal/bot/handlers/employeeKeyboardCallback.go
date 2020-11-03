package handlers

import (
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const (
	EmployeeName       = "EmployeeName"
	EmployeeSecondName = "EmployeeSecondName"
	Surname            = "Surname"
	PhotoUrl           = "PhotoUrl"
	Position           = "Position"
	CompanyID          = "CompanyID"
	HireDate           = "HireDate"
	Save               = "Save"
	Break              = "Break"
)

var employeeNumericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Name", EmployeeName),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Second Name", EmployeeSecondName),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Surname", Surname),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee PhotoUrl", PhotoUrl),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee Position", Position),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Employee CompanyID", CompanyID),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Save", Save),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Break", Break),
	),
)

func (u Handlers) employeeKeyboard(empl *presenter.Employee, msg tgbotapi.MessageConfig, ch chan *presenter.Employee) {
	oldEmployee := &presenter.Employee{
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

	u.Active[int(msg.ChatID)].ButtonInput = msgChan
	msg1 := <-msgChan
	u.Active[int(msg.ChatID)].ButtonInput = nil
	u.switchEmployeeCallBack(msg1, oldEmployee, ch, msg)
}

func (u Handlers) switchEmployeeCallBack(msg1 tgbotapi.CallbackQuery, oldEmployee *presenter.Employee, ch chan *presenter.Employee, msg tgbotapi.MessageConfig) {
	switch msg1.Data {
	case Break:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		oldEmployee = nil
		ch <- oldEmployee
		return

	case Save:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		ch <- oldEmployee
		return

	case EmployeeName:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)

		msg.Text = "Enter new Employee Name:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.Name {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.Name = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.Name)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case EmployeeSecondName:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		msg.Text = "Enter new Employee Second Name:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}

		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.SecondName {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.SecondName = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.SecondName)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case Surname:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		msg.Text = "Enter new Employee Surname:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.Surname {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.Surname = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.Surname)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case PhotoUrl:

		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		msg.Text = "Enter new Photo Url:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.PhotoUrl {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.PhotoUrl = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.PhotoUrl)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case HireDate:

		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		msg.Text = "Enter new Hire Date:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.HireDate {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.HireDate = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.HireDate)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case Position:

		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)
		msg.Text = "Enter new Position:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1
		if msg1.Text == oldEmployee.Position {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.Position = msg1.Text
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.Position)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	case CompanyID:
		logger.Log.Debugf("User %v put button %v", msg1.From.UserName, msg1.Data)

		msg.Text = "Enter new Company ID:"
		_, err := u.Bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Can't send message to user: %v", err)
			return
		}
		mshChan1 := make(chan tgbotapi.Message, 1)
		u.Active[int(msg.ChatID)].SimpleInput = mshChan1
		msg1 := <-mshChan1

		if !IsNumericAndPositive(msg1.Text) {
			logger.Log.Debug("Data is not numeric and positive: %v")
			msg.Text = "Please, try again\nInput is not correct"
			_, err := u.Bot.Send(msg)
			if err != nil {
				logger.Log.Errorf("Can't send message to user: %v", err)
				return
			}
			u.Active[int(msg.ChatID)].SimpleInput = nil
			u.employeeKeyboard(oldEmployee, msg, ch)
			return
		}

		_, response := u.interService.GetCompany(msg1.Text)
		fmt.Println(response)

		if response == CompanyNotFound {
			msg.Text = "Please, try again\nCompany with such ID not found"
			_, err := u.Bot.Send(msg)
			if err != nil {
				logger.Log.Errorf("Can't send message to user: %v", err)
				return
			}
			logger.Log.Info("Company not found")
			u.Active[int(msg.ChatID)].SimpleInput = nil
			u.employeeKeyboard(oldEmployee, msg, ch)
			return
		}

		id, _ := strconv.Atoi(msg1.Text)
		if int64(id) == oldEmployee.CompanyID {
			logger.Log.Debugf("Name has not been changed\n")
		} else {
			oldEmployee.CompanyID = int64(id)
			logger.Log.Debugf("Company LegalForm with id %v changed to %v \n", oldEmployee.ID, oldEmployee.CompanyID)
		}
		u.Active[int(msg.ChatID)].SimpleInput = nil
		u.employeeKeyboard(oldEmployee, msg, ch)
		return

	default:
		logger.Log.Debugf("User %v put unknown button", msg1.From.UserName)
		oldEmployee = nil
		ch <- oldEmployee
		return
	}
}
