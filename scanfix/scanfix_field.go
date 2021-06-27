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
	Format   f0.Format
	Tag      int // Tag is a unique number of the last successfully tokenized FIX field.
	splitter int // splitter is an index of a "=" character of a FIX field, splitter is a tag/value separator.
}

// Split is a Protoscan splitter which splits field of the FIX message.
// Returns hint is a number of bytes hinted to read
// and returns advance is a needed number of bytes by which the carriage is to shift
// and an error if occurs.
// Each token is a value of the FIX message.
func (scan *Field) Split(data []byte, tokens *[]byte, indexes *[]int, gaps *[]byte, atEOF bool) (int, int, error) {
	// Read at least 2 bytes, for example: "8=".
	if hint := 2 - len(data); hint > 0 {
		return hint, 0, nil
	}

	if scan.splitter == 0 {
		// The splitter character (=) is not at the tail, so reads one more byte.
		if data[len(data)-1] != '=' {
			return 1, 0, nil
		}

		scan.splitter = len(data) - 1
		s := string(data[:scan.splitter])

		i, err := strconv.Atoi(s)
		if err != nil {
			err = fmt.Errorf("parsing FIX tag: %q, substring: %q, error: %w", data[:len(data)-1], data, err)
			return 0, 0, err
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

	l := len(data) - scan.splitter - 2
	if l > codec.Sizer().Max() {
		err := fmt.Errorf(
			"unexpected value length %d of the tag %d %q, maximum value length %d, field length %d, field: %q",
			l, scan.Tag, f0.TagText[scan.Tag], codec.Sizer().Max(), len(data), data,
		)
		return 0, 0, err
	}

	if hint := codec.Sizer().Min() + scan.splitter + 2 - len(data); hint > 0 {
		return hint, 0, nil
	}

	if data[len(data)-1] != 0x01 {
		return 1, 0, nil
	}

	token := append([]byte{}, data[scan.splitter+1:len(data)-1]...)

	gap := append([]byte{}, data[:scan.splitter+1]...)
	gap = append(gap, data[len(data)-1:]...)

	*tokens = append(*tokens, token...)
	*indexes = append(*indexes, len(token))
	*gaps = append(*gaps, gap...)

	scan.splitter = 0

	return 0, len(data), nil
}
