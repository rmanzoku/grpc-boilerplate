// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: healthcheck.proto

package healthcheck

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MessageOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MessageOutput) Reset() {
	*x = MessageOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOutput) ProtoMessage() {}

func (x *MessageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOutput.ProtoReflect.Descriptor instead.
func (*MessageOutput) Descriptor() ([]byte, []int) {
	return file_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOutput) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[1]
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
	return file_healthcheck_proto_rawDescGZIP(), []int{1}
}

var File_healthcheck_proto protoreflect.FileDescriptor

var file_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21,
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x69, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x12, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6d, 0x61, 0x6e, 0x7a, 0x6f, 0x6b, 0x75, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_healthcheck_proto_rawDescOnce sync.Once
	file_healthcheck_proto_rawDescData = file_healthcheck_proto_rawDesc
)

func file_healthcheck_proto_rawDescGZIP() []byte {
	file_healthcheck_proto_rawDescOnce.Do(func() {
		file_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_healthcheck_proto_rawDescData)
	})
	return file_healthcheck_proto_rawDescData
}

var file_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_healthcheck_proto_goTypes = []interface{}{
	(*MessageOutput)(nil), // 0: healthcheck.MessageOutput
	(*Empty)(nil),         // 1: healthcheck.Empty
}
var file_healthcheck_proto_depIdxs = []int32{
	1, // 0: healthcheck.HealthcheckService.Healthcheck:input_type -> healthcheck.Empty
	0, // 1: healthcheck.HealthcheckService.Healthcheck:output_type -> healthcheck.MessageOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_healthcheck_proto_init() }
func file_healthcheck_proto_init() {
	if File_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOutput); i {
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
		file_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_healthcheck_proto_goTypes,
		DependencyIndexes: file_healthcheck_proto_depIdxs,
		MessageInfos:      file_healthcheck_proto_msgTypes,
	}.Build()
	File_healthcheck_proto = out.File
	file_healthcheck_proto_rawDesc = nil
	file_healthcheck_proto_goTypes = nil
	file_healthcheck_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthcheckServiceClient is the client API for HealthcheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthcheckServiceClient interface {
	Healthcheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MessageOutput, error)
}

type healthcheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckServiceClient(cc grpc.ClientConnInterface) HealthcheckServiceClient {
	return &healthcheckServiceClient{cc}
}

func (c *healthcheckServiceClient) Healthcheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MessageOutput, error) {
	out := new(MessageOutput)
	err := c.cc.Invoke(ctx, "/healthcheck.HealthcheckService/Healthcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckServiceServer is the server API for HealthcheckService service.
type HealthcheckServiceServer interface {
	Healthcheck(context.Context, *Empty) (*MessageOutput, error)
}

// UnimplementedHealthcheckServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthcheckServiceServer struct {
}

func (*UnimplementedHealthcheckServiceServer) Healthcheck(context.Context, *Empty) (*MessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthcheck not implemented")
}

func RegisterHealthcheckServiceServer(s *grpc.Server, srv HealthcheckServiceServer) {
	s.RegisterService(&_HealthcheckService_serviceDesc, srv)
}

func _HealthcheckService_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServiceServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.HealthcheckService/Healthcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServiceServer).Healthcheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthcheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "healthcheck.HealthcheckService",
	HandlerType: (*HealthcheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _HealthcheckService_Healthcheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthcheck.proto",
}
