package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"strconv"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"

	"github.com/gorilla/mux"

	"net/http"
)

type employeeService struct {
	interServiceEmpl InterServiceEmployee
}

func InitializeEmployee(interService InterServiceEmployee) *employeeService {
	return &employeeService{
		interServiceEmpl: interService,
	}
}

func (e *employeeService) CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error adding employee"
		var input presenter.Employee

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Errorf("Can't get employee struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		id, err := e.interServiceEmpl.CreateEmployee(&input)
		if err != nil {
			logger.Log.Errorf("Creating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		}
		createdEmployee := &presenter.Employee{
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
		if err := json.NewEncoder(w).Encode(createdEmployee); err != nil {
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
		if IsNumericAndPositive(mux.Vars(r)["id"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		id := mux.Vars(r)["id"]
		employee, err := e.interServiceEmpl.GetEmployee(id)
		if err != nil {
			logger.Log.Errorf("Error getting employee: %v", err)
			respondWithError(w, errorMessage)
		}

		if err := json.NewEncoder(w).Encode(employee); err != nil {
			logger.Log.Errorf("Can't display employee: %v", err)
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
		if IsNumericAndPositive(mux.Vars(r)["id"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		id := mux.Vars(r)["id"]

		err := e.interServiceEmpl.DeleteEmployee(id)
		if err != nil {
			logger.Log.Errorf("Deleting of employee failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful delete of employee")
		}
	}
}

func (e *employeeService) UpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error updating employee"
		var update presenter.Employee

		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		err = e.interServiceEmpl.UpdateEmployee(&update)

		if err != nil {
			logger.Log.Errorf("Updating of employee failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update of employee")
		}
	}
}

func (e *employeeService) FormUpdateEmployee() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error updating employee"
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}

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

		update := presenter.Employee{
			ID:         int64(id),
			Name:       r.Form.Get("name"),
			SecondName: r.Form.Get("second_name"),
			Surname:    r.Form.Get("surname"),
			PhotoUrl:   r.Form.Get("photo_url"),
			HireDate:   r.Form.Get("hire_date"),
			Position:   r.Form.Get("position"),
			CompanyID:  int64(companyID),
		}

		err = e.interServiceEmpl.UpdateEmployee(&update)

		if err != nil {
			logger.Log.Errorf("Updating of employee failed: %v", err)
			respondWithError(w, errorMessage)
			return
		} else {
			logger.Log.Infof("Successful update of employee")
		}
	}
}
