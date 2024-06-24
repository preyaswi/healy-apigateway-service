// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/pb/patient/patient.proto

package patient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoogleSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoogleId     string `protobuf:"bytes,1,opt,name=google_id,json=googleId,proto3" json:"google_id,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AccessToken  string `protobuf:"bytes,4,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,5,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Tokenexpiry  string `protobuf:"bytes,6,opt,name=tokenexpiry,proto3" json:"tokenexpiry,omitempty"`
}

func (x *GoogleSignInRequest) Reset() {
	*x = GoogleSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleSignInRequest) ProtoMessage() {}

func (x *GoogleSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleSignInRequest.ProtoReflect.Descriptor instead.
func (*GoogleSignInRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleSignInRequest) GetGoogleId() string {
	if x != nil {
		return x.GoogleId
	}
	return ""
}

func (x *GoogleSignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GoogleSignInRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleSignInRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GoogleSignInRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *GoogleSignInRequest) GetTokenexpiry() string {
	if x != nil {
		return x.Tokenexpiry
	}
	return ""
}

type GoogleSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GoogleId string `protobuf:"bytes,2,opt,name=google_id,json=googleId,proto3" json:"google_id,omitempty"`
	Fullname string `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GoogleSignInResponse) Reset() {
	*x = GoogleSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleSignInResponse) ProtoMessage() {}

func (x *GoogleSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleSignInResponse.ProtoReflect.Descriptor instead.
func (*GoogleSignInResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleSignInResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GoogleSignInResponse) GetGoogleId() string {
	if x != nil {
		return x.GoogleId
	}
	return ""
}

func (x *GoogleSignInResponse) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *GoogleSignInResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PatientSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname        string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Confirmpassword string `protobuf:"bytes,4,opt,name=confirmpassword,proto3" json:"confirmpassword,omitempty"`
	Gender          string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber   string `protobuf:"bytes,6,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *PatientSignUpRequest) Reset() {
	*x = PatientSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientSignUpRequest) ProtoMessage() {}

func (x *PatientSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientSignUpRequest.ProtoReflect.Descriptor instead.
func (*PatientSignUpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{2}
}

func (x *PatientSignUpRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *PatientSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PatientSignUpRequest) GetConfirmpassword() string {
	if x != nil {
		return x.Confirmpassword
	}
	return ""
}

func (x *PatientSignUpRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientSignUpRequest) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type PatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname      string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Gender        string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber string `protobuf:"bytes,5,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *PatientDetails) Reset() {
	*x = PatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientDetails) ProtoMessage() {}

func (x *PatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientDetails.ProtoReflect.Descriptor instead.
func (*PatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{3}
}

func (x *PatientDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatientDetails) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *PatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientDetails) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientDetails) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type PatientSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientDetails *GoogleSignInResponse `protobuf:"bytes,1,opt,name=PatientDetails,proto3" json:"PatientDetails,omitempty"`
	AccessToken    string                `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken   string                `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *PatientSignUpResponse) Reset() {
	*x = PatientSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientSignUpResponse) ProtoMessage() {}

func (x *PatientSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientSignUpResponse.ProtoReflect.Descriptor instead.
func (*PatientSignUpResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{4}
}

func (x *PatientSignUpResponse) GetPatientDetails() *GoogleSignInResponse {
	if x != nil {
		return x.PatientDetails
	}
	return nil
}

func (x *PatientSignUpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PatientSignUpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type PatientLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PatientLoginRequest) Reset() {
	*x = PatientLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientLoginRequest) ProtoMessage() {}

func (x *PatientLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientLoginRequest.ProtoReflect.Descriptor instead.
func (*PatientLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{5}
}

func (x *PatientLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PatientLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientDetails *PatientDetails `protobuf:"bytes,1,opt,name=PatientDetails,proto3" json:"PatientDetails,omitempty"`
	AccessToken    string          `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken   string          `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *PatientLoginResponse) Reset() {
	*x = PatientLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientLoginResponse) ProtoMessage() {}

func (x *PatientLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientLoginResponse.ProtoReflect.Descriptor instead.
func (*PatientLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{6}
}

func (x *PatientLoginResponse) GetPatientDetails() *PatientDetails {
	if x != nil {
		return x.PatientDetails
	}
	return nil
}

func (x *PatientLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PatientLoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Idreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Idreq) Reset() {
	*x = Idreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idreq) ProtoMessage() {}

func (x *Idreq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idreq.ProtoReflect.Descriptor instead.
func (*Idreq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{7}
}

func (x *Idreq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type InPatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname      string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Gender        string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber string `protobuf:"bytes,4,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *InPatientDetails) Reset() {
	*x = InPatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InPatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InPatientDetails) ProtoMessage() {}

func (x *InPatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InPatientDetails.ProtoReflect.Descriptor instead.
func (*InPatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{8}
}

func (x *InPatientDetails) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *InPatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InPatientDetails) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *InPatientDetails) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId        string            `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	InPatientDetails *InPatientDetails `protobuf:"bytes,2,opt,name=InPatientDetails,proto3" json:"InPatientDetails,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *UpdateRequest) GetInPatientDetails() *InPatientDetails {
	if x != nil {
		return x.InPatientDetails
	}
	return nil
}

type Listpares struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pa []*PatientDetails `protobuf:"bytes,1,rep,name=pa,proto3" json:"pa,omitempty"`
}

func (x *Listpares) Reset() {
	*x = Listpares{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listpares) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listpares) ProtoMessage() {}

func (x *Listpares) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listpares.ProtoReflect.Descriptor instead.
func (*Listpares) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{10}
}

func (x *Listpares) GetPa() []*PatientDetails {
	if x != nil {
		return x.Pa
	}
	return nil
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{11}
}

var File_pkg_pb_patient_patient_proto protoreflect.FileDescriptor

var file_pkg_pb_patient_patient_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x75,
	0x0a, 0x14, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47,
	0x0a, 0x13, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x05, 0x69, 0x64, 0x72, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x49, 0x6e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x75,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x10, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x10, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x34, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x61, 0x72,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x02, 0x70, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x02, 0x70, 0x61, 0x22, 0x05, 0x0a, 0x03, 0x72,
	0x65, 0x71, 0x32, 0x9a, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x4e,
	0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1c,
	0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x11, 0x49, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x64,
	0x72, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x61, 0x72, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_patient_patient_proto_rawDescOnce sync.Once
	file_pkg_pb_patient_patient_proto_rawDescData = file_pkg_pb_patient_patient_proto_rawDesc
)

func file_pkg_pb_patient_patient_proto_rawDescGZIP() []byte {
	file_pkg_pb_patient_patient_proto_rawDescOnce.Do(func() {
		file_pkg_pb_patient_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_patient_patient_proto_rawDescData)
	})
	return file_pkg_pb_patient_patient_proto_rawDescData
}

var file_pkg_pb_patient_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_pb_patient_patient_proto_goTypes = []interface{}{
	(*GoogleSignInRequest)(nil),   // 0: patient.GoogleSignInRequest
	(*GoogleSignInResponse)(nil),  // 1: patient.GoogleSignInResponse
	(*PatientSignUpRequest)(nil),  // 2: patient.PatientSignUpRequest
	(*PatientDetails)(nil),        // 3: patient.PatientDetails
	(*PatientSignUpResponse)(nil), // 4: patient.PatientSignUpResponse
	(*PatientLoginRequest)(nil),   // 5: patient.PatientLoginRequest
	(*PatientLoginResponse)(nil),  // 6: patient.PatientLoginResponse
	(*Idreq)(nil),                 // 7: patient.idreq
	(*InPatientDetails)(nil),      // 8: patient.InPatientDetails
	(*UpdateRequest)(nil),         // 9: patient.UpdateRequest
	(*Listpares)(nil),             // 10: patient.listpares
	(*Req)(nil),                   // 11: patient.req
}
var file_pkg_pb_patient_patient_proto_depIdxs = []int32{
	1,  // 0: patient.PatientSignUpResponse.PatientDetails:type_name -> patient.GoogleSignInResponse
	3,  // 1: patient.PatientLoginResponse.PatientDetails:type_name -> patient.PatientDetails
	8,  // 2: patient.UpdateRequest.InPatientDetails:type_name -> patient.InPatientDetails
	3,  // 3: patient.listpares.pa:type_name -> patient.PatientDetails
	0,  // 4: patient.Patient.GoogleSignIn:input_type -> patient.GoogleSignInRequest
	7,  // 5: patient.Patient.IndPatientDetails:input_type -> patient.idreq
	9,  // 6: patient.Patient.UpdatePatientDetails:input_type -> patient.UpdateRequest
	11, // 7: patient.Patient.ListPatients:input_type -> patient.req
	4,  // 8: patient.Patient.GoogleSignIn:output_type -> patient.PatientSignUpResponse
	3,  // 9: patient.Patient.IndPatientDetails:output_type -> patient.PatientDetails
	8,  // 10: patient.Patient.UpdatePatientDetails:output_type -> patient.InPatientDetails
	10, // 11: patient.Patient.ListPatients:output_type -> patient.listpares
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_patient_patient_proto_init() }
func file_pkg_pb_patient_patient_proto_init() {
	if File_pkg_pb_patient_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_patient_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleSignInRequest); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleSignInResponse); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientSignUpRequest); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientDetails); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientSignUpResponse); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientLoginRequest); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientLoginResponse); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idreq); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InPatientDetails); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listpares); i {
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
		file_pkg_pb_patient_patient_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
			RawDescriptor: file_pkg_pb_patient_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_patient_patient_proto_goTypes,
		DependencyIndexes: file_pkg_pb_patient_patient_proto_depIdxs,
		MessageInfos:      file_pkg_pb_patient_patient_proto_msgTypes,
	}.Build()
	File_pkg_pb_patient_patient_proto = out.File
	file_pkg_pb_patient_patient_proto_rawDesc = nil
	file_pkg_pb_patient_patient_proto_goTypes = nil
	file_pkg_pb_patient_patient_proto_depIdxs = nil
}
