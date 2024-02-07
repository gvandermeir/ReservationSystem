package server

import (
	"context"

	"github.com/google/uuid"
	converters "github.com/gvandermeir/ReservationSystem/Internal/DataTypes/converters/generated"
	pb "github.com/gvandermeir/ReservationSystem/Internal/DataTypes/protos"
	handlers "github.com/gvandermeir/ReservationSystem/Internal/Handlers"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedReservationServiceServer
	h *handlers.Handlers
}

// NewServer returns a new server
func NewGRPCServer() *grpc.Server {
	gsrv := grpc.NewServer()
	handler, err := handlers.NewHandlers()
	if err != nil {
		return nil
	}
	s := server{h: handler}
	pb.RegisterReservationServiceServer(gsrv, &s)
	return gsrv
}

func (s *server) SetAvailability(ctx context.Context, in *pb.SetAvailabilityRequest) (*pb.SetAvailabilityResponse, error) {
	converter := converters.AvailabilityImpl{}
	availabilities := converter.FromProtos(in.Availabilities)
	if len(availabilities) == 0 {
		return &pb.SetAvailabilityResponse{Success: false}, nil
	}
	provider, err := uuid.ParseBytes(in.ProviderId)
	if err != nil {
		return &pb.SetAvailabilityResponse{Success: false}, err
	}

	err = s.h.SetAvailability(ctx, provider, availabilities)
	if err != nil {
		return &pb.SetAvailabilityResponse{Success: false}, err
	}
	return &pb.SetAvailabilityResponse{Success: true}, nil
}

func (s *server) GetAppointments(ctx context.Context, in *pb.GetAppointmentsRequest) (*pb.GetAppointmentsResponse, error) {
	providers := []*uuid.UUID{}
	for _, provider := range in.ProviderIds {
		id, err := uuid.ParseBytes(provider)
		if err != nil {
			return &pb.GetAppointmentsResponse{}, err
		}
		providers = append(providers, &id)
	}
	start := in.StartTime.AsTime()
	end := in.EndTime.AsTime()

	appointments, err := s.h.GetAppointments(ctx, providers, start, end)

	if err != nil {
		return &pb.GetAppointmentsResponse{}, err
	}

	converter := converters.AppointmentImpl{}
	protoAppointments := converter.ToProtos(appointments)

	return &pb.GetAppointmentsResponse{Appointments: protoAppointments}, nil

}

func (s *server) ReserveAppointment(ctx context.Context, in *pb.ReserveAppointmentRequest) (*pb.ReserveAppointmentResponse, error) {
	converter := converters.AppointmentImpl{}
	appointment := converter.FromProto(in.Appointment)
	patient, err := uuid.ParseBytes(in.PatientID)
	if err != nil {
		return &pb.ReserveAppointmentResponse{Success: false}, err
	}
	err = s.h.ReserveAppointment(ctx, patient, appointment)

	if err != nil {
		return &pb.ReserveAppointmentResponse{Success: false}, err
	}

	return &pb.ReserveAppointmentResponse{Success: true}, nil

}

func (s *server) ConfirmAppointment(ctx context.Context, in *pb.ConfrimAppointmentRequest) (*pb.ConfrimAppointmentResponse, error) {
	appointment, err := uuid.ParseBytes(in.AppointmentId)
	if err != nil {
		return &pb.ConfrimAppointmentResponse{Success: false}, err
	}
	patient, err := uuid.ParseBytes(in.PatientID)
	if err != nil {
		return &pb.ConfrimAppointmentResponse{Success: false}, err
	}
	err = s.h.ConfirmAppointment(ctx, patient, appointment)

	if err != nil {
		return &pb.ConfrimAppointmentResponse{Success: false}, err
	}

	return &pb.ConfrimAppointmentResponse{Success: true}, nil
}