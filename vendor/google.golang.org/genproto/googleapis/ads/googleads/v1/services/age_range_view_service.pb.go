// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/age_range_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Request message for [AgeRangeViewService.GetAgeRangeView][google.ads.googleads.v1.services.AgeRangeViewService.GetAgeRangeView].
type GetAgeRangeViewRequest struct {
	// The resource name of the age range view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgeRangeViewRequest) Reset()         { *m = GetAgeRangeViewRequest{} }
func (m *GetAgeRangeViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgeRangeViewRequest) ProtoMessage()    {}
func (*GetAgeRangeViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_age_range_view_service_7c05e971116a806c, []int{0}
}
func (m *GetAgeRangeViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgeRangeViewRequest.Unmarshal(m, b)
}
func (m *GetAgeRangeViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgeRangeViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetAgeRangeViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgeRangeViewRequest.Merge(dst, src)
}
func (m *GetAgeRangeViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgeRangeViewRequest.Size(m)
}
func (m *GetAgeRangeViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgeRangeViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgeRangeViewRequest proto.InternalMessageInfo

func (m *GetAgeRangeViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAgeRangeViewRequest)(nil), "google.ads.googleads.v1.services.GetAgeRangeViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgeRangeViewServiceClient is the client API for AgeRangeViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgeRangeViewServiceClient interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error)
}

type ageRangeViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgeRangeViewServiceClient(cc *grpc.ClientConn) AgeRangeViewServiceClient {
	return &ageRangeViewServiceClient{cc}
}

func (c *ageRangeViewServiceClient) GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error) {
	out := new(resources.AgeRangeView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AgeRangeViewService/GetAgeRangeView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgeRangeViewServiceServer is the server API for AgeRangeViewService service.
type AgeRangeViewServiceServer interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(context.Context, *GetAgeRangeViewRequest) (*resources.AgeRangeView, error)
}

func RegisterAgeRangeViewServiceServer(s *grpc.Server, srv AgeRangeViewServiceServer) {
	s.RegisterService(&_AgeRangeViewService_serviceDesc, srv)
}

func _AgeRangeViewService_GetAgeRangeView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgeRangeViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AgeRangeViewService/GetAgeRangeView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, req.(*GetAgeRangeViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgeRangeViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AgeRangeViewService",
	HandlerType: (*AgeRangeViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgeRangeView",
			Handler:    _AgeRangeViewService_GetAgeRangeView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/age_range_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/age_range_view_service.proto", fileDescriptor_age_range_view_service_7c05e971116a806c)
}

var fileDescriptor_age_range_view_service_7c05e971116a806c = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x49, 0x04, 0xc1, 0xa0, 0x08, 0x11, 0xa4, 0x14, 0x87, 0x52, 0x3b, 0x48, 0x87, 0x3b,
	0xa2, 0xa0, 0x72, 0xd2, 0x21, 0x5d, 0xea, 0x24, 0xa5, 0x42, 0x06, 0x09, 0x84, 0xb3, 0x79, 0x39,
	0x02, 0x4d, 0xae, 0xde, 0x9b, 0xa6, 0x83, 0xb8, 0xf8, 0x15, 0xfc, 0x06, 0x8e, 0xee, 0x7e, 0x09,
	0xc1, 0xc9, 0xaf, 0xe0, 0xe4, 0x97, 0x50, 0xd2, 0xeb, 0x85, 0xaa, 0x2d, 0xdd, 0x1e, 0xde, 0x3c,
	0xbf, 0xf7, 0xcf, 0x93, 0x73, 0x3a, 0x42, 0x4a, 0x31, 0x02, 0xca, 0x63, 0xa4, 0x5a, 0x96, 0xaa,
	0xf0, 0x28, 0x82, 0x2a, 0x92, 0x21, 0x20, 0xe5, 0x02, 0x22, 0xc5, 0x33, 0x01, 0x51, 0x91, 0xc0,
	0x34, 0x9a, 0xd7, 0xc9, 0x58, 0xc9, 0x5c, 0xba, 0x0d, 0xcd, 0x10, 0x1e, 0x23, 0xa9, 0x70, 0x52,
	0x78, 0xc4, 0xe0, 0xf5, 0xd3, 0x55, 0x03, 0x14, 0xa0, 0x9c, 0xa8, 0xff, 0x13, 0x74, 0xe7, 0xfa,
	0x81, 0xe1, 0xc6, 0x09, 0xe5, 0x59, 0x26, 0x73, 0x9e, 0x27, 0x32, 0x43, 0xfd, 0xb5, 0xd9, 0x71,
	0xf6, 0x7b, 0x90, 0xfb, 0x02, 0x06, 0x25, 0x17, 0x24, 0x30, 0x1d, 0xc0, 0xdd, 0x04, 0x30, 0x77,
	0x0f, 0x9d, 0x1d, 0xd3, 0x39, 0xca, 0x78, 0x0a, 0x35, 0xab, 0x61, 0x1d, 0x6d, 0x0d, 0xb6, 0x4d,
	0xf1, 0x8a, 0xa7, 0x70, 0xfc, 0x6e, 0x39, 0x7b, 0x8b, 0xf0, 0xb5, 0xde, 0xd6, 0x7d, 0xb5, 0x9c,
	0xdd, 0x3f, 0x7d, 0xdd, 0x73, 0xb2, 0xee, 0x46, 0xb2, 0x7c, 0x95, 0x3a, 0x5d, 0x49, 0x56, 0xb7,
	0x93, 0x45, 0xae, 0x79, 0xf6, 0xf8, 0xf1, 0xf9, 0x64, 0x7b, 0x2e, 0x2d, 0xf3, 0xb9, 0xff, 0x75,
	0x46, 0x67, 0x38, 0xc1, 0x5c, 0xa6, 0xa0, 0x90, 0xb6, 0xcb, 0xc0, 0x2a, 0x08, 0x69, 0xfb, 0xa1,
	0xfb, 0x6d, 0x39, 0xad, 0xa1, 0x4c, 0xd7, 0x6e, 0xda, 0xad, 0x2d, 0xb9, 0xba, 0x5f, 0x26, 0xda,
	0xb7, 0x6e, 0x2e, 0xe7, 0xb4, 0x90, 0x23, 0x9e, 0x09, 0x22, 0x95, 0xa0, 0x02, 0xb2, 0x59, 0xde,
	0xe6, 0xcf, 0x8d, 0x13, 0x5c, 0xfd, 0x52, 0x2e, 0x8c, 0x78, 0xb6, 0x37, 0x7a, 0xbe, 0xff, 0x62,
	0x37, 0x7a, 0xba, 0xa1, 0x1f, 0x23, 0xd1, 0xb2, 0x54, 0x81, 0x47, 0xe6, 0x83, 0xf1, 0xcd, 0x58,
	0x42, 0x3f, 0xc6, 0xb0, 0xb2, 0x84, 0x81, 0x17, 0x1a, 0xcb, 0x97, 0xdd, 0xd2, 0x75, 0xc6, 0xfc,
	0x18, 0x19, 0xab, 0x4c, 0x8c, 0x05, 0x1e, 0x63, 0xc6, 0x76, 0xbb, 0x39, 0xdb, 0xf3, 0xe4, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x12, 0xf3, 0x69, 0xd0, 0x02, 0x00, 0x00,
}