package presenter


type Employee struct {
	ID         int64     `json:"id"`
	Name       string `json:"name";validate:"name,required"`
	SecondName string `json:"second_name"`
	Surname    string `json:"surname"`
	PhotoUrl   string `json:"photo_url";validate:"photo_url,required"`
	HireDate   string `json:"hire_date"`
	Position   string `json:"position"`
	CompanyID  int64  `json:"company_id"`
}

