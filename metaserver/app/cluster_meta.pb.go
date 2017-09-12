// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cluster_meta.proto

/*
	Package main is a generated protocol buffer package.

	It is generated from these files:
		cluster_meta.proto

	It has these top-level messages:
		ClusterMeta
		InstanceNameList
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import gxredis "github.com/AlexStocks/goext/database/redis"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ClusterMeta struct {
	Version   int32                        `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Instances map[string]*gxredis.Instance `protobuf:"bytes,2,rep,name=Instances" json:"Instances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ClusterMeta) Reset()                    { *m = ClusterMeta{} }
func (*ClusterMeta) ProtoMessage()               {}
func (*ClusterMeta) Descriptor() ([]byte, []int) { return fileDescriptorClusterMeta, []int{0} }

type InstanceNameList struct {
	List []string `protobuf:"bytes,1,rep,name=List" json:"List,omitempty"`
}

func (m *InstanceNameList) Reset()                    { *m = InstanceNameList{} }
func (*InstanceNameList) ProtoMessage()               {}
func (*InstanceNameList) Descriptor() ([]byte, []int) { return fileDescriptorClusterMeta, []int{1} }

func init() {
	proto.RegisterType((*ClusterMeta)(nil), "main.ClusterMeta")
	proto.RegisterType((*InstanceNameList)(nil), "main.InstanceNameList")
}
func (this *ClusterMeta) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClusterMeta)
	if !ok {
		that2, ok := that.(ClusterMeta)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ClusterMeta")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClusterMeta but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClusterMeta but is not nil && this == nil")
	}
	if this.Version != that1.Version {
		return fmt.Errorf("Version this(%v) Not Equal that(%v)", this.Version, that1.Version)
	}
	if len(this.Instances) != len(that1.Instances) {
		return fmt.Errorf("Instances this(%v) Not Equal that(%v)", len(this.Instances), len(that1.Instances))
	}
	for i := range this.Instances {
		if !this.Instances[i].Equal(that1.Instances[i]) {
			return fmt.Errorf("Instances this[%v](%v) Not Equal that[%v](%v)", i, this.Instances[i], i, that1.Instances[i])
		}
	}
	return nil
}
func (this *ClusterMeta) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClusterMeta)
	if !ok {
		that2, ok := that.(ClusterMeta)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if len(this.Instances) != len(that1.Instances) {
		return false
	}
	for i := range this.Instances {
		if !this.Instances[i].Equal(that1.Instances[i]) {
			return false
		}
	}
	return true
}
func (this *InstanceNameList) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*InstanceNameList)
	if !ok {
		that2, ok := that.(InstanceNameList)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *InstanceNameList")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *InstanceNameList but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *InstanceNameList but is not nil && this == nil")
	}
	if len(this.List) != len(that1.List) {
		return fmt.Errorf("List this(%v) Not Equal that(%v)", len(this.List), len(that1.List))
	}
	for i := range this.List {
		if this.List[i] != that1.List[i] {
			return fmt.Errorf("List this[%v](%v) Not Equal that[%v](%v)", i, this.List[i], i, that1.List[i])
		}
	}
	return nil
}
func (this *InstanceNameList) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceNameList)
	if !ok {
		that2, ok := that.(InstanceNameList)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.List) != len(that1.List) {
		return false
	}
	for i := range this.List {
		if this.List[i] != that1.List[i] {
			return false
		}
	}
	return true
}
func (this *ClusterMeta) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.ClusterMeta{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	keysForInstances := make([]string, 0, len(this.Instances))
	for k, _ := range this.Instances {
		keysForInstances = append(keysForInstances, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForInstances)
	mapStringForInstances := "map[string]*gxredis.Instance{"
	for _, k := range keysForInstances {
		mapStringForInstances += fmt.Sprintf("%#v: %#v,", k, this.Instances[k])
	}
	mapStringForInstances += "}"
	if this.Instances != nil {
		s = append(s, "Instances: "+mapStringForInstances+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceNameList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.InstanceNameList{")
	s = append(s, "List: "+fmt.Sprintf("%#v", this.List)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClusterMeta(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ClusterMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClusterMeta(dAtA, i, uint64(m.Version))
	}
	if len(m.Instances) > 0 {
		for k, _ := range m.Instances {
			dAtA[i] = 0x12
			i++
			v := m.Instances[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovClusterMeta(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovClusterMeta(uint64(len(k))) + msgSize
			i = encodeVarintClusterMeta(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintClusterMeta(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintClusterMeta(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func (m *InstanceNameList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceNameList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, s := range m.List {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64ClusterMeta(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ClusterMeta(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClusterMeta(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClusterMeta) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovClusterMeta(uint64(m.Version))
	}
	if len(m.Instances) > 0 {
		for k, v := range m.Instances {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovClusterMeta(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovClusterMeta(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovClusterMeta(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *InstanceNameList) Size() (n int) {
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, s := range m.List {
			l = len(s)
			n += 1 + l + sovClusterMeta(uint64(l))
		}
	}
	return n
}

func sovClusterMeta(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClusterMeta(x uint64) (n int) {
	return sovClusterMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterMeta) String() string {
	if this == nil {
		return "nil"
	}
	keysForInstances := make([]string, 0, len(this.Instances))
	for k, _ := range this.Instances {
		keysForInstances = append(keysForInstances, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForInstances)
	mapStringForInstances := "map[string]*gxredis.Instance{"
	for _, k := range keysForInstances {
		mapStringForInstances += fmt.Sprintf("%v: %v,", k, this.Instances[k])
	}
	mapStringForInstances += "}"
	s := strings.Join([]string{`&ClusterMeta{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Instances:` + mapStringForInstances + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceNameList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceNameList{`,
		`List:` + fmt.Sprintf("%v", this.List) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClusterMeta(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterMeta
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
			return fmt.Errorf("proto: ClusterMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMeta
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
				return ErrInvalidLengthClusterMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Instances == nil {
				m.Instances = make(map[string]*gxredis.Instance)
			}
			var mapkey string
			var mapvalue *gxredis.Instance
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClusterMeta
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClusterMeta
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthClusterMeta
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClusterMeta
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthClusterMeta
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthClusterMeta
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &gxredis.Instance{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipClusterMeta(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthClusterMeta
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Instances[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusterMeta
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
func (m *InstanceNameList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterMeta
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
			return fmt.Errorf("proto: InstanceNameList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceNameList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMeta
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
				return ErrInvalidLengthClusterMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusterMeta
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
func skipClusterMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterMeta
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
					return 0, ErrIntOverflowClusterMeta
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
					return 0, ErrIntOverflowClusterMeta
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
				return 0, ErrInvalidLengthClusterMeta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClusterMeta
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
				next, err := skipClusterMeta(dAtA[start:])
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
	ErrInvalidLengthClusterMeta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterMeta   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cluster_meta.proto", fileDescriptorClusterMeta) }

var fileDescriptorClusterMeta = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x29, 0x2d,
	0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98,
	0x05, 0xd1, 0x24, 0x25, 0x50, 0x94, 0x9a, 0x92, 0x59, 0x8c, 0x64, 0x8c, 0xd2, 0x0e, 0x46, 0x2e,
	0x6e, 0x67, 0x88, 0xe9, 0xbe, 0xa9, 0x25, 0x89, 0x42, 0x12, 0x5c, 0xec, 0x61, 0xa9, 0x45, 0xc5,
	0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x30, 0xae, 0x90, 0x1d, 0x17, 0xa7,
	0x67, 0x5e, 0x71, 0x49, 0x62, 0x5e, 0x72, 0x6a, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x82, 0x1e, 0xc8, 0x11, 0x7a, 0x48, 0xfa, 0xf5, 0xe0, 0x4a, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83,
	0x10, 0x5a, 0xa4, 0xfc, 0xb9, 0xf8, 0x50, 0x25, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0xc1,
	0xf6, 0x70, 0x06, 0x81, 0x98, 0x42, 0xea, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x82, 0x7a, 0xe9, 0x15, 0x60, 0x17, 0xc3, 0x8d, 0x0d, 0x82, 0xc8,
	0x5b, 0x31, 0x59, 0x30, 0x2a, 0xa9, 0x71, 0x09, 0xc0, 0x84, 0xfd, 0x12, 0x73, 0x53, 0x7d, 0x32,
	0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x40, 0xb4, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x98,
	0xed, 0x64, 0x72, 0xe2, 0xa1, 0x1c, 0xc3, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x3c,
	0x78, 0x28, 0xc7, 0xf8, 0xe1, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4,
	0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x1c, 0x3e, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x3e, 0x7c, 0x8c, 0x7c, 0x01, 0x00, 0x00,
}
