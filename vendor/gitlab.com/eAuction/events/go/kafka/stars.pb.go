// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stars.proto

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

type Star_EventType int32

const (
	Star_EVENT_TYPE_UNKNOWN  Star_EventType = 0
	Star_EVENT_TYPE_CREATED  Star_EventType = 1
	Star_EVENT_TYPE_DELETED  Star_EventType = 2
	Star_EVENT_TYPE_EXPIRING Star_EventType = 3
	Star_EVENT_TYPE_EXPIRED  Star_EventType = 4
)

var Star_EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNKNOWN",
	1: "EVENT_TYPE_CREATED",
	2: "EVENT_TYPE_DELETED",
	3: "EVENT_TYPE_EXPIRING",
	4: "EVENT_TYPE_EXPIRED",
}

var Star_EventType_value = map[string]int32{
	"EVENT_TYPE_UNKNOWN":  0,
	"EVENT_TYPE_CREATED":  1,
	"EVENT_TYPE_DELETED":  2,
	"EVENT_TYPE_EXPIRING": 3,
	"EVENT_TYPE_EXPIRED":  4,
}

func (x Star_EventType) String() string {
	return proto.EnumName(Star_EventType_name, int32(x))
}

func (Star_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1, 0}
}

type StarEvent struct {
	Type                 Star_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=events.Star_EventType" json:"type,omitempty"`
	Created              *Star_Created  `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Deleted              *Star_Deleted  `protobuf:"bytes,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Expiring             *Star_Expiring `protobuf:"bytes,4,opt,name=expiring,proto3" json:"expiring,omitempty"`
	Expired              *Star_Expired  `protobuf:"bytes,5,opt,name=expired,proto3" json:"expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StarEvent) Reset()         { *m = StarEvent{} }
func (m *StarEvent) String() string { return proto.CompactTextString(m) }
func (*StarEvent) ProtoMessage()    {}
func (*StarEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{0}
}

func (m *StarEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StarEvent.Unmarshal(m, b)
}
func (m *StarEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StarEvent.Marshal(b, m, deterministic)
}
func (m *StarEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StarEvent.Merge(m, src)
}
func (m *StarEvent) XXX_Size() int {
	return xxx_messageInfo_StarEvent.Size(m)
}
func (m *StarEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StarEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StarEvent proto.InternalMessageInfo

func (m *StarEvent) GetType() Star_EventType {
	if m != nil {
		return m.Type
	}
	return Star_EVENT_TYPE_UNKNOWN
}

func (m *StarEvent) GetCreated() *Star_Created {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *StarEvent) GetDeleted() *Star_Deleted {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *StarEvent) GetExpiring() *Star_Expiring {
	if m != nil {
		return m.Expiring
	}
	return nil
}

func (m *StarEvent) GetExpired() *Star_Expired {
	if m != nil {
		return m.Expired
	}
	return nil
}

type Star struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star) Reset()         { *m = Star{} }
func (m *Star) String() string { return proto.CompactTextString(m) }
func (*Star) ProtoMessage()    {}
func (*Star) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1}
}

func (m *Star) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star.Unmarshal(m, b)
}
func (m *Star) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star.Marshal(b, m, deterministic)
}
func (m *Star) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star.Merge(m, src)
}
func (m *Star) XXX_Size() int {
	return xxx_messageInfo_Star.Size(m)
}
func (m *Star) XXX_DiscardUnknown() {
	xxx_messageInfo_Star.DiscardUnknown(m)
}

var xxx_messageInfo_Star proto.InternalMessageInfo

type Star_Created struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LotId                int64    `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star_Created) Reset()         { *m = Star_Created{} }
func (m *Star_Created) String() string { return proto.CompactTextString(m) }
func (*Star_Created) ProtoMessage()    {}
func (*Star_Created) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1, 0}
}

func (m *Star_Created) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star_Created.Unmarshal(m, b)
}
func (m *Star_Created) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star_Created.Marshal(b, m, deterministic)
}
func (m *Star_Created) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star_Created.Merge(m, src)
}
func (m *Star_Created) XXX_Size() int {
	return xxx_messageInfo_Star_Created.Size(m)
}
func (m *Star_Created) XXX_DiscardUnknown() {
	xxx_messageInfo_Star_Created.DiscardUnknown(m)
}

var xxx_messageInfo_Star_Created proto.InternalMessageInfo

func (m *Star_Created) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Star_Created) GetLotId() int64 {
	if m != nil {
		return m.LotId
	}
	return 0
}

func (m *Star_Created) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Star_Created) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Star_Deleted struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star_Deleted) Reset()         { *m = Star_Deleted{} }
func (m *Star_Deleted) String() string { return proto.CompactTextString(m) }
func (*Star_Deleted) ProtoMessage()    {}
func (*Star_Deleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1, 1}
}

func (m *Star_Deleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star_Deleted.Unmarshal(m, b)
}
func (m *Star_Deleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star_Deleted.Marshal(b, m, deterministic)
}
func (m *Star_Deleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star_Deleted.Merge(m, src)
}
func (m *Star_Deleted) XXX_Size() int {
	return xxx_messageInfo_Star_Deleted.Size(m)
}
func (m *Star_Deleted) XXX_DiscardUnknown() {
	xxx_messageInfo_Star_Deleted.DiscardUnknown(m)
}

var xxx_messageInfo_Star_Deleted proto.InternalMessageInfo

func (m *Star_Deleted) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Star_Expiring struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LotId                int64    `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	LotName              string   `protobuf:"bytes,4,opt,name=lot_name,json=lotName,proto3" json:"lot_name,omitempty"`
	ExpiringAt           int64    `protobuf:"varint,5,opt,name=expiring_at,json=expiringAt,proto3" json:"expiring_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star_Expiring) Reset()         { *m = Star_Expiring{} }
func (m *Star_Expiring) String() string { return proto.CompactTextString(m) }
func (*Star_Expiring) ProtoMessage()    {}
func (*Star_Expiring) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1, 2}
}

func (m *Star_Expiring) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star_Expiring.Unmarshal(m, b)
}
func (m *Star_Expiring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star_Expiring.Marshal(b, m, deterministic)
}
func (m *Star_Expiring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star_Expiring.Merge(m, src)
}
func (m *Star_Expiring) XXX_Size() int {
	return xxx_messageInfo_Star_Expiring.Size(m)
}
func (m *Star_Expiring) XXX_DiscardUnknown() {
	xxx_messageInfo_Star_Expiring.DiscardUnknown(m)
}

var xxx_messageInfo_Star_Expiring proto.InternalMessageInfo

func (m *Star_Expiring) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Star_Expiring) GetLotId() int64 {
	if m != nil {
		return m.LotId
	}
	return 0
}

func (m *Star_Expiring) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Star_Expiring) GetLotName() string {
	if m != nil {
		return m.LotName
	}
	return ""
}

func (m *Star_Expiring) GetExpiringAt() int64 {
	if m != nil {
		return m.ExpiringAt
	}
	return 0
}

type Star_Expired struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LotId                int64    `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star_Expired) Reset()         { *m = Star_Expired{} }
func (m *Star_Expired) String() string { return proto.CompactTextString(m) }
func (*Star_Expired) ProtoMessage()    {}
func (*Star_Expired) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ac58a8ca3678fc, []int{1, 3}
}

func (m *Star_Expired) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star_Expired.Unmarshal(m, b)
}
func (m *Star_Expired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star_Expired.Marshal(b, m, deterministic)
}
func (m *Star_Expired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star_Expired.Merge(m, src)
}
func (m *Star_Expired) XXX_Size() int {
	return xxx_messageInfo_Star_Expired.Size(m)
}
func (m *Star_Expired) XXX_DiscardUnknown() {
	xxx_messageInfo_Star_Expired.DiscardUnknown(m)
}

var xxx_messageInfo_Star_Expired proto.InternalMessageInfo

func (m *Star_Expired) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Star_Expired) GetLotId() int64 {
	if m != nil {
		return m.LotId
	}
	return 0
}

func (m *Star_Expired) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.Star_EventType", Star_EventType_name, Star_EventType_value)
	proto.RegisterType((*StarEvent)(nil), "events.StarEvent")
	proto.RegisterType((*Star)(nil), "events.Star")
	proto.RegisterType((*Star_Created)(nil), "events.Star.Created")
	proto.RegisterType((*Star_Deleted)(nil), "events.Star.Deleted")
	proto.RegisterType((*Star_Expiring)(nil), "events.Star.Expiring")
	proto.RegisterType((*Star_Expired)(nil), "events.Star.Expired")
}

func init() { proto.RegisterFile("stars.proto", fileDescriptor_d8ac58a8ca3678fc) }

var fileDescriptor_d8ac58a8ca3678fc = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x8b, 0x9b, 0x40,
	0x18, 0xc6, 0xeb, 0x9f, 0xc4, 0xf8, 0x06, 0x16, 0x99, 0x76, 0xb7, 0xae, 0x50, 0xba, 0xa4, 0x14,
	0x96, 0x1e, 0x94, 0x6e, 0x3f, 0x81, 0x5d, 0x87, 0x12, 0x5a, 0xdc, 0x65, 0x6a, 0xff, 0x5e, 0x64,
	0xa2, 0xd3, 0x20, 0x31, 0x8e, 0x98, 0x49, 0x69, 0xee, 0x3d, 0xf6, 0x03, 0x97, 0x9e, 0xca, 0x8c,
	0x1a, 0x36, 0x92, 0x63, 0x6e, 0xfa, 0x3c, 0xbf, 0xf1, 0x7d, 0x9e, 0x97, 0x11, 0xa6, 0x1b, 0x41,
	0x9b, 0x8d, 0x5f, 0x37, 0x5c, 0x70, 0x34, 0x66, 0x3f, 0x59, 0x25, 0x36, 0xb3, 0x7f, 0x1a, 0xd8,
	0x1f, 0x05, 0x6d, 0xb0, 0x7c, 0x45, 0xaf, 0xc0, 0x14, 0xbb, 0x9a, 0xb9, 0xda, 0x95, 0x76, 0x7d,
	0x76, 0x73, 0xe1, 0xb7, 0x90, 0x2f, 0x01, 0x5f, 0x11, 0xc9, 0xae, 0x66, 0x44, 0x31, 0xc8, 0x07,
	0x2b, 0x6b, 0x18, 0x15, 0x2c, 0x77, 0xf5, 0x2b, 0xed, 0x7a, 0x7a, 0xf3, 0xe4, 0x00, 0xbf, 0x6d,
	0x3d, 0xd2, 0x43, 0x92, 0xcf, 0x59, 0xc9, 0x24, 0x6f, 0x1c, 0xe1, 0xa3, 0xd6, 0x23, 0x3d, 0x84,
	0x5e, 0xc3, 0x84, 0xfd, 0xaa, 0x8b, 0xa6, 0xa8, 0x96, 0xae, 0xa9, 0x0e, 0x9c, 0x1f, 0xe6, 0xe9,
	0x4c, 0xb2, 0xc7, 0xe4, 0x08, 0xf5, 0xcc, 0x72, 0x77, 0x74, 0x64, 0x04, 0x6e, 0x3d, 0xd2, 0x43,
	0xb3, 0xbf, 0x06, 0x98, 0xd2, 0xf1, 0x2a, 0xb0, 0xba, 0xbc, 0xe8, 0x0c, 0xf4, 0x22, 0x57, 0x0b,
	0x30, 0x88, 0x5e, 0xe4, 0xe8, 0x1c, 0xc6, 0x25, 0x17, 0x69, 0xd1, 0xb6, 0x34, 0xc8, 0xa8, 0xe4,
	0x62, 0x9e, 0xa3, 0x67, 0x00, 0x34, 0xcb, 0xf8, 0xb6, 0x52, 0x96, 0x2c, 0x64, 0x13, 0xbb, 0x53,
	0x5a, 0xbb, 0xeb, 0x9d, 0x52, 0xa1, 0xe2, 0x1b, 0xc4, 0xee, 0x94, 0x50, 0x78, 0x97, 0x60, 0x75,
	0x7d, 0x87, 0xf3, 0xbc, 0x3f, 0x1a, 0x4c, 0xfa, 0x6a, 0x27, 0x0a, 0x73, 0x09, 0x13, 0x79, 0xaa,
	0xa2, 0x6b, 0xa6, 0xa2, 0xd8, 0xc4, 0x2a, 0xb9, 0x88, 0xe9, 0x9a, 0xa1, 0xe7, 0x30, 0xed, 0xb7,
	0x27, 0x83, 0x8e, 0xd4, 0x57, 0xa1, 0x97, 0x42, 0xe1, 0xdd, 0x81, 0xd5, 0xad, 0xed, 0x34, 0x61,
	0x66, 0xbf, 0x35, 0xb0, 0xf7, 0x57, 0x09, 0x5d, 0x00, 0xc2, 0x9f, 0x71, 0x9c, 0xa4, 0xc9, 0xb7,
	0x7b, 0x9c, 0x7e, 0x8a, 0xdf, 0xc7, 0x77, 0x5f, 0x62, 0xe7, 0xd1, 0x40, 0xbf, 0x25, 0x38, 0x4c,
	0x70, 0xe4, 0x68, 0x03, 0x3d, 0xc2, 0x1f, 0xb0, 0xd4, 0x75, 0xf4, 0x14, 0x1e, 0x3f, 0xd0, 0xf1,
	0xd7, 0xfb, 0x39, 0x99, 0xc7, 0xef, 0x1c, 0x63, 0x70, 0x40, 0x19, 0x38, 0x72, 0xcc, 0xb7, 0x2f,
	0xbf, 0xbf, 0x58, 0x16, 0xa2, 0xa4, 0x0b, 0x3f, 0xe3, 0xeb, 0x80, 0x85, 0xdb, 0x4c, 0x14, 0xbc,
	0x0a, 0xda, 0xeb, 0x12, 0x2c, 0x79, 0xb0, 0xa2, 0x3f, 0x56, 0x74, 0x31, 0x56, 0x7f, 0xcb, 0x9b,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xf8, 0x4b, 0x33, 0x3c, 0x03, 0x00, 0x00,
}
