package presenter

//import (
//	"github.com/google/uuid"
//)
//
//type ID = uuid.UUID

type Company struct {
	ID        int64     `json:"id"`
	Name      string `json:"name"`
	Legalform string `json:"legal_form"`
}

