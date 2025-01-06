// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: okrs.proto

package proto

import (
	_type "github.com/thylong/rft-backend/pkg/proto/google/type"
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

// Kr represents a kr that occured in the company
type Kr struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Sponsor       string                 `protobuf:"bytes,4,opt,name=sponsor,proto3" json:"sponsor,omitempty"`
	Kpis          string                 `protobuf:"bytes,5,opt,name=kpis,proto3" json:"kpis,omitempty"`
	Scope         string                 `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	Initiatives   string                 `protobuf:"bytes,7,opt,name=initiatives,proto3" json:"initiatives,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Kr) Reset() {
	*x = Kr{}
	mi := &file_okrs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Kr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kr) ProtoMessage() {}

func (x *Kr) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kr.ProtoReflect.Descriptor instead.
func (*Kr) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{0}
}

func (x *Kr) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Kr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Kr) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Kr) GetSponsor() string {
	if x != nil {
		return x.Sponsor
	}
	return ""
}

func (x *Kr) GetKpis() string {
	if x != nil {
		return x.Kpis
	}
	return ""
}

func (x *Kr) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Kr) GetInitiatives() string {
	if x != nil {
		return x.Initiatives
	}
	return ""
}

// Okr represents an okr that occured in the company
type Okr struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Number        int32                  `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Year          int32                  `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	EmbeddedChild []*Kr                  `protobuf:"bytes,7,rep,name=embedded_child,json=embeddedChild,proto3" json:"embedded_child,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Okr) Reset() {
	*x = Okr{}
	mi := &file_okrs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Okr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Okr) ProtoMessage() {}

func (x *Okr) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Okr.ProtoReflect.Descriptor instead.
func (*Okr) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{1}
}

func (x *Okr) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Okr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Okr) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Okr) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Okr) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Okr) GetEmbeddedChild() []*Kr {
	if x != nil {
		return x.EmbeddedChild
	}
	return nil
}

type GetOkrsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pagination parameters (optional)
	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                         // The page number
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // Number of okrs per page
	// Filtering parameters (optional)
	Search        string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"` // Search query for filtering okrs
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOkrsRequest) Reset() {
	*x = GetOkrsRequest{}
	mi := &file_okrs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOkrsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOkrsRequest) ProtoMessage() {}

func (x *GetOkrsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOkrsRequest.ProtoReflect.Descriptor instead.
func (*GetOkrsRequest) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{2}
}

func (x *GetOkrsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOkrsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetOkrsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

// Response message for GetOkrs
type GetOkrsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Okrs  []*Okr                 `protobuf:"bytes,1,rep,name=okrs,proto3" json:"okrs,omitempty"`
	// Pagination metadata
	TotalCount    int32 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` // Total number of okrs
	Page          int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`                               // Current page
	PageSize      int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`       // Number of okrs per page
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOkrsResponse) Reset() {
	*x = GetOkrsResponse{}
	mi := &file_okrs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOkrsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOkrsResponse) ProtoMessage() {}

func (x *GetOkrsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOkrsResponse.ProtoReflect.Descriptor instead.
func (*GetOkrsResponse) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{3}
}

func (x *GetOkrsResponse) GetOkrs() []*Okr {
	if x != nil {
		return x.Okrs
	}
	return nil
}

func (x *GetOkrsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetOkrsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOkrsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetOkrRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOkrRequest) Reset() {
	*x = GetOkrRequest{}
	mi := &file_okrs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOkrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOkrRequest) ProtoMessage() {}

func (x *GetOkrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOkrRequest.ProtoReflect.Descriptor instead.
func (*GetOkrRequest) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{4}
}

func (x *GetOkrRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOkrResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Okr           *Okr                   `protobuf:"bytes,1,opt,name=okr,proto3" json:"okr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOkrResponse) Reset() {
	*x = GetOkrResponse{}
	mi := &file_okrs_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOkrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOkrResponse) ProtoMessage() {}

func (x *GetOkrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOkrResponse.ProtoReflect.Descriptor instead.
func (*GetOkrResponse) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{5}
}

func (x *GetOkrResponse) GetOkr() *Okr {
	if x != nil {
		return x.Okr
	}
	return nil
}

type PutOkrRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type          string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Department    string                 `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Regions       []string               `protobuf:"bytes,7,rep,name=regions,proto3" json:"regions,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	StartAt       *_type.DateTime        `protobuf:"bytes,9,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutOkrRequest) Reset() {
	*x = PutOkrRequest{}
	mi := &file_okrs_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutOkrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutOkrRequest) ProtoMessage() {}

func (x *PutOkrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutOkrRequest.ProtoReflect.Descriptor instead.
func (*PutOkrRequest) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{6}
}

func (x *PutOkrRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutOkrRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutOkrRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PutOkrRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PutOkrRequest) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *PutOkrRequest) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *PutOkrRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PutOkrRequest) GetStartAt() *_type.DateTime {
	if x != nil {
		return x.StartAt
	}
	return nil
}

type PutOkrResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Okr           *Okr                   `protobuf:"bytes,1,opt,name=okr,proto3" json:"okr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutOkrResponse) Reset() {
	*x = PutOkrResponse{}
	mi := &file_okrs_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutOkrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutOkrResponse) ProtoMessage() {}

func (x *PutOkrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutOkrResponse.ProtoReflect.Descriptor instead.
func (*PutOkrResponse) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{7}
}

func (x *PutOkrResponse) GetOkr() *Okr {
	if x != nil {
		return x.Okr
	}
	return nil
}

type DeleteOkrRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOkrRequest) Reset() {
	*x = DeleteOkrRequest{}
	mi := &file_okrs_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOkrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOkrRequest) ProtoMessage() {}

func (x *DeleteOkrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOkrRequest.ProtoReflect.Descriptor instead.
func (*DeleteOkrRequest) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteOkrRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOkrResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOkrResponse) Reset() {
	*x = DeleteOkrResponse{}
	mi := &file_okrs_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOkrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOkrResponse) ProtoMessage() {}

func (x *DeleteOkrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okrs_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOkrResponse.ProtoReflect.Descriptor instead.
func (*DeleteOkrResponse) Descriptor() ([]byte, []int) {
	return file_okrs_proto_rawDescGZIP(), []int{9}
}

var File_okrs_proto protoreflect.FileDescriptor

var file_okrs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6f, 0x6b,
	0x72, 0x73, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x02, 0x4b, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x70, 0x69, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x70, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x22, 0xa8, 0x01, 0x0a, 0x03, 0x4f, 0x6b, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0e, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x4b, 0x72, 0x52, 0x0d, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x22, 0x59, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x6b, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6f,
	0x6b, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6f, 0x6b, 0x72, 0x73,
	0x2e, 0x4f, 0x6b, 0x72, 0x52, 0x04, 0x6f, 0x6b, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x1f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x03, 0x6f, 0x6b, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6f,
	0x6b, 0x72, 0x73, 0x2e, 0x4f, 0x6b, 0x72, 0x52, 0x03, 0x6f, 0x6b, 0x72, 0x22, 0xe9, 0x01, 0x0a,
	0x0d, 0x50, 0x75, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x4f,
	0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x6f, 0x6b,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x4f,
	0x6b, 0x72, 0x52, 0x03, 0x6f, 0x6b, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xf4, 0x01, 0x0a, 0x0a, 0x4f, 0x6b, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x6f, 0x6b, 0x72,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x4f, 0x6b, 0x72, 0x12, 0x13, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6b,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x4f, 0x6b, 0x72, 0x12, 0x13, 0x2e, 0x6f, 0x6b, 0x72,
	0x73, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x6b, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x6b, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f,
	0x6b, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6b, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x6b, 0x72, 0x73, 0x42, 0x09, 0x4f, 0x6b, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x79,
	0x6c, 0x6f, 0x6e, 0x67, 0x2f, 0x72, 0x66, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58,
	0xaa, 0x02, 0x04, 0x4f, 0x6b, 0x72, 0x73, 0xca, 0x02, 0x04, 0x4f, 0x6b, 0x72, 0x73, 0xe2, 0x02,
	0x10, 0x4f, 0x6b, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x04, 0x4f, 0x6b, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_okrs_proto_rawDescOnce sync.Once
	file_okrs_proto_rawDescData = file_okrs_proto_rawDesc
)

func file_okrs_proto_rawDescGZIP() []byte {
	file_okrs_proto_rawDescOnce.Do(func() {
		file_okrs_proto_rawDescData = protoimpl.X.CompressGZIP(file_okrs_proto_rawDescData)
	})
	return file_okrs_proto_rawDescData
}

var file_okrs_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_okrs_proto_goTypes = []any{
	(*Kr)(nil),                // 0: okrs.Kr
	(*Okr)(nil),               // 1: okrs.Okr
	(*GetOkrsRequest)(nil),    // 2: okrs.GetOkrsRequest
	(*GetOkrsResponse)(nil),   // 3: okrs.GetOkrsResponse
	(*GetOkrRequest)(nil),     // 4: okrs.GetOkrRequest
	(*GetOkrResponse)(nil),    // 5: okrs.GetOkrResponse
	(*PutOkrRequest)(nil),     // 6: okrs.PutOkrRequest
	(*PutOkrResponse)(nil),    // 7: okrs.PutOkrResponse
	(*DeleteOkrRequest)(nil),  // 8: okrs.DeleteOkrRequest
	(*DeleteOkrResponse)(nil), // 9: okrs.DeleteOkrResponse
	(*_type.DateTime)(nil),    // 10: google.type.DateTime
}
var file_okrs_proto_depIdxs = []int32{
	0,  // 0: okrs.Okr.embedded_child:type_name -> okrs.Kr
	1,  // 1: okrs.GetOkrsResponse.okrs:type_name -> okrs.Okr
	1,  // 2: okrs.GetOkrResponse.okr:type_name -> okrs.Okr
	10, // 3: okrs.PutOkrRequest.start_at:type_name -> google.type.DateTime
	1,  // 4: okrs.PutOkrResponse.okr:type_name -> okrs.Okr
	2,  // 5: okrs.OkrService.GetOkrs:input_type -> okrs.GetOkrsRequest
	4,  // 6: okrs.OkrService.GetOkr:input_type -> okrs.GetOkrRequest
	6,  // 7: okrs.OkrService.PutOkr:input_type -> okrs.PutOkrRequest
	8,  // 8: okrs.OkrService.DeleteOkr:input_type -> okrs.DeleteOkrRequest
	3,  // 9: okrs.OkrService.GetOkrs:output_type -> okrs.GetOkrsResponse
	5,  // 10: okrs.OkrService.GetOkr:output_type -> okrs.GetOkrResponse
	7,  // 11: okrs.OkrService.PutOkr:output_type -> okrs.PutOkrResponse
	9,  // 12: okrs.OkrService.DeleteOkr:output_type -> okrs.DeleteOkrResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_okrs_proto_init() }
func file_okrs_proto_init() {
	if File_okrs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_okrs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_okrs_proto_goTypes,
		DependencyIndexes: file_okrs_proto_depIdxs,
		MessageInfos:      file_okrs_proto_msgTypes,
	}.Build()
	File_okrs_proto = out.File
	file_okrs_proto_rawDesc = nil
	file_okrs_proto_goTypes = nil
	file_okrs_proto_depIdxs = nil
}
