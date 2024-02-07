package handlers

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (h *Handlers) ConfirmAppointment(ctx context.Context, patient uuid.UUID, appointmentID uuid.UUID) error {
	appointment, err := h.dh.GetAppointmentByID(ctx, appointmentID)
	if err != nil {
		return err
	}
	if appointment.PatientID != patient {
		return errors.New("Appointment expired and is no longer available")
	}

	appointment.Confirmed = true

	// Create the appointment
	return h.dh.UpsertAppointment(ctx, appointment)
}