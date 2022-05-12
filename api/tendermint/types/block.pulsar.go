// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Block             protoreflect.MessageDescriptor
	fd_Block_header      protoreflect.FieldDescriptor
	fd_Block_data        protoreflect.FieldDescriptor
	fd_Block_evidence    protoreflect.FieldDescriptor
	fd_Block_last_commit protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_block_proto_init()
	md_Block = File_tendermint_types_block_proto.Messages().ByName("Block")
	fd_Block_header = md_Block.Fields().ByName("header")
	fd_Block_data = md_Block.Fields().ByName("data")
	fd_Block_evidence = md_Block.Fields().ByName("evidence")
	fd_Block_last_commit = md_Block.Fields().ByName("last_commit")
}

var _ protoreflect.Message = (*fastReflection_Block)(nil)

type fastReflection_Block Block

func (x *Block) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Block)(x)
}

func (x *Block) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Block_messageType fastReflection_Block_messageType
var _ protoreflect.MessageType = fastReflection_Block_messageType{}

type fastReflection_Block_messageType struct{}

func (x fastReflection_Block_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Block)(nil)
}
func (x fastReflection_Block_messageType) New() protoreflect.Message {
	return new(fastReflection_Block)
}
func (x fastReflection_Block_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Block
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Block) Descriptor() protoreflect.MessageDescriptor {
	return md_Block
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Block) Type() protoreflect.MessageType {
	return _fastReflection_Block_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Block) New() protoreflect.Message {
	return new(fastReflection_Block)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Block) Interface() protoreflect.ProtoMessage {
	return (*Block)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Block) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Header != nil {
		value := protoreflect.ValueOfMessage(x.Header.ProtoReflect())
		if !f(fd_Block_header, value) {
			return
		}
	}
	if x.Data != nil {
		value := protoreflect.ValueOfMessage(x.Data.ProtoReflect())
		if !f(fd_Block_data, value) {
			return
		}
	}
	if x.Evidence != nil {
		value := protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
		if !f(fd_Block_evidence, value) {
			return
		}
	}
	if x.LastCommit != nil {
		value := protoreflect.ValueOfMessage(x.LastCommit.ProtoReflect())
		if !f(fd_Block_last_commit, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Block) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Block.header":
		return x.Header != nil
	case "tendermint.types.Block.data":
		return x.Data != nil
	case "tendermint.types.Block.evidence":
		return x.Evidence != nil
	case "tendermint.types.Block.last_commit":
		return x.LastCommit != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Block) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Block.header":
		x.Header = nil
	case "tendermint.types.Block.data":
		x.Data = nil
	case "tendermint.types.Block.evidence":
		x.Evidence = nil
	case "tendermint.types.Block.last_commit":
		x.LastCommit = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Block) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Block.header":
		value := x.Header
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Block.data":
		value := x.Data
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Block.evidence":
		value := x.Evidence
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Block.last_commit":
		value := x.LastCommit
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Block) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Block.header":
		x.Header = value.Message().Interface().(*Header)
	case "tendermint.types.Block.data":
		x.Data = value.Message().Interface().(*Data)
	case "tendermint.types.Block.evidence":
		x.Evidence = value.Message().Interface().(*EvidenceList)
	case "tendermint.types.Block.last_commit":
		x.LastCommit = value.Message().Interface().(*Commit)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Block) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Block.header":
		if x.Header == nil {
			x.Header = new(Header)
		}
		return protoreflect.ValueOfMessage(x.Header.ProtoReflect())
	case "tendermint.types.Block.data":
		if x.Data == nil {
			x.Data = new(Data)
		}
		return protoreflect.ValueOfMessage(x.Data.ProtoReflect())
	case "tendermint.types.Block.evidence":
		if x.Evidence == nil {
			x.Evidence = new(EvidenceList)
		}
		return protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
	case "tendermint.types.Block.last_commit":
		if x.LastCommit == nil {
			x.LastCommit = new(Commit)
		}
		return protoreflect.ValueOfMessage(x.LastCommit.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Block) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Block.header":
		m := new(Header)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Block.data":
		m := new(Data)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Block.evidence":
		m := new(EvidenceList)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Block.last_commit":
		m := new(Commit)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Block"))
		}
		panic(fmt.Errorf("message tendermint.types.Block does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Block) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Block", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Block) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Block) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Block) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Block) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Block)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Header != nil {
			l = options.Size(x.Header)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Data != nil {
			l = options.Size(x.Data)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Evidence != nil {
			l = options.Size(x.Evidence)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LastCommit != nil {
			l = options.Size(x.LastCommit)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Block)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.LastCommit != nil {
			encoded, err := options.Marshal(x.LastCommit)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.Evidence != nil {
			encoded, err := options.Marshal(x.Evidence)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Data != nil {
			encoded, err := options.Marshal(x.Data)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Header != nil {
			encoded, err := options.Marshal(x.Header)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Block)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Block: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Header == nil {
					x.Header = &Header{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Header); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Data == nil {
					x.Data = &Data{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Data); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Evidence == nil {
					x.Evidence = &EvidenceList{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Evidence); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastCommit", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.LastCommit == nil {
					x.LastCommit = &Commit{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastCommit); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: tendermint/types/block.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *Header       `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data       *Data         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Evidence   *EvidenceList `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence,omitempty"`
	LastCommit *Commit       `protobuf:"bytes,4,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_tendermint_types_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Block) GetEvidence() *EvidenceList {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *Block) GetLastCommit() *Commit {
	if x != nil {
		return x.LastCommit
	}
	return nil
}

var File_tendermint_types_block_proto protoreflect.FileDescriptor

var file_tendermint_types_block_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x08, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0xa6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42,
	0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x54, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x1c, 0x54,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_block_proto_rawDescOnce sync.Once
	file_tendermint_types_block_proto_rawDescData = file_tendermint_types_block_proto_rawDesc
)

func file_tendermint_types_block_proto_rawDescGZIP() []byte {
	file_tendermint_types_block_proto_rawDescOnce.Do(func() {
		file_tendermint_types_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_block_proto_rawDescData)
	})
	return file_tendermint_types_block_proto_rawDescData
}

var file_tendermint_types_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tendermint_types_block_proto_goTypes = []interface{}{
	(*Block)(nil),        // 0: tendermint.types.Block
	(*Header)(nil),       // 1: tendermint.types.Header
	(*Data)(nil),         // 2: tendermint.types.Data
	(*EvidenceList)(nil), // 3: tendermint.types.EvidenceList
	(*Commit)(nil),       // 4: tendermint.types.Commit
}
var file_tendermint_types_block_proto_depIdxs = []int32{
	1, // 0: tendermint.types.Block.header:type_name -> tendermint.types.Header
	2, // 1: tendermint.types.Block.data:type_name -> tendermint.types.Data
	3, // 2: tendermint.types.Block.evidence:type_name -> tendermint.types.EvidenceList
	4, // 3: tendermint.types.Block.last_commit:type_name -> tendermint.types.Commit
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tendermint_types_block_proto_init() }
func file_tendermint_types_block_proto_init() {
	if File_tendermint_types_block_proto != nil {
		return
	}
	file_tendermint_types_types_proto_init()
	file_tendermint_types_evidence_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_types_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_block_proto_goTypes,
		DependencyIndexes: file_tendermint_types_block_proto_depIdxs,
		MessageInfos:      file_tendermint_types_block_proto_msgTypes,
	}.Build()
	File_tendermint_types_block_proto = out.File
	file_tendermint_types_block_proto_rawDesc = nil
	file_tendermint_types_block_proto_goTypes = nil
	file_tendermint_types_block_proto_depIdxs = nil
}
