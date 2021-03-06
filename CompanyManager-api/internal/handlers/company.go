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
	CompanyGETRequest     = "CompanyGETRequest"
	CompanyPOSTRequest    = "CompanyPOSTRequest"
	CompanyPUTRequest     = "CompanyPUTRequest"
	CompanyDeleteRequest  = "CompanyDeleteRequest"
	CompanyGETResponse    = "CompanyGETResponse"
	CompanyPOSTResponse   = "CompanyPOSTResponse"
	CompanyPUTResponse    = "CompanyPUTResponse"
	CompanyDeleteResponse = "CompanyDeleteResponse"
	EmployeeByCompanyGETResponse = "EmployeeByCompanyGETResponse"
	EmployeeByCompanyGETRequest = "EmployeeByCompanyGETRequest"
)

type companyService struct {
	kafka kafka.KafkaRep
}

func InitializeCompany(kafka kafka.KafkaRep) *companyService {
	return &companyService{
		kafka: kafka,
	}
}

func (c *companyService) CreateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error adding company"
		var input struct {
			Name      string `json:"name"`
			Legalform string `json:"legal_form"`
		}

		//todo: mzh: what HTTP response would be returned in case of error?

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		comp, err := json.Marshal(input)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		//todo: mzh: Clean architecture + SRP violation. HTTP handler handles broker's request-reply logic
		byteUUID, err := c.kafka.KafkaSend(comp, CompanyPOSTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(CompanyPOSTResponse, byteUUID)
		if err != nil {
			respondWithError(w, errorMessage)
			logger.Log.Errorf("Error sending message to env: %v", err)
			return
		}
		// todo: mzh: why parse kafka message here? SRP violation
		id, err  := ByteToInt64(msg)
		if err != nil {
			respondWithError(w, errorMessage)
			logger.Log.Errorf("Can't convert byte to int: %v", err)
			return
		}
		//todo: mzh: strange naming of var
		toJ := &presenter.Company{
			ID:        id,
			Name:      input.Name,
			Legalform: input.Legalform,
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

func (c *companyService) GetCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading company"
		var company *presenter.Company

		byteUUID, err := c.kafka.KafkaSend([]byte(mux.Vars(r)["companyId"]), CompanyGETRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(CompanyGETResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		company, err = JsonToCompany(msg)
		if err != nil {
			logger.Log.Errorf("Can't convert json to company struct: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if err := json.NewEncoder(w).Encode(company); err != nil {
			logger.Log.Errorf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Get handler completed successfully")
		}
	}
}

func (c *companyService) DeleteCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error deleting company"
		byteUUID, err := c.kafka.KafkaSend([]byte(mux.Vars(r)["companyId"]), CompanyDeleteRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(CompanyDeleteResponse, byteUUID)
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

func (c *companyService) UpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating company"
		var update struct {
			Name      string `json:"name"`
			Legalform string `json:"legal_form"`
		}

		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		comp, err := json.Marshal(update)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		byteUUID, err := c.kafka.KafkaSend(comp, CompanyPUTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(CompanyPUTResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			logger.Log.Errorf("Updating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}

func (c *companyService) FormUpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating company"

		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			logger.Log.Errorf("Can't convert string to int: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		update := presenter.Company{
			ID:        int64(id),
			Name:      r.Form.Get("name"),
			Legalform: r.Form.Get("legal_form"),
		}

		comp, err := json.Marshal(update)
		if err != nil {
			logger.Log.Errorf("Can't prepare company struct for sending to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		byteUUID, err := c.kafka.KafkaSend(comp, CompanyPUTRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(CompanyPUTResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			logger.Log.Errorf("Updating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}


func (c *companyService) GetEmployeesByCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading employees"

		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		var employee []presenter.Employee
		byteUUID, err := c.kafka.KafkaSend([]byte(mux.Vars(r)["companyId"]), EmployeeByCompanyGETRequest)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := c.kafka.KafkaGet(EmployeeByCompanyGETResponse, byteUUID)
		if err != nil {
			logger.Log.Errorf("Error sending message to env: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		employee, err = JsonToEmployeeArr(msg)
		if err != nil {
			logger.Log.Errorf("Can't convert json to employee array: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		if err := json.NewEncoder(w).Encode(employee); err != nil {
			logger.Log.Errorf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Get(employees by company) handler completed successfully")
		}
	}
}
