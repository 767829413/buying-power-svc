// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lots.proto

package kafka

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

type LotUpdatedEvent_Type int32

const (
	LotUpdatedEvent_UNKNOWN       LotUpdatedEvent_Type = 0
	LotUpdatedEvent_FINISHED      LotUpdatedEvent_Type = 1
	LotUpdatedEvent_SOLD          LotUpdatedEvent_Type = 2
	LotUpdatedEvent_INVOICE_PAID  LotUpdatedEvent_Type = 3
	LotUpdatedEvent_DEAL_CANCELED LotUpdatedEvent_Type = 4
)

var LotUpdatedEvent_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "FINISHED",
	2: "SOLD",
	3: "INVOICE_PAID",
	4: "DEAL_CANCELED",
}

var LotUpdatedEvent_Type_value = map[string]int32{
	"UNKNOWN":       0,
	"FINISHED":      1,
	"SOLD":          2,
	"INVOICE_PAID":  3,
	"DEAL_CANCELED": 4,
}

func (x LotUpdatedEvent_Type) String() string {
	return proto.EnumName(LotUpdatedEvent_Type_name, int32(x))
}

func (LotUpdatedEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5bf3a7ea55784cef, []int{0, 0}
}

type LotUpdatedEvent struct {
	Type                 LotUpdatedEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.LotUpdatedEvent_Type" json:"type,omitempty"`
	LotId                int64                `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LotUpdatedEvent) Reset()         { *m = LotUpdatedEvent{} }
func (m *LotUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*LotUpdatedEvent) ProtoMessage()    {}
func (*LotUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf3a7ea55784cef, []int{0}
}

func (m *LotUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LotUpdatedEvent.Unmarshal(m, b)
}
func (m *LotUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LotUpdatedEvent.Marshal(b, m, deterministic)
}
func (m *LotUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotUpdatedEvent.Merge(m, src)
}
func (m *LotUpdatedEvent) XXX_Size() int {
	return xxx_messageInfo_LotUpdatedEvent.Size(m)
}
func (m *LotUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LotUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LotUpdatedEvent proto.InternalMessageInfo

func (m *LotUpdatedEvent) GetType() LotUpdatedEvent_Type {
	if m != nil {
		return m.Type
	}
	return LotUpdatedEvent_UNKNOWN
}

func (m *LotUpdatedEvent) GetLotId() int64 {
	if m != nil {
		return m.LotId
	}
	return 0
}

func init() {
	proto.RegisterEnum("events.LotUpdatedEvent_Type", LotUpdatedEvent_Type_name, LotUpdatedEvent_Type_value)
	proto.RegisterType((*LotUpdatedEvent)(nil), "events.LotUpdatedEvent")
}

func init() { proto.RegisterFile("lots.proto", fileDescriptor_5bf3a7ea55784cef) }

var fileDescriptor_5bf3a7ea55784cef = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x2f, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x56, 0x5a,
	0xc3, 0xc8, 0xc5, 0xef, 0x93, 0x5f, 0x12, 0x5a, 0x90, 0x92, 0x58, 0x92, 0x9a, 0xe2, 0x0a, 0x12,
	0x14, 0x32, 0xe0, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x92,
	0xd1, 0x83, 0x28, 0xd5, 0x43, 0x53, 0xa6, 0x17, 0x52, 0x59, 0x90, 0x1a, 0x04, 0x56, 0x29, 0x24,
	0xca, 0xc5, 0x96, 0x93, 0x5f, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4,
	0x9a, 0x93, 0x5f, 0xe2, 0x99, 0xa2, 0x14, 0xc0, 0xc5, 0x02, 0x52, 0x24, 0xc4, 0xcd, 0xc5, 0x1e,
	0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x20, 0xc4, 0xc3, 0xc5, 0xe1, 0xe6, 0xe9, 0xe7,
	0x19, 0xec, 0xe1, 0xea, 0x22, 0xc0, 0x28, 0xc4, 0xc1, 0xc5, 0x12, 0xec, 0xef, 0xe3, 0x22, 0xc0,
	0x24, 0x24, 0xc0, 0xc5, 0xe3, 0xe9, 0x17, 0xe6, 0xef, 0xe9, 0xec, 0x1a, 0x1f, 0xe0, 0xe8, 0xe9,
	0x22, 0xc0, 0x2c, 0x24, 0xc8, 0xc5, 0xeb, 0xe2, 0xea, 0xe8, 0x13, 0xef, 0xec, 0xe8, 0xe7, 0xec,
	0xea, 0xe3, 0xea, 0x22, 0xc0, 0xe2, 0xa4, 0x1a, 0xa5, 0x9c, 0x9e, 0x59, 0x92, 0x93, 0x98, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x9f, 0xea, 0x58, 0x9a, 0x5c, 0x92, 0x99, 0x9f, 0xa7, 0x0f, 0x71, 0xa1,
	0x7e, 0x7a, 0xbe, 0x7e, 0x76, 0x62, 0x5a, 0x76, 0x62, 0x12, 0x1b, 0xd8, 0x93, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6b, 0x14, 0x0c, 0x4d, 0xf2, 0x00, 0x00, 0x00,
}
