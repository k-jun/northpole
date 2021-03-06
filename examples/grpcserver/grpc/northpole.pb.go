// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: northpole.proto

package grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RoomStatus int32

const (
	RoomStatus_Open  RoomStatus = 0
	RoomStatus_Close RoomStatus = 1
)

// Enum value maps for RoomStatus.
var (
	RoomStatus_name = map[int32]string{
		0: "Open",
		1: "Close",
	}
	RoomStatus_value = map[string]int32{
		"Open":  0,
		"Close": 1,
	}
)

func (x RoomStatus) Enum() *RoomStatus {
	p := new(RoomStatus)
	*p = x
	return p
}

func (x RoomStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_northpole_proto_enumTypes[0].Descriptor()
}

func (RoomStatus) Type() protoreflect.EnumType {
	return &file_northpole_proto_enumTypes[0]
}

func (x RoomStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomStatus.Descriptor instead.
func (RoomStatus) EnumDescriptor() ([]byte, []int) {
	return file_northpole_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_northpole_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_northpole_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_northpole_proto_rawDescGZIP(), []int{0}
}

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               RoomStatus `protobuf:"varint,2,opt,name=status,proto3,enum=RoomStatus" json:"status,omitempty"`
	CurrentNumberOfUsers int64      `protobuf:"varint,3,opt,name=current_number_of_users,json=currentNumberOfUsers,proto3" json:"current_number_of_users,omitempty"`
	MaxNumberOfUsers     int64      `protobuf:"varint,4,opt,name=max_number_of_users,json=maxNumberOfUsers,proto3" json:"max_number_of_users,omitempty"`
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_northpole_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_northpole_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_northpole_proto_rawDescGZIP(), []int{1}
}

func (x *RoomInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomInfo) GetStatus() RoomStatus {
	if x != nil {
		return x.Status
	}
	return RoomStatus_Open
}

func (x *RoomInfo) GetCurrentNumberOfUsers() int64 {
	if x != nil {
		return x.CurrentNumberOfUsers
	}
	return 0
}

func (x *RoomInfo) GetMaxNumberOfUsers() int64 {
	if x != nil {
		return x.MaxNumberOfUsers
	}
	return 0
}

type RoomCreateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId           string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId           string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MaxNumberOfUsers int64  `protobuf:"varint,3,opt,name=max_number_of_users,json=maxNumberOfUsers,proto3" json:"max_number_of_users,omitempty"`
}

func (x *RoomCreateInfo) Reset() {
	*x = RoomCreateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_northpole_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateInfo) ProtoMessage() {}

func (x *RoomCreateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_northpole_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateInfo.ProtoReflect.Descriptor instead.
func (*RoomCreateInfo) Descriptor() ([]byte, []int) {
	return file_northpole_proto_rawDescGZIP(), []int{2}
}

func (x *RoomCreateInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RoomCreateInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoomCreateInfo) GetMaxNumberOfUsers() int64 {
	if x != nil {
		return x.MaxNumberOfUsers
	}
	return 0
}

type MatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MatchInfo) Reset() {
	*x = MatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_northpole_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchInfo) ProtoMessage() {}

func (x *MatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_northpole_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchInfo.ProtoReflect.Descriptor instead.
func (*MatchInfo) Descriptor() ([]byte, []int) {
	return file_northpole_proto_rawDescGZIP(), []int{3}
}

func (x *MatchInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *MatchInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_northpole_proto protoreflect.FileDescriptor

var file_northpole_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x70, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa5, 0x01, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x17,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f,
	0x66, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x71, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x2a, 0x21, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x10, 0x01, 0x32, 0xbe, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x72, 0x74,
	0x68, 0x50, 0x6f, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0f, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0a, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x2d, 0x6a, 0x75, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_northpole_proto_rawDescOnce sync.Once
	file_northpole_proto_rawDescData = file_northpole_proto_rawDesc
)

func file_northpole_proto_rawDescGZIP() []byte {
	file_northpole_proto_rawDescOnce.Do(func() {
		file_northpole_proto_rawDescData = protoimpl.X.CompressGZIP(file_northpole_proto_rawDescData)
	})
	return file_northpole_proto_rawDescData
}

var file_northpole_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_northpole_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_northpole_proto_goTypes = []interface{}{
	(RoomStatus)(0),        // 0: RoomStatus
	(*Empty)(nil),          // 1: Empty
	(*RoomInfo)(nil),       // 2: RoomInfo
	(*RoomCreateInfo)(nil), // 3: RoomCreateInfo
	(*MatchInfo)(nil),      // 4: MatchInfo
}
var file_northpole_proto_depIdxs = []int32{
	0, // 0: RoomInfo.status:type_name -> RoomStatus
	4, // 1: NorthPole.JoinPublicRoom:input_type -> MatchInfo
	3, // 2: NorthPole.CreatePrivateRoom:input_type -> RoomCreateInfo
	4, // 3: NorthPole.JoinPrivateRoom:input_type -> MatchInfo
	4, // 4: NorthPole.LeaveRoom:input_type -> MatchInfo
	2, // 5: NorthPole.JoinPublicRoom:output_type -> RoomInfo
	2, // 6: NorthPole.CreatePrivateRoom:output_type -> RoomInfo
	2, // 7: NorthPole.JoinPrivateRoom:output_type -> RoomInfo
	1, // 8: NorthPole.LeaveRoom:output_type -> Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_northpole_proto_init() }
func file_northpole_proto_init() {
	if File_northpole_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_northpole_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_northpole_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
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
		file_northpole_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreateInfo); i {
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
		file_northpole_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchInfo); i {
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
			RawDescriptor: file_northpole_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_northpole_proto_goTypes,
		DependencyIndexes: file_northpole_proto_depIdxs,
		EnumInfos:         file_northpole_proto_enumTypes,
		MessageInfos:      file_northpole_proto_msgTypes,
	}.Build()
	File_northpole_proto = out.File
	file_northpole_proto_rawDesc = nil
	file_northpole_proto_goTypes = nil
	file_northpole_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NorthPoleClient is the client API for NorthPole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NorthPoleClient interface {
	JoinPublicRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (NorthPole_JoinPublicRoomClient, error)
	CreatePrivateRoom(ctx context.Context, in *RoomCreateInfo, opts ...grpc.CallOption) (NorthPole_CreatePrivateRoomClient, error)
	JoinPrivateRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (NorthPole_JoinPrivateRoomClient, error)
	LeaveRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (*Empty, error)
}

type northPoleClient struct {
	cc grpc.ClientConnInterface
}

func NewNorthPoleClient(cc grpc.ClientConnInterface) NorthPoleClient {
	return &northPoleClient{cc}
}

func (c *northPoleClient) JoinPublicRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (NorthPole_JoinPublicRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NorthPole_serviceDesc.Streams[0], "/NorthPole/JoinPublicRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &northPoleJoinPublicRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NorthPole_JoinPublicRoomClient interface {
	Recv() (*RoomInfo, error)
	grpc.ClientStream
}

type northPoleJoinPublicRoomClient struct {
	grpc.ClientStream
}

func (x *northPoleJoinPublicRoomClient) Recv() (*RoomInfo, error) {
	m := new(RoomInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *northPoleClient) CreatePrivateRoom(ctx context.Context, in *RoomCreateInfo, opts ...grpc.CallOption) (NorthPole_CreatePrivateRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NorthPole_serviceDesc.Streams[1], "/NorthPole/CreatePrivateRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &northPoleCreatePrivateRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NorthPole_CreatePrivateRoomClient interface {
	Recv() (*RoomInfo, error)
	grpc.ClientStream
}

type northPoleCreatePrivateRoomClient struct {
	grpc.ClientStream
}

func (x *northPoleCreatePrivateRoomClient) Recv() (*RoomInfo, error) {
	m := new(RoomInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *northPoleClient) JoinPrivateRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (NorthPole_JoinPrivateRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NorthPole_serviceDesc.Streams[2], "/NorthPole/JoinPrivateRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &northPoleJoinPrivateRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NorthPole_JoinPrivateRoomClient interface {
	Recv() (*RoomInfo, error)
	grpc.ClientStream
}

type northPoleJoinPrivateRoomClient struct {
	grpc.ClientStream
}

func (x *northPoleJoinPrivateRoomClient) Recv() (*RoomInfo, error) {
	m := new(RoomInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *northPoleClient) LeaveRoom(ctx context.Context, in *MatchInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/NorthPole/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NorthPoleServer is the server API for NorthPole service.
type NorthPoleServer interface {
	JoinPublicRoom(*MatchInfo, NorthPole_JoinPublicRoomServer) error
	CreatePrivateRoom(*RoomCreateInfo, NorthPole_CreatePrivateRoomServer) error
	JoinPrivateRoom(*MatchInfo, NorthPole_JoinPrivateRoomServer) error
	LeaveRoom(context.Context, *MatchInfo) (*Empty, error)
}

// UnimplementedNorthPoleServer can be embedded to have forward compatible implementations.
type UnimplementedNorthPoleServer struct {
}

func (*UnimplementedNorthPoleServer) JoinPublicRoom(*MatchInfo, NorthPole_JoinPublicRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinPublicRoom not implemented")
}
func (*UnimplementedNorthPoleServer) CreatePrivateRoom(*RoomCreateInfo, NorthPole_CreatePrivateRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePrivateRoom not implemented")
}
func (*UnimplementedNorthPoleServer) JoinPrivateRoom(*MatchInfo, NorthPole_JoinPrivateRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinPrivateRoom not implemented")
}
func (*UnimplementedNorthPoleServer) LeaveRoom(context.Context, *MatchInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}

func RegisterNorthPoleServer(s *grpc.Server, srv NorthPoleServer) {
	s.RegisterService(&_NorthPole_serviceDesc, srv)
}

func _NorthPole_JoinPublicRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NorthPoleServer).JoinPublicRoom(m, &northPoleJoinPublicRoomServer{stream})
}

type NorthPole_JoinPublicRoomServer interface {
	Send(*RoomInfo) error
	grpc.ServerStream
}

type northPoleJoinPublicRoomServer struct {
	grpc.ServerStream
}

func (x *northPoleJoinPublicRoomServer) Send(m *RoomInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _NorthPole_CreatePrivateRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomCreateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NorthPoleServer).CreatePrivateRoom(m, &northPoleCreatePrivateRoomServer{stream})
}

type NorthPole_CreatePrivateRoomServer interface {
	Send(*RoomInfo) error
	grpc.ServerStream
}

type northPoleCreatePrivateRoomServer struct {
	grpc.ServerStream
}

func (x *northPoleCreatePrivateRoomServer) Send(m *RoomInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _NorthPole_JoinPrivateRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NorthPoleServer).JoinPrivateRoom(m, &northPoleJoinPrivateRoomServer{stream})
}

type NorthPole_JoinPrivateRoomServer interface {
	Send(*RoomInfo) error
	grpc.ServerStream
}

type northPoleJoinPrivateRoomServer struct {
	grpc.ServerStream
}

func (x *northPoleJoinPrivateRoomServer) Send(m *RoomInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _NorthPole_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NorthPoleServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NorthPole/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NorthPoleServer).LeaveRoom(ctx, req.(*MatchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _NorthPole_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NorthPole",
	HandlerType: (*NorthPoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaveRoom",
			Handler:    _NorthPole_LeaveRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinPublicRoom",
			Handler:       _NorthPole_JoinPublicRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreatePrivateRoom",
			Handler:       _NorthPole_CreatePrivateRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "JoinPrivateRoom",
			Handler:       _NorthPole_JoinPrivateRoom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "northpole.proto",
}
