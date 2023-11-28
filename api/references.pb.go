// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.2
// source: references.proto

package api

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

type RType int32

const (
	RType_transportType        RType = 0
	RType_loadingType          RType = 1
	RType_addition             RType = 2
	RType_serviceCategory      RType = 3
	RType_autoBrand            RType = 4
	RType_busBrand             RType = 5
	RType_specialBrand         RType = 6
	RType_busTransportType     RType = 7
	RType_specialTransportType RType = 8
	RType_cargoType            RType = 9
)

// Enum value maps for RType.
var (
	RType_name = map[int32]string{
		0: "transportType",
		1: "loadingType",
		2: "addition",
		3: "serviceCategory",
		4: "autoBrand",
		5: "busBrand",
		6: "specialBrand",
		7: "busTransportType",
		8: "specialTransportType",
		9: "cargoType",
	}
	RType_value = map[string]int32{
		"transportType":        0,
		"loadingType":          1,
		"addition":             2,
		"serviceCategory":      3,
		"autoBrand":            4,
		"busBrand":             5,
		"specialBrand":         6,
		"busTransportType":     7,
		"specialTransportType": 8,
		"cargoType":            9,
	}
)

func (x RType) Enum() *RType {
	p := new(RType)
	*p = x
	return p
}

func (x RType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RType) Descriptor() protoreflect.EnumDescriptor {
	return file_references_proto_enumTypes[0].Descriptor()
}

func (RType) Type() protoreflect.EnumType {
	return &file_references_proto_enumTypes[0]
}

func (x RType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RType.Descriptor instead.
func (RType) EnumDescriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{0}
}

type SortType int32

const (
	SortType_r_name     SortType = 0
	SortType_id         SortType = 1
	SortType_created_at SortType = 2
)

// Enum value maps for SortType.
var (
	SortType_name = map[int32]string{
		0: "r_name",
		1: "id",
		2: "created_at",
	}
	SortType_value = map[string]int32{
		"r_name":     0,
		"id":         1,
		"created_at": 2,
	}
)

func (x SortType) Enum() *SortType {
	p := new(SortType)
	*p = x
	return p
}

func (x SortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortType) Descriptor() protoreflect.EnumDescriptor {
	return file_references_proto_enumTypes[1].Descriptor()
}

func (SortType) Type() protoreflect.EnumType {
	return &file_references_proto_enumTypes[1]
}

func (x SortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortType.Descriptor instead.
func (SortType) EnumDescriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{1}
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type RType  `protobuf:"varint,3,opt,name=type,proto3,enum=api.RType" json:"type,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{0}
}

func (x *Reference) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Reference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reference) GetType() RType {
	if x != nil {
		return x.Type
	}
	return RType_transportType
}

type ReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type RType  `protobuf:"varint,2,opt,name=type,proto3,enum=api.RType" json:"type,omitempty"`
}

func (x *ReferenceRequest) Reset() {
	*x = ReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceRequest) ProtoMessage() {}

func (x *ReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceRequest.ProtoReflect.Descriptor instead.
func (*ReferenceRequest) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{1}
}

func (x *ReferenceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReferenceRequest) GetType() RType {
	if x != nil {
		return x.Type
	}
	return RType_transportType
}

type ReferenceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type RType    `protobuf:"varint,1,opt,name=type,proto3,enum=api.RType" json:"type,omitempty"`
	Sort SortType `protobuf:"varint,2,opt,name=sort,proto3,enum=api.SortType" json:"sort,omitempty"`
}

func (x *ReferenceListRequest) Reset() {
	*x = ReferenceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceListRequest) ProtoMessage() {}

func (x *ReferenceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceListRequest.ProtoReflect.Descriptor instead.
func (*ReferenceListRequest) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{2}
}

func (x *ReferenceListRequest) GetType() RType {
	if x != nil {
		return x.Type
	}
	return RType_transportType
}

func (x *ReferenceListRequest) GetSort() SortType {
	if x != nil {
		return x.Sort
	}
	return SortType_r_name
}

type ReferenceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference []*Reference `protobuf:"bytes,1,rep,name=reference,proto3" json:"reference,omitempty"`
	Found     uint32       `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *ReferenceListResponse) Reset() {
	*x = ReferenceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceListResponse) ProtoMessage() {}

func (x *ReferenceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceListResponse.ProtoReflect.Descriptor instead.
func (*ReferenceListResponse) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{3}
}

func (x *ReferenceListResponse) GetReference() []*Reference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *ReferenceListResponse) GetFound() uint32 {
	if x != nil {
		return x.Found
	}
	return 0
}

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positive   bool   `protobuf:"varint,1,opt,name=positive,proto3" json:"positive,omitempty"`
	EntityId   uint64 `protobuf:"varint,2,opt,name=entityId,proto3" json:"entityId,omitempty"`
	EntityType string `protobuf:"bytes,3,opt,name=entityType,proto3" json:"entityType,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{4}
}

func (x *Like) GetPositive() bool {
	if x != nil {
		return x.Positive
	}
	return false
}

func (x *Like) GetEntityId() uint64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *Like) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId   uint64 `protobuf:"varint,2,opt,name=entityId,proto3" json:"entityId,omitempty"`
	EntityType string `protobuf:"bytes,3,opt,name=entityType,proto3" json:"entityType,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{5}
}

func (x *LikeRequest) GetEntityId() uint64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *LikeRequest) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type LikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positive int64 `protobuf:"varint,1,opt,name=positive,proto3" json:"positive,omitempty"`
	Negative int64 `protobuf:"varint,2,opt,name=negative,proto3" json:"negative,omitempty"`
}

func (x *LikeResponse) Reset() {
	*x = LikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResponse) ProtoMessage() {}

func (x *LikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResponse.ProtoReflect.Descriptor instead.
func (*LikeResponse) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{6}
}

func (x *LikeResponse) GetPositive() int64 {
	if x != nil {
		return x.Positive
	}
	return 0
}

func (x *LikeResponse) GetNegative() int64 {
	if x != nil {
		return x.Negative
	}
	return 0
}

type O struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Bin          string `protobuf:"bytes,2,opt,name=Bin,proto3" json:"Bin,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	RegisterDate string `protobuf:"bytes,4,opt,name=RegisterDate,proto3" json:"RegisterDate,omitempty"`
	OkedCode     string `protobuf:"bytes,5,opt,name=OkedCode,proto3" json:"OkedCode,omitempty"`
	OkedName     string `protobuf:"bytes,6,opt,name=OkedName,proto3" json:"OkedName,omitempty"`
	SecondOkeds  string `protobuf:"bytes,7,opt,name=SecondOkeds,proto3" json:"SecondOkeds,omitempty"`
	KrpCode      string `protobuf:"bytes,8,opt,name=KrpCode,proto3" json:"KrpCode,omitempty"`
	KrpName      string `protobuf:"bytes,9,opt,name=KrpName,proto3" json:"KrpName,omitempty"`
	KrpBfCode    string `protobuf:"bytes,10,opt,name=KrpBfCode,proto3" json:"KrpBfCode,omitempty"`
	KrpBfName    string `protobuf:"bytes,11,opt,name=KrpBfName,proto3" json:"KrpBfName,omitempty"`
	KatoCode     string `protobuf:"bytes,12,opt,name=KatoCode,proto3" json:"KatoCode,omitempty"`
	KatoId       int64  `protobuf:"varint,13,opt,name=KatoId,proto3" json:"KatoId,omitempty"`
	KatoAddress  string `protobuf:"bytes,14,opt,name=KatoAddress,proto3" json:"KatoAddress,omitempty"`
	Fio          string `protobuf:"bytes,15,opt,name=Fio,proto3" json:"Fio,omitempty"`
	Ip           string `protobuf:"bytes,16,opt,name=Ip,proto3" json:"Ip,omitempty"`
}

func (x *O) Reset() {
	*x = O{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *O) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*O) ProtoMessage() {}

func (x *O) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use O.ProtoReflect.Descriptor instead.
func (*O) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{7}
}

func (x *O) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *O) GetBin() string {
	if x != nil {
		return x.Bin
	}
	return ""
}

func (x *O) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *O) GetRegisterDate() string {
	if x != nil {
		return x.RegisterDate
	}
	return ""
}

func (x *O) GetOkedCode() string {
	if x != nil {
		return x.OkedCode
	}
	return ""
}

func (x *O) GetOkedName() string {
	if x != nil {
		return x.OkedName
	}
	return ""
}

func (x *O) GetSecondOkeds() string {
	if x != nil {
		return x.SecondOkeds
	}
	return ""
}

func (x *O) GetKrpCode() string {
	if x != nil {
		return x.KrpCode
	}
	return ""
}

func (x *O) GetKrpName() string {
	if x != nil {
		return x.KrpName
	}
	return ""
}

func (x *O) GetKrpBfCode() string {
	if x != nil {
		return x.KrpBfCode
	}
	return ""
}

func (x *O) GetKrpBfName() string {
	if x != nil {
		return x.KrpBfName
	}
	return ""
}

func (x *O) GetKatoCode() string {
	if x != nil {
		return x.KatoCode
	}
	return ""
}

func (x *O) GetKatoId() int64 {
	if x != nil {
		return x.KatoId
	}
	return 0
}

func (x *O) GetKatoAddress() string {
	if x != nil {
		return x.KatoAddress
	}
	return ""
}

func (x *O) GetFio() string {
	if x != nil {
		return x.Fio
	}
	return ""
}

func (x *O) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type StatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Obj         *O     `protobuf:"bytes,2,opt,name=Obj,proto3" json:"Obj,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *StatResponse) Reset() {
	*x = StatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponse) ProtoMessage() {}

func (x *StatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponse.ProtoReflect.Descriptor instead.
func (*StatResponse) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{8}
}

func (x *StatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StatResponse) GetObj() *O {
	if x != nil {
		return x.Obj
	}
	return nil
}

func (x *StatResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type StatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bin string `protobuf:"bytes,1,opt,name=bin,proto3" json:"bin,omitempty"`
}

func (x *StatRequest) Reset() {
	*x = StatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_references_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatRequest) ProtoMessage() {}

func (x *StatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_references_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatRequest.ProtoReflect.Descriptor instead.
func (*StatRequest) Descriptor() ([]byte, []int) {
	return file_references_proto_rawDescGZIP(), []int{9}
}

func (x *StatRequest) GetBin() string {
	if x != nil {
		return x.Bin
	}
	return ""
}

var File_references_proto protoreflect.FileDescriptor

var file_references_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x50, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x5a, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x5b,
	0x0a, 0x15, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x5e, 0x0a, 0x04, 0x4c,
	0x69, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x4c,
	0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x22, 0x9f,
	0x03, 0x0a, 0x01, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x42, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x6b,
	0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x6b,
	0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4f, 0x6b, 0x65, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x4f, 0x6b, 0x65, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x72, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x72, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x72, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x72, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x4b, 0x72, 0x70, 0x42, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4b, 0x72, 0x70, 0x42, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4b, 0x72,
	0x70, 0x42, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4b,
	0x72, 0x70, 0x42, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x61, 0x74, 0x6f,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x61, 0x74, 0x6f,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x61, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4b, 0x61, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x4b, 0x61, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4b, 0x61, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x46, 0x69, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x69, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70,
	0x22, 0x64, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x03, 0x4f, 0x62,
	0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x52,
	0x03, 0x4f, 0x62, 0x6a, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6e, 0x2a, 0xbd, 0x01, 0x0a, 0x06, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x62, 0x75, 0x73, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x07, 0x12, 0x18,
	0x0a, 0x14, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x10, 0x09, 0x2a, 0x2e, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x10, 0x02, 0x32, 0x90, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x46,
	0x69, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_references_proto_rawDescOnce sync.Once
	file_references_proto_rawDescData = file_references_proto_rawDesc
)

func file_references_proto_rawDescGZIP() []byte {
	file_references_proto_rawDescOnce.Do(func() {
		file_references_proto_rawDescData = protoimpl.X.CompressGZIP(file_references_proto_rawDescData)
	})
	return file_references_proto_rawDescData
}

var file_references_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_references_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_references_proto_goTypes = []interface{}{
	(RType)(0),                    // 0: api.r_type
	(SortType)(0),                 // 1: api.sortType
	(*Reference)(nil),             // 2: api.Reference
	(*ReferenceRequest)(nil),      // 3: api.ReferenceRequest
	(*ReferenceListRequest)(nil),  // 4: api.ReferenceListRequest
	(*ReferenceListResponse)(nil), // 5: api.ReferenceListResponse
	(*Like)(nil),                  // 6: api.Like
	(*LikeRequest)(nil),           // 7: api.LikeRequest
	(*LikeResponse)(nil),          // 8: api.LikeResponse
	(*O)(nil),                     // 9: api.o
	(*StatResponse)(nil),          // 10: api.StatResponse
	(*StatRequest)(nil),           // 11: api.StatRequest
}
var file_references_proto_depIdxs = []int32{
	0,  // 0: api.Reference.type:type_name -> api.r_type
	0,  // 1: api.ReferenceRequest.type:type_name -> api.r_type
	0,  // 2: api.ReferenceListRequest.type:type_name -> api.r_type
	1,  // 3: api.ReferenceListRequest.sort:type_name -> api.sortType
	2,  // 4: api.ReferenceListResponse.reference:type_name -> api.Reference
	9,  // 5: api.StatResponse.Obj:type_name -> api.o
	11, // 6: api.ReferenceService.Stat:input_type -> api.StatRequest
	3,  // 7: api.ReferenceService.Get:input_type -> api.ReferenceRequest
	4,  // 8: api.ReferenceService.Find:input_type -> api.ReferenceListRequest
	6,  // 9: api.ReferenceService.AddLike:input_type -> api.Like
	7,  // 10: api.ReferenceService.GetLikes:input_type -> api.LikeRequest
	10, // 11: api.ReferenceService.Stat:output_type -> api.StatResponse
	2,  // 12: api.ReferenceService.Get:output_type -> api.Reference
	5,  // 13: api.ReferenceService.Find:output_type -> api.ReferenceListResponse
	8,  // 14: api.ReferenceService.AddLike:output_type -> api.LikeResponse
	8,  // 15: api.ReferenceService.GetLikes:output_type -> api.LikeResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_references_proto_init() }
func file_references_proto_init() {
	if File_references_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_references_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_references_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceRequest); i {
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
		file_references_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceListRequest); i {
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
		file_references_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceListResponse); i {
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
		file_references_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_references_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
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
		file_references_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeResponse); i {
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
		file_references_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*O); i {
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
		file_references_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatResponse); i {
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
		file_references_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatRequest); i {
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
			RawDescriptor: file_references_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_references_proto_goTypes,
		DependencyIndexes: file_references_proto_depIdxs,
		EnumInfos:         file_references_proto_enumTypes,
		MessageInfos:      file_references_proto_msgTypes,
	}.Build()
	File_references_proto = out.File
	file_references_proto_rawDesc = nil
	file_references_proto_goTypes = nil
	file_references_proto_depIdxs = nil
}