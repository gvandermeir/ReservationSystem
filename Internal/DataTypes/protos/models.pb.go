// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: models.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Availability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId []byte                 `protobuf:"bytes,1,opt,name=providerId,proto3" json:"providerId,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *Availability) Reset() {
	*x = Availability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Availability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Availability) ProtoMessage() {}

func (x *Availability) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Availability.ProtoReflect.Descriptor instead.
func (*Availability) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Availability) GetProviderId() []byte {
	if x != nil {
		return x.ProviderId
	}
	return nil
}

func (x *Availability) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Availability) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *Patient) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Patient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Appointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProviderId []byte                 `protobuf:"bytes,2,opt,name=providerId,proto3" json:"providerId,omitempty"`
	PatientId  []byte                 `protobuf:"bytes,3,opt,name=patientId,proto3" json:"patientId,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	BookedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=bookedAt,proto3" json:"bookedAt,omitempty"`
	Confirmed  bool                   `protobuf:"varint,7,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (x *Appointment) Reset() {
	*x = Appointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appointment) ProtoMessage() {}

func (x *Appointment) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appointment.ProtoReflect.Descriptor instead.
func (*Appointment) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *Appointment) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Appointment) GetProviderId() []byte {
	if x != nil {
		return x.ProviderId
	}
	return nil
}

func (x *Appointment) GetPatientId() []byte {
	if x != nil {
		return x.PatientId
	}
	return nil
}

func (x *Appointment) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Appointment) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Appointment) GetBookedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.BookedAt
	}
	return nil
}

func (x *Appointment) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

type SetAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId     []byte          `protobuf:"bytes,1,opt,name=providerId,proto3" json:"providerId,omitempty"`
	Availabilities []*Availability `protobuf:"bytes,2,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
}

func (x *SetAvailabilityRequest) Reset() {
	*x = SetAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvailabilityRequest) ProtoMessage() {}

func (x *SetAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*SetAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{4}
}

func (x *SetAvailabilityRequest) GetProviderId() []byte {
	if x != nil {
		return x.ProviderId
	}
	return nil
}

func (x *SetAvailabilityRequest) GetAvailabilities() []*Availability {
	if x != nil {
		return x.Availabilities
	}
	return nil
}

type SetAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SetAvailabilityResponse) Reset() {
	*x = SetAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvailabilityResponse) ProtoMessage() {}

func (x *SetAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*SetAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{5}
}

func (x *SetAvailabilityResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAppointmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderIds [][]byte               `protobuf:"bytes,1,rep,name=providerIds,proto3" json:"providerIds,omitempty"`
	StartTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *GetAppointmentsRequest) Reset() {
	*x = GetAppointmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppointmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppointmentsRequest) ProtoMessage() {}

func (x *GetAppointmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppointmentsRequest.ProtoReflect.Descriptor instead.
func (*GetAppointmentsRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppointmentsRequest) GetProviderIds() [][]byte {
	if x != nil {
		return x.ProviderIds
	}
	return nil
}

func (x *GetAppointmentsRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetAppointmentsRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetAppointmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointments []*Appointment `protobuf:"bytes,1,rep,name=appointments,proto3" json:"appointments,omitempty"`
}

func (x *GetAppointmentsResponse) Reset() {
	*x = GetAppointmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppointmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppointmentsResponse) ProtoMessage() {}

func (x *GetAppointmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppointmentsResponse.ProtoReflect.Descriptor instead.
func (*GetAppointmentsResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppointmentsResponse) GetAppointments() []*Appointment {
	if x != nil {
		return x.Appointments
	}
	return nil
}

type ReserveAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientID   []byte       `protobuf:"bytes,1,opt,name=patientID,proto3" json:"patientID,omitempty"`
	Appointment *Appointment `protobuf:"bytes,2,opt,name=appointment,proto3" json:"appointment,omitempty"`
}

func (x *ReserveAppointmentRequest) Reset() {
	*x = ReserveAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveAppointmentRequest) ProtoMessage() {}

func (x *ReserveAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveAppointmentRequest.ProtoReflect.Descriptor instead.
func (*ReserveAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{8}
}

func (x *ReserveAppointmentRequest) GetPatientID() []byte {
	if x != nil {
		return x.PatientID
	}
	return nil
}

func (x *ReserveAppointmentRequest) GetAppointment() *Appointment {
	if x != nil {
		return x.Appointment
	}
	return nil
}

type ReserveAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ReserveAppointmentResponse) Reset() {
	*x = ReserveAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveAppointmentResponse) ProtoMessage() {}

func (x *ReserveAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveAppointmentResponse.ProtoReflect.Descriptor instead.
func (*ReserveAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{9}
}

func (x *ReserveAppointmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ConfrimAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientID     []byte `protobuf:"bytes,1,opt,name=patientID,proto3" json:"patientID,omitempty"`
	AppointmentId []byte `protobuf:"bytes,2,opt,name=appointmentId,proto3" json:"appointmentId,omitempty"`
}

func (x *ConfrimAppointmentRequest) Reset() {
	*x = ConfrimAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfrimAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfrimAppointmentRequest) ProtoMessage() {}

func (x *ConfrimAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfrimAppointmentRequest.ProtoReflect.Descriptor instead.
func (*ConfrimAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{10}
}

func (x *ConfrimAppointmentRequest) GetPatientID() []byte {
	if x != nil {
		return x.PatientID
	}
	return nil
}

func (x *ConfrimAppointmentRequest) GetAppointmentId() []byte {
	if x != nil {
		return x.AppointmentId
	}
	return nil
}

type ConfrimAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ConfrimAppointmentResponse) Reset() {
	*x = ConfrimAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfrimAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfrimAppointmentResponse) ProtoMessage() {}

func (x *ConfrimAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfrimAppointmentResponse.ProtoReflect.Descriptor instead.
func (*ConfrimAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{11}
}

func (x *ConfrimAppointmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x9e, 0x01,
	0x0a, 0x0c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x43,
	0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x76, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x33, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x35, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x5f, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x72, 0x69, 0x6d, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x36, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x66, 0x72, 0x69, 0x6d, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfe, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x54, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x66, 0x72, 0x69, 0x6d, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x72, 0x69,
	0x6d, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x72, 0x69, 0x6d, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x76, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x65, 0x69, 0x72, 0x2f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_models_proto_goTypes = []interface{}{
	(*Provider)(nil),                   // 0: protos.Provider
	(*Availability)(nil),               // 1: protos.Availability
	(*Patient)(nil),                    // 2: protos.Patient
	(*Appointment)(nil),                // 3: protos.Appointment
	(*SetAvailabilityRequest)(nil),     // 4: protos.SetAvailabilityRequest
	(*SetAvailabilityResponse)(nil),    // 5: protos.SetAvailabilityResponse
	(*GetAppointmentsRequest)(nil),     // 6: protos.GetAppointmentsRequest
	(*GetAppointmentsResponse)(nil),    // 7: protos.GetAppointmentsResponse
	(*ReserveAppointmentRequest)(nil),  // 8: protos.ReserveAppointmentRequest
	(*ReserveAppointmentResponse)(nil), // 9: protos.ReserveAppointmentResponse
	(*ConfrimAppointmentRequest)(nil),  // 10: protos.ConfrimAppointmentRequest
	(*ConfrimAppointmentResponse)(nil), // 11: protos.ConfrimAppointmentResponse
	(*timestamppb.Timestamp)(nil),      // 12: google.protobuf.Timestamp
}
var file_models_proto_depIdxs = []int32{
	12, // 0: protos.Availability.startTime:type_name -> google.protobuf.Timestamp
	12, // 1: protos.Availability.endTime:type_name -> google.protobuf.Timestamp
	12, // 2: protos.Appointment.startTime:type_name -> google.protobuf.Timestamp
	12, // 3: protos.Appointment.endTime:type_name -> google.protobuf.Timestamp
	12, // 4: protos.Appointment.bookedAt:type_name -> google.protobuf.Timestamp
	1,  // 5: protos.SetAvailabilityRequest.availabilities:type_name -> protos.Availability
	12, // 6: protos.GetAppointmentsRequest.startTime:type_name -> google.protobuf.Timestamp
	12, // 7: protos.GetAppointmentsRequest.endTime:type_name -> google.protobuf.Timestamp
	3,  // 8: protos.GetAppointmentsResponse.appointments:type_name -> protos.Appointment
	3,  // 9: protos.ReserveAppointmentRequest.appointment:type_name -> protos.Appointment
	4,  // 10: protos.ReservationService.SetAvailability:input_type -> protos.SetAvailabilityRequest
	6,  // 11: protos.ReservationService.GetAppointments:input_type -> protos.GetAppointmentsRequest
	8,  // 12: protos.ReservationService.ReserveAppointment:input_type -> protos.ReserveAppointmentRequest
	10, // 13: protos.ReservationService.ConfrimAppointment:input_type -> protos.ConfrimAppointmentRequest
	5,  // 14: protos.ReservationService.SetAvailability:output_type -> protos.SetAvailabilityResponse
	7,  // 15: protos.ReservationService.GetAppointments:output_type -> protos.GetAppointmentsResponse
	9,  // 16: protos.ReservationService.ReserveAppointment:output_type -> protos.ReserveAppointmentResponse
	11, // 17: protos.ReservationService.ConfrimAppointment:output_type -> protos.ConfrimAppointmentResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Availability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appointment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvailabilityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvailabilityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppointmentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppointmentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveAppointmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveAppointmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfrimAppointmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfrimAppointmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}