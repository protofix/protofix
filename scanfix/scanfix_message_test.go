// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix_test

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/danil/equal4"
	"github.com/danil/protoscan"
	"github.com/protofix/protofix/scanfix"
)

var scanMessageTestCases = []struct {
	name              string
	line              int
	scanner           scanfix.Message
	input             string
	expectedTokens    string
	expectedGaps      []string
	expectedBodyLens  []string
	expectedMsgTypes  []string
	expectedChecksums []string
	expectedError     error
}{
	{
		name:              "one message",
		line:              line(),
		scanner:           scanfix.Message{BeginString: "FIX.4.4"},
		input:             "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|",
		expectedTokens:    "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|",
		expectedBodyLens:  []string{"70"},
		expectedMsgTypes:  []string{"A"},
		expectedChecksums: []string{"173"},
	},
	{
		name:    "three messages",
		line:    line(),
		scanner: scanfix.Message{BeginString: "FIX.4.4"},
		input: "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|" +
			"8=FIX.4.4|9=100|35=4|49=FG|56=tgFZctx20008c|34=1|43=Y|52=20210224-12:11:42.403|122=20210224-12:11:42.403|123=Y|36=6|10=157|" +
			"8=FIX.4.4|9=58|35=0|49=FG|56=tgFZctx20008c|34=6|52=20210224-12:12:12.162|10=129|",
		expectedTokens: "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|" +
			"8=FIX.4.4|9=100|35=4|49=FG|56=tgFZctx20008c|34=1|43=Y|52=20210224-12:11:42.403|122=20210224-12:11:42.403|123=Y|36=6|10=157|" +
			"8=FIX.4.4|9=58|35=0|49=FG|56=tgFZctx20008c|34=6|52=20210224-12:12:12.162|10=129|",
		expectedBodyLens:  []string{"70", "100", "58"},
		expectedMsgTypes:  []string{"A", "4", "0"},
		expectedChecksums: []string{"173", "157", "129"},
	},
	{
		name:    "immediate EOF",
		line:    line(),
		scanner: scanfix.Message{BeginString: "FIX.4.4"},
	},
	{
		name:              "one message with garbage before head and after tail",
		line:              line(),
		scanner:           scanfix.Message{BeginString: "FIX.4.4"},
		input:             "foo" + "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|" + "bar",
		expectedTokens:    "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|",
		expectedGaps:      []string{"foo", "bar"},
		expectedBodyLens:  []string{"70"},
		expectedMsgTypes:  []string{"A"},
		expectedChecksums: []string{"173"},
	},
	{
		name:    "two messages with garbage before, after and between",
		line:    line(),
		scanner: scanfix.Message{BeginString: "FIX.4.4"},
		input: "foo" +
			"8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|" +
			"bar" +
			"8=FIX.4.4|9=100|35=4|49=FG|56=tgFZctx20008c|34=1|43=Y|52=20210224-12:11:42.403|122=20210224-12:11:42.403|123=Y|36=6|10=157|" +
			"xyz",
		expectedTokens: "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|" +
			"8=FIX.4.4|9=100|35=4|49=FG|56=tgFZctx20008c|34=1|43=Y|52=20210224-12:11:42.403|122=20210224-12:11:42.403|123=Y|36=6|10=157|",
		expectedGaps:      []string{"foo", "bar", "xyz"},
		expectedBodyLens:  []string{"70", "100"},
		expectedMsgTypes:  []string{"A", "4"},
		expectedChecksums: []string{"173", "157"},
	},
	{
		name:         "wrong begin string",
		line:         line(),
		scanner:      scanfix.Message{BeginString: "FOO.BAR"},
		input:        "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|",
		expectedGaps: []string{"8=FIX.4.4\x019=70\x0135=A\x0149=FG\x0156=tgFZctx20008c\x0134=5\x0152=20210224-12:11:42.394\x0198=0\x01108=30\x0110=173\x01"},
	},
	{
		name:          "missing begin string",
		line:          line(),
		scanner:       scanfix.Message{BeginString: ""},
		input:         "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.394|98=0|108=30|10=173|",
		expectedError: errors.New("missing begin string"),
	},
}

func TestScanMessage(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range scanMessageTestCases {
		tc := tc
		t.Run(fmt.Sprintf("scan FIX 4.4 %s %d", tc.name, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			input := strings.ReplaceAll(tc.input, "|", "\x01")
			r := bytes.NewReader([]byte(input))
			split := &tc.scanner
			scan := protoscan.Protoscan{
				Reader: r,
				Split:  split.Split,
			}

			var (
				tokens    []byte
				gaps      []string
				bodyLens  []string
				msgTypes  []string
				checksums []string
			)

			for scan.Scan() {
				token := make([]byte, len(scan.Token()))
				copy(token, scan.Token())
				tokens = append(tokens, token...)

				for _, g := range scan.Gap() {
					gaps = append(gaps, string(g))
				}

				bodyLens = append(bodyLens, strconv.Itoa(split.BodyLength))

				msgTypes = append(msgTypes, string(split.MsgType))
				checksums = append(checksums, string(split.Checksum))
			}

			err := scan.Err()
			if !equal4.ErrorEqual(err, tc.expectedError) {
				t.Errorf("unexpected scan error, expected: %s, received: %s %s", tc.expectedError, err, linkToExample)
			}

			for _, g := range scan.Gap() {
				gaps = append(gaps, string(g))
			}

			expectedTokens := strings.ReplaceAll(tc.expectedTokens, "|", "\x01")
			if string(tokens) != expectedTokens {
				t.Errorf("unexpected tokens, expected: %q, received: %q %s", expectedTokens, tokens, linkToExample)
			}

			if strings.Join(gaps, "") != strings.Join(tc.expectedGaps, "") {
				t.Errorf("unexpected gaps, expected: %q, received: %q %s", tc.expectedGaps, gaps, linkToExample)
			}

			if strings.Join(bodyLens, "") != strings.Join(tc.expectedBodyLens, "") {
				t.Errorf("unexpected body length, expected: %#v, received: %#v %s", tc.expectedBodyLens, bodyLens, linkToExample)
			}

			if strings.Join(msgTypes, "") != strings.Join(tc.expectedMsgTypes, "") {
				t.Errorf("unexpected message type, expected: %#v, received: %#v %s", tc.expectedMsgTypes, msgTypes, linkToExample)
			}

			if strings.Join(checksums, "") != strings.Join(tc.expectedChecksums, "") {
				t.Errorf("unexpected checksum, expected: %#v, received: %#v %s", tc.expectedChecksums, checksums, linkToExample)
			}
		})
	}
}
