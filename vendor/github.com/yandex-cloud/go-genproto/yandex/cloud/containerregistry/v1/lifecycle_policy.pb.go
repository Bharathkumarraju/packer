// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/lifecycle_policy.proto

package containerregistry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type LifecyclePolicy_Status int32

const (
	LifecyclePolicy_STATUS_UNSPECIFIED LifecyclePolicy_Status = 0
	LifecyclePolicy_ACTIVE             LifecyclePolicy_Status = 1
	LifecyclePolicy_DISABLED           LifecyclePolicy_Status = 2
)

var LifecyclePolicy_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "ACTIVE",
	2: "DISABLED",
}

var LifecyclePolicy_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"ACTIVE":             1,
	"DISABLED":           2,
}

func (x LifecyclePolicy_Status) String() string {
	return proto.EnumName(LifecyclePolicy_Status_name, int32(x))
}

func (LifecyclePolicy_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{0, 0}
}

type LifecyclePolicy struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RepositoryId         string                 `protobuf:"bytes,3,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status               LifecyclePolicy_Status `protobuf:"varint,5,opt,name=status,proto3,enum=yandex.cloud.containerregistry.v1.LifecyclePolicy_Status" json:"status,omitempty"`
	CreatedAt            *timestamp.Timestamp   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Rules                []*LifecycleRule       `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LifecyclePolicy) Reset()         { *m = LifecyclePolicy{} }
func (m *LifecyclePolicy) String() string { return proto.CompactTextString(m) }
func (*LifecyclePolicy) ProtoMessage()    {}
func (*LifecyclePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{0}
}

func (m *LifecyclePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecyclePolicy.Unmarshal(m, b)
}
func (m *LifecyclePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecyclePolicy.Marshal(b, m, deterministic)
}
func (m *LifecyclePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecyclePolicy.Merge(m, src)
}
func (m *LifecyclePolicy) XXX_Size() int {
	return xxx_messageInfo_LifecyclePolicy.Size(m)
}
func (m *LifecyclePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecyclePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_LifecyclePolicy proto.InternalMessageInfo

func (m *LifecyclePolicy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LifecyclePolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LifecyclePolicy) GetRepositoryId() string {
	if m != nil {
		return m.RepositoryId
	}
	return ""
}

func (m *LifecyclePolicy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LifecyclePolicy) GetStatus() LifecyclePolicy_Status {
	if m != nil {
		return m.Status
	}
	return LifecyclePolicy_STATUS_UNSPECIFIED
}

func (m *LifecyclePolicy) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LifecyclePolicy) GetRules() []*LifecycleRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type LifecycleRule struct {
	Description          string             `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	ExpirePeriod         *duration.Duration `protobuf:"bytes,2,opt,name=expire_period,json=expirePeriod,proto3" json:"expire_period,omitempty"`
	TagRegexp            string             `protobuf:"bytes,3,opt,name=tag_regexp,json=tagRegexp,proto3" json:"tag_regexp,omitempty"`
	Untagged             bool               `protobuf:"varint,4,opt,name=untagged,proto3" json:"untagged,omitempty"`
	RetainedTop          int64              `protobuf:"varint,5,opt,name=retained_top,json=retainedTop,proto3" json:"retained_top,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LifecycleRule) Reset()         { *m = LifecycleRule{} }
func (m *LifecycleRule) String() string { return proto.CompactTextString(m) }
func (*LifecycleRule) ProtoMessage()    {}
func (*LifecycleRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{1}
}

func (m *LifecycleRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleRule.Unmarshal(m, b)
}
func (m *LifecycleRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleRule.Marshal(b, m, deterministic)
}
func (m *LifecycleRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleRule.Merge(m, src)
}
func (m *LifecycleRule) XXX_Size() int {
	return xxx_messageInfo_LifecycleRule.Size(m)
}
func (m *LifecycleRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleRule.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleRule proto.InternalMessageInfo

func (m *LifecycleRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LifecycleRule) GetExpirePeriod() *duration.Duration {
	if m != nil {
		return m.ExpirePeriod
	}
	return nil
}

func (m *LifecycleRule) GetTagRegexp() string {
	if m != nil {
		return m.TagRegexp
	}
	return ""
}

func (m *LifecycleRule) GetUntagged() bool {
	if m != nil {
		return m.Untagged
	}
	return false
}

func (m *LifecycleRule) GetRetainedTop() int64 {
	if m != nil {
		return m.RetainedTop
	}
	return 0
}

func init() {
	proto.RegisterEnum("yandex.cloud.containerregistry.v1.LifecyclePolicy_Status", LifecyclePolicy_Status_name, LifecyclePolicy_Status_value)
	proto.RegisterType((*LifecyclePolicy)(nil), "yandex.cloud.containerregistry.v1.LifecyclePolicy")
	proto.RegisterType((*LifecycleRule)(nil), "yandex.cloud.containerregistry.v1.LifecycleRule")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/lifecycle_policy.proto", fileDescriptor_400d7b415ccde714)
}

var fileDescriptor_400d7b415ccde714 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x8f, 0xd2, 0x40,
	0x14, 0xc5, 0x2d, 0x2c, 0xec, 0x32, 0xc0, 0x4a, 0xe6, 0xc1, 0x54, 0x12, 0x15, 0x31, 0x26, 0x44,
	0xb3, 0xed, 0x52, 0xff, 0xc4, 0x55, 0xd9, 0x04, 0x16, 0x36, 0x69, 0xb2, 0x31, 0x58, 0x58, 0x4d,
	0x7c, 0x69, 0x86, 0xce, 0xdd, 0xee, 0x24, 0xa5, 0x33, 0x99, 0x4e, 0x09, 0xbc, 0xf9, 0xec, 0x83,
	0x5f, 0x6b, 0xfd, 0x4a, 0xee, 0x93, 0x71, 0x0a, 0x0a, 0x8b, 0xc9, 0xc6, 0xb7, 0xce, 0xbd, 0xf7,
	0xdc, 0x9e, 0xfe, 0x4e, 0x07, 0xbd, 0x59, 0x90, 0x98, 0xc2, 0xdc, 0x0e, 0x22, 0x9e, 0x52, 0x3b,
	0xe0, 0xb1, 0x22, 0x2c, 0x06, 0x29, 0x21, 0x64, 0x89, 0x92, 0x0b, 0x7b, 0xd6, 0xb6, 0x23, 0x76,
	0x01, 0xc1, 0x22, 0x88, 0xc0, 0x17, 0x3c, 0x62, 0xc1, 0xc2, 0x12, 0x92, 0x2b, 0x8e, 0x1f, 0x67,
	0x4a, 0x4b, 0x2b, 0xad, 0x2d, 0xa5, 0x35, 0x6b, 0xd7, 0x1f, 0x6c, 0x2c, 0x9f, 0x91, 0x88, 0x51,
	0xa2, 0x18, 0x8f, 0xb3, 0x0d, 0xf5, 0x87, 0x21, 0xe7, 0x61, 0x04, 0xb6, 0x3e, 0x4d, 0xd2, 0x0b,
	0x9b, 0xa6, 0x72, 0xbd, 0xff, 0xe8, 0x66, 0x5f, 0xb1, 0x29, 0x24, 0x8a, 0x4c, 0x45, 0x36, 0xd0,
	0xfc, 0x9e, 0x47, 0x77, 0xcf, 0x56, 0xee, 0x86, 0xda, 0x1c, 0xde, 0x47, 0x39, 0x46, 0x4d, 0xa3,
	0x61, 0xb4, 0x4a, 0x5e, 0x8e, 0x51, 0x8c, 0xd1, 0x4e, 0x4c, 0xa6, 0x60, 0xe6, 0x74, 0x45, 0x3f,
	0xe3, 0x27, 0xa8, 0x2a, 0x41, 0xf0, 0x84, 0x29, 0x2e, 0x17, 0x3e, 0xa3, 0x66, 0x5e, 0x37, 0x2b,
	0x7f, 0x8b, 0x2e, 0xc5, 0x0d, 0x54, 0xa6, 0x90, 0x04, 0x92, 0x89, 0xdf, 0x96, 0xcc, 0x1d, 0x3d,
	0xb2, 0x5e, 0xc2, 0x1f, 0x51, 0x31, 0x51, 0x44, 0xa5, 0x89, 0x59, 0x68, 0x18, 0xad, 0x7d, 0xe7,
	0xc8, 0xba, 0x15, 0x89, 0x75, 0xc3, 0xae, 0x35, 0xd2, 0x0b, 0xbc, 0xe5, 0x22, 0x7c, 0x84, 0x50,
	0x20, 0x81, 0x28, 0xa0, 0x3e, 0x51, 0x66, 0xb1, 0x61, 0xb4, 0xca, 0x4e, 0xdd, 0xca, 0x38, 0x58,
	0x2b, 0x0e, 0xd6, 0x78, 0xc5, 0xc1, 0x2b, 0x2d, 0xa7, 0xbb, 0x0a, 0x9f, 0xa2, 0x82, 0x4c, 0x23,
	0x48, 0xcc, 0xdd, 0x46, 0xbe, 0x55, 0x76, 0x0e, 0xff, 0xc7, 0x8c, 0x97, 0x46, 0xe0, 0x65, 0xf2,
	0xe6, 0x5b, 0x54, 0xcc, 0x4c, 0xe1, 0x7b, 0x08, 0x8f, 0xc6, 0xdd, 0xf1, 0xf9, 0xc8, 0x3f, 0xff,
	0x30, 0x1a, 0x0e, 0x4e, 0xdc, 0x53, 0x77, 0xd0, 0xaf, 0xdd, 0xc1, 0x08, 0x15, 0xbb, 0x27, 0x63,
	0xf7, 0xd3, 0xa0, 0x66, 0xe0, 0x0a, 0xda, 0xeb, 0xbb, 0xa3, 0x6e, 0xef, 0x6c, 0xd0, 0xaf, 0xe5,
	0x9a, 0x3f, 0x0d, 0x54, 0xdd, 0x58, 0x8a, 0x9f, 0x6f, 0x52, 0xd4, 0xb9, 0xf4, 0x4a, 0xdf, 0x7e,
	0xb4, 0x0b, 0xef, 0x3b, 0xce, 0xab, 0xd7, 0x9b, 0x40, 0x5d, 0x54, 0x85, 0xb9, 0x60, 0x12, 0x7c,
	0x01, 0x92, 0x71, 0xaa, 0x43, 0x2b, 0x3b, 0xf7, 0xb7, 0x00, 0xf4, 0x97, 0x3f, 0x4a, 0xaf, 0x74,
	0x7d, 0xd5, 0x2e, 0x1c, 0x77, 0x9c, 0x97, 0x97, 0x5e, 0x25, 0x93, 0x0e, 0xb5, 0x12, 0xb7, 0x10,
	0x52, 0x24, 0xf4, 0x25, 0x84, 0x30, 0x17, 0x59, 0xbe, 0xeb, 0xaf, 0x2d, 0x29, 0x12, 0x7a, 0xba,
	0x87, 0xeb, 0x68, 0x2f, 0x8d, 0x15, 0x09, 0x43, 0xa0, 0x3a, 0xe4, 0x3d, 0xef, 0xcf, 0x19, 0x3f,
	0x43, 0x15, 0x09, 0x9a, 0x1b, 0xf5, 0x15, 0x17, 0x3a, 0xe7, 0x7c, 0x6f, 0xf7, 0xfa, 0xaa, 0x9d,
	0x3f, 0xee, 0x1c, 0x7a, 0xe5, 0x55, 0x73, 0xcc, 0x45, 0xef, 0xab, 0x81, 0x9e, 0x6e, 0x20, 0x27,
	0x82, 0xfd, 0x13, 0xfb, 0x97, 0xcf, 0x21, 0x53, 0x97, 0xe9, 0xc4, 0x0a, 0xf8, 0xd4, 0xce, 0x14,
	0x07, 0xd9, 0x0d, 0x09, 0xf9, 0x41, 0x08, 0xb1, 0xfe, 0x4a, 0xfb, 0xd6, 0x7b, 0xf9, 0x6e, 0xab,
	0x38, 0x29, 0x6a, 0xe9, 0x8b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0xa1, 0xf6, 0x5d, 0xd5,
	0x03, 0x00, 0x00,
}