// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/proto/quote/quote_service.proto

package quotepb

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

type CreateQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateQuoteRequest) Reset() {
	*x = CreateQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuoteRequest) ProtoMessage() {}

func (x *CreateQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuoteRequest.ProtoReflect.Descriptor instead.
func (*CreateQuoteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateQuoteRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreateQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateQuoteResponse) Reset() {
	*x = CreateQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuoteResponse) ProtoMessage() {}

func (x *CreateQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuoteResponse.ProtoReflect.Descriptor instead.
func (*CreateQuoteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuoteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQuoteRequest) Reset() {
	*x = ListQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuoteRequest) ProtoMessage() {}

func (x *ListQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuoteRequest.ProtoReflect.Descriptor instead.
func (*ListQuoteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{2}
}

type ListQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotes []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *ListQuoteResponse) Reset() {
	*x = ListQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuoteResponse) ProtoMessage() {}

func (x *ListQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuoteResponse.ProtoReflect.Descriptor instead.
func (*ListQuoteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListQuoteResponse) GetQuotes() []*Quote {
	if x != nil {
		return x.Quotes
	}
	return nil
}

type DetailQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetailQuoteRequest) Reset() {
	*x = DetailQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailQuoteRequest) ProtoMessage() {}

func (x *DetailQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailQuoteRequest.ProtoReflect.Descriptor instead.
func (*DetailQuoteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{4}
}

func (x *DetailQuoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DetailQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *DetailQuoteResponse) Reset() {
	*x = DetailQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailQuoteResponse) ProtoMessage() {}

func (x *DetailQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailQuoteResponse.ProtoReflect.Descriptor instead.
func (*DetailQuoteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{5}
}

func (x *DetailQuoteResponse) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

type UpdateQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateQuoteRequest) Reset() {
	*x = UpdateQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuoteRequest) ProtoMessage() {}

func (x *UpdateQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuoteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateQuoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQuoteRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type UpdateQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateQuoteResponse) Reset() {
	*x = UpdateQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuoteResponse) ProtoMessage() {}

func (x *UpdateQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuoteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateQuoteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuoteRequest) Reset() {
	*x = DeleteQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuoteRequest) ProtoMessage() {}

func (x *DeleteQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuoteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteQuoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuoteResponse) Reset() {
	*x = DeleteQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_quote_quote_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuoteResponse) ProtoMessage() {}

func (x *DeleteQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_quote_quote_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuoteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_quote_quote_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteQuoteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_proto_quote_quote_service_proto protoreflect.FileDescriptor

var file_api_proto_quote_quote_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xca, 0x03, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x58, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x50, 0x69, 0x63, 0x6b, 0x48, 0x44, 0x2f, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x2d, 0x65,
	0x63, 0x68, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x3b, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_quote_quote_service_proto_rawDescOnce sync.Once
	file_api_proto_quote_quote_service_proto_rawDescData = file_api_proto_quote_quote_service_proto_rawDesc
)

func file_api_proto_quote_quote_service_proto_rawDescGZIP() []byte {
	file_api_proto_quote_quote_service_proto_rawDescOnce.Do(func() {
		file_api_proto_quote_quote_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_quote_quote_service_proto_rawDescData)
	})
	return file_api_proto_quote_quote_service_proto_rawDescData
}

var file_api_proto_quote_quote_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_quote_quote_service_proto_goTypes = []interface{}{
	(*CreateQuoteRequest)(nil),  // 0: api.proto.quote.CreateQuoteRequest
	(*CreateQuoteResponse)(nil), // 1: api.proto.quote.CreateQuoteResponse
	(*ListQuoteRequest)(nil),    // 2: api.proto.quote.ListQuoteRequest
	(*ListQuoteResponse)(nil),   // 3: api.proto.quote.ListQuoteResponse
	(*DetailQuoteRequest)(nil),  // 4: api.proto.quote.DetailQuoteRequest
	(*DetailQuoteResponse)(nil), // 5: api.proto.quote.DetailQuoteResponse
	(*UpdateQuoteRequest)(nil),  // 6: api.proto.quote.UpdateQuoteRequest
	(*UpdateQuoteResponse)(nil), // 7: api.proto.quote.UpdateQuoteResponse
	(*DeleteQuoteRequest)(nil),  // 8: api.proto.quote.DeleteQuoteRequest
	(*DeleteQuoteResponse)(nil), // 9: api.proto.quote.DeleteQuoteResponse
	(*Quote)(nil),               // 10: api.proto.Quote
}
var file_api_proto_quote_quote_service_proto_depIdxs = []int32{
	10, // 0: api.proto.quote.ListQuoteResponse.quotes:type_name -> api.proto.quote.Quote
	10, // 1: api.proto.quote.DetailQuoteResponse.quote:type_name -> api.proto.quote.Quote
	0,  // 2: api.proto.quote.QuoteService.CreateQuote:input_type -> api.proto.quote.CreateQuoteRequest
	2,  // 3: api.proto.quote.QuoteService.ListQuote:input_type -> api.proto.quote.ListQuoteRequest
	4,  // 4: api.proto.quote.QuoteService.DetailQuote:input_type -> api.proto.quote.DetailQuoteRequest
	6,  // 5: api.proto.quote.QuoteService.UpdateQuote:input_type -> api.proto.quote.UpdateQuoteRequest
	8,  // 6: api.proto.quote.QuoteService.DeleteQuote:input_type -> api.proto.quote.DeleteQuoteRequest
	1,  // 7: api.proto.quote.QuoteService.CreateQuote:output_type -> api.proto.quote.CreateQuoteResponse
	3,  // 8: api.proto.quote.QuoteService.ListQuote:output_type -> api.proto.quote.ListQuoteResponse
	5,  // 9: api.proto.quote.QuoteService.DetailQuote:output_type -> api.proto.quote.DetailQuoteResponse
	7,  // 10: api.proto.quote.QuoteService.UpdateQuote:output_type -> api.proto.quote.UpdateQuoteResponse
	9,  // 11: api.proto.quote.QuoteService.DeleteQuote:output_type -> api.proto.quote.DeleteQuoteResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_quote_quote_service_proto_init() }
func file_api_proto_quote_quote_service_proto_init() {
	if File_api_proto_quote_quote_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_quote_quote_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuoteRequest); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuoteResponse); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuoteRequest); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuoteResponse); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailQuoteRequest); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailQuoteResponse); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuoteRequest); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuoteResponse); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuoteRequest); i {
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
		file_api_proto_quote_quote_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuoteResponse); i {
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
			RawDescriptor: file_api_proto_quote_quote_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_quote_quote_service_proto_goTypes,
		DependencyIndexes: file_api_proto_quote_quote_service_proto_depIdxs,
		MessageInfos:      file_api_proto_quote_quote_service_proto_msgTypes,
	}.Build()
	File_api_proto_quote_quote_service_proto = out.File
	file_api_proto_quote_quote_service_proto_rawDesc = nil
	file_api_proto_quote_quote_service_proto_goTypes = nil
	file_api_proto_quote_quote_service_proto_depIdxs = nil
}
