// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.2
// source: service_station.proto

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

type ServiceStation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Categories  []uint64 `protobuf:"varint,4,rep,packed,name=categories,proto3" json:"categories,omitempty"`
	Location    *Geo     `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Address     string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Images      []string `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	Phone       []string `protobuf:"bytes,8,rep,name=phone,proto3" json:"phone,omitempty"`
	Email       []string `protobuf:"bytes,9,rep,name=email,proto3" json:"email,omitempty"`
	Lat         float64  `protobuf:"fixed64,10,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon         float64  `protobuf:"fixed64,11,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *ServiceStation) Reset() {
	*x = ServiceStation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStation) ProtoMessage() {}

func (x *ServiceStation) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStation.ProtoReflect.Descriptor instead.
func (*ServiceStation) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceStation) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceStation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ServiceStation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceStation) GetCategories() []uint64 {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ServiceStation) GetLocation() *Geo {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ServiceStation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ServiceStation) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ServiceStation) GetPhone() []string {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *ServiceStation) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *ServiceStation) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *ServiceStation) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type ServiceStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServiceStationRequest) Reset() {
	*x = ServiceStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStationRequest) ProtoMessage() {}

func (x *ServiceStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStationRequest.ProtoReflect.Descriptor instead.
func (*ServiceStationRequest) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceStationRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ServiceStationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ServiceStationResponse) Reset() {
	*x = ServiceStationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStationResponse) ProtoMessage() {}

func (x *ServiceStationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStationResponse.ProtoReflect.Descriptor instead.
func (*ServiceStationResponse) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceStationResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceStationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FindServiceStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location   []*Geo   `protobuf:"bytes,1,rep,name=location,proto3" json:"location,omitempty"`
	Categories []string `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Limit      uint32   `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     uint32   `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FindServiceStationRequest) Reset() {
	*x = FindServiceStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServiceStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServiceStationRequest) ProtoMessage() {}

func (x *FindServiceStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServiceStationRequest.ProtoReflect.Descriptor instead.
func (*FindServiceStationRequest) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{3}
}

func (x *FindServiceStationRequest) GetLocation() []*Geo {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *FindServiceStationRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *FindServiceStationRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindServiceStationRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FindServiceStationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   string                             `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string                             `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Location    *FindServiceStationResult_ShortGeo `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Title       string                             `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Description string                             `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Categories  []string                           `protobuf:"bytes,11,rep,name=categories,proto3" json:"categories,omitempty"`
	Images      []string                           `protobuf:"bytes,12,rep,name=images,proto3" json:"images,omitempty"`
	Phone       []string                           `protobuf:"bytes,13,rep,name=phone,proto3" json:"phone,omitempty"`
	Email       []string                           `protobuf:"bytes,14,rep,name=email,proto3" json:"email,omitempty"`
	Owner       *FindServiceStationResult_Owner    `protobuf:"bytes,17,opt,name=owner,proto3" json:"owner,omitempty"`
	Likes       int32                              `protobuf:"varint,18,opt,name=likes,proto3" json:"likes,omitempty"`
	Dislikes    int32                              `protobuf:"varint,19,opt,name=dislikes,proto3" json:"dislikes,omitempty"`
}

func (x *FindServiceStationResult) Reset() {
	*x = FindServiceStationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServiceStationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServiceStationResult) ProtoMessage() {}

func (x *FindServiceStationResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServiceStationResult.ProtoReflect.Descriptor instead.
func (*FindServiceStationResult) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{4}
}

func (x *FindServiceStationResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindServiceStationResult) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FindServiceStationResult) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FindServiceStationResult) GetLocation() *FindServiceStationResult_ShortGeo {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *FindServiceStationResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FindServiceStationResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FindServiceStationResult) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *FindServiceStationResult) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *FindServiceStationResult) GetPhone() []string {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *FindServiceStationResult) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *FindServiceStationResult) GetOwner() *FindServiceStationResult_Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *FindServiceStationResult) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *FindServiceStationResult) GetDislikes() int32 {
	if x != nil {
		return x.Dislikes
	}
	return 0
}

type FindServiceStationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceStations []*FindServiceStationResult `protobuf:"bytes,1,rep,name=ServiceStations,proto3" json:"ServiceStations,omitempty"`
	Found           uint64                      `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *FindServiceStationResponse) Reset() {
	*x = FindServiceStationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServiceStationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServiceStationResponse) ProtoMessage() {}

func (x *FindServiceStationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServiceStationResponse.ProtoReflect.Descriptor instead.
func (*FindServiceStationResponse) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{5}
}

func (x *FindServiceStationResponse) GetServiceStations() []*FindServiceStationResult {
	if x != nil {
		return x.ServiceStations
	}
	return nil
}

func (x *FindServiceStationResponse) GetFound() uint64 {
	if x != nil {
		return x.Found
	}
	return 0
}

type FindServiceStationResult_ShortGeo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    uint32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Address string  `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Lat     float64 `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon     float64 `protobuf:"fixed64,6,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *FindServiceStationResult_ShortGeo) Reset() {
	*x = FindServiceStationResult_ShortGeo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServiceStationResult_ShortGeo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServiceStationResult_ShortGeo) ProtoMessage() {}

func (x *FindServiceStationResult_ShortGeo) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServiceStationResult_ShortGeo.ProtoReflect.Descriptor instead.
func (*FindServiceStationResult_ShortGeo) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{4, 0}
}

func (x *FindServiceStationResult_ShortGeo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindServiceStationResult_ShortGeo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindServiceStationResult_ShortGeo) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FindServiceStationResult_ShortGeo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *FindServiceStationResult_ShortGeo) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *FindServiceStationResult_ShortGeo) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type FindServiceStationResult_Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FindServiceStationResult_Owner) Reset() {
	*x = FindServiceStationResult_Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_station_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServiceStationResult_Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServiceStationResult_Owner) ProtoMessage() {}

func (x *FindServiceStationResult_Owner) ProtoReflect() protoreflect.Message {
	mi := &file_service_station_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServiceStationResult_Owner.ProtoReflect.Descriptor instead.
func (*FindServiceStationResult_Owner) Descriptor() ([]byte, []int) {
	return file_service_station_proto_rawDescGZIP(), []int{4, 1}
}

func (x *FindServiceStationResult_Owner) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindServiceStationResult_Owner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_service_station_proto protoreflect.FileDescriptor

var file_service_station_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0b, 0x63, 0x61,
	0x72, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x46, 0x69,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x6f, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xe5, 0x04, 0x0a, 0x18,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x6f,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x1a, 0x80, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x1a, 0x2b, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x32, 0xe7, 0x02, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_station_proto_rawDescOnce sync.Once
	file_service_station_proto_rawDescData = file_service_station_proto_rawDesc
)

func file_service_station_proto_rawDescGZIP() []byte {
	file_service_station_proto_rawDescOnce.Do(func() {
		file_service_station_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_station_proto_rawDescData)
	})
	return file_service_station_proto_rawDescData
}

var file_service_station_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_station_proto_goTypes = []interface{}{
	(*ServiceStation)(nil),                    // 0: api.ServiceStation
	(*ServiceStationRequest)(nil),             // 1: api.ServiceStationRequest
	(*ServiceStationResponse)(nil),            // 2: api.ServiceStationResponse
	(*FindServiceStationRequest)(nil),         // 3: api.FindServiceStationRequest
	(*FindServiceStationResult)(nil),          // 4: api.FindServiceStationResult
	(*FindServiceStationResponse)(nil),        // 5: api.FindServiceStationResponse
	(*FindServiceStationResult_ShortGeo)(nil), // 6: api.FindServiceStationResult.ShortGeo
	(*FindServiceStationResult_Owner)(nil),    // 7: api.FindServiceStationResult.Owner
	(*Geo)(nil),                               // 8: api.Geo
}
var file_service_station_proto_depIdxs = []int32{
	8,  // 0: api.ServiceStation.location:type_name -> api.Geo
	8,  // 1: api.FindServiceStationRequest.location:type_name -> api.Geo
	6,  // 2: api.FindServiceStationResult.location:type_name -> api.FindServiceStationResult.ShortGeo
	7,  // 3: api.FindServiceStationResult.owner:type_name -> api.FindServiceStationResult.Owner
	4,  // 4: api.FindServiceStationResponse.ServiceStations:type_name -> api.FindServiceStationResult
	1,  // 5: api.ServiceStationService.Get:input_type -> api.ServiceStationRequest
	3,  // 6: api.ServiceStationService.Find:input_type -> api.FindServiceStationRequest
	0,  // 7: api.ServiceStationService.Create:input_type -> api.ServiceStation
	0,  // 8: api.ServiceStationService.Update:input_type -> api.ServiceStation
	1,  // 9: api.ServiceStationService.Delete:input_type -> api.ServiceStationRequest
	4,  // 10: api.ServiceStationService.Get:output_type -> api.FindServiceStationResult
	5,  // 11: api.ServiceStationService.Find:output_type -> api.FindServiceStationResponse
	2,  // 12: api.ServiceStationService.Create:output_type -> api.ServiceStationResponse
	2,  // 13: api.ServiceStationService.Update:output_type -> api.ServiceStationResponse
	2,  // 14: api.ServiceStationService.Delete:output_type -> api.ServiceStationResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_service_station_proto_init() }
func file_service_station_proto_init() {
	if File_service_station_proto != nil {
		return
	}
	file_cargo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_station_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStation); i {
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
		file_service_station_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStationRequest); i {
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
		file_service_station_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStationResponse); i {
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
		file_service_station_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServiceStationRequest); i {
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
		file_service_station_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServiceStationResult); i {
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
		file_service_station_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServiceStationResponse); i {
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
		file_service_station_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServiceStationResult_ShortGeo); i {
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
		file_service_station_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServiceStationResult_Owner); i {
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
			RawDescriptor: file_service_station_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_station_proto_goTypes,
		DependencyIndexes: file_service_station_proto_depIdxs,
		MessageInfos:      file_service_station_proto_msgTypes,
	}.Build()
	File_service_station_proto = out.File
	file_service_station_proto_rawDesc = nil
	file_service_station_proto_goTypes = nil
	file_service_station_proto_depIdxs = nil
}
