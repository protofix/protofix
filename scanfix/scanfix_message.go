// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// Message is a FIX messages splitter for the Protoscan.
type Message struct {
	BeginString string // BeginString contains the FIX version for the field tag 8.
	BodyLength  int    // BodyLength of the last successfully tokenized FIX message <https://en.wikipedia.org/wiki/Financial_Information_eXchange#Body_length>.
	MsgType     []byte // MsgType is a message type of the last successfully tokenized FIX message <https://en.wikipedia.org/wiki/Financial_Information_eXchange#FIX_tagvalue_message_format>.
	CheckSum    []byte // CheckSum of the last successfully tokenized FIX message (checksum validation is not perform) <https://en.wikipedia.org/wiki/Financial_Information_eXchange#Trailer:_Checksum>.
	Gaps        []byte // Last skipped bytes copied by Split. Gaps intend to holds buffered and unprocessed bytes for example if EOF occurs.
	length      int    // length is a body length of a FIX message being tokenizing.
}

// Split is a Protoscan splitter which splits a FIX messages.
// Returns hint is a number of bytes hinted to read
// and returns advance is a needed number of bytes by which the carriage is to shift
// and returns token if available
// and an error if occurs.
// Each token is a FIX message.
func (scan *Message) Split(data []byte, atEOF bool) (int, int, []byte, error) {
	scan.Gaps = scan.Gaps[:0]

	if scan.BeginString == "" {
		return 0, 0, nil, errors.New("missing begin string")
	}

	if atEOF && len(data) > 0 {
		scan.MsgType = nil
		scan.CheckSum = nil
		scan.BodyLength = 0
		scan.length = 0
		scan.Gaps = append(scan.Gaps, data...)
		return 0, len(data), nil, nil
	}

	// Read at least 14 bytes: "8=FIX.?.?|9=?|".
	if hint := 14 - len(data); hint > 0 {
		return hint, 0, nil, nil
	}

	tag8 := bytes.Index(data, []byte("8="+scan.BeginString+"\x019="))
	if tag8 == -1 {
		return 1, 0, nil, nil
	}

	tag9 := tag8 + 10

	if hint := tag9 + 3 - len(data); hint > 0 {
		return hint, 0, nil, nil
	}

	// soh9 is an index of a 0x01 or \x01 SOH character (start of heading)
	// of a FIX tag 9 <https://en.wikipedia.org/wiki/C0_and_C1_control_codes#SOH>.
	// Actually soh9 is a end of tag.
	soh9 := bytes.IndexByte(data[tag9+3:], 0x01)
	if soh9 == -1 {
		return 1, 0, nil, nil
	}

	soh9 += tag9 + 3

	if scan.length == 0 {
		s := string(data[tag9+2 : soh9])
		l, err := strconv.Atoi(s)
		if err != nil {
			err = fmt.Errorf("parsing FIX body length: %q, substring: %q, error: %w", s, data[tag8:soh9], err)
			return 0, 0, nil, err
		}
		scan.length = l
	}

	// hint equal to the body length plus tag 10 length (FIX Message Checksum)
	// <https://en.wikipedia.org/wiki/Financial_Information_eXchange#Body_length>.
	if hint := scan.length - len(data[soh9+1:]) + 7; hint > 0 {
		return hint, 0, nil, nil
	}

	tag10 := tag8 + len(data[tag8:soh9+1+scan.length+7]) - 7
	soh10 := tag8 + len(data[tag8:soh9+1+scan.length+7]) - 1

	if !bytes.Equal(data[tag10-1:tag10+3], []byte("\x01"+"10=")) {
		return 0, 0, nil, fmt.Errorf("missing a tag of the field 10, message: %q", data[tag8:soh9+1+scan.length+7])
	}

	if data[soh10] != 0x01 {
		return 0, 0, nil, fmt.Errorf("missing a SOH of the field 10, message: %q", data[tag8:soh9+1+scan.length+7])
	}

	i := bytes.Index(data[soh9:soh10], []byte("\x01"+"35="))
	if i == -1 {
		return 0, 0, nil, fmt.Errorf("missing a tag of the field 35, message: %q", data[tag8:soh9+1+scan.length+7])
	}

	tag35 := soh9 + i + 1

	soh35 := 0
	for i := tag35 + 3; i < tag10; i++ {
		if data[i] == 0x01 {
			soh35 = i
			break
		}
	}

	if soh35 == 0 {
		return 0, 0, nil, fmt.Errorf("missing a SOH of the field 35, message: %q", data[tag8:soh9+1+scan.length+7])
	}

	if len(data[:tag8]) > 0 {
		scan.Gaps = append(scan.Gaps, data[:tag8]...)
	}

	scan.BodyLength = scan.length
	scan.MsgType = data[tag35+3 : soh35]
	scan.CheckSum = data[tag10+3 : soh10]

	length := scan.length
	scan.length = 0

	return 0, tag8 + len(data[tag8:soh9+1+length+7]), data[tag8 : soh9+1+length+7], nil
}
