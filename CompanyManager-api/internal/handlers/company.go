package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/kafka/producers"
	"log"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error adding book"
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

//func DeleteCompany() http.HandlerFunc {
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
//		_, err = c.DeleteCompany(context.Background(), &presenter.Id{Id: id64})
//
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "Company not found")
//			return
//		}
//
//	}
//}
//
func UpdateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding book"
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
		producers.KafkaSendStruct(comp, "CompanyPUTRequest")
		msg := consumers.KafkaGetStruct("CompanyPUTResponse")
		id := ByteToInt64(msg)
		toJ := &presenter.Company{
			ID:        id,
			Name:      input.Name,
			Legalform: input.Legalform,
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	}
}
//
//func FormUpdateCompany() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		var comp presenter.Company
//
//		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//
//		id64 := int64(id)
//
//		err = parseJsonToStruct(w, r, &comp)
//		if err != nil {
//			return
//		}
//
//		companyProtocol := &presenter.Company{
//			ID:        id64,
//			Name:      comp.Name,
//			Legalform: comp.Legalform,
//		}
//		_, err = c.FormUpdateCompany(r.Context(), companyProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//	}
//}
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
