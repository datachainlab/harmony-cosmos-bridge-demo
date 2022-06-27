// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/chains/harmony/config.proto

package harmony

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ChainConfig struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// human name of a chain-id
	HarmonyChainId string `protobuf:"bytes,2,opt,name=harmony_chain_id,json=harmonyChainId,proto3" json:"harmony_chain_id,omitempty"`
	ShardId        uint32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	ShardRpcAddr   string `protobuf:"bytes,4,opt,name=shard_rpc_addr,json=shardRpcAddr,proto3" json:"shard_rpc_addr,omitempty"`
	// if shard_id = 0, set the same address as shard_rpc_addr
	BeaconRpcAddr string `protobuf:"bytes,5,opt,name=beacon_rpc_addr,json=beaconRpcAddr,proto3" json:"beacon_rpc_addr,omitempty"`
	// use for relayer
	ShardPrivateKey string `protobuf:"bytes,6,opt,name=shard_private_key,json=shardPrivateKey,proto3" json:"shard_private_key,omitempty"`
	// if shard_id = 0, set the same key as shard_private_key
	BeaconPrivateKey  string `protobuf:"bytes,7,opt,name=beacon_private_key,json=beaconPrivateKey,proto3" json:"beacon_private_key,omitempty"`
	IbcHostAddress    string `protobuf:"bytes,8,opt,name=ibc_host_address,json=ibcHostAddress,proto3" json:"ibc_host_address,omitempty"`
	IbcHandlerAddress string `protobuf:"bytes,9,opt,name=ibc_handler_address,json=ibcHandlerAddress,proto3" json:"ibc_handler_address,omitempty"`
	GasLimit          uint64 `protobuf:"varint,10,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice          int64  `protobuf:"varint,11,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d1a31f40ed93c46, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

type ProverConfig struct {
	TrustingPeriod string `protobuf:"bytes,1,opt,name=trusting_period,json=trustingPeriod,proto3" json:"trusting_period,omitempty"`
}

func (m *ProverConfig) Reset()         { *m = ProverConfig{} }
func (m *ProverConfig) String() string { return proto.CompactTextString(m) }
func (*ProverConfig) ProtoMessage()    {}
func (*ProverConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d1a31f40ed93c46, []int{1}
}
func (m *ProverConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProverConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProverConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProverConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProverConfig.Merge(m, src)
}
func (m *ProverConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProverConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProverConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProverConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.harmony.config.ChainConfig")
	proto.RegisterType((*ProverConfig)(nil), "relayer.chains.harmony.config.ProverConfig")
}

func init() {
	proto.RegisterFile("relayer/chains/harmony/config.proto", fileDescriptor_3d1a31f40ed93c46)
}

var fileDescriptor_3d1a31f40ed93c46 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x07, 0xf0, 0x86, 0x8e, 0xb5, 0xf5, 0xb6, 0xb6, 0x33, 0x1c, 0x32, 0x10, 0x51, 0x35, 0x10,
	0x44, 0x88, 0x26, 0x07, 0x0e, 0x9c, 0x47, 0x2f, 0x54, 0x70, 0xa8, 0x7a, 0xe4, 0x12, 0x39, 0xb6,
	0x71, 0x2c, 0x9a, 0x38, 0xfa, 0xec, 0x4d, 0xca, 0x5b, 0xf0, 0x1a, 0xbc, 0xc9, 0x8e, 0x3b, 0x72,
	0x84, 0xf6, 0x45, 0x50, 0x3e, 0x27, 0x63, 0x07, 0x6e, 0xfd, 0xfe, 0xff, 0x9f, 0xbf, 0x2a, 0xb2,
	0xc9, 0x4b, 0x90, 0x3b, 0xd6, 0x48, 0x48, 0x79, 0xc1, 0x74, 0x65, 0xd3, 0x82, 0x41, 0x69, 0xaa,
	0x26, 0xe5, 0xa6, 0xfa, 0xa6, 0x55, 0x52, 0x83, 0x71, 0x86, 0xbe, 0xe8, 0x50, 0xe2, 0x51, 0xd2,
	0xa1, 0xc4, 0xa3, 0x67, 0x4f, 0x95, 0x51, 0x06, 0x65, 0xda, 0xfe, 0xf2, 0x87, 0x2e, 0x7f, 0x0e,
	0xc9, 0xc9, 0xaa, 0xf5, 0x2b, 0x54, 0xf4, 0x82, 0x8c, 0xf1, 0x78, 0xa6, 0x45, 0x18, 0x2c, 0x82,
	0x78, 0xb2, 0x1d, 0xe1, 0xbc, 0x16, 0x34, 0x26, 0xf3, 0x6e, 0x65, 0x76, 0x4f, 0x1e, 0x21, 0x99,
	0x76, 0xf9, 0xaa, 0x93, 0x17, 0x64, 0x6c, 0x0b, 0x06, 0xa2, 0x15, 0xc3, 0x45, 0x10, 0x9f, 0x6d,
	0x47, 0x38, 0xaf, 0x05, 0x7d, 0x45, 0xa6, 0xbe, 0x82, 0x9a, 0x67, 0x4c, 0x08, 0x08, 0x8f, 0x70,
	0xc5, 0x29, 0xa6, 0xdb, 0x9a, 0x5f, 0x09, 0x01, 0xf4, 0x35, 0x99, 0xe5, 0x92, 0x71, 0x53, 0xfd,
	0x63, 0x8f, 0x91, 0x9d, 0xf9, 0xb8, 0x77, 0x6f, 0xc9, 0xb9, 0xdf, 0x56, 0x83, 0xbe, 0x61, 0x4e,
	0x66, 0xdf, 0x65, 0x13, 0x1e, 0xa3, 0x9c, 0x61, 0xb1, 0xf1, 0xf9, 0x67, 0xd9, 0xd0, 0x77, 0x84,
	0x76, 0x3b, 0x1f, 0xe2, 0x11, 0xe2, 0xb9, 0x6f, 0x1e, 0xe8, 0x98, 0xcc, 0x75, 0xce, 0xb3, 0xc2,
	0x58, 0x87, 0xff, 0x2f, 0xad, 0x0d, 0xc7, 0xfe, 0x63, 0x75, 0xce, 0x3f, 0x19, 0xeb, 0xae, 0x7c,
	0x4a, 0x13, 0xf2, 0x04, 0x25, 0xab, 0xc4, 0x4e, 0xc2, 0x3d, 0x9e, 0x20, 0x3e, 0x6f, 0xb1, 0x6f,
	0x7a, 0xff, 0x9c, 0x4c, 0x14, 0xb3, 0xd9, 0x4e, 0x97, 0xda, 0x85, 0x64, 0x11, 0xc4, 0x47, 0xdb,
	0xb1, 0x62, 0xf6, 0x4b, 0x3b, 0xf7, 0x65, 0x0d, 0x9a, 0xcb, 0xf0, 0x64, 0x11, 0xc4, 0x43, 0x2c,
	0x37, 0xed, 0x7c, 0xf9, 0x81, 0x9c, 0x6e, 0xc0, 0xdc, 0x48, 0xe8, 0xee, 0xea, 0x0d, 0x99, 0x39,
	0xb8, 0xb6, 0x4e, 0x57, 0x2a, 0xab, 0x25, 0x68, 0xd3, 0x5f, 0xd9, 0xb4, 0x8f, 0x37, 0x98, 0x7e,
	0x54, 0xb7, 0x7f, 0xa2, 0xc1, 0xed, 0x3e, 0x0a, 0xee, 0xf6, 0x51, 0xf0, 0x7b, 0x1f, 0x05, 0x3f,
	0x0e, 0xd1, 0xe0, 0xee, 0x10, 0x0d, 0x7e, 0x1d, 0xa2, 0xc1, 0xd7, 0xb5, 0xd2, 0xae, 0xb8, 0xce,
	0x13, 0x6e, 0xca, 0x54, 0x30, 0xc7, 0xf0, 0x72, 0x77, 0x2c, 0xef, 0x5f, 0xd9, 0x92, 0x1b, 0x5b,
	0x1a, 0xbb, 0xcc, 0x41, 0x0b, 0x25, 0x97, 0x42, 0x96, 0x26, 0xfd, 0xff, 0x7b, 0xcc, 0x8f, 0xf1,
	0x51, 0xbd, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x37, 0xaf, 0x08, 0xb0, 0x02, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasPrice != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x58
	}
	if m.GasLimit != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x50
	}
	if len(m.IbcHandlerAddress) > 0 {
		i -= len(m.IbcHandlerAddress)
		copy(dAtA[i:], m.IbcHandlerAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHandlerAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IbcHostAddress) > 0 {
		i -= len(m.IbcHostAddress)
		copy(dAtA[i:], m.IbcHostAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHostAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BeaconPrivateKey) > 0 {
		i -= len(m.BeaconPrivateKey)
		copy(dAtA[i:], m.BeaconPrivateKey)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.BeaconPrivateKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ShardPrivateKey) > 0 {
		i -= len(m.ShardPrivateKey)
		copy(dAtA[i:], m.ShardPrivateKey)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ShardPrivateKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BeaconRpcAddr) > 0 {
		i -= len(m.BeaconRpcAddr)
		copy(dAtA[i:], m.BeaconRpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.BeaconRpcAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ShardRpcAddr) > 0 {
		i -= len(m.ShardRpcAddr)
		copy(dAtA[i:], m.ShardRpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ShardRpcAddr)))
		i--
		dAtA[i] = 0x22
	}
	if m.ShardId != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.ShardId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.HarmonyChainId) > 0 {
		i -= len(m.HarmonyChainId)
		copy(dAtA[i:], m.HarmonyChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.HarmonyChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProverConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProverConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProverConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrustingPeriod) > 0 {
		i -= len(m.TrustingPeriod)
		copy(dAtA[i:], m.TrustingPeriod)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.TrustingPeriod)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.HarmonyChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ShardId != 0 {
		n += 1 + sovConfig(uint64(m.ShardId))
	}
	l = len(m.ShardRpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.BeaconRpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ShardPrivateKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.BeaconPrivateKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHostAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHandlerAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovConfig(uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		n += 1 + sovConfig(uint64(m.GasPrice))
	}
	return n
}

func (m *ProverConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TrustingPeriod)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HarmonyChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HarmonyChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardId", wireType)
			}
			m.ShardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardRpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardRpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeaconRpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeaconRpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeaconPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeaconPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHostAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHandlerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHandlerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProverConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProverConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProverConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustingPeriod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustingPeriod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
