// Code generated by protoc-gen-go.
// source: communication.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	communication.proto

It has these top-level messages:
	Action
*/
package messages

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

type Action_FaceDetectionState int32

const (
	Action_FACEDETECTION_UNCHANGED Action_FaceDetectionState = 0
	Action_FACEDETECTION_ENABLE    Action_FaceDetectionState = 1
	Action_FACEDETECTION_DISABLE   Action_FaceDetectionState = 2
)

var Action_FaceDetectionState_name = map[int32]string{
	0: "FACEDETECTION_UNCHANGED",
	1: "FACEDETECTION_ENABLE",
	2: "FACEDETECTION_DISABLE",
}
var Action_FaceDetectionState_value = map[string]int32{
	"FACEDETECTION_UNCHANGED": 0,
	"FACEDETECTION_ENABLE":    1,
	"FACEDETECTION_DISABLE":   2,
}

func (x Action_FaceDetectionState) String() string {
	return proto.EnumName(Action_FaceDetectionState_name, int32(x))
}
func (Action_FaceDetectionState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Action_RenderingMode int32

const (
	Action_RENDERINGMODE_UNCHANGED Action_RenderingMode = 0
	Action_RENDERINGMODE_NORMAL    Action_RenderingMode = 1
	Action_RENDERINGMODE_FUN       Action_RenderingMode = 2
)

var Action_RenderingMode_name = map[int32]string{
	0: "RENDERINGMODE_UNCHANGED",
	1: "RENDERINGMODE_NORMAL",
	2: "RENDERINGMODE_FUN",
}
var Action_RenderingMode_value = map[string]int32{
	"RENDERINGMODE_UNCHANGED": 0,
	"RENDERINGMODE_NORMAL":    1,
	"RENDERINGMODE_FUN":       2,
}

func (x Action_RenderingMode) String() string {
	return proto.EnumName(Action_RenderingMode_name, int32(x))
}
func (Action_RenderingMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Action struct {
	FaceDetection Action_FaceDetectionState `protobuf:"varint,1,opt,name=faceDetection,enum=messages.Action_FaceDetectionState" json:"faceDetection,omitempty"`
	RenderingMode Action_RenderingMode      `protobuf:"varint,2,opt,name=renderingMode,enum=messages.Action_RenderingMode" json:"renderingMode,omitempty"`
	Camera        int32                     `protobuf:"varint,3,opt,name=Camera,json=camera" json:"Camera,omitempty"`
	QuitServer    bool                      `protobuf:"varint,4,opt,name=QuitServer,json=quitServer" json:"QuitServer,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Action)(nil), "messages.Action")
	proto.RegisterEnum("messages.Action_FaceDetectionState", Action_FaceDetectionState_name, Action_FaceDetectionState_value)
	proto.RegisterEnum("messages.Action_RenderingMode", Action_RenderingMode_name, Action_RenderingMode_value)
}

func init() { proto.RegisterFile("communication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0x51, 0x4f, 0xb3, 0x30,
	0x14, 0x86, 0x3f, 0xd8, 0x27, 0x59, 0x4e, 0x82, 0xc1, 0xea, 0x14, 0x63, 0xb2, 0x18, 0xbc, 0xf1,
	0x8a, 0x0b, 0xfd, 0x05, 0x48, 0xcb, 0x24, 0x19, 0x25, 0x96, 0xed, 0xd2, 0x98, 0xca, 0xce, 0x16,
	0x2e, 0x00, 0x2d, 0x9d, 0xbf, 0xc5, 0x9f, 0x2b, 0xc3, 0x68, 0x56, 0x77, 0xd9, 0xf3, 0x3c, 0xef,
	0x79, 0x4f, 0x52, 0x38, 0x2d, 0xdb, 0xba, 0xde, 0x36, 0x55, 0x29, 0x75, 0xd5, 0x36, 0xe1, 0x9b,
	0x6a, 0x75, 0x4b, 0xc6, 0x35, 0x76, 0x9d, 0xdc, 0x60, 0x17, 0x7c, 0x8e, 0xc0, 0x89, 0xca, 0x1d,
	0x22, 0x29, 0xb8, 0x6b, 0x59, 0x22, 0x45, 0x8d, 0xc3, 0xc0, 0xb7, 0xae, 0xad, 0xdb, 0xe3, 0xbb,
	0x9b, 0xf0, 0x47, 0x0e, 0xbf, 0xc5, 0x30, 0xd9, 0xb7, 0x0a, 0x2d, 0x35, 0x0a, 0x33, 0x49, 0x28,
	0xb8, 0x0a, 0x9b, 0x15, 0xaa, 0xaa, 0xd9, 0x64, 0xed, 0x0a, 0x7d, 0x7b, 0x58, 0x35, 0x3d, 0x58,
	0x25, 0xf6, 0x2d, 0x61, 0x86, 0xc8, 0x39, 0x38, 0xb1, 0xac, 0x51, 0x49, 0x7f, 0xd4, 0xc7, 0x8f,
	0x84, 0x53, 0x0e, 0x2f, 0x32, 0x05, 0x78, 0xda, 0x56, 0xba, 0x40, 0xf5, 0x81, 0xca, 0xff, 0xdf,
	0xb3, 0xb1, 0x80, 0xf7, 0xdf, 0x49, 0xb0, 0x06, 0x72, 0x78, 0x22, 0xb9, 0x82, 0x8b, 0x24, 0x8a,
	0x19, 0x65, 0x0b, 0x16, 0x2f, 0xd2, 0x9c, 0xbf, 0x2c, 0x79, 0xfc, 0x18, 0xf1, 0x19, 0xa3, 0xde,
	0x3f, 0xe2, 0xc3, 0x99, 0x09, 0x19, 0x8f, 0x1e, 0xe6, 0xcc, 0xb3, 0xc8, 0x25, 0x4c, 0x4c, 0x42,
	0xd3, 0x62, 0x40, 0x76, 0xf0, 0x0c, 0xae, 0x71, 0xff, 0xae, 0x42, 0x30, 0x4e, 0x99, 0x48, 0xf9,
	0x2c, 0xcb, 0x29, 0xfb, 0x5b, 0x61, 0x42, 0x9e, 0x8b, 0x2c, 0x9a, 0xf7, 0x15, 0x13, 0x38, 0x31,
	0x49, 0xb2, 0xe4, 0x9e, 0xfd, 0xea, 0x0c, 0x7f, 0x75, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x4b, 0x42, 0xce, 0xc2, 0x01, 0x00, 0x00,
}
