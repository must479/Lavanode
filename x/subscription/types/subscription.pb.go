// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/subscription.proto

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

type Subscription struct {
	Creator            string              `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Consumer           string              `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Block              uint64              `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	PlanIndex          string              `protobuf:"bytes,4,opt,name=plan_index,json=planIndex,proto3" json:"plan_index,omitempty"`
	PlanBlock          uint64              `protobuf:"varint,5,opt,name=plan_block,json=planBlock,proto3" json:"plan_block,omitempty"`
	DurationBought     uint64              `protobuf:"varint,6,opt,name=duration_bought,json=durationBought,proto3" json:"duration_bought,omitempty"`
	DurationLeft       uint64              `protobuf:"varint,7,opt,name=duration_left,json=durationLeft,proto3" json:"duration_left,omitempty"`
	MonthExpiryTime    uint64              `protobuf:"varint,8,opt,name=month_expiry_time,json=monthExpiryTime,proto3" json:"month_expiry_time,omitempty"`
	MonthCuTotal       uint64              `protobuf:"varint,10,opt,name=month_cu_total,json=monthCuTotal,proto3" json:"month_cu_total,omitempty"`
	MonthCuLeft        uint64              `protobuf:"varint,11,opt,name=month_cu_left,json=monthCuLeft,proto3" json:"month_cu_left,omitempty"`
	Cluster            string              `protobuf:"bytes,13,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DurationTotal      uint64              `protobuf:"varint,14,opt,name=duration_total,json=durationTotal,proto3" json:"duration_total,omitempty"`
	AutoRenewal        bool                `protobuf:"varint,15,opt,name=auto_renewal,json=autoRenewal,proto3" json:"auto_renewal,omitempty"`
	FutureSubscription *FutureSubscription `protobuf:"bytes,16,opt,name=future_subscription,json=futureSubscription,proto3" json:"future_subscription,omitempty"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3bc5507ca237d79, []int{0}
}
func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return m.Size()
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Subscription) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *Subscription) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Subscription) GetPlanIndex() string {
	if m != nil {
		return m.PlanIndex
	}
	return ""
}

func (m *Subscription) GetPlanBlock() uint64 {
	if m != nil {
		return m.PlanBlock
	}
	return 0
}

func (m *Subscription) GetDurationBought() uint64 {
	if m != nil {
		return m.DurationBought
	}
	return 0
}

func (m *Subscription) GetDurationLeft() uint64 {
	if m != nil {
		return m.DurationLeft
	}
	return 0
}

func (m *Subscription) GetMonthExpiryTime() uint64 {
	if m != nil {
		return m.MonthExpiryTime
	}
	return 0
}

func (m *Subscription) GetMonthCuTotal() uint64 {
	if m != nil {
		return m.MonthCuTotal
	}
	return 0
}

func (m *Subscription) GetMonthCuLeft() uint64 {
	if m != nil {
		return m.MonthCuLeft
	}
	return 0
}

func (m *Subscription) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Subscription) GetDurationTotal() uint64 {
	if m != nil {
		return m.DurationTotal
	}
	return 0
}

func (m *Subscription) GetAutoRenewal() bool {
	if m != nil {
		return m.AutoRenewal
	}
	return false
}

func (m *Subscription) GetFutureSubscription() *FutureSubscription {
	if m != nil {
		return m.FutureSubscription
	}
	return nil
}

type FutureSubscription struct {
	PlanIndex      string `protobuf:"bytes,1,opt,name=plan_index,json=planIndex,proto3" json:"plan_index,omitempty"`
	PlanBlock      uint64 `protobuf:"varint,2,opt,name=plan_block,json=planBlock,proto3" json:"plan_block,omitempty"`
	DurationBought uint64 `protobuf:"varint,3,opt,name=duration_bought,json=durationBought,proto3" json:"duration_bought,omitempty"`
}

func (m *FutureSubscription) Reset()         { *m = FutureSubscription{} }
func (m *FutureSubscription) String() string { return proto.CompactTextString(m) }
func (*FutureSubscription) ProtoMessage()    {}
func (*FutureSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3bc5507ca237d79, []int{1}
}
func (m *FutureSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FutureSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FutureSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FutureSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FutureSubscription.Merge(m, src)
}
func (m *FutureSubscription) XXX_Size() int {
	return m.Size()
}
func (m *FutureSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_FutureSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_FutureSubscription proto.InternalMessageInfo

func (m *FutureSubscription) GetPlanIndex() string {
	if m != nil {
		return m.PlanIndex
	}
	return ""
}

func (m *FutureSubscription) GetPlanBlock() uint64 {
	if m != nil {
		return m.PlanBlock
	}
	return 0
}

func (m *FutureSubscription) GetDurationBought() uint64 {
	if m != nil {
		return m.DurationBought
	}
	return 0
}

func init() {
	proto.RegisterType((*Subscription)(nil), "lavanet.lava.subscription.Subscription")
	proto.RegisterType((*FutureSubscription)(nil), "lavanet.lava.subscription.FutureSubscription")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/subscription.proto", fileDescriptor_c3bc5507ca237d79)
}

var fileDescriptor_c3bc5507ca237d79 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0xad, 0xdb, 0x3a, 0x93, 0xbf, 0x2c, 0x1c, 0x16, 0x24, 0xac, 0x10, 0x40, 0x44,
	0xa8, 0x38, 0x12, 0xbc, 0x41, 0x10, 0x95, 0xa8, 0x38, 0x99, 0x9e, 0x38, 0x60, 0xd9, 0xee, 0xa6,
	0xb1, 0xb0, 0xbd, 0xd6, 0x7a, 0x16, 0x52, 0xf1, 0x12, 0x3c, 0x08, 0x0f, 0xc2, 0xb1, 0x47, 0x8e,
	0x28, 0x79, 0x11, 0xe4, 0x71, 0x62, 0x39, 0x04, 0x50, 0x4f, 0xab, 0xf9, 0xcd, 0xf7, 0x79, 0xbc,
	0xfb, 0x69, 0xe0, 0x34, 0x09, 0x3e, 0x07, 0x99, 0xc4, 0x69, 0x79, 0x4e, 0x0b, 0x13, 0x16, 0x91,
	0x8e, 0x73, 0x8c, 0x55, 0xb6, 0x53, 0xb8, 0xb9, 0x56, 0xa8, 0xf8, 0xfd, 0x8d, 0xda, 0x2d, 0x4f,
	0xb7, 0x29, 0x18, 0x7f, 0xb7, 0xa0, 0xfb, 0xbe, 0x01, 0xb8, 0x80, 0x93, 0x48, 0xcb, 0x00, 0x95,
	0x16, 0x6c, 0xc4, 0x26, 0x6d, 0x6f, 0x5b, 0xf2, 0x07, 0x60, 0x47, 0x2a, 0x2b, 0x4c, 0x2a, 0xb5,
	0x38, 0xa0, 0x56, 0x5d, 0xf3, 0x7b, 0x70, 0x14, 0x26, 0x2a, 0xfa, 0x24, 0x0e, 0x47, 0x6c, 0x62,
	0x79, 0x55, 0xc1, 0x1f, 0x02, 0xe4, 0x49, 0x90, 0xf9, 0x71, 0x76, 0x29, 0x97, 0xc2, 0x22, 0x4f,
	0xbb, 0x24, 0x6f, 0x4b, 0x50, 0xb7, 0x2b, 0xe7, 0x11, 0x39, 0xa9, 0x3d, 0x23, 0xf7, 0x33, 0x18,
	0x5c, 0x1a, 0x1d, 0x94, 0x7f, 0xe5, 0x87, 0xca, 0x5c, 0x2d, 0x50, 0x1c, 0x93, 0xa6, 0xbf, 0xc5,
	0x33, 0xa2, 0xfc, 0x31, 0xf4, 0x6a, 0x61, 0x22, 0xe7, 0x28, 0x4e, 0x48, 0xd6, 0xdd, 0xc2, 0x77,
	0x72, 0x8e, 0xfc, 0x39, 0xdc, 0x49, 0x55, 0x86, 0x0b, 0x5f, 0x2e, 0xf3, 0x58, 0x5f, 0xfb, 0x18,
	0xa7, 0x52, 0xd8, 0x24, 0x1c, 0x50, 0xe3, 0x0d, 0xf1, 0x8b, 0x38, 0x95, 0xfc, 0x09, 0xf4, 0x2b,
	0x6d, 0x64, 0x7c, 0x54, 0x18, 0x24, 0x02, 0xaa, 0x2f, 0x12, 0x7d, 0x6d, 0x2e, 0x4a, 0xc6, 0xc7,
	0xd0, 0xab, 0x55, 0x34, 0xb6, 0x43, 0xa2, 0xce, 0x46, 0x44, 0x53, 0xcb, 0xd7, 0x4c, 0x4c, 0x81,
	0x52, 0x8b, 0xde, 0xe6, 0x35, 0xab, 0x92, 0x3f, 0x85, 0xfa, 0x1a, 0x9b, 0x19, 0x7d, 0xb2, 0xd7,
	0x57, 0xa9, 0x86, 0x3c, 0x82, 0x6e, 0x60, 0x50, 0xf9, 0x5a, 0x66, 0xf2, 0x4b, 0x90, 0x88, 0xc1,
	0x88, 0x4d, 0x6c, 0xaf, 0x53, 0x32, 0xaf, 0x42, 0xfc, 0x23, 0xdc, 0x9d, 0x1b, 0x34, 0x5a, 0xfa,
	0xcd, 0x64, 0xc5, 0x70, 0xc4, 0x26, 0x9d, 0x97, 0x2f, 0xdc, 0x7f, 0x66, 0xef, 0x9e, 0x91, 0xab,
	0x99, 0xbe, 0xc7, 0xe7, 0x7b, 0xec, 0xdc, 0xb2, 0xdb, 0x43, 0x38, 0xb7, 0xec, 0xee, 0xb0, 0x37,
	0xfe, 0x0a, 0x7c, 0xdf, 0xf5, 0x47, 0xce, 0xec, 0xff, 0x39, 0x1f, 0xdc, 0x22, 0xe7, 0xc3, 0xbf,
	0xe5, 0x3c, 0x3b, 0xfb, 0xb1, 0x72, 0xd8, 0xcd, 0xca, 0x61, 0xbf, 0x56, 0x0e, 0xfb, 0xb6, 0x76,
	0x5a, 0x37, 0x6b, 0xa7, 0xf5, 0x73, 0xed, 0xb4, 0x3e, 0x9c, 0x5e, 0xc5, 0xb8, 0x30, 0xa1, 0x1b,
	0xa9, 0x74, 0xba, 0xb3, 0x19, 0xcb, 0xdd, 0xdd, 0xc0, 0xeb, 0x5c, 0x16, 0xe1, 0x31, 0x6d, 0xc5,
	0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0xeb, 0x7c, 0x12, 0x45, 0x03, 0x00, 0x00,
}

func (m *Subscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FutureSubscription != nil {
		{
			size, err := m.FutureSubscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubscription(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.AutoRenewal {
		i--
		if m.AutoRenewal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if m.DurationTotal != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationTotal))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0x6a
	}
	if m.MonthCuLeft != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthCuLeft))
		i--
		dAtA[i] = 0x58
	}
	if m.MonthCuTotal != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthCuTotal))
		i--
		dAtA[i] = 0x50
	}
	if m.MonthExpiryTime != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthExpiryTime))
		i--
		dAtA[i] = 0x40
	}
	if m.DurationLeft != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationLeft))
		i--
		dAtA[i] = 0x38
	}
	if m.DurationBought != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationBought))
		i--
		dAtA[i] = 0x30
	}
	if m.PlanBlock != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PlanIndex) > 0 {
		i -= len(m.PlanIndex)
		copy(dAtA[i:], m.PlanIndex)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.PlanIndex)))
		i--
		dAtA[i] = 0x22
	}
	if m.Block != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FutureSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FutureSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FutureSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DurationBought != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationBought))
		i--
		dAtA[i] = 0x18
	}
	if m.PlanBlock != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanBlock))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PlanIndex) > 0 {
		i -= len(m.PlanIndex)
		copy(dAtA[i:], m.PlanIndex)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.PlanIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovSubscription(uint64(m.Block))
	}
	l = len(m.PlanIndex)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanBlock != 0 {
		n += 1 + sovSubscription(uint64(m.PlanBlock))
	}
	if m.DurationBought != 0 {
		n += 1 + sovSubscription(uint64(m.DurationBought))
	}
	if m.DurationLeft != 0 {
		n += 1 + sovSubscription(uint64(m.DurationLeft))
	}
	if m.MonthExpiryTime != 0 {
		n += 1 + sovSubscription(uint64(m.MonthExpiryTime))
	}
	if m.MonthCuTotal != 0 {
		n += 1 + sovSubscription(uint64(m.MonthCuTotal))
	}
	if m.MonthCuLeft != 0 {
		n += 1 + sovSubscription(uint64(m.MonthCuLeft))
	}
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.DurationTotal != 0 {
		n += 1 + sovSubscription(uint64(m.DurationTotal))
	}
	if m.AutoRenewal {
		n += 2
	}
	if m.FutureSubscription != nil {
		l = m.FutureSubscription.Size()
		n += 2 + l + sovSubscription(uint64(l))
	}
	return n
}

func (m *FutureSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlanIndex)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanBlock != 0 {
		n += 1 + sovSubscription(uint64(m.PlanBlock))
	}
	if m.DurationBought != 0 {
		n += 1 + sovSubscription(uint64(m.DurationBought))
	}
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: Subscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanBlock", wireType)
			}
			m.PlanBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationBought", wireType)
			}
			m.DurationBought = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationBought |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationLeft", wireType)
			}
			m.DurationLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthExpiryTime", wireType)
			}
			m.MonthExpiryTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthExpiryTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthCuTotal", wireType)
			}
			m.MonthCuTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthCuTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthCuLeft", wireType)
			}
			m.MonthCuLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthCuLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationTotal", wireType)
			}
			m.DurationTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoRenewal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoRenewal = bool(v != 0)
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureSubscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FutureSubscription == nil {
				m.FutureSubscription = &FutureSubscription{}
			}
			if err := m.FutureSubscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func (m *FutureSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: FutureSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FutureSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanBlock", wireType)
			}
			m.PlanBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationBought", wireType)
			}
			m.DurationBought = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationBought |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)
