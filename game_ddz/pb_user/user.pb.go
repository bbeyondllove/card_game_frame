// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package pb_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	BasicInfo
*/
package pb_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BasicInfo struct {
	Nickname string `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Sex      string `protobuf:"bytes,2,opt,name=sex" json:"sex,omitempty"`
	Imageid  int64  `protobuf:"varint,3,opt,name=imageid" json:"imageid,omitempty"`
}

func (m *BasicInfo) Reset()                    { *m = BasicInfo{} }
func (m *BasicInfo) String() string            { return proto.CompactTextString(m) }
func (*BasicInfo) ProtoMessage()               {}
func (*BasicInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*BasicInfo)(nil), "pb_user.BasicInfo")
}

var fileDescriptor0 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0x8a, 0x07, 0x71, 0x95, 0xac, 0xb9,
	0x38, 0x9d, 0x12, 0x8b, 0x33, 0x93, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0x04, 0xb8, 0x38, 0xf2, 0x32,
	0x93, 0xb3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x85, 0xb8, 0xb9, 0x98,
	0x8b, 0x53, 0x2b, 0x24, 0x98, 0xc0, 0x1c, 0x7e, 0x2e, 0xf6, 0xcc, 0xdc, 0xc4, 0xf4, 0xd4, 0xcc,
	0x14, 0x09, 0x66, 0xa0, 0x00, 0x73, 0x12, 0x1b, 0xd8, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5b, 0xad, 0x96, 0x2c, 0x5a, 0x00, 0x00, 0x00,
}
