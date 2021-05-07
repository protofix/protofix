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
	tagRawsPool = sync.Pool{New: func() interface{} { m := map[int][]byte{}; return &m }}
	unkTagsPool = sync.Pool{New: func() interface{} { return new([]int) }}
)

func (marsh Marshal) marshal(v interface{}) ([]byte, map[int]interface{}, []error, error) {
	stru := reflect.ValueOf(v).Elem()
	var unknown map[int]interface{}

	if marsh.Tag == "" {
		return nil, nil, nil, errors.New("missing tag of the unmarshal")
	}

	err := marsh.Format.Validate()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshaler format validation: %w", err)
	}

	tagRaws := *tagRawsPool.Get().(*map[int][]byte)
	for tag := range tagRaws {
		delete(tagRaws, tag)
	}
	defer tagRawsPool.Put(&tagRaws)

	var bodyLen9, chkSum10 reflect.Value
	var warns []error

	for i := 0; i < stru.NumField(); i++ {
		fldVal := stru.Field(i)
		fldTyp := stru.Type().Field(i)

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
		// Head 9.
		case f0.BodyLength9:
			bodyLen9 = val
			// Postpone the body length encoding.
			continue

		// Head+Body.
		default:
			codec = marsh.Format.Fields[num]

		// Trail 10.
		case f0.CheckSum10:
			chkSum10 = val
			// Postpone the checksum encoding.
			continue
		}

		if codec == nil || codec.Sizer() == nil {
			if val.Interface() == nil {
				continue
			}

			if unknown == nil {
				unknown = map[int]interface{}{}
			}
			unknown[num] = val.Interface()

			err = fmt.Errorf("missing codec of the field tag %d %q", num, f0.TagText[num])
			warns = append(warns, err)

			codec = marsh.Format.Unknown
		}

		raw, err := codec.Encode(val)
		if err != nil {
			err = fmt.Errorf("encode field tag %d %q: %w", num, f0.TagText[num], err)
			warns = append(warns, err)
		}

		if raw == nil {
			continue
		}

		tagRaws[num] = Concatenate(tag, raw)
	}

	enc9 := marsh.Format.BodyLength9

	if enc9.Crypt == nil || enc9.Size == nil {
		err := fmt.Errorf("missing codec of the mandatory field tag %d %q", f0.BodyLength9, f0.TagText[f0.BodyLength9])
		return nil, unknown, warns, err

	} else {
		enc9.Serial.Fields = tagRaws

		raw, err := enc9.Encode(bodyLen9)
		if err != nil {
			err = fmt.Errorf("encode field tag %d %q: %w", f0.BodyLength9, f0.TagText[f0.BodyLength9], err)
			warns = append(warns, err)
		}

		if raw != nil {
			tagRaws[f0.BodyLength9] = Concatenate(strconv.Itoa(f0.BodyLength9), raw)
		}
	}

	enc10 := marsh.Format.CheckSum10

	if enc10.Crypt == nil || enc10.Size == nil {
		err := fmt.Errorf("missing codec of the mandatory field tag %d %q", f0.CheckSum10, f0.TagText[f0.CheckSum10])
		return nil, unknown, warns, err

	} else {
		enc10.Serial.Fields = tagRaws

		raw, err := enc10.Encode(chkSum10)
		if err != nil {
			err = fmt.Errorf("encode field tag %d %q: %w", f0.CheckSum10, f0.TagText[f0.CheckSum10], err)
			warns = append(warns, err)
		}

		if raw != nil {
			tagRaws[f0.CheckSum10] = Concatenate(strconv.Itoa(f0.CheckSum10), raw)
		}
	}

	for tag, field := range marsh.Format.Fields {
		if _, ok := tagRaws[tag]; !ok && field.Required() {
			err := fmt.Errorf("missing value of the mandatory field tag %d %q", tag, f0.TagText[tag])
			warns = append(warns, err)
		}
	}

	var buf bytes.Buffer

	// Writing fields with the exception of the checksum field, it is the last field.
	for _, tag := range marsh.Format.Sort[:len(marsh.Format.Sort)-1] {
		raw, ok := tagRaws[tag]
		if !ok {
			continue
		}

		_, err := buf.Write(raw)
		if err != nil {
			err = fmt.Errorf("write field tag %d %q: %w", tag, f0.TagText[tag], err)
			return nil, unknown, warns, err
		}

		delete(tagRaws, tag)
	}

	unkTags := *unkTagsPool.Get().(*[]int)
	unkTags = unkTags[:0]
	defer unkTagsPool.Put(&unkTags)

	for tag := range tagRaws {
		if tag == f0.CheckSum10 {
			continue
		}
		unkTags = append(unkTags, tag)
	}

	sort.Ints(unkTags)

	// Writing remaining unknown fields. We do not write checksum field here.
	for _, tag := range unkTags {
		_, err := buf.Write(tagRaws[tag])
		if err != nil {
			err = fmt.Errorf("write unknown field tag %d %q: %w", tag, f0.TagText[tag], err)
			return nil, unknown, warns, err
		}
	}

	// Finally writing checksum field.
	_, err = buf.Write(tagRaws[f0.CheckSum10])
	if err != nil {
		err = fmt.Errorf("write field tag %d %q: %w", f0.CheckSum10, f0.TagText[f0.CheckSum10], err)
		return nil, unknown, warns, err
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
