// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: birthday.proto

package birthday_proto

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

type Requests int32

const (
	Requests_POST   Requests = 0
	Requests_GET    Requests = 1
	Requests_PUT    Requests = 2
	Requests_DELETE Requests = 3
)

// Enum value maps for Requests.
var (
	Requests_name = map[int32]string{
		0: "POST",
		1: "GET",
		2: "PUT",
		3: "DELETE",
	}
	Requests_value = map[string]int32{
		"POST":   0,
		"GET":    1,
		"PUT":    2,
		"DELETE": 3,
	}
)

func (x Requests) Enum() *Requests {
	p := new(Requests)
	*p = x
	return p
}

func (x Requests) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Requests) Descriptor() protoreflect.EnumDescriptor {
	return file_birthday_proto_enumTypes[0].Descriptor()
}

func (Requests) Type() protoreflect.EnumType {
	return &file_birthday_proto_enumTypes[0]
}

func (x Requests) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Requests.Descriptor instead.
func (Requests) EnumDescriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{0}
}

type CreateBirthdayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the name of the person
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//the date of birth of the person
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	//the identification number of the person
	PersonalNumber string `protobuf:"bytes,3,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
}

func (x *CreateBirthdayRequest) Reset() {
	*x = CreateBirthdayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBirthdayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBirthdayRequest) ProtoMessage() {}

func (x *CreateBirthdayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBirthdayRequest.ProtoReflect.Descriptor instead.
func (*CreateBirthdayRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBirthdayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBirthdayRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CreateBirthdayRequest) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

type DeleteBirthdayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBirthdayResponse) Reset() {
	*x = DeleteBirthdayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBirthdayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBirthdayResponse) ProtoMessage() {}

func (x *DeleteBirthdayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBirthdayResponse.ProtoReflect.Descriptor instead.
func (*DeleteBirthdayResponse) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{1}
}

type GetBirthdayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the identification number of the person
	PersonalNumber string `protobuf:"bytes,1,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
}

func (x *GetBirthdayRequest) Reset() {
	*x = GetBirthdayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBirthdayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBirthdayRequest) ProtoMessage() {}

func (x *GetBirthdayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBirthdayRequest.ProtoReflect.Descriptor instead.
func (*GetBirthdayRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{2}
}

func (x *GetBirthdayRequest) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

type UpdateBirthdayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the identification number of the person
	PersonalNumber string `protobuf:"bytes,1,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
	//the name of the person
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//the date of birth of the person
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *UpdateBirthdayRequest) Reset() {
	*x = UpdateBirthdayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBirthdayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBirthdayRequest) ProtoMessage() {}

func (x *UpdateBirthdayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBirthdayRequest.ProtoReflect.Descriptor instead.
func (*UpdateBirthdayRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBirthdayRequest) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

func (x *UpdateBirthdayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBirthdayRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type DeleteBirthdayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the identification number of the person
	PersonalNumber string `protobuf:"bytes,1,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
}

func (x *DeleteBirthdayRequest) Reset() {
	*x = DeleteBirthdayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBirthdayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBirthdayRequest) ProtoMessage() {}

func (x *DeleteBirthdayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBirthdayRequest.ProtoReflect.Descriptor instead.
func (*DeleteBirthdayRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBirthdayRequest) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

type BirthdayObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the name of the person
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//the date of birth of the person
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	//the identification number of the person
	PersonalNumber string `protobuf:"bytes,3,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
	//the unique id that's created for the person
	ID string `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *BirthdayObject) Reset() {
	*x = BirthdayObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BirthdayObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BirthdayObject) ProtoMessage() {}

func (x *BirthdayObject) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BirthdayObject.ProtoReflect. Descriptor instead.
func (*BirthdayObject) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{5}
}

func (x *BirthdayObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BirthdayObject) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BirthdayObject) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

func (x *BirthdayObject) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_birthday_proto protoreflect.FileDescriptor

var file_birthday_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x67, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x67, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x3f, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x0e, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x2a, 0x32, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x32, 0xd1, 0x02, 0x0a, 0x11,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12,
	0x1c, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x72, 0x69, 0x6e, 0x42, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x79, 0x2f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_birthday_proto_rawDescOnce sync.Once
	file_birthday_proto_rawDescData = file_birthday_proto_rawDesc
)

func file_birthday_proto_rawDescGZIP() []byte {
	file_birthday_proto_rawDescOnce.Do(func() {
		file_birthday_proto_rawDescData = protoimpl.X.CompressGZIP(file_birthday_proto_rawDescData)
	})
	return file_birthday_proto_rawDescData
}

var file_birthday_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_birthday_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_birthday_proto_goTypes = []interface{}{
	(Requests)(0),                  // 0: birthday.Requests
	(*CreateBirthdayRequest)(nil),  // 1: birthday.CreateBirthdayRequest
	(*DeleteBirthdayResponse)(nil), // 2: birthday.DeleteBirthdayResponse
	(*GetBirthdayRequest)(nil),     // 3: birthday.GetBirthdayRequest
	(*UpdateBirthdayRequest)(nil),  // 4: birthday.UpdateBirthdayRequest
	(*DeleteBirthdayRequest)(nil),  // 5: birthday.DeleteBirthdayRequest
	(*BirthdayObject)(nil),         // 6: birthday.BirthdayObject
}
var file_birthday_proto_depIdxs = []int32{
	1, // 0: birthday.BirthdayFunctions.CreateBirthday:input_type -> birthday.CreateBirthdayRequest
	3, // 1: birthday.BirthdayFunctions.GetBirthday:input_type -> birthday.GetBirthdayRequest
	4, // 2: birthday.BirthdayFunctions.UpdateBirthday:input_type -> birthday.UpdateBirthdayRequest
	5, // 3: birthday.BirthdayFunctions.DeleteBirthday:input_type -> birthday.DeleteBirthdayRequest
	6, // 4: birthday.BirthdayFunctions.CreateBirthday:output_type -> birthday.BirthdayObject
	6, // 5: birthday.BirthdayFunctions.GetBirthday:output_type -> birthday.BirthdayObject
	6, // 6: birthday.BirthdayFunctions.UpdateBirthday:output_type -> birthday.BirthdayObject
	2, // 7: birthday.BirthdayFunctions.DeleteBirthday:output_type -> birthday.DeleteBirthdayResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_birthday_proto_init() }
func file_birthday_proto_init() {
	if File_birthday_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_birthday_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBirthdayRequest); i {
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
		file_birthday_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBirthdayResponse); i {
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
		file_birthday_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBirthdayRequest); i {
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
		file_birthday_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBirthdayRequest); i {
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
		file_birthday_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBirthdayRequest); i {
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
		file_birthday_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BirthdayObject); i {
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
			RawDescriptor: file_birthday_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_birthday_proto_goTypes,
		DependencyIndexes: file_birthday_proto_depIdxs,
		EnumInfos:         file_birthday_proto_enumTypes,
		MessageInfos:      file_birthday_proto_msgTypes,
	}.Build()
	File_birthday_proto = out.File
	file_birthday_proto_rawDesc = nil
	file_birthday_proto_goTypes = nil
	file_birthday_proto_depIdxs = nil
}
