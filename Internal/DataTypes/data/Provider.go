package data

import (
	"time"
	"github.com/google/uuid"
)

type Provider struct {
	ID    uuid.UUID    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Availability struct {
	ProviderID uuid.UUID       `json:"provider_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time   `json:"end_time"`
}
