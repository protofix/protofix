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
	"strconv"
	"strings"

	"github.com/danil/protoscan"
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/scanfix"
)

type Unmarshal struct {
	Tag    string
	Format f0.Format
}

const unmarshalTitle = "FIX unmarshal"

func (marsh Unmarshal) Unmarshal(p []byte, v interface{}) (map[int][]byte, []error, error) {
	unknown, warns, err := marsh.unmarshal(p, v)
	for i := 0; i < len(warns); i++ {
		warns[i] = fmt.Errorf("%s: %w", unmarshalTitle, warns[i])
	}
	if err != nil {
		return unknown, warns, fmt.Errorf("%s: %w, struct: %+v", unmarshalTitle, err, v)
	}
	return unknown, warns, nil
}

func (marsh Unmarshal) unmarshal(p []byte, v interface{}) (map[int][]byte, []error, error) {
	dst := reflect.ValueOf(v).Elem()
	tagVals := map[int]reflect.Value{}

	if marsh.Tag == "" {
		return nil, nil, errors.New("missing tag of the unmarshal")
	}

	err := marsh.Format.Validate()
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshaler format validation: %w", err)
	}

	var warns []error

	for i := 0; i < dst.NumField(); i++ {
		fldVal := dst.Field(i)
		fldTyp := dst.Type().Field(i)

		tag := strings.Split(fldTyp.Tag.Get(marsh.Tag), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if tag == "" {
			continue
		}

		n, err := strconv.ParseInt(tag, 10, 64)
		if err != nil {
			err = fmt.Errorf("parse struct tag %q: %s", tag, err)
			warns = append(warns, err)
			continue
		}

		num := int(n)
		tagVals[num] = fldVal
	}

	scan := &scanfix.Field{Format: marsh.Format}
	var buf []byte
	var unknown map[int][]byte

	scanner := protoscan.Protoscan{
		Reader: bytes.NewReader(p),
		Split:  scan.Split,
		Buffer: buf,
	}

	for scanner.Scan() {
		val, ok := tagVals[scan.Tag]
		if !ok {
			if unknown == nil {
				unknown = map[int][]byte{}
			}
			// we cannot unmarshal without destination, so remember it ...
			unknown[scan.Tag] = scanner.Token()

			err := fmt.Errorf("missing destination for the field tag %d %q", scan.Tag, f0.TagText[scan.Tag])
			warns = append(warns, err)

			// ... and go forth
			continue
		}

		var codec f0.Codec

		switch scan.Tag {
		// Head 9.
		case f0.BodyLength9:
			codec = scan.Format.BodyLength9

		// Head+Body.
		default:
			codec = scan.Format.Fields[scan.Tag]

		// Trail 10.
		case f0.CheckSum10:
			codec = scan.Format.CheckSum10
		}

		if codec == nil || codec.Sizer() == nil {
			if unknown == nil {
				unknown = map[int][]byte{}
			}
			token := make([]byte, len(scanner.Token()))
			copy(token, scanner.Token())
			unknown[scan.Tag] = token

			err = fmt.Errorf("missing codec of the field tag %d %q", scan.Tag, f0.TagText[scan.Tag])
			warns = append(warns, err)

			codec = marsh.Format.Unknown
		}

		recievedGap := fmt.Sprintf("%d=\x01", scan.Tag)
		expectedGap := string(bytes.Join(scanner.Gap(), []byte{}))
		if recievedGap != expectedGap {
			err := fmt.Errorf("unexpected scan gap, expected, %q, recieved: %q", expectedGap, recievedGap)
			return unknown, warns, err
		}

		err := codec.Decode(val, scanner.Token())
		if err != nil {
			err = fmt.Errorf("decode field tag %d %q: %w", scan.Tag, f0.TagText[scan.Tag], err)
			warns = append(warns, err)
		}
	}

	return unknown, warns, scanner.Err()
}
