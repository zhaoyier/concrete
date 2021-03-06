// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/etwebapi/webapi.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = strings.TrimPrefix

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	//
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (m *HelloRequest) Validate() error {
	return nil
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a03d461c0aaef6a, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	// common.Result result = 1;
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (m *HelloReply) Validate() error {
	return nil
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a03d461c0aaef6a, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetPictureListReq struct {
	//
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	//
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
}

func (m *GetPictureListReq) Validate() error {
	return nil
}

func (m *GetPictureListReq) Reset()         { *m = GetPictureListReq{} }
func (m *GetPictureListReq) String() string { return proto.CompactTextString(m) }
func (*GetPictureListReq) ProtoMessage()    {}
func (*GetPictureListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a03d461c0aaef6a, []int{2}
}

func (m *GetPictureListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPictureListReq.Unmarshal(m, b)
}
func (m *GetPictureListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPictureListReq.Marshal(b, m, deterministic)
}
func (m *GetPictureListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPictureListReq.Merge(m, src)
}
func (m *GetPictureListReq) XXX_Size() int {
	return xxx_messageInfo_GetPictureListReq.Size(m)
}
func (m *GetPictureListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPictureListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPictureListReq proto.InternalMessageInfo

func (m *GetPictureListReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetPictureListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetPictureListResp struct {
	//
	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode"`
	//
	Data []*Picture `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
	//
	ErrMsg string `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg"`
}

func (m *GetPictureListResp) Validate() error {
	return nil
}

func (m *GetPictureListResp) Reset()         { *m = GetPictureListResp{} }
func (m *GetPictureListResp) String() string { return proto.CompactTextString(m) }
func (*GetPictureListResp) ProtoMessage()    {}
func (*GetPictureListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a03d461c0aaef6a, []int{3}
}

func (m *GetPictureListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPictureListResp.Unmarshal(m, b)
}
func (m *GetPictureListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPictureListResp.Marshal(b, m, deterministic)
}
func (m *GetPictureListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPictureListResp.Merge(m, src)
}
func (m *GetPictureListResp) XXX_Size() int {
	return xxx_messageInfo_GetPictureListResp.Size(m)
}
func (m *GetPictureListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPictureListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPictureListResp proto.InternalMessageInfo

func (m *GetPictureListResp) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetPictureListResp) GetData() []*Picture {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetPictureListResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type Picture struct {
	//
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	//
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	//
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	//
	ImgNum int32 `protobuf:"varint,4,opt,name=imgNum,proto3" json:"imgNum"`
	//
	ImgSrc string `protobuf:"bytes,5,opt,name=imgSrc,proto3" json:"imgSrc"`
}

func (m *Picture) Validate() error {
	return nil
}

func (m *Picture) Reset()         { *m = Picture{} }
func (m *Picture) String() string { return proto.CompactTextString(m) }
func (*Picture) ProtoMessage()    {}
func (*Picture) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a03d461c0aaef6a, []int{4}
}

func (m *Picture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Picture.Unmarshal(m, b)
}
func (m *Picture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Picture.Marshal(b, m, deterministic)
}
func (m *Picture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Picture.Merge(m, src)
}
func (m *Picture) XXX_Size() int {
	return xxx_messageInfo_Picture.Size(m)
}
func (m *Picture) XXX_DiscardUnknown() {
	xxx_messageInfo_Picture.DiscardUnknown(m)
}

var xxx_messageInfo_Picture proto.InternalMessageInfo

func (m *Picture) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Picture) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Picture) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Picture) GetImgNum() int32 {
	if m != nil {
		return m.ImgNum
	}
	return 0
}

func (m *Picture) GetImgSrc() string {
	if m != nil {
		return m.ImgSrc
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "etwebapi.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "etwebapi.HelloReply")
	proto.RegisterType((*GetPictureListReq)(nil), "etwebapi.GetPictureListReq")
	proto.RegisterType((*GetPictureListResp)(nil), "etwebapi.GetPictureListResp")
	proto.RegisterType((*Picture)(nil), "etwebapi.Picture")
}

func init() { proto.RegisterFile("proto/etwebapi/webapi.proto", fileDescriptor_0a03d461c0aaef6a) }

var fileDescriptor_0a03d461c0aaef6a = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x9a, 0xb4, 0x3a, 0x8a, 0xd0, 0x41, 0x64, 0x41, 0x90, 0xb2, 0xa0, 0xf4, 0x94,
	0x82, 0x9e, 0x3d, 0x79, 0xd0, 0x83, 0x8a, 0xac, 0x9f, 0x60, 0xdb, 0x0c, 0x61, 0x21, 0x31, 0x69,
	0x66, 0x82, 0xf8, 0xed, 0x25, 0xfb, 0x47, 0xc5, 0x53, 0xe6, 0xbd, 0xbc, 0xfd, 0xed, 0xcc, 0x2c,
	0x5c, 0xf6, 0x43, 0x27, 0xdd, 0x96, 0xe4, 0x93, 0x76, 0xb6, 0x77, 0xdb, 0xf0, 0x29, 0xbd, 0x8b,
	0x47, 0xc9, 0xd6, 0x1a, 0x4e, 0x9f, 0xa8, 0x69, 0x3a, 0x43, 0x87, 0x91, 0x58, 0x10, 0x21, 0xff,
	0xb0, 0x2d, 0xa9, 0xd9, 0x7a, 0xb6, 0x39, 0x36, 0xbe, 0xd6, 0x37, 0x00, 0x31, 0xd3, 0x37, 0x5f,
	0xa8, 0x60, 0xd9, 0x12, 0xb3, 0xad, 0x49, 0x65, 0x3e, 0x94, 0xa4, 0xbe, 0x87, 0xd5, 0x23, 0xc9,
	0x9b, 0xdb, 0xcb, 0x38, 0xd0, 0xb3, 0x63, 0x31, 0x74, 0x98, 0x80, 0xfd, 0x94, 0x9d, 0x80, 0x85,
	0xf1, 0x35, 0x9e, 0x43, 0xd1, 0xb8, 0xd6, 0x89, 0x07, 0x14, 0x26, 0x08, 0xcd, 0x80, 0xff, 0x8f,
	0x73, 0x8f, 0x57, 0x00, 0x2c, 0x56, 0x46, 0x7e, 0xe8, 0xaa, 0x44, 0xf9, 0xe3, 0xe0, 0x35, 0xe4,
	0x95, 0x15, 0xab, 0xb2, 0xf5, 0x7c, 0x73, 0x72, 0xbb, 0x2a, 0xd3, 0x64, 0x65, 0x04, 0x19, 0xff,
	0x1b, 0x2f, 0x60, 0x41, 0xc3, 0xf0, 0xc2, 0xb5, 0x9a, 0xfb, 0xa6, 0xa3, 0xd2, 0x0c, 0xcb, 0x18,
	0xc4, 0x33, 0xc8, 0x5c, 0x15, 0x07, 0xcf, 0x5c, 0x35, 0x75, 0x29, 0x4e, 0x9a, 0x34, 0x66, 0x10,
	0x3f, 0x0b, 0x9a, 0xff, 0x2e, 0x68, 0x82, 0xbb, 0xb6, 0x7e, 0x1d, 0x5b, 0x95, 0xfb, 0xfe, 0xa2,
	0x8a, 0xfe, 0xfb, 0xb0, 0x57, 0x45, 0xb8, 0x34, 0xa8, 0xdd, 0xc2, 0xbf, 0xc2, 0xdd, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x77, 0xc1, 0xa3, 0xa4, 0x01, 0x00, 0x00,
}
