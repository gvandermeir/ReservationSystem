package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
)

func (h *Handlers) GetAppointments(ctx context.Context, providers []*uuid.UUID, start time.Time, end time.Time) ([]*data.Appointment, error){
	minTime := time.Now().Add(time.Hour * 24)
	if start.After(minTime) {
		start = minTime
		if end.Before(start) {
			return nil, errors.New("Appointments must be at least 24 hours in the future")
		}
	}
    return h.dh.GetAppointmentsByProviders(ctx, providers, start, end)
}