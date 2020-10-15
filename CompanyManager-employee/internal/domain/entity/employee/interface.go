package employee

type Reader interface {
	GetEmployee(id int64) (*Employee, error)
	//Search(query string) ([]*employee.Employee, error)
	//List() ([]*employee.Employee, error)
}


type Writer interface {
	//Create(e *employee.Employee) (int64, error)
	//Update(e *employee.Employee) error
	//Delete(id int64) error
}

//repository interface
type Worker interface {
	Reader
	Writer
}

type EmployeeRepository struct {
	repo Worker
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{}
}
