// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_label.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A relationship between an ad group and a label.
type AdGroupLabel struct {
	// The resource name of the ad group label.
	// Ad group label resource names have the form:
	// `customers/{customer_id}/adGroupLabels/{ad_group_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group to which the label is attached.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The label assigned to the ad group.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupLabel) Reset()         { *m = AdGroupLabel{} }
func (m *AdGroupLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupLabel) ProtoMessage()    {}
func (*AdGroupLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_90328d22925d96e4, []int{0}
}

func (m *AdGroupLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupLabel.Unmarshal(m, b)
}
func (m *AdGroupLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupLabel.Merge(m, src)
}
func (m *AdGroupLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupLabel.Size(m)
}
func (m *AdGroupLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupLabel proto.InternalMessageInfo

func (m *AdGroupLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupLabel) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupLabel)(nil), "google.ads.googleads.v1.resources.AdGroupLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_label.proto", fileDescriptor_90328d22925d96e4)
}

var fileDescriptor_90328d22925d96e4 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x69, 0xc7, 0xff, 0xaf, 0xd6, 0x79, 0xb0, 0xa7, 0x31, 0x86, 0x6c, 0xca, 0x60, 0xa7,
	0x84, 0x4e, 0x50, 0x88, 0xa7, 0xee, 0x32, 0x10, 0x91, 0x31, 0xa1, 0x07, 0x29, 0x8c, 0x77, 0x4b,
	0x0c, 0x85, 0x36, 0x29, 0x49, 0x3b, 0xbf, 0x8e, 0x78, 0xf4, 0xa3, 0xf8, 0x51, 0xfc, 0x10, 0x22,
	0x6d, 0x9a, 0xe0, 0x49, 0xbd, 0x3d, 0xed, 0xfb, 0x7b, 0x9e, 0xe7, 0x4d, 0x12, 0x5c, 0x71, 0x29,
	0x79, 0xce, 0x30, 0x50, 0x8d, 0x8d, 0x6c, 0xd4, 0x3e, 0xc2, 0x8a, 0x69, 0x59, 0xab, 0x1d, 0xd3,
	0x18, 0xe8, 0x86, 0x2b, 0x59, 0x97, 0x9b, 0x1c, 0xb6, 0x2c, 0x47, 0xa5, 0x92, 0x95, 0x0c, 0x27,
	0x06, 0x46, 0x40, 0x35, 0x72, 0x3e, 0xb4, 0x8f, 0x90, 0xf3, 0x0d, 0x47, 0x36, 0xba, 0xcc, 0x30,
	0x08, 0x21, 0x2b, 0xa8, 0x32, 0x29, 0xb4, 0x09, 0x18, 0x9e, 0x75, 0xd3, 0xf6, 0x6b, 0x5b, 0x3f,
	0xe1, 0x67, 0x05, 0x65, 0xc9, 0x54, 0x37, 0x3f, 0x7f, 0xf1, 0x82, 0x7e, 0x4c, 0x97, 0x4d, 0xf1,
	0x5d, 0xd3, 0x1b, 0x5e, 0x04, 0x27, 0x36, 0x7b, 0x23, 0xa0, 0x60, 0x03, 0x6f, 0xec, 0xcd, 0x8e,
	0xd6, 0x7d, 0xfb, 0xf3, 0x1e, 0x0a, 0x16, 0x5e, 0x07, 0x87, 0x76, 0xdd, 0x81, 0x3f, 0xf6, 0x66,
	0xc7, 0xf3, 0x51, 0xb7, 0x1e, 0xb2, 0x45, 0xe8, 0xa1, 0x52, 0x99, 0xe0, 0x09, 0xe4, 0x35, 0x5b,
	0x1f, 0x80, 0xa9, 0x08, 0xe7, 0xc1, 0xbf, 0xf6, 0x78, 0x83, 0xde, 0x1f, 0x5c, 0x06, 0x5d, 0x7c,
	0x7a, 0xc1, 0x74, 0x27, 0x0b, 0xf4, 0xeb, 0x55, 0x2c, 0x4e, 0xbf, 0x9f, 0x64, 0xd5, 0x44, 0xae,
	0xbc, 0xc7, 0xdb, 0xce, 0xc7, 0x65, 0x0e, 0x82, 0x23, 0xa9, 0x38, 0xe6, 0x4c, 0xb4, 0x85, 0xf6,
	0x29, 0xca, 0x4c, 0xff, 0xf0, 0x32, 0x37, 0x4e, 0xbd, 0xfa, 0xbd, 0x65, 0x1c, 0xbf, 0xf9, 0x93,
	0xa5, 0x89, 0x8c, 0xa9, 0x46, 0x46, 0x36, 0x2a, 0x89, 0xd0, 0xda, 0x92, 0xef, 0x96, 0x49, 0x63,
	0xaa, 0x53, 0xc7, 0xa4, 0x49, 0x94, 0x3a, 0xe6, 0xc3, 0x9f, 0x9a, 0x01, 0x21, 0x31, 0xd5, 0x84,
	0x38, 0x8a, 0x90, 0x24, 0x22, 0xc4, 0x71, 0xdb, 0xff, 0xed, 0xb2, 0x97, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0xf7, 0x4d, 0x99, 0x45, 0x02, 0x00, 0x00,
}