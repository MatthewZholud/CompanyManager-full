package repository

import (
	"database/sql"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
)

type postgresRepo struct {
	db *sql.DB
}

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
		rows.Scan(&company.ID, &company.Name, &company.Legalform)
	}
	return &company, nil
}

func (s *postgresRepo) Create(c *company.Company) (string, error) {
	var compId string
	err := s.db.QueryRow("INSERT INTO company(name, legal_form) VALUES ($1, $2) RETURNING company_id", c.Name, c.Legalform).Scan(&compId)
	if err != nil {
		return compId, err
	}
	return compId, nil
}

func (s *postgresRepo) Delete(id int64) (string, error) {
	_, err := s.db.Exec("DELETE FROM company WHERE company_id = $1", id)
	if err != nil {
		return "", err
	}
	companyReply :=  "Successful delete"
	return companyReply, nil
}

func (s *postgresRepo) Update(c *company.Company) (string, error) {
	_, err := s.db.Exec("UPDATE company set name = $1, legal_form = $2 where company_id = $3;", c.Name, c.Legalform, c.ID)
	if err != nil {
		return  "", err
	}
	companyReply := "Successful update"
	return companyReply, nil
}

