// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package evidencev1beta1

import (
	_ "github.com/verzth/cosmos-sdk/api/amino"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Equivocation                   protoreflect.MessageDescriptor
	fd_Equivocation_height            protoreflect.FieldDescriptor
	fd_Equivocation_time              protoreflect.FieldDescriptor
	fd_Equivocation_power             protoreflect.FieldDescriptor
	fd_Equivocation_consensus_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_evidence_v1beta1_evidence_proto_init()
	md_Equivocation = File_cosmos_evidence_v1beta1_evidence_proto.Messages().ByName("Equivocation")
	fd_Equivocation_height = md_Equivocation.Fields().ByName("height")
	fd_Equivocation_time = md_Equivocation.Fields().ByName("time")
	fd_Equivocation_power = md_Equivocation.Fields().ByName("power")
	fd_Equivocation_consensus_address = md_Equivocation.Fields().ByName("consensus_address")
}

var _ protoreflect.Message = (*fastReflection_Equivocation)(nil)

type fastReflection_Equivocation Equivocation

func (x *Equivocation) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Equivocation)(x)
}

func (x *Equivocation) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_evidence_v1beta1_evidence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Equivocation_messageType fastReflection_Equivocation_messageType
var _ protoreflect.MessageType = fastReflection_Equivocation_messageType{}

type fastReflection_Equivocation_messageType struct{}

func (x fastReflection_Equivocation_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Equivocation)(nil)
}
func (x fastReflection_Equivocation_messageType) New() protoreflect.Message {
	return new(fastReflection_Equivocation)
}
func (x fastReflection_Equivocation_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Equivocation
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Equivocation) Descriptor() protoreflect.MessageDescriptor {
	return md_Equivocation
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Equivocation) Type() protoreflect.MessageType {
	return _fastReflection_Equivocation_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Equivocation) New() protoreflect.Message {
	return new(fastReflection_Equivocation)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Equivocation) Interface() protoreflect.ProtoMessage {
	return (*Equivocation)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Equivocation) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Equivocation_height, value) {
			return
		}
	}
	if x.Time != nil {
		value := protoreflect.ValueOfMessage(x.Time.ProtoReflect())
		if !f(fd_Equivocation_time, value) {
			return
		}
	}
	if x.Power != int64(0) {
		value := protoreflect.ValueOfInt64(x.Power)
		if !f(fd_Equivocation_power, value) {
			return
		}
	}
	if x.ConsensusAddress != "" {
		value := protoreflect.ValueOfString(x.ConsensusAddress)
		if !f(fd_Equivocation_consensus_address, value) {
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
func (x *fastReflection_Equivocation) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.height":
		return x.Height != int64(0)
	case "cosmos.evidence.v1beta1.Equivocation.time":
		return x.Time != nil
	case "cosmos.evidence.v1beta1.Equivocation.power":
		return x.Power != int64(0)
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		return x.ConsensusAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Equivocation) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.height":
		x.Height = int64(0)
	case "cosmos.evidence.v1beta1.Equivocation.time":
		x.Time = nil
	case "cosmos.evidence.v1beta1.Equivocation.power":
		x.Power = int64(0)
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		x.ConsensusAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Equivocation) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "cosmos.evidence.v1beta1.Equivocation.time":
		value := x.Time
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.evidence.v1beta1.Equivocation.power":
		value := x.Power
		return protoreflect.ValueOfInt64(value)
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		value := x.ConsensusAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Equivocation) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.height":
		x.Height = value.Int()
	case "cosmos.evidence.v1beta1.Equivocation.time":
		x.Time = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.evidence.v1beta1.Equivocation.power":
		x.Power = value.Int()
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		x.ConsensusAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Equivocation) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.time":
		if x.Time == nil {
			x.Time = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Time.ProtoReflect())
	case "cosmos.evidence.v1beta1.Equivocation.height":
		panic(fmt.Errorf("field height of message cosmos.evidence.v1beta1.Equivocation is not mutable"))
	case "cosmos.evidence.v1beta1.Equivocation.power":
		panic(fmt.Errorf("field power of message cosmos.evidence.v1beta1.Equivocation is not mutable"))
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		panic(fmt.Errorf("field consensus_address of message cosmos.evidence.v1beta1.Equivocation is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Equivocation) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.Equivocation.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.evidence.v1beta1.Equivocation.time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.evidence.v1beta1.Equivocation.power":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.evidence.v1beta1.Equivocation.consensus_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.Equivocation"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.Equivocation does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Equivocation) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.evidence.v1beta1.Equivocation", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Equivocation) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Equivocation) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Equivocation) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Equivocation) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Equivocation)
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
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Time != nil {
			l = options.Size(x.Time)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Power != 0 {
			n += 1 + runtime.Sov(uint64(x.Power))
		}
		l = len(x.ConsensusAddress)
		if l > 0 {
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
		x := input.Message.Interface().(*Equivocation)
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
		if len(x.ConsensusAddress) > 0 {
			i -= len(x.ConsensusAddress)
			copy(dAtA[i:], x.ConsensusAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ConsensusAddress)))
			i--
			dAtA[i] = 0x22
		}
		if x.Power != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Power))
			i--
			dAtA[i] = 0x18
		}
		if x.Time != nil {
			encoded, err := options.Marshal(x.Time)
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
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*Equivocation)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Equivocation: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Equivocation: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
				if x.Time == nil {
					x.Time = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Time); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
				}
				x.Power = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Power |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ConsensusAddress = string(dAtA[iNdEx:postIndex])
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
// source: cosmos/evidence/v1beta1/evidence.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Equivocation implements the Evidence interface and defines evidence of double
// signing misbehavior.
type Equivocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// height is the equivocation height.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// time is the equivocation time.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// power is the equivocation validator power.
	Power int64 `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	// consensus_address is the equivocation validator consensus address.
	ConsensusAddress string `protobuf:"bytes,4,opt,name=consensus_address,json=consensusAddress,proto3" json:"consensus_address,omitempty"`
}

func (x *Equivocation) Reset() {
	*x = Equivocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_evidence_v1beta1_evidence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equivocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equivocation) ProtoMessage() {}

// Deprecated: Use Equivocation.ProtoReflect.Descriptor instead.
func (*Equivocation) Descriptor() ([]byte, []int) {
	return file_cosmos_evidence_v1beta1_evidence_proto_rawDescGZIP(), []int{0}
}

func (x *Equivocation) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Equivocation) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Equivocation) GetPower() int64 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Equivocation) GetConsensusAddress() string {
	if x != nil {
		return x.ConsensusAddress
	}
	return ""
}

var File_cosmos_evidence_v1beta1_evidence_proto protoreflect.FileDescriptor

var file_cosmos_evidence_v1beta1_evidence_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x45, 0x71, 0x75, 0x69, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x3d, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0d, 0xc8, 0xde, 0x1f, 0x00, 0x90,
	0xdf, 0x1f, 0x01, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x24, 0x88, 0xa0, 0x1f,
	0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x8a, 0xe7, 0xb0, 0x2a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x45, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0xe8, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0d, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x45, 0x58, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa8, 0xe2, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_evidence_v1beta1_evidence_proto_rawDescOnce sync.Once
	file_cosmos_evidence_v1beta1_evidence_proto_rawDescData = file_cosmos_evidence_v1beta1_evidence_proto_rawDesc
)

func file_cosmos_evidence_v1beta1_evidence_proto_rawDescGZIP() []byte {
	file_cosmos_evidence_v1beta1_evidence_proto_rawDescOnce.Do(func() {
		file_cosmos_evidence_v1beta1_evidence_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_evidence_v1beta1_evidence_proto_rawDescData)
	})
	return file_cosmos_evidence_v1beta1_evidence_proto_rawDescData
}

var file_cosmos_evidence_v1beta1_evidence_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_evidence_v1beta1_evidence_proto_goTypes = []interface{}{
	(*Equivocation)(nil),          // 0: cosmos.evidence.v1beta1.Equivocation
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_cosmos_evidence_v1beta1_evidence_proto_depIdxs = []int32{
	1, // 0: cosmos.evidence.v1beta1.Equivocation.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_evidence_v1beta1_evidence_proto_init() }
func file_cosmos_evidence_v1beta1_evidence_proto_init() {
	if File_cosmos_evidence_v1beta1_evidence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_evidence_v1beta1_evidence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equivocation); i {
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
			RawDescriptor: file_cosmos_evidence_v1beta1_evidence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_evidence_v1beta1_evidence_proto_goTypes,
		DependencyIndexes: file_cosmos_evidence_v1beta1_evidence_proto_depIdxs,
		MessageInfos:      file_cosmos_evidence_v1beta1_evidence_proto_msgTypes,
	}.Build()
	File_cosmos_evidence_v1beta1_evidence_proto = out.File
	file_cosmos_evidence_v1beta1_evidence_proto_rawDesc = nil
	file_cosmos_evidence_v1beta1_evidence_proto_goTypes = nil
	file_cosmos_evidence_v1beta1_evidence_proto_depIdxs = nil
}
