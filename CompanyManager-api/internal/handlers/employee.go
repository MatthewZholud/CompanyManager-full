package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"

	"github.com/gorilla/mux"

	"net/http"
)

func CreateEmployee() http.HandlerFunc {
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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		empl, err := json.Marshal(input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(empl, "EmployeePOSTRequest")
		msg := consumers.KafkaGetStruct("EmployeePOSTResponse")
		id := ByteToInt64(msg)
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
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func GetEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var employee presenter.Employee
		producers.KafkaSendId(mux.Vars(r)["id"], "EmployeeGETRequest")
		msg := consumers.KafkaGetStruct("EmployeeGETResponse")
		employee = JsonToEmployee(msg)
		json.NewEncoder(w).Encode(employee)
	}
}

func DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error deleting employee"
		producers.KafkaSendId(mux.Vars(r)["id"], "EmployeeDeleteRequest")
		msg := consumers.KafkaGetStruct("EmployeeDeleteResponse")
		if string(msg) != "Successful delete" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}


func UpdateEmployee() http.HandlerFunc {

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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		empl, err := json.Marshal(update)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(empl, "EmployeePUTRequest")
		msg := consumers.KafkaGetStruct("EmployeePUTResponse")
		if string(msg) != "Successful update" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func FormUpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating employee"

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}
		companyID, err := strconv.Atoi(r.Form.Get("company_id"))
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(empl, "EmployeePUTRequest")
		msg := consumers.KafkaGetStruct("EmployeePUTResponse")
		if string(msg) != "Successful update" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}
