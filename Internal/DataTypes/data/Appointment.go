package data

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID         uuid.UUID       `json:"id"`
	ProviderID uuid.UUID       `json:"provider_id"`
	PatientID  uuid.UUID       `json:"patient_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time    `json:"end_time"`
	BookedAt   time.Time `json:"booked_at"`
	Confirmed  bool      `json:"confirmed"`
}

func NewAppointment(providerID, patientID uuid.UUID, startTime, endTime time.Time) *Appointment {
	return &Appointment{ID:uuid.New(), ProviderID: providerID, PatientID: patientID, StartTime: startTime, EndTime: endTime, Confirmed: false}
}