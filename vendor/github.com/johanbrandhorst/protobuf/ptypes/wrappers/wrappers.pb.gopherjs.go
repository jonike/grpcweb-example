// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: wrappers/wrappers.proto

/*
	Package wrappers is a generated protocol buffer package.

	It is generated from these files:
		wrappers/wrappers.proto

	It has these top-level messages:
		DoubleValue
		FloatValue
		Int64Value
		UInt64Value
		Int32Value
		UInt32Value
		BoolValue
		StringValue
		BytesValue
*/
package wrappers

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value float64
}

// GetValue gets the Value of the DoubleValue.
func (m *DoubleValue) GetValue() (x float64) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals DoubleValue to the provided writer.
func (m *DoubleValue) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteFloat64(1, m.Value)
	}

	return
}

// Marshal marshals DoubleValue to a slice of bytes.
func (m *DoubleValue) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a DoubleValue from the provided reader.
func (m *DoubleValue) UnmarshalFromReader(reader jspb.Reader) *DoubleValue {
	for reader.Next() {
		if m == nil {
			m = &DoubleValue{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadFloat64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a DoubleValue from a slice of bytes.
func (m *DoubleValue) Unmarshal(rawBytes []byte) (*DoubleValue, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value float32
}

// GetValue gets the Value of the FloatValue.
func (m *FloatValue) GetValue() (x float32) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals FloatValue to the provided writer.
func (m *FloatValue) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteFloat32(1, m.Value)
	}

	return
}

// Marshal marshals FloatValue to a slice of bytes.
func (m *FloatValue) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a FloatValue from the provided reader.
func (m *FloatValue) UnmarshalFromReader(reader jspb.Reader) *FloatValue {
	for reader.Next() {
		if m == nil {
			m = &FloatValue{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadFloat32()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a FloatValue from a slice of bytes.
func (m *FloatValue) Unmarshal(rawBytes []byte) (*FloatValue, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value int64
}

// GetValue gets the Value of the Int64Value.
func (m *Int64Value) GetValue() (x int64) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals Int64Value to the provided writer.
func (m *Int64Value) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteInt64(1, m.Value)
	}

	return
}

// Marshal marshals Int64Value to a slice of bytes.
func (m *Int64Value) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Int64Value from the provided reader.
func (m *Int64Value) UnmarshalFromReader(reader jspb.Reader) *Int64Value {
	for reader.Next() {
		if m == nil {
			m = &Int64Value{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadInt64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Int64Value from a slice of bytes.
func (m *Int64Value) Unmarshal(rawBytes []byte) (*Int64Value, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value uint64
}

// GetValue gets the Value of the UInt64Value.
func (m *UInt64Value) GetValue() (x uint64) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals UInt64Value to the provided writer.
func (m *UInt64Value) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteUint64(1, m.Value)
	}

	return
}

// Marshal marshals UInt64Value to a slice of bytes.
func (m *UInt64Value) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a UInt64Value from the provided reader.
func (m *UInt64Value) UnmarshalFromReader(reader jspb.Reader) *UInt64Value {
	for reader.Next() {
		if m == nil {
			m = &UInt64Value{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadUint64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a UInt64Value from a slice of bytes.
func (m *UInt64Value) Unmarshal(rawBytes []byte) (*UInt64Value, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value int32
}

// GetValue gets the Value of the Int32Value.
func (m *Int32Value) GetValue() (x int32) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals Int32Value to the provided writer.
func (m *Int32Value) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteInt32(1, m.Value)
	}

	return
}

// Marshal marshals Int32Value to a slice of bytes.
func (m *Int32Value) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Int32Value from the provided reader.
func (m *Int32Value) UnmarshalFromReader(reader jspb.Reader) *Int32Value {
	for reader.Next() {
		if m == nil {
			m = &Int32Value{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadInt32()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Int32Value from a slice of bytes.
func (m *Int32Value) Unmarshal(rawBytes []byte) (*Int32Value, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value uint32
}

// GetValue gets the Value of the UInt32Value.
func (m *UInt32Value) GetValue() (x uint32) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals UInt32Value to the provided writer.
func (m *UInt32Value) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value != 0 {
		writer.WriteUint32(1, m.Value)
	}

	return
}

// Marshal marshals UInt32Value to a slice of bytes.
func (m *UInt32Value) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a UInt32Value from the provided reader.
func (m *UInt32Value) UnmarshalFromReader(reader jspb.Reader) *UInt32Value {
	for reader.Next() {
		if m == nil {
			m = &UInt32Value{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadUint32()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a UInt32Value from a slice of bytes.
func (m *UInt32Value) Unmarshal(rawBytes []byte) (*UInt32Value, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value bool
}

// GetValue gets the Value of the BoolValue.
func (m *BoolValue) GetValue() (x bool) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals BoolValue to the provided writer.
func (m *BoolValue) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Value {
		writer.WriteBool(1, m.Value)
	}

	return
}

// Marshal marshals BoolValue to a slice of bytes.
func (m *BoolValue) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a BoolValue from the provided reader.
func (m *BoolValue) UnmarshalFromReader(reader jspb.Reader) *BoolValue {
	for reader.Next() {
		if m == nil {
			m = &BoolValue{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadBool()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a BoolValue from a slice of bytes.
func (m *BoolValue) Unmarshal(rawBytes []byte) (*BoolValue, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value string
}

// GetValue gets the Value of the StringValue.
func (m *StringValue) GetValue() (x string) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals StringValue to the provided writer.
func (m *StringValue) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Value) > 0 {
		writer.WriteString(1, m.Value)
	}

	return
}

// Marshal marshals StringValue to a slice of bytes.
func (m *StringValue) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a StringValue from the provided reader.
func (m *StringValue) UnmarshalFromReader(reader jspb.Reader) *StringValue {
	for reader.Next() {
		if m == nil {
			m = &StringValue{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a StringValue from a slice of bytes.
func (m *StringValue) Unmarshal(rawBytes []byte) (*StringValue, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value []byte
}

// GetValue gets the Value of the BytesValue.
func (m *BytesValue) GetValue() (x []byte) {
	if m == nil {
		return x
	}
	return m.Value
}

// MarshalToWriter marshals BytesValue to the provided writer.
func (m *BytesValue) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Value) > 0 {
		writer.WriteBytes(1, m.Value)
	}

	return
}

// Marshal marshals BytesValue to a slice of bytes.
func (m *BytesValue) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a BytesValue from the provided reader.
func (m *BytesValue) UnmarshalFromReader(reader jspb.Reader) *BytesValue {
	for reader.Next() {
		if m == nil {
			m = &BytesValue{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Value = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a BytesValue from a slice of bytes.
func (m *BytesValue) Unmarshal(rawBytes []byte) (*BytesValue, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}