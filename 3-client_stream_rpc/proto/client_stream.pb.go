// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_stream.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 定义流式请求信息
type StreamRequest struct {
	// 参数类型 参数名称 标识号
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_729fd3a5e8f9e1ff, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// 定义响应信息
type SimpleResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_729fd3a5e8f9e1ff, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SimpleResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*SimpleResponse)(nil), "proto.SimpleResponse")
}

func init() { proto.RegisterFile("client_stream.proto", fileDescriptor_729fd3a5e8f9e1ff) }

var fileDescriptor_729fd3a5e8f9e1ff = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x89, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c, 0xbc, 0xc1, 0x60, 0xe1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x5b, 0xc9, 0x8a, 0x8b, 0x2f, 0x38, 0x33, 0xb7, 0x20, 0x27, 0x35, 0x28, 0xb5, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x15, 0xa4, 0x2a, 0x39, 0x3f, 0x25, 0x15, 0xac, 0x8a, 0x35, 0x08, 0xcc, 0x16,
	0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02, 0x6b, 0x85, 0x70, 0x8c, 0xbc,
	0x60, 0x16, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x72, 0xb1, 0x05, 0xa5, 0x26,
	0xe7, 0x17, 0xa5, 0x08, 0x89, 0x40, 0x9c, 0xa2, 0x87, 0xe2, 0x00, 0x29, 0x51, 0x98, 0x28, 0x8a,
	0x8d, 0x4a, 0x0c, 0x1a, 0x8c, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0xbb, 0x36, 0xaf, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	// 客户端流式RPC，需要在请求数据前加stream
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordClient{stream}
	return x, nil
}

type StreamService_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type streamServiceRecordClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	// 客户端流式RPC，需要在请求数据前加stream
	Record(StreamService_RecordServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) Record(srv StreamService_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Record(&streamServiceRecordServer{stream})
}

type StreamService_RecordServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Record",
			Handler:       _StreamService_Record_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_stream.proto",
}
