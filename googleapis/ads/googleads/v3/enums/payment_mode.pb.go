// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/payment_mode.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible payment modes.
type PaymentModeEnum_PaymentMode int32

const (
	// Not specified.
	PaymentModeEnum_UNSPECIFIED PaymentModeEnum_PaymentMode = 0
	// Used for return value only. Represents value unknown in this version.
	PaymentModeEnum_UNKNOWN PaymentModeEnum_PaymentMode = 1
	// Pay per click.
	PaymentModeEnum_CLICKS PaymentModeEnum_PaymentMode = 4
	// Pay per conversion value. This mode is only supported by campaigns with
	// AdvertisingChannelType.HOTEL, BiddingStrategyType.COMMISSION, and
	// BudgetType.HOTEL_ADS_COMMISSION.
	PaymentModeEnum_CONVERSION_VALUE PaymentModeEnum_PaymentMode = 5
	// Pay per conversion. This mode is only supported by campaigns with
	// AdvertisingChannelType.DISPLAY (excluding
	// AdvertisingChannelSubType.DISPLAY_GMAIL), BiddingStrategyType.TARGET_CPA,
	// and BudgetType.FIXED_CPA. The customer must also be eligible for this
	// mode. See Customer.eligibility_failure_reasons for details.
	PaymentModeEnum_CONVERSIONS PaymentModeEnum_PaymentMode = 6
	// Pay per guest stay value. This mode is only supported by campaigns with
	// AdvertisingChannelType.HOTEL, BiddingStrategyType.COMMISSION, and
	// BudgetType.HOTEL_ADS_COMMISSION.
	PaymentModeEnum_GUEST_STAY PaymentModeEnum_PaymentMode = 7
)

var PaymentModeEnum_PaymentMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	4: "CLICKS",
	5: "CONVERSION_VALUE",
	6: "CONVERSIONS",
	7: "GUEST_STAY",
}

var PaymentModeEnum_PaymentMode_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"CLICKS":           4,
	"CONVERSION_VALUE": 5,
	"CONVERSIONS":      6,
	"GUEST_STAY":       7,
}

func (x PaymentModeEnum_PaymentMode) String() string {
	return proto.EnumName(PaymentModeEnum_PaymentMode_name, int32(x))
}

func (PaymentModeEnum_PaymentMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7bf76f63298a9af6, []int{0, 0}
}

// Container for enum describing possible payment modes.
type PaymentModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentModeEnum) Reset()         { *m = PaymentModeEnum{} }
func (m *PaymentModeEnum) String() string { return proto.CompactTextString(m) }
func (*PaymentModeEnum) ProtoMessage()    {}
func (*PaymentModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf76f63298a9af6, []int{0}
}

func (m *PaymentModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentModeEnum.Unmarshal(m, b)
}
func (m *PaymentModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentModeEnum.Marshal(b, m, deterministic)
}
func (m *PaymentModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentModeEnum.Merge(m, src)
}
func (m *PaymentModeEnum) XXX_Size() int {
	return xxx_messageInfo_PaymentModeEnum.Size(m)
}
func (m *PaymentModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PaymentModeEnum_PaymentMode", PaymentModeEnum_PaymentMode_name, PaymentModeEnum_PaymentMode_value)
	proto.RegisterType((*PaymentModeEnum)(nil), "google.ads.googleads.v3.enums.PaymentModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/payment_mode.proto", fileDescriptor_7bf76f63298a9af6)
}

var fileDescriptor_7bf76f63298a9af6 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0x7e, 0xbb, 0x57, 0x37, 0xc8, 0xc0, 0x85, 0xe2, 0x49, 0xdc, 0x61, 0xfb, 0x00, 0xa9, 0xd0,
	0x5b, 0x3c, 0x65, 0xb5, 0x8e, 0xb2, 0xd9, 0x15, 0xbb, 0x56, 0x94, 0xc2, 0x88, 0x26, 0x84, 0xc1,
	0x9a, 0x94, 0xa5, 0x1b, 0x78, 0xf4, 0xab, 0x78, 0xf4, 0xa3, 0xf8, 0x3d, 0xbc, 0xf8, 0x29, 0xa4,
	0x89, 0x9b, 0xbb, 0xe8, 0x25, 0x3c, 0xfc, 0x9e, 0x3f, 0x3c, 0x79, 0xc0, 0x85, 0x50, 0x4a, 0xac,
	0xb8, 0x47, 0x99, 0xf6, 0x2c, 0x6c, 0xd0, 0xd6, 0xf7, 0xb8, 0xdc, 0x94, 0xda, 0xab, 0xe8, 0x73,
	0xc9, 0x65, 0xbd, 0x28, 0x15, 0xe3, 0xa8, 0x5a, 0xab, 0x5a, 0xb9, 0x7d, 0x2b, 0x43, 0x94, 0x69,
	0xb4, 0x77, 0xa0, 0xad, 0x8f, 0x8c, 0xe3, 0xec, 0x7c, 0x17, 0x58, 0x2d, 0x3d, 0x2a, 0xa5, 0xaa,
	0x69, 0xbd, 0x54, 0x52, 0x5b, 0xf3, 0xf0, 0xc5, 0x01, 0xbd, 0xc4, 0x66, 0xde, 0x28, 0xc6, 0x43,
	0xb9, 0x29, 0x87, 0x12, 0x74, 0x0f, 0x4e, 0x6e, 0x0f, 0x74, 0xb3, 0x38, 0x4d, 0xc2, 0x20, 0xba,
	0x8e, 0xc2, 0x2b, 0xf8, 0xcf, 0xed, 0x82, 0x4e, 0x16, 0x4f, 0xe2, 0xd9, 0x5d, 0x0c, 0x1d, 0x17,
	0x80, 0x76, 0x30, 0x8d, 0x82, 0x49, 0x0a, 0x8f, 0xdc, 0x53, 0x00, 0x83, 0x59, 0x9c, 0x87, 0xb7,
	0x69, 0x34, 0x8b, 0x17, 0x39, 0x99, 0x66, 0x21, 0x3c, 0x6e, 0xfc, 0x3f, 0xd7, 0x14, 0xb6, 0xdd,
	0x13, 0x00, 0xc6, 0x59, 0x98, 0xce, 0x17, 0xe9, 0x9c, 0xdc, 0xc3, 0xce, 0xe8, 0xc3, 0x01, 0x83,
	0x27, 0x55, 0xa2, 0x3f, 0xff, 0x31, 0x82, 0x07, 0x9d, 0x92, 0xa6, 0x7b, 0xe2, 0x3c, 0x8c, 0xbe,
	0x2d, 0x42, 0xad, 0xa8, 0x14, 0x48, 0xad, 0x85, 0x27, 0xb8, 0x34, 0x3f, 0xdb, 0x8d, 0x57, 0x2d,
	0xf5, 0x2f, 0x5b, 0x5e, 0x9a, 0xf7, 0xb5, 0xf5, 0x7f, 0x4c, 0xc8, 0x5b, 0xab, 0x3f, 0xb6, 0x51,
	0x84, 0x69, 0x64, 0x61, 0x83, 0x72, 0x1f, 0x35, 0x93, 0xe8, 0xf7, 0x1d, 0x5f, 0x10, 0xa6, 0x8b,
	0x3d, 0x5f, 0xe4, 0x7e, 0x61, 0xf8, 0xcf, 0xd6, 0xc0, 0x1e, 0x31, 0x26, 0x4c, 0x63, 0xbc, 0x57,
	0x60, 0x9c, 0xfb, 0x18, 0x1b, 0xcd, 0x63, 0xdb, 0x14, 0xf3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x70, 0x40, 0x8a, 0xe3, 0x01, 0x00, 0x00,
}