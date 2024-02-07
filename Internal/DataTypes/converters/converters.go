package converters

import (
	"time"

	"github.com/google/uuid"
	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/data"
	"github.com/gvandermeir/ReservationSystem/Internal/DataTypes/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:extend BytesToUUID
// goverter:extend UUIDToBytes
// goverter:matchIgnoreCase
type Provider interface {
	FromProto(*protos.Provider) *data.Provider
	// goverter:ignore state sizeCache unknownFields
	ToProto(*data.Provider) *protos.Provider

	FromProtos([]*protos.Provider) []*data.Provider
	ToProtos([]*data.Provider) []*protos.Provider
}

// goverter:converter
// goverter:extend BytesToUUID
// goverter:extend UUIDToBytes
// goverter:extend TimestamppbToTime
// goverter:extend TimeToTimestamppb
// goverter:matchIgnoreCase
type Appointment interface {
	FromProto(*protos.Appointment) *data.Appointment
	// goverter:ignore state sizeCache unknownFields
	ToProto(*data.Appointment) *protos.Appointment

	FromProtos([]*protos.Appointment) []*data.Appointment
	ToProtos([]*data.Appointment) []*protos.Appointment
}

// goverter:converter
// goverter:extend BytesToUUID
// goverter:extend UUIDToBytes
// goverter:matchIgnoreCase
type Patient interface {
	FromProto(*protos.Patient) *data.Patient
	// goverter:ignore state sizeCache unknownFields
	ToProto(*data.Patient) *protos.Patient

	FromProtos([]*protos.Patient) []*data.Patient
	ToProtos([]*data.Patient) []*protos.Patient
}

// goverter:converter
// goverter:extend BytesToUUID
// goverter:extend UUIDToBytes
// goverter:extend TimestamppbToTime
// goverter:extend TimeToTimestamppb
// goverter:matchIgnoreCase
type Availability interface {
	FromProto(*protos.Availability) *data.Availability
	// goverter:ignore state sizeCache unknownFields
	ToProto(*data.Availability) *protos.Availability

	FromProtos([]*protos.Availability) []*data.Availability
	ToProtos([]*data.Availability) []*protos.Availability
}

func BytesToUUID(b []byte) uuid.UUID {
	return uuid.UUID(b)
}

func UUIDToBytes(u uuid.UUID) []byte {
	return u[:]
}

func TimestamppbToTime(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}

func TimeToTimestamppb(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}