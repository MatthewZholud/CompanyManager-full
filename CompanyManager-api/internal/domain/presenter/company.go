package presenter

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

type Company struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Legalform string `json:"legal_form"`
}

