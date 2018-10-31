// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/collections.proto

package collections // import "github.com/airbloc/airbloc-go/collections"

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

type Policy struct {
	DataProducer         float32  `protobuf:"fixed32,1,opt,name=DataProducer,proto3" json:"DataProducer,omitempty"`
	DataProcessor        float32  `protobuf:"fixed32,2,opt,name=DataProcessor,proto3" json:"DataProcessor,omitempty"`
	DataRelayer          float32  `protobuf:"fixed32,3,opt,name=DataRelayer,proto3" json:"DataRelayer,omitempty"`
	DataSource           float32  `protobuf:"fixed32,4,opt,name=DataSource,proto3" json:"DataSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetDataProducer() float32 {
	if m != nil {
		return m.DataProducer
	}
	return 0
}

func (m *Policy) GetDataProcessor() float32 {
	if m != nil {
		return m.DataProcessor
	}
	return 0
}

func (m *Policy) GetDataRelayer() float32 {
	if m != nil {
		return m.DataRelayer
	}
	return 0
}

func (m *Policy) GetDataSource() float32 {
	if m != nil {
		return m.DataSource
	}
	return 0
}

type CreateCollectionRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	SchemaId             string   `protobuf:"bytes,2,opt,name=schemaId,proto3" json:"schemaId,omitempty"`
	Policy               *Policy  `protobuf:"bytes,11,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionRequest) Reset()         { *m = CreateCollectionRequest{} }
func (m *CreateCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionRequest) ProtoMessage()    {}
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{1}
}
func (m *CreateCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionRequest.Unmarshal(m, b)
}
func (m *CreateCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionRequest.Merge(dst, src)
}
func (m *CreateCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionRequest.Size(m)
}
func (m *CreateCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionRequest proto.InternalMessageInfo

func (m *CreateCollectionRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateCollectionRequest) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *CreateCollectionRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type CreateCollectionResponse struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionResponse) Reset()         { *m = CreateCollectionResponse{} }
func (m *CreateCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionResponse) ProtoMessage()    {}
func (*CreateCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{2}
}
func (m *CreateCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionResponse.Unmarshal(m, b)
}
func (m *CreateCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionResponse.Merge(dst, src)
}
func (m *CreateCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionResponse.Size(m)
}
func (m *CreateCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionResponse proto.InternalMessageInfo

func (m *CreateCollectionResponse) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

type ListCollectionRequest struct {
	PageIndex            int32    `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	ItemsPerPage         int32    `protobuf:"varint,2,opt,name=itemsPerPage,proto3" json:"itemsPerPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCollectionRequest) Reset()         { *m = ListCollectionRequest{} }
func (m *ListCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*ListCollectionRequest) ProtoMessage()    {}
func (*ListCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{3}
}
func (m *ListCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCollectionRequest.Unmarshal(m, b)
}
func (m *ListCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCollectionRequest.Marshal(b, m, deterministic)
}
func (dst *ListCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCollectionRequest.Merge(dst, src)
}
func (m *ListCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_ListCollectionRequest.Size(m)
}
func (m *ListCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCollectionRequest proto.InternalMessageInfo

func (m *ListCollectionRequest) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *ListCollectionRequest) GetItemsPerPage() int32 {
	if m != nil {
		return m.ItemsPerPage
	}
	return 0
}

type CollectionResponse struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	SchemaId             string   `protobuf:"bytes,2,opt,name=schemaId,proto3" json:"schemaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionResponse) Reset()         { *m = CollectionResponse{} }
func (m *CollectionResponse) String() string { return proto.CompactTextString(m) }
func (*CollectionResponse) ProtoMessage()    {}
func (*CollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{4}
}
func (m *CollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionResponse.Unmarshal(m, b)
}
func (m *CollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionResponse.Marshal(b, m, deterministic)
}
func (dst *CollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionResponse.Merge(dst, src)
}
func (m *CollectionResponse) XXX_Size() int {
	return xxx_messageInfo_CollectionResponse.Size(m)
}
func (m *CollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionResponse proto.InternalMessageInfo

func (m *CollectionResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CollectionResponse) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

type ListCollectionResponse struct {
	Total                int32                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Collections          []*CollectionResponse `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCollectionResponse) Reset()         { *m = ListCollectionResponse{} }
func (m *ListCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*ListCollectionResponse) ProtoMessage()    {}
func (*ListCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_collections_66ee75dfb019f958, []int{5}
}
func (m *ListCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCollectionResponse.Unmarshal(m, b)
}
func (m *ListCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCollectionResponse.Marshal(b, m, deterministic)
}
func (dst *ListCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCollectionResponse.Merge(dst, src)
}
func (m *ListCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_ListCollectionResponse.Size(m)
}
func (m *ListCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCollectionResponse proto.InternalMessageInfo

func (m *ListCollectionResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListCollectionResponse) GetCollections() []*CollectionResponse {
	if m != nil {
		return m.Collections
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "airbloc.collections.Policy")
	proto.RegisterType((*CreateCollectionRequest)(nil), "airbloc.collections.CreateCollectionRequest")
	proto.RegisterType((*CreateCollectionResponse)(nil), "airbloc.collections.CreateCollectionResponse")
	proto.RegisterType((*ListCollectionRequest)(nil), "airbloc.collections.ListCollectionRequest")
	proto.RegisterType((*CollectionResponse)(nil), "airbloc.collections.CollectionResponse")
	proto.RegisterType((*ListCollectionResponse)(nil), "airbloc.collections.ListCollectionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectionClient is the client API for Collection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectionClient interface {
	Create(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	List(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error)
}

type collectionClient struct {
	cc *grpc.ClientConn
}

func NewCollectionClient(cc *grpc.ClientConn) CollectionClient {
	return &collectionClient{cc}
}

func (c *collectionClient) Create(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, "/airbloc.collections.Collection/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) List(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error) {
	out := new(ListCollectionResponse)
	err := c.cc.Invoke(ctx, "/airbloc.collections.Collection/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServer is the server API for Collection service.
type CollectionServer interface {
	Create(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	List(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error)
}

func RegisterCollectionServer(s *grpc.Server, srv CollectionServer) {
	s.RegisterService(&_Collection_serviceDesc, srv)
}

func _Collection_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.collections.Collection/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).Create(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.collections.Collection/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).List(ctx, req.(*ListCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Collection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.collections.Collection",
	HandlerType: (*CollectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Collection_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Collection_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/collections.proto",
}

func init() {
	proto.RegisterFile("proto/collections.proto", fileDescriptor_collections_66ee75dfb019f958)
}

var fileDescriptor_collections_66ee75dfb019f958 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0xfd, 0x09, 0xed, 0x8d, 0x82, 0x8c, 0x3f, 0x0d, 0x55, 0x24, 0x04, 0xc1, 0x6a, 0x6d,
	0x0a, 0xed, 0xde, 0x85, 0x15, 0x21, 0xe2, 0x22, 0x8c, 0x2b, 0xdd, 0xc8, 0x74, 0x72, 0x49, 0x03,
	0x69, 0x27, 0xce, 0x4c, 0xc0, 0xbe, 0x80, 0x4f, 0xe0, 0xa3, 0xf9, 0x40, 0x92, 0x49, 0x6c, 0x92,
	0xaf, 0x59, 0xf4, 0x5b, 0x85, 0x7b, 0x72, 0xce, 0xcd, 0x39, 0x27, 0x33, 0x30, 0xcb, 0xa5, 0xd0,
	0x62, 0xcd, 0x45, 0x96, 0x21, 0xd7, 0xa9, 0x38, 0xa9, 0xc0, 0x20, 0xe4, 0x31, 0x4b, 0xe5, 0x3e,
	0x13, 0x3c, 0x68, 0xbd, 0xf2, 0xff, 0x58, 0x60, 0x47, 0x22, 0x4b, 0xf9, 0x99, 0xf8, 0xf0, 0xe0,
	0x23, 0xd3, 0x2c, 0x92, 0x22, 0x2e, 0x38, 0x4a, 0xd7, 0xf2, 0xac, 0xc5, 0x80, 0x76, 0x30, 0xf2,
	0x0a, 0x1e, 0xd6, 0x33, 0x47, 0xa5, 0x84, 0x74, 0x07, 0x86, 0xd4, 0x05, 0x89, 0x07, 0x4e, 0x09,
	0x50, 0xcc, 0xd8, 0x19, 0xa5, 0x3b, 0x34, 0x9c, 0x36, 0x44, 0x5e, 0x02, 0x94, 0xe3, 0x57, 0x51,
	0x48, 0x8e, 0xee, 0xc8, 0x10, 0x5a, 0x88, 0xff, 0xdb, 0x82, 0xd9, 0x4e, 0x22, 0xd3, 0xb8, 0xbb,
	0x98, 0xa5, 0xf8, 0xb3, 0x40, 0xa5, 0xc9, 0x13, 0x18, 0xb3, 0x3c, 0x0f, 0x63, 0x63, 0x70, 0x4a,
	0xab, 0x81, 0xcc, 0x61, 0xa2, 0xf8, 0x01, 0x8f, 0x2c, 0x8c, 0x8d, 0xa9, 0x29, 0xbd, 0xcc, 0x64,
	0x0b, 0x76, 0x6e, 0x32, 0xba, 0x8e, 0x67, 0x2d, 0x9c, 0xcd, 0xf3, 0xa0, 0xa7, 0x8a, 0xa0, 0xaa,
	0x81, 0xd6, 0xd4, 0xcf, 0xa3, 0xc9, 0xf0, 0x91, 0xe3, 0xbf, 0x07, 0xf7, 0xda, 0x87, 0xca, 0xc5,
	0x49, 0x61, 0x59, 0x58, 0xa3, 0xbf, 0xf8, 0xe9, 0x60, 0xfe, 0x37, 0x78, 0xfa, 0x25, 0x55, 0xfa,
	0x3a, 0xc5, 0x0b, 0x98, 0xe6, 0x2c, 0xc1, 0xf0, 0x14, 0xe3, 0x2f, 0xa3, 0x1c, 0xd3, 0x06, 0x28,
	0x57, 0xa7, 0x1a, 0x8f, 0x2a, 0x42, 0x19, 0xb1, 0x04, 0x4d, 0xa2, 0x31, 0xed, 0x60, 0xfe, 0x27,
	0x20, 0x3d, 0xa6, 0xee, 0xdd, 0x8e, 0x7f, 0x86, 0x67, 0x77, 0x2d, 0x36, 0xbb, 0xb4, 0xd0, 0x2c,
	0xab, 0xfd, 0x55, 0x03, 0x09, 0xc1, 0x69, 0xd5, 0xe6, 0x0e, 0xbc, 0xe1, 0xc2, 0xd9, 0xbc, 0xee,
	0xad, 0xf4, 0x7a, 0x27, 0x6d, 0x6b, 0x37, 0x7f, 0x2d, 0x80, 0x86, 0x43, 0x10, 0xec, 0xaa, 0x6c,
	0xf2, 0xae, 0x7f, 0x5d, 0xff, 0x89, 0x98, 0xaf, 0x6e, 0x64, 0xd7, 0xb1, 0x7e, 0xc0, 0xa8, 0x0c,
	0x4c, 0xde, 0xf6, 0xca, 0x7a, 0x7f, 0xd7, 0x7c, 0x79, 0x13, 0xb7, 0xfa, 0xc0, 0x87, 0xe5, 0xf7,
	0x37, 0x49, 0xaa, 0x0f, 0xc5, 0x3e, 0xe0, 0xe2, 0xb8, 0xae, 0x85, 0xff, 0x9f, 0xab, 0xa4, 0x73,
	0x39, 0xf7, 0xb6, 0xb9, 0x9d, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x69, 0xf4, 0x19,
	0xb8, 0x03, 0x00, 0x00,
}
