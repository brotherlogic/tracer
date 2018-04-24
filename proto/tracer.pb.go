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
	Creator string `protobuf:"bytes,4,opt,name=creator" json:"creator,omitempty"`
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

func (m *ContextProperties) GetCreator() string {
	if m != nil {
		return m.Creator
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
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x24, 0x4b, 0x46, 0xe3, 0x0f, 0x94, 0x69, 0x12, 0xd4, 0x50, 0xa8, 0xd1, 0xc9, 0xbd,
	0xa4, 0xd4, 0x81, 0x42, 0x7a, 0x28, 0x04, 0xc7, 0x85, 0x12, 0xa2, 0x94, 0x8d, 0x42, 0x8f, 0x45,
	0x96, 0x86, 0xb2, 0x20, 0x7b, 0xd5, 0xd5, 0xb6, 0x34, 0xe4, 0xd4, 0xfe, 0xe2, 0xfe, 0x84, 0xe2,
	0x95, 0xd6, 0x52, 0x9c, 0xf6, 0x94, 0xdb, 0xcc, 0xe3, 0xcd, 0xbc, 0x99, 0xb7, 0xb3, 0x30, 0x54,
	0x32, 0xcd, 0x48, 0x9e, 0x94, 0x52, 0x28, 0x81, 0x5e, 0x9d, 0x45, 0xbf, 0x2c, 0xd8, 0x9f, 0x8b,
	0xb5, 0xa2, 0x9f, 0xea, 0x93, 0x14, 0x25, 0x49, 0xc5, 0xa9, 0xc2, 0x31, 0xd8, 0x3c, 0x0f, 0xad,
	0x89, 0x35, 0xf5, 0x99, 0xcd, 0x73, 0x0c, 0xa1, 0x9f, 0x49, 0x4a, 0x15, 0xe5, 0xa1, 0x3d, 0xb1,
	0xa6, 0x0e, 0x33, 0x29, 0x22, 0xf4, 0x72, 0x4e, 0x79, 0xe8, 0x68, 0x58, 0xc7, 0x5b, 0xb6, 0x90,
	0x61, 0x4f, 0xb7, 0x30, 0x29, 0x1e, 0x80, 0x5b, 0xa4, 0x4b, 0x2a, 0x42, 0x57, 0xe3, 0x75, 0x12,
	0xfd, 0xb1, 0xc0, 0xbf, 0xe2, 0x05, 0x55, 0x4a, 0xac, 0xa9, 0xe5, 0xd8, 0x1d, 0x0e, 0xbe, 0x00,
	0x5f, 0xf1, 0x15, 0x55, 0x2a, 0x5d, 0x95, 0x8d, 0x58, 0x0b, 0xe0, 0x11, 0x78, 0x42, 0xf2, 0xaf,
	0x7c, 0xdd, 0x08, 0x36, 0x19, 0x9e, 0x42, 0x4f, 0xdd, 0x95, 0xa4, 0xe5, 0xc6, 0xb3, 0x97, 0x27,
	0x8d, 0x05, 0x5b, 0xb1, 0x36, 0x4a, 0xee, 0x4a, 0x62, 0x9a, 0x1c, 0x2d, 0x61, 0xf4, 0x00, 0xc6,
	0x01, 0xf4, 0x6f, 0xe3, 0xcb, 0xf8, 0xfa, 0x73, 0x1c, 0xec, 0xa1, 0x0f, 0xee, 0x4d, 0x72, 0xce,
	0x92, 0xc0, 0x42, 0x84, 0xb1, 0x0e, 0xbf, 0x7c, 0xb8, 0x8d, 0xe7, 0xc9, 0xc7, 0xeb, 0x38, 0xb0,
	0x31, 0x80, 0xe1, 0x22, 0xbe, 0x68, 0x11, 0x07, 0xfb, 0xe0, 0x2c, 0xe2, 0x8b, 0xa0, 0x87, 0x00,
	0xde, 0xd5, 0x39, 0xbb, 0x5c, 0xb0, 0xc0, 0x8d, 0xee, 0x61, 0xd0, 0xb8, 0x3e, 0x4f, 0x8b, 0x02,
	0xcf, 0x00, 0xca, 0xad, 0xfb, 0xda, 0xf7, 0xc1, 0xec, 0xb9, 0x99, 0xf6, 0xd1, 0xf3, 0xb0, 0x0e,
	0x19, 0xdf, 0x00, 0xac, 0xcc, 0xb4, 0x55, 0x68, 0x4f, 0x9c, 0xe9, 0x60, 0xb6, 0xff, 0x68, 0x51,
	0xd6, 0x21, 0x45, 0xf7, 0x30, 0x62, 0x94, 0x09, 0x99, 0x33, 0xfa, 0xf6, 0x9d, 0x2a, 0xf5, 0x14,
	0xf9, 0xd7, 0xe0, 0x6f, 0x3b, 0xeb, 0x17, 0xfb, 0xa7, 0x7a, 0xcb, 0x89, 0x02, 0x18, 0x1b, 0xf1,
	0xaa, 0x14, 0xeb, 0x8a, 0xa2, 0xf7, 0x30, 0x4c, 0x36, 0x05, 0x66, 0x9a, 0xce, 0xf9, 0x58, 0xff,
	0x39, 0x9f, 0xee, 0x69, 0x44, 0xef, 0x60, 0xd4, 0xd4, 0xd7, 0x0d, 0xf1, 0x15, 0xb8, 0x59, 0x5a,
	0x14, 0x9b, 0x4d, 0x36, 0x6e, 0x3c, 0xdb, 0xd9, 0x64, 0xe3, 0x38, 0xab, 0x19, 0xb3, 0xdf, 0x56,
	0x53, 0x2c, 0x6f, 0x48, 0xfe, 0xe0, 0x19, 0xe1, 0x19, 0x78, 0xf5, 0x7c, 0x78, 0x68, 0xea, 0x1e,
	0x98, 0x75, 0x7c, 0xb4, 0x0b, 0x37, 0x6b, 0xec, 0xe1, 0x5b, 0x70, 0x75, 0x2f, 0x3c, 0x30, 0x94,
	0xee, 0x5e, 0xc7, 0x87, 0x3b, 0xa8, 0xa9, 0x5b, 0x7a, 0xfa, 0x4b, 0x9e, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0xdc, 0xb9, 0xef, 0xa2, 0x03, 0x00, 0x00,
}
