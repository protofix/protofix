// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix

import (
	"fmt"
	"strconv"

	f0 "github.com/protofix/protofix/codecfix"
)

// Field is a Protoscan splitter which splits field of the FIX message.
type Field struct {
	Format f0.Format
	Tag    int // Tag is a unique number of the last successfully tokenized FIX field.
	equal  int // equal is an index of a "=" character of a FIX field, equal is a tag/value separator.
}

// Split is a Protoscan splitter which splits field of the FIX message.
// Returns hint is a number of bytes hinted to read
// and returns advance is a needed number of bytes by which the carriage is to shift
// and returns a token and an error if occurs.
// Each token is a value of the FIX message.
func (scan *Field) Split(data []byte, atEOF bool) (int, [][]byte, int, []byte, error) {
	// Read at least 2 bytes, for example: "8=".
	if hint := 2 - len(data); hint > 0 {
		return hint, nil, 0, nil, nil
	}

	if scan.equal == 0 {
		// The equal character (=) is not at the tail, so reads one more byte.
		if data[len(data)-1] != '=' {
			return 1, nil, 0, nil, nil
		}

		scan.equal = len(data) - 1
		s := string(data[:scan.equal])

		i, err := strconv.Atoi(s)
		if err != nil {
			err = fmt.Errorf("parsing FIX tag: %q, substring: %q, error: %w", data[:len(data)-1], data, err)
			return 0, nil, 0, nil, err
		}

		scan.Tag = i
	}

	var codec f0.Codec

	switch scan.Tag {
	// Head 9.
	case f0.BodyLength9:
		if scan.Format.BodyLength9.Size != nil {
			codec = scan.Format.BodyLength9
		}

	// Head+Body.
	default:
		codec = scan.Format.Fields[scan.Tag]

	// Trail 10.
	case f0.CheckSum10:
		if scan.Format.CheckSum10.Size != nil {
			codec = scan.Format.CheckSum10
		}
	}

	if codec == nil {
		codec = scan.Format.Unknown0
	}

	l := len(data) - scan.equal - 2
	if l > codec.Sizer().Max() {
		err := fmt.Errorf(
			"unexpected value length %d of the tag %d %q, maximum value length %d, field length %d, field: %q",
			l, scan.Tag, f0.TagText[scan.Tag], codec.Sizer().Max(), len(data), data,
		)
		return 0, nil, 0, nil, err
	}

	if hint := codec.Sizer().Min() + scan.equal + 2 - len(data); hint > 0 {
		return hint, nil, 0, nil, nil
	}

	if data[len(data)-1] != 0x01 {
		return 1, nil, 0, nil, nil
	}

	gap := [][]byte{data[:scan.equal+1], data[len(data)-1:]}
	token := data[scan.equal+1 : len(data)-1]
	scan.equal = 0

	return 0, gap, len(data), token, nil
}
