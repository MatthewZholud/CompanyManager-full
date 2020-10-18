package repository

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
)

type postgresRepo struct {
	db *sql.DB
}

//NewMySQLRepository create new repository
func NewPostgresRepository(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}


func (s *postgresRepo) Get(id int64) (*company.Company, error) {
	var company company.Company

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
	return &company, nil
}

func (s *postgresRepo) Create(c *company.Company) (string, error) {
	var compId string
	err := s.db.QueryRow("INSERT INTO company(name, legal_form) VALUES ($1, $2) RETURNING company_id", c.Name, c.Legalform).Scan(compId)
	if err != nil {
		return compId, err
	}
	return compId, nil
}
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