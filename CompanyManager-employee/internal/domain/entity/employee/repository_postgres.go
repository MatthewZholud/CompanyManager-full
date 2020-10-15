package employee

import "database/sql"

type postgresRepo struct {
	db *sql.DB
}

//NewMySQLRepository create new repository
func NewPostgresRepository(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}

//func (s *postgresRepo) GetEmployee(id int64) (*Employee, error) {
//	var employee Employee
//	rows, err := s.db.Query("SELECT * from employees WHERE employee_id = $1", id)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
//			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID);
//			err != nil {
//			return nil, err
//		}
//	}
//
//	employeeProto := ToProtoEmployee(employee)
//	return &employeeProto, nil
//}

//func (s *postgresRepo) CreateEmployee(e *Employee) (ID, error) {
//	var empId int64
//	err := s.db.QueryRow("INSERT INTO employees(name, secondName, surname, photoUrl, hireDate, position, company_id) "+
//		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING employee_id", e.Name, e.SecondName, e.Surname, e.PhotoUrl, e.HireDate, e.Position, e.CompanyID).Scan(empId)
//	if err != nil {
//		return e.ID, err
//	}
//	return e.ID, nil
//}
//
//func (s *postgresRepo) DeleteEmployee(ctx context.Context, in *Id) (*EmployeeReply, error) {
//	_, err := s.db.Exec("DELETE FROM employees WHERE employee_id = $1", in.Id)
//	if err != nil {
//		return nil, err
//	}
//	employeeReply := EmployeeReply{Message: "Successful delete"}
//	return &employeeReply, nil
//}
//
//func (s *postgresRepo) UpdateEmployee(ctx context.Context, in *EmployeeProto) (*EmployeeReply, error) {
//	_, err := s.db.Exec("UPDATE employees set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5,"+
//		" position = $6, company_id = $7 where employee_id = $7;", in.Name, in.SecondName, in.Surname,
//		in.PhotoUrl, in.HireDate, in.Position, in.CompanyId, in.Id)
//	if err != nil {
//		return nil, err
//	}
//	employeeReply := EmployeeReply{Message: "Successful update"}
//	return &employeeReply, nil
//}
//
//func (s *postgresRepo) FormUpdateEmployee(ctx context.Context, in *EmployeeProto) (*EmployeeReply, error) {
//	_, err := s.db.Exec("UPDATE employees set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5,"+
//		" position = $6, company_id = $7 where employee_id = $7;", in.Name, in.SecondName, in.Surname,
//		in.PhotoUrl, in.HireDate, in.Position, in.CompanyId, in.Id)
//	if err != nil {
//		return nil, err
//	}
//	employeeReply := EmployeeReply{Message: "Successful update"}
//
//	return &employeeReply, nil
//}