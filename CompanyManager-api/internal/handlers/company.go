package handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/logger"
	"strconv"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-api/internal/presenter"
	"github.com/gorilla/mux"
	"net/http"
)


type companyService struct {
	interServiceComp InterServiceCompany
}

func InitializeCompany(interService InterServiceCompany) *companyService {
	return &companyService{
		interServiceComp: interService,
	}
}

func (c *companyService) CreateCompany() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error adding company"
		var input presenter.Company

		//todo: mzh: what HTTP response would be returned in case of error?

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		id, err := c.interServiceComp.CreateCompany(&input)
		if err != nil {
			logger.Log.Errorf("Creating failed: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		createdCompany := &presenter.Company{
			ID:        id,
			Name:      input.Name,
			Legalform: input.Legalform,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(createdCompany); err != nil {
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
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		id := mux.Vars(r)["companyId"]
		company, err := c.interServiceComp.GetCompany(id)
		if err != nil {
			logger.Log.Errorf("Error getting company: %v", err)
			respondWithError(w, errorMessage)
		}
		if err := json.NewEncoder(w).Encode(company); err != nil {
			logger.Log.Errorf("Can't display company: %v", err)
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
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
		id := mux.Vars(r)["companyId"]
		err := c.interServiceComp.DeleteCompany(id)
		if err != nil {
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
		var update presenter.Company
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			logger.Log.Errorf("Can't get company struct from body: %v", err)
			respondWithError(w, errorMessage)
			return
		}

		err = c.interServiceComp.UpdateCompany(&update)
		if err != nil {
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
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			logger.Log.Errorf("Id is not numeric and positive: %v")
			respondWithError(w, errorMessage)
			return
		}
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

		err = c.interServiceComp.UpdateCompany(&update)

		if err != nil {
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

		id := mux.Vars(r)["companyId"]

		employee, err := c.interServiceComp.GetEmployeesByCompany(id)
		if err != nil {
			logger.Log.Errorf("Error getting employees by company: %v", err)
			respondWithError(w, errorMessage)
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
