// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/type.proto

package ptype

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	source_context "google.golang.org/genproto/protobuf/source_context"
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

// The syntax in which a protocol buffer element is defined.
type Syntax int32

const (
	// Syntax `proto2`.
	Syntax_SYNTAX_PROTO2 Syntax = 0
	// Syntax `proto3`.
	Syntax_SYNTAX_PROTO3 Syntax = 1
)

var Syntax_name = map[int32]string{
	0: "SYNTAX_PROTO2",
	1: "SYNTAX_PROTO3",
}

var Syntax_value = map[string]int32{
	"SYNTAX_PROTO2": 0,
	"SYNTAX_PROTO3": 1,
}

func (x Syntax) String() string {
	return proto.EnumName(Syntax_name, int32(x))
}

func (Syntax) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{0}
}

// Basic field types.
type Field_Kind int32

const (
	// Field type unknown.
	Field_TYPE_UNKNOWN Field_Kind = 0
	// Field type double.
	Field_TYPE_DOUBLE Field_Kind = 1
	// Field type float.
	Field_TYPE_FLOAT Field_Kind = 2
	// Field type int64.
	Field_TYPE_INT64 Field_Kind = 3
	// Field type uint64.
	Field_TYPE_UINT64 Field_Kind = 4
	// Field type int32.
	Field_TYPE_INT32 Field_Kind = 5
	// Field type fixed64.
	Field_TYPE_FIXED64 Field_Kind = 6
	// Field type fixed32.
	Field_TYPE_FIXED32 Field_Kind = 7
	// Field type bool.
	Field_TYPE_BOOL Field_Kind = 8
	// Field type string.
	Field_TYPE_STRING Field_Kind = 9
	// Field type group. Proto2 syntax only, and deprecated.
	Field_TYPE_GROUP Field_Kind = 10
	// Field type message.
	Field_TYPE_MESSAGE Field_Kind = 11
	// Field type bytes.
	Field_TYPE_BYTES Field_Kind = 12
	// Field type uint32.
	Field_TYPE_UINT32 Field_Kind = 13
	// Field type enum.
	Field_TYPE_ENUM Field_Kind = 14
	// Field type sfixed32.
	Field_TYPE_SFIXED32 Field_Kind = 15
	// Field type sfixed64.
	Field_TYPE_SFIXED64 Field_Kind = 16
	// Field type sint32.
	Field_TYPE_SINT32 Field_Kind = 17
	// Field type sint64.
	Field_TYPE_SINT64 Field_Kind = 18
)

var Field_Kind_name = map[int32]string{
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}

var Field_Kind_value = map[string]int32{
	"TYPE_UNKNOWN":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Kind) String() string {
	return proto.EnumName(Field_Kind_name, int32(x))
}

func (Field_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{1, 0}
}

// Whether a field is optional, required, or repeated.
type Field_Cardinality int32

const (
	// For fields with unknown cardinality.
	Field_CARDINALITY_UNKNOWN Field_Cardinality = 0
	// For optional fields.
	Field_CARDINALITY_OPTIONAL Field_Cardinality = 1
	// For required fields. Proto2 syntax only.
	Field_CARDINALITY_REQUIRED Field_Cardinality = 2
	// For repeated fields.
	Field_CARDINALITY_REPEATED Field_Cardinality = 3
)

var Field_Cardinality_name = map[int32]string{
	0: "CARDINALITY_UNKNOWN",
	1: "CARDINALITY_OPTIONAL",
	2: "CARDINALITY_REQUIRED",
	3: "CARDINALITY_REPEATED",
}

var Field_Cardinality_value = map[string]int32{
	"CARDINALITY_UNKNOWN":  0,
	"CARDINALITY_OPTIONAL": 1,
	"CARDINALITY_REQUIRED": 2,
	"CARDINALITY_REPEATED": 3,
}

func (x Field_Cardinality) String() string {
	return proto.EnumName(Field_Cardinality_name, int32(x))
}

func (Field_Cardinality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{1, 1}
}

// A protocol buffer message type.
type Type struct {
	// The fully qualified message name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of fields.
	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	// The list of types appearing in `oneof` definitions in this type.
	Oneofs []string `protobuf:"bytes,3,rep,name=oneofs,proto3" json:"oneofs,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	// The source context.
	SourceContext *source_context.SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// The source syntax.
	Syntax               Syntax   `protobuf:"varint,6,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{0}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetOneofs() []string {
	if m != nil {
		return m.Oneofs
	}
	return nil
}

func (m *Type) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Type) GetSourceContext() *source_context.SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Type) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// A single field of a message type.
type Field struct {
	// The field type.
	Kind Field_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=google.protobuf.Field_Kind" json:"kind,omitempty"`
	// The field cardinality.
	Cardinality Field_Cardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=google.protobuf.Field_Cardinality" json:"cardinality,omitempty"`
	// The field number.
	Number int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// The field name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The field type URL, without the scheme, for message or enumeration
	// types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`.
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// The index of the field type in `Type.oneofs`, for message or enumeration
	// types. The first type has index 1; zero means the type is not in the list.
	OneofIndex int32 `protobuf:"varint,7,opt,name=oneof_index,json=oneofIndex,proto3" json:"oneof_index,omitempty"`
	// Whether to use alternative packed wire representation.
	Packed bool `protobuf:"varint,8,opt,name=packed,proto3" json:"packed,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,9,rep,name=options,proto3" json:"options,omitempty"`
	// The field JSON name.
	JsonName string `protobuf:"bytes,10,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	// The string value of the default value of this field. Proto2 syntax only.
	DefaultValue         string   `protobuf:"bytes,11,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{1}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetKind() Field_Kind {
	if m != nil {
		return m.Kind
	}
	return Field_TYPE_UNKNOWN
}

func (m *Field) GetCardinality() Field_Cardinality {
	if m != nil {
		return m.Cardinality
	}
	return Field_CARDINALITY_UNKNOWN
}

func (m *Field) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Field) GetOneofIndex() int32 {
	if m != nil {
		return m.OneofIndex
	}
	return 0
}

func (m *Field) GetPacked() bool {
	if m != nil {
		return m.Packed
	}
	return false
}

func (m *Field) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Field) GetJsonName() string {
	if m != nil {
		return m.JsonName
	}
	return ""
}

func (m *Field) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

// Enum type definition.
type Enum struct {
	// Enum type name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value definitions.
	Enumvalue []*EnumValue `protobuf:"bytes,2,rep,name=enumvalue,proto3" json:"enumvalue,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	// The source context.
	SourceContext *source_context.SourceContext `protobuf:"bytes,4,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// The source syntax.
	Syntax               Syntax   `protobuf:"varint,5,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Enum) Reset()         { *m = Enum{} }
func (m *Enum) String() string { return proto.CompactTextString(m) }
func (*Enum) ProtoMessage()    {}
func (*Enum) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{2}
}

func (m *Enum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enum.Unmarshal(m, b)
}
func (m *Enum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enum.Marshal(b, m, deterministic)
}
func (m *Enum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enum.Merge(m, src)
}
func (m *Enum) XXX_Size() int {
	return xxx_messageInfo_Enum.Size(m)
}
func (m *Enum) XXX_DiscardUnknown() {
	xxx_messageInfo_Enum.DiscardUnknown(m)
}

var xxx_messageInfo_Enum proto.InternalMessageInfo

func (m *Enum) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Enum) GetEnumvalue() []*EnumValue {
	if m != nil {
		return m.Enumvalue
	}
	return nil
}

func (m *Enum) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Enum) GetSourceContext() *source_context.SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Enum) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// Enum value definition.
type EnumValue struct {
	// Enum value name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value number.
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// Protocol buffer options.
	Options              []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{3}
}

func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (m *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(m, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnumValue) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *EnumValue) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// A protocol buffer option, which can be attached to a message, field,
// enumeration, etc.
type Option struct {
	// The option's name. For protobuf built-in options (options defined in
	// descriptor.proto), this is the short name. For example, `"map_entry"`.
	// For custom options, it should be the fully-qualified name. For example,
	// `"google.api.http"`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The option's value packed in an Any message. If the value is a primitive,
	// the corresponding wrapper type defined in google/protobuf/wrappers.proto
	// should be used. If the value is an enum, it should be stored as an int32
	// value using the google.protobuf.Int32Value type.
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd271cc1e348c538, []int{4}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Option) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.protobuf.Syntax", Syntax_name, Syntax_value)
	proto.RegisterEnum("google.protobuf.Field_Kind", Field_Kind_name, Field_Kind_value)
	proto.RegisterEnum("google.protobuf.Field_Cardinality", Field_Cardinality_name, Field_Cardinality_value)
	proto.RegisterType((*Type)(nil), "google.protobuf.Type")
	proto.RegisterType((*Field)(nil), "google.protobuf.Field")
	proto.RegisterType((*Enum)(nil), "google.protobuf.Enum")
	proto.RegisterType((*EnumValue)(nil), "google.protobuf.EnumValue")
	proto.RegisterType((*Option)(nil), "google.protobuf.Option")
}

func init() { proto.RegisterFile("google/protobuf/type.proto", fileDescriptor_dd271cc1e348c538) }

var fileDescriptor_dd271cc1e348c538 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x8e, 0xda, 0x56,
	0x14, 0x8e, 0x8d, 0xf1, 0xe0, 0xc3, 0xc0, 0xdc, 0xdc, 0x44, 0x89, 0x33, 0x91, 0x52, 0x44, 0xbb,
	0x40, 0x59, 0x80, 0x0a, 0xa3, 0x51, 0xa5, 0xae, 0x60, 0xf0, 0x50, 0x6b, 0x88, 0xed, 0x5e, 0x4c,
	0x93, 0xe9, 0x06, 0x79, 0xe0, 0x0e, 0x22, 0x31, 0xd7, 0x08, 0xdb, 0xed, 0xb0, 0xe8, 0x23, 0xf4,
	0x25, 0xba, 0xec, 0xba, 0x0f, 0xd1, 0x47, 0xea, 0xae, 0xd5, 0xbd, 0x06, 0x63, 0x7e, 0x2a, 0x4d,
	0x9b, 0xcd, 0x68, 0xce, 0xf7, 0x7d, 0xe7, 0xf7, 0x1e, 0x8e, 0xe1, 0x7c, 0x1a, 0x04, 0x53, 0x9f,
	0x36, 0x16, 0xcb, 0x20, 0x0a, 0xee, 0xe2, 0xfb, 0x46, 0xb4, 0x5a, 0xd0, 0xba, 0xb0, 0xf0, 0x59,
	0xc2, 0xd5, 0x37, 0xdc, 0xf9, 0xab, 0x7d, 0xb1, 0xc7, 0x56, 0x09, 0x7b, 0xfe, 0xd5, 0x3e, 0x15,
	0x06, 0xf1, 0x72, 0x4c, 0x47, 0xe3, 0x80, 0x45, 0xf4, 0x21, 0x4a, 0x54, 0xd5, 0x5f, 0x65, 0x50,
	0xdc, 0xd5, 0x82, 0x62, 0x0c, 0x0a, 0xf3, 0xe6, 0x54, 0x97, 0x2a, 0x52, 0x4d, 0x23, 0xe2, 0x7f,
	0x5c, 0x07, 0xf5, 0x7e, 0x46, 0xfd, 0x49, 0xa8, 0xcb, 0x95, 0x5c, 0xad, 0xd8, 0x7c, 0x51, 0xdf,
	0xcb, 0x5f, 0xbf, 0xe6, 0x34, 0x59, 0xab, 0xf0, 0x0b, 0x50, 0x03, 0x46, 0x83, 0xfb, 0x50, 0xcf,
	0x55, 0x72, 0x35, 0x8d, 0xac, 0x2d, 0xfc, 0x35, 0x9c, 0x04, 0x8b, 0x68, 0x16, 0xb0, 0x50, 0x57,
	0x44, 0xa0, 0x97, 0x07, 0x81, 0x6c, 0xc1, 0x93, 0x8d, 0x0e, 0x1b, 0x50, 0xde, 0xad, 0x57, 0xcf,
	0x57, 0xa4, 0x5a, 0xb1, 0xf9, 0xe6, 0xc0, 0x73, 0x20, 0x64, 0x57, 0x89, 0x8a, 0x94, 0xc2, 0xac,
	0x89, 0x1b, 0xa0, 0x86, 0x2b, 0x16, 0x79, 0x0f, 0xba, 0x5a, 0x91, 0x6a, 0xe5, 0x23, 0x89, 0x07,
	0x82, 0x26, 0x6b, 0x59, 0xf5, 0x0f, 0x15, 0xf2, 0xa2, 0x29, 0xdc, 0x00, 0xe5, 0xd3, 0x8c, 0x4d,
	0xc4, 0x40, 0xca, 0xcd, 0xd7, 0xc7, 0x5b, 0xaf, 0xdf, 0xcc, 0xd8, 0x84, 0x08, 0x21, 0xee, 0x42,
	0x71, 0xec, 0x2d, 0x27, 0x33, 0xe6, 0xf9, 0xb3, 0x68, 0xa5, 0xcb, 0xc2, 0xaf, 0xfa, 0x2f, 0x7e,
	0x57, 0x5b, 0x25, 0xc9, 0xba, 0xf1, 0x19, 0xb2, 0x78, 0x7e, 0x47, 0x97, 0x7a, 0xae, 0x22, 0xd5,
	0xf2, 0x64, 0x6d, 0xa5, 0xef, 0xa3, 0x64, 0xde, 0xe7, 0x15, 0x14, 0xf8, 0x72, 0x8c, 0xe2, 0xa5,
	0x2f, 0xfa, 0xd3, 0xc8, 0x09, 0xb7, 0x87, 0x4b, 0x1f, 0x7f, 0x01, 0x45, 0x31, 0xfc, 0xd1, 0x8c,
	0x4d, 0xe8, 0x83, 0x7e, 0x22, 0x62, 0x81, 0x80, 0x4c, 0x8e, 0xf0, 0x3c, 0x0b, 0x6f, 0xfc, 0x89,
	0x4e, 0xf4, 0x42, 0x45, 0xaa, 0x15, 0xc8, 0xda, 0xca, 0xbe, 0x95, 0xf6, 0xc8, 0xb7, 0x7a, 0x0d,
	0xda, 0xc7, 0x30, 0x60, 0x23, 0x51, 0x1f, 0x88, 0x3a, 0x0a, 0x1c, 0xb0, 0x78, 0x8d, 0x5f, 0x42,
	0x69, 0x42, 0xef, 0xbd, 0xd8, 0x8f, 0x46, 0x3f, 0x79, 0x7e, 0x4c, 0xf5, 0xa2, 0x10, 0x9c, 0xae,
	0xc1, 0x1f, 0x38, 0x56, 0xfd, 0x53, 0x06, 0x85, 0x4f, 0x12, 0x23, 0x38, 0x75, 0x6f, 0x1d, 0x63,
	0x34, 0xb4, 0x6e, 0x2c, 0xfb, 0xbd, 0x85, 0x9e, 0xe0, 0x33, 0x28, 0x0a, 0xa4, 0x6b, 0x0f, 0x3b,
	0x7d, 0x03, 0x49, 0xb8, 0x0c, 0x20, 0x80, 0xeb, 0xbe, 0xdd, 0x76, 0x91, 0x9c, 0xda, 0xa6, 0xe5,
	0x5e, 0x5e, 0xa0, 0x5c, 0xea, 0x30, 0x4c, 0x00, 0x25, 0x2b, 0x68, 0x35, 0x51, 0x3e, 0xcd, 0x71,
	0x6d, 0x7e, 0x30, 0xba, 0x97, 0x17, 0x48, 0xdd, 0x45, 0x5a, 0x4d, 0x74, 0x82, 0x4b, 0xa0, 0x09,
	0xa4, 0x63, 0xdb, 0x7d, 0x54, 0x48, 0x63, 0x0e, 0x5c, 0x62, 0x5a, 0x3d, 0xa4, 0xa5, 0x31, 0x7b,
	0xc4, 0x1e, 0x3a, 0x08, 0xd2, 0x08, 0xef, 0x8c, 0xc1, 0xa0, 0xdd, 0x33, 0x50, 0x31, 0x55, 0x74,
	0x6e, 0x5d, 0x63, 0x80, 0x4e, 0x77, 0xca, 0x6a, 0x35, 0x51, 0x29, 0x4d, 0x61, 0x58, 0xc3, 0x77,
	0xa8, 0x8c, 0x9f, 0x42, 0x29, 0x49, 0xb1, 0x29, 0xe2, 0x6c, 0x0f, 0xba, 0xbc, 0x40, 0x68, 0x5b,
	0x48, 0x12, 0xe5, 0xe9, 0x0e, 0x70, 0x79, 0x81, 0x70, 0x35, 0x82, 0x62, 0x66, 0xb7, 0xf0, 0x4b,
	0x78, 0x76, 0xd5, 0x26, 0x5d, 0xd3, 0x6a, 0xf7, 0x4d, 0xf7, 0x36, 0x33, 0x57, 0x1d, 0x9e, 0x67,
	0x09, 0xdb, 0x71, 0x4d, 0xdb, 0x6a, 0xf7, 0x91, 0xb4, 0xcf, 0x10, 0xe3, 0xfb, 0xa1, 0x49, 0x8c,
	0x2e, 0x92, 0x0f, 0x19, 0xc7, 0x68, 0xbb, 0x46, 0x17, 0xe5, 0xaa, 0x7f, 0x4b, 0xa0, 0x18, 0x2c,
	0x9e, 0x1f, 0x3d, 0x23, 0xdf, 0x80, 0x46, 0x59, 0x3c, 0x4f, 0x9e, 0x3f, 0xb9, 0x24, 0xe7, 0x07,
	0x4b, 0xc5, 0xbd, 0xc5, 0x32, 0x90, 0xad, 0x38, 0xbb, 0x8c, 0xb9, 0xff, 0x7d, 0x38, 0x94, 0xcf,
	0x3b, 0x1c, 0xf9, 0xc7, 0x1d, 0x8e, 0x8f, 0xa0, 0xa5, 0x2d, 0x1c, 0x9d, 0xc2, 0xf6, 0x87, 0x2d,
	0xef, 0xfc, 0xb0, 0xff, 0x7b, 0x8f, 0xd5, 0xef, 0x40, 0x4d, 0xa0, 0xa3, 0x89, 0xde, 0x42, 0x7e,
	0x33, 0x6a, 0xde, 0xf8, 0xf3, 0x83, 0x70, 0x6d, 0xb6, 0x22, 0x89, 0xe4, 0x6d, 0x1d, 0xd4, 0xa4,
	0x0f, 0xbe, 0x6c, 0x83, 0x5b, 0xcb, 0x6d, 0x7f, 0x18, 0x39, 0xc4, 0x76, 0xed, 0x26, 0x7a, 0xb2,
	0x0f, 0xb5, 0x90, 0xd4, 0xf9, 0x05, 0x9e, 0x8d, 0x83, 0xf9, 0x7e, 0xc4, 0x8e, 0xc6, 0x3f, 0x21,
	0x0e, 0xb7, 0x1c, 0xe9, 0xc7, 0xc6, 0x9a, 0x9d, 0x06, 0xbe, 0xc7, 0xa6, 0xf5, 0x60, 0x39, 0x6d,
	0x4c, 0x29, 0x13, 0xda, 0xed, 0xc7, 0x68, 0xc1, 0x0f, 0xd5, 0xb7, 0xe2, 0xef, 0x5f, 0x92, 0xf4,
	0x9b, 0x9c, 0xeb, 0x39, 0x9d, 0xdf, 0xe5, 0x37, 0xbd, 0xc4, 0xd5, 0xd9, 0x94, 0xfa, 0x9e, 0xfa,
	0xfe, 0x0d, 0x0b, 0x7e, 0x66, 0x3c, 0x41, 0x78, 0xa7, 0x0a, 0xff, 0xd6, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0x2b, 0xc0, 0xd8, 0x24, 0x07, 0x00, 0x00,
}
