// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stellarstation/api/v1/orbit/orbit.proto

package orbit // import "github.com/infostellarinc/go-stellarstation/api/v1/orbit"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Unparsed TLE data for a satellite - https://en.wikipedia.org/wiki/Two-line_element_set
type Tle struct {
	// The first line of the TLE. Not a title line.
	Line_1 string `protobuf:"bytes,1,opt,name=line_1,json=line1,proto3" json:"line_1,omitempty"`
	// The second line of the TLE.
	Line_2               string   `protobuf:"bytes,2,opt,name=line_2,json=line2,proto3" json:"line_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tle) Reset()         { *m = Tle{} }
func (m *Tle) String() string { return proto.CompactTextString(m) }
func (*Tle) ProtoMessage()    {}
func (*Tle) Descriptor() ([]byte, []int) {
	return fileDescriptor_orbit_b15e55ba0dabff28, []int{0}
}
func (m *Tle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tle.Unmarshal(m, b)
}
func (m *Tle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tle.Marshal(b, m, deterministic)
}
func (dst *Tle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tle.Merge(dst, src)
}
func (m *Tle) XXX_Size() int {
	return xxx_messageInfo_Tle.Size(m)
}
func (m *Tle) XXX_DiscardUnknown() {
	xxx_messageInfo_Tle.DiscardUnknown(m)
}

var xxx_messageInfo_Tle proto.InternalMessageInfo

func (m *Tle) GetLine_1() string {
	if m != nil {
		return m.Line_1
	}
	return ""
}

func (m *Tle) GetLine_2() string {
	if m != nil {
		return m.Line_2
	}
	return ""
}

func init() {
	proto.RegisterType((*Tle)(nil), "stellarstation.api.v1.orbit.Tle")
}

func init() {
	proto.RegisterFile("stellarstation/api/v1/orbit/orbit.proto", fileDescriptor_orbit_b15e55ba0dabff28)
}

var fileDescriptor_orbit_b15e55ba0dabff28 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2e, 0x49, 0xcd,
	0xc9, 0x49, 0x2c, 0x2a, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f,
	0x33, 0xd4, 0xcf, 0x2f, 0x4a, 0xca, 0x2c, 0x81, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xd2, 0xa8, 0x0a, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xca, 0x0c, 0xf5, 0xc0, 0x4a, 0x94, 0x8c, 0xb9,
	0x98, 0x43, 0x72, 0x52, 0x85, 0x44, 0xb9, 0xd8, 0x72, 0x32, 0xf3, 0x52, 0xe3, 0x0d, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x58, 0x41, 0x3c, 0x43, 0xb8, 0xb0, 0x91, 0x04, 0x13, 0x42, 0xd8,
	0xc8, 0x29, 0x93, 0x4b, 0x3e, 0x39, 0x3f, 0x57, 0x0f, 0x8f, 0xb9, 0x4e, 0x5c, 0xfe, 0x20, 0x2a,
	0x00, 0xe4, 0x80, 0x00, 0xc6, 0x28, 0x8b, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0xfd, 0xcc, 0xbc, 0xb4, 0x7c, 0xa8, 0xce, 0xcc, 0xbc, 0x64, 0xfd, 0xf4, 0x7c, 0x5d, 0x3c,
	0x1e, 0x49, 0x62, 0x03, 0xfb, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x1b, 0xd6, 0x67,
	0xee, 0x00, 0x00, 0x00,
}
