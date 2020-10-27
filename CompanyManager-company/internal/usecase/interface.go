package usecase

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
)

//type Reader interface {
//	Get(id int64) (*company.Company, error)
//}
//
//type Writer interface {
//	Create(e *company.Company) (string, error)
//	Update(e *company.Company) (string, error)
//	Delete(id int64) (string, error)
//}

//repository interface
type Repository interface {
	Get(id int64) (*entity.Company, error)
	GetAll() (*[]entity.Company, error)


	Create(e *entity.Company) (string, error)
	Update(e *entity.Company) (string, error)
	Delete(id int64) (string, error)
}

type UseCase interface {
	GetCompany(message []byte) ([]byte, error)
	GetAllCompany() ([]byte, error)

	//SearchBooks(query string) ([]*presenter.Book, error)
	//ListBooks() ([]*presenter.Book, error)
	CreateCompany(message []byte) ([]byte, error)
	UpdateCompany(message []byte) ([]byte, error)
	DeleteCompany(message []byte) ([]byte, error)
}


