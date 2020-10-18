package handlers

import (
	"encoding/json"
	"log"

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

		errorMessage := "Error adding book"
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

//func DeleteEmployee() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		id, err := strconv.Atoi(mux.Vars(r)["id"])
//		if err != nil {
//			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//			return
//		}
//		id64 := int64(id)
//
//		_, err = e.DeleteEmployee(context.Background(), &employee.Id{Id: id64})
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "Company not found")
//			return
//		}
//	}
//}
//


func UpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")


		errorMessage := "Error updating book"
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
		id := ByteToInt64(msg)
		toJ := &presenter.Employee{
			ID:         id,
			Name:       update.Name,
			SecondName: update.SecondName,
			Surname:    update.Surname,
			PhotoUrl:   update.PhotoUrl,
			HireDate:   update.HireDate,
			Position:   update.Position,
			CompanyID:  update.CompanyID,
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}
//
//func FormUpdateEmployee() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		var empl employee.Employee
//
//		id, err := strconv.Atoi(mux.Vars(r)["id"])
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//		id64 := int64(id)
//
//		err = parseJsonToStruct(w, r, &empl)
//		if err != nil {
//			return
//		}
//
//		employeeProtocol := &employee.EmployeeProto{
//			Id:         id64,
//			Name:       empl.Name,
//			SecondName: empl.SecondName,
//			Surname:    empl.Surname,
//			PhotoUrl:   empl.PhotoUrl,
//			HireDate:   empl.HireDate,
//			Position:   empl.Position,
//			CompanyId:  empl.CompanyID,
//		}
//		_, err = e.FormUpdateEmployee(r.Context(), employeeProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//	}
//}
