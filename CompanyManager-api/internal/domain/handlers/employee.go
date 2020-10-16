package handlers

import (
	"encoding/json"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/kafka/consumers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/kafka/producers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"

	"github.com/gorilla/mux"

	"net/http"
)

//func CreateEmployee() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		var empl employee.Employee
//
//		err := parseJsonToStruct(w, r, &empl)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//
//		employeeProtocol := &employee.EmployeeProto{
//			Name:       empl.Name,
//			SecondName: empl.SecondName,
//			Surname:    empl.Surname,
//			PhotoUrl:   empl.PhotoUrl,
//			HireDate:   empl.HireDate,
//			Position:   empl.Position,
//			CompanyId:  empl.CompanyID,
//		}
//		newEmployeeIDProto, err := e.CreateEmployee(r.Context(), employeeProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//		Respond(w, newEmployeeIDProto)
//		return
//	}
//}

func GetEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var employee presenter.Employee
		producers.KafkaSendId(mux.Vars(r)["id"], "getEmployee", 0)
		msg := consumers.KafkaGetStruct("sendEmployee")
		employee = domain.JsonToEmployee(msg)
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
//func UpdateEmployee() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//
//		var empl employee.Employee
//
//		err := parseJsonToStruct(w, r, &empl)
//		if err != nil {
//			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//			return
//		}
//
//		employeeProtocol := &employee.EmployeeProto{
//			Id:         empl.ID,
//			Name:       empl.Name,
//			SecondName: empl.SecondName,
//			Surname:    empl.Surname,
//			PhotoUrl:   empl.PhotoUrl,
//			HireDate:   empl.HireDate,
//			Position:   empl.Position,
//			CompanyId:  empl.CompanyID,
//		}
//		_, err = e.UpdateEmployee(r.Context(), employeeProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "Employee not found")
//			return
//		}
//	}
//}
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
