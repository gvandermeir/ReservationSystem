package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
)

func (h *Handlers) SetAvailability(ctx context.Context, provider uuid.UUID, availabilities []*data.Availability) (error){
	var appointments []*data.Appointment
	for _, availability := range availabilities {
		start := availability.StartTime
        for current := start; !current.After(availability.EndTime); current = current.Add(15 * time.Minute) {
            appointments = append(appointments, data.NewAppointment(provider, uuid.MustParse("00000000-0000-0000-0000-000000000000"), current, current.Add(15 * time.Minute)))
        }
    }
    // Save appointments to database
    if len(appointments) > 0 {
        return h.dh.InsertAppointments(ctx, appointments)
    }
    return errors.New("No appointments to insert")
}
