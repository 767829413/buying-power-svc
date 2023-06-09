// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dialogs.proto

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

type DialogEvent_Type int32

const (
	DialogEvent_EVENT_TYPE_UNKNOWN     DialogEvent_Type = 0
	DialogEvent_EVENT_TYPE_CREATED     DialogEvent_Type = 1
	DialogEvent_EVENT_TYPE_NEW_MESSAGE DialogEvent_Type = 2
)

var DialogEvent_Type_name = map[int32]string{
	0: "EVENT_TYPE_UNKNOWN",
	1: "EVENT_TYPE_CREATED",
	2: "EVENT_TYPE_NEW_MESSAGE",
}

var DialogEvent_Type_value = map[string]int32{
	"EVENT_TYPE_UNKNOWN":     0,
	"EVENT_TYPE_CREATED":     1,
	"EVENT_TYPE_NEW_MESSAGE": 2,
}

func (x DialogEvent_Type) String() string {
	return proto.EnumName(DialogEvent_Type_name, int32(x))
}

func (DialogEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{0, 0}
}

type DialogEvent struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Platform             string                  `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Type                 DialogEvent_Type        `protobuf:"varint,3,opt,name=type,proto3,enum=events.DialogEvent_Type" json:"type,omitempty"`
	Created              *DialogEvent_Created    `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	NewMessage           *DialogEvent_NewMessage `protobuf:"bytes,5,opt,name=new_message,json=newMessage,proto3" json:"new_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DialogEvent) Reset()         { *m = DialogEvent{} }
func (m *DialogEvent) String() string { return proto.CompactTextString(m) }
func (*DialogEvent) ProtoMessage()    {}
func (*DialogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{0}
}

func (m *DialogEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialogEvent.Unmarshal(m, b)
}
func (m *DialogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialogEvent.Marshal(b, m, deterministic)
}
func (m *DialogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialogEvent.Merge(m, src)
}
func (m *DialogEvent) XXX_Size() int {
	return xxx_messageInfo_DialogEvent.Size(m)
}
func (m *DialogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DialogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DialogEvent proto.InternalMessageInfo

func (m *DialogEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DialogEvent) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DialogEvent) GetType() DialogEvent_Type {
	if m != nil {
		return m.Type
	}
	return DialogEvent_EVENT_TYPE_UNKNOWN
}

func (m *DialogEvent) GetCreated() *DialogEvent_Created {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *DialogEvent) GetNewMessage() *DialogEvent_NewMessage {
	if m != nil {
		return m.NewMessage
	}
	return nil
}

type DialogEvent_Created struct {
	Dialog               *Dialog  `protobuf:"bytes,1,opt,name=dialog,proto3" json:"dialog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DialogEvent_Created) Reset()         { *m = DialogEvent_Created{} }
func (m *DialogEvent_Created) String() string { return proto.CompactTextString(m) }
func (*DialogEvent_Created) ProtoMessage()    {}
func (*DialogEvent_Created) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{0, 0}
}

func (m *DialogEvent_Created) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialogEvent_Created.Unmarshal(m, b)
}
func (m *DialogEvent_Created) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialogEvent_Created.Marshal(b, m, deterministic)
}
func (m *DialogEvent_Created) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialogEvent_Created.Merge(m, src)
}
func (m *DialogEvent_Created) XXX_Size() int {
	return xxx_messageInfo_DialogEvent_Created.Size(m)
}
func (m *DialogEvent_Created) XXX_DiscardUnknown() {
	xxx_messageInfo_DialogEvent_Created.DiscardUnknown(m)
}

var xxx_messageInfo_DialogEvent_Created proto.InternalMessageInfo

func (m *DialogEvent_Created) GetDialog() *Dialog {
	if m != nil {
		return m.Dialog
	}
	return nil
}

type DialogEvent_NewMessage struct {
	IsFromUser           bool     `protobuf:"varint,1,opt,name=is_from_user,json=isFromUser,proto3" json:"is_from_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DialogEvent_NewMessage) Reset()         { *m = DialogEvent_NewMessage{} }
func (m *DialogEvent_NewMessage) String() string { return proto.CompactTextString(m) }
func (*DialogEvent_NewMessage) ProtoMessage()    {}
func (*DialogEvent_NewMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{0, 1}
}

func (m *DialogEvent_NewMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialogEvent_NewMessage.Unmarshal(m, b)
}
func (m *DialogEvent_NewMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialogEvent_NewMessage.Marshal(b, m, deterministic)
}
func (m *DialogEvent_NewMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialogEvent_NewMessage.Merge(m, src)
}
func (m *DialogEvent_NewMessage) XXX_Size() int {
	return xxx_messageInfo_DialogEvent_NewMessage.Size(m)
}
func (m *DialogEvent_NewMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DialogEvent_NewMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DialogEvent_NewMessage proto.InternalMessageInfo

func (m *DialogEvent_NewMessage) GetIsFromUser() bool {
	if m != nil {
		return m.IsFromUser
	}
	return false
}

type Dialog struct {
	Initiator            *Dialog_Initiator `protobuf:"bytes,3,opt,name=initiator,proto3" json:"initiator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Dialog) Reset()         { *m = Dialog{} }
func (m *Dialog) String() string { return proto.CompactTextString(m) }
func (*Dialog) ProtoMessage()    {}
func (*Dialog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{1}
}

func (m *Dialog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dialog.Unmarshal(m, b)
}
func (m *Dialog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dialog.Marshal(b, m, deterministic)
}
func (m *Dialog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dialog.Merge(m, src)
}
func (m *Dialog) XXX_Size() int {
	return xxx_messageInfo_Dialog.Size(m)
}
func (m *Dialog) XXX_DiscardUnknown() {
	xxx_messageInfo_Dialog.DiscardUnknown(m)
}

var xxx_messageInfo_Dialog proto.InternalMessageInfo

func (m *Dialog) GetInitiator() *Dialog_Initiator {
	if m != nil {
		return m.Initiator
	}
	return nil
}

type Dialog_Initiator struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	AvatarLink           string   `protobuf:"bytes,3,opt,name=avatar_link,json=avatarLink,proto3" json:"avatar_link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dialog_Initiator) Reset()         { *m = Dialog_Initiator{} }
func (m *Dialog_Initiator) String() string { return proto.CompactTextString(m) }
func (*Dialog_Initiator) ProtoMessage()    {}
func (*Dialog_Initiator) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbd424ed875e1b0, []int{1, 0}
}

func (m *Dialog_Initiator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dialog_Initiator.Unmarshal(m, b)
}
func (m *Dialog_Initiator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dialog_Initiator.Marshal(b, m, deterministic)
}
func (m *Dialog_Initiator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dialog_Initiator.Merge(m, src)
}
func (m *Dialog_Initiator) XXX_Size() int {
	return xxx_messageInfo_Dialog_Initiator.Size(m)
}
func (m *Dialog_Initiator) XXX_DiscardUnknown() {
	xxx_messageInfo_Dialog_Initiator.DiscardUnknown(m)
}

var xxx_messageInfo_Dialog_Initiator proto.InternalMessageInfo

func (m *Dialog_Initiator) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Dialog_Initiator) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Dialog_Initiator) GetAvatarLink() string {
	if m != nil {
		return m.AvatarLink
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.DialogEvent_Type", DialogEvent_Type_name, DialogEvent_Type_value)
	proto.RegisterType((*DialogEvent)(nil), "events.DialogEvent")
	proto.RegisterType((*DialogEvent_Created)(nil), "events.DialogEvent.Created")
	proto.RegisterType((*DialogEvent_NewMessage)(nil), "events.DialogEvent.NewMessage")
	proto.RegisterType((*Dialog)(nil), "events.Dialog")
	proto.RegisterType((*Dialog_Initiator)(nil), "events.Dialog.Initiator")
}

func init() { proto.RegisterFile("dialogs.proto", fileDescriptor_ccbd424ed875e1b0) }

var fileDescriptor_ccbd424ed875e1b0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x57, 0xba, 0xe5, 0x04, 0xaa, 0xca, 0x17, 0x53, 0x94, 0x09, 0xa8, 0x8a, 0x40,
	0xbd, 0x40, 0x89, 0x28, 0x82, 0x5b, 0x54, 0x36, 0x83, 0x10, 0x2c, 0x20, 0x2f, 0x63, 0x82, 0x1b,
	0xcb, 0x6b, 0x9c, 0x60, 0x25, 0xb1, 0x23, 0xdb, 0xdb, 0xb4, 0x17, 0xe1, 0x21, 0x78, 0x4a, 0x54,
	0x27, 0xed, 0xc6, 0xb4, 0xbb, 0xe4, 0xf7, 0xff, 0xd0, 0xf1, 0xb1, 0xe1, 0x71, 0x2e, 0x58, 0xad,
	0x4a, 0x13, 0xb7, 0x5a, 0x59, 0x85, 0x46, 0xfc, 0x92, 0x4b, 0x6b, 0x66, 0x7f, 0x76, 0x20, 0x38,
	0x72, 0x0a, 0x5e, 0x03, 0x34, 0x86, 0x81, 0xc8, 0x43, 0x6f, 0xea, 0xcd, 0x7d, 0x32, 0x10, 0x39,
	0x8a, 0x60, 0xaf, 0xad, 0x99, 0x2d, 0x94, 0x6e, 0xc2, 0x81, 0xa3, 0xdb, 0x7f, 0xf4, 0x0a, 0x86,
	0xf6, 0xba, 0xe5, 0xe1, 0xce, 0xd4, 0x9b, 0x8f, 0x17, 0x61, 0xdc, 0x55, 0xc6, 0xb7, 0xea, 0xe2,
	0xec, 0xba, 0xe5, 0xc4, 0xb9, 0xd0, 0x5b, 0xd8, 0x5d, 0x69, 0xce, 0x2c, 0xcf, 0xc3, 0xe1, 0xd4,
	0x9b, 0x07, 0x8b, 0x83, 0xfb, 0x02, 0x87, 0x9d, 0x85, 0x6c, 0xbc, 0xe8, 0x3d, 0x04, 0x92, 0x5f,
	0xd1, 0x86, 0x1b, 0xc3, 0x4a, 0x1e, 0x3e, 0x74, 0xd1, 0xa7, 0xf7, 0x45, 0x53, 0x7e, 0x75, 0xdc,
	0xb9, 0x08, 0xc8, 0xed, 0x77, 0xf4, 0x1a, 0x76, 0xfb, 0x52, 0xf4, 0x12, 0x46, 0xdd, 0x16, 0xdc,
	0x01, 0x83, 0xc5, 0xf8, 0xff, 0x1a, 0xd2, 0xab, 0x51, 0x0c, 0x70, 0x53, 0x86, 0xa6, 0xf0, 0x48,
	0x18, 0x5a, 0x68, 0xd5, 0xd0, 0x0b, 0xc3, 0xb5, 0xcb, 0xee, 0x11, 0x10, 0xe6, 0xa3, 0x56, 0xcd,
	0xa9, 0xe1, 0x7a, 0x46, 0x60, 0xb8, 0x3e, 0x28, 0xda, 0x07, 0x84, 0x7f, 0xe0, 0x34, 0xa3, 0xd9,
	0xcf, 0xef, 0x98, 0x9e, 0xa6, 0x5f, 0xd2, 0x6f, 0x67, 0xe9, 0xe4, 0xc1, 0x1d, 0x7e, 0x48, 0xf0,
	0x32, 0xc3, 0x47, 0x13, 0x0f, 0x45, 0xb0, 0x7f, 0x8b, 0xa7, 0xf8, 0x8c, 0x1e, 0xe3, 0x93, 0x93,
	0xe5, 0x27, 0x3c, 0x19, 0xcc, 0xfe, 0x7a, 0x30, 0xea, 0xc6, 0x42, 0xef, 0xc0, 0x17, 0x52, 0x58,
	0xc1, 0xac, 0xd2, 0x6e, 0xd9, 0xc1, 0xdd, 0x65, 0xc7, 0x9f, 0x37, 0x3a, 0xb9, 0xb1, 0x46, 0xbf,
	0xc1, 0xdf, 0x72, 0xf4, 0x04, 0xa0, 0x10, 0xda, 0x58, 0x2a, 0x59, 0xc3, 0xfb, 0x0b, 0xf6, 0x1d,
	0x49, 0x59, 0xc3, 0xd1, 0x01, 0xf8, 0x35, 0xdb, 0xa8, 0xfd, 0x45, 0xaf, 0x81, 0x13, 0x9f, 0x41,
	0xc0, 0x2e, 0x99, 0x65, 0x9a, 0xd6, 0x42, 0x56, 0x6e, 0x04, 0x9f, 0x40, 0x87, 0xbe, 0x0a, 0x59,
	0x7d, 0x78, 0xf1, 0xeb, 0x79, 0x29, 0x6c, 0xcd, 0xce, 0xe3, 0x95, 0x6a, 0x12, 0xbe, 0xbc, 0x58,
	0x59, 0xa1, 0x64, 0xd2, 0xcd, 0x98, 0x94, 0x2a, 0xa9, 0x58, 0x51, 0xb1, 0xf3, 0x91, 0x7b, 0x7b,
	0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x03, 0xe4, 0x4f, 0x0d, 0x8c, 0x02, 0x00, 0x00,
}
