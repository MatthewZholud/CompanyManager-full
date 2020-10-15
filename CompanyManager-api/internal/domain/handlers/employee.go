package handlers

import (
	"context"
	//"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	kafka "github.com/segmentio/kafka-go"
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
		id := mux.Vars(r)["id"]


		topic := "my-topic"
		partition := 0

		conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(id)},

		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
		//incomingEmployee, err := e.GetEmployee(context.Background(), &employee.Id{Id: id64})
		//if err != nil {
		//	respondWithError(w, http.StatusNotFound, "Employee not found")
		//	return
		//}
		//
		//Respond(w, presenter.Employee{
		//	ID:         incomingEmployee.Id,
		//	Name:       incomingEmployee.Name,
		//	SecondName: incomingEmployee.SecondName,
		//	Surname:    incomingEmployee.Surname,
		//	PhotoUrl:   incomingEmployee.PhotoUrl,
		//	HireDate:   incomingEmployee.HireDate,
		//	Position:   incomingEmployee.Position,
		//	CompanyID:  incomingEmployee.CompanyId,
		//})


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
