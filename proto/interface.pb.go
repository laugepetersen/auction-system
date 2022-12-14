// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: proto/interface.proto

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/interface.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VoteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestHost        int32   `protobuf:"varint,1,opt,name=bestHost,proto3" json:"bestHost,omitempty"`
	BestHostLamport int32   `protobuf:"varint,2,opt,name=bestHostLamport,proto3" json:"bestHostLamport,omitempty"`
	Queue           []int32 `protobuf:"varint,3,rep,packed,name=queue,proto3" json:"queue,omitempty"`
}

func (x *VoteReq) Reset() {
	*x = VoteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteReq) ProtoMessage() {}

func (x *VoteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteReq.ProtoReflect.Descriptor instead.
func (*VoteReq) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{0}
}

func (x *VoteReq) GetBestHost() int32 {
	if x != nil {
		return x.BestHost
	}
	return 0
}

func (x *VoteReq) GetBestHostLamport() int32 {
	if x != nil {
		return x.BestHostLamport
	}
	return 0
}

func (x *VoteReq) GetQueue() []int32 {
	if x != nil {
		return x.Queue
	}
	return nil
}

type VoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack bool `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *VoteReply) Reset() {
	*x = VoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteReply) ProtoMessage() {}

func (x *VoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteReply.ProtoReflect.Descriptor instead.
func (*VoteReply) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{1}
}

func (x *VoteReply) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host    int32 `protobuf:"varint,1,opt,name=host,proto3" json:"host,omitempty"`
	Lamport int32 `protobuf:"varint,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetHost() int32 {
	if x != nil {
		return x.Host
	}
	return 0
}

func (x *Ping) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type BidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Host    int32 `protobuf:"varint,2,opt,name=host,proto3" json:"host,omitempty"`
	Lamport int32 `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *BidMessage) Reset() {
	*x = BidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidMessage) ProtoMessage() {}

func (x *BidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidMessage.ProtoReflect.Descriptor instead.
func (*BidMessage) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{3}
}

func (x *BidMessage) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidMessage) GetHost() int32 {
	if x != nil {
		return x.Host
	}
	return 0
}

func (x *BidMessage) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack     bool   `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Ack) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type ElectorAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lamport int32 `protobuf:"varint,1,opt,name=lamport,proto3" json:"lamport,omitempty"`
	Port    int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Answer  bool  `protobuf:"varint,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *ElectorAnswer) Reset() {
	*x = ElectorAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectorAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectorAnswer) ProtoMessage() {}

func (x *ElectorAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectorAnswer.ProtoReflect.Descriptor instead.
func (*ElectorAnswer) Descriptor() ([]byte, []int) {
	return file_proto_interface_proto_rawDescGZIP(), []int{5}
}

func (x *ElectorAnswer) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *ElectorAnswer) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ElectorAnswer) GetAnswer() bool {
	if x != nil {
		return x.Answer
	}
	return false
}

var File_proto_interface_proto protoreflect.FileDescriptor

var file_proto_interface_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x65, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73,
	0x74, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x62, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x61, 0x63, 0x6b, 0x22, 0x34, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x0a, 0x42, 0x69,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4b,
	0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x55, 0x0a, 0x0d, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c,
	0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x32, 0x9b, 0x02, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x10, 0x41, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x6b, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x63, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_interface_proto_rawDescOnce sync.Once
	file_proto_interface_proto_rawDescData = file_proto_interface_proto_rawDesc
)

func file_proto_interface_proto_rawDescGZIP() []byte {
	file_proto_interface_proto_rawDescOnce.Do(func() {
		file_proto_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_interface_proto_rawDescData)
	})
	return file_proto_interface_proto_rawDescData
}

var file_proto_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_interface_proto_goTypes = []interface{}{
	(*VoteReq)(nil),       // 0: proto.VoteReq
	(*VoteReply)(nil),     // 1: proto.VoteReply
	(*Ping)(nil),          // 2: proto.Ping
	(*BidMessage)(nil),    // 3: proto.BidMessage
	(*Ack)(nil),           // 4: proto.Ack
	(*ElectorAnswer)(nil), // 5: proto.ElectorAnswer
}
var file_proto_interface_proto_depIdxs = []int32{
	3, // 0: proto.AuctionService.Bid:input_type -> proto.BidMessage
	2, // 1: proto.AuctionService.GetResult:input_type -> proto.Ping
	2, // 2: proto.AuctionService.PingClient:input_type -> proto.Ping
	0, // 3: proto.AuctionService.AskForLeadership:input_type -> proto.VoteReq
	0, // 4: proto.AuctionService.SetNewLeader:input_type -> proto.VoteReq
	4, // 5: proto.AuctionService.Message:input_type -> proto.Ack
	4, // 6: proto.AuctionService.Bid:output_type -> proto.Ack
	4, // 7: proto.AuctionService.GetResult:output_type -> proto.Ack
	2, // 8: proto.AuctionService.PingClient:output_type -> proto.Ping
	1, // 9: proto.AuctionService.AskForLeadership:output_type -> proto.VoteReply
	1, // 10: proto.AuctionService.SetNewLeader:output_type -> proto.VoteReply
	4, // 11: proto.AuctionService.Message:output_type -> proto.Ack
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_interface_proto_init() }
func file_proto_interface_proto_init() {
	if File_proto_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectorAnswer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_interface_proto_goTypes,
		DependencyIndexes: file_proto_interface_proto_depIdxs,
		MessageInfos:      file_proto_interface_proto_msgTypes,
	}.Build()
	File_proto_interface_proto = out.File
	file_proto_interface_proto_rawDesc = nil
	file_proto_interface_proto_goTypes = nil
	file_proto_interface_proto_depIdxs = nil
}
