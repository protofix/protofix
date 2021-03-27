// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package marshfix is a FIX protocol marshaler/unmarshaler.
package marshfix

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	f0 "github.com/danil/protofix/codecfix"
)

type Marshal struct {
	Tag    string
	Format f0.Format
}

const marshalTitle = "FIX marshal"

func (marsh Marshal) Marshal(v interface{}) ([]byte, map[int]interface{}, []error, error) {
	enc, unknown, warns, err := marsh.marshal(v)
	for i := 0; i < len(warns); i++ {
		warns[i] = fmt.Errorf("%s: %w", marshalTitle, warns[i])
	}
	if err != nil {
		return enc, unknown, warns, fmt.Errorf("%s: %w, struct: %+v", marshalTitle, err, v)
	}
	return enc, unknown, warns, nil
}

var (
	headNumsPool = sync.Pool{New: func() interface{} { return new([]int) }}
	bodyNumsPool = sync.Pool{New: func() interface{} { return new([]int) }}
	tagRawsPool  = sync.Pool{New: func() interface{} { m := map[int][]byte{}; return &m }}
)

func (marsh Marshal) marshal(v interface{}) ([]byte, map[int]interface{}, []error, error) {
	str := reflect.ValueOf(v).Elem()
	var unknown map[int]interface{}

	if marsh.Tag == "" {
		return nil, nil, nil, errors.New("missing tag of the unmarshal")
	}

	err := marsh.Format.Validate()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshaler format validation: %w", err)
	}

	headNums := *headNumsPool.Get().(*[]int)
	headNums = headNums[:0]
	defer headNumsPool.Put(&headNums)

	bodyNums := *bodyNumsPool.Get().(*[]int)
	bodyNums = bodyNums[:0]
	defer bodyNumsPool.Put(&bodyNums)

	tagRaws := *tagRawsPool.Get().(*map[int][]byte)
	for k := range tagRaws {
		delete(tagRaws, k)
	}
	defer tagRawsPool.Put(&tagRaws)

	var bodyLen9, chkSum10 reflect.Value
	var warns []error

	for i := 0; i < str.NumField(); i++ {
		fldVal := str.Field(i)
		fldTyp := str.Type().Field(i)

		tag := strings.Split(fldTyp.Tag.Get(marsh.Tag), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if tag == "" {
			continue
		}

		n, err := strconv.ParseInt(tag, 10, 64)
		if err != nil {
			err = fmt.Errorf("parse struct tag %q: %w", tag, err)
			warns = append(warns, err)
			continue
		}

		num := int(n)
		f := fldVal.Interface()
		val := reflect.ValueOf(f)

		var codec f0.Codec

		switch num {
		// Head.
		case f0.BeginString8, f0.MsgType35, f0.SenderCompID49, f0.TargetCompID56, f0.ApplVerID1128:
			codec = marsh.Format.Head[num]
			headNums = append(headNums, num)

		// Head 9.
		case f0.BodyLength9:
			bodyLen9 = val
			// Postpone the body length encoding.
			continue

		// Body.
		default:
			codec = marsh.Format.Body[num]
			bodyNums = append(bodyNums, num)

		// Trail 10.
		case f0.Checksum10:
			chkSum10 = val
			// Postpone the checksum encoding.
			continue
		}

		if codec == nil || codec.Varyer() == nil {
			if val.Interface() == nil {
				continue
			}

			if unknown == nil {
				unknown = map[int]interface{}{}
			}
			unknown[num] = val.Interface()

			err = fmt.Errorf("missing codec of the field tag %d: %q", num, f0.FldText[num])
			warns = append(warns, err)

			codec = marsh.Format.Unknown
		}

		raw, err := codec.Encode(val)
		if err != nil {
			err = fmt.Errorf("encode field tag %d: %q: %w", num, f0.FldText[num], err)
			warns = append(warns, err)
		}

		if raw == nil {
			continue
		}

		tagRaws[num] = Concatenate(tag, raw)
	}

	enc9 := marsh.Format.BodyLength9

	if enc9.Crypt == nil || enc9.Var == nil {
		err := fmt.Errorf(
			"missing codec of the mandatory field tag %d: %q",
			f0.BodyLength9, f0.FldText[f0.BodyLength9],
		)
		return nil, unknown, warns, err

	} else {
		enc9.Serial.Fields = tagRaws

		raw, err := enc9.Encode(bodyLen9)
		if err != nil {
			err = fmt.Errorf("encode field tag %d: %q: %w", f0.BodyLength9, f0.FldText[f0.BodyLength9], err)
			warns = append(warns, err)
		}

		if raw != nil {
			headNums = append(headNums, f0.BodyLength9)
			tagRaws[f0.BodyLength9] = Concatenate(strconv.Itoa(f0.BodyLength9), raw)
		}
	}

	sort.Ints(headNums)
	sort.Ints(bodyNums)

	enc10 := marsh.Format.Checksum10

	if enc10.Crypt == nil || enc10.Var == nil {
		err := fmt.Errorf(
			"missing codec of the mandatory field tag %d: %q",
			f0.Checksum10, f0.FldText[f0.Checksum10],
		)
		return nil, unknown, warns, err

	} else {
		enc10.Serial.Fields = tagRaws

		raw, err := enc10.Encode(chkSum10)
		if err != nil {
			err = fmt.Errorf("encode field tag %d: %q: %w", f0.Checksum10, f0.FldText[f0.Checksum10], err)
			warns = append(warns, err)
		}

		if raw != nil {
			// The checksum of a FIX message is always the last field in the message.
			// So appending after sorting.
			// <https://en.wikipedia.org/wiki/Financial_Information_eXchange#Trailer:_Checksum>.
			bodyNums = append(bodyNums, f0.Checksum10)
			tagRaws[f0.Checksum10] = Concatenate(strconv.Itoa(f0.Checksum10), raw)
		}
	}

	for k, v := range marsh.Format.Head {
		if _, ok := tagRaws[k]; !ok && v.Required() {
			err := fmt.Errorf("missing value of the mandatory header field tag %d: %q", k, f0.FldText[k])
			warns = append(warns, err)
		}
	}

	for k, v := range marsh.Format.Body {
		if _, ok := tagRaws[k]; !ok && v.Required() {
			err := fmt.Errorf("missing value of the mandatory body field tag %d: %q", k, f0.FldText[k])
			warns = append(warns, err)
		}
	}

	var buf bytes.Buffer

	for _, n := range append(headNums, bodyNums...) {
		n, err := buf.Write(tagRaws[n])
		if err != nil {
			err = fmt.Errorf("write field tag %d: %q: %w", n, f0.FldText[n], err)
			return nil, unknown, warns, err
		}
	}

	return buf.Bytes(), unknown, warns, nil
}

func Concatenate(tag string, val []byte) []byte {
	const (
		// soh is the 0x01 or \x01 Start of Heading (SOH) ASCII character
		// <https://en.wikipedia.org/wiki/C0_and_C1_control_codes#SOH>.
		// Actually soh9 is a end of tag.
		soh = 0x01
		// equal is the equal ASCII character.
		equal = 0x3D
	)
	return append([]byte(tag), append([]byte{equal}, append(val, soh)...)...)
}
