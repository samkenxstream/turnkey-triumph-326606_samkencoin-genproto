// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/type/timeofday/timeofday.proto
// DO NOT EDIT!

/*
Package google_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/type/timeofday/timeofday.proto

It has these top-level messages:
	TimeOfDay
*/
package google_type // import "google.golang.org/genproto/googleapis/type/timeofday"

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

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may chose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	// to allow the value "24:00:00" for scenarios like business closing time.
	Hours int32 `protobuf:"varint,1,opt,name=hours" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes" json:"minutes,omitempty"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,3,opt,name=seconds" json:"seconds,omitempty"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos int32 `protobuf:"varint,4,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *TimeOfDay) Reset()                    { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string            { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()               {}
func (*TimeOfDay) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*TimeOfDay)(nil), "google.type.TimeOfDay")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/type/timeofday/timeofday.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0xea, 0x97, 0x64, 0xe6, 0xa6, 0xe6, 0xa7, 0xa5, 0x24, 0x56, 0x22, 0x58, 0x7a, 0x60,
	0x95, 0x42, 0xdc, 0x50, 0x53, 0x40, 0xca, 0x94, 0xb2, 0xb9, 0x38, 0x43, 0x80, 0xf2, 0xfe, 0x69,
	0x2e, 0x89, 0x95, 0x42, 0x22, 0x5c, 0xac, 0x19, 0xf9, 0xa5, 0x45, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x10, 0x8e, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x66, 0x5e, 0x69, 0x49, 0x6a, 0xb1,
	0x04, 0x13, 0x58, 0x1c, 0xc6, 0x05, 0xc9, 0x14, 0xa7, 0x26, 0xe7, 0xe7, 0xa5, 0x14, 0x4b, 0x30,
	0x43, 0x64, 0xa0, 0x5c, 0x90, 0x49, 0x79, 0x89, 0x79, 0xf9, 0xc5, 0x12, 0x2c, 0x10, 0x93, 0xc0,
	0x1c, 0x27, 0x4d, 0x2e, 0xfe, 0xe4, 0xfc, 0x5c, 0x3d, 0x24, 0xfb, 0x9d, 0xf8, 0xe0, 0xb6, 0x07,
	0x80, 0x1c, 0x17, 0xc0, 0xb8, 0x88, 0x89, 0xd9, 0x3d, 0x24, 0x20, 0x89, 0x0d, 0xec, 0x56, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x92, 0x39, 0x10, 0xf3, 0x00, 0x00, 0x00,
}