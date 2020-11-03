package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"strconv"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"

	"github.com/gorilla/mux"

	"net/http"
)

const (
	EmployeeGETRequest     = "EmployeeGETRequest"
	EmployeePOSTRequest    = "EmployeePOSTRequest"
	EmployeePUTRequest     = "EmployeePUTRequest"
	EmployeeDeleteRequest  = "EmployeeDeleteRequest"
	EmployeeGETResponse    = "EmployeeGETResponse"
	EmployeePOSTResponse   = "EmployeePOSTResponse"
	EmployeePUTResponse    = "EmployeePUTResponse"
	EmployeeDeleteResponse = "EmployeeDeleteResponse"
)

type employeeService struct {
	kafka kafka.KafkaRep
}

func InitializeEmployee(kafka kafka.KafkaRep) *employeeService {
	return &employeeService{
		kafka: kafka,
	}
}

func (e *employeeService) CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error adding employee"
		var input struct {
			Name       string `json:"name"`
			SecondName string `json:"second_name"`
			Surname    string `json:"surname"`
			PhotoUrl   string `json:"photo_url"`
			HireDate   string `json:"hire_date"`
			Position   string `json:"position"`
			CompanyID  int64  `json:"company_id"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Errorf("Can't get employee struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		empl, err := json.Marshal(input)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		byteUUID, err := e.kafka.KafkaSend(empl, EmployeePOSTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := e.kafka.KafkaGet(EmployeePOSTResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		id, err := ByteToInt64(msg)
		if err != nil {
			respondWithError(w, errorMessage)
			logger.Log.Errorf("Can't convert byte to int: %v", err)
			return
		}
		toJ := &presenter.Employee{
			ID:         id,
			Name:       input.Name,
			SecondName: input.SecondName,
			Surname:    input.Surname,
			PhotoUrl:   input.PhotoUrl,
			HireDate:   input.HireDate,
			Position:   input.Position,
			CompanyID:  input.CompanyID,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			logger.Log.Errorf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Create handler completed successfully")
		}
	}
}

func (e *employeeService) GetEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading employees"
		var employee *presenter.Employee
		byteUUID, err := e.kafka.KafkaSend([]byte(mux.Vars(r)["id"]), EmployeeGETRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := e.kafka.KafkaGet(EmployeeGETResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		employee, err = JsonToEmployee(msg)
		if err != nil {
			logger.Log.Errorf("Can't convert json to employee struct: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if err := json.NewEncoder(w).Encode(employee); err != nil {
			logger.Log.Errorf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Get handler completed successfully")
		}
	}
}

func (e *employeeService) DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error deleting employee"
		byteUUID, err := e.kafka.KafkaSend([]byte(mux.Vars(r)["id"]), EmployeeDeleteRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := e.kafka.KafkaGet(EmployeeDeleteResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful delete" {
			logger.Log.Errorf("Deleting failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful delete")
		}
	}
}

func (e *employeeService) UpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error updating employee"
		var update struct {
			Name       string `json:"name"`
			SecondName string `json:"second_name"`
			Surname    string `json:"surname"`
			PhotoUrl   string `json:"photo_url"`
			HireDate   string `json:"hire_date"`
			Position   string `json:"position"`
			CompanyID  int64  `json:"company_id"`
		}

		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		empl, err := json.Marshal(update)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		byteUUID, err := e.kafka.KafkaSend(empl, EmployeePUTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := e.kafka.KafkaGet(EmployeePUTResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			respondWithError(w, errorMessage)
			logger.Log.Errorf("Updating failed: %v", err)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}

func (e *employeeService) FormUpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating employee"
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logger.Log.Errorf("Can't convert string to int: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		companyID, err := strconv.Atoi(r.Form.Get("company_id"))
		if err != nil {
			logger.Log.Errorf("Can't convert string to int: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		update := &presenter.Employee{
			ID:         int64(id),
			Name:       r.Form.Get("name"),
			SecondName: r.Form.Get("second_name"),
			Surname:    r.Form.Get("surname"),
			PhotoUrl:   r.Form.Get("photo_url"),
			HireDate:   r.Form.Get("hire_date"),
			Position:   r.Form.Get("position"),
			CompanyID:  int64(companyID),
		}
		empl, err := json.Marshal(update)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		byteUUID, err := e.kafka.KafkaSend(empl, EmployeePUTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := e.kafka.KafkaGet(EmployeePUTResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			respondWithError(w, errorMessage)
			logger.Log.Errorf("Updating failed: %v", err)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}
