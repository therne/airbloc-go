// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/data.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type BatchRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batchId,proto3" json:"batchId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{0}
}

func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest.Unmarshal(m, b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
}
func (m *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(m, src)
}
func (m *BatchRequest) XXX_Size() int {
	return xxx_messageInfo_BatchRequest.Size(m)
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

type DataId struct {
	DataId               string   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataId) Reset()         { *m = DataId{} }
func (m *DataId) String() string { return proto.CompactTextString(m) }
func (*DataId) ProtoMessage()    {}
func (*DataId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{1}
}

func (m *DataId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataId.Unmarshal(m, b)
}
func (m *DataId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataId.Marshal(b, m, deterministic)
}
func (m *DataId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataId.Merge(m, src)
}
func (m *DataId) XXX_Size() int {
	return xxx_messageInfo_DataId.Size(m)
}
func (m *DataId) XXX_DiscardUnknown() {
	xxx_messageInfo_DataId.DiscardUnknown(m)
}

var xxx_messageInfo_DataId proto.InternalMessageInfo

func (m *DataId) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

type DataResult struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	OwnerUserAid         string   `protobuf:"bytes,2,opt,name=ownerUserAid,proto3" json:"ownerUserAid,omitempty"`
	IngestedAt           int64    `protobuf:"varint,3,opt,name=ingestedAt,proto3" json:"ingestedAt,omitempty"`
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResult) Reset()         { *m = DataResult{} }
func (m *DataResult) String() string { return proto.CompactTextString(m) }
func (*DataResult) ProtoMessage()    {}
func (*DataResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{2}
}

func (m *DataResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResult.Unmarshal(m, b)
}
func (m *DataResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResult.Marshal(b, m, deterministic)
}
func (m *DataResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResult.Merge(m, src)
}
func (m *DataResult) XXX_Size() int {
	return xxx_messageInfo_DataResult.Size(m)
}
func (m *DataResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResult.DiscardUnknown(m)
}

var xxx_messageInfo_DataResult proto.InternalMessageInfo

func (m *DataResult) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *DataResult) GetOwnerUserAid() string {
	if m != nil {
		return m.OwnerUserAid
	}
	return ""
}

func (m *DataResult) GetIngestedAt() int64 {
	if m != nil {
		return m.IngestedAt
	}
	return 0
}

func (m *DataResult) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type GetBatchResult struct {
	Data                 []*DataResult `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBatchResult) Reset()         { *m = GetBatchResult{} }
func (m *GetBatchResult) String() string { return proto.CompactTextString(m) }
func (*GetBatchResult) ProtoMessage()    {}
func (*GetBatchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{3}
}

func (m *GetBatchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBatchResult.Unmarshal(m, b)
}
func (m *GetBatchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBatchResult.Marshal(b, m, deterministic)
}
func (m *GetBatchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBatchResult.Merge(m, src)
}
func (m *GetBatchResult) XXX_Size() int {
	return xxx_messageInfo_GetBatchResult.Size(m)
}
func (m *GetBatchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBatchResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetBatchResult proto.InternalMessageInfo

func (m *GetBatchResult) GetData() []*DataResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type SetDataPermissionRequest struct {
	DataId               string   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	ConsumerId           string   `protobuf:"bytes,2,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	AllowAccess          bool     `protobuf:"varint,3,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDataPermissionRequest) Reset()         { *m = SetDataPermissionRequest{} }
func (m *SetDataPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataPermissionRequest) ProtoMessage()    {}
func (*SetDataPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{4}
}

func (m *SetDataPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataPermissionRequest.Unmarshal(m, b)
}
func (m *SetDataPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataPermissionRequest.Marshal(b, m, deterministic)
}
func (m *SetDataPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataPermissionRequest.Merge(m, src)
}
func (m *SetDataPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_SetDataPermissionRequest.Size(m)
}
func (m *SetDataPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataPermissionRequest proto.InternalMessageInfo

func (m *SetDataPermissionRequest) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *SetDataPermissionRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *SetDataPermissionRequest) GetAllowAccess() bool {
	if m != nil {
		return m.AllowAccess
	}
	return false
}

type SetBatchDataPermissionRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batchId,proto3" json:"batchId,omitempty"`
	ConsumerId           string   `protobuf:"bytes,2,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	AllowAccess          bool     `protobuf:"varint,3,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBatchDataPermissionRequest) Reset()         { *m = SetBatchDataPermissionRequest{} }
func (m *SetBatchDataPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*SetBatchDataPermissionRequest) ProtoMessage()    {}
func (*SetBatchDataPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2002f69fb9f05dba, []int{5}
}

func (m *SetBatchDataPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Unmarshal(m, b)
}
func (m *SetBatchDataPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Marshal(b, m, deterministic)
}
func (m *SetBatchDataPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBatchDataPermissionRequest.Merge(m, src)
}
func (m *SetBatchDataPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Size(m)
}
func (m *SetBatchDataPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBatchDataPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBatchDataPermissionRequest proto.InternalMessageInfo

func (m *SetBatchDataPermissionRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

func (m *SetBatchDataPermissionRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *SetBatchDataPermissionRequest) GetAllowAccess() bool {
	if m != nil {
		return m.AllowAccess
	}
	return false
}

func init() {
	proto.RegisterType((*BatchRequest)(nil), "airbloc.rpc.v1.BatchRequest")
	proto.RegisterType((*DataId)(nil), "airbloc.rpc.v1.DataId")
	proto.RegisterType((*DataResult)(nil), "airbloc.rpc.v1.DataResult")
	proto.RegisterType((*GetBatchResult)(nil), "airbloc.rpc.v1.GetBatchResult")
	proto.RegisterType((*SetDataPermissionRequest)(nil), "airbloc.rpc.v1.SetDataPermissionRequest")
	proto.RegisterType((*SetBatchDataPermissionRequest)(nil), "airbloc.rpc.v1.SetBatchDataPermissionRequest")
}

func init() { proto.RegisterFile("proto/rpc/v1/server/data.proto", fileDescriptor_2002f69fb9f05dba) }

var fileDescriptor_2002f69fb9f05dba = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x49, 0x94, 0x96, 0x49, 0xe9, 0x61, 0x0f, 0x91, 0x65, 0x20, 0xb2, 0x7c, 0xf2, 0xa5,
	0x6b, 0xa5, 0xbd, 0xc0, 0x89, 0xa6, 0x14, 0x95, 0xdc, 0xc0, 0x11, 0x17, 0x38, 0xad, 0xd7, 0x83,
	0x6b, 0x69, 0xe3, 0x35, 0xbb, 0xeb, 0x54, 0x15, 0xff, 0x81, 0x7f, 0xc4, 0x7f, 0x43, 0xeb, 0x0f,
	0xea, 0x44, 0x71, 0x85, 0xe8, 0xc9, 0x9e, 0xb7, 0x6f, 0x67, 0xde, 0xbc, 0x19, 0x1b, 0xe6, 0xa5,
	0x92, 0x46, 0x46, 0xaa, 0xe4, 0xd1, 0x76, 0x11, 0x69, 0x54, 0x5b, 0x54, 0x51, 0xca, 0x0c, 0xa3,
	0xf5, 0x01, 0x39, 0x65, 0xb9, 0x4a, 0x84, 0xe4, 0x54, 0x95, 0x9c, 0x6e, 0x17, 0xde, 0xcb, 0x4c,
	0xca, 0x4c, 0x60, 0x54, 0x9f, 0x26, 0xd5, 0xf7, 0x08, 0x37, 0xa5, 0xb9, 0x6f, 0xc8, 0x41, 0x08,
	0x27, 0x57, 0xcc, 0xf0, 0xdb, 0x18, 0x7f, 0x54, 0xa8, 0x0d, 0x71, 0xe1, 0x28, 0xb1, 0xf1, 0x2a,
	0x75, 0x1d, 0xdf, 0x09, 0x9f, 0xc7, 0x5d, 0x18, 0xf8, 0x30, 0xb9, 0x66, 0x86, 0xad, 0x52, 0x32,
	0x83, 0x49, 0x5a, 0xbf, 0xb5, 0x94, 0x36, 0x0a, 0x7e, 0x39, 0x00, 0x96, 0x12, 0xa3, 0xae, 0x84,
	0x21, 0x01, 0x9c, 0x70, 0x29, 0x04, 0x72, 0x93, 0xcb, 0xe2, 0x2f, 0x79, 0x07, 0xb3, 0x1c, 0x79,
	0x57, 0xa0, 0xfa, 0xa2, 0x51, 0x2d, 0xf3, 0xd4, 0x7d, 0xd6, 0x70, 0xfa, 0x18, 0x99, 0x03, 0xe4,
	0x45, 0x86, 0xda, 0x60, 0xba, 0x34, 0xee, 0xc8, 0x77, 0xc2, 0x51, 0xdc, 0x43, 0xac, 0xe4, 0x92,
	0xdd, 0x0b, 0xc9, 0x52, 0x77, 0xdc, 0x48, 0x6e, 0xc3, 0xe0, 0x12, 0x4e, 0x6f, 0xd0, 0xb4, 0xfd,
	0xd5, 0x9a, 0x28, 0x8c, 0xad, 0x58, 0xd7, 0xf1, 0x47, 0xe1, 0xf4, 0xdc, 0xa3, 0xbb, 0x56, 0xd1,
	0x07, 0xf5, 0x71, 0xcd, 0x0b, 0x0c, 0xb8, 0x6b, 0x34, 0x16, 0xfe, 0x84, 0x6a, 0x93, 0x6b, 0x9d,
	0xcb, 0xa2, 0xb3, 0x6a, 0xc0, 0x06, 0xab, 0x97, 0xcb, 0x42, 0x57, 0x1b, 0x54, 0xab, 0xae, 0xa3,
	0x1e, 0x42, 0x7c, 0x98, 0x32, 0x21, 0xe4, 0xdd, 0x92, 0x73, 0xd4, 0xba, 0x6e, 0xe8, 0x38, 0xee,
	0x43, 0xc1, 0x4f, 0x78, 0xbd, 0x6e, 0x75, 0x1f, 0x2e, 0x3d, 0x38, 0xa5, 0xa7, 0x17, 0x3f, 0xff,
	0x3d, 0x86, 0xb1, 0xad, 0x4a, 0xde, 0xc2, 0xe8, 0x06, 0x0d, 0x99, 0x1d, 0x32, 0x69, 0x95, 0x7a,
	0x8f, 0x98, 0x47, 0x3e, 0xc2, 0x71, 0x67, 0x3c, 0x79, 0xb5, 0xcf, 0xeb, 0xef, 0x9b, 0x37, 0xdf,
	0x3f, 0xdd, 0x1b, 0xd8, 0x67, 0x78, 0xb1, 0x46, 0xf3, 0xe0, 0x00, 0x09, 0xf7, 0x2f, 0x0c, 0xcd,
	0xc7, 0x9b, 0xd1, 0x66, 0xf1, 0x69, 0xb7, 0xf8, 0xf4, 0x83, 0x5d, 0x7c, 0xf2, 0x0d, 0xc8, 0x4e,
	0xca, 0x46, 0xe6, 0xd9, 0x81, 0xbc, 0xc3, 0x13, 0x18, 0x4c, 0xfe, 0x06, 0x26, 0xd7, 0x28, 0xd0,
	0xe0, 0xa0, 0x6f, 0x43, 0x37, 0xdf, 0xc3, 0xb4, 0xb9, 0xf9, 0x2f, 0xb6, 0x0d, 0x25, 0xb9, 0x84,
	0xc9, 0x1a, 0xed, 0xe7, 0x35, 0x58, 0xfe, 0xd1, 0xbc, 0xa1, 0x43, 0xde, 0xc1, 0x51, 0x8c, 0x02,
	0x99, 0xc6, 0xff, 0x93, 0x70, 0x75, 0xf1, 0x75, 0x91, 0xe5, 0xe6, 0xb6, 0x4a, 0x28, 0x97, 0x9b,
	0xa8, 0xcd, 0xd0, 0x3d, 0xcf, 0x32, 0x19, 0x1d, 0xf8, 0x7d, 0x25, 0x93, 0x1a, 0xbc, 0xf8, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0x47, 0x8c, 0xd0, 0xdc, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	// Get returns information of the data from your warehouse or purchased data.
	Get(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*DataResult, error)
	GetBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*GetBatchResult, error)
	// SetPermission allows a consumer to access the given data.
	SetPermission(ctx context.Context, in *SetDataPermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetPermissionBatch(ctx context.Context, in *SetBatchDataPermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete removes specific data from your warehouse or purchased data.
	Delete(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Select receives stream of data ID and creates a batch of data IDs.
	// the batch can be used as a input of batch operations.
	Select(ctx context.Context, opts ...grpc.CallOption) (Data_SelectClient, error)
	// Release unselects and removes the batch.
	Release(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Get(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*DataResult, error) {
	out := new(DataResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*GetBatchResult, error) {
	out := new(GetBatchResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) SetPermission(ctx context.Context, in *SetDataPermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/SetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) SetPermissionBatch(ctx context.Context, in *SetBatchDataPermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/SetPermissionBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Delete(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DeleteBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/DeleteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Select(ctx context.Context, opts ...grpc.CallOption) (Data_SelectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[0], "/airbloc.rpc.v1.Data/Select", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataSelectClient{stream}
	return x, nil
}

type Data_SelectClient interface {
	Send(*DataId) error
	CloseAndRecv() (*BatchRequest, error)
	grpc.ClientStream
}

type dataSelectClient struct {
	grpc.ClientStream
}

func (x *dataSelectClient) Send(m *DataId) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataSelectClient) CloseAndRecv() (*BatchRequest, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BatchRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) Release(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Data/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	// Get returns information of the data from your warehouse or purchased data.
	Get(context.Context, *DataId) (*DataResult, error)
	GetBatch(context.Context, *BatchRequest) (*GetBatchResult, error)
	// SetPermission allows a consumer to access the given data.
	SetPermission(context.Context, *SetDataPermissionRequest) (*empty.Empty, error)
	SetPermissionBatch(context.Context, *SetBatchDataPermissionRequest) (*empty.Empty, error)
	// Delete removes specific data from your warehouse or purchased data.
	Delete(context.Context, *DataId) (*empty.Empty, error)
	DeleteBatch(context.Context, *BatchRequest) (*empty.Empty, error)
	// Select receives stream of data ID and creates a batch of data IDs.
	// the batch can be used as a input of batch operations.
	Select(Data_SelectServer) error
	// Release unselects and removes the batch.
	Release(context.Context, *BatchRequest) (*empty.Empty, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Get(ctx, req.(*DataId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_SetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDataPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).SetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/SetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).SetPermission(ctx, req.(*SetDataPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_SetPermissionBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBatchDataPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).SetPermissionBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/SetPermissionBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).SetPermissionBatch(ctx, req.(*SetBatchDataPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Delete(ctx, req.(*DataId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_DeleteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).DeleteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/DeleteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).DeleteBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Select_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).Select(&dataSelectServer{stream})
}

type Data_SelectServer interface {
	SendAndClose(*BatchRequest) error
	Recv() (*DataId, error)
	grpc.ServerStream
}

type dataSelectServer struct {
	grpc.ServerStream
}

func (x *dataSelectServer) SendAndClose(m *BatchRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataSelectServer) Recv() (*DataId, error) {
	m := new(DataId)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Data_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Data/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Release(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Data_Get_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _Data_GetBatch_Handler,
		},
		{
			MethodName: "SetPermission",
			Handler:    _Data_SetPermission_Handler,
		},
		{
			MethodName: "SetPermissionBatch",
			Handler:    _Data_SetPermissionBatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Data_Delete_Handler,
		},
		{
			MethodName: "DeleteBatch",
			Handler:    _Data_DeleteBatch_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _Data_Release_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Select",
			Handler:       _Data_Select_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/rpc/v1/server/data.proto",
}
