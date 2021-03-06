// Code generated by protoc-gen-go. DO NOT EDIT.
// source: winch/config/auth.proto

/*
Package winch_config is a generated protocol buffer package.

It is generated from these files:
	winch/config/auth.proto
	winch/config/mapper.proto

It has these top-level messages:
	AuthConfig
	AuthSource
	KubernetesAccess
	OIDCAccess
	DummyAccess
	MapperConfig
	Route
	DirectRoute
	RegexpRoute
*/
package winch_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / AuthConfig is the top level configuration message for a winch auth.
type AuthConfig struct {
	AuthSources []*AuthSource `protobuf:"bytes,1,rep,name=auth_sources,json=authSources" json:"auth_sources,omitempty"`
}

func (m *AuthConfig) Reset()                    { *m = AuthConfig{} }
func (m *AuthConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthConfig) ProtoMessage()               {}
func (*AuthConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthConfig) GetAuthSources() []*AuthSource {
	if m != nil {
		return m.AuthSources
	}
	return nil
}

// / AuthSource specifies the kind of the backend auth we need to inject on winch reqeuest.
type AuthSource struct {
	// name is an ID of auth source. It can be referenced inside winch routing.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*AuthSource_Dummy
	//	*AuthSource_Kube
	//	*AuthSource_Oidc
	Type isAuthSource_Type `protobuf_oneof:"type"`
}

func (m *AuthSource) Reset()                    { *m = AuthSource{} }
func (m *AuthSource) String() string            { return proto.CompactTextString(m) }
func (*AuthSource) ProtoMessage()               {}
func (*AuthSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isAuthSource_Type interface {
	isAuthSource_Type()
}

type AuthSource_Dummy struct {
	Dummy *DummyAccess `protobuf:"bytes,2,opt,name=dummy,oneof"`
}
type AuthSource_Kube struct {
	Kube *KubernetesAccess `protobuf:"bytes,3,opt,name=kube,oneof"`
}
type AuthSource_Oidc struct {
	Oidc *OIDCAccess `protobuf:"bytes,4,opt,name=oidc,oneof"`
}

func (*AuthSource_Dummy) isAuthSource_Type() {}
func (*AuthSource_Kube) isAuthSource_Type()  {}
func (*AuthSource_Oidc) isAuthSource_Type()  {}

func (m *AuthSource) GetType() isAuthSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AuthSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthSource) GetDummy() *DummyAccess {
	if x, ok := m.GetType().(*AuthSource_Dummy); ok {
		return x.Dummy
	}
	return nil
}

func (m *AuthSource) GetKube() *KubernetesAccess {
	if x, ok := m.GetType().(*AuthSource_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *AuthSource) GetOidc() *OIDCAccess {
	if x, ok := m.GetType().(*AuthSource_Oidc); ok {
		return x.Oidc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthSource_OneofMarshaler, _AuthSource_OneofUnmarshaler, _AuthSource_OneofSizer, []interface{}{
		(*AuthSource_Dummy)(nil),
		(*AuthSource_Kube)(nil),
		(*AuthSource_Oidc)(nil),
	}
}

func _AuthSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dummy); err != nil {
			return err
		}
	case *AuthSource_Kube:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *AuthSource_Oidc:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Oidc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthSource.Type has unexpected type %T", x)
	}
	return nil
}

func _AuthSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthSource)
	switch tag {
	case 2: // type.dummy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DummyAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Dummy{msg}
		return true, err
	case 3: // type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KubernetesAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Kube{msg}
		return true, err
	case 4: // type.oidc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OIDCAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Oidc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		s := proto.Size(x.Dummy)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Kube:
		s := proto.Size(x.Kube)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Oidc:
		s := proto.Size(x.Oidc)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / KubernetesAccess is an convenient way of specifying auth for backend. It grabs the data inside already used
// / ~/.kube/config (or any specified config path) and deducts the auth type based on that. NOTE that only these types are
// / supported:
// / - OIDC
type KubernetesAccess struct {
	// User to reference access credentials from.
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// By default ~/.kube/config as usual.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *KubernetesAccess) Reset()                    { *m = KubernetesAccess{} }
func (m *KubernetesAccess) String() string            { return proto.CompactTextString(m) }
func (*KubernetesAccess) ProtoMessage()               {}
func (*KubernetesAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KubernetesAccess) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *KubernetesAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type OIDCAccess struct {
	Provider string   `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Secret   string   `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	Scopes   []string `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	Path     string   `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
	// login_callback_path specified URL path for redirect URL to specify when doing OIDC login.
	// If empty login will be disabled which means in case of no refresh token or not valid one, error will be returned
	// thus not needing user interaction.
	LoginCallbackPath string `protobuf:"bytes,6,opt,name=login_callback_path,json=loginCallbackPath" json:"login_callback_path,omitempty"`
}

func (m *OIDCAccess) Reset()                    { *m = OIDCAccess{} }
func (m *OIDCAccess) String() string            { return proto.CompactTextString(m) }
func (*OIDCAccess) ProtoMessage()               {}
func (*OIDCAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OIDCAccess) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *OIDCAccess) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OIDCAccess) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OIDCAccess) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *OIDCAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *OIDCAccess) GetLoginCallbackPath() string {
	if m != nil {
		return m.LoginCallbackPath
	}
	return ""
}

// DummyAccess just directly passes specified value into auth header. If value is not specified it will return error.
type DummyAccess struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DummyAccess) Reset()                    { *m = DummyAccess{} }
func (m *DummyAccess) String() string            { return proto.CompactTextString(m) }
func (*DummyAccess) ProtoMessage()               {}
func (*DummyAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DummyAccess) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthConfig)(nil), "winch.config.AuthConfig")
	proto.RegisterType((*AuthSource)(nil), "winch.config.AuthSource")
	proto.RegisterType((*KubernetesAccess)(nil), "winch.config.KubernetesAccess")
	proto.RegisterType((*OIDCAccess)(nil), "winch.config.OIDCAccess")
	proto.RegisterType((*DummyAccess)(nil), "winch.config.DummyAccess")
}

func init() { proto.RegisterFile("winch/config/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x8e, 0x94, 0x40,
	0x10, 0x86, 0x65, 0x87, 0x21, 0x4e, 0xb1, 0x07, 0x6d, 0x8d, 0xe2, 0x1e, 0x56, 0xc2, 0x5e, 0xb8,
	0x2c, 0xc4, 0xd5, 0x78, 0xf1, 0xb4, 0x33, 0x7b, 0x70, 0xe2, 0x41, 0x83, 0x0f, 0x40, 0x9a, 0xa6,
	0x85, 0xce, 0x00, 0x4d, 0xe8, 0xee, 0x99, 0xec, 0xe3, 0xf9, 0x08, 0x3e, 0x81, 0x89, 0x4f, 0x62,
	0xa8, 0x1e, 0x61, 0x77, 0xbc, 0x55, 0xd7, 0xff, 0xd5, 0x5f, 0xfd, 0xd3, 0xc0, 0xeb, 0x83, 0xe8,
	0x58, 0x9d, 0x32, 0xd9, 0xfd, 0x10, 0x55, 0x4a, 0x8d, 0xae, 0x93, 0x7e, 0x90, 0x5a, 0x92, 0x73,
	0x14, 0x12, 0x2b, 0x5c, 0x7c, 0xac, 0x84, 0xae, 0x4d, 0x91, 0x30, 0xd9, 0xa6, 0xed, 0x41, 0xe8,
	0x9d, 0x3c, 0xa4, 0x95, 0xbc, 0x46, 0xf4, 0x7a, 0x4f, 0x1b, 0x51, 0x52, 0x2d, 0x07, 0x95, 0x4e,
	0xa5, 0x75, 0x89, 0xb6, 0x00, 0xb7, 0x46, 0xd7, 0x1b, 0x74, 0x21, 0x9f, 0xe0, 0x7c, 0xdc, 0x90,
	0x2b, 0x69, 0x06, 0xc6, 0x55, 0xe0, 0x84, 0x8b, 0xd8, 0xbf, 0x09, 0x92, 0x87, 0xab, 0x92, 0x91,
	0xff, 0x8e, 0x40, 0xe6, 0xd3, 0xa9, 0x56, 0xd1, 0x4f, 0xc7, 0x7a, 0xd9, 0x33, 0x21, 0xe0, 0x76,
	0xb4, 0xe5, 0x81, 0x13, 0x3a, 0xf1, 0x2a, 0xc3, 0x9a, 0xbc, 0x83, 0x65, 0x69, 0xda, 0xf6, 0x3e,
	0x38, 0x0b, 0x9d, 0xd8, 0xbf, 0x79, 0xf3, 0xd8, 0xf8, 0x6e, 0x94, 0x6e, 0x19, 0xe3, 0x4a, 0x7d,
	0x7e, 0x92, 0x59, 0x92, 0x7c, 0x00, 0x77, 0x67, 0x0a, 0x1e, 0x2c, 0x70, 0xe2, 0xf2, 0xf1, 0xc4,
	0x17, 0x53, 0xf0, 0xa1, 0xe3, 0x9a, 0xab, 0x69, 0x0c, 0x69, 0x92, 0x80, 0x2b, 0x45, 0xc9, 0x02,
	0x17, 0xa7, 0x4e, 0x02, 0x7c, 0xdd, 0xde, 0x6d, 0x66, 0x7e, 0xe4, 0xd6, 0x1e, 0xb8, 0xfa, 0xbe,
	0xe7, 0xd1, 0x1a, 0x9e, 0x9d, 0x7a, 0x92, 0x0b, 0x70, 0x8d, 0xe2, 0x83, 0x0d, 0xb2, 0xf6, 0xfe,
	0xfc, 0x7e, 0x7b, 0x16, 0x3a, 0x19, 0xf6, 0xc6, 0x90, 0x3d, 0xd5, 0x35, 0xe6, 0x59, 0x65, 0x58,
	0x47, 0xbf, 0x1c, 0x80, 0x79, 0x05, 0x89, 0xe0, 0x69, 0x3f, 0xc8, 0xbd, 0x28, 0xff, 0xb3, 0x98,
	0xfa, 0xe4, 0x0a, 0x56, 0xac, 0x11, 0xbc, 0xd3, 0xb9, 0x28, 0xad, 0xd7, 0x0c, 0x59, 0x61, 0x5b,
	0x92, 0x4b, 0xf0, 0x14, 0x67, 0x03, 0xd7, 0xf8, 0x2d, 0x66, 0xe2, 0xd8, 0x25, 0xaf, 0xc0, 0x53,
	0x4c, 0xf6, 0x5c, 0x05, 0x6e, 0xb8, 0x88, 0x57, 0xd9, 0xf1, 0x34, 0xdd, 0x71, 0x39, 0xdf, 0x91,
	0x24, 0xf0, 0xa2, 0x91, 0x95, 0xe8, 0x72, 0x46, 0x9b, 0xa6, 0xa0, 0x6c, 0x97, 0x23, 0xe2, 0x21,
	0xf2, 0x1c, 0xa5, 0xcd, 0x51, 0xf9, 0x36, 0x66, 0xba, 0x02, 0xff, 0xc1, 0xeb, 0x90, 0x97, 0xb0,
	0xdc, 0xd3, 0xc6, 0xfc, 0x7b, 0x5c, 0x7b, 0x28, 0x3c, 0xfc, 0xa5, 0xde, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x32, 0x95, 0xd3, 0xb3, 0x02, 0x00, 0x00,
}
