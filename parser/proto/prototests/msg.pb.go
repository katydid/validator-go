// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package prototests

import (
	bytes "bytes"
	compress_gzip "compress/gzip"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	io_ioutil "io/ioutil"
	math "math"
	reflect "reflect"
	strings "strings"
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

// BigMsg contains a field and a message field.
type BigMsg struct {
	Field            *int64    `protobuf:"varint,1,opt,name=Field" json:"Field,omitempty"`
	Msg              *SmallMsg `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BigMsg) Reset()         { *m = BigMsg{} }
func (m *BigMsg) String() string { return proto.CompactTextString(m) }
func (*BigMsg) ProtoMessage()    {}
func (*BigMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}
func (m *BigMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigMsg.Unmarshal(m, b)
}
func (m *BigMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigMsg.Marshal(b, m, deterministic)
}
func (m *BigMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigMsg.Merge(m, src)
}
func (m *BigMsg) XXX_Size() int {
	return xxx_messageInfo_BigMsg.Size(m)
}
func (m *BigMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BigMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BigMsg proto.InternalMessageInfo

func (m *BigMsg) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BigMsg) GetMsg() *SmallMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SmallMsg only contains some native fields.
type SmallMsg struct {
	ScarBusStop      *string  `protobuf:"bytes,1,opt,name=ScarBusStop" json:"ScarBusStop,omitempty"`
	FlightParachute  []uint32 `protobuf:"fixed32,12,rep,name=FlightParachute" json:"FlightParachute,omitempty"`
	MapShark         *string  `protobuf:"bytes,18,opt,name=MapShark" json:"MapShark,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMsg) Reset()         { *m = SmallMsg{} }
func (m *SmallMsg) String() string { return proto.CompactTextString(m) }
func (*SmallMsg) ProtoMessage()    {}
func (*SmallMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}
func (m *SmallMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmallMsg.Unmarshal(m, b)
}
func (m *SmallMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmallMsg.Marshal(b, m, deterministic)
}
func (m *SmallMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmallMsg.Merge(m, src)
}
func (m *SmallMsg) XXX_Size() int {
	return xxx_messageInfo_SmallMsg.Size(m)
}
func (m *SmallMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SmallMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SmallMsg proto.InternalMessageInfo

func (m *SmallMsg) GetScarBusStop() string {
	if m != nil && m.ScarBusStop != nil {
		return *m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func (m *SmallMsg) GetMapShark() string {
	if m != nil && m.MapShark != nil {
		return *m.MapShark
	}
	return ""
}

// Packed contains some repeated packed fields.
type Packed struct {
	Ints             []int64   `protobuf:"varint,4,rep,packed,name=Ints" json:"Ints,omitempty"`
	Floats           []float64 `protobuf:"fixed64,5,rep,packed,name=Floats" json:"Floats,omitempty"`
	Uints            []uint32  `protobuf:"varint,6,rep,packed,name=Uints" json:"Uints,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Packed) Reset()         { *m = Packed{} }
func (m *Packed) String() string { return proto.CompactTextString(m) }
func (*Packed) ProtoMessage()    {}
func (*Packed) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}
func (m *Packed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packed.Unmarshal(m, b)
}
func (m *Packed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packed.Marshal(b, m, deterministic)
}
func (m *Packed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packed.Merge(m, src)
}
func (m *Packed) XXX_Size() int {
	return xxx_messageInfo_Packed.Size(m)
}
func (m *Packed) XXX_DiscardUnknown() {
	xxx_messageInfo_Packed.DiscardUnknown(m)
}

var xxx_messageInfo_Packed proto.InternalMessageInfo

func (m *Packed) GetInts() []int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *Packed) GetFloats() []float64 {
	if m != nil {
		return m.Floats
	}
	return nil
}

func (m *Packed) GetUints() []uint32 {
	if m != nil {
		return m.Uints
	}
	return nil
}

func init() {
	proto.RegisterType((*BigMsg)(nil), "prototests.BigMsg")
	proto.RegisterType((*SmallMsg)(nil), "prototests.SmallMsg")
	proto.RegisterType((*Packed)(nil), "prototests.Packed")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8d, 0x3f, 0x4e, 0xc3, 0x30,
	0x1c, 0x85, 0x6b, 0xdc, 0x86, 0xf6, 0x57, 0x10, 0xc8, 0x8a, 0x90, 0x95, 0x21, 0xb2, 0x32, 0x20,
	0x2f, 0xa4, 0x12, 0x47, 0xc8, 0x10, 0x89, 0x21, 0x52, 0x95, 0x08, 0x76, 0x37, 0x0d, 0x4e, 0xd4,
	0x04, 0x47, 0xb6, 0x73, 0x2f, 0x8e, 0xc2, 0xc8, 0x5a, 0x7a, 0x01, 0x46, 0x46, 0x54, 0x87, 0x7f,
	0x62, 0xb2, 0xbf, 0xef, 0xf9, 0x3d, 0xc3, 0xa2, 0x33, 0x32, 0xee, 0xb5, 0xb2, 0x8a, 0x80, 0x3b,
	0x6c, 0x65, 0xac, 0x09, 0x6e, 0x64, 0x63, 0xeb, 0x61, 0x13, 0x97, 0xaa, 0x5b, 0x49, 0x25, 0xd5,
	0xca, 0x65, 0x9b, 0xe1, 0xd1, 0x91, 0x03, 0x77, 0x1b, 0xab, 0x51, 0x0a, 0x5e, 0xd2, 0xc8, 0xcc,
	0x48, 0xe2, 0xc3, 0x2c, 0x6d, 0xaa, 0x76, 0x4b, 0x11, 0x43, 0x1c, 0xe7, 0x23, 0x90, 0x6b, 0xc0,
	0x99, 0x91, 0x14, 0x33, 0xc4, 0x97, 0xb7, 0x7e, 0xfc, 0xfb, 0x51, 0x5c, 0x74, 0xa2, 0x6d, 0x33,
	0x23, 0xf3, 0xe3, 0x83, 0x48, 0xc3, 0xfc, 0x5b, 0x10, 0x06, 0xcb, 0xa2, 0x14, 0x3a, 0x19, 0x4c,
	0x61, 0x55, 0xef, 0xf6, 0x16, 0xf9, 0x5f, 0x45, 0x38, 0x5c, 0xa4, 0x6d, 0x23, 0x6b, 0xbb, 0x16,
	0x5a, 0x94, 0xf5, 0x60, 0x2b, 0x7a, 0xc6, 0x30, 0x3f, 0xcd, 0xff, 0x6b, 0x12, 0xc0, 0x3c, 0x13,
	0x7d, 0x51, 0x0b, 0xbd, 0xa3, 0xc4, 0x0d, 0xfd, 0x70, 0xf4, 0x00, 0xde, 0x5a, 0x94, 0xbb, 0x6a,
	0x4b, 0xae, 0x60, 0x7a, 0xf7, 0x64, 0x0d, 0x9d, 0x32, 0xcc, 0x71, 0x72, 0x72, 0x89, 0x72, 0xc7,
	0x24, 0x00, 0x2f, 0x6d, 0x95, 0xb0, 0x86, 0xce, 0x18, 0xe6, 0xc8, 0x25, 0x5f, 0x86, 0x50, 0x98,
	0xdd, 0x37, 0xc7, 0x92, 0xc7, 0x30, 0x3f, 0x77, 0xd1, 0x28, 0x12, 0xff, 0x7d, 0x1f, 0xa2, 0x8f,
	0x7d, 0x88, 0x9e, 0xdf, 0x42, 0xf4, 0x72, 0x08, 0x27, 0xaf, 0x87, 0x70, 0xf2, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x8b, 0xe2, 0xad, 0x70, 0x70, 0x01, 0x00, 0x00,
}

func (this *BigMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *SmallMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *Packed) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func MsgDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3979 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x4b, 0x70, 0x23, 0xd7,
		0x75, 0x15, 0xbe, 0x04, 0x2e, 0x40, 0xb0, 0xf9, 0x48, 0x71, 0x30, 0x94, 0x35, 0xc3, 0x81, 0x7e,
		0x94, 0x64, 0x71, 0x5c, 0xa3, 0x99, 0x91, 0x84, 0x89, 0xad, 0x80, 0x24, 0x48, 0x43, 0x21, 0x48,
		0xb8, 0x41, 0xea, 0xe7, 0x4a, 0x75, 0x3d, 0x36, 0x1e, 0x81, 0x9e, 0x69, 0x74, 0xb7, 0xbb, 0x1b,
		0x33, 0xe2, 0x54, 0x16, 0x4a, 0x29, 0x9f, 0x72, 0xe5, 0xff, 0xa9, 0x8a, 0xad, 0xc8, 0x4a, 0xec,
		0x54, 0xac, 0xc4, 0xf9, 0xd9, 0xf9, 0x38, 0xb1, 0xb3, 0xc9, 0xc6, 0x89, 0x57, 0xa9, 0x78, 0x97,
		0x45, 0x16, 0x71, 0xa4, 0xaa, 0xfc, 0x94, 0x44, 0x49, 0x66, 0xe1, 0x2a, 0x6d, 0x5c, 0xef, 0xd7,
		0x68, 0x34, 0xc0, 0x69, 0xd0, 0x55, 0x92, 0x57, 0x64, 0xdf, 0x77, 0xcf, 0xe9, 0xfb, 0xee, 0xbb,
		0xef, 0xde, 0xfb, 0x5e, 0x03, 0xbe, 0x5d, 0x85, 0x95, 0xae, 0x6d, 0x77, 0x4d, 0x72, 0xd1, 0x71,
		0x6d, 0xdf, 0x3e, 0x1c, 0x1c, 0x5d, 0xec, 0x10, 0x4f, 0x77, 0x0d, 0xc7, 0xb7, 0xdd, 0x35, 0x26,
		0x43, 0x73, 0x5c, 0x63, 0x4d, 0x6a, 0x54, 0x9a, 0x30, 0xbf, 0x65, 0x98, 0x64, 0x33, 0x50, 0x6c,
		0x13, 0x1f, 0x3d, 0x0d, 0xe9, 0x23, 0xc3, 0x24, 0xe5, 0xc4, 0x4a, 0x6a, 0xb5, 0x70, 0xe9, 0xc1,
		0xb5, 0x08, 0x68, 0x6d, 0x14, 0xd1, 0xa2, 0x62, 0x95, 0x21, 0x2a, 0xef, 0xa4, 0x61, 0x61, 0xc2,
		0x28, 0x42, 0x90, 0xb6, 0x70, 0x9f, 0x32, 0x26, 0x56, 0xf3, 0x2a, 0xfb, 0x1f, 0x95, 0x61, 0xc6,
		0xc1, 0xfa, 0x0d, 0xdc, 0x25, 0xe5, 0x24, 0x13, 0xcb, 0x47, 0x74, 0x0e, 0xa0, 0x43, 0x1c, 0x62,
		0x75, 0x88, 0xa5, 0x1f, 0x97, 0x53, 0x2b, 0xa9, 0xd5, 0xbc, 0x1a, 0x92, 0xa0, 0xc7, 0x61, 0xde,
		0x19, 0x1c, 0x9a, 0x86, 0xae, 0x85, 0xd4, 0x60, 0x25, 0xb5, 0x9a, 0x51, 0x15, 0x3e, 0xb0, 0x39,
		0x54, 0x7e, 0x04, 0xe6, 0x6e, 0x11, 0x7c, 0x23, 0xac, 0x5a, 0x60, 0xaa, 0x25, 0x2a, 0x0e, 0x29,
		0x6e, 0x40, 0xb1, 0x4f, 0x3c, 0x0f, 0x77, 0x89, 0xe6, 0x1f, 0x3b, 0xa4, 0x9c, 0x66, 0xb3, 0x5f,
		0x19, 0x9b, 0x7d, 0x74, 0xe6, 0x05, 0x81, 0xda, 0x3f, 0x76, 0x08, 0xaa, 0x41, 0x9e, 0x58, 0x83,
		0x3e, 0x67, 0xc8, 0x9c, 0xe0, 0xbf, 0xba, 0x35, 0xe8, 0x47, 0x59, 0x72, 0x14, 0x26, 0x28, 0x66,
		0x3c, 0xe2, 0xde, 0x34, 0x74, 0x52, 0xce, 0x32, 0x82, 0x47, 0xc6, 0x08, 0xda, 0x7c, 0x3c, 0xca,
		0x21, 0x71, 0x68, 0x03, 0xf2, 0xe4, 0x15, 0x9f, 0x58, 0x9e, 0x61, 0x5b, 0xe5, 0x19, 0x46, 0xf2,
		0xd0, 0x84, 0x55, 0x24, 0x66, 0x27, 0x4a, 0x31, 0xc4, 0xa1, 0xab, 0x30, 0x63, 0x3b, 0xbe, 0x61,
		0x5b, 0x5e, 0x39, 0xb7, 0x92, 0x58, 0x2d, 0x5c, 0xfa, 0xc8, 0xc4, 0x40, 0xd8, 0xe3, 0x3a, 0xaa,
		0x54, 0x46, 0x0d, 0x50, 0x3c, 0x7b, 0xe0, 0xea, 0x44, 0xd3, 0xed, 0x0e, 0xd1, 0x0c, 0xeb, 0xc8,
		0x2e, 0xe7, 0x19, 0xc1, 0xf9, 0xf1, 0x89, 0x30, 0xc5, 0x0d, 0xbb, 0x43, 0x1a, 0xd6, 0x91, 0xad,
		0x96, 0xbc, 0x91, 0x67, 0xb4, 0x04, 0x59, 0xef, 0xd8, 0xf2, 0xf1, 0x2b, 0xe5, 0x22, 0x8b, 0x10,
		0xf1, 0x54, 0xf9, 0x46, 0x16, 0xe6, 0xa6, 0x09, 0xb1, 0x6b, 0x90, 0x39, 0xa2, 0xb3, 0x2c, 0x27,
		0x4f, 0xe3, 0x03, 0x8e, 0x19, 0x75, 0x62, 0xf6, 0x07, 0x74, 0x62, 0x0d, 0x0a, 0x16, 0xf1, 0x7c,
		0xd2, 0xe1, 0x11, 0x91, 0x9a, 0x32, 0xa6, 0x80, 0x83, 0xc6, 0x43, 0x2a, 0xfd, 0x03, 0x85, 0xd4,
		0x8b, 0x30, 0x17, 0x98, 0xa4, 0xb9, 0xd8, 0xea, 0xca, 0xd8, 0xbc, 0x18, 0x67, 0xc9, 0x5a, 0x5d,
		0xe2, 0x54, 0x0a, 0x53, 0x4b, 0x64, 0xe4, 0x19, 0x6d, 0x02, 0xd8, 0x16, 0xb1, 0x8f, 0xb4, 0x0e,
		0xd1, 0xcd, 0x72, 0xee, 0x04, 0x2f, 0xed, 0x51, 0x95, 0x31, 0x2f, 0xd9, 0x5c, 0xaa, 0x9b, 0xe8,
		0x99, 0x61, 0xa8, 0xcd, 0x9c, 0x10, 0x29, 0x4d, 0xbe, 0xc9, 0xc6, 0xa2, 0xed, 0x00, 0x4a, 0x2e,
		0xa1, 0x71, 0x4f, 0x3a, 0x62, 0x66, 0x79, 0x66, 0xc4, 0x5a, 0xec, 0xcc, 0x54, 0x01, 0xe3, 0x13,
		0x9b, 0x75, 0xc3, 0x8f, 0xe8, 0x01, 0x08, 0x04, 0x1a, 0x0b, 0x2b, 0x60, 0x59, 0xa8, 0x28, 0x85,
		0xbb, 0xb8, 0x4f, 0x96, 0x6f, 0x43, 0x69, 0xd4, 0x3d, 0x68, 0x11, 0x32, 0x9e, 0x8f, 0x5d, 0x9f,
		0x45, 0x61, 0x46, 0xe5, 0x0f, 0x48, 0x81, 0x14, 0xb1, 0x3a, 0x2c, 0xcb, 0x65, 0x54, 0xfa, 0x2f,
		0xfa, 0xd1, 0xe1, 0x84, 0x53, 0x6c, 0xc2, 0x0f, 0x8f, 0xaf, 0xe8, 0x08, 0x73, 0x74, 0xde, 0xcb,
		0x4f, 0xc1, 0xec, 0xc8, 0x04, 0xa6, 0x7d, 0x75, 0xe5, 0x27, 0xe0, 0xde, 0x89, 0xd4, 0xe8, 0x45,
		0x58, 0x1c, 0x58, 0x86, 0xe5, 0x13, 0xd7, 0x71, 0x09, 0x8d, 0x58, 0xfe, 0xaa, 0xf2, 0xbf, 0xce,
		0x9c, 0x10, 0x73, 0x07, 0x61, 0x6d, 0xce, 0xa2, 0x2e, 0x0c, 0xc6, 0x85, 0x8f, 0xe5, 0x73, 0xff,
		0x36, 0xa3, 0xbc, 0xfa, 0xea, 0xab, 0xaf, 0x26, 0x2b, 0x9f, 0xcb, 0xc2, 0xe2, 0xa4, 0x3d, 0x33,
		0x71, 0xfb, 0x2e, 0x41, 0xd6, 0x1a, 0xf4, 0x0f, 0x89, 0xcb, 0x9c, 0x94, 0x51, 0xc5, 0x13, 0xaa,
		0x41, 0xc6, 0xc4, 0x87, 0xc4, 0x2c, 0xa7, 0x57, 0x12, 0xab, 0xa5, 0x4b, 0x8f, 0x4f, 0xb5, 0x2b,
		0xd7, 0x76, 0x28, 0x44, 0xe5, 0x48, 0xf4, 0x09, 0x48, 0x8b, 0x14, 0x4d, 0x19, 0x1e, 0x9b, 0x8e,
		0x81, 0xee, 0x25, 0x95, 0xe1, 0xd0, 0x7d, 0x90, 0xa7, 0x7f, 0x79, 0x6c, 0x64, 0x99, 0xcd, 0x39,
		0x2a, 0xa0, 0x71, 0x81, 0x96, 0x21, 0xc7, 0xb6, 0x49, 0x87, 0xc8, 0xd2, 0x16, 0x3c, 0xd3, 0xc0,
		0xea, 0x90, 0x23, 0x3c, 0x30, 0x7d, 0xed, 0x26, 0x36, 0x07, 0x84, 0x05, 0x7c, 0x5e, 0x2d, 0x0a,
		0xe1, 0xf3, 0x54, 0x86, 0xce, 0x43, 0x81, 0xef, 0x2a, 0xc3, 0xea, 0x90, 0x57, 0x58, 0xf6, 0xcc,
		0xa8, 0x7c, 0xa3, 0x35, 0xa8, 0x84, 0xbe, 0xfe, 0xba, 0x67, 0x5b, 0x32, 0x34, 0xd9, 0x2b, 0xa8,
		0x80, 0xbd, 0xfe, 0xa9, 0x68, 0xe2, 0xbe, 0x7f, 0xf2, 0xf4, 0xa2, 0x31, 0x55, 0xf9, 0x7a, 0x12,
		0xd2, 0x2c, 0x5f, 0xcc, 0x41, 0x61, 0xff, 0xa5, 0x56, 0x5d, 0xdb, 0xdc, 0x3b, 0x58, 0xdf, 0xa9,
		0x2b, 0x09, 0x54, 0x02, 0x60, 0x82, 0xad, 0x9d, 0xbd, 0xda, 0xbe, 0x92, 0x0c, 0x9e, 0x1b, 0xbb,
		0xfb, 0x57, 0x2f, 0x2b, 0xa9, 0x00, 0x70, 0xc0, 0x05, 0xe9, 0xb0, 0xc2, 0x93, 0x97, 0x94, 0x0c,
		0x52, 0xa0, 0xc8, 0x09, 0x1a, 0x2f, 0xd6, 0x37, 0xaf, 0x5e, 0x56, 0xb2, 0xa3, 0x92, 0x27, 0x2f,
		0x29, 0x33, 0x68, 0x16, 0xf2, 0x4c, 0xb2, 0xbe, 0xb7, 0xb7, 0xa3, 0xe4, 0x02, 0xce, 0xf6, 0xbe,
		0xda, 0xd8, 0xdd, 0x56, 0xf2, 0x01, 0xe7, 0xb6, 0xba, 0x77, 0xd0, 0x52, 0x20, 0x60, 0x68, 0xd6,
		0xdb, 0xed, 0xda, 0x76, 0x5d, 0x29, 0x04, 0x1a, 0xeb, 0x2f, 0xed, 0xd7, 0xdb, 0x4a, 0x71, 0xc4,
		0xac, 0x27, 0x2f, 0x29, 0xb3, 0xc1, 0x2b, 0xea, 0xbb, 0x07, 0x4d, 0xa5, 0x84, 0xe6, 0x61, 0x96,
		0xbf, 0x42, 0x1a, 0x31, 0x17, 0x11, 0x5d, 0xbd, 0xac, 0x28, 0x43, 0x43, 0x38, 0xcb, 0xfc, 0x88,
		0xe0, 0xea, 0x65, 0x05, 0x55, 0x36, 0x20, 0xc3, 0xa2, 0x0b, 0x21, 0x28, 0xed, 0xd4, 0xd6, 0xeb,
		0x3b, 0xda, 0x5e, 0x6b, 0xbf, 0xb1, 0xb7, 0x5b, 0xdb, 0x51, 0x12, 0x43, 0x99, 0x5a, 0xff, 0xd4,
		0x41, 0x43, 0xad, 0x6f, 0x2a, 0xc9, 0xb0, 0xac, 0x55, 0xaf, 0xed, 0xd7, 0x37, 0x95, 0x54, 0x45,
		0x87, 0xc5, 0x49, 0x79, 0x72, 0xe2, 0xce, 0x08, 0x2d, 0x71, 0xf2, 0x84, 0x25, 0x66, 0x5c, 0x63,
		0x4b, 0xfc, 0x76, 0x12, 0x16, 0x26, 0xd4, 0x8a, 0x89, 0x2f, 0x79, 0x16, 0x32, 0x3c, 0x44, 0x79,
		0xf5, 0x7c, 0x74, 0x62, 0xd1, 0x61, 0x01, 0x3b, 0x56, 0x41, 0x19, 0x2e, 0xdc, 0x41, 0xa4, 0x4e,
		0xe8, 0x20, 0x28, 0xc5, 0x58, 0x4e, 0xff, 0xf1, 0xb1, 0x9c, 0xce, 0xcb, 0xde, 0xd5, 0x69, 0xca,
		0x1e, 0x93, 0x9d, 0x2e, 0xb7, 0x67, 0x26, 0xe4, 0xf6, 0x6b, 0x30, 0x3f, 0x46, 0x34, 0x75, 0x8e,
		0x7d, 0x2d, 0x01, 0xe5, 0x93, 0x9c, 0x13, 0x93, 0xe9, 0x92, 0x23, 0x99, 0xee, 0x5a, 0xd4, 0x83,
		0x17, 0x4e, 0x5e, 0x84, 0xb1, 0xb5, 0x7e, 0x2b, 0x01, 0x4b, 0x93, 0x3b, 0xc5, 0x89, 0x36, 0x7c,
		0x02, 0xb2, 0x7d, 0xe2, 0xf7, 0x6c, 0xd9, 0x2d, 0x3d, 0x3c, 0xa1, 0x06, 0xd3, 0xe1, 0xe8, 0x62,
		0x0b, 0x54, 0xb8, 0x88, 0xa7, 0x4e, 0x6a, 0xf7, 0xb8, 0x35, 0x63, 0x96, 0x7e, 0x36, 0x09, 0xf7,
		0x4e, 0x24, 0x9f, 0x68, 0xe8, 0xfd, 0x00, 0x86, 0xe5, 0x0c, 0x7c, 0xde, 0x11, 0xf1, 0x04, 0x9b,
		0x67, 0x12, 0x96, 0xbc, 0x68, 0xf2, 0x1c, 0xf8, 0xc1, 0x78, 0x8a, 0x8d, 0x03, 0x17, 0x31, 0x85,
		0xa7, 0x87, 0x86, 0xa6, 0x99, 0xa1, 0xe7, 0x4e, 0x98, 0xe9, 0x58, 0x60, 0x7e, 0x0c, 0x14, 0xdd,
		0x34, 0x88, 0xe5, 0x6b, 0x9e, 0xef, 0x12, 0xdc, 0x37, 0xac, 0x2e, 0xab, 0x20, 0xb9, 0x6a, 0xe6,
		0x08, 0x9b, 0x1e, 0x51, 0xe7, 0xf8, 0x70, 0x5b, 0x8e, 0x52, 0x04, 0x0b, 0x20, 0x37, 0x84, 0xc8,
		0x8e, 0x20, 0xf8, 0x70, 0x80, 0xa8, 0xfc, 0x7c, 0x1e, 0x0a, 0xa1, 0xbe, 0x1a, 0x5d, 0x80, 0xe2,
		0x75, 0x7c, 0x13, 0x6b, 0xf2, 0xac, 0xc4, 0x3d, 0x51, 0xa0, 0xb2, 0x96, 0x38, 0x2f, 0x7d, 0x0c,
		0x16, 0x99, 0x8a, 0x3d, 0xf0, 0x89, 0xab, 0xe9, 0x26, 0xf6, 0x3c, 0xe6, 0xb4, 0x1c, 0x53, 0x45,
		0x74, 0x6c, 0x8f, 0x0e, 0x6d, 0xc8, 0x11, 0x74, 0x05, 0x16, 0x18, 0xa2, 0x3f, 0x30, 0x7d, 0xc3,
		0x31, 0x89, 0x46, 0x4f, 0x6f, 0x1e, 0xab, 0x24, 0x81, 0x65, 0xf3, 0x54, 0xa3, 0x29, 0x14, 0xa8,
		0x45, 0x1e, 0xda, 0x84, 0xfb, 0x19, 0xac, 0x4b, 0x2c, 0xe2, 0x62, 0x9f, 0x68, 0xe4, 0x33, 0x03,
		0x6c, 0x7a, 0x1a, 0xb6, 0x3a, 0x5a, 0x0f, 0x7b, 0xbd, 0xf2, 0x22, 0x25, 0x58, 0x4f, 0x96, 0x13,
		0xea, 0x59, 0xaa, 0xb8, 0x2d, 0xf4, 0xea, 0x4c, 0xad, 0x66, 0x75, 0x3e, 0x89, 0xbd, 0x1e, 0xaa,
		0xc2, 0x12, 0x63, 0xf1, 0x7c, 0xd7, 0xb0, 0xba, 0x9a, 0xde, 0x23, 0xfa, 0x0d, 0x6d, 0xe0, 0x1f,
		0x3d, 0x5d, 0xbe, 0x2f, 0xfc, 0x7e, 0x66, 0x61, 0x9b, 0xe9, 0x6c, 0x50, 0x95, 0x03, 0xff, 0xe8,
		0x69, 0xd4, 0x86, 0x22, 0x5d, 0x8c, 0xbe, 0x71, 0x9b, 0x68, 0x47, 0xb6, 0xcb, 0x4a, 0x63, 0x69,
		0x42, 0x6a, 0x0a, 0x79, 0x70, 0x6d, 0x4f, 0x00, 0x9a, 0x76, 0x87, 0x54, 0x33, 0xed, 0x56, 0xbd,
		0xbe, 0xa9, 0x16, 0x24, 0xcb, 0x96, 0xed, 0xd2, 0x80, 0xea, 0xda, 0x81, 0x83, 0x0b, 0x3c, 0xa0,
		0xba, 0xb6, 0x74, 0xef, 0x15, 0x58, 0xd0, 0x75, 0x3e, 0x67, 0x43, 0xd7, 0xc4, 0x19, 0xcb, 0x2b,
		0x2b, 0x23, 0xce, 0xd2, 0xf5, 0x6d, 0xae, 0x20, 0x62, 0xdc, 0x43, 0xcf, 0xc0, 0xbd, 0x43, 0x67,
		0x85, 0x81, 0xf3, 0x63, 0xb3, 0x8c, 0x42, 0xaf, 0xc0, 0x82, 0x73, 0x3c, 0x0e, 0x44, 0x23, 0x6f,
		0x74, 0x8e, 0xa3, 0xb0, 0xa7, 0x60, 0xd1, 0xe9, 0x39, 0xe3, 0xb8, 0xc7, 0xc2, 0x38, 0xe4, 0xf4,
		0x9c, 0x28, 0xf0, 0x21, 0x76, 0xe0, 0x76, 0x89, 0x8e, 0x7d, 0xd2, 0x29, 0x9f, 0x09, 0xab, 0x87,
		0x06, 0xd0, 0x45, 0x50, 0x74, 0x5d, 0x23, 0x16, 0x3e, 0x34, 0x89, 0x86, 0x5d, 0x62, 0x61, 0xaf,
		0x7c, 0x3e, 0xac, 0x5c, 0xd2, 0xf5, 0x3a, 0x1b, 0xad, 0xb1, 0x41, 0xf4, 0x18, 0xcc, 0xdb, 0x87,
		0xd7, 0x75, 0x1e, 0x92, 0x9a, 0xe3, 0x92, 0x23, 0xe3, 0x95, 0xf2, 0x83, 0xcc, 0xbf, 0x73, 0x74,
		0x80, 0x05, 0x64, 0x8b, 0x89, 0xd1, 0xa3, 0xa0, 0xe8, 0x5e, 0x0f, 0xbb, 0x0e, 0xcb, 0xc9, 0x9e,
		0x83, 0x75, 0x52, 0x7e, 0x88, 0xab, 0x72, 0xf9, 0xae, 0x14, 0xd3, 0x2d, 0xe1, 0xdd, 0x32, 0x8e,
		0x7c, 0xc9, 0xf8, 0x08, 0xdf, 0x12, 0x4c, 0x26, 0xd8, 0x56, 0x41, 0xa1, 0xae, 0x18, 0x79, 0xf1,
		0x2a, 0x53, 0x2b, 0x39, 0x3d, 0x27, 0xfc, 0xde, 0x07, 0x60, 0x96, 0x6a, 0x0e, 0x5f, 0xfa, 0x28,
		0x6f, 0xc8, 0x9c, 0x5e, 0xe8, 0x8d, 0x97, 0x61, 0x89, 0x2a, 0xf5, 0x89, 0x8f, 0x3b, 0xd8, 0xc7,
		0x21, 0xed, 0x8f, 0x32, 0x6d, 0xea, 0xf7, 0xa6, 0x18, 0x1c, 0xb1, 0xd3, 0x1d, 0x1c, 0x1e, 0x07,
		0x91, 0xf5, 0x04, 0xb7, 0x93, 0xca, 0x64, 0x6c, 0x7d, 0x60, 0x4d, 0x77, 0xa5, 0x0a, 0xc5, 0x70,
		0xe0, 0xa3, 0x3c, 0xf0, 0xd0, 0x57, 0x12, 0xb4, 0x0b, 0xda, 0xd8, 0xdb, 0xa4, 0xfd, 0xcb, 0xcb,
		0x75, 0x25, 0x49, 0xfb, 0xa8, 0x9d, 0xc6, 0x7e, 0x5d, 0x53, 0x0f, 0x76, 0xf7, 0x1b, 0xcd, 0xba,
		0x92, 0x0a, 0x37, 0xec, 0xdf, 0x4a, 0x42, 0x69, 0xf4, 0xec, 0x85, 0x7e, 0x04, 0xce, 0xc8, 0x8b,
		0x12, 0x8f, 0xf8, 0xda, 0x2d, 0xc3, 0x65, 0x7b, 0xb1, 0x8f, 0x79, 0x5d, 0x0c, 0xa2, 0x61, 0x51,
		0x68, 0xb5, 0x89, 0xff, 0x82, 0xe1, 0xd2, 0x9d, 0xd6, 0xc7, 0x3e, 0xda, 0x81, 0xf3, 0x96, 0xad,
		0x79, 0x3e, 0xb6, 0x3a, 0xd8, 0xed, 0x68, 0xc3, 0x2b, 0x2a, 0x0d, 0xeb, 0x3a, 0xf1, 0x3c, 0x9b,
		0xd7, 0xc0, 0x80, 0xe5, 0x23, 0x96, 0xdd, 0x16, 0xca, 0xc3, 0xe2, 0x50, 0x13, 0xaa, 0x91, 0xc8,
		0x4d, 0x9d, 0x14, 0xb9, 0xf7, 0x41, 0xbe, 0x8f, 0x1d, 0x8d, 0x58, 0xbe, 0x7b, 0xcc, 0x3a, 0xee,
		0x9c, 0x9a, 0xeb, 0x63, 0xa7, 0x4e, 0x9f, 0x3f, 0x9c, 0x83, 0xcf, 0x3f, 0xa5, 0xa0, 0x18, 0xee,
		0xba, 0xe9, 0x21, 0x46, 0x67, 0x05, 0x2a, 0xc1, 0x52, 0xd8, 0x03, 0x77, 0xed, 0xd1, 0xd7, 0x36,
		0x68, 0xe5, 0xaa, 0x66, 0x79, 0x2f, 0xac, 0x72, 0x24, 0xed, 0x1a, 0x68, 0x68, 0x11, 0xde, 0x7b,
		0xe4, 0x54, 0xf1, 0x84, 0xb6, 0x21, 0x7b, 0xdd, 0x63, 0xdc, 0x59, 0xc6, 0xfd, 0xe0, 0xdd, 0xb9,
		0x9f, 0x6b, 0x33, 0xf2, 0xfc, 0x73, 0x6d, 0x6d, 0x77, 0x4f, 0x6d, 0xd6, 0x76, 0x54, 0x01, 0x47,
		0x67, 0x21, 0x6d, 0xe2, 0xdb, 0xc7, 0xa3, 0x35, 0x8e, 0x89, 0xa6, 0x75, 0xfc, 0x59, 0x48, 0xdf,
		0x22, 0xf8, 0xc6, 0x68, 0x65, 0x61, 0xa2, 0x0f, 0x30, 0xf4, 0x2f, 0x42, 0x86, 0xf9, 0x0b, 0x01,
		0x08, 0x8f, 0x29, 0xf7, 0xa0, 0x1c, 0xa4, 0x37, 0xf6, 0x54, 0x1a, 0xfe, 0x0a, 0x14, 0xb9, 0x54,
		0x6b, 0x35, 0xea, 0x1b, 0x75, 0x25, 0x59, 0xb9, 0x02, 0x59, 0xee, 0x04, 0xba, 0x35, 0x02, 0x37,
		0x28, 0xf7, 0x88, 0x47, 0xc1, 0x91, 0x90, 0xa3, 0x07, 0xcd, 0xf5, 0xba, 0xaa, 0x24, 0xc3, 0xcb,
		0xeb, 0x41, 0x31, 0xdc, 0x70, 0x7f, 0x38, 0x31, 0xf5, 0xcd, 0x04, 0x14, 0x42, 0x0d, 0x34, 0xed,
		0x7c, 0xb0, 0x69, 0xda, 0xb7, 0x34, 0x6c, 0x1a, 0xd8, 0x13, 0x41, 0x01, 0x4c, 0x54, 0xa3, 0x92,
		0x69, 0x17, 0xed, 0x43, 0x31, 0xfe, 0xcd, 0x04, 0x28, 0xd1, 0xde, 0x35, 0x62, 0x60, 0xe2, 0x87,
		0x6a, 0xe0, 0x1b, 0x09, 0x28, 0x8d, 0x36, 0xac, 0x11, 0xf3, 0x2e, 0xfc, 0x50, 0xcd, 0xfb, 0xe7,
		0x24, 0xcc, 0x8e, 0xb4, 0xa9, 0xd3, 0x5a, 0xf7, 0x19, 0x98, 0x37, 0x3a, 0xa4, 0xef, 0xd8, 0x3e,
		0xb1, 0xf4, 0x63, 0xcd, 0x24, 0x37, 0x89, 0x59, 0xae, 0xb0, 0x44, 0x71, 0xf1, 0xee, 0x8d, 0xf0,
		0x5a, 0x63, 0x88, 0xdb, 0xa1, 0xb0, 0xea, 0x42, 0x63, 0xb3, 0xde, 0x6c, 0xed, 0xed, 0xd7, 0x77,
		0x37, 0x5e, 0xd2, 0x0e, 0x76, 0x7f, 0x6c, 0x77, 0xef, 0x85, 0x5d, 0x55, 0x31, 0x22, 0x6a, 0x1f,
		0xe0, 0x56, 0x6f, 0x81, 0x12, 0x35, 0x0a, 0x9d, 0x81, 0x49, 0x66, 0x29, 0xf7, 0xa0, 0x05, 0x98,
		0xdb, 0xdd, 0xd3, 0xda, 0x8d, 0xcd, 0xba, 0x56, 0xdf, 0xda, 0xaa, 0x6f, 0xec, 0xb7, 0xf9, 0xd5,
		0x46, 0xa0, 0xbd, 0x3f, 0xba, 0xa9, 0x5f, 0x4f, 0xc1, 0xc2, 0x04, 0x4b, 0x50, 0x4d, 0x1c, 0x4a,
		0xf8, 0x39, 0xe9, 0x89, 0x69, 0xac, 0x5f, 0xa3, 0x5d, 0x41, 0x0b, 0xbb, 0xbe, 0x38, 0xc3, 0x3c,
		0x0a, 0xd4, 0x4b, 0x96, 0x6f, 0x1c, 0x19, 0xc4, 0x15, 0x37, 0x41, 0xfc, 0xa4, 0x32, 0x37, 0x94,
		0xf3, 0xcb, 0xa0, 0x8f, 0x02, 0x72, 0x6c, 0xcf, 0xf0, 0x8d, 0x9b, 0x44, 0x33, 0x2c, 0x79, 0x6d,
		0x44, 0x4f, 0x2e, 0x69, 0x55, 0x91, 0x23, 0x0d, 0xcb, 0x0f, 0xb4, 0x2d, 0xd2, 0xc5, 0x11, 0x6d,
		0x9a, 0xc0, 0x53, 0xaa, 0x22, 0x47, 0x02, 0xed, 0x0b, 0x50, 0xec, 0xd8, 0x03, 0xda, 0xce, 0x71,
		0x3d, 0x5a, 0x2f, 0x12, 0x6a, 0x81, 0xcb, 0x02, 0x15, 0xd1, 0xa8, 0x0f, 0xef, 0xab, 0x8a, 0x6a,
		0x81, 0xcb, 0xb8, 0xca, 0x23, 0x30, 0x87, 0xbb, 0x5d, 0x97, 0x92, 0x4b, 0x22, 0x7e, 0xf4, 0x28,
		0x05, 0x62, 0xa6, 0xb8, 0xfc, 0x1c, 0xe4, 0xa4, 0x1f, 0x68, 0x49, 0xa6, 0x9e, 0xd0, 0x1c, 0x7e,
		0x9e, 0x4e, 0xae, 0xe6, 0xd5, 0x9c, 0x25, 0x07, 0x2f, 0x40, 0xd1, 0xf0, 0xb4, 0xe1, 0xf5, 0x7b,
		0x72, 0x25, 0xb9, 0x9a, 0x53, 0x0b, 0x86, 0x17, 0x5c, 0x5d, 0x56, 0xde, 0x4a, 0x42, 0x69, 0xf4,
		0xf3, 0x01, 0xda, 0x84, 0x9c, 0x69, 0xeb, 0x98, 0x85, 0x16, 0xff, 0x76, 0xb5, 0x1a, 0xf3, 0xc5,
		0x61, 0x6d, 0x47, 0xe8, 0xab, 0x01, 0x72, 0xf9, 0xef, 0x13, 0x90, 0x93, 0x62, 0xb4, 0x04, 0x69,
		0x07, 0xfb, 0x3d, 0x46, 0x97, 0x59, 0x4f, 0x2a, 0x09, 0x95, 0x3d, 0x53, 0xb9, 0xe7, 0x60, 0x8b,
		0x85, 0x80, 0x90, 0xd3, 0x67, 0xba, 0xae, 0x26, 0xc1, 0x1d, 0x76, 0xae, 0xb1, 0xfb, 0x7d, 0x62,
		0xf9, 0x9e, 0x5c, 0x57, 0x21, 0xdf, 0x10, 0x62, 0xf4, 0x38, 0xcc, 0xfb, 0x2e, 0x36, 0xcc, 0x11,
		0xdd, 0x34, 0xd3, 0x55, 0xe4, 0x40, 0xa0, 0x5c, 0x85, 0xb3, 0x92, 0xb7, 0x43, 0x7c, 0xac, 0xf7,
		0x48, 0x67, 0x08, 0xca, 0xb2, 0xfb, 0x8b, 0x33, 0x42, 0x61, 0x53, 0x8c, 0x4b, 0x6c, 0xe5, 0x3b,
		0x09, 0x98, 0x97, 0x27, 0xb1, 0x4e, 0xe0, 0xac, 0x26, 0x00, 0xb6, 0x2c, 0xdb, 0x0f, 0xbb, 0x6b,
		0x3c, 0x94, 0xc7, 0x70, 0x6b, 0xb5, 0x00, 0xa4, 0x86, 0x08, 0x96, 0xfb, 0x00, 0xc3, 0x91, 0x13,
		0xdd, 0x76, 0x1e, 0x0a, 0xe2, 0xdb, 0x10, 0xfb, 0xc0, 0xc8, 0xcf, 0xee, 0xc0, 0x45, 0xf4, 0xc8,
		0x86, 0x16, 0x21, 0x73, 0x48, 0xba, 0x86, 0x25, 0x6e, 0x7c, 0xf9, 0x83, 0xbc, 0x61, 0x49, 0x07,
		0x37, 0x2c, 0xeb, 0x9f, 0x86, 0x05, 0xdd, 0xee, 0x47, 0xcd, 0x5d, 0x57, 0x22, 0xf7, 0x07, 0xde,
		0x27, 0x13, 0x2f, 0xc3, 0xb0, 0xc5, 0xfc, 0x5e, 0x22, 0xf1, 0xa5, 0x64, 0x6a, 0xbb, 0xb5, 0xfe,
		0x95, 0xe4, 0xf2, 0x36, 0x87, 0xb6, 0xe4, 0x4c, 0x55, 0x72, 0x64, 0x12, 0x9d, 0x5a, 0x0f, 0x5f,
		0x7e, 0x1c, 0x9e, 0xe8, 0x1a, 0x7e, 0x6f, 0x70, 0xb8, 0xa6, 0xdb, 0xfd, 0x8b, 0x5d, 0xbb, 0x6b,
		0x0f, 0xbf, 0xa9, 0xd2, 0x27, 0xf6, 0xc0, 0xfe, 0x13, 0xdf, 0x55, 0xf3, 0x81, 0x74, 0x39, 0xf6,
		0x23, 0x6c, 0x75, 0x17, 0x16, 0x84, 0xb2, 0xc6, 0x3e, 0xec, 0xf0, 0xe3, 0x09, 0xba, 0xeb, 0xe5,
		0x58, 0xf9, 0x6b, 0xef, 0xb0, 0x72, 0xad, 0xce, 0x0b, 0x28, 0x1d, 0xe3, 0x27, 0x98, 0xaa, 0x0a,
		0xf7, 0x8e, 0xf0, 0xf1, 0xad, 0x49, 0xdc, 0x18, 0xc6, 0x6f, 0x09, 0xc6, 0x85, 0x10, 0x63, 0x5b,
		0x40, 0xab, 0x1b, 0x30, 0x7b, 0x1a, 0xae, 0xbf, 0x15, 0x5c, 0x45, 0x12, 0x26, 0xd9, 0x86, 0x39,
		0x46, 0xa2, 0x0f, 0x3c, 0xdf, 0xee, 0xb3, 0xbc, 0x77, 0x77, 0x9a, 0xbf, 0x7b, 0x87, 0xef, 0x95,
		0x12, 0x85, 0x6d, 0x04, 0xa8, 0x6a, 0x15, 0xd8, 0xb7, 0xac, 0x0e, 0xd1, 0xcd, 0x18, 0x86, 0x6f,
		0x0b, 0x43, 0x02, 0xfd, 0xea, 0xf3, 0xb0, 0x48, 0xff, 0x67, 0x69, 0x29, 0x6c, 0x49, 0xfc, 0x4d,
		0x5a, 0xf9, 0x3b, 0xaf, 0xf1, 0xed, 0xb8, 0x10, 0x10, 0x84, 0x6c, 0x0a, 0xad, 0x62, 0x97, 0xf8,
		0x3e, 0x71, 0x3d, 0x0d, 0x9b, 0x93, 0xcc, 0x0b, 0x5d, 0x45, 0x94, 0x3f, 0xff, 0xee, 0xe8, 0x2a,
		0x6e, 0x73, 0x64, 0xcd, 0x34, 0xab, 0x07, 0x70, 0x66, 0x42, 0x54, 0x4c, 0xc1, 0xf9, 0xba, 0xe0,
		0x5c, 0x1c, 0x8b, 0x0c, 0x4a, 0xdb, 0x02, 0x29, 0x0f, 0xd6, 0x72, 0x0a, 0xce, 0xdf, 0x14, 0x9c,
		0x48, 0x60, 0xe5, 0x92, 0x52, 0xc6, 0xe7, 0x60, 0xfe, 0x26, 0x71, 0x0f, 0x6d, 0x4f, 0x5c, 0xff,
		0x4c, 0x41, 0xf7, 0x86, 0xa0, 0x9b, 0x13, 0x40, 0x76, 0x1f, 0x44, 0xb9, 0x9e, 0x81, 0xdc, 0x11,
		0xd6, 0xc9, 0x14, 0x14, 0x5f, 0x10, 0x14, 0x33, 0x54, 0x9f, 0x42, 0x6b, 0x50, 0xec, 0xda, 0xa2,
		0x32, 0xc5, 0xc3, 0xdf, 0x14, 0xf0, 0x82, 0xc4, 0x08, 0x0a, 0xc7, 0x76, 0x06, 0x26, 0x2d, 0x5b,
		0xf1, 0x14, 0xbf, 0x25, 0x29, 0x24, 0x46, 0x50, 0x9c, 0xc2, 0xad, 0xbf, 0x2d, 0x29, 0xbc, 0x90,
		0x3f, 0x9f, 0x85, 0x82, 0x6d, 0x99, 0xc7, 0xb6, 0x35, 0x8d, 0x11, 0x5f, 0x14, 0x0c, 0x20, 0x20,
		0x94, 0xe0, 0x1a, 0xe4, 0xa7, 0x5d, 0x88, 0xdf, 0x7d, 0x57, 0x6e, 0x0f, 0xb9, 0x02, 0xdb, 0x30,
		0x27, 0x13, 0x94, 0x61, 0x5b, 0x53, 0x50, 0x7c, 0x59, 0x50, 0x94, 0x42, 0x30, 0x31, 0x0d, 0x9f,
		0x78, 0x7e, 0x97, 0x4c, 0x43, 0xf2, 0x96, 0x9c, 0x86, 0x80, 0x08, 0x57, 0x1e, 0x12, 0x4b, 0xef,
		0x4d, 0xc7, 0xf0, 0x7b, 0xd2, 0x95, 0x12, 0x43, 0x29, 0x36, 0x60, 0xb6, 0x8f, 0x5d, 0xaf, 0x87,
		0xcd, 0xa9, 0x96, 0xe3, 0xf7, 0x05, 0x47, 0x31, 0x00, 0x09, 0x8f, 0x0c, 0xac, 0xd3, 0xd0, 0x7c,
		0x45, 0x7a, 0x24, 0x04, 0x13, 0x5b, 0xcf, 0xf3, 0xd9, 0x5d, 0xd9, 0x69, 0xd8, 0xfe, 0x40, 0x6e,
		0x3d, 0x8e, 0x6d, 0x86, 0x19, 0xaf, 0x41, 0xde, 0x33, 0x6e, 0x4f, 0x45, 0xf3, 0x87, 0x72, 0xa5,
		0x19, 0x80, 0x82, 0x5f, 0x82, 0xb3, 0x13, 0xcb, 0xc4, 0x14, 0x64, 0x7f, 0x24, 0xc8, 0x96, 0x26,
		0x94, 0x0a, 0x91, 0x12, 0x4e, 0x4b, 0xf9, 0xc7, 0x32, 0x25, 0x90, 0x08, 0x57, 0x8b, 0x9e, 0x15,
		0x3c, 0x7c, 0x74, 0x3a, 0xaf, 0xfd, 0x89, 0xf4, 0x1a, 0xc7, 0x8e, 0x78, 0x6d, 0x1f, 0x96, 0x04,
		0xe3, 0xe9, 0xd6, 0xf5, 0xab, 0x32, 0xb1, 0x72, 0xf4, 0xc1, 0xe8, 0xea, 0x7e, 0x1a, 0x96, 0x03,
		0x77, 0xca, 0xa6, 0xd4, 0xd3, 0xfa, 0xd8, 0x99, 0x82, 0xf9, 0x6b, 0x82, 0x59, 0x66, 0xfc, 0xa0,
		0xab, 0xf5, 0x9a, 0xd8, 0xa1, 0xe4, 0x2f, 0x42, 0x59, 0x92, 0x0f, 0x2c, 0x97, 0xe8, 0x76, 0xd7,
		0x32, 0x6e, 0x93, 0xce, 0x14, 0xd4, 0x7f, 0x1a, 0x59, 0xaa, 0x83, 0x10, 0x9c, 0x32, 0x37, 0x40,
		0x09, 0x7a, 0x15, 0xcd, 0xe8, 0x3b, 0xb6, 0xeb, 0xc7, 0x30, 0xfe, 0x99, 0x5c, 0xa9, 0x00, 0xd7,
		0x60, 0xb0, 0x6a, 0x1d, 0x4a, 0xec, 0x71, 0xda, 0x90, 0xfc, 0x73, 0x41, 0x34, 0x3b, 0x44, 0x89,
		0xc4, 0xa1, 0xdb, 0x7d, 0x07, 0xbb, 0xd3, 0xe4, 0xbf, 0xbf, 0x90, 0x89, 0x43, 0x40, 0x44, 0xe2,
		0xf0, 0x8f, 0x1d, 0x42, 0xab, 0xfd, 0x14, 0x0c, 0x5f, 0x97, 0x89, 0x43, 0x62, 0x04, 0x85, 0x6c,
		0x18, 0xa6, 0xa0, 0xf8, 0x4b, 0x49, 0x21, 0x31, 0x94, 0xe2, 0x53, 0xc3, 0x42, 0xeb, 0x92, 0xae,
		0xe1, 0xf9, 0x2e, 0x6f, 0x85, 0xef, 0x4e, 0xf5, 0x57, 0xef, 0x8e, 0x36, 0x61, 0x6a, 0x08, 0x4a,
		0x33, 0x91, 0xb8, 0x42, 0x65, 0x27, 0xa5, 0x78, 0xc3, 0xbe, 0x21, 0x33, 0x51, 0x08, 0x46, 0x6d,
		0x0b, 0x75, 0x88, 0xd4, 0xed, 0x3a, 0x3d, 0x1f, 0x4c, 0x41, 0xf7, 0xcd, 0x88, 0x71, 0x6d, 0x89,
		0xa5, 0x9c, 0xa1, 0xfe, 0x67, 0x60, 0xdd, 0x20, 0xc7, 0x53, 0x45, 0xe7, 0x5f, 0x47, 0xfa, 0x9f,
		0x03, 0x8e, 0xe4, 0x39, 0x64, 0x2e, 0xd2, 0x4f, 0xa1, 0xb8, 0x5f, 0x01, 0x95, 0x7f, 0xf2, 0x8e,
		0x98, 0xef, 0x68, 0x3b, 0x55, 0xdd, 0xa1, 0x41, 0x3e, 0xda, 0xf4, 0xc4, 0x93, 0xbd, 0x76, 0x27,
		0x88, 0xf3, 0x91, 0x9e, 0xa7, 0xba, 0x05, 0xb3, 0x23, 0x0d, 0x4f, 0x3c, 0xd5, 0x4f, 0x09, 0xaa,
		0x62, 0xb8, 0xdf, 0xa9, 0x5e, 0x81, 0x34, 0x6d, 0x5e, 0xe2, 0xe1, 0x3f, 0x2d, 0xe0, 0x4c, 0xbd,
		0xfa, 0x71, 0xc8, 0xc9, 0xa6, 0x25, 0x1e, 0xfa, 0x33, 0x02, 0x1a, 0x40, 0x28, 0x5c, 0x36, 0x2c,
		0xf1, 0xf0, 0x9f, 0x95, 0x70, 0x09, 0xa1, 0xf0, 0xe9, 0x5d, 0xf8, 0x37, 0x3f, 0x97, 0x16, 0x45,
		0x47, 0xfa, 0xee, 0x1a, 0xcc, 0x88, 0x4e, 0x25, 0x1e, 0xfd, 0x59, 0xf1, 0x72, 0x89, 0xa8, 0x3e,
		0x05, 0x99, 0x29, 0x1d, 0xfe, 0x0b, 0x02, 0xca, 0xf5, 0xab, 0x1b, 0x50, 0x08, 0x75, 0x27, 0xf1,
		0xf0, 0x5f, 0x14, 0xf0, 0x30, 0x8a, 0x9a, 0x2e, 0xba, 0x93, 0x78, 0x82, 0x5f, 0x92, 0xa6, 0x0b,
		0x04, 0x75, 0x9b, 0x6c, 0x4c, 0xe2, 0xd1, 0xbf, 0x2c, 0xbd, 0x2e, 0x21, 0xd5, 0x67, 0x21, 0x1f,
		0x14, 0x9b, 0x78, 0xfc, 0xaf, 0x08, 0xfc, 0x10, 0x43, 0x3d, 0x10, 0x2a, 0x76, 0xf1, 0x14, 0xbf,
		0x2a, 0x3d, 0x10, 0x42, 0xd1, 0x6d, 0x14, 0x6d, 0x60, 0xe2, 0x99, 0x7e, 0x4d, 0x6e, 0xa3, 0x48,
		0xff, 0x42, 0x57, 0x93, 0xe5, 0xfc, 0x78, 0x8a, 0x5f, 0x97, 0xab, 0xc9, 0xf4, 0xa9, 0x19, 0xd1,
		0x8e, 0x20, 0x9e, 0xe3, 0x37, 0xa4, 0x19, 0x91, 0x86, 0xa0, 0xda, 0x02, 0x34, 0xde, 0x0d, 0xc4,
		0xf3, 0x7d, 0x4e, 0xf0, 0xcd, 0x8f, 0x35, 0x03, 0xd5, 0x17, 0x60, 0x69, 0x72, 0x27, 0x10, 0xcf,
		0xfa, 0xf9, 0x3b, 0x91, 0xb3, 0x5b, 0xb8, 0x11, 0xa8, 0xee, 0x0f, 0x4b, 0x4a, 0xb8, 0x0b, 0x88,
		0xa7, 0x7d, 0xfd, 0xce, 0x68, 0xe2, 0x0e, 0x37, 0x01, 0xd5, 0x1a, 0xc0, 0xb0, 0x00, 0xc7, 0x73,
		0xbd, 0x21, 0xb8, 0x42, 0x20, 0xba, 0x35, 0x44, 0xfd, 0x8d, 0xc7, 0x7f, 0x41, 0x6e, 0x0d, 0x81,
		0xa0, 0x5b, 0x43, 0x96, 0xde, 0x78, 0xf4, 0x9b, 0x72, 0x6b, 0x48, 0x08, 0x8d, 0xec, 0x50, 0x75,
		0x8b, 0x67, 0xf8, 0xa2, 0x8c, 0xec, 0x10, 0xaa, 0xba, 0x0b, 0xf3, 0x63, 0x05, 0x31, 0x9e, 0xea,
		0x4b, 0x82, 0x4a, 0x89, 0xd6, 0xc3, 0x70, 0xf1, 0x12, 0xc5, 0x30, 0x9e, 0xed, 0x77, 0x22, 0xc5,
		0x4b, 0xd4, 0xc2, 0xea, 0x35, 0xc8, 0x59, 0x03, 0xd3, 0xa4, 0x9b, 0x07, 0xdd, 0xfd, 0x97, 0x7b,
		0xe5, 0x7f, 0x7f, 0x5f, 0x78, 0x47, 0x02, 0xaa, 0x57, 0x20, 0x43, 0xfa, 0x87, 0xa4, 0x13, 0x87,
		0xfc, 0x8f, 0xf7, 0x65, 0xc2, 0xa4, 0xda, 0xd5, 0x67, 0x01, 0xf8, 0xd5, 0x08, 0xfb, 0xec, 0x17,
		0x83, 0xfd, 0xcf, 0xf7, 0xc5, 0x6f, 0x6a, 0x86, 0x90, 0x21, 0x01, 0xff, 0x85, 0xce, 0xdd, 0x09,
		0xde, 0x1d, 0x25, 0x60, 0x2b, 0xf2, 0x0c, 0xcc, 0x5c, 0xf7, 0x6c, 0xcb, 0xc7, 0xdd, 0x38, 0xf4,
		0x7f, 0x09, 0xb4, 0xd4, 0xa7, 0x0e, 0xeb, 0xdb, 0x2e, 0xf1, 0x71, 0xd7, 0x8b, 0xc3, 0xfe, 0xb7,
		0xc0, 0x06, 0x00, 0x0a, 0xd6, 0xb1, 0xe7, 0x4f, 0x33, 0xef, 0xff, 0x91, 0x60, 0x09, 0xa0, 0x46,
		0xd3, 0xff, 0x6f, 0x90, 0xe3, 0x38, 0xec, 0x7b, 0xd2, 0x68, 0xa1, 0x5f, 0xfd, 0x38, 0xe4, 0xe9,
		0xbf, 0xfc, 0x87, 0x72, 0x31, 0xe0, 0xff, 0x15, 0xe0, 0x21, 0x82, 0xbe, 0xd9, 0xf3, 0x3b, 0xbe,
		0x11, 0xef, 0xec, 0xff, 0x13, 0x2b, 0x2d, 0xf5, 0xab, 0x35, 0x28, 0x78, 0x7e, 0xa7, 0x33, 0x10,
		0xfd, 0x69, 0x0c, 0xfc, 0xff, 0xdf, 0x0f, 0xae, 0x2c, 0x02, 0x0c, 0x5d, 0xed, 0x5b, 0x37, 0x7c,
		0xc7, 0x66, 0x9f, 0x39, 0xe2, 0x18, 0xee, 0x08, 0x86, 0x10, 0x64, 0xbd, 0x3e, 0xf9, 0xfa, 0x16,
		0xb6, 0xed, 0x6d, 0x9b, 0x5f, 0xdc, 0xbe, 0x5c, 0x89, 0xbf, 0x81, 0x85, 0xf7, 0x92, 0x90, 0xef,
		0x7b, 0x5d, 0x71, 0x09, 0xcb, 0x33, 0x14, 0xad, 0xc0, 0xde, 0xf2, 0xe9, 0xee, 0x6f, 0x2b, 0x5b,
		0x90, 0x5d, 0x37, 0xba, 0x4d, 0xaf, 0x8b, 0x16, 0x21, 0xc3, 0xac, 0x67, 0x1f, 0x1f, 0x53, 0x2a,
		0x7f, 0x40, 0x0f, 0x43, 0xaa, 0xe9, 0x75, 0xc5, 0xef, 0xda, 0x16, 0xd7, 0x86, 0x2f, 0x5a, 0x6b,
		0xf7, 0xb1, 0x69, 0x36, 0xbd, 0xae, 0x4a, 0x15, 0x2a, 0x2e, 0xe4, 0xa4, 0x00, 0xad, 0x40, 0xa1,
		0xad, 0x63, 0x77, 0x7d, 0xe0, 0xb5, 0x7d, 0xdb, 0x91, 0xbf, 0xdb, 0x0a, 0x89, 0xd0, 0x2a, 0xcc,
		0x6d, 0x99, 0x46, 0xb7, 0xe7, 0xb7, 0xb0, 0x8b, 0xf5, 0xde, 0xc0, 0x27, 0xe5, 0xe2, 0x4a, 0x6a,
		0x75, 0x46, 0x8d, 0x8a, 0xd1, 0x32, 0xe4, 0x9a, 0xd8, 0x69, 0xf7, 0xb0, 0x7b, 0x83, 0xfd, 0x0a,
		0x28, 0xaf, 0x06, 0xcf, 0x95, 0xe7, 0x21, 0xdb, 0xe2, 0xdf, 0xfd, 0x97, 0x20, 0xdd, 0xe0, 0x1f,
		0x11, 0x52, 0xab, 0x29, 0x7e, 0xeb, 0x4e, 0x9f, 0xd1, 0x32, 0x64, 0xb7, 0x4c, 0x1b, 0xfb, 0x1e,
		0xfb, 0xa5, 0x63, 0x82, 0x8d, 0x08, 0x09, 0x2a, 0x43, 0xe6, 0xc0, 0x90, 0x1f, 0x11, 0x66, 0xd9,
		0x10, 0x17, 0xac, 0x2f, 0xbe, 0xf7, 0xdd, 0x73, 0x89, 0xef, 0x7d, 0xf7, 0x5c, 0xe2, 0xab, 0xff,
		0x72, 0x2e, 0xf1, 0x0f, 0x6f, 0x9f, 0xbb, 0xe7, 0x1f, 0xdf, 0x3e, 0x77, 0xcf, 0xf7, 0x03, 0x00,
		0x00, 0xff, 0xff, 0xef, 0xca, 0x82, 0x66, 0x67, 0x34, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *BigMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.BigMsg{")
	if this.Field != nil {
		s = append(s, "Field: "+valueToGoStringMsg(this.Field, "int64")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.SmallMsg{")
	if this.ScarBusStop != nil {
		s = append(s, "ScarBusStop: "+valueToGoStringMsg(this.ScarBusStop, "string")+",\n")
	}
	if this.FlightParachute != nil {
		s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	}
	if this.MapShark != nil {
		s = append(s, "MapShark: "+valueToGoStringMsg(this.MapShark, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Packed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.Packed{")
	if this.Ints != nil {
		s = append(s, "Ints: "+fmt.Sprintf("%#v", this.Ints)+",\n")
	}
	if this.Floats != nil {
		s = append(s, "Floats: "+fmt.Sprintf("%#v", this.Floats)+",\n")
	}
	if this.Uints != nil {
		s = append(s, "Uints: "+fmt.Sprintf("%#v", this.Uints)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedBigMsg(r randyMsg, easy bool) *BigMsg {
	this := &BigMsg{}
	if r.Intn(5) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field = &v1
	}
	if r.Intn(5) != 0 {
		this.Msg = NewPopulatedSmallMsg(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 4)
	}
	return this
}

func NewPopulatedSmallMsg(r randyMsg, easy bool) *SmallMsg {
	this := &SmallMsg{}
	if r.Intn(5) != 0 {
		v2 := string(randStringMsg(r))
		this.ScarBusStop = &v2
	}
	if r.Intn(5) != 0 {
		v3 := r.Intn(10)
		this.FlightParachute = make([]uint32, v3)
		for i := 0; i < v3; i++ {
			this.FlightParachute[i] = uint32(r.Uint32())
		}
	}
	if r.Intn(5) != 0 {
		v4 := string(randStringMsg(r))
		this.MapShark = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 19)
	}
	return this
}

func NewPopulatedPacked(r randyMsg, easy bool) *Packed {
	this := &Packed{}
	if r.Intn(5) != 0 {
		v5 := r.Intn(10)
		this.Ints = make([]int64, v5)
		for i := 0; i < v5; i++ {
			this.Ints[i] = int64(r.Int63())
			if r.Intn(2) == 0 {
				this.Ints[i] *= -1
			}
		}
	}
	if r.Intn(5) != 0 {
		v6 := r.Intn(10)
		this.Floats = make([]float64, v6)
		for i := 0; i < v6; i++ {
			this.Floats[i] = float64(r.Float64())
			if r.Intn(2) == 0 {
				this.Floats[i] *= -1
			}
		}
	}
	if r.Intn(5) != 0 {
		v7 := r.Intn(10)
		this.Uints = make([]uint32, v7)
		for i := 0; i < v7; i++ {
			this.Uints[i] = uint32(r.Uint32())
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 7)
	}
	return this
}

type randyMsg interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMsg(r randyMsg) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMsg(r randyMsg) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneMsg(r)
	}
	return string(tmps)
}
func randUnrecognizedMsg(r randyMsg, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMsg(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMsg(dAtA []byte, r randyMsg, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMsg(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
