// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/adjustment.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Adjustment struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	AdjustedUsage uint64 `protobuf:"varint,2,opt,name=adjustedUsage,proto3" json:"adjustedUsage,omitempty"`
	TotalUsage    uint64 `protobuf:"varint,3,opt,name=totalUsage,proto3" json:"totalUsage,omitempty"`
}

func (m *Adjustment) Reset()         { *m = Adjustment{} }
func (m *Adjustment) String() string { return proto.CompactTextString(m) }
func (*Adjustment) ProtoMessage()    {}
func (*Adjustment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6061843cba96837b, []int{0}
}
func (m *Adjustment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Adjustment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Adjustment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Adjustment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adjustment.Merge(m, src)
}
func (m *Adjustment) XXX_Size() int {
	return m.Size()
}
func (m *Adjustment) XXX_DiscardUnknown() {
	xxx_messageInfo_Adjustment.DiscardUnknown(m)
}

var xxx_messageInfo_Adjustment proto.InternalMessageInfo

func (m *Adjustment) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Adjustment) GetAdjustedUsage() uint64 {
	if m != nil {
		return m.AdjustedUsage
	}
	return 0
}

func (m *Adjustment) GetTotalUsage() uint64 {
	if m != nil {
		return m.TotalUsage
	}
	return 0
}

func init() {
	proto.RegisterType((*Adjustment)(nil), "lavanet.lava.subscription.Adjustment")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/adjustment.proto", fileDescriptor_6061843cba96837b)
}

var fileDescriptor_6061843cba96837b = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xc5, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25,
	0x99, 0xf9, 0x79, 0xfa, 0x89, 0x29, 0x59, 0xa5, 0xc5, 0x25, 0xb9, 0xa9, 0x79, 0x25, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x50, 0xb5, 0x7a, 0x20, 0x5a, 0x0f, 0x59, 0xad, 0x52, 0x06,
	0x17, 0x97, 0x23, 0x5c, 0xb9, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0xa4, 0xc2, 0xc5, 0x0b, 0x31, 0x32, 0x35, 0x25, 0xb4,
	0x38, 0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0x55, 0x50, 0x48, 0x8e, 0x8b,
	0xab, 0x24, 0xbf, 0x24, 0x31, 0x07, 0xa2, 0x84, 0x19, 0xac, 0x04, 0x49, 0xc4, 0xc9, 0xed, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x51, 0x7c, 0x55, 0x81, 0xea, 0xaf, 0x92, 0xca, 0x82, 0xd4, 0xe2,
	0x24, 0x36, 0xb0, 0x9f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x8e, 0x73, 0xb8, 0x01,
	0x01, 0x00, 0x00,
}

func (m *Adjustment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Adjustment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Adjustment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalUsage != 0 {
		i = encodeVarintAdjustment(dAtA, i, uint64(m.TotalUsage))
		i--
		dAtA[i] = 0x18
	}
	if m.AdjustedUsage != 0 {
		i = encodeVarintAdjustment(dAtA, i, uint64(m.AdjustedUsage))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintAdjustment(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdjustment(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdjustment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Adjustment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovAdjustment(uint64(l))
	}
	if m.AdjustedUsage != 0 {
		n += 1 + sovAdjustment(uint64(m.AdjustedUsage))
	}
	if m.TotalUsage != 0 {
		n += 1 + sovAdjustment(uint64(m.TotalUsage))
	}
	return n
}

func sovAdjustment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdjustment(x uint64) (n int) {
	return sovAdjustment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Adjustment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdjustment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Adjustment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Adjustment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdjustment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdjustment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdjustment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustedUsage", wireType)
			}
			m.AdjustedUsage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdjustment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjustedUsage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUsage", wireType)
			}
			m.TotalUsage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdjustment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalUsage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAdjustment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdjustment
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
func skipAdjustment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdjustment
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
					return 0, ErrIntOverflowAdjustment
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdjustment
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
			if length < 0 {
				return 0, ErrInvalidLengthAdjustment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdjustment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdjustment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdjustment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdjustment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdjustment = fmt.Errorf("proto: unexpected end of group")
)
