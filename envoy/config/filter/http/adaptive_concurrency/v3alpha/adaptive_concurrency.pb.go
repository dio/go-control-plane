// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto

package envoy_config_filter_http_adaptive_concurrency_v3alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AdaptiveConcurrency struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdaptiveConcurrency) Reset()         { *m = AdaptiveConcurrency{} }
func (m *AdaptiveConcurrency) String() string { return proto.CompactTextString(m) }
func (*AdaptiveConcurrency) ProtoMessage()    {}
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_8176400d3397e66d, []int{0}
}

func (m *AdaptiveConcurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdaptiveConcurrency.Unmarshal(m, b)
}
func (m *AdaptiveConcurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdaptiveConcurrency.Marshal(b, m, deterministic)
}
func (m *AdaptiveConcurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdaptiveConcurrency.Merge(m, src)
}
func (m *AdaptiveConcurrency) XXX_Size() int {
	return xxx_messageInfo_AdaptiveConcurrency.Size(m)
}
func (m *AdaptiveConcurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AdaptiveConcurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AdaptiveConcurrency proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdaptiveConcurrency)(nil), "envoy.config.filter.http.adaptive_concurrency.v3alpha.AdaptiveConcurrency")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto", fileDescriptor_8176400d3397e66d)
}

var fileDescriptor_8176400d3397e66d = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x4c, 0x49, 0x2c, 0x28, 0xc9, 0x2c, 0x4b, 0x8d, 0x4f, 0xce,
	0xcf, 0x4b, 0x2e, 0x2d, 0x2a, 0x4a, 0xcd, 0x4b, 0xae, 0xd4, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8,
	0x48, 0xc4, 0x2a, 0xa9, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x64, 0x0a, 0x36, 0x51, 0x0f, 0x62,
	0xa2, 0x1e, 0xc4, 0x44, 0x3d, 0x90, 0x89, 0x7a, 0x58, 0x35, 0x41, 0x4d, 0x54, 0x12, 0xe5, 0x12,
	0x76, 0x84, 0xca, 0x3b, 0x23, 0xa4, 0x9d, 0x12, 0xb9, 0x9c, 0x33, 0xf3, 0xf5, 0xc0, 0x46, 0x16,
	0x14, 0xe5, 0x57, 0x54, 0xea, 0x91, 0x65, 0xba, 0x93, 0x04, 0x16, 0xb3, 0x03, 0x40, 0xce, 0x0d,
	0x60, 0x4c, 0x62, 0x03, 0xbb, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xb3, 0x1b, 0x4a,
	0x0b, 0x01, 0x00, 0x00,
}