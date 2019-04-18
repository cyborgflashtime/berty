// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/deviceinfo/application.proto

package deviceinfo

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Application_State int32

const (
	Application_Unknown    Application_State = 0
	Application_Kill       Application_State = 1
	Application_Background Application_State = 2
	Application_Foreground Application_State = 3
)

var Application_State_name = map[int32]string{
	0: "Unknown",
	1: "Kill",
	2: "Background",
	3: "Foreground",
}

var Application_State_value = map[string]int32{
	"Unknown":    0,
	"Kill":       1,
	"Background": 2,
	"Foreground": 3,
}

func (x Application_State) String() string {
	return proto.EnumName(Application_State_name, int32(x))
}

func (Application_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56744db749073b57, []int{0, 0}
}

type Application struct {
	State                Application_State `protobuf:"varint,1,opt,name=state,proto3,enum=berty.pkg.deviceinfo.Application_State" json:"state,omitempty"`
	Route                string            `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_56744db749073b57, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Application.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return m.Size()
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetState() Application_State {
	if m != nil {
		return m.State
	}
	return Application_Unknown
}

func (m *Application) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func init() {
	proto.RegisterEnum("berty.pkg.deviceinfo.Application_State", Application_State_name, Application_State_value)
	proto.RegisterType((*Application)(nil), "berty.pkg.deviceinfo.Application")
}

func init() { proto.RegisterFile("pkg/deviceinfo/application.proto", fileDescriptor_56744db749073b57) }

var fileDescriptor_56744db749073b57 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc8, 0x4e, 0xd7,
	0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x4f, 0x2c, 0x28, 0xc8, 0xc9,
	0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0x4a,
	0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0xc8, 0x4e, 0xd7, 0x43, 0xa8, 0x53, 0x5a, 0xc4, 0xc8, 0xc5, 0xed,
	0x88, 0x50, 0x2b, 0x64, 0xcb, 0xc5, 0x5a, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x67, 0xa4, 0xae, 0x87, 0x4d, 0x97, 0x1e, 0x92, 0x0e, 0xbd, 0x60, 0x90, 0xf2, 0x20, 0x88,
	0x2e, 0x21, 0x11, 0x2e, 0xd6, 0xa2, 0xfc, 0xd2, 0x92, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x08, 0x47, 0xc9, 0x8e, 0x8b, 0x15, 0xac, 0x4a, 0x88, 0x9b, 0x8b, 0x3d, 0x34, 0x2f, 0x3b,
	0x2f, 0xbf, 0x3c, 0x4f, 0x80, 0x41, 0x88, 0x83, 0x8b, 0xc5, 0x3b, 0x33, 0x27, 0x47, 0x80, 0x51,
	0x88, 0x8f, 0x8b, 0xcb, 0x29, 0x31, 0x39, 0x3b, 0xbd, 0x28, 0xbf, 0x34, 0x2f, 0x45, 0x80, 0x09,
	0xc4, 0x77, 0xcb, 0x2f, 0x4a, 0x85, 0xf2, 0x99, 0x9d, 0x0c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0xe4, 0x20, 0xce, 0x2a,
	0x49, 0x4d, 0xce, 0xd0, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x47, 0xf5, 0x7e, 0x12, 0x1b, 0xd8, 0xcf,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x72, 0x23, 0xf4, 0x17, 0x01, 0x00, 0x00,
}

func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApplication(dAtA, i, uint64(m.State))
	}
	if len(m.Route) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Route)))
		i += copy(dAtA[i:], m.Route)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Application) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovApplication(uint64(m.State))
	}
	l = len(m.Route)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Application) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Application_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Route = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthApplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplication
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
				next, err := skipApplication(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthApplication
				}
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
	ErrInvalidLengthApplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication   = fmt.Errorf("proto: integer overflow")
)
