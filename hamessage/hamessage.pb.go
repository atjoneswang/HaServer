// Code generated by protoc-gen-go.
// source: hamessage.proto
// DO NOT EDIT!

/*
Package hamessage is a generated protocol buffer package.

It is generated from these files:
	hamessage.proto

It has these top-level messages:
	Hamessage
*/
package hamessage

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

type Hamessage struct {
	Event        string `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Profileimage string `protobuf:"bytes,3,opt,name=profileimage" json:"profileimage,omitempty"`
	Text         string `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
	Email        string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
}

func (m *Hamessage) Reset()                    { *m = Hamessage{} }
func (m *Hamessage) String() string            { return proto.CompactTextString(m) }
func (*Hamessage) ProtoMessage()               {}
func (*Hamessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Hamessage)(nil), "Hamessage")
}

func init() { proto.RegisterFile("hamessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0xcc, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x6a, 0x66, 0xe4, 0xe2,
	0xf4, 0x80, 0x89, 0x09, 0x89, 0x70, 0xb1, 0xa6, 0x96, 0xa5, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60,
	0x41, 0x30, 0x5b, 0x48, 0x89, 0x8b, 0xa7, 0xa0, 0x28, 0x3f, 0x2d, 0x33, 0x27, 0x35, 0x33, 0x37,
	0x31, 0x3d, 0x55, 0x82, 0x19, 0x2c, 0x87, 0x22, 0x06, 0xd2, 0x57, 0x92, 0x5a, 0x51, 0x22, 0xc1,
	0x02, 0xd1, 0x07, 0x62, 0x83, 0x6d, 0xc8, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x85, 0xda, 0x00, 0xe2,
	0x24, 0xb1, 0x81, 0x1d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x55, 0x8f, 0x2f, 0x9f,
	0x00, 0x00, 0x00,
}
