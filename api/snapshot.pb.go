// Code generated by protoc-gen-gogo.
// source: snapshot.proto
// DO NOT EDIT!

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Snapshot_Version int32

const (
	// V0 is the initial version of the StoreSnapshot message.
	Snapshot_V0 Snapshot_Version = 0
)

var Snapshot_Version_name = map[int32]string{
	0: "V0",
}
var Snapshot_Version_value = map[string]int32{
	"V0": 0,
}

func (x Snapshot_Version) String() string {
	return proto.EnumName(Snapshot_Version_name, int32(x))
}
func (Snapshot_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{2, 0} }

// StoreSnapshot is used to store snapshots of the store.
type StoreSnapshot struct {
	Nodes    []*Node    `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Services []*Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
	Networks []*Network `protobuf:"bytes,3,rep,name=networks" json:"networks,omitempty"`
	Tasks    []*Task    `protobuf:"bytes,4,rep,name=tasks" json:"tasks,omitempty"`
	Clusters []*Cluster `protobuf:"bytes,5,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *StoreSnapshot) Reset()                    { *m = StoreSnapshot{} }
func (*StoreSnapshot) ProtoMessage()               {}
func (*StoreSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{0} }

// ClusterSnapshot stores cluster membership information in snapshots.
type ClusterSnapshot struct {
	Members []*RaftMember `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
	Removed []uint64      `protobuf:"varint,2,rep,name=removed" json:"removed,omitempty"`
}

func (m *ClusterSnapshot) Reset()                    { *m = ClusterSnapshot{} }
func (*ClusterSnapshot) ProtoMessage()               {}
func (*ClusterSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{1} }

type Snapshot struct {
	Version    Snapshot_Version `protobuf:"varint,1,opt,name=version,proto3,enum=docker.cluster.api.Snapshot_Version" json:"version,omitempty"`
	Membership ClusterSnapshot  `protobuf:"bytes,2,opt,name=membership" json:"membership"`
	Store      StoreSnapshot    `protobuf:"bytes,3,opt,name=store" json:"store"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{2} }

func init() {
	proto.RegisterType((*StoreSnapshot)(nil), "docker.cluster.api.StoreSnapshot")
	proto.RegisterType((*ClusterSnapshot)(nil), "docker.cluster.api.ClusterSnapshot")
	proto.RegisterType((*Snapshot)(nil), "docker.cluster.api.Snapshot")
	proto.RegisterEnum("docker.cluster.api.Snapshot_Version", Snapshot_Version_name, Snapshot_Version_value)
}

func (m *StoreSnapshot) Copy() *StoreSnapshot {
	if m == nil {
		return nil
	}

	o := &StoreSnapshot{}

	if m.Nodes != nil {
		o.Nodes = make([]*Node, 0, len(m.Nodes))
		for _, v := range m.Nodes {
			o.Nodes = append(o.Nodes, v.Copy())
		}
	}

	if m.Services != nil {
		o.Services = make([]*Service, 0, len(m.Services))
		for _, v := range m.Services {
			o.Services = append(o.Services, v.Copy())
		}
	}

	if m.Networks != nil {
		o.Networks = make([]*Network, 0, len(m.Networks))
		for _, v := range m.Networks {
			o.Networks = append(o.Networks, v.Copy())
		}
	}

	if m.Tasks != nil {
		o.Tasks = make([]*Task, 0, len(m.Tasks))
		for _, v := range m.Tasks {
			o.Tasks = append(o.Tasks, v.Copy())
		}
	}

	if m.Clusters != nil {
		o.Clusters = make([]*Cluster, 0, len(m.Clusters))
		for _, v := range m.Clusters {
			o.Clusters = append(o.Clusters, v.Copy())
		}
	}

	return o
}

func (m *ClusterSnapshot) Copy() *ClusterSnapshot {
	if m == nil {
		return nil
	}

	o := &ClusterSnapshot{}

	if m.Members != nil {
		o.Members = make([]*RaftMember, 0, len(m.Members))
		for _, v := range m.Members {
			o.Members = append(o.Members, v.Copy())
		}
	}

	if m.Removed != nil {
		o.Removed = make([]uint64, 0, len(m.Removed))
		for _, v := range m.Removed {
			o.Removed = append(o.Removed, v)
		}
	}

	return o
}

func (m *Snapshot) Copy() *Snapshot {
	if m == nil {
		return nil
	}

	o := &Snapshot{
		Version:    m.Version,
		Membership: *m.Membership.Copy(),
		Store:      *m.Store.Copy(),
	}

	return o
}

func (this *StoreSnapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&api.StoreSnapshot{")
	if this.Nodes != nil {
		s = append(s, "Nodes: "+fmt.Sprintf("%#v", this.Nodes)+",\n")
	}
	if this.Services != nil {
		s = append(s, "Services: "+fmt.Sprintf("%#v", this.Services)+",\n")
	}
	if this.Networks != nil {
		s = append(s, "Networks: "+fmt.Sprintf("%#v", this.Networks)+",\n")
	}
	if this.Tasks != nil {
		s = append(s, "Tasks: "+fmt.Sprintf("%#v", this.Tasks)+",\n")
	}
	if this.Clusters != nil {
		s = append(s, "Clusters: "+fmt.Sprintf("%#v", this.Clusters)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterSnapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&api.ClusterSnapshot{")
	if this.Members != nil {
		s = append(s, "Members: "+fmt.Sprintf("%#v", this.Members)+",\n")
	}
	s = append(s, "Removed: "+fmt.Sprintf("%#v", this.Removed)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Snapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&api.Snapshot{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Membership: "+strings.Replace(this.Membership.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Store: "+strings.Replace(this.Store.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSnapshot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringSnapshot(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *StoreSnapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreSnapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			data[i] = 0xa
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Services) > 0 {
		for _, msg := range m.Services {
			data[i] = 0x12
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Networks) > 0 {
		for _, msg := range m.Networks {
			data[i] = 0x1a
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tasks) > 0 {
		for _, msg := range m.Tasks {
			data[i] = 0x22
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Clusters) > 0 {
		for _, msg := range m.Clusters {
			data[i] = 0x2a
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterSnapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClusterSnapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			data[i] = 0xa
			i++
			i = encodeVarintSnapshot(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Removed) > 0 {
		for _, num := range m.Removed {
			data[i] = 0x10
			i++
			i = encodeVarintSnapshot(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *Snapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Snapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintSnapshot(data, i, uint64(m.Version))
	}
	data[i] = 0x12
	i++
	i = encodeVarintSnapshot(data, i, uint64(m.Membership.Size()))
	n1, err := m.Membership.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintSnapshot(data, i, uint64(m.Store.Size()))
	n2, err := m.Store.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeFixed64Snapshot(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Snapshot(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSnapshot(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

func (m *StoreSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	return n
}

func (m *ClusterSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		for _, e := range m.Removed {
			n += 1 + sovSnapshot(uint64(e))
		}
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovSnapshot(uint64(m.Version))
	}
	l = m.Membership.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	l = m.Store.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func sovSnapshot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StoreSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreSnapshot{`,
		`Nodes:` + strings.Replace(fmt.Sprintf("%v", this.Nodes), "Node", "Node", 1) + `,`,
		`Services:` + strings.Replace(fmt.Sprintf("%v", this.Services), "Service", "Service", 1) + `,`,
		`Networks:` + strings.Replace(fmt.Sprintf("%v", this.Networks), "Network", "Network", 1) + `,`,
		`Tasks:` + strings.Replace(fmt.Sprintf("%v", this.Tasks), "Task", "Task", 1) + `,`,
		`Clusters:` + strings.Replace(fmt.Sprintf("%v", this.Clusters), "Cluster", "Cluster", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterSnapshot{`,
		`Members:` + strings.Replace(fmt.Sprintf("%v", this.Members), "RaftMember", "RaftMember", 1) + `,`,
		`Removed:` + fmt.Sprintf("%v", this.Removed) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Snapshot{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Membership:` + strings.Replace(strings.Replace(this.Membership.String(), "ClusterSnapshot", "ClusterSnapshot", 1), `&`, ``, 1) + `,`,
		`Store:` + strings.Replace(strings.Replace(this.Store.String(), "StoreSnapshot", "StoreSnapshot", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSnapshot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StoreSnapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Networks = append(m.Networks, &Network{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &Task{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &Cluster{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *ClusterSnapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &RaftMember{})
			if err := m.Members[len(m.Members)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Removed = append(m.Removed, v)
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *Snapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (Snapshot_Version(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Membership.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Store.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSnapshot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSnapshot
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipSnapshot(data[start:])
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
	ErrInvalidLengthSnapshot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorSnapshot = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x4e, 0xe3, 0x40,
	0x10, 0xc6, 0x63, 0xe7, 0x8f, 0xa3, 0x89, 0x92, 0xbb, 0x5b, 0x5d, 0x61, 0xe5, 0x4e, 0x06, 0x0c,
	0x45, 0x2a, 0x83, 0x42, 0x01, 0x0d, 0x14, 0xa1, 0xa2, 0x20, 0xc5, 0x06, 0x45, 0xb4, 0x8e, 0xb3,
	0x24, 0x26, 0xc4, 0x6b, 0x79, 0x97, 0x20, 0x3a, 0x9e, 0x83, 0x27, 0x4a, 0x49, 0x49, 0x85, 0x08,
	0x0d, 0x2d, 0x8f, 0xc0, 0xd8, 0x6b, 0x5b, 0x91, 0x70, 0x28, 0x46, 0x9e, 0x5d, 0xff, 0xbe, 0xf9,
	0xc6, 0xe3, 0x81, 0x96, 0x08, 0xdc, 0x50, 0x4c, 0xb9, 0x74, 0xc2, 0x88, 0x4b, 0x4e, 0xc8, 0x98,
	0x7b, 0x33, 0x16, 0x39, 0xde, 0xed, 0x9d, 0x90, 0xf8, 0x74, 0x43, 0xbf, 0xdd, 0xe4, 0xa3, 0x1b,
	0xe6, 0x49, 0xa1, 0x90, 0x76, 0x43, 0x3e, 0x84, 0x2c, 0x3b, 0xfc, 0x9d, 0xf0, 0x09, 0x4f, 0xd2,
	0xfd, 0x38, 0x53, 0xb7, 0xf6, 0x93, 0x0e, 0xcd, 0x81, 0xe4, 0x11, 0x1b, 0xa4, 0xd5, 0x89, 0x03,
	0xd5, 0x80, 0x8f, 0x99, 0x30, 0xb5, 0xed, 0x72, 0xa7, 0xd1, 0x35, 0x9d, 0xef, 0x3e, 0x4e, 0x1f,
	0x01, 0xaa, 0x30, 0x72, 0x04, 0x75, 0xc1, 0xa2, 0x85, 0xef, 0xa1, 0x44, 0x4f, 0x24, 0xff, 0x8a,
	0x24, 0x03, 0xc5, 0xd0, 0x1c, 0x8e, 0x85, 0x01, 0x93, 0xf7, 0x3c, 0x9a, 0x09, 0xb3, 0xbc, 0x59,
	0xd8, 0x57, 0x0c, 0xcd, 0xe1, 0xb8, 0x43, 0xe9, 0x0a, 0x54, 0x55, 0x36, 0x77, 0x78, 0x89, 0x00,
	0x55, 0x58, 0x6c, 0x94, 0xbe, 0x12, 0x66, 0x75, 0xb3, 0xd1, 0x99, 0xca, 0x69, 0x0e, 0xdb, 0x0c,
	0x7e, 0xa5, 0x97, 0xf9, 0x74, 0x8e, 0xc1, 0x98, 0xb3, 0xf9, 0x28, 0x2e, 0xa5, 0xe6, 0x63, 0x15,
	0x95, 0xa2, 0xee, 0xb5, 0xbc, 0x48, 0x30, 0x9a, 0xe1, 0xc4, 0x04, 0x23, 0x62, 0x73, 0xbe, 0x60,
	0xe3, 0x64, 0x4c, 0x15, 0x9a, 0x1d, 0xed, 0x0f, 0x0d, 0xea, 0xb9, 0xc1, 0x29, 0x18, 0x0b, 0xc4,
	0x7d, 0x1e, 0xa0, 0x81, 0xd6, 0x69, 0x75, 0xf7, 0x0a, 0xa7, 0x99, 0xed, 0xc2, 0x50, 0xb1, 0x34,
	0x13, 0x91, 0x73, 0x80, 0xd4, 0x71, 0xea, 0x87, 0xe8, 0xa4, 0x61, 0x8f, 0xbb, 0x3f, 0x7c, 0x6e,
	0x56, 0xa9, 0x57, 0x59, 0xbe, 0x6e, 0x95, 0xe8, 0x9a, 0x98, 0x9c, 0x40, 0x55, 0xc4, 0xab, 0x81,
	0x7f, 0x27, 0xae, 0xb2, 0x53, 0xd8, 0xc8, 0xfa, 0xee, 0xa4, 0x35, 0x94, 0xca, 0xfe, 0x03, 0x46,
	0xda, 0x1d, 0xa9, 0x81, 0x3e, 0x3c, 0xf8, 0x5d, 0xea, 0xfd, 0x5f, 0xae, 0xac, 0xd2, 0x0b, 0xc6,
	0xe7, 0xca, 0xd2, 0x1e, 0xdf, 0x2d, 0x6d, 0x89, 0xf1, 0x8c, 0xf1, 0x86, 0x71, 0xa5, 0x8f, 0x6a,
	0xc9, 0x52, 0x1e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x11, 0xf6, 0x39, 0x1c, 0xec, 0x02, 0x00,
	0x00,
}
