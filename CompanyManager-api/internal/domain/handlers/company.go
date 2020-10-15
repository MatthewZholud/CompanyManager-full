package handlers

//import (
//	"context"
//	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/domain/presenter"
//	"net/http"
//	"strconv"
//
//	"github.com/gorilla/mux"
//)

//func CreateCompany() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		var comp presenter.Company
//
//		err := parseJsonToStruct(w, r, &comp)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//
//		companyProtocol := &presenter.Company{
//			Name:      comp.Name,
//			Legalform: comp.Legalform,
//		}
//		newCompanyIDProto, err := c.CreateCompany(r.Context(), companyProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//			return
//		}
//		Respond(w, newCompanyIDProto)
//	}
//}

//func GetCompany() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//
//		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//		if err != nil {
//			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//			return
//		}
//
//		id64 := int64(id)
//
//
//		incomingCompany, err := c.GetCompany(context.Background(), &presenter.Id{Id: id64})
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "Employee not found")
//			return
//		}
//
//		Respond(w, presenter.Company{
//			ID:        incomingCompany.Id,
//			Name:      incomingCompany.Name,
//			Legalform: incomingCompany.Legalform,
//		})
//	}
//}

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
//func UpdateCompany() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		var comp presenter.Company
//
//		err := parseJsonToStruct(w, r, &comp)
//		if err != nil {
//			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//			return
//		}
//
//		companyProtocol := &presenter.Company{
//			Id:        comp.ID,
//			Name:      comp.Name,
//			Legalform: comp.Legalform,
//		}
//		_, err = c.UpdateCompany(r.Context(), companyProtocol)
//		if err != nil {
//			respondWithError(w, http.StatusNotFound, "Employee not found")
//			return
//		}
//	}
//}
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
