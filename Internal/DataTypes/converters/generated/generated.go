// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	converters "github.com/gvandermeir/ReservationSystem/Internal/DataTypes/converters"
	data "github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
	protos "github.com/gvandermeir/ReservationSystem/Internal/DataTypes/protos"
)

type AppointmentImpl struct{}

func (c *AppointmentImpl) FromProto(source *protos.Appointment) *data.Appointment {
	var pDataAppointment *data.Appointment
	if source != nil {
		var dataAppointment data.Appointment
		dataAppointment.ID = converters.BytesToUUID((*source).Id)
		dataAppointment.ProviderID = converters.BytesToUUID((*source).ProviderId)
		dataAppointment.PatientID = converters.BytesToUUID((*source).PatientId)
		dataAppointment.StartTime = converters.TimestamppbToTime((*source).StartTime)
		dataAppointment.EndTime = converters.TimestamppbToTime((*source).EndTime)
		dataAppointment.BookedAt = converters.TimestamppbToTime((*source).BookedAt)
		dataAppointment.Confirmed = (*source).Confirmed
		pDataAppointment = &dataAppointment
	}
	return pDataAppointment
}
func (c *AppointmentImpl) FromProtos(source []*protos.Appointment) []*data.Appointment {
	var pDataAppointmentList []*data.Appointment
	if source != nil {
		pDataAppointmentList = make([]*data.Appointment, len(source))
		for i := 0; i < len(source); i++ {
			pDataAppointmentList[i] = c.FromProto(source[i])
		}
	}
	return pDataAppointmentList
}
func (c *AppointmentImpl) ToProto(source *data.Appointment) *protos.Appointment {
	var pProtosAppointment *protos.Appointment
	if source != nil {
		var protosAppointment protos.Appointment
		protosAppointment.Id = converters.UUIDToBytes((*source).ID)
		protosAppointment.ProviderId = converters.UUIDToBytes((*source).ProviderID)
		protosAppointment.PatientId = converters.UUIDToBytes((*source).PatientID)
		protosAppointment.StartTime = converters.TimeToTimestamppb((*source).StartTime)
		protosAppointment.EndTime = converters.TimeToTimestamppb((*source).EndTime)
		protosAppointment.BookedAt = converters.TimeToTimestamppb((*source).BookedAt)
		protosAppointment.Confirmed = (*source).Confirmed
		pProtosAppointment = &protosAppointment
	}
	return pProtosAppointment
}
func (c *AppointmentImpl) ToProtos(source []*data.Appointment) []*protos.Appointment {
	var pProtosAppointmentList []*protos.Appointment
	if source != nil {
		pProtosAppointmentList = make([]*protos.Appointment, len(source))
		for i := 0; i < len(source); i++ {
			pProtosAppointmentList[i] = c.ToProto(source[i])
		}
	}
	return pProtosAppointmentList
}

type AvailabilityImpl struct{}

func (c *AvailabilityImpl) FromProto(source *protos.Availability) *data.Availability {
	var pDataAvailability *data.Availability
	if source != nil {
		var dataAvailability data.Availability
		dataAvailability.ProviderID = converters.BytesToUUID((*source).ProviderId)
		dataAvailability.StartTime = converters.TimestamppbToTime((*source).StartTime)
		dataAvailability.EndTime = converters.TimestamppbToTime((*source).EndTime)
		pDataAvailability = &dataAvailability
	}
	return pDataAvailability
}
func (c *AvailabilityImpl) FromProtos(source []*protos.Availability) []*data.Availability {
	var pDataAvailabilityList []*data.Availability
	if source != nil {
		pDataAvailabilityList = make([]*data.Availability, len(source))
		for i := 0; i < len(source); i++ {
			pDataAvailabilityList[i] = c.FromProto(source[i])
		}
	}
	return pDataAvailabilityList
}
func (c *AvailabilityImpl) ToProto(source *data.Availability) *protos.Availability {
	var pProtosAvailability *protos.Availability
	if source != nil {
		var protosAvailability protos.Availability
		protosAvailability.ProviderId = converters.UUIDToBytes((*source).ProviderID)
		protosAvailability.StartTime = converters.TimeToTimestamppb((*source).StartTime)
		protosAvailability.EndTime = converters.TimeToTimestamppb((*source).EndTime)
		pProtosAvailability = &protosAvailability
	}
	return pProtosAvailability
}
func (c *AvailabilityImpl) ToProtos(source []*data.Availability) []*protos.Availability {
	var pProtosAvailabilityList []*protos.Availability
	if source != nil {
		pProtosAvailabilityList = make([]*protos.Availability, len(source))
		for i := 0; i < len(source); i++ {
			pProtosAvailabilityList[i] = c.ToProto(source[i])
		}
	}
	return pProtosAvailabilityList
}

type PatientImpl struct{}

func (c *PatientImpl) FromProto(source *protos.Patient) *data.Patient {
	var pDataPatient *data.Patient
	if source != nil {
		var dataPatient data.Patient
		dataPatient.ID = converters.BytesToUUID((*source).Id)
		dataPatient.Email = (*source).Email
		dataPatient.Name = (*source).Name
		pDataPatient = &dataPatient
	}
	return pDataPatient
}
func (c *PatientImpl) FromProtos(source []*protos.Patient) []*data.Patient {
	var pDataPatientList []*data.Patient
	if source != nil {
		pDataPatientList = make([]*data.Patient, len(source))
		for i := 0; i < len(source); i++ {
			pDataPatientList[i] = c.FromProto(source[i])
		}
	}
	return pDataPatientList
}
func (c *PatientImpl) ToProto(source *data.Patient) *protos.Patient {
	var pProtosPatient *protos.Patient
	if source != nil {
		var protosPatient protos.Patient
		protosPatient.Id = converters.UUIDToBytes((*source).ID)
		protosPatient.Name = (*source).Name
		protosPatient.Email = (*source).Email
		pProtosPatient = &protosPatient
	}
	return pProtosPatient
}
func (c *PatientImpl) ToProtos(source []*data.Patient) []*protos.Patient {
	var pProtosPatientList []*protos.Patient
	if source != nil {
		pProtosPatientList = make([]*protos.Patient, len(source))
		for i := 0; i < len(source); i++ {
			pProtosPatientList[i] = c.ToProto(source[i])
		}
	}
	return pProtosPatientList
}

type ProviderImpl struct{}

func (c *ProviderImpl) FromProto(source *protos.Provider) *data.Provider {
	var pDataProvider *data.Provider
	if source != nil {
		var dataProvider data.Provider
		dataProvider.ID = converters.BytesToUUID((*source).Id)
		dataProvider.Email = (*source).Email
		dataProvider.Name = (*source).Name
		pDataProvider = &dataProvider
	}
	return pDataProvider
}
func (c *ProviderImpl) FromProtos(source []*protos.Provider) []*data.Provider {
	var pDataProviderList []*data.Provider
	if source != nil {
		pDataProviderList = make([]*data.Provider, len(source))
		for i := 0; i < len(source); i++ {
			pDataProviderList[i] = c.FromProto(source[i])
		}
	}
	return pDataProviderList
}
func (c *ProviderImpl) ToProto(source *data.Provider) *protos.Provider {
	var pProtosProvider *protos.Provider
	if source != nil {
		var protosProvider protos.Provider
		protosProvider.Id = converters.UUIDToBytes((*source).ID)
		protosProvider.Name = (*source).Name
		protosProvider.Email = (*source).Email
		pProtosProvider = &protosProvider
	}
	return pProtosProvider
}
func (c *ProviderImpl) ToProtos(source []*data.Provider) []*protos.Provider {
	var pProtosProviderList []*protos.Provider
	if source != nil {
		pProtosProviderList = make([]*protos.Provider, len(source))
		for i := 0; i < len(source); i++ {
			pProtosProviderList[i] = c.ToProto(source[i])
		}
	}
	return pProtosProviderList
}
