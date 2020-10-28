package repository

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}

func (s *postgresRepo) Get(id int64) (*entity.Employee, error) {
	var employee entity.Employee

	rows, err := s.db.Query("SELECT * from employees WHERE employee_id = $1", id)
	if err != nil {
		logger.Log.Debug("Get query to Db was failed")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID)
	}
	return &employee, nil
}

func (s *postgresRepo) GetAll() (*[]entity.Employee, error) {
	rows, err := s.db.Query("SELECT * from employees")
	if err != nil {
		logger.Log.Debug("Get all query to Db was failed")
		return nil, err
	}
	defer rows.Close()
	employees := []entity.Employee{}

	for rows.Next() {
		employee := entity.Employee{}
		rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID)
		employees = append(employees, employee)
	}
	return &employees, nil
}

func (s *postgresRepo) Create(e *entity.Employee) (string, error) {
	var empId string
	err := s.db.QueryRow("INSERT INTO employees(name, secondName, surname, photoUrl, hireDate, position, company_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING employee_id", e.Name, e.SecondName, e.Surname, e.PhotoUrl, e.HireDate, e.Position, e.CompanyID).Scan(empId)
	if err != nil {
		logger.Log.Debug("Create query to Db was failed")
		return empId, err
	}
	return empId, nil
}

func (s *postgresRepo) Delete(id int64) (string, error) {
	_, err := s.db.Exec("DELETE FROM employees WHERE employee_id = $1", id)
	if err != nil {
		logger.Log.Debug("Delete query to Db was failed")
		return "", err
	}
	employeeReply := "Successful delete"
	return employeeReply, nil
}

func (s *postgresRepo) Update(e *entity.Employee) (string, error) {
	_, err := s.db.Exec("UPDATE employees set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5,"+
		" position = $6, company_id = $7 where employee_id = $8;", e.Name, e.SecondName, e.Surname,
		e.PhotoUrl, e.HireDate, e.Position, e.CompanyID, e.ID)
	if err != nil {
		logger.Log.Debug("Update query to Db was failed")
		return  "", err
	}
	employeeReply := "Successful update"
	return employeeReply, nil
}

func (s *postgresRepo) GetEmployeesByCompany(id int64) (*[]entity.Employee, error) {

	rows, err := s.db.Query("SELECT * from employees WHERE company_id = $1", id)
	if err != nil {
		logger.Log.Debug("Get(employee by company) query to Db was failed")
		return  nil, err
	}
	defer rows.Close()
	employees := []entity.Employee{}

	for rows.Next() {
		employee := entity.Employee{}
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}
