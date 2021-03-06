package entity

//
//type ID = uuid.UUID
//
//func NewID() ID {
//	return ID(uuid.New())
//}
//
//func StringToID(s string) (ID, error) {
//	id, err := uuid.Parse(s)
//	return ID(id), err
//}

type Company struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Legalform string `json:"legal_form"`
}

//func NewCompany(name, legalForm string, id int64) *Company {
//	comp := &Company{
//		ID:        id,
//		Name:      name,
//		Legalform: legalForm,
//	}
//	return comp
//}

type Message struct {
	Key []byte
	Value []byte
}
