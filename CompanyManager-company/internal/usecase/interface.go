package usecase

import "github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"

type Reader interface {
	Get(id int64) (*company.Company, error)
	//Search(query string) ([]*company.Company, error)
	//List() ([]*company.Company, error)
}

//Writer book writer
type Writer interface {
	Create(e *company.Company) (string, error)
	//Update(e *company.Company) error
	//Delete(id int64) error
}

//repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetCompany(message string)
	//SearchBooks(query string) ([]*entity.Book, error)
	//ListBooks() ([]*entity.Book, error)
	CreateCompany(message []byte)
	//UpdateBook(e *entity.Book) error
	//DeleteBook(id entity.ID) error
}