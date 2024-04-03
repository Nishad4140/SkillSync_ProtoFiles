// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: project.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateGigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	FreelancerId  string  `protobuf:"bytes,2,opt,name=freelancerId,proto3" json:"freelancerId,omitempty"`
	CategoryId    int32   `protobuf:"varint,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	SkillId       int32   `protobuf:"varint,4,opt,name=skillId,proto3" json:"skillId,omitempty"`
	Description   string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PackageTypeId int32   `protobuf:"varint,6,opt,name=packageTypeId,proto3" json:"packageTypeId,omitempty"`
	Price         float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	DelivaryDays  string  `protobuf:"bytes,8,opt,name=delivaryDays,proto3" json:"delivaryDays,omitempty"`
}

func (x *CreateGigRequest) Reset() {
	*x = CreateGigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGigRequest) ProtoMessage() {}

func (x *CreateGigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGigRequest.ProtoReflect.Descriptor instead.
func (*CreateGigRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGigRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateGigRequest) GetFreelancerId() string {
	if x != nil {
		return x.FreelancerId
	}
	return ""
}

func (x *CreateGigRequest) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateGigRequest) GetSkillId() int32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *CreateGigRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGigRequest) GetPackageTypeId() int32 {
	if x != nil {
		return x.PackageTypeId
	}
	return 0
}

func (x *CreateGigRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateGigRequest) GetDelivaryDays() string {
	if x != nil {
		return x.DelivaryDays
	}
	return ""
}

type GigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	FreelancerId  string   `protobuf:"bytes,3,opt,name=freelancerId,proto3" json:"freelancerId,omitempty"`
	CategoryId    int32    `protobuf:"varint,4,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	SkillId       int32    `protobuf:"varint,5,opt,name=skillId,proto3" json:"skillId,omitempty"`
	Description   string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	PackageTypeId int32    `protobuf:"varint,7,opt,name=packageTypeId,proto3" json:"packageTypeId,omitempty"`
	Price         float32  `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	DelivaryDays  string   `protobuf:"bytes,9,opt,name=delivaryDays,proto3" json:"delivaryDays,omitempty"`
	Images        []string `protobuf:"bytes,10,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *GigResponse) Reset() {
	*x = GigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GigResponse) ProtoMessage() {}

func (x *GigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GigResponse.ProtoReflect.Descriptor instead.
func (*GigResponse) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *GigResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GigResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GigResponse) GetFreelancerId() string {
	if x != nil {
		return x.FreelancerId
	}
	return ""
}

func (x *GigResponse) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GigResponse) GetSkillId() int32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *GigResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GigResponse) GetPackageTypeId() int32 {
	if x != nil {
		return x.PackageTypeId
	}
	return 0
}

func (x *GigResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GigResponse) GetDelivaryDays() string {
	if x != nil {
		return x.DelivaryDays
	}
	return ""
}

func (x *GigResponse) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type GetById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetById) Reset() {
	*x = GetById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetById) ProtoMessage() {}

func (x *GetById) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetById.ProtoReflect.Descriptor instead.
func (*GetById) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *GetById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByUserId) Reset() {
	*x = GetByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserId) ProtoMessage() {}

func (x *GetByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserId.ProtoReflect.Descriptor instead.
func (*GetByUserId) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{3}
}

func (x *GetByUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddPackageTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageType string `protobuf:"bytes,1,opt,name=packageType,proto3" json:"packageType,omitempty"`
}

func (x *AddPackageTypeRequest) Reset() {
	*x = AddPackageTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPackageTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPackageTypeRequest) ProtoMessage() {}

func (x *AddPackageTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPackageTypeRequest.ProtoReflect.Descriptor instead.
func (*AddPackageTypeRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{4}
}

func (x *AddPackageTypeRequest) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

type PackageTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PackageType string `protobuf:"bytes,2,opt,name=packageType,proto3" json:"packageType,omitempty"`
}

func (x *PackageTypeResponse) Reset() {
	*x = PackageTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageTypeResponse) ProtoMessage() {}

func (x *PackageTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageTypeResponse.ProtoReflect.Descriptor instead.
func (*PackageTypeResponse) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *PackageTypeResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PackageTypeResponse) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

type PackageTypeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PackageTypeId) Reset() {
	*x = PackageTypeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageTypeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageTypeId) ProtoMessage() {}

func (x *PackageTypeId) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageTypeId.ProtoReflect.Descriptor instead.
func (*PackageTypeId) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{6}
}

func (x *PackageTypeId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddClientGigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CategoryId   string `protobuf:"bytes,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	SkillId      string `protobuf:"bytes,4,opt,name=skillId,proto3" json:"skillId,omitempty"`
	Price        string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Description  string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	DeliveryDate string `protobuf:"bytes,7,opt,name=deliveryDate,proto3" json:"deliveryDate,omitempty"`
}

func (x *AddClientGigRequest) Reset() {
	*x = AddClientGigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientGigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientGigRequest) ProtoMessage() {}

func (x *AddClientGigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientGigRequest.ProtoReflect.Descriptor instead.
func (*AddClientGigRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{7}
}

func (x *AddClientGigRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AddClientGigRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddClientGigRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *AddClientGigRequest) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *AddClientGigRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *AddClientGigRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddClientGigRequest) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

type ClientRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId     string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Title        string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	CategoryId   string   `protobuf:"bytes,4,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	SkillId      string   `protobuf:"bytes,5,opt,name=skillId,proto3" json:"skillId,omitempty"`
	Price        string   `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	Description  string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	DeliveryDate string   `protobuf:"bytes,8,opt,name=deliveryDate,proto3" json:"deliveryDate,omitempty"`
	Images       []string `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ClientRequestResponse) Reset() {
	*x = ClientRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequestResponse) ProtoMessage() {}

func (x *ClientRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequestResponse.ProtoReflect.Descriptor instead.
func (*ClientRequestResponse) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{8}
}

func (x *ClientRequestResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientRequestResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientRequestResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ClientRequestResponse) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *ClientRequestResponse) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *ClientRequestResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ClientRequestResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ClientRequestResponse) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *ClientRequestResponse) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x61, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x61, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73,
	0x22, 0xab, 0x02, 0x0a, 0x0b, 0x47, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x79, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x61,
	0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x19,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x1f, 0x0a, 0x0d,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0x87, 0x02,
	0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x32, 0x8e, 0x06, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x47, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x12,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x46, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x47, 0x69, 0x67, 0x73, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x47, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x10, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_project_proto_goTypes = []interface{}{
	(*CreateGigRequest)(nil),      // 0: project.CreateGigRequest
	(*GigResponse)(nil),           // 1: project.GigResponse
	(*GetById)(nil),               // 2: project.GetById
	(*GetByUserId)(nil),           // 3: project.GetByUserId
	(*AddPackageTypeRequest)(nil), // 4: project.AddPackageTypeRequest
	(*PackageTypeResponse)(nil),   // 5: project.PackageTypeResponse
	(*PackageTypeId)(nil),         // 6: project.PackageTypeId
	(*AddClientGigRequest)(nil),   // 7: project.AddClientGigRequest
	(*ClientRequestResponse)(nil), // 8: project.ClientRequestResponse
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_project_proto_depIdxs = []int32{
	0,  // 0: project.ProjectService.CreateGig:input_type -> project.CreateGigRequest
	1,  // 1: project.ProjectService.UpdateGig:input_type -> project.GigResponse
	2,  // 2: project.ProjectService.GetGig:input_type -> project.GetById
	3,  // 3: project.ProjectService.GetAllFreelancerGigs:input_type -> project.GetByUserId
	4,  // 4: project.ProjectService.AddPackageType:input_type -> project.AddPackageTypeRequest
	5,  // 5: project.ProjectService.EditPackageType:input_type -> project.PackageTypeResponse
	9,  // 6: project.ProjectService.GetPackageType:input_type -> google.protobuf.Empty
	7,  // 7: project.ProjectService.ClientAddRequest:input_type -> project.AddClientGigRequest
	8,  // 8: project.ProjectService.ClientUpdateRequest:input_type -> project.ClientRequestResponse
	2,  // 9: project.ProjectService.GetClientRequest:input_type -> project.GetById
	3,  // 10: project.ProjectService.GetAllClientRequest:input_type -> project.GetByUserId
	9,  // 11: project.ProjectService.CreateGig:output_type -> google.protobuf.Empty
	9,  // 12: project.ProjectService.UpdateGig:output_type -> google.protobuf.Empty
	1,  // 13: project.ProjectService.GetGig:output_type -> project.GigResponse
	1,  // 14: project.ProjectService.GetAllFreelancerGigs:output_type -> project.GigResponse
	9,  // 15: project.ProjectService.AddPackageType:output_type -> google.protobuf.Empty
	9,  // 16: project.ProjectService.EditPackageType:output_type -> google.protobuf.Empty
	5,  // 17: project.ProjectService.GetPackageType:output_type -> project.PackageTypeResponse
	9,  // 18: project.ProjectService.ClientAddRequest:output_type -> google.protobuf.Empty
	9,  // 19: project.ProjectService.ClientUpdateRequest:output_type -> google.protobuf.Empty
	8,  // 20: project.ProjectService.GetClientRequest:output_type -> project.ClientRequestResponse
	8,  // 21: project.ProjectService.GetAllClientRequest:output_type -> project.ClientRequestResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGigRequest); i {
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
		file_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GigResponse); i {
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
		file_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetById); i {
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
		file_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserId); i {
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
		file_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPackageTypeRequest); i {
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
		file_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageTypeResponse); i {
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
		file_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageTypeId); i {
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
		file_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientGigRequest); i {
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
		file_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequestResponse); i {
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
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
