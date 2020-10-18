package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/producers"
	"log"
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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		comp, err := json.Marshal(input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(comp, "CompanyPOSTRequest")
		msg := consumers.KafkaGetStruct("CompanyPOSTResponse")
		id := ByteToInt64(msg)
		toJ := &presenter.Company{
			ID:        id,
			Name:      input.Name,
			Legalform: input.Legalform,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func GetCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var company presenter.Company
		producers.KafkaSendId(mux.Vars(r)["companyId"], "CompanyGETRequest")
		msg := consumers.KafkaGetStruct("CompanyGETResponse")
		company = JsonToCompany(msg)
		json.NewEncoder(w).Encode(company)
	}
}

func DeleteCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error deleting company"
		producers.KafkaSendId(mux.Vars(r)["companyId"], "CompanyDeleteRequest")
		msg := consumers.KafkaGetStruct("CompanyDeleteResponse")
		if string(msg) != "Successful delete" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		comp, err := json.Marshal(update)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(comp, "CompanyPUTRequest")
		msg := consumers.KafkaGetStruct("CompanyPUTResponse")
		if string(msg) != "Successful update" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func FormUpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating company"

		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}

		update := presenter.Company{
			ID:        int64(id),
			Name:      r.Form.Get("name"),
			Legalform: r.Form.Get("legal_form"),
		}

		comp, err := json.Marshal(update)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		producers.KafkaSendStruct(comp, "CompanyPUTRequest")
		msg := consumers.KafkaGetStruct("CompanyPUTResponse")
		if string(msg) != "Successful update" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

//
//func GetEmployeesByCompany() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//		if err != nil {
//			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//			return
//		}
//
//		id64 := int64(id)
//
//		employeesProto, err := c.GetEmployeesByCompany(context.Background(), &presenter.Id{Id: id64})
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "ICompany not found")
//			return
//		}
//		var employees []presenter.Employee
//		for _, empl := range employeesProto.Employees {
//			employees = append(employees, presenter.Employee{
//				ID:         empl.Id,
//				Name:       empl.Name,
//				SecondName: empl.SecondName,
//				Surname:    empl.Surname,
//				PhotoUrl:   empl.PhotoUrl,
//				HireDate:   empl.HireDate,
//				Position:   empl.Position,
//				CompanyID:  empl.CompanyId,
//			})
//		}
//
//		Respond(w, employees)
//	}
//}
