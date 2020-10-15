package company

import "github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/domain/entity/company"

type Reader interface {
	GetCompany(id int64) (*company.Company, error)
	Search(query string) ([]*company.Company, error)
	List() ([]*company.Company, error)
}

//Writer book writer
type Writer interface {
	Create(e *company.Company) (int64, error)
	Update(e *company.Company) error
	Delete(id int64) error
}

//repository interface
type Repository interface {
	Reader
	Writer
}

type CompanyRepository struct {
	repo Repository
}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

