// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	bid.proto
	capabilities.proto
	hub.proto
	insonmnia.proto
	miner.proto

It has these top-level messages:
	Bid
	Capabilities
	CPUDevice
	RAMDevice
	GPUDevice
	ListRequest
	ListReply
	HubInfoRequest
	TaskRequirements
	HubStartTaskRequest
	HubStartTaskReply
	HubStatusMapRequest
	HubStatusRequest
	HubStatusReply
	DealRequest
	DealReply
	PingRequest
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	StopTaskRequest
	StopTaskReply
	TaskStatusRequest
	TaskStatusReply
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	TaskResourceRequirements
	Timestamp
	MinerInfoRequest
	MinerHandshakeRequest
	MinerHandshakeReply
	MinerStartRequest
	MinerStartReply
	MinerStatusMapRequest
*/
package sonm

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

type OrderType int32

const (
	OrderType_BID OrderType = 0
	OrderType_ASK OrderType = 1
)

var OrderType_name = map[int32]string{
	0: "BID",
	1: "ASK",
}
var OrderType_value = map[string]int32{
	"BID": 0,
	"ASK": 1,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Bid struct {
	// Virtual computer start of life (hour grained).
	StartTime *Timestamp `protobuf:"bytes,1,opt,name=startTime" json:"startTime,omitempty"`
	// Virtual computer end of life (hour grained).
	EndTime *Timestamp `protobuf:"bytes,2,opt,name=endTime" json:"endTime,omitempty"`
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	RatingBuyer float32 `protobuf:"fixed32,3,opt,name=ratingBuyer" json:"ratingBuyer,omitempty"`
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	RatingSupplier float32 `protobuf:"fixed32,4,opt,name=ratingSupplier" json:"ratingSupplier,omitempty"`
	// CPU core count.
	CpuQtty int64 `protobuf:"varint,5,opt,name=cpuQtty" json:"cpuQtty,omitempty"`
	// CPU type.
	CpuType int64 `protobuf:"varint,6,opt,name=cpuType" json:"cpuType,omitempty"`
	// Virtual computer RAM, in GB
	RamQtty int64 `protobuf:"varint,7,opt,name=ramQtty" json:"ramQtty,omitempty"`
	// Storage size, in GB.
	StorageQtty int64 `protobuf:"varint,8,opt,name=storageQtty" json:"storageQtty,omitempty"`
	// Storage type.
	StorageType int64 `protobuf:"varint,9,opt,name=storageType" json:"storageType,omitempty"`
	// Inbound or outbound traffic (the higher value), in GB.
	NetworkQtty int64 `protobuf:"varint,10,opt,name=networkQtty" json:"networkQtty,omitempty"`
	// Network latency in ms, the lower is better.
	NetworkLat int64 `protobuf:"varint,11,opt,name=networkLat" json:"networkLat,omitempty"`
	// GPU count.
	GpuQtty int64 `protobuf:"varint,12,opt,name=gpu_qtty,json=gpuQtty" json:"gpu_qtty,omitempty"`
	// GPU manufacturer.
	GpuManufact string `protobuf:"bytes,13,opt,name=gpuManufact" json:"gpuManufact,omitempty"`
	// GPU type.
	GpuType int64 `protobuf:"varint,14,opt,name=gpuType" json:"gpuType,omitempty"`
	// GPU RAM, in GB.
	GpuRam int64 `protobuf:"varint,15,opt,name=gpuRam" json:"gpuRam,omitempty"`
	// Order price.
	Price int64 `protobuf:"varint,16,opt,name=price" json:"price,omitempty"`
	// Order type.
	OrderType OrderType `protobuf:"varint,17,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
}

func (m *Bid) Reset()                    { *m = Bid{} }
func (m *Bid) String() string            { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()               {}
func (*Bid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bid) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Bid) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Bid) GetRatingBuyer() float32 {
	if m != nil {
		return m.RatingBuyer
	}
	return 0
}

func (m *Bid) GetRatingSupplier() float32 {
	if m != nil {
		return m.RatingSupplier
	}
	return 0
}

func (m *Bid) GetCpuQtty() int64 {
	if m != nil {
		return m.CpuQtty
	}
	return 0
}

func (m *Bid) GetCpuType() int64 {
	if m != nil {
		return m.CpuType
	}
	return 0
}

func (m *Bid) GetRamQtty() int64 {
	if m != nil {
		return m.RamQtty
	}
	return 0
}

func (m *Bid) GetStorageQtty() int64 {
	if m != nil {
		return m.StorageQtty
	}
	return 0
}

func (m *Bid) GetStorageType() int64 {
	if m != nil {
		return m.StorageType
	}
	return 0
}

func (m *Bid) GetNetworkQtty() int64 {
	if m != nil {
		return m.NetworkQtty
	}
	return 0
}

func (m *Bid) GetNetworkLat() int64 {
	if m != nil {
		return m.NetworkLat
	}
	return 0
}

func (m *Bid) GetGpuQtty() int64 {
	if m != nil {
		return m.GpuQtty
	}
	return 0
}

func (m *Bid) GetGpuManufact() string {
	if m != nil {
		return m.GpuManufact
	}
	return ""
}

func (m *Bid) GetGpuType() int64 {
	if m != nil {
		return m.GpuType
	}
	return 0
}

func (m *Bid) GetGpuRam() int64 {
	if m != nil {
		return m.GpuRam
	}
	return 0
}

func (m *Bid) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Bid) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_BID
}

func init() {
	proto.RegisterType((*Bid)(nil), "sonm.Bid")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4a, 0xeb, 0x40,
	0x10, 0x86, 0x4f, 0x9a, 0xb6, 0x69, 0x26, 0xe7, 0xb4, 0x3d, 0x8b, 0xc8, 0x2a, 0x28, 0x8b, 0x17,
	0x12, 0x05, 0x7b, 0xa1, 0x4f, 0x60, 0xf1, 0x46, 0x54, 0xc4, 0xb4, 0xf7, 0xb2, 0x6d, 0xd6, 0xb0,
	0x68, 0x92, 0x75, 0xb3, 0x8b, 0xf4, 0xb1, 0x7c, 0x43, 0xd9, 0x49, 0xd2, 0x46, 0xf1, 0x2a, 0xf9,
	0xbf, 0xff, 0x9f, 0x61, 0x86, 0x59, 0x08, 0x57, 0x32, 0x9d, 0x29, 0x5d, 0x9a, 0x92, 0xf4, 0xab,
	0xb2, 0xc8, 0x0f, 0x27, 0xb2, 0x70, 0xdf, 0x42, 0xf2, 0x1a, 0x9f, 0x7c, 0xf6, 0xc1, 0x9f, 0xcb,
	0x94, 0x5c, 0x40, 0x58, 0x19, 0xae, 0xcd, 0x52, 0xe6, 0x82, 0x7a, 0xcc, 0x8b, 0xa3, 0xcb, 0xc9,
	0xcc, 0x45, 0x67, 0x8e, 0x54, 0x86, 0xe7, 0x2a, 0xd9, 0x25, 0xc8, 0x19, 0x04, 0xa2, 0x48, 0x31,
	0xdc, 0xfb, 0x3d, 0xdc, 0xfa, 0x84, 0x41, 0xa4, 0xb9, 0x91, 0x45, 0x36, 0xb7, 0x1b, 0xa1, 0xa9,
	0xcf, 0xbc, 0xb8, 0x97, 0x74, 0x11, 0x39, 0x85, 0x71, 0x2d, 0x17, 0x56, 0xa9, 0x37, 0x29, 0x34,
	0xed, 0x63, 0xe8, 0x07, 0x25, 0x14, 0x82, 0xb5, 0xb2, 0x4f, 0xc6, 0x6c, 0xe8, 0x80, 0x79, 0xb1,
	0x9f, 0xb4, 0xb2, 0x71, 0x96, 0x1b, 0x25, 0xe8, 0x70, 0xeb, 0x38, 0xe9, 0x1c, 0xcd, 0x73, 0xac,
	0x09, 0x6a, 0xa7, 0x91, 0x6e, 0xae, 0xca, 0x94, 0x9a, 0x67, 0x02, 0xdd, 0x11, 0xba, 0x5d, 0xd4,
	0x49, 0x60, 0xe7, 0xf0, 0x5b, 0x02, 0xbb, 0x33, 0x88, 0x0a, 0x61, 0x3e, 0x4a, 0xfd, 0x8a, 0x3d,
	0xa0, 0x4e, 0x74, 0x10, 0x39, 0x06, 0x68, 0xe4, 0x3d, 0x37, 0x34, 0xc2, 0x40, 0x87, 0x90, 0x03,
	0x18, 0x65, 0xca, 0x3e, 0xbf, 0xbb, 0xf2, 0xbf, 0xf5, 0x80, 0x59, 0xb3, 0x14, 0x83, 0x28, 0x53,
	0xf6, 0x81, 0x17, 0xf6, 0x85, 0xaf, 0x0d, 0xfd, 0xc7, 0xbc, 0x38, 0x4c, 0xba, 0xc8, 0x2d, 0x97,
	0x35, 0x6b, 0x8f, 0xb7, 0xb5, 0x38, 0xd8, 0x3e, 0x0c, 0x33, 0x65, 0x13, 0x9e, 0xd3, 0x09, 0x1a,
	0x8d, 0x22, 0x7b, 0x30, 0x50, 0x5a, 0xae, 0x05, 0x9d, 0x22, 0xae, 0x85, 0x3b, 0x7e, 0xa9, 0x53,
	0xa1, 0xb1, 0xd3, 0x7f, 0xe6, 0xc5, 0xe3, 0xf6, 0x9e, 0x8f, 0x2d, 0x4e, 0x76, 0x89, 0xf3, 0x23,
	0x08, 0xb7, 0x9c, 0x04, 0xe0, 0xcf, 0x6f, 0x6f, 0xa6, 0x7f, 0xdc, 0xcf, 0xf5, 0xe2, 0x6e, 0xea,
	0xad, 0x86, 0xf8, 0xb2, 0xae, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xcf, 0x3c, 0xad, 0x7d,
	0x02, 0x00, 0x00,
}
