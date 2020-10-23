package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"strconv"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error adding company"
		var input struct {
			Name      string `json:"name"`
			Legalform string `json:"legal_form"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Fatalf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		comp, err := json.Marshal(input)
		if err != nil {
			logger.Log.Fatalf("Can't prepare company struct for sending to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		err = producers.KafkaSend(comp, "CompanyPOSTRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("CompanyPOSTResponse")
		if err != nil {
			respondWithError(w, errorMessage)
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			return
		}
		id, err  := ByteToInt64(msg)
		if err != nil {
			respondWithError(w, errorMessage)
			logger.Log.Fatalf("Can't convert byte to int: %v", err)
			return
		}
		toJ := &presenter.Company{
			ID:        id,
			Name:      input.Name,
			Legalform: input.Legalform,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			logger.Log.Fatalf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Create handler completed successfully")
		}
	}
}

func GetCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading company"
		var company *presenter.Company

		err := producers.KafkaSend([]byte(mux.Vars(r)["companyId"]), "CompanyGETRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("CompanyGETResponse")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		company, err = JsonToCompany(msg)
		if err != nil {
			logger.Log.Fatalf("Can't convert json to company struct: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if err := json.NewEncoder(w).Encode(company); err != nil {
			logger.Log.Fatalf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Get handler completed successfully")
		}
	}
}

func DeleteCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error deleting company"
		err := producers.KafkaSend([]byte(mux.Vars(r)["companyId"]), "CompanyDeleteRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("CompanyDeleteResponse")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful delete" {
			logger.Log.Fatalf("Deleting failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful delete")
		}
	}
}

func UpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating company"
		var update struct {
			Name      string `json:"name"`
			Legalform string `json:"legal_form"`
		}

		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			logger.Log.Fatalf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		comp, err := json.Marshal(update)
		if err != nil {
			logger.Log.Fatalf("Can't prepare company struct for sending to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		err = producers.KafkaSend(comp, "CompanyPUTRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("CompanyPUTResponse")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			logger.Log.Fatalf("Updating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}

func FormUpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating company"

		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			logger.Log.Fatalf("Can't convert string to int: %v", err)
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
			logger.Log.Fatalf("Can't prepare company struct for sending to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		err = producers.KafkaSend(comp, "CompanyPUTRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("CompanyPUTResponse")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		if string(msg) != "Successful update" {
			logger.Log.Fatalf("Updating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update")
		}
	}
}


func GetEmployeesByCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading employees"

		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		var employee []presenter.Employee
		err := producers.KafkaSend([]byte(mux.Vars(r)["companyId"]), "EmployeeByCompanyGETRequest")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		msg, err := consumers.KafkaGetStruct("EmployeeByCompanyGETResponse")
		if err != nil {
			logger.Log.Fatalf("Error sending message to kafka: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		employee, err = JsonToEmployeeArr(msg)
		if err != nil {
			logger.Log.Fatalf("Can't convert json to employee array: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		if err := json.NewEncoder(w).Encode(employee); err != nil {
			logger.Log.Fatalf("Can't display: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Get(employees by company) handler completed successfully")
		}
	}
}
