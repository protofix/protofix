// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package codecfix is a FIX protocol codec.
package codecfix

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Codec interface {
	Encode(rvalue reflect.Value) (encoded []byte, err error)
	Decode(rvalue reflect.Value, encoded []byte) (err error)
	Required() bool
	Sizer() Sizer
}

// Fld is a field codec.
type Fld struct {
	Require bool // Required or optional.
	Crypt   Cryptor
	Serial  Serializer
	Size    Sizer
}

func (f Fld) Encode(rval reflect.Value) ([]byte, error)   { return encode(f, rval) }
func (f Fld) Decode(rval reflect.Value, enc []byte) error { return decode(f, rval, enc) }
func (f Fld) Sizer() Sizer                                { return f.Size }
func (f Fld) Required() bool                              { return f.Require }

func (f Fld) Cryptor() Cryptor       { return f.Crypt }
func (f Fld) Serializer() Serializer { return f.Serial }

// BodyLengthFld is a codec for the BodyLength field. Body length is always mandatory.
type BodyLengthFld struct {
	Crypt  Cryptor
	Serial BodyLengthSrlz
	Size   Sizer
}

func (f BodyLengthFld) Encode(rval reflect.Value) ([]byte, error)   { return encode(f, rval) }
func (f BodyLengthFld) Decode(rval reflect.Value, enc []byte) error { return decode(f, rval, enc) }
func (f BodyLengthFld) Required() bool                              { return true }
func (f BodyLengthFld) Sizer() Sizer                                { return f.Size }

func (f BodyLengthFld) Cryptor() Cryptor       { return f.Crypt }
func (f BodyLengthFld) Serializer() Serializer { return f.Serial }

// ChecksumStringFld is a codec for the Checksum field. Checksum is always mandatory.
type ChecksumStringFld struct {
	Crypt  Cryptor
	Serial ChecksumStringSrlz
	Size   Sizer
}

func (f ChecksumStringFld) Encode(rval reflect.Value) ([]byte, error)   { return encode(f, rval) }
func (f ChecksumStringFld) Decode(rval reflect.Value, enc []byte) error { return decode(f, rval, enc) }
func (f ChecksumStringFld) Required() bool                              { return true }
func (f ChecksumStringFld) Sizer() Sizer                                { return f.Size }

func (f ChecksumStringFld) Cryptor() Cryptor       { return f.Crypt }
func (f ChecksumStringFld) Serializer() Serializer { return f.Serial }

// UnknownFld is a codec for the unknown field.
type UnknownFld struct {
	Crypt  Cryptor
	Serial Serializer
	Size   Sizer
}

func (f UnknownFld) Encode(rval reflect.Value) ([]byte, error)   { return encode(f, rval) }
func (f UnknownFld) Decode(rval reflect.Value, enc []byte) error { return decode(f, rval, enc) }
func (f UnknownFld) Required() bool                              { return false }
func (f UnknownFld) Sizer() Sizer                                { return f.Size }

func (f UnknownFld) Cryptor() Cryptor       { return f.Crypt }
func (f UnknownFld) Serializer() Serializer { return f.Serial }

type fielder interface {
	Cryptor() Cryptor
	Serializer() Serializer
	Sizer() Sizer
}

func encode(codec fielder, rval reflect.Value) ([]byte, error) {
	dec, err := codec.Serializer().Serialize(rval)
	if err != nil && len(dec) == 0 {
		return nil, err
	}

	enc, err2 := codec.Cryptor().Encrypt(dec)
	if err == nil {
		err = err2
	}

	return enc, err
}

func decode(codec fielder, rval reflect.Value, enc []byte) error {
	dec, err := codec.Cryptor().Decrypt(enc)
	if err != nil && len(dec) == 0 {
		return err
	}

	err2 := codec.Serializer().Deserialize(rval, dec)
	if err == nil {
		err = err2
	}

	return err
}

type Sizer interface {
	// Min is a minimum length of a value of a FIX field.
	Min() int
	// Max is a maximum length of a value of a FIX field.
	Max() int
	// Hint is a number of bytes of a field to read.
	Hint() int
}

// Con implements Sizer for the FIX fields with constant length.
type Con struct {
	Size int
}

func (c Con) Min() int  { return c.Size }
func (c Con) Max() int  { return c.Size }
func (c Con) Hint() int { return c.Size }

// Con implements Sizer for the FIX fields with variable length.
type Var struct {
	MinSize int
	MaxSize int
}

func (c Var) Min() int  { return c.MinSize }
func (c Var) Max() int  { return c.MaxSize }
func (c Var) Hint() int { return c.MinSize }

const (
	MaxBytes  = 65536
	MaxInt    = 18
	MaxString = MaxBytes
)

type Serializer interface {
	// Serialize serializes the FIX fields.
	Serialize(rvalue reflect.Value) (serialized []byte, err error)
	// Deserialize deserializes the FIX fields.
	Deserialize(rvalue reflect.Value, serialized []byte) error
}

//
// Byte slice serializers/deserializers.
//

// Bytes returns byte slice serializer with constraint.
func Bytes(cstr ...[]byte) BytesSrlz { return BytesSrlz{Constraint: cstr} }

// BytesSrlz intends to serialize/deserialize byte slice value only.
type BytesSrlz struct {
	Constraint [][]byte
}

// Serialize serializes byte slice value.
func (ser BytesSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	v := rval.Interface()

	val, ok := v.([]byte)
	if !ok {
		err := errors.New("unexpected struct field type, expected byte slice")
		return nil, err
	}

	if val == nil {
		return nil, nil
	}

	if len(ser.Constraint) != 0 {
		for _, c := range ser.Constraint {
			if bytes.Equal(c, val) {
				goto srlzValidBytes
			}
		}
		return nil, fmt.Errorf("unexpected value %q, allow: %v", val, ser.Constraint)
	srlzValidBytes:
	}

	return val, nil
}

// Deserialize deserializes byte slice value.
func (ser BytesSrlz) Deserialize(rval reflect.Value, p []byte) error {
	if len(ser.Constraint) != 0 {
		for _, c := range ser.Constraint {
			if bytes.Equal(c, p) {
				goto dsrlzValidBytes
			}
		}
		return fmt.Errorf("unexpected value %q, allow: %v", p, ser.Constraint)
	dsrlzValidBytes:
	}

	val := reflect.ValueOf(p)
	rval.Set(val)
	return nil
}

//
// Bool serializers/deserializers.
//

// Bool returns bool serializer with constraint.
func Bool(cstr ...bool) BoolSrlz { return BoolSrlz{Constraint: cstr} }

// BoolSrlz intends to serialize/deserialize bool value only.
type BoolSrlz struct {
	Constraint []bool
}

// Serialize serializes bool value.
func (ser BoolSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.Bool()
	if !val {
		return nil, nil
	}

	return SerializeBool(val, ser.Constraint)
}

// Deserialize deserializes bool value.
func (ser BoolSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeBool(rval, p, ser.Constraint)
}

// BoolDefault returns bool serializer with constraint with default value false.
func BoolDefault(cstr ...bool) BoolDefaultSrlz { return BoolDefaultSrlz{Constraint: cstr} }

// BoolDefaultSrlz intends to serialize/deserialize bool value with default value false.
type BoolDefaultSrlz struct {
	Constraint []bool
}

// Serialize serializes bool value with default value false.
func (ser BoolDefaultSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.Bool()
	return SerializeBool(val, ser.Constraint)
}

// Deserialize deserializes bool value with default value false.
func (ser BoolDefaultSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeBool(rval, p, ser.Constraint)
}

func SerializeBool(val bool, cstr []bool) ([]byte, error) {
	if len(cstr) == 1 && !val {
		val = cstr[0]

	} else if len(cstr) != 0 {
		for _, c := range cstr {
			if val == c {
				goto srlzValidBool
			}
		}
		return nil, fmt.Errorf("unexpected value %t, allow: %v", val, cstr)
	srlzValidBool:
	}

	if val {
		return []byte{BoolRawCstr[1]}, nil
	}

	return []byte{BoolRawCstr[0]}, nil
}

func DeserializeBool(rval reflect.Value, p []byte, cstr []bool) error {
	var val bool
	if len(p) == 1 && p[0] == BoolRawCstr[1] {
		val = true
	} else if len(p) != 1 || p[0] != BoolRawCstr[0] {
		return fmt.Errorf("unexpected raw value %q, allow: %v", p, BoolRawCstr)
	}

	if len(cstr) != 0 {
		for _, c := range cstr {
			if val == c {
				goto dsrlzValidBool
			}
		}
		return fmt.Errorf("unexpected value %t, allow: %v", val, cstr)
	dsrlzValidBool:
	}

	rval.SetBool(val)
	return nil
}

var BoolRawCstr = [2]byte{'N', 'Y'}

//
// Int serializers/deserializers.
//

// Int returns int serializer with constraint.
func Int(cstr ...int) IntSrlz { return IntSrlz{Constraint: cstr} }

// IntSrlz intends to serialize/deserialize int value only.
type IntSrlz struct {
	Constraint []int
}

// Serialize serializes int value.
func (ser IntSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := int(rval.Int())
	if val == 0 {
		return nil, nil
	}
	return SerializeInt(val, ser.Constraint)
}

// Deserialize deserializes int value.
func (ser IntSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeInt(rval, p, ser.Constraint)
}

// IntDefault returns int serializer with constraint with default zero.
func IntDefault(cstr ...int) IntDefaultSrlz { return IntDefaultSrlz{Constraint: cstr} }

// IntDefaultSrlz intends to serialize/deserialize int value with default zero.
type IntDefaultSrlz struct {
	Constraint []int
}

// Serialize serializes int value with default zero.
func (ser IntDefaultSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	return SerializeInt(int(rval.Int()), ser.Constraint)
}

// Deserialize deserializes int value with default zero.
func (ser IntDefaultSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeInt(rval, p, ser.Constraint)
}

// Length returns int serializer with constraint.
func Length(cstr ...int) LengthSrlz { return LengthSrlz{Constraint: cstr} }

// LengthSrlz intends to serialize/deserialize int value.
type LengthSrlz struct {
	Constraint []int
}

// Serialize serializes int value.
func (ser LengthSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := int(rval.Int())
	if val == 0 {
		return nil, nil
	}
	if val < 0 {
		err := fmt.Errorf("unexpected length value %d, must be positive", val)
		return nil, err
	}
	return SerializeInt(val, ser.Constraint)
}

// Deserialize deserializes int value.
func (ser LengthSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeInt(rval, p, ser.Constraint)
}

// SeqNum returns int serializer with constraint.
func SeqNum(cstr ...int) SeqNumSrlz { return SeqNumSrlz{Constraint: cstr} }

// SeqNumSrlz intends to serialize/deserialize int value.
type SeqNumSrlz struct {
	Constraint []int
}

// Serialize serializes int value.
func (ser SeqNumSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := int(rval.Int())
	if val < 0 {
		err := fmt.Errorf("unexpected length value %d, must be positive", val)
		return nil, err
	}
	return SerializeInt(val, ser.Constraint)
}

// Deserialize deserializes int value.
func (ser SeqNumSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeInt(rval, p, ser.Constraint)
}

// BodyLength returns int serializer with constraint.
func BodyLength(cstr ...int) BodyLengthSrlz { return BodyLengthSrlz{Constraint: cstr} }

// BodyLengthSrlz intends to serialize/deserialize int value only.
type BodyLengthSrlz struct {
	Constraint []int
	Fields     map[int][]byte
}

// Serialize serializes int value.
func (ser BodyLengthSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := int(rval.Int())
	if val != 0 {
		return SerializeInt(val, ser.Constraint)
	}

	for k, v := range ser.Fields {
		if k != BeginString8 {
			val += len(v)
		}
	}

	expected := 5
	if val < expected {
		err := fmt.Errorf("unexpected body length value %d, should be at least: %d", val, expected)
		return nil, err
	}

	return SerializeInt(val, ser.Constraint)
}

// Deserialize deserializes int value.
func (ser BodyLengthSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeInt(rval, p, ser.Constraint)
}

func SerializeInt(val int, cstr []int) ([]byte, error) {
	if len(cstr) == 1 && val == 0 {
		val = cstr[0]

	} else if len(cstr) != 0 {
		for _, c := range cstr {
			if c == val {
				goto srlzValidInt
			}
		}
		return nil, fmt.Errorf("unexpected value %q, allow: %v", val, cstr)
	srlzValidInt:
	}

	return []byte(strconv.Itoa(val)), nil
}

func DeserializeInt(rval reflect.Value, p []byte, cstr []int) error {
	val, err := strconv.Atoi(string(p))
	if err != nil {
		return err
	}

	if len(cstr) != 0 {
		for _, c := range cstr {
			if c == val {
				goto dsrlzValidInt
			}
		}
		return fmt.Errorf("unexpected value %q, allowable values: %v", val, cstr)
	dsrlzValidInt:
	}

	rval.SetInt(int64(val))
	return nil
}

//
// Time duration serializers/deserializers.
//

// SecondsDuration returns string serializer with constraint.
func SecondsDuration(cstr ...time.Duration) SecondsDurationSrlz {
	return SecondsDurationSrlz{Constraint: cstr}
}

// SecondsDurationSrlz intends to serialize/deserialize duration value only.
type SecondsDurationSrlz struct {
	Constraint []time.Duration
}

// Serialize serializes duration value.
func (ser SecondsDurationSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	i := rval.Int()
	if i == 0 {
		return nil, nil
	}

	val := time.Duration(i)
	var err error

	if len(ser.Constraint) != 0 {
		for _, c := range ser.Constraint {
			if c == val {
				goto srlzValidSecondsDuration
			}
		}
		err = fmt.Errorf("unexpected value %q, allow: %v", val, ser.Constraint)
	srlzValidSecondsDuration:
	}

	s := strconv.FormatInt(int64(val.Seconds()), 10)

	return []byte(s), err
}

// Deserialize deserializes duration value.
func (ser SecondsDurationSrlz) Deserialize(rval reflect.Value, p []byte) error {
	i, err := strconv.ParseInt(string(p), 10, 64)
	if err != nil {
		return fmt.Errorf("parse duration seconds value %q: %v", p, err)
	}

	val := time.Duration(i) * time.Second

	if len(ser.Constraint) != 0 {
		for _, c := range ser.Constraint {
			if c == val {
				goto dsrlzValidSecondsDuration
			}
		}
		err = fmt.Errorf("unexpected value %q, allow: %v", val, ser.Constraint)
	dsrlzValidSecondsDuration:
	}

	rval.SetInt(val.Nanoseconds())
	return err
}

//
// String serializers/deserializers.
//

// String returns string serializer with constraint.
func String(cstr ...string) StringSrlz { return StringSrlz{Constraint: cstr} }

// StringSrlz intends to serialize/deserialize string value only.
type StringSrlz struct {
	Constraint []string
}

// Serialize serializes string value.
func (ser StringSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.String()
	if val == "" {
		return nil, nil
	}
	return SerializeString(rval.String(), ser.Constraint)
}

// Deserialize deserializes string value.
func (ser StringSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeString(rval, p, ser.Constraint)
}

// StringDefault returns string serializer with constraint.
func StringDefault(cstr ...string) StringDefaultSrlz { return StringDefaultSrlz{Constraint: cstr} }

// StringDefaultSrlz intends to serialize/deserialize string value only.
type StringDefaultSrlz struct {
	Constraint []string
}

// Serialize serializes string value.
func (ser StringDefaultSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.String()
	return SerializeString(val, ser.Constraint)
}

// Deserialize deserializes string value.
func (ser StringDefaultSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeString(rval, p, ser.Constraint)
}

// ChecksumString returns string serializer with constraint.
func ChecksumString(cstr ...string) ChecksumStringSrlz {
	return ChecksumStringSrlz{Constraint: cstr}
}

// ChecksumStringSrlz intends to serialize/deserialize string value only.
type ChecksumStringSrlz struct {
	Constraint []string
	Fields     map[int][]byte
}

// Serialize serializes string value.
func (ser ChecksumStringSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.String()
	if val != "" {
		return SerializeString(rval.String(), ser.Constraint)
	}

	sum := 0
	for _, v := range ser.Fields {
		for _, b := range v {
			sum += int(b)
		}
	}

	val = fmt.Sprintf("%03d", sum%256)

	expected := 1
	if sum < expected {
		err := fmt.Errorf(`unexpected checksum value %q, should be at least: "%03d"`, val, expected)
		return nil, err
	}

	return SerializeString(val, ser.Constraint)
}

// Deserialize deserializes string value.
func (ser ChecksumStringSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return DeserializeString(rval, p, ser.Constraint)
}

func SerializeString(val string, cstr []string) ([]byte, error) {
	if len(cstr) == 1 && val == "" {
		val = cstr[0]

	} else if len(cstr) != 0 {
		for _, c := range cstr {
			if val == c {
				goto srlzValidString
			}
		}
		return nil, fmt.Errorf("unexpected value %q, allow: %v", val, cstr)
	srlzValidString:
	}

	return []byte(val), nil
}

func DeserializeString(rval reflect.Value, p []byte, cstr []string) error {
	val := string(p)
	if len(cstr) != 0 {
		for _, c := range cstr {
			if c == val {
				goto dsrlzValidString
			}
		}
		return fmt.Errorf("unexpected value %q, allowable values: %v", val, cstr)
	dsrlzValidString:
	}

	rval.SetString(val)
	return nil
}

//
// Unknown serializers/deserializers.
//

// Unknown returns serializer for unknown.
func Unknown(cstr ...[]byte) UnknownSrlz { return UnknownSrlz{} }

// UnknownSrlz intends to serialize/deserialize byte slice value only.
type UnknownSrlz struct{}

// Serialize serializes unknown value.
func (ser UnknownSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	val := rval.Interface()
	return []byte(fmt.Sprint(val)), nil
}

// Deserialize deserializes unknown value.
func (ser UnknownSrlz) Deserialize(rval reflect.Value, p []byte) error {
	return errors.New("cannot deserialize unknown value")
}

//
// UTC timestamp serializers/deserializers.
//

// UTCTimestampNanosecondTime returns string serializer.
func UTCTimestampNanosecondTime(cstr ...string) UTCTimestampNanosecondTimeSrlz {
	return UTCTimestampNanosecondTimeSrlz{}
}

// UTCTimestampNanosecondTimeSrlz intends to serialize/deserialize time value only.
type UTCTimestampNanosecondTimeSrlz struct{}

// Serialize serializes time value.
func (ser UTCTimestampNanosecondTimeSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	v := rval.Interface()

	val, ok := v.(time.Time)
	if !ok {
		err := errors.New("unexpected struct field type, expected time")
		return nil, err
	}

	if val.IsZero() {
		return nil, nil
	}

	s := fmt.Sprintf("%s.%09d", val.Format("20060102-15:04:05"), val.Nanosecond())

	return []byte(s), nil
}

// Deserialize deserializes time value.
func (ser UTCTimestampNanosecondTimeSrlz) Deserialize(rval reflect.Value, p []byte) error {
	val, err := time.Parse("20060102-15:04:05.999999999", string(p))
	if err != nil {
		return fmt.Errorf("parse time value %q: %v", p, err)
	}

	rval.Set(reflect.ValueOf(val))
	return nil
}

// UTCTimestampMicrosecondTime returns string serializer.
func UTCTimestampMicrosecondTime(cstr ...string) UTCTimestampMicrosecondTimeSrlz {
	return UTCTimestampMicrosecondTimeSrlz{}
}

// UTCTimestampMicrosecondTimeSrlz intends to serialize/deserialize time value only.
type UTCTimestampMicrosecondTimeSrlz struct{}

// Serialize serializes time value.
func (ser UTCTimestampMicrosecondTimeSrlz) Serialize(rval reflect.Value) ([]byte, error) {
	v := rval.Interface()

	val, ok := v.(time.Time)
	if !ok {
		err := errors.New("unexpected struct field type, expected time")
		return nil, err
	}

	if val.IsZero() {
		return nil, nil
	}

	s := fmt.Sprintf("%s.%06d", val.Format("20060102-15:04:05"), val.Nanosecond())

	return []byte(s), nil
}

// Deserialize deserializes time value.
func (ser UTCTimestampMicrosecondTimeSrlz) Deserialize(rval reflect.Value, p []byte) error {
	val, err := time.Parse("20060102-15:04:05.999999", string(p))
	if err != nil {
		return fmt.Errorf("parse time value %q: %v", p, err)
	}

	rval.Set(reflect.ValueOf(val))
	return nil
}
