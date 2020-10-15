package company

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


func (s *postgresRepo) GetCompany(id int64) (*Company, error) {
	var company Company

	rows, err := s.db.Query("SELECT * from company WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&company.ID, &company.Name, &company.Legalform);
			err != nil {
			return nil, err
		}
	}
	companyProto := ToProtoCompany(company)
	return &companyProto, nil
}

//CreateOrder CreateOrder
//func (s *Server) CreateCompany(ctx context.Context, in *CompanyProto) (*Id, error) {
//	var compId int64
//	err := s.Database.Db.QueryRow("INSERT INTO company(name, legal_form) VALUES ($1, $2) RETURNING company_id", in.Name, in.Legalform).Scan(compId)
//	if err != nil {
//		return nil, err
//	}
//	return &Id{Id: compId}, nil
//}
//
//func (s *Server) DeleteCompany(ctx context.Context, in *Id) (*CompanyReply, error) {
//	_, err := s.Database.Db.Exec("DELETE FROM company WHERE company_id = $1", in.Id)
//	if err != nil {
//		return nil, err
//	}
//	companyReply := CompanyReply{Message: "Successful deletion"}
//	return &companyReply, nil
//}
//
//func (s *Server) UpdateCompany(ctx context.Context, in *CompanyProto) (*CompanyReply, error) {
//	_, err := s.Database.Db.Exec("UPDATE company set name = $1, legal_form = $2 where company_id = $3;", in.Name, in.Legalform, in.Id)
//	if err != nil {
//		return nil, err
//	}
//	companyReply := CompanyReply{Message: "Successful update"}
//
//	return &companyReply, nil
//}
//
//func (s *Server) FormUpdateCompany(ctx context.Context, in *CompanyProto) (*CompanyReply, error) {
//	_, err := s.Database.Db.Exec("UPDATE company set name = $1, legal_form = $2 where company_id = $3;", in.Name, in.Legalform, in.Id)
//	if err != nil {
//		return nil, err
//	}
//	companyReply := CompanyReply{Message: "Successful update"}
//
//	return &companyReply, nil
//}
//
//func (s *Server) GetEmployeesByCompany(ctx context.Context, in *Id) (*Employees, error) {
//
//	rows, err := s.Database.Db.Query("SELECT * from employees WHERE company_id = $1", in.Id)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	employees := []employee.Employee{}
//
//	for rows.Next() {
//		employee := employee.Employee{}
//
//		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
//			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID); err != nil {
//			return nil, err
//		}
//		employees = append(employees, employee)
//	}
//	employeesProto := ToMultipleProtoEmployee(employees)
//	return &employeesProto, nil
//}