// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flow/layers/dhcpv4.proto

package layers

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

// LayerDHCPv4 wrapper to generate extra layer
type DHCPv4 struct {
	Contents     []byte `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	Payload      []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	HardwareLen  uint8  `protobuf:"varint,3,opt,name=hardware_len,json=hardwareLen,proto3,casttype=uint8" json:"hardware_len,omitempty"`
	HardwareOpts uint8  `protobuf:"varint,4,opt,name=hardware_opts,json=hardwareOpts,proto3,casttype=uint8" json:"hardware_opts,omitempty"`
	Xid          uint32 `protobuf:"varint,5,opt,name=xid,proto3" json:"xid,omitempty"`
	Secs         uint16 `protobuf:"varint,6,opt,name=secs,proto3,casttype=uint16" json:"secs,omitempty"`
	Flags        uint16 `protobuf:"varint,7,opt,name=flags,proto3,casttype=uint16" json:"flags,omitempty"`
	ServerName   []byte `protobuf:"bytes,8,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	File         []byte `protobuf:"bytes,9,opt,name=file,proto3" json:"file,omitempty"`
}

func (m *DHCPv4) Reset()         { *m = DHCPv4{} }
func (m *DHCPv4) String() string { return proto.CompactTextString(m) }
func (*DHCPv4) ProtoMessage()    {}
func (*DHCPv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_212f4ad6f78ca751, []int{0}
}
func (m *DHCPv4) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DHCPv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DHCPv4.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DHCPv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPv4.Merge(m, src)
}
func (m *DHCPv4) XXX_Size() int {
	return m.ProtoSize()
}
func (m *DHCPv4) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPv4.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPv4 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DHCPv4)(nil), "layers.DHCPv4")
}

func init() { proto.RegisterFile("flow/layers/dhcpv4.proto", fileDescriptor_212f4ad6f78ca751) }

var fileDescriptor_212f4ad6f78ca751 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x1c, 0xc5, 0x5b, 0x28, 0x05, 0x0e, 0x48, 0xcc, 0xc5, 0xe1, 0xc2, 0x70, 0x25, 0x4e, 0x0c, 0xda,
	0xc6, 0x48, 0x8c, 0x71, 0x44, 0x07, 0x07, 0xa2, 0xa6, 0xa3, 0x0b, 0x39, 0xda, 0xa3, 0x54, 0x4b,
	0xaf, 0xb9, 0x3b, 0x8a, 0x7c, 0x03, 0x47, 0x3f, 0x02, 0xbb, 0x5f, 0xc4, 0x91, 0xd1, 0x89, 0x18,
	0xfa, 0x2d, 0x98, 0x4c, 0xaf, 0xa1, 0x31, 0xba, 0xbd, 0xf7, 0x7e, 0xef, 0x7f, 0xc9, 0xe5, 0x01,
	0x34, 0x8d, 0xd8, 0xd2, 0x89, 0xc8, 0x8a, 0x72, 0xe1, 0xf8, 0x33, 0x2f, 0x49, 0x07, 0x76, 0xc2,
	0x99, 0x64, 0xd0, 0x2c, 0xc2, 0xee, 0x71, 0xc0, 0x02, 0xa6, 0x22, 0x27, 0x57, 0x05, 0x3d, 0xf9,
	0xa8, 0x00, 0xf3, 0xf6, 0xee, 0xe6, 0x31, 0x1d, 0xc0, 0x2e, 0x68, 0x78, 0x2c, 0x96, 0x34, 0x96,
	0x02, 0xe9, 0x3d, 0xbd, 0xdf, 0x76, 0x4b, 0x0f, 0x11, 0xa8, 0x27, 0x64, 0x15, 0x31, 0xe2, 0xa3,
	0x8a, 0x42, 0x07, 0x0b, 0x4f, 0x41, 0x7b, 0x46, 0xb8, 0xbf, 0x24, 0x9c, 0x8e, 0x23, 0x1a, 0xa3,
	0x6a, 0x4f, 0xef, 0x77, 0x86, 0xcd, 0xfd, 0xd6, 0xaa, 0x2d, 0xc2, 0x58, 0x5e, 0xb9, 0xad, 0x03,
	0x1e, 0xd1, 0x18, 0xda, 0xa0, 0x53, 0xb6, 0x59, 0x22, 0x05, 0x32, 0xfe, 0xd6, 0xcb, 0xd7, 0x1e,
	0x12, 0x29, 0xe0, 0x11, 0xa8, 0xbe, 0x86, 0x3e, 0xaa, 0xe5, 0x2d, 0x37, 0x97, 0x10, 0x03, 0x43,
	0x50, 0x4f, 0x20, 0x53, 0x1d, 0x82, 0xfd, 0xd6, 0x32, 0xf3, 0xc3, 0xf3, 0x4b, 0x57, 0xe5, 0xb0,
	0x07, 0x6a, 0xd3, 0x88, 0x04, 0x02, 0xd5, 0xff, 0x15, 0x0a, 0x00, 0x2d, 0xd0, 0x12, 0x94, 0xa7,
	0x94, 0x8f, 0x63, 0x32, 0xa7, 0xa8, 0xa1, 0xfe, 0x03, 0x8a, 0xe8, 0x9e, 0xcc, 0x29, 0x84, 0xc0,
	0x98, 0x86, 0x11, 0x45, 0x4d, 0x45, 0x94, 0xbe, 0x36, 0xde, 0xd6, 0x96, 0x36, 0x1c, 0x7d, 0xee,
	0xb0, 0xbe, 0xd9, 0x61, 0xfd, 0x7b, 0x87, 0xb5, 0xf7, 0x0c, 0x6b, 0xeb, 0x0c, 0xeb, 0x9b, 0x0c,
	0x6b, 0x5f, 0x19, 0xd6, 0x9e, 0xec, 0x20, 0x94, 0xb3, 0xc5, 0xc4, 0xf6, 0xd8, 0xdc, 0x11, 0x2f,
	0x2b, 0x3f, 0x4c, 0xe9, 0x59, 0xc2, 0xd9, 0x33, 0xf5, 0xe4, 0xc1, 0x3b, 0xbf, 0x66, 0x9a, 0x98,
	0x6a, 0x82, 0x8b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xec, 0x9c, 0x95, 0xbc, 0x01, 0x00,
	0x00,
}

func (m *DHCPv4) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DHCPv4) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DHCPv4) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.File) > 0 {
		i -= len(m.File)
		copy(dAtA[i:], m.File)
		i = encodeVarintDhcpv4(dAtA, i, uint64(len(m.File)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ServerName) > 0 {
		i -= len(m.ServerName)
		copy(dAtA[i:], m.ServerName)
		i = encodeVarintDhcpv4(dAtA, i, uint64(len(m.ServerName)))
		i--
		dAtA[i] = 0x42
	}
	if m.Flags != 0 {
		i = encodeVarintDhcpv4(dAtA, i, uint64(m.Flags))
		i--
		dAtA[i] = 0x38
	}
	if m.Secs != 0 {
		i = encodeVarintDhcpv4(dAtA, i, uint64(m.Secs))
		i--
		dAtA[i] = 0x30
	}
	if m.Xid != 0 {
		i = encodeVarintDhcpv4(dAtA, i, uint64(m.Xid))
		i--
		dAtA[i] = 0x28
	}
	if m.HardwareOpts != 0 {
		i = encodeVarintDhcpv4(dAtA, i, uint64(m.HardwareOpts))
		i--
		dAtA[i] = 0x20
	}
	if m.HardwareLen != 0 {
		i = encodeVarintDhcpv4(dAtA, i, uint64(m.HardwareLen))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintDhcpv4(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contents) > 0 {
		i -= len(m.Contents)
		copy(dAtA[i:], m.Contents)
		i = encodeVarintDhcpv4(dAtA, i, uint64(len(m.Contents)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDhcpv4(dAtA []byte, offset int, v uint64) int {
	offset -= sovDhcpv4(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DHCPv4) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contents)
	if l > 0 {
		n += 1 + l + sovDhcpv4(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovDhcpv4(uint64(l))
	}
	if m.HardwareLen != 0 {
		n += 1 + sovDhcpv4(uint64(m.HardwareLen))
	}
	if m.HardwareOpts != 0 {
		n += 1 + sovDhcpv4(uint64(m.HardwareOpts))
	}
	if m.Xid != 0 {
		n += 1 + sovDhcpv4(uint64(m.Xid))
	}
	if m.Secs != 0 {
		n += 1 + sovDhcpv4(uint64(m.Secs))
	}
	if m.Flags != 0 {
		n += 1 + sovDhcpv4(uint64(m.Flags))
	}
	l = len(m.ServerName)
	if l > 0 {
		n += 1 + l + sovDhcpv4(uint64(l))
	}
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovDhcpv4(uint64(l))
	}
	return n
}

func sovDhcpv4(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDhcpv4(x uint64) (n int) {
	return sovDhcpv4(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DHCPv4) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDhcpv4
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
			return fmt.Errorf("proto: DHCPv4: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DHCPv4: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDhcpv4
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDhcpv4
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents[:0], dAtA[iNdEx:postIndex]...)
			if m.Contents == nil {
				m.Contents = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDhcpv4
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDhcpv4
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardwareLen", wireType)
			}
			m.HardwareLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HardwareLen |= uint8(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardwareOpts", wireType)
			}
			m.HardwareOpts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HardwareOpts |= uint8(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Xid", wireType)
			}
			m.Xid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Xid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secs", wireType)
			}
			m.Secs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Secs |= uint16(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flags |= uint16(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDhcpv4
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDhcpv4
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerName = append(m.ServerName[:0], dAtA[iNdEx:postIndex]...)
			if m.ServerName == nil {
				m.ServerName = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDhcpv4
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDhcpv4
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDhcpv4
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = append(m.File[:0], dAtA[iNdEx:postIndex]...)
			if m.File == nil {
				m.File = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDhcpv4(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDhcpv4
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDhcpv4
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
func skipDhcpv4(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDhcpv4
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
					return 0, ErrIntOverflowDhcpv4
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
					return 0, ErrIntOverflowDhcpv4
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
				return 0, ErrInvalidLengthDhcpv4
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDhcpv4
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDhcpv4
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDhcpv4        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDhcpv4          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDhcpv4 = fmt.Errorf("proto: unexpected end of group")
)
