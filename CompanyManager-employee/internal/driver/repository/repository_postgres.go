package repository

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity/employee"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}

func (s *postgresRepo) Get(id int64) (*employee.Employee, error) {
	var employee employee.Employee

	rows, err1 := s.db.Query("SELECT * from employees WHERE employee_id = $1", id)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID);
			err != nil {
			return nil, err
		}
	}

	return &employee, nil
}

func (s *postgresRepo) Create(e *employee.Employee) (string, error) {
	var empId string
	err := s.db.QueryRow("INSERT INTO employees(name, secondName, surname, photoUrl, hireDate, position, company_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING employee_id", e.Name, e.SecondName, e.Surname, e.PhotoUrl, e.HireDate, e.Position, e.CompanyID).Scan(empId)
	if err != nil {
		return empId, err
	}
	return empId, nil
}

func (s *postgresRepo) Delete(id int64) (string, error) {
	_, err := s.db.Exec("DELETE FROM employees WHERE employee_id = $1", id)
	if err != nil {
		return "", err
	}
	employeeReply := "Successful delete"
	return employeeReply, nil
}

func (s *postgresRepo) Update(e *employee.Employee) (string, error) {
	_, err := s.db.Exec("UPDATE employees set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5,"+
		" position = $6, company_id = $7 where employee_id = $7;", e.Name, e.SecondName, e.Surname,
		e.PhotoUrl, e.HireDate, e.Position, e.CompanyID, e.ID)
	if err != nil {
		return "", err
	}
	employeeReply := "Successful update"
	return employeeReply, nil
}

func (s *postgresRepo) GetEmployeesByCompany(id int64) (*[]employee.Employee, error) {

	rows, err := s.db.Query("SELECT * from employees WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees := []employee.Employee{}

	for rows.Next() {
		employee := employee.Employee{}

		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}
