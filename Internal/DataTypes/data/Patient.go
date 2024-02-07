package data

import (
	"github.com/google/uuid"
)

type Patient struct {
	ID    uuid.UUID    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
