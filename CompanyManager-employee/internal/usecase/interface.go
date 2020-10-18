package usecase

import "github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/entity/employee"

type Reader interface {
	Get(id int64) (*employee.Employee, error)
	//Search(query string) ([]*employee.Employee, error)
	//List() ([]*employee.Employee, error)
}


type Writer interface {
	Create(e *employee.Employee) (string, error)
	//Update(e *employee.Employee) error
	//Delete(id int64) error
}

//repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetEmployee(message string)
	//SearchBooks(query string) ([]*entity.Book, error)
	//ListBooks() ([]*entity.Book, error)
	CreateEmployee(message []byte)
	//UpdateBook(e *entity.Book) error
	//DeleteBook(id entity.ID) error
}