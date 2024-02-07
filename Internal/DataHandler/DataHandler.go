package datahandler

import (
	"context"
	"database/sql"
	"time"

	_ "embed"

	"github.com/google/uuid"
	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
)

type DataHandler struct {
	db *sql.DB
}

func NewDataHandler() (*DataHandler, error) {
	db, err := sql.Open("postgres", "postgres://user:xxxx@xxxxxxxxxxxx")
	return &DataHandler{db: db}, err
}

//go:embed Queries/insertAppointments.sql
var insertAppointmentsQuery string

func (d *DataHandler) InsertAppointments(ctx context.Context, appointments []*data.Appointment) error {
    _, err := d.db.Query(insertAppointmentsQuery, )
    if err != nil {
        return err
    }
	return nil
}

//go:embed Queries/getAppointments.sql
var getAppointmentsQuery string
func (d *DataHandler) GetAppointmentsByProviders(ctx context.Context,providerIDs []*uuid.UUID, start, end time.Time) ([]*data.Appointment, error){
    appointments, err := d.db.Query(getAppointmentsQuery, start, end, time.Now().Add(-30*time.Minute), providerIDs)
    if err != nil {
        return nil, err
    }
	convertedAppointments := []*data.Appointment{}

	for appointments.Next() {
		var appointment data.Appointment
		err = appointments.Scan(&appointment.ID, &appointment.ProviderID, &appointment.PatientID, &appointment.StartTime, &appointment.EndTime, &appointment.BookedAt, &appointment.Confirmed)
		if err != nil {
			return nil, err
		}
		convertedAppointments = append(convertedAppointments, &appointment)
	}

	return convertedAppointments, nil
}

//go:embed Queries/upsertAppointments.sql
var upsertAppointmentsQuery string
func (d *DataHandler) UpsertAppointment(ctx context.Context, appointment *data.Appointment) error{
	_, err := d.db.Query(upsertAppointmentsQuery, appointment.PatientID, appointment.BookedAt, appointment.Confirmed)
	if err != nil {
		return err
	}
	return nil
}

//go:embed Queries/getAppointmentsByID.sql
var getAppointmentsByIDQuery string
func (d *DataHandler) GetAppointmentByID(ctx context.Context, appointmentID uuid.UUID) (*data.Appointment, error){
	appointment, err := d.db.Query(getAppointmentsByIDQuery, appointmentID)
	if err != nil {
		return nil, err
	}
	var convertedAppointment data.Appointment
	err = appointment.Scan(&convertedAppointment.ID, &convertedAppointment.ProviderID, &convertedAppointment.PatientID, &convertedAppointment.StartTime, &convertedAppointment.EndTime, &convertedAppointment.BookedAt, &convertedAppointment.Confirmed)
	if err != nil {
		return nil, err
	}
	return &convertedAppointment, nil
}