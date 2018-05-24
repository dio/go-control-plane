// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/rbac/v2/rbac.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/rbac/v2/rbac.proto

	It has these top-level messages:
		RBAC
		RBACPerRoute
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_config_rbac_v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2alpha"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RBAC struct {
	// Specify the RBAC rules to be applied globally
	Rules *envoy_config_rbac_v2alpha.RBAC `protobuf:"bytes,1,opt,name=rules" json:"rules,omitempty"`
}

func (m *RBAC) Reset()                    { *m = RBAC{} }
func (m *RBAC) String() string            { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()               {}
func (*RBAC) Descriptor() ([]byte, []int) { return fileDescriptorRbac, []int{0} }

func (m *RBAC) GetRules() *envoy_config_rbac_v2alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

type RBACPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*RBACPerRoute_Disabled
	//	*RBACPerRoute_Rbac
	Override isRBACPerRoute_Override `protobuf_oneof:"override"`
}

func (m *RBACPerRoute) Reset()                    { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string            { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()               {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) { return fileDescriptorRbac, []int{1} }

type isRBACPerRoute_Override interface {
	isRBACPerRoute_Override()
	MarshalTo([]byte) (int, error)
	Size() int
}

type RBACPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}
type RBACPerRoute_Rbac struct {
	Rbac *RBAC `protobuf:"bytes,2,opt,name=rbac,oneof"`
}

func (*RBACPerRoute_Disabled) isRBACPerRoute_Override() {}
func (*RBACPerRoute_Rbac) isRBACPerRoute_Override()     {}

func (m *RBACPerRoute) GetOverride() isRBACPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *RBACPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*RBACPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *RBACPerRoute) GetRbac() *RBAC {
	if x, ok := m.GetOverride().(*RBACPerRoute_Rbac); ok {
		return x.Rbac
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RBACPerRoute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RBACPerRoute_OneofMarshaler, _RBACPerRoute_OneofUnmarshaler, _RBACPerRoute_OneofSizer, []interface{}{
		(*RBACPerRoute_Disabled)(nil),
		(*RBACPerRoute_Rbac)(nil),
	}
}

func _RBACPerRoute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RBACPerRoute)
	// override
	switch x := m.Override.(type) {
	case *RBACPerRoute_Disabled:
		t := uint64(0)
		if x.Disabled {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *RBACPerRoute_Rbac:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rbac); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RBACPerRoute.Override has unexpected type %T", x)
	}
	return nil
}

func _RBACPerRoute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RBACPerRoute)
	switch tag {
	case 1: // override.disabled
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Override = &RBACPerRoute_Disabled{x != 0}
		return true, err
	case 2: // override.rbac
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RBAC)
		err := b.DecodeMessage(msg)
		m.Override = &RBACPerRoute_Rbac{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RBACPerRoute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RBACPerRoute)
	// override
	switch x := m.Override.(type) {
	case *RBACPerRoute_Disabled:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *RBACPerRoute_Rbac:
		s := proto.Size(x.Rbac)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.http.rbac.v2.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.config.filter.http.rbac.v2.RBACPerRoute")
}
func (m *RBAC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBAC) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Rules != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRbac(dAtA, i, uint64(m.Rules.Size()))
		n1, err := m.Rules.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *RBACPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBACPerRoute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Override != nil {
		nn2, err := m.Override.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *RBACPerRoute_Disabled) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *RBACPerRoute_Rbac) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Rbac != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRbac(dAtA, i, uint64(m.Rbac.Size()))
		n3, err := m.Rbac.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func encodeVarintRbac(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RBAC) Size() (n int) {
	var l int
	_ = l
	if m.Rules != nil {
		l = m.Rules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	return n
}

func (m *RBACPerRoute) Size() (n int) {
	var l int
	_ = l
	if m.Override != nil {
		n += m.Override.Size()
	}
	return n
}

func (m *RBACPerRoute_Disabled) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *RBACPerRoute_Rbac) Size() (n int) {
	var l int
	_ = l
	if m.Rbac != nil {
		l = m.Rbac.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	return n
}

func sovRbac(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRbac(x uint64) (n int) {
	return sovRbac(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RBAC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RBAC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBAC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = &envoy_config_rbac_v2alpha.RBAC{}
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RBACPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RBACPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBACPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Override = &RBACPerRoute_Disabled{b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rbac", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RBAC{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Override = &RBACPerRoute_Rbac{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRbac(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRbac
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRbac
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRbac(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRbac = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRbac   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/config/filter/http/rbac/v2/rbac.proto", fileDescriptorRbac) }

var fileDescriptorRbac = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x4a, 0x4a, 0x4c, 0xd6, 0x2f, 0x33, 0x02, 0xd3, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0x60, 0xc5, 0x7a, 0x10, 0xc5, 0x7a, 0x10, 0xc5, 0x7a, 0x20,
	0xc5, 0x7a, 0x60, 0x45, 0x65, 0x46, 0x52, 0x2a, 0x28, 0xc6, 0x41, 0x8d, 0x48, 0xcc, 0x29, 0xc8,
	0x48, 0x44, 0x32, 0x47, 0x4a, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6,
	0x80, 0x4a, 0x88, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0xc9, 0x93,
	0x8b, 0x25, 0xc8, 0xc9, 0xd1, 0x59, 0xc8, 0x91, 0x8b, 0xb5, 0xa8, 0x34, 0x27, 0xb5, 0x58, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5e, 0x0f, 0xc5, 0x39, 0x50, 0x27, 0x80, 0x2d, 0xd3, 0x03,
	0xa9, 0x77, 0xe2, 0xda, 0xf5, 0xf2, 0x00, 0x33, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0x44,
	0xa7, 0xd2, 0x14, 0x46, 0x2e, 0x1e, 0x90, 0x5c, 0x40, 0x6a, 0x51, 0x50, 0x7e, 0x69, 0x49, 0xaa,
	0x90, 0x3a, 0x17, 0x47, 0x4a, 0x66, 0x71, 0x62, 0x52, 0x4e, 0x6a, 0x0a, 0xd8, 0x58, 0x0e, 0x27,
	0x4e, 0x90, 0x2e, 0x96, 0x2c, 0x26, 0x0e, 0x46, 0x0f, 0x86, 0x20, 0xb8, 0xa4, 0x90, 0x07, 0x17,
	0x0b, 0xc8, 0x06, 0x09, 0x26, 0xb0, 0xdd, 0x6a, 0x7a, 0x84, 0x82, 0x02, 0xc3, 0x09, 0x1e, 0x0c,
	0x41, 0x60, 0x13, 0x9c, 0x04, 0xb9, 0x38, 0xf2, 0xcb, 0x52, 0x8b, 0x8a, 0x32, 0x53, 0x52, 0x85,
	0x58, 0x77, 0xbc, 0x3c, 0xc0, 0xcc, 0xe8, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x31, 0x95, 0x19, 0x25, 0xb1, 0x81, 0xbd, 0x6e, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xee, 0xca, 0xd5, 0xa0, 0x01, 0x00, 0x00,
}
