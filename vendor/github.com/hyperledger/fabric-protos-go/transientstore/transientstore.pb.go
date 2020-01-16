// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transientstore/transientstore.proto

package transientstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rwset "github.com/hyperledger/fabric-protos-go/ledger/rwset"
	peer "github.com/hyperledger/fabric-protos-go/peer"
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

// TxPvtReadWriteSetWithConfigInfo encapsulates the transaction's private
// read-write set and additional information about the configurations such as
// the latest collection config when the transaction is simulated
type TxPvtReadWriteSetWithConfigInfo struct {
	EndorsedAt           uint64                                   `protobuf:"varint,1,opt,name=endorsed_at,json=endorsedAt,proto3" json:"endorsed_at,omitempty"`
	PvtRwset             *rwset.TxPvtReadWriteSet                 `protobuf:"bytes,2,opt,name=pvt_rwset,json=pvtRwset,proto3" json:"pvt_rwset,omitempty"`
	CollectionConfigs    map[string]*peer.CollectionConfigPackage `protobuf:"bytes,3,rep,name=collection_configs,json=collectionConfigs,proto3" json:"collection_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *TxPvtReadWriteSetWithConfigInfo) Reset()         { *m = TxPvtReadWriteSetWithConfigInfo{} }
func (m *TxPvtReadWriteSetWithConfigInfo) String() string { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSetWithConfigInfo) ProtoMessage()    {}
func (*TxPvtReadWriteSetWithConfigInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fca243668b157b9e, []int{0}
}

func (m *TxPvtReadWriteSetWithConfigInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo.Unmarshal(m, b)
}
func (m *TxPvtReadWriteSetWithConfigInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo.Marshal(b, m, deterministic)
}
func (m *TxPvtReadWriteSetWithConfigInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo.Merge(m, src)
}
func (m *TxPvtReadWriteSetWithConfigInfo) XXX_Size() int {
	return xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo.Size(m)
}
func (m *TxPvtReadWriteSetWithConfigInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxPvtReadWriteSetWithConfigInfo proto.InternalMessageInfo

func (m *TxPvtReadWriteSetWithConfigInfo) GetEndorsedAt() uint64 {
	if m != nil {
		return m.EndorsedAt
	}
	return 0
}

func (m *TxPvtReadWriteSetWithConfigInfo) GetPvtRwset() *rwset.TxPvtReadWriteSet {
	if m != nil {
		return m.PvtRwset
	}
	return nil
}

func (m *TxPvtReadWriteSetWithConfigInfo) GetCollectionConfigs() map[string]*peer.CollectionConfigPackage {
	if m != nil {
		return m.CollectionConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*TxPvtReadWriteSetWithConfigInfo)(nil), "transientstore.TxPvtReadWriteSetWithConfigInfo")
	proto.RegisterMapType((map[string]*peer.CollectionConfigPackage)(nil), "transientstore.TxPvtReadWriteSetWithConfigInfo.CollectionConfigsEntry")
}

func init() {
	proto.RegisterFile("transientstore/transientstore.proto", fileDescriptor_fca243668b157b9e)
}

var fileDescriptor_fca243668b157b9e = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x65, 0xdb, 0xef, 0x13, 0x9b, 0x82, 0x68, 0x40, 0x59, 0x7a, 0x69, 0xd1, 0x4b, 0x0f, 0x36,
	0x0b, 0x95, 0x8a, 0x78, 0xd3, 0xa2, 0xe0, 0xad, 0x44, 0xa1, 0xe0, 0xa5, 0xa4, 0xd9, 0xd9, 0x34,
	0x74, 0x4d, 0x96, 0x64, 0xba, 0xda, 0x5f, 0xea, 0xdf, 0x91, 0x6e, 0xac, 0xda, 0x56, 0xf0, 0x12,
	0x92, 0x37, 0xf3, 0xde, 0x9b, 0x97, 0x84, 0x9c, 0xa1, 0x13, 0xc6, 0x6b, 0x30, 0xe8, 0xd1, 0x3a,
	0x48, 0x36, 0x8f, 0xac, 0x70, 0x16, 0x2d, 0x3d, 0xd8, 0x44, 0x5b, 0x71, 0x0e, 0xa9, 0x02, 0x97,
	0xb8, 0x57, 0x0f, 0x18, 0xd6, 0xd0, 0xd9, 0x3a, 0x2e, 0x00, 0x5c, 0x22, 0x6d, 0x9e, 0x83, 0x44,
	0x6d, 0x4d, 0x80, 0x4f, 0xdf, 0x6b, 0xa4, 0xfd, 0xf4, 0x36, 0x2a, 0x91, 0x83, 0x48, 0xc7, 0x4e,
	0x23, 0x3c, 0x02, 0x8e, 0x35, 0xce, 0x86, 0xd6, 0x64, 0x5a, 0x3d, 0x98, 0xcc, 0xd2, 0x36, 0x69,
	0x82, 0x49, 0xad, 0xf3, 0x90, 0x4e, 0x04, 0xc6, 0x51, 0x27, 0xea, 0xfe, 0xe3, 0x64, 0x0d, 0xdd,
	0x20, 0x1d, 0x90, 0x46, 0x51, 0xe2, 0xa4, 0xb2, 0x8b, 0x6b, 0x9d, 0xa8, 0xdb, 0xec, 0xc7, 0x2c,
	0x98, 0xef, 0x68, 0xf3, 0xfd, 0xa2, 0x44, 0xbe, 0xaa, 0xd1, 0x05, 0xa1, 0xdf, 0xf3, 0x4c, 0x64,
	0x65, 0xe8, 0xe3, 0x7a, 0xa7, 0xde, 0x6d, 0xf6, 0xef, 0xd9, 0x56, 0xde, 0x3f, 0x86, 0x64, 0xc3,
	0x2f, 0xa5, 0x00, 0xfa, 0x3b, 0x83, 0x6e, 0xc9, 0x8f, 0xe4, 0x36, 0xde, 0x02, 0x72, 0xf2, 0x7b,
	0x33, 0x3d, 0x24, 0xf5, 0x39, 0x2c, 0xab, 0x80, 0x0d, 0xbe, 0xda, 0xd2, 0x01, 0xf9, 0x5f, 0x8a,
	0x7c, 0x01, 0x9f, 0xa9, 0xda, 0xe1, 0xd6, 0xfc, 0x8e, 0xdb, 0x48, 0xc8, 0xb9, 0x50, 0xc0, 0x43,
	0xf7, 0x75, 0xed, 0x2a, 0xba, 0xcd, 0xc8, 0xb9, 0x75, 0x8a, 0xcd, 0x96, 0x05, 0xb8, 0xf0, 0x2a,
	0x2c, 0x13, 0x53, 0xa7, 0xe5, 0x5a, 0x63, 0x33, 0xe0, 0xf3, 0xa5, 0xd2, 0x38, 0x5b, 0x4c, 0x99,
	0xb4, 0x2f, 0xc9, 0x0f, 0x52, 0x12, 0x48, 0xbd, 0x40, 0xea, 0x29, 0xbb, 0xf5, 0x11, 0xa6, 0x7b,
	0x55, 0xe5, 0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x71, 0x0c, 0xce, 0xd5, 0x30, 0x02, 0x00, 0x00,
}