// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messenger_commands.proto

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

type MessengerCommand_Type int32

const (
	MessengerCommand_TYPE_UNKNOWN         MessengerCommand_Type = 0
	MessengerCommand_TYPE_BROKERS_UPDATED MessengerCommand_Type = 1
)

var MessengerCommand_Type_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "TYPE_BROKERS_UPDATED",
}

var MessengerCommand_Type_value = map[string]int32{
	"TYPE_UNKNOWN":         0,
	"TYPE_BROKERS_UPDATED": 1,
}

func (x MessengerCommand_Type) String() string {
	return proto.EnumName(MessengerCommand_Type_name, int32(x))
}

func (MessengerCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0eed1cdedd280fa5, []int{0, 0}
}

// DEPRECATED:
type MessengerCommand struct {
	Type                 MessengerCommand_Type            `protobuf:"varint,1,opt,name=type,proto3,enum=events.MessengerCommand_Type" json:"type,omitempty"`
	BrokersUpdated       *MessengerCommand_BrokersUpdated `protobuf:"bytes,2,opt,name=brokers_updated,json=brokersUpdated,proto3" json:"brokers_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MessengerCommand) Reset()         { *m = MessengerCommand{} }
func (m *MessengerCommand) String() string { return proto.CompactTextString(m) }
func (*MessengerCommand) ProtoMessage()    {}
func (*MessengerCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eed1cdedd280fa5, []int{0}
}

func (m *MessengerCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessengerCommand.Unmarshal(m, b)
}
func (m *MessengerCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessengerCommand.Marshal(b, m, deterministic)
}
func (m *MessengerCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessengerCommand.Merge(m, src)
}
func (m *MessengerCommand) XXX_Size() int {
	return xxx_messageInfo_MessengerCommand.Size(m)
}
func (m *MessengerCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_MessengerCommand.DiscardUnknown(m)
}

var xxx_messageInfo_MessengerCommand proto.InternalMessageInfo

func (m *MessengerCommand) GetType() MessengerCommand_Type {
	if m != nil {
		return m.Type
	}
	return MessengerCommand_TYPE_UNKNOWN
}

func (m *MessengerCommand) GetBrokersUpdated() *MessengerCommand_BrokersUpdated {
	if m != nil {
		return m.BrokersUpdated
	}
	return nil
}

type MessengerCommand_BrokersUpdated struct {
	DialogId             string   `protobuf:"bytes,1,opt,name=dialog_id,json=dialogId,proto3" json:"dialog_id,omitempty"`
	Brokers              []string `protobuf:"bytes,2,rep,name=brokers,proto3" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessengerCommand_BrokersUpdated) Reset()         { *m = MessengerCommand_BrokersUpdated{} }
func (m *MessengerCommand_BrokersUpdated) String() string { return proto.CompactTextString(m) }
func (*MessengerCommand_BrokersUpdated) ProtoMessage()    {}
func (*MessengerCommand_BrokersUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eed1cdedd280fa5, []int{0, 0}
}

func (m *MessengerCommand_BrokersUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessengerCommand_BrokersUpdated.Unmarshal(m, b)
}
func (m *MessengerCommand_BrokersUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessengerCommand_BrokersUpdated.Marshal(b, m, deterministic)
}
func (m *MessengerCommand_BrokersUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessengerCommand_BrokersUpdated.Merge(m, src)
}
func (m *MessengerCommand_BrokersUpdated) XXX_Size() int {
	return xxx_messageInfo_MessengerCommand_BrokersUpdated.Size(m)
}
func (m *MessengerCommand_BrokersUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_MessengerCommand_BrokersUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_MessengerCommand_BrokersUpdated proto.InternalMessageInfo

func (m *MessengerCommand_BrokersUpdated) GetDialogId() string {
	if m != nil {
		return m.DialogId
	}
	return ""
}

func (m *MessengerCommand_BrokersUpdated) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.MessengerCommand_Type", MessengerCommand_Type_name, MessengerCommand_Type_value)
	proto.RegisterType((*MessengerCommand)(nil), "events.MessengerCommand")
	proto.RegisterType((*MessengerCommand_BrokersUpdated)(nil), "events.MessengerCommand.BrokersUpdated")
}

func init() { proto.RegisterFile("messenger_commands.proto", fileDescriptor_0eed1cdedd280fa5) }

var fileDescriptor_0eed1cdedd280fa5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x61, 0x4b, 0x84, 0x30,
	0x1c, 0x06, 0xf0, 0xb4, 0xe3, 0xca, 0x15, 0x26, 0xa3, 0x17, 0xa3, 0x08, 0xe4, 0x22, 0xf2, 0x95,
	0x92, 0x7d, 0x82, 0xb3, 0x93, 0x88, 0x23, 0x4f, 0x96, 0x12, 0xf5, 0x46, 0xa6, 0xfb, 0x27, 0xe2,
	0xe9, 0x44, 0x77, 0xc1, 0x7d, 0x8f, 0x3e, 0x70, 0xe4, 0xba, 0x17, 0x17, 0xf4, 0xf2, 0x79, 0xf8,
	0x6d, 0x3c, 0x1b, 0x22, 0x0d, 0x0c, 0x03, 0xb4, 0x25, 0xf4, 0x59, 0x21, 0x9a, 0x86, 0xb5, 0x7c,
	0x70, 0xbb, 0x5e, 0x48, 0x81, 0xa7, 0xf0, 0x09, 0xad, 0x1c, 0x66, 0x5f, 0x3a, 0xb2, 0x9e, 0x77,
	0xe8, 0x41, 0x19, 0x7c, 0x87, 0x26, 0x72, 0xdb, 0x01, 0xd1, 0x6c, 0xcd, 0x31, 0xfd, 0x2b, 0x57,
	0x59, 0xf7, 0xaf, 0x73, 0x93, 0x6d, 0x07, 0x74, 0xa4, 0x38, 0x46, 0x67, 0x79, 0x2f, 0x6a, 0xe8,
	0x87, 0x6c, 0xd3, 0x71, 0x26, 0x81, 0x13, 0xdd, 0xd6, 0x9c, 0x13, 0xff, 0xf6, 0xdf, 0xd3, 0x81,
	0xf2, 0xa9, 0xe2, 0xd4, 0xcc, 0xf7, 0xf2, 0xc5, 0x23, 0x32, 0xf7, 0x05, 0xbe, 0x44, 0x06, 0xaf,
	0xd8, 0x5a, 0x94, 0x59, 0xc5, 0xc7, 0x6d, 0x06, 0x3d, 0x56, 0xc5, 0x13, 0xc7, 0x04, 0x1d, 0xfd,
	0x5e, 0x40, 0x74, 0xfb, 0xd0, 0x31, 0xe8, 0x2e, 0xce, 0x7c, 0x34, 0xf9, 0x19, 0x8a, 0x2d, 0x74,
	0x9a, 0xbc, 0xc5, 0x61, 0x96, 0x46, 0xcb, 0x68, 0xf5, 0x1a, 0x59, 0x07, 0x98, 0xa0, 0xf3, 0xb1,
	0x09, 0xe8, 0x6a, 0x19, 0xd2, 0x97, 0x2c, 0x8d, 0x17, 0xf3, 0x24, 0x5c, 0x58, 0x5a, 0x70, 0xf3,
	0x7e, 0x5d, 0x56, 0x72, 0xcd, 0x72, 0xb7, 0x10, 0x8d, 0x07, 0xf3, 0x4d, 0x21, 0x2b, 0xd1, 0x7a,
	0xea, 0x29, 0x5e, 0x29, 0xbc, 0x9a, 0x7d, 0xd4, 0x2c, 0x9f, 0x8e, 0x9f, 0x79, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0xf4, 0x8b, 0x1e, 0x68, 0x01, 0x00, 0x00,
}
