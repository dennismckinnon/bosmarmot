// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		rpc.proto

	It has these top-level messages:
		ResultStatus
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import crypto "github.com/hyperledger/burrow/crypto"
import tendermint "github.com/hyperledger/burrow/consensus/tendermint"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResultStatus struct {
	ChainID           string                                        `protobuf:"bytes,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	GenesisHash       github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,4,opt,name=GenesisHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"GenesisHash"`
	NodeInfo          *tendermint.NodeInfo                          `protobuf:"bytes,2,opt,name=NodeInfo" json:"NodeInfo,omitempty"`
	BurrowVersion     string                                        `protobuf:"bytes,3,opt,name=BurrowVersion,proto3" json:"BurrowVersion,omitempty"`
	PublicKey         crypto.PublicKey                              `protobuf:"bytes,5,opt,name=PublicKey" json:"PublicKey"`
	LatestBlockHash   github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,6,opt,name=LatestBlockHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"LatestBlockHash"`
	LatestBlockHeight uint64                                        `protobuf:"varint,7,opt,name=LatestBlockHeight,proto3" json:""`
	LatestBlockTime   time.Time                                     `protobuf:"bytes,8,opt,name=LatestBlockTime,stdtime" json:"LatestBlockTime"`
}

func (m *ResultStatus) Reset()                    { *m = ResultStatus{} }
func (m *ResultStatus) String() string            { return proto.CompactTextString(m) }
func (*ResultStatus) ProtoMessage()               {}
func (*ResultStatus) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

func (m *ResultStatus) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ResultStatus) GetNodeInfo() *tendermint.NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *ResultStatus) GetBurrowVersion() string {
	if m != nil {
		return m.BurrowVersion
	}
	return ""
}

func (m *ResultStatus) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func (m *ResultStatus) GetLatestBlockHeight() uint64 {
	if m != nil {
		return m.LatestBlockHeight
	}
	return 0
}

func (m *ResultStatus) GetLatestBlockTime() time.Time {
	if m != nil {
		return m.LatestBlockTime
	}
	return time.Time{}
}

func (*ResultStatus) XXX_MessageName() string {
	return "rpc.ResultStatus"
}
func init() {
	proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
	golang_proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
}
func (m *ResultStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	if m.NodeInfo != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.NodeInfo.Size()))
		n1, err := m.NodeInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.BurrowVersion) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.BurrowVersion)))
		i += copy(dAtA[i:], m.BurrowVersion)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.GenesisHash.Size()))
	n2, err := m.GenesisHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x2a
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.PublicKey.Size()))
	n3, err := m.PublicKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.LatestBlockHash.Size()))
	n4, err := m.LatestBlockHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.LatestBlockHeight != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.LatestBlockHeight))
	}
	dAtA[i] = 0x42
	i++
	i = encodeVarintRpc(dAtA, i, uint64(types.SizeOfStdTime(m.LatestBlockTime)))
	n5, err := types.StdTimeMarshalTo(m.LatestBlockTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResultStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.BurrowVersion)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = m.GenesisHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	l = m.PublicKey.Size()
	n += 1 + l + sovRpc(uint64(l))
	l = m.LatestBlockHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	if m.LatestBlockHeight != 0 {
		n += 1 + sovRpc(uint64(m.LatestBlockHeight))
	}
	l = types.SizeOfStdTime(m.LatestBlockTime)
	n += 1 + l + sovRpc(uint64(l))
	return n
}

func sovRpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResultStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResultStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &tendermint.NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurrowVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurrowVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestBlockHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHeight", wireType)
			}
			m.LatestBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlockHeight |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.LatestBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }
func init() { golang_proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc6, 0x34, 0xdc, 0xb5, 0xbe, 0x22, 0x38, 0x8b, 0xc1, 0xea, 0x90, 0x04, 0xc4, 0x90, 0x05,
	0x07, 0x1d, 0x3a, 0xb1, 0x1b, 0x24, 0xee, 0x04, 0x3a, 0xa1, 0x70, 0x02, 0x89, 0x05, 0x25, 0xe9,
	0xbb, 0xc4, 0x22, 0x89, 0x23, 0xdb, 0x11, 0xe4, 0x5f, 0xf0, 0x93, 0x18, 0x3b, 0x32, 0x33, 0x1c,
	0xa8, 0xdd, 0xd8, 0xd9, 0x51, 0xdd, 0xa6, 0x4d, 0x8b, 0xc4, 0xc2, 0xe6, 0xef, 0x7d, 0xcf, 0xdf,
	0xfb, 0xde, 0xfb, 0xf0, 0x48, 0xd5, 0x29, 0xab, 0x95, 0x34, 0x92, 0x0c, 0x54, 0x9d, 0x4e, 0x1e,
	0x65, 0xc2, 0xe4, 0x4d, 0xc2, 0x52, 0x59, 0x86, 0x99, 0xcc, 0x64, 0x68, 0xb9, 0xa4, 0xb9, 0xb2,
	0xc8, 0x02, 0xfb, 0x5a, 0xfd, 0x99, 0x8c, 0x53, 0xd5, 0xd6, 0xa6, 0x43, 0x77, 0x0d, 0x54, 0x53,
	0x50, 0xa5, 0xa8, 0xcc, 0xba, 0xe2, 0x65, 0x52, 0x66, 0x05, 0x6c, 0x55, 0x8c, 0x28, 0x41, 0x9b,
	0xb8, 0xac, 0x57, 0x0d, 0x0f, 0x7e, 0x0f, 0xf0, 0x38, 0x02, 0xdd, 0x14, 0xe6, 0x8d, 0x89, 0x4d,
	0xa3, 0x09, 0xc5, 0x87, 0xcf, 0xf2, 0x58, 0x54, 0xe7, 0xcf, 0x29, 0xf2, 0x51, 0x30, 0x8a, 0x3a,
	0x48, 0x1e, 0xe3, 0xe1, 0x85, 0x9c, 0xc2, 0x79, 0x75, 0x25, 0xe9, 0x4d, 0x1f, 0x05, 0x47, 0x27,
	0xf7, 0x58, 0x6f, 0x60, 0xc7, 0x45, 0x9b, 0x2e, 0xf2, 0x10, 0xdf, 0xe6, 0x8d, 0x52, 0xf2, 0xd3,
	0x5b, 0x50, 0x5a, 0xc8, 0x8a, 0x0e, 0xac, 0xe2, 0x6e, 0x91, 0xbc, 0xc3, 0x47, 0x2f, 0xa0, 0x02,
	0x2d, 0xf4, 0x59, 0xac, 0x73, 0xea, 0xf8, 0x28, 0x18, 0xf3, 0xd3, 0xd9, 0xb5, 0x77, 0xe3, 0xfb,
	0xb5, 0xd7, 0xbf, 0x47, 0xde, 0xd6, 0xa0, 0x0a, 0x98, 0x66, 0xa0, 0xc2, 0xc4, 0x4a, 0x84, 0x89,
	0xa8, 0x62, 0xd5, 0xb2, 0x33, 0xf8, 0xcc, 0x5b, 0x03, 0x3a, 0xea, 0x2b, 0x91, 0x53, 0x3c, 0x7a,
	0xdd, 0x24, 0x85, 0x48, 0x5f, 0x42, 0x4b, 0x6f, 0x59, 0xc7, 0xc7, 0x6c, 0x7d, 0xb0, 0x0d, 0xc1,
	0x9d, 0xe5, 0xa4, 0x68, 0xdb, 0x49, 0x3e, 0xe0, 0x3b, 0xaf, 0x62, 0x03, 0xda, 0xf0, 0x42, 0xa6,
	0x1f, 0xad, 0xa7, 0x83, 0xff, 0xf1, 0xb4, 0xaf, 0x46, 0x4e, 0xf0, 0x71, 0xbf, 0x04, 0x22, 0xcb,
	0x0d, 0x3d, 0xf4, 0x51, 0xe0, 0x70, 0xe7, 0xd7, 0xd2, 0xcc, 0xdf, 0x34, 0xb9, 0xd8, 0x31, 0x75,
	0x29, 0x4a, 0xa0, 0x43, 0xbb, 0xd1, 0x84, 0xad, 0x22, 0x66, 0x5d, 0xc4, 0xec, 0xb2, 0x8b, 0x98,
	0x0f, 0x97, 0x86, 0xbf, 0xfc, 0xf0, 0x50, 0xb4, 0xff, 0x99, 0x3f, 0x9d, 0xcd, 0x5d, 0xf4, 0x6d,
	0xee, 0xa2, 0x9f, 0x73, 0x17, 0x7d, 0x5d, 0xb8, 0x68, 0xb6, 0x70, 0xd1, 0xfb, 0xfb, 0xff, 0xde,
	0x4c, 0xd5, 0x69, 0x72, 0x60, 0xe7, 0x3c, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xca, 0x39, 0xb6,
	0x91, 0xb9, 0x02, 0x00, 0x00,
}
