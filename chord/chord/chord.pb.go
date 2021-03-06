// Code generated by protoc-gen-go.
// source: chord.proto
// DO NOT EDIT!

/*
Package chord is a generated protocol buffer package.

It is generated from these files:
	chord.proto

It has these top-level messages:
	Empty
	KeyValueReqMsg
	TransferReqMsg
	KeyValueReplyMsg
	RemoteNodeMsg
	IdReplyMsg
	UpdateReqMsg
	NotifyReqMsg
	RemoteQueryMsg
	RpcOkayMsg
*/
package chord

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type KeyValueReqMsg struct {
	NodeId []byte `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	Value  string `protobuf:"bytes,3,opt,name=Value" json:"Value,omitempty"`
}

func (m *KeyValueReqMsg) Reset()                    { *m = KeyValueReqMsg{} }
func (m *KeyValueReqMsg) String() string            { return proto.CompactTextString(m) }
func (*KeyValueReqMsg) ProtoMessage()               {}
func (*KeyValueReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyValueReqMsg) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *KeyValueReqMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValueReqMsg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TransferReqMsg struct {
	NodeId   []byte `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	FromId   []byte `protobuf:"bytes,2,opt,name=FromId,proto3" json:"FromId,omitempty"`
	FromAddr string `protobuf:"bytes,3,opt,name=FromAddr" json:"FromAddr,omitempty"`
	PredId   []byte `protobuf:"bytes,4,opt,name=PredId,proto3" json:"PredId,omitempty"`
}

func (m *TransferReqMsg) Reset()                    { *m = TransferReqMsg{} }
func (m *TransferReqMsg) String() string            { return proto.CompactTextString(m) }
func (*TransferReqMsg) ProtoMessage()               {}
func (*TransferReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransferReqMsg) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *TransferReqMsg) GetFromId() []byte {
	if m != nil {
		return m.FromId
	}
	return nil
}

func (m *TransferReqMsg) GetFromAddr() string {
	if m != nil {
		return m.FromAddr
	}
	return ""
}

func (m *TransferReqMsg) GetPredId() []byte {
	if m != nil {
		return m.PredId
	}
	return nil
}

type KeyValueReplyMsg struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *KeyValueReplyMsg) Reset()                    { *m = KeyValueReplyMsg{} }
func (m *KeyValueReplyMsg) String() string            { return proto.CompactTextString(m) }
func (*KeyValueReplyMsg) ProtoMessage()               {}
func (*KeyValueReplyMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeyValueReplyMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValueReplyMsg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RemoteNodeMsg struct {
	Id   []byte `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=Addr" json:"Addr,omitempty"`
}

func (m *RemoteNodeMsg) Reset()                    { *m = RemoteNodeMsg{} }
func (m *RemoteNodeMsg) String() string            { return proto.CompactTextString(m) }
func (*RemoteNodeMsg) ProtoMessage()               {}
func (*RemoteNodeMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoteNodeMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RemoteNodeMsg) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type IdReplyMsg struct {
	Id    []byte `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=Addr" json:"Addr,omitempty"`
	Valid bool   `protobuf:"varint,3,opt,name=Valid" json:"Valid,omitempty"`
}

func (m *IdReplyMsg) Reset()                    { *m = IdReplyMsg{} }
func (m *IdReplyMsg) String() string            { return proto.CompactTextString(m) }
func (*IdReplyMsg) ProtoMessage()               {}
func (*IdReplyMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IdReplyMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *IdReplyMsg) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *IdReplyMsg) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type UpdateReqMsg struct {
	FromId     []byte `protobuf:"bytes,1,opt,name=FromId,proto3" json:"FromId,omitempty"`
	UpdateId   []byte `protobuf:"bytes,2,opt,name=UpdateId,proto3" json:"UpdateId,omitempty"`
	UpdateAddr string `protobuf:"bytes,3,opt,name=UpdateAddr" json:"UpdateAddr,omitempty"`
}

func (m *UpdateReqMsg) Reset()                    { *m = UpdateReqMsg{} }
func (m *UpdateReqMsg) String() string            { return proto.CompactTextString(m) }
func (*UpdateReqMsg) ProtoMessage()               {}
func (*UpdateReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateReqMsg) GetFromId() []byte {
	if m != nil {
		return m.FromId
	}
	return nil
}

func (m *UpdateReqMsg) GetUpdateId() []byte {
	if m != nil {
		return m.UpdateId
	}
	return nil
}

func (m *UpdateReqMsg) GetUpdateAddr() string {
	if m != nil {
		return m.UpdateAddr
	}
	return ""
}

type NotifyReqMsg struct {
	NodeId     []byte `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	NodeAddr   string `protobuf:"bytes,2,opt,name=NodeAddr" json:"NodeAddr,omitempty"`
	UpdateId   []byte `protobuf:"bytes,3,opt,name=UpdateId,proto3" json:"UpdateId,omitempty"`
	UpdateAddr string `protobuf:"bytes,4,opt,name=UpdateAddr" json:"UpdateAddr,omitempty"`
}

func (m *NotifyReqMsg) Reset()                    { *m = NotifyReqMsg{} }
func (m *NotifyReqMsg) String() string            { return proto.CompactTextString(m) }
func (*NotifyReqMsg) ProtoMessage()               {}
func (*NotifyReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NotifyReqMsg) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NotifyReqMsg) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

func (m *NotifyReqMsg) GetUpdateId() []byte {
	if m != nil {
		return m.UpdateId
	}
	return nil
}

func (m *NotifyReqMsg) GetUpdateAddr() string {
	if m != nil {
		return m.UpdateAddr
	}
	return ""
}

type RemoteQueryMsg struct {
	FromId []byte `protobuf:"bytes,1,opt,name=FromId,proto3" json:"FromId,omitempty"`
	Id     []byte `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *RemoteQueryMsg) Reset()                    { *m = RemoteQueryMsg{} }
func (m *RemoteQueryMsg) String() string            { return proto.CompactTextString(m) }
func (*RemoteQueryMsg) ProtoMessage()               {}
func (*RemoteQueryMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoteQueryMsg) GetFromId() []byte {
	if m != nil {
		return m.FromId
	}
	return nil
}

func (m *RemoteQueryMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type RpcOkayMsg struct {
	Ok bool `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
}

func (m *RpcOkayMsg) Reset()                    { *m = RpcOkayMsg{} }
func (m *RpcOkayMsg) String() string            { return proto.CompactTextString(m) }
func (*RpcOkayMsg) ProtoMessage()               {}
func (*RpcOkayMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RpcOkayMsg) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "chord.Empty")
	proto.RegisterType((*KeyValueReqMsg)(nil), "chord.KeyValueReqMsg")
	proto.RegisterType((*TransferReqMsg)(nil), "chord.TransferReqMsg")
	proto.RegisterType((*KeyValueReplyMsg)(nil), "chord.KeyValueReplyMsg")
	proto.RegisterType((*RemoteNodeMsg)(nil), "chord.RemoteNodeMsg")
	proto.RegisterType((*IdReplyMsg)(nil), "chord.IdReplyMsg")
	proto.RegisterType((*UpdateReqMsg)(nil), "chord.UpdateReqMsg")
	proto.RegisterType((*NotifyReqMsg)(nil), "chord.NotifyReqMsg")
	proto.RegisterType((*RemoteQueryMsg)(nil), "chord.RemoteQueryMsg")
	proto.RegisterType((*RpcOkayMsg)(nil), "chord.RpcOkayMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChordRPC service

type ChordRPCClient interface {
	GetPredecessorIdCaller(ctx context.Context, in *RemoteNodeMsg, opts ...grpc.CallOption) (*IdReplyMsg, error)
	GetSuccessorIdCaller(ctx context.Context, in *RemoteNodeMsg, opts ...grpc.CallOption) (*IdReplyMsg, error)
	SetPredecessorIdCaller(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error)
	SetSuccessorIdCaller(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error)
	NotifyCaller(ctx context.Context, in *NotifyReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error)
	FindSuccessorCaller(ctx context.Context, in *RemoteQueryMsg, opts ...grpc.CallOption) (*IdReplyMsg, error)
	ClosestPrecedingFingerCaller(ctx context.Context, in *RemoteQueryMsg, opts ...grpc.CallOption) (*IdReplyMsg, error)
	GetCaller(ctx context.Context, in *KeyValueReqMsg, opts ...grpc.CallOption) (*KeyValueReplyMsg, error)
	PutCaller(ctx context.Context, in *KeyValueReqMsg, opts ...grpc.CallOption) (*KeyValueReplyMsg, error)
	TransferKeysCaller(ctx context.Context, in *TransferReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error)
}

type chordRPCClient struct {
	cc *grpc.ClientConn
}

func NewChordRPCClient(cc *grpc.ClientConn) ChordRPCClient {
	return &chordRPCClient{cc}
}

func (c *chordRPCClient) GetPredecessorIdCaller(ctx context.Context, in *RemoteNodeMsg, opts ...grpc.CallOption) (*IdReplyMsg, error) {
	out := new(IdReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/GetPredecessorIdCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) GetSuccessorIdCaller(ctx context.Context, in *RemoteNodeMsg, opts ...grpc.CallOption) (*IdReplyMsg, error) {
	out := new(IdReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/GetSuccessorIdCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) SetPredecessorIdCaller(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error) {
	out := new(RpcOkayMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/SetPredecessorIdCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) SetSuccessorIdCaller(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error) {
	out := new(RpcOkayMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/SetSuccessorIdCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) NotifyCaller(ctx context.Context, in *NotifyReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error) {
	out := new(RpcOkayMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/NotifyCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) FindSuccessorCaller(ctx context.Context, in *RemoteQueryMsg, opts ...grpc.CallOption) (*IdReplyMsg, error) {
	out := new(IdReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/FindSuccessorCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) ClosestPrecedingFingerCaller(ctx context.Context, in *RemoteQueryMsg, opts ...grpc.CallOption) (*IdReplyMsg, error) {
	out := new(IdReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/ClosestPrecedingFingerCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) GetCaller(ctx context.Context, in *KeyValueReqMsg, opts ...grpc.CallOption) (*KeyValueReplyMsg, error) {
	out := new(KeyValueReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/GetCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) PutCaller(ctx context.Context, in *KeyValueReqMsg, opts ...grpc.CallOption) (*KeyValueReplyMsg, error) {
	out := new(KeyValueReplyMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/PutCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordRPCClient) TransferKeysCaller(ctx context.Context, in *TransferReqMsg, opts ...grpc.CallOption) (*RpcOkayMsg, error) {
	out := new(RpcOkayMsg)
	err := grpc.Invoke(ctx, "/chord.ChordRPC/TransferKeysCaller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChordRPC service

type ChordRPCServer interface {
	GetPredecessorIdCaller(context.Context, *RemoteNodeMsg) (*IdReplyMsg, error)
	GetSuccessorIdCaller(context.Context, *RemoteNodeMsg) (*IdReplyMsg, error)
	SetPredecessorIdCaller(context.Context, *UpdateReqMsg) (*RpcOkayMsg, error)
	SetSuccessorIdCaller(context.Context, *UpdateReqMsg) (*RpcOkayMsg, error)
	NotifyCaller(context.Context, *NotifyReqMsg) (*RpcOkayMsg, error)
	FindSuccessorCaller(context.Context, *RemoteQueryMsg) (*IdReplyMsg, error)
	ClosestPrecedingFingerCaller(context.Context, *RemoteQueryMsg) (*IdReplyMsg, error)
	GetCaller(context.Context, *KeyValueReqMsg) (*KeyValueReplyMsg, error)
	PutCaller(context.Context, *KeyValueReqMsg) (*KeyValueReplyMsg, error)
	TransferKeysCaller(context.Context, *TransferReqMsg) (*RpcOkayMsg, error)
}

func RegisterChordRPCServer(s *grpc.Server, srv ChordRPCServer) {
	s.RegisterService(&_ChordRPC_serviceDesc, srv)
}

func _ChordRPC_GetPredecessorIdCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteNodeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).GetPredecessorIdCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/GetPredecessorIdCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).GetPredecessorIdCaller(ctx, req.(*RemoteNodeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_GetSuccessorIdCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteNodeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).GetSuccessorIdCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/GetSuccessorIdCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).GetSuccessorIdCaller(ctx, req.(*RemoteNodeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_SetPredecessorIdCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).SetPredecessorIdCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/SetPredecessorIdCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).SetPredecessorIdCaller(ctx, req.(*UpdateReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_SetSuccessorIdCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).SetSuccessorIdCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/SetSuccessorIdCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).SetSuccessorIdCaller(ctx, req.(*UpdateReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_NotifyCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).NotifyCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/NotifyCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).NotifyCaller(ctx, req.(*NotifyReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_FindSuccessorCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).FindSuccessorCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/FindSuccessorCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).FindSuccessorCaller(ctx, req.(*RemoteQueryMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_ClosestPrecedingFingerCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).ClosestPrecedingFingerCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/ClosestPrecedingFingerCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).ClosestPrecedingFingerCaller(ctx, req.(*RemoteQueryMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_GetCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).GetCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/GetCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).GetCaller(ctx, req.(*KeyValueReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_PutCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).PutCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/PutCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).PutCaller(ctx, req.(*KeyValueReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordRPC_TransferKeysCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordRPCServer).TransferKeysCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ChordRPC/TransferKeysCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordRPCServer).TransferKeysCaller(ctx, req.(*TransferReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChordRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chord.ChordRPC",
	HandlerType: (*ChordRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPredecessorIdCaller",
			Handler:    _ChordRPC_GetPredecessorIdCaller_Handler,
		},
		{
			MethodName: "GetSuccessorIdCaller",
			Handler:    _ChordRPC_GetSuccessorIdCaller_Handler,
		},
		{
			MethodName: "SetPredecessorIdCaller",
			Handler:    _ChordRPC_SetPredecessorIdCaller_Handler,
		},
		{
			MethodName: "SetSuccessorIdCaller",
			Handler:    _ChordRPC_SetSuccessorIdCaller_Handler,
		},
		{
			MethodName: "NotifyCaller",
			Handler:    _ChordRPC_NotifyCaller_Handler,
		},
		{
			MethodName: "FindSuccessorCaller",
			Handler:    _ChordRPC_FindSuccessorCaller_Handler,
		},
		{
			MethodName: "ClosestPrecedingFingerCaller",
			Handler:    _ChordRPC_ClosestPrecedingFingerCaller_Handler,
		},
		{
			MethodName: "GetCaller",
			Handler:    _ChordRPC_GetCaller_Handler,
		},
		{
			MethodName: "PutCaller",
			Handler:    _ChordRPC_PutCaller_Handler,
		},
		{
			MethodName: "TransferKeysCaller",
			Handler:    _ChordRPC_TransferKeysCaller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord.proto",
}

func init() { proto.RegisterFile("chord.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x9c, 0xa4, 0xa4, 0x43, 0x88, 0xca, 0x36, 0x84, 0x28, 0xaa, 0x10, 0xf2, 0x89, 0x53,
	0x0f, 0x54, 0x42, 0x08, 0xa4, 0x96, 0x2a, 0x22, 0x10, 0x21, 0xda, 0xb0, 0x01, 0xee, 0xae, 0x77,
	0x1a, 0xa2, 0x3a, 0x59, 0xb3, 0x76, 0x0e, 0xbe, 0xf0, 0x0f, 0xf8, 0xcf, 0xec, 0x97, 0xb7, 0xeb,
	0xd4, 0x6a, 0x4b, 0x6f, 0xf3, 0x26, 0xfb, 0x66, 0xde, 0xbc, 0x71, 0x06, 0x1e, 0xc7, 0xbf, 0xb8,
	0x60, 0x87, 0xa9, 0xe0, 0x39, 0x27, 0x6d, 0x0d, 0xc2, 0x47, 0xd0, 0xfe, 0xb8, 0x4a, 0xf3, 0x22,
	0x9c, 0x41, 0xef, 0x0b, 0x16, 0x3f, 0xa3, 0x64, 0x83, 0x14, 0x7f, 0x7f, 0xcd, 0x16, 0x64, 0x00,
	0x3b, 0x67, 0x9c, 0xe1, 0x94, 0x0d, 0x1b, 0x2f, 0x1b, 0xaf, 0xba, 0xd4, 0x22, 0xb2, 0x07, 0x4d,
	0xf9, 0x72, 0x18, 0xc8, 0xe4, 0x2e, 0x55, 0x21, 0xe9, 0x43, 0x5b, 0x13, 0x87, 0x4d, 0x9d, 0x33,
	0x20, 0xcc, 0xa1, 0xf7, 0x5d, 0x44, 0xeb, 0xec, 0x12, 0xc5, 0x1d, 0x15, 0x65, 0x7e, 0x22, 0xf8,
	0x4a, 0xe6, 0x03, 0x93, 0x37, 0x88, 0x8c, 0xa0, 0xa3, 0xa2, 0x53, 0xc6, 0x84, 0x2d, 0xed, 0xb0,
	0xe2, 0xcc, 0x04, 0x32, 0xc9, 0x69, 0x19, 0x8e, 0x41, 0xe1, 0x3b, 0xd8, 0xbb, 0x9e, 0x23, 0x4d,
	0x0a, 0xd5, 0xd7, 0x2a, 0x6e, 0xd4, 0x28, 0x0e, 0x7c, 0xc5, 0x47, 0xf0, 0x84, 0xe2, 0x8a, 0xe7,
	0xa8, 0x74, 0x29, 0x62, 0x0f, 0x02, 0x27, 0x56, 0x46, 0x84, 0x40, 0x4b, 0x8b, 0x31, 0x2c, 0x1d,
	0x87, 0x13, 0x80, 0x29, 0x73, 0xad, 0xee, 0xc1, 0xb0, 0xcd, 0x97, 0x4c, 0xcf, 0xd4, 0xa1, 0x06,
	0x84, 0x17, 0xd0, 0xfd, 0x91, 0xb2, 0x28, 0xf7, 0xec, 0xb7, 0xa6, 0x34, 0xb6, 0x4d, 0x31, 0xef,
	0x9c, 0x5d, 0x0e, 0x93, 0x17, 0x00, 0x26, 0xf6, 0x2c, 0xf3, 0x32, 0xe1, 0x1f, 0xe8, 0x9e, 0xf1,
	0x7c, 0x79, 0x59, 0xdc, 0xb1, 0x10, 0xd9, 0x43, 0x45, 0x9e, 0x72, 0x87, 0x2b, 0xfd, 0x9b, 0xb7,
	0xf6, 0x6f, 0xdd, 0xe8, 0xff, 0x16, 0x7a, 0xc6, 0xe0, 0x6f, 0x1b, 0x14, 0xc5, 0x6d, 0x53, 0x1a,
	0x1f, 0x83, 0xd2, 0xc7, 0xf0, 0x00, 0x80, 0xa6, 0xf1, 0xf9, 0x55, 0x54, 0xba, 0x7c, 0x7e, 0xa5,
	0x19, 0x1d, 0x2a, 0xa3, 0xd7, 0x7f, 0xdb, 0xd0, 0x19, 0xab, 0xef, 0x99, 0xce, 0xc6, 0xe4, 0x14,
	0x06, 0x9f, 0x30, 0x57, 0x9f, 0x03, 0xc6, 0x98, 0x65, 0x5c, 0x4c, 0xd9, 0x38, 0x4a, 0x12, 0x94,
	0xc6, 0x1f, 0x9a, 0x7f, 0x40, 0x65, 0xc9, 0xa3, 0xa7, 0x36, 0xeb, 0x6d, 0xf1, 0x04, 0xfa, 0xb2,
	0xc4, 0x7c, 0x13, 0x3f, 0xb4, 0xc0, 0x07, 0x18, 0xcc, 0xeb, 0x35, 0xec, 0xdb, 0xc7, 0xfe, 0xae,
	0x5d, 0x05, 0x6f, 0xc4, 0x63, 0xe8, 0xcf, 0xeb, 0x24, 0xdc, 0x97, 0xff, 0xa6, 0x5c, 0xf5, 0x16,
	0xcf, 0xdf, 0x7f, 0x1d, 0xef, 0x04, 0xf6, 0x27, 0xcb, 0x35, 0x73, 0x8d, 0x2d, 0xfd, 0x59, 0x65,
	0xf2, 0x72, 0x7d, 0x75, 0xa3, 0x7f, 0x86, 0x83, 0x71, 0xc2, 0x33, 0xcc, 0xd4, 0xf8, 0x31, 0xb2,
	0xe5, 0x7a, 0x21, 0x0b, 0x2e, 0xf0, 0xff, 0x2b, 0xbd, 0x87, 0x5d, 0xb9, 0x85, 0x2d, 0x5a, 0xf5,
	0x48, 0x8d, 0x9e, 0xdf, 0x48, 0x5f, 0x93, 0x67, 0x9b, 0x87, 0x92, 0x8f, 0x81, 0x94, 0xa7, 0x4b,
	0xfe, 0x96, 0x6d, 0x55, 0xa9, 0x5e, 0xb5, 0x1a, 0x13, 0x2f, 0x76, 0xf4, 0x8d, 0x3d, 0xfa, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x21, 0xfb, 0x4f, 0x6d, 0x72, 0x05, 0x00, 0x00,
}
