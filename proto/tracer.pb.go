// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracer.proto

/*
Package tracer is a generated protocol buffer package.

It is generated from these files:
	tracer.proto

It has these top-level messages:
	ContextProperties
	Milestone
	ContextCall
	RecordRequest
	RecordResponse
	TraceRequest
	TraceResponse
*/
package tracer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Milestone_MilestoneType int32

const (
	Milestone_UNKNOWN        Milestone_MilestoneType = 0
	Milestone_START          Milestone_MilestoneType = 1
	Milestone_START_FUNCTION Milestone_MilestoneType = 2
	Milestone_END_FUNCTION   Milestone_MilestoneType = 3
	Milestone_END            Milestone_MilestoneType = 4
	Milestone_MARKER         Milestone_MilestoneType = 5
)

var Milestone_MilestoneType_name = map[int32]string{
	0: "UNKNOWN",
	1: "START",
	2: "START_FUNCTION",
	3: "END_FUNCTION",
	4: "END",
	5: "MARKER",
}
var Milestone_MilestoneType_value = map[string]int32{
	"UNKNOWN":        0,
	"START":          1,
	"START_FUNCTION": 2,
	"END_FUNCTION":   3,
	"END":            4,
	"MARKER":         5,
}

func (x Milestone_MilestoneType) String() string {
	return proto.EnumName(Milestone_MilestoneType_name, int32(x))
}
func (Milestone_MilestoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ContextProperties struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Created int64  `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Died    int64  `protobuf:"varint,3,opt,name=died" json:"died,omitempty"`
	Origin  string `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Label   string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
}

func (m *ContextProperties) Reset()                    { *m = ContextProperties{} }
func (m *ContextProperties) String() string            { return proto.CompactTextString(m) }
func (*ContextProperties) ProtoMessage()               {}
func (*ContextProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContextProperties) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContextProperties) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ContextProperties) GetDied() int64 {
	if m != nil {
		return m.Died
	}
	return 0
}

func (m *ContextProperties) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ContextProperties) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type Milestone struct {
	Label     string                  `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Timestamp int64                   `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Origin    string                  `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Type      Milestone_MilestoneType `protobuf:"varint,5,opt,name=type,enum=tracer.Milestone_MilestoneType" json:"type,omitempty"`
}

func (m *Milestone) Reset()                    { *m = Milestone{} }
func (m *Milestone) String() string            { return proto.CompactTextString(m) }
func (*Milestone) ProtoMessage()               {}
func (*Milestone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Milestone) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Milestone) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Milestone) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Milestone) GetType() Milestone_MilestoneType {
	if m != nil {
		return m.Type
	}
	return Milestone_UNKNOWN
}

type ContextCall struct {
	Properties *ContextProperties `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
	Milestones []*Milestone       `protobuf:"bytes,2,rep,name=milestones" json:"milestones,omitempty"`
}

func (m *ContextCall) Reset()                    { *m = ContextCall{} }
func (m *ContextCall) String() string            { return proto.CompactTextString(m) }
func (*ContextCall) ProtoMessage()               {}
func (*ContextCall) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContextCall) GetProperties() *ContextProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ContextCall) GetMilestones() []*Milestone {
	if m != nil {
		return m.Milestones
	}
	return nil
}

type RecordRequest struct {
	Properties *ContextProperties `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
	Milestone  *Milestone         `protobuf:"bytes,2,opt,name=milestone" json:"milestone,omitempty"`
}

func (m *RecordRequest) Reset()                    { *m = RecordRequest{} }
func (m *RecordRequest) String() string            { return proto.CompactTextString(m) }
func (*RecordRequest) ProtoMessage()               {}
func (*RecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RecordRequest) GetProperties() *ContextProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *RecordRequest) GetMilestone() *Milestone {
	if m != nil {
		return m.Milestone
	}
	return nil
}

type RecordResponse struct {
}

func (m *RecordResponse) Reset()                    { *m = RecordResponse{} }
func (m *RecordResponse) String() string            { return proto.CompactTextString(m) }
func (*RecordResponse) ProtoMessage()               {}
func (*RecordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type TraceRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator" json:"creator,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TraceRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TraceRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type TraceResponse struct {
	Calls []*ContextCall `protobuf:"bytes,1,rep,name=calls" json:"calls,omitempty"`
}

func (m *TraceResponse) Reset()                    { *m = TraceResponse{} }
func (m *TraceResponse) String() string            { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()               {}
func (*TraceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TraceResponse) GetCalls() []*ContextCall {
	if m != nil {
		return m.Calls
	}
	return nil
}

func init() {
	proto.RegisterType((*ContextProperties)(nil), "tracer.ContextProperties")
	proto.RegisterType((*Milestone)(nil), "tracer.Milestone")
	proto.RegisterType((*ContextCall)(nil), "tracer.ContextCall")
	proto.RegisterType((*RecordRequest)(nil), "tracer.RecordRequest")
	proto.RegisterType((*RecordResponse)(nil), "tracer.RecordResponse")
	proto.RegisterType((*TraceRequest)(nil), "tracer.TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "tracer.TraceResponse")
	proto.RegisterEnum("tracer.Milestone_MilestoneType", Milestone_MilestoneType_name, Milestone_MilestoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TracerService service

type TracerServiceClient interface {
	Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error)
}

type tracerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTracerServiceClient(cc *grpc.ClientConn) TracerServiceClient {
	return &tracerServiceClient{cc}
}

func (c *tracerServiceClient) Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/tracer.TracerService/Record", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracerServiceClient) Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error) {
	out := new(TraceResponse)
	err := grpc.Invoke(ctx, "/tracer.TracerService/Trace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TracerService service

type TracerServiceServer interface {
	Record(context.Context, *RecordRequest) (*RecordResponse, error)
	Trace(context.Context, *TraceRequest) (*TraceResponse, error)
}

func RegisterTracerServiceServer(s *grpc.Server, srv TracerServiceServer) {
	s.RegisterService(&_TracerService_serviceDesc, srv)
}

func _TracerService_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServiceServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracer.TracerService/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServiceServer).Record(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracerService_Trace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServiceServer).Trace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracer.TracerService/Trace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServiceServer).Trace(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TracerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tracer.TracerService",
	HandlerType: (*TracerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _TracerService_Record_Handler,
		},
		{
			MethodName: "Trace",
			Handler:    _TracerService_Trace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracer.proto",
}

func init() { proto.RegisterFile("tracer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x24, 0x4b, 0x46, 0xe3, 0x0f, 0x94, 0x69, 0x12, 0xd4, 0x50, 0xa8, 0xd1, 0xc9, 0xbd,
	0xa4, 0xd4, 0x81, 0x42, 0x7a, 0x28, 0x04, 0xc7, 0x85, 0x12, 0xa2, 0x94, 0x8d, 0x42, 0x8f, 0x45,
	0x96, 0x86, 0xb2, 0x20, 0x7b, 0xd5, 0xd5, 0xb6, 0x34, 0xe4, 0x50, 0xe8, 0x2f, 0xee, 0x4f, 0x28,
	0x5e, 0x69, 0x2d, 0xc5, 0x49, 0x4f, 0xbd, 0xed, 0x3c, 0xde, 0xcc, 0x9b, 0xf7, 0x34, 0x82, 0xa1,
	0x92, 0x69, 0x46, 0xf2, 0xa4, 0x94, 0x42, 0x09, 0xf4, 0xea, 0x2a, 0xfa, 0x05, 0xfb, 0x73, 0xb1,
	0x56, 0xf4, 0x53, 0x7d, 0x92, 0xa2, 0x24, 0xa9, 0x38, 0x55, 0x38, 0x06, 0x9b, 0xe7, 0xa1, 0x35,
	0xb1, 0xa6, 0x3e, 0xb3, 0x79, 0x8e, 0x21, 0xf4, 0x33, 0x49, 0xa9, 0xa2, 0x3c, 0xb4, 0x27, 0xd6,
	0xd4, 0x61, 0xa6, 0x44, 0x84, 0x5e, 0xce, 0x29, 0x0f, 0x1d, 0x0d, 0xeb, 0x37, 0x1e, 0x81, 0x27,
	0x24, 0xff, 0xca, 0xd7, 0x61, 0x4f, 0x4f, 0x68, 0x2a, 0x3c, 0x00, 0xb7, 0x48, 0x97, 0x54, 0x84,
	0xae, 0x86, 0xeb, 0x22, 0xfa, 0x63, 0x81, 0x7f, 0xc5, 0x0b, 0xaa, 0x94, 0x58, 0x53, 0xcb, 0xb1,
	0x3b, 0x1c, 0x7c, 0x01, 0xbe, 0xe2, 0x2b, 0xaa, 0x54, 0xba, 0x2a, 0x1b, 0xa9, 0x16, 0xf8, 0xa7,
	0xde, 0x29, 0xf4, 0xd4, 0x5d, 0x49, 0x5a, 0x6e, 0x3c, 0x7b, 0x79, 0xd2, 0xf8, 0xdf, 0x8a, 0xb5,
	0xaf, 0xe4, 0xae, 0x24, 0xa6, 0xc9, 0xd1, 0x12, 0x46, 0x0f, 0x60, 0x1c, 0x40, 0xff, 0x36, 0xbe,
	0x8c, 0xaf, 0x3f, 0xc7, 0xc1, 0x1e, 0xfa, 0xe0, 0xde, 0x24, 0xe7, 0x2c, 0x09, 0x2c, 0x44, 0x18,
	0xeb, 0xe7, 0x97, 0x0f, 0xb7, 0xf1, 0x3c, 0xf9, 0x78, 0x1d, 0x07, 0x36, 0x06, 0x30, 0x5c, 0xc4,
	0x17, 0x2d, 0xe2, 0x60, 0x1f, 0x9c, 0x45, 0x7c, 0x11, 0xf4, 0x10, 0xc0, 0xbb, 0x3a, 0x67, 0x97,
	0x0b, 0x16, 0xb8, 0xd1, 0x3d, 0x0c, 0x9a, 0xcc, 0xe7, 0x69, 0x51, 0xe0, 0x19, 0x40, 0xb9, 0xcd,
	0x5e, 0xa7, 0x3e, 0x98, 0x3d, 0x37, 0xdb, 0x3e, 0xfa, 0x38, 0xac, 0x43, 0xc6, 0x37, 0x00, 0x2b,
	0xb3, 0x6d, 0x15, 0xda, 0x13, 0x67, 0x3a, 0x98, 0xed, 0x3f, 0x32, 0xca, 0x3a, 0xa4, 0xe8, 0x1e,
	0x46, 0x8c, 0x32, 0x21, 0x73, 0x46, 0xdf, 0xbe, 0x53, 0xa5, 0xfe, 0x47, 0xfe, 0x35, 0xf8, 0xdb,
	0xc9, 0xfa, 0x8b, 0x3d, 0xa9, 0xde, 0x72, 0xa2, 0x00, 0xc6, 0x46, 0xbc, 0x2a, 0xc5, 0xba, 0xa2,
	0xe8, 0x3d, 0x0c, 0x93, 0x4d, 0x83, 0xd9, 0xc6, 0x9c, 0x9a, 0x90, 0xcd, 0xfd, 0x99, 0xf2, 0xe9,
	0xd3, 0x88, 0xde, 0xc1, 0xa8, 0xe9, 0xaf, 0x07, 0xe2, 0x2b, 0x70, 0xb3, 0xb4, 0x28, 0x36, 0x4e,
	0x36, 0x69, 0x3c, 0xdb, 0x71, 0xb2, 0x49, 0x9c, 0xd5, 0x8c, 0xd9, 0x6f, 0xab, 0x69, 0x96, 0x37,
	0x24, 0x7f, 0xf0, 0x8c, 0xf0, 0x0c, 0xbc, 0x7a, 0x3f, 0x3c, 0x34, 0x7d, 0x0f, 0xc2, 0x3a, 0x3e,
	0xda, 0x85, 0x1b, 0x1b, 0x7b, 0xf8, 0x16, 0x5c, 0x3d, 0x0b, 0x0f, 0x0c, 0xa5, 0xeb, 0xeb, 0xf8,
	0x70, 0x07, 0x35, 0x7d, 0x4b, 0x4f, 0xff, 0x8f, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x41,
	0xd8, 0x40, 0x82, 0x9f, 0x03, 0x00, 0x00,
}
