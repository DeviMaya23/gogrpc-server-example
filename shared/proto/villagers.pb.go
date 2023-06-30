// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: villagers.proto

package proto

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

type Villager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Personality string `protobuf:"bytes,2,opt,name=Personality,proto3" json:"Personality,omitempty"`
}

func (x *Villager) Reset() {
	*x = Villager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_villagers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Villager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Villager) ProtoMessage() {}

func (x *Villager) ProtoReflect() protoreflect.Message {
	mi := &file_villagers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Villager.ProtoReflect.Descriptor instead.
func (*Villager) Descriptor() ([]byte, []int) {
	return file_villagers_proto_rawDescGZIP(), []int{0}
}

func (x *Villager) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Villager) GetPersonality() string {
	if x != nil {
		return x.Personality
	}
	return ""
}

type FindByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FindByNameRequest) Reset() {
	*x = FindByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_villagers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByNameRequest) ProtoMessage() {}

func (x *FindByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_villagers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByNameRequest.ProtoReflect.Descriptor instead.
func (*FindByNameRequest) Descriptor() ([]byte, []int) {
	return file_villagers_proto_rawDescGZIP(), []int{1}
}

func (x *FindByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Villagers []*Villager `protobuf:"bytes,1,rep,name=Villagers,proto3" json:"Villagers,omitempty"`
}

func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_villagers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_villagers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_villagers_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllResponse) GetVillagers() []*Villager {
	if x != nil {
		return x.Villagers
	}
	return nil
}

type FindStreamClientSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FindStreamClientSideRequest) Reset() {
	*x = FindStreamClientSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_villagers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStreamClientSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStreamClientSideRequest) ProtoMessage() {}

func (x *FindStreamClientSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_villagers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStreamClientSideRequest.ProtoReflect.Descriptor instead.
func (*FindStreamClientSideRequest) Descriptor() ([]byte, []int) {
	return file_villagers_proto_rawDescGZIP(), []int{3}
}

func (x *FindStreamClientSideRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindStreamClientSideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Villagers []*Villager `protobuf:"bytes,1,rep,name=Villagers,proto3" json:"Villagers,omitempty"`
}

func (x *FindStreamClientSideResponse) Reset() {
	*x = FindStreamClientSideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_villagers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStreamClientSideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStreamClientSideResponse) ProtoMessage() {}

func (x *FindStreamClientSideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_villagers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStreamClientSideResponse.ProtoReflect.Descriptor instead.
func (*FindStreamClientSideResponse) Descriptor() ([]byte, []int) {
	return file_villagers_proto_rawDescGZIP(), []int{4}
}

func (x *FindStreamClientSideResponse) GetVillagers() []*Villager {
	if x != nil {
		return x.Villagers
	}
	return nil
}

var File_villagers_proto protoreflect.FileDescriptor

var file_villagers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x08, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x40, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x52, 0x09, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x22, 0x31, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x52, 0x09, 0x56, 0x69, 0x6c, 0x6c, 0x61,
	0x67, 0x65, 0x72, 0x73, 0x32, 0xf6, 0x02, 0x0a, 0x10, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x17, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x72, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x17, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x69, 0x74,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x72, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_villagers_proto_rawDescOnce sync.Once
	file_villagers_proto_rawDescData = file_villagers_proto_rawDesc
)

func file_villagers_proto_rawDescGZIP() []byte {
	file_villagers_proto_rawDescOnce.Do(func() {
		file_villagers_proto_rawDescData = protoimpl.X.CompressGZIP(file_villagers_proto_rawDescData)
	})
	return file_villagers_proto_rawDescData
}

var file_villagers_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_villagers_proto_goTypes = []interface{}{
	(*Villager)(nil),                     // 0: proto.Villager
	(*FindByNameRequest)(nil),            // 1: proto.FindByNameRequest
	(*FindAllResponse)(nil),              // 2: proto.FindAllResponse
	(*FindStreamClientSideRequest)(nil),  // 3: proto.FindStreamClientSideRequest
	(*FindStreamClientSideResponse)(nil), // 4: proto.FindStreamClientSideResponse
	(*emptypb.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_villagers_proto_depIdxs = []int32{
	0, // 0: proto.FindAllResponse.Villagers:type_name -> proto.Villager
	0, // 1: proto.FindStreamClientSideResponse.Villagers:type_name -> proto.Villager
	5, // 2: proto.VillagersService.FindAll:input_type -> google.protobuf.Empty
	1, // 3: proto.VillagersService.FindByName:input_type -> proto.FindByNameRequest
	5, // 4: proto.VillagersService.FindAllStreamServerSide:input_type -> google.protobuf.Empty
	3, // 5: proto.VillagersService.FindStreamClientSide:input_type -> proto.FindStreamClientSideRequest
	3, // 6: proto.VillagersService.FindStreamBidirecitonal:input_type -> proto.FindStreamClientSideRequest
	2, // 7: proto.VillagersService.FindAll:output_type -> proto.FindAllResponse
	0, // 8: proto.VillagersService.FindByName:output_type -> proto.Villager
	0, // 9: proto.VillagersService.FindAllStreamServerSide:output_type -> proto.Villager
	2, // 10: proto.VillagersService.FindStreamClientSide:output_type -> proto.FindAllResponse
	0, // 11: proto.VillagersService.FindStreamBidirecitonal:output_type -> proto.Villager
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_villagers_proto_init() }
func file_villagers_proto_init() {
	if File_villagers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_villagers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Villager); i {
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
		file_villagers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByNameRequest); i {
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
		file_villagers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllResponse); i {
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
		file_villagers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStreamClientSideRequest); i {
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
		file_villagers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStreamClientSideResponse); i {
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
			RawDescriptor: file_villagers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_villagers_proto_goTypes,
		DependencyIndexes: file_villagers_proto_depIdxs,
		MessageInfos:      file_villagers_proto_msgTypes,
	}.Build()
	File_villagers_proto = out.File
	file_villagers_proto_rawDesc = nil
	file_villagers_proto_goTypes = nil
	file_villagers_proto_depIdxs = nil
}
