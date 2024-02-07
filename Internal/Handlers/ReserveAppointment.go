package handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
)

func (h *Handlers) ReserveAppointment(ctx context.Context, patient uuid.UUID, appointment *data.Appointment) error {
	appointment.PatientID = patient
	appointment.BookedAt = time.Now()
	// Create the appointment
	return h.dh.UpsertAppointment(ctx, appointment)
}