// Code generated by protoc-gen-go.
// source: contact.proto
// DO NOT EDIT!

package ricochet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Contact_Status int32

const (
	Contact_OFFLINE  Contact_Status = 0
	Contact_ONLINE   Contact_Status = 1
	Contact_REQUEST  Contact_Status = 2
	Contact_REJECTED Contact_Status = 3
)

var Contact_Status_name = map[int32]string{
	0: "OFFLINE",
	1: "ONLINE",
	2: "REQUEST",
	3: "REJECTED",
}
var Contact_Status_value = map[string]int32{
	"OFFLINE":  0,
	"ONLINE":   1,
	"REQUEST":  2,
	"REJECTED": 3,
}

func (x Contact_Status) String() string {
	return proto.EnumName(Contact_Status_name, int32(x))
}
func (Contact_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type ContactRequest_Direction int32

const (
	ContactRequest_INBOUND  ContactRequest_Direction = 0
	ContactRequest_OUTBOUND ContactRequest_Direction = 1
)

var ContactRequest_Direction_name = map[int32]string{
	0: "INBOUND",
	1: "OUTBOUND",
}
var ContactRequest_Direction_value = map[string]int32{
	"INBOUND":  0,
	"OUTBOUND": 1,
}

func (x ContactRequest_Direction) String() string {
	return proto.EnumName(ContactRequest_Direction_name, int32(x))
}
func (ContactRequest_Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type ContactEvent_Type int32

const (
	ContactEvent_NULL     ContactEvent_Type = 0
	ContactEvent_POPULATE ContactEvent_Type = 1
	ContactEvent_ADD      ContactEvent_Type = 2
	ContactEvent_UPDATE   ContactEvent_Type = 3
	ContactEvent_DELETE   ContactEvent_Type = 4
)

var ContactEvent_Type_name = map[int32]string{
	0: "NULL",
	1: "POPULATE",
	2: "ADD",
	3: "UPDATE",
	4: "DELETE",
}
var ContactEvent_Type_value = map[string]int32{
	"NULL":     0,
	"POPULATE": 1,
	"ADD":      2,
	"UPDATE":   3,
	"DELETE":   4,
}

func (x ContactEvent_Type) String() string {
	return proto.EnumName(ContactEvent_Type_name, int32(x))
}
func (ContactEvent_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type Contact struct {
	Id            int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address       string          `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Nickname      string          `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	WhenCreated   string          `protobuf:"bytes,4,opt,name=whenCreated" json:"whenCreated,omitempty"`
	LastConnected string          `protobuf:"bytes,5,opt,name=lastConnected" json:"lastConnected,omitempty"`
	Request       *ContactRequest `protobuf:"bytes,6,opt,name=request" json:"request,omitempty"`
	Status        Contact_Status  `protobuf:"varint,10,opt,name=status,enum=ricochet.Contact_Status" json:"status,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Contact) GetRequest() *ContactRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type ContactRequest struct {
	Direction ContactRequest_Direction `protobuf:"varint,1,opt,name=direction,enum=ricochet.ContactRequest_Direction" json:"direction,omitempty"`
	Address   string                   `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Nickname  string                   `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Text      string                   `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
}

func (m *ContactRequest) Reset()                    { *m = ContactRequest{} }
func (m *ContactRequest) String() string            { return proto.CompactTextString(m) }
func (*ContactRequest) ProtoMessage()               {}
func (*ContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type MonitorContactsRequest struct {
}

func (m *MonitorContactsRequest) Reset()                    { *m = MonitorContactsRequest{} }
func (m *MonitorContactsRequest) String() string            { return proto.CompactTextString(m) }
func (*MonitorContactsRequest) ProtoMessage()               {}
func (*MonitorContactsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ContactEvent struct {
	Type ContactEvent_Type `protobuf:"varint,1,opt,name=type,enum=ricochet.ContactEvent_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Subject:
	//	*ContactEvent_Contact
	//	*ContactEvent_Request
	Subject isContactEvent_Subject `protobuf_oneof:"subject"`
}

func (m *ContactEvent) Reset()                    { *m = ContactEvent{} }
func (m *ContactEvent) String() string            { return proto.CompactTextString(m) }
func (*ContactEvent) ProtoMessage()               {}
func (*ContactEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isContactEvent_Subject interface {
	isContactEvent_Subject()
}

type ContactEvent_Contact struct {
	Contact *Contact `protobuf:"bytes,2,opt,name=contact,oneof"`
}
type ContactEvent_Request struct {
	Request *ContactRequest `protobuf:"bytes,3,opt,name=request,oneof"`
}

func (*ContactEvent_Contact) isContactEvent_Subject() {}
func (*ContactEvent_Request) isContactEvent_Subject() {}

func (m *ContactEvent) GetSubject() isContactEvent_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *ContactEvent) GetContact() *Contact {
	if x, ok := m.GetSubject().(*ContactEvent_Contact); ok {
		return x.Contact
	}
	return nil
}

func (m *ContactEvent) GetRequest() *ContactRequest {
	if x, ok := m.GetSubject().(*ContactEvent_Request); ok {
		return x.Request
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ContactEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ContactEvent_OneofMarshaler, _ContactEvent_OneofUnmarshaler, _ContactEvent_OneofSizer, []interface{}{
		(*ContactEvent_Contact)(nil),
		(*ContactEvent_Request)(nil),
	}
}

func _ContactEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ContactEvent)
	// subject
	switch x := m.Subject.(type) {
	case *ContactEvent_Contact:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Contact); err != nil {
			return err
		}
	case *ContactEvent_Request:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ContactEvent.Subject has unexpected type %T", x)
	}
	return nil
}

func _ContactEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ContactEvent)
	switch tag {
	case 2: // subject.contact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Contact)
		err := b.DecodeMessage(msg)
		m.Subject = &ContactEvent_Contact{msg}
		return true, err
	case 3: // subject.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContactRequest)
		err := b.DecodeMessage(msg)
		m.Subject = &ContactEvent_Request{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ContactEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ContactEvent)
	// subject
	switch x := m.Subject.(type) {
	case *ContactEvent_Contact:
		s := proto.Size(x.Contact)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ContactEvent_Request:
		s := proto.Size(x.Request)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AddContactReply struct {
}

func (m *AddContactReply) Reset()                    { *m = AddContactReply{} }
func (m *AddContactReply) String() string            { return proto.CompactTextString(m) }
func (*AddContactReply) ProtoMessage()               {}
func (*AddContactReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type DeleteContactRequest struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *DeleteContactRequest) Reset()                    { *m = DeleteContactRequest{} }
func (m *DeleteContactRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactRequest) ProtoMessage()               {}
func (*DeleteContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DeleteContactReply struct {
}

func (m *DeleteContactReply) Reset()                    { *m = DeleteContactReply{} }
func (m *DeleteContactReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactReply) ProtoMessage()               {}
func (*DeleteContactReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type RejectInboundRequestReply struct {
}

func (m *RejectInboundRequestReply) Reset()                    { *m = RejectInboundRequestReply{} }
func (m *RejectInboundRequestReply) String() string            { return proto.CompactTextString(m) }
func (*RejectInboundRequestReply) ProtoMessage()               {}
func (*RejectInboundRequestReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*Contact)(nil), "ricochet.Contact")
	proto.RegisterType((*ContactRequest)(nil), "ricochet.ContactRequest")
	proto.RegisterType((*MonitorContactsRequest)(nil), "ricochet.MonitorContactsRequest")
	proto.RegisterType((*ContactEvent)(nil), "ricochet.ContactEvent")
	proto.RegisterType((*AddContactReply)(nil), "ricochet.AddContactReply")
	proto.RegisterType((*DeleteContactRequest)(nil), "ricochet.DeleteContactRequest")
	proto.RegisterType((*DeleteContactReply)(nil), "ricochet.DeleteContactReply")
	proto.RegisterType((*RejectInboundRequestReply)(nil), "ricochet.RejectInboundRequestReply")
	proto.RegisterEnum("ricochet.Contact_Status", Contact_Status_name, Contact_Status_value)
	proto.RegisterEnum("ricochet.ContactRequest_Direction", ContactRequest_Direction_name, ContactRequest_Direction_value)
	proto.RegisterEnum("ricochet.ContactEvent_Type", ContactEvent_Type_name, ContactEvent_Type_value)
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x9a, 0x2c, 0x69, 0x6f, 0xb7, 0xd2, 0x59, 0x13, 0x0a, 0xec, 0x65, 0xb2, 0x10, 0xe2,
	0x85, 0x80, 0x0a, 0x8f, 0x48, 0xac, 0x6b, 0x32, 0x51, 0x14, 0x92, 0xe2, 0x25, 0x1f, 0x90, 0x26,
	0x96, 0x16, 0x28, 0x4e, 0x49, 0x5c, 0x60, 0xdf, 0xc6, 0x2b, 0x1f, 0xc4, 0x27, 0x60, 0x3b, 0x4e,
	0xc7, 0x36, 0x40, 0x68, 0x6f, 0xd7, 0xf7, 0x9c, 0xeb, 0x7b, 0xcf, 0x3d, 0x36, 0xec, 0xe7, 0x15,
	0xe3, 0x59, 0xce, 0xbd, 0x75, 0x5d, 0xf1, 0x0a, 0xf5, 0xeb, 0x32, 0xaf, 0xf2, 0x0b, 0xca, 0xf1,
	0xf7, 0x1e, 0x38, 0xb3, 0x16, 0x43, 0x23, 0xe8, 0x95, 0x85, 0x6b, 0x1c, 0x1b, 0x4f, 0x76, 0x89,
	0x88, 0x90, 0x0b, 0x4e, 0x56, 0x14, 0x35, 0x6d, 0x1a, 0xb7, 0x27, 0x92, 0x03, 0xd2, 0x1d, 0xd1,
	0x43, 0xe8, 0xb3, 0x32, 0xff, 0xc8, 0xb2, 0x4f, 0xd4, 0x35, 0x15, 0xb4, 0x3d, 0xa3, 0x63, 0x18,
	0x7e, 0xbd, 0xa0, 0x6c, 0x56, 0xd3, 0x8c, 0xd3, 0xc2, 0xb5, 0x14, 0xfc, 0x7b, 0x0a, 0x3d, 0x82,
	0xfd, 0x55, 0xd6, 0x70, 0xd1, 0x96, 0xd1, 0x5c, 0x72, 0x76, 0x15, 0xe7, 0x7a, 0x12, 0x4d, 0xc0,
	0xa9, 0xe9, 0xe7, 0x0d, 0x6d, 0xb8, 0x6b, 0x0b, 0x7c, 0x38, 0x71, 0xbd, 0x6e, 0x6a, 0x4f, 0x4f,
	0x4c, 0x5a, 0x9c, 0x74, 0x44, 0xf4, 0x1c, 0xec, 0x86, 0x67, 0x7c, 0xd3, 0xb8, 0x20, 0x4a, 0x46,
	0x7f, 0x28, 0xf1, 0xce, 0x15, 0x4e, 0x34, 0x0f, 0xbf, 0x02, 0xbb, 0xcd, 0xa0, 0x21, 0x38, 0xf1,
	0xd9, 0x59, 0x38, 0x8f, 0x82, 0xf1, 0x0e, 0x02, 0xb0, 0xe3, 0x48, 0xc5, 0x86, 0x04, 0x48, 0xf0,
	0x3e, 0x0d, 0xce, 0x93, 0x71, 0x0f, 0xed, 0x41, 0x9f, 0x04, 0x6f, 0x83, 0x59, 0x12, 0xf8, 0x63,
	0x13, 0xff, 0x30, 0x60, 0x74, 0x7d, 0x16, 0x74, 0x02, 0x83, 0xa2, 0xac, 0x85, 0x84, 0xb2, 0x62,
	0x6a, 0x97, 0xa3, 0x09, 0xfe, 0xdb, 0xe0, 0x9e, 0xdf, 0x31, 0xc9, 0x55, 0xd1, 0x1d, 0xd7, 0x8e,
	0xc0, 0xe2, 0xf4, 0x1b, 0xd7, 0xfb, 0x56, 0x31, 0x7e, 0x0c, 0x83, 0x6d, 0x07, 0x29, 0x63, 0x1e,
	0x9d, 0xc6, 0x69, 0xe4, 0x0b, 0x7d, 0x42, 0x46, 0x9c, 0x26, 0xed, 0xc9, 0xc0, 0x2e, 0xdc, 0x7f,
	0x57, 0xb1, 0x92, 0x57, 0xb5, 0x9e, 0xaf, 0xd1, 0x03, 0xe2, 0x9f, 0x06, 0xec, 0xe9, 0x5c, 0xf0,
	0x85, 0x32, 0x8e, 0x9e, 0x89, 0x36, 0x97, 0x6b, 0xaa, 0x95, 0x1d, 0xdd, 0x52, 0xa6, 0x58, 0x5e,
	0x22, 0x28, 0x44, 0x11, 0xd1, 0x53, 0x70, 0xf4, 0xdb, 0x53, 0x6a, 0x86, 0x93, 0x83, 0x5b, 0x35,
	0x6f, 0x76, 0x48, 0xc7, 0x41, 0x2f, 0xaf, 0x5c, 0x37, 0xff, 0xed, 0xba, 0xac, 0xd2, 0x54, 0xfc,
	0x1a, 0x2c, 0xd9, 0x12, 0xf5, 0xc1, 0x8a, 0xd2, 0x30, 0x6c, 0x05, 0x2e, 0xe2, 0x45, 0x1a, 0x4e,
	0x13, 0x69, 0xa1, 0x03, 0xe6, 0xd4, 0xf7, 0x85, 0x7d, 0xc2, 0xd7, 0x74, 0xe1, 0xcb, 0xa4, 0x29,
	0x63, 0x3f, 0x08, 0x03, 0x11, 0x5b, 0xa7, 0x03, 0x70, 0x9a, 0xcd, 0xf2, 0x83, 0x58, 0x15, 0x3e,
	0x80, 0x7b, 0xd3, 0xa2, 0xd8, 0xf6, 0x5a, 0xaf, 0x2e, 0xf1, 0x09, 0x1c, 0xfa, 0x74, 0x45, 0x39,
	0xbd, 0xe1, 0xf5, 0x7f, 0x7f, 0x18, 0x7c, 0x08, 0xe8, 0xc6, 0x0d, 0xf2, 0xde, 0x23, 0x78, 0x40,
	0xa8, 0x6c, 0x3a, 0x67, 0xcb, 0x6a, 0xc3, 0x8a, 0xee, 0x3d, 0x4b, 0x70, 0x69, 0xab, 0xaf, 0xfa,
	0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xab, 0x53, 0xdb, 0xbb, 0x03, 0x00, 0x00,
}
