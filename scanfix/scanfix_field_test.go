// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	f0 "github.com/danil/protofix/codecfix"
	"github.com/danil/protofix/scanfix"
	"github.com/danil/protoscan"
)

var ScanFieldTestCases = []struct {
	line         int
	input        string
	expectedTags []string
	expectedVals []string
}{
	{
		line:         line(),
		input:        "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.000394000|98=0|108=30|10=173|",
		expectedTags: []string{"8", "9", "35", "49", "56", "34", "52", "98", "108", "10"},
		expectedVals: []string{"FIX.4.4", "70", "A", "FG", "tgFZctx20008c", "5", "20210224-12:11:42.000394000", "0", "30", "173"},
	},
}

func TestScanField(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ScanFieldTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			input := strings.ReplaceAll(tc.input, "|", "\x01")
			r := bytes.NewReader([]byte(input))
			split := &scanfix.Field{Format: format}
			scan := protoscan.Protoscan{
				Reader: r,
				Split:  split.Split,
			}

			var tags []string
			var vars []string

			for scan.Scan() {
				tags = append(tags, strconv.Itoa(split.Tag))
				vars = append(vars, string(scan.Token()))

				expectedGap := fmt.Sprintf("%d=\x01", split.Tag)
				recievedGap := string(bytes.Join(scan.Gap(), []byte{}))
				if recievedGap != expectedGap {
					t.Fatalf("unexpected scan gap, expected, %q, recieved: %q %s", expectedGap, recievedGap, linkToExample)
				}
			}

			err := scan.Err()
			if err != nil {
				t.Fatalf("unexpected scan error: %v %s", err, linkToExample)
			}

			gap := scan.Gap()
			if gap != nil {
				t.Fatalf("unexpected scan gap: %v %s", gap, linkToExample)
			}

			if strings.Join(tags, "") != strings.Join(tc.expectedTags, "") {
				t.Errorf("unexpected tag, expected: %#v, received: %#v %s", tc.expectedTags, tags, linkToExample)
			}

			if strings.Join(vars, "") != strings.Join(tc.expectedVals, "") {
				t.Errorf("unexpected variables, expected: %q, received: %q %s", tc.expectedVals, vars, linkToExample)
			}
		})
	}
}

var format = f0.Format{
	Head: map[int]f0.Codec{
		f0.BeginString8:   f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		f0.MsgType35:      f0.Fld{Req, f0.ASCII, f0.StringDefault(), f0.Var{1, 2}},
		f0.SenderCompID49: f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 12}},
		f0.TargetCompID56: f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, f0.MaxString}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, f0.MaxInt}},
	Body: map[int]f0.Codec{
		f0.MsgSeqNum34:     f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, f0.MaxInt}},
		f0.SendingTime52:   f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Con{27}},
		f0.EncryptMethod98: f0.Fld{Req, f0.ASCII, f0.IntDefault(0), f0.Con{1}},
		f0.HeartBtInt108:   f0.Fld{Req, f0.ASCII, f0.SecondsDuration(30 * time.Second), f0.Var{1, f0.MaxInt}},
		f0.Password554:     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 8}},
	},
	Checksum10: f0.ChecksumStringFld{f0.ASCII, f0.ChecksumString(), f0.Con{3}},
	Unknown:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, f0.MaxBytes}},
}

const Req = true
