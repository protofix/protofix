// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moex44_test

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/danil/equal4"
	f0 "github.com/danil/protofix/codecfix"
	"github.com/danil/protofix/marshfix"
	"github.com/danil/protofix/moex44"
)

var UnmarshalTestCases = []struct {
	line            int
	name            string
	unmarshaler     marshfix.Unmarshal
	input           string
	expected        moex44Logon
	expectedUnknown map[int][]byte
	expectedWarns   []error
	expectedError   error
	benchmark       bool
}{
	{
		line:        line(),
		name:        "readme example",
		unmarshaler: moex44.MOEX44LogonUnmarshaler,
		input:       "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
		expected: moex44Logon{
			BeginString8:    "FIX.4.4",
			BodyLength9:     103,
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
			CheckSum10:      "060",
		},
	},
	{
		line:        line(),
		name:        "many fields",
		unmarshaler: moex44.MOEX44LogonUnmarshaler,
		input:       "8=FIX.4.4|9=162|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-15:25:21.000000123|98=0|108=42|122=20210312-12:35:12.000000000|141=Y|554=87654321|925=87654321|1409=3|6867=A|6936=R|10=166|",
		expected: moex44Logon{
			BeginString8:           "FIX.4.4",
			BodyLength9:            162,
			MsgType35:              "A",
			SenderCompID49:         "Foo",
			TargetCompID56:         "Bar",
			MsgSeqNum34:            1,
			PossDupFlag43:          false,
			PossResend97:           false,
			SendingTime52:          time.Date(2021, time.March, 12, 15, 25, 21, 123, time.UTC),
			OrigSendingTime122:     time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98:        0,
			HeartBtInt108:          42 * time.Second,
			ResetSeqNumFlag141:     true,
			Password554:            "87654321",
			NewPassword925:         "87654321",
			SessionStatus1409:      3,
			CancelOnDisconnect6867: "A",
			LanguageID6936:         "R",
			CheckSum10:             "166",
		},
	},
	{
		line:        line(),
		name:        `invalid 108/"heart beat interval" tag value (more than 60 seconds)`,
		unmarshaler: moex44.MOEX44LogonUnmarshaler,
		input:       "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=100500|141=N|554=87654321|1409=0|10=060|",
		expected: moex44Logon{
			BeginString8:    "FIX.4.4",
			BodyLength9:     103,
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   100500 * time.Second,
			Password554:     "87654321",
			CheckSum10:      "060",
		},
		expectedWarns: []error{
			errors.New(`FIX unmarshal: decode field tag 108 "HeartBtInt": unexpected value "27h55m0s", allow: [1s 2s 3s 4s 5s 6s 7s 8s 9s 10s 11s 12s 13s 14s 15s 16s 17s 18s 19s 20s 21s 22s 23s 24s 25s 26s 27s 28s 29s 30s 31s 32s 33s 34s 35s 36s 37s 38s 39s 40s 41s 42s 43s 44s 45s 46s 47s 48s 49s 50s 51s 52s 53s 54s 55s 56s 57s 58s 59s 1m0s]`),
		},
	},
	{
		line: line(),
		name: `missing codec of the field tag 35/"message type"`,
		unmarshaler: func() marshfix.Unmarshal {
			head := map[int]f0.Codec{}
			for k, v := range moex44.MOEX44LogonUnmarshaler.Format.Fields {
				head[k] = v
			}
			delete(head, f0.MsgType35)
			marsh := moex44.MOEX44LogonUnmarshaler
			marsh.Format.Fields = head
			return marsh
		}(),
		input: "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
		expected: moex44Logon{
			BeginString8:    "FIX.4.4",
			BodyLength9:     103,
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
			CheckSum10:      "060",
		},
		expectedUnknown: map[int][]byte{f0.MsgType35: []byte("A")},
		expectedWarns: []error{
			errors.New(`FIX unmarshal: missing codec of the field tag 35 "MsgType"`),
			errors.New(`FIX unmarshal: decode field tag 35 "MsgType": cannot deserialize unknown value`),
		},
	},
	{
		line: line(),
		name: `missing codec of the field tag 9/"body length"`,
		unmarshaler: func() marshfix.Unmarshal {
			marsh := moex44.MOEX44LogonUnmarshaler
			marsh.Format.BodyLength9 = f0.BodyLengthFld{}
			return marsh
		}(),
		input: "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
		expected: moex44Logon{
			BeginString8:    "FIX.4.4",
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
			CheckSum10:      "060",
		},
		expectedUnknown: map[int][]byte{f0.BodyLength9: []byte("103")},
		expectedWarns: []error{
			errors.New(`FIX unmarshal: missing codec of the field tag 9 "BodyLength"`),
			errors.New(`FIX unmarshal: decode field tag 9 "BodyLength": cannot deserialize unknown value`),
		},
	},
	{
		line: line(),
		name: "missing codec of the field tag 10/checksum",
		unmarshaler: func() marshfix.Unmarshal {
			marsh := moex44.MOEX44LogonUnmarshaler
			marsh.Format.CheckSum10 = f0.ChecksumStringFld{}
			return marsh
		}(),
		input: "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
		expected: moex44Logon{
			BeginString8:    "FIX.4.4",
			BodyLength9:     103,
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
		},
		expectedUnknown: map[int][]byte{f0.CheckSum10: []byte("060")},
		expectedWarns: []error{
			errors.New(`FIX unmarshal: missing codec of the field tag 10 "Checksum"`),
			errors.New(`FIX unmarshal: decode field tag 10 "Checksum": cannot deserialize unknown value`),
		},
	},
	{
		line: line(),
		name: "missing tag of the unmarshaler",
		unmarshaler: func() marshfix.Unmarshal {
			marsh := moex44.MOEX44LogonUnmarshaler
			marsh.Tag = ""
			return marsh
		}(),
		input:         "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=100500|141=N|554=87654321|1409=0|10=060|",
		expectedError: errors.New(`FIX unmarshal: missing tag of the unmarshal, struct: &{BeginString8: BodyLength9:0 MsgType35: SenderCompID49: TargetCompID56: MsgSeqNum34:0 PossDupFlag43:false PossResend97:false SendingTime52:0001-01-01 00:00:00 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:0s ResetSeqNumFlag141:false Password554: NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
	{
		line: line(),
		name: "missing format of the unmarshaler",
		unmarshaler: func() marshfix.Unmarshal {
			marsh := moex44.MOEX44LogonUnmarshaler
			marsh.Format = f0.Format{}
			return marsh
		}(),
		input:         "8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|97=N|52=20210312-12:35:12.000000000|56=Bar|98=0|108=100500|141=N|554=87654321|1409=0|10=060|",
		expectedError: errors.New(`FIX unmarshal: unmarshaler format validation: missing fields, missing sorting order, unexpected sorting order length: 0, struct: &{BeginString8: BodyLength9:0 MsgType35: SenderCompID49: TargetCompID56: MsgSeqNum34:0 PossDupFlag43:false PossResend97:false SendingTime52:0001-01-01 00:00:00 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:0s ResetSeqNumFlag141:false Password554: NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
}

func TestUmmarshal(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range UnmarshalTestCases {
		tc := tc
		t.Run(fmt.Sprintf("unmarshal %d", tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			logon := moex44Logon{}
			input := bytes.ReplaceAll([]byte(tc.input), []byte{'|'}, []byte{0x01})
			unknown, warns, err := tc.unmarshaler.Unmarshal(input, &logon)
			if !equal4.ErrorEqual(err, tc.expectedError) {
				t.Errorf("unexpected unmarshal error, expected: %s, received: %s %s", tc.expectedError, err, linkToExample)
			}

			if logon != tc.expected {
				t.Errorf("unexpected FIX message, expected: %v, received: %v, input: %q  %s", tc.expected, logon, tc.input, linkToExample)
			}

			if !reflect.DeepEqual(unknown, tc.expectedUnknown) {
				t.Errorf("unexpected unknown fields , expected: %v, received: %v, input: %q  %s", tc.expectedUnknown, unknown, tc.input, linkToExample)
			}

			if len(warns) != len(tc.expectedWarns) {
				t.Fatalf("unexpected validation result, expected: len=%d %+v, received: len=%d %+v %s", len(tc.expectedWarns), tc.expectedWarns, len(warns), warns, linkToExample)
			}

			for i := 0; i < len(warns); i++ {
				if !equal4.ErrorEqual(warns[i], tc.expectedWarns[i]) {
					t.Errorf("unexpected validation error, expected: %s, received: %s %s", tc.expectedWarns[i], warns[i], linkToExample)
				}
			}
		})
	}
}

var MarshalTestCases = []struct {
	line            int
	name            string
	marshaler       marshfix.Marshal
	input           moex44Logon
	expected        string
	expectedUnknown map[int]interface{}
	expectedWarns   []error
	expectedError   error
	benchmark       bool
}{
	{
		line:      line(),
		name:      "readme example",
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
		},
		expected: "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
	},
	{
		line:      line(),
		name:      "many fields",
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			MsgType35:              "A",
			SenderCompID49:         "Foo",
			TargetCompID56:         "Bar",
			MsgSeqNum34:            1,
			PossDupFlag43:          false,
			PossResend97:           false,
			SendingTime52:          time.Date(2021, time.March, 12, 15, 25, 21, 123, time.UTC),
			OrigSendingTime122:     time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98:        0,
			HeartBtInt108:          42 * time.Second,
			ResetSeqNumFlag141:     true,
			Password554:            "87654321",
			NewPassword925:         "87654321",
			SessionStatus1409:      3,
			CancelOnDisconnect6867: "A",
			LanguageID6936:         "R",
		},
		expected:  "8=FIX.4.4|9=162|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-15:25:21.000000123|122=20210312-12:35:12.000000000|98=0|108=42|141=Y|554=87654321|925=87654321|1409=3|6867=A|6936=R|10=166|",
		benchmark: true,
	},
	{
		line: line(),
		name: `missing codec of the field tag 35/"message type"`,
		marshaler: func() marshfix.Marshal {
			head := map[int]f0.Codec{}
			for k, v := range moex44.MOEX44LogonMarshaler.Format.Fields {
				head[k] = v
			}
			delete(head, f0.MsgType35)
			marsh := moex44.MOEX44LogonMarshaler
			marsh.Format.Fields = head
			return marsh
		}(),
		input: moex44Logon{
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
		},
		expected:        "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
		expectedUnknown: map[int]interface{}{f0.MsgType35: "A"},
		expectedWarns: []error{
			errors.New(`FIX marshal: missing codec of the field tag 35 "MsgType"`),
		},
	},
	{
		line:      line(),
		name:      `values of 9/"body length" and 10/checksum will calculated`,
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			BodyLength9:    0,
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
			CheckSum10:     "",
		},
		expected: "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
	},
	{
		line:      line(),
		name:      `8/"begin string" and 35/"message type" and 98/"encrypt method" supplied by default`,
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			BeginString8:    "",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
		},
		expected: "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
	},
	{
		line:      line(),
		name:      `explicit 8/"begin string" and 35/"message type" is ok`,
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			BeginString8:   "FIX.4.4",
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
		},
		expected: "8=FIX.4.4|9=103|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=060|",
	},
	{
		line:      line(),
		name:      `wrong values of the of 9/"body length" and 10/checksum not overwrite`,
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			BodyLength9:    1234567890,
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
			CheckSum10:     "666",
		},
		expected: "8=FIX.4.4|9=1234567890|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|10=666|",
	},
	{
		line:      line(),
		name:      `invalid 108/"heart beat interval" tag value (less than 1 second)`,
		marshaler: moex44.MOEX44LogonMarshaler,
		input: moex44Logon{
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  100500 * time.Second,
			Password554:    "87654321",
		},
		expected: "8=FIX.4.4|9=107|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=100500|141=N|554=87654321|1409=0|10=000|",
		expectedWarns: []error{
			errors.New(`FIX marshal: encode field tag 108 "HeartBtInt": unexpected value "27h55m0s", allow: [1s 2s 3s 4s 5s 6s 7s 8s 9s 10s 11s 12s 13s 14s 15s 16s 17s 18s 19s 20s 21s 22s 23s 24s 25s 26s 27s 28s 29s 30s 31s 32s 33s 34s 35s 36s 37s 38s 39s 40s 41s 42s 43s 44s 45s 46s 47s 48s 49s 50s 51s 52s 53s 54s 55s 56s 57s 58s 59s 1m0s]`),
		},
	},
	{
		line: line(),
		name: `missing 9/"body length" codec`,
		marshaler: func() marshfix.Marshal {
			marsh := moex44.MOEX44LogonMarshaler
			marsh.Format.BodyLength9 = f0.BodyLengthFld{}
			return marsh
		}(),
		input: moex44Logon{
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
		},
		expectedError: errors.New(`FIX marshal: missing codec of the mandatory field tag 9 "BodyLength", struct: &{BeginString8: BodyLength9:0 MsgType35:A SenderCompID49:Foo TargetCompID56:Bar MsgSeqNum34:1 PossDupFlag43:false PossResend97:false SendingTime52:2021-03-12 12:35:12 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:42s ResetSeqNumFlag141:false Password554:87654321 NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
	{
		line: line(),
		name: "missing 10/checksum codec",
		marshaler: func() marshfix.Marshal {
			marsh := moex44.MOEX44LogonMarshaler
			marsh.Format.CheckSum10 = f0.ChecksumStringFld{}
			return marsh
		}(),
		input: moex44Logon{
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
		},
		expectedError: errors.New(`FIX marshal: missing codec of the mandatory field tag 10 "Checksum", struct: &{BeginString8: BodyLength9:0 MsgType35:A SenderCompID49:Foo TargetCompID56:Bar MsgSeqNum34:1 PossDupFlag43:false PossResend97:false SendingTime52:2021-03-12 12:35:12 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:42s ResetSeqNumFlag141:false Password554:87654321 NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
	{
		line: line(),
		name: "missing tag of the unmarshaler",
		marshaler: func() marshfix.Marshal {
			marsh := moex44.MOEX44LogonMarshaler
			marsh.Tag = ""
			return marsh
		}(),
		input: moex44Logon{
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
		},
		expectedError: errors.New(`FIX marshal: missing tag of the unmarshal, struct: &{BeginString8: BodyLength9:0 MsgType35:A SenderCompID49:Foo TargetCompID56:Bar MsgSeqNum34:1 PossDupFlag43:false PossResend97:false SendingTime52:2021-03-12 12:35:12 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:42s ResetSeqNumFlag141:false Password554:87654321 NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
	{
		line: line(),
		name: "missing format of the unmarshaler",
		marshaler: func() marshfix.Marshal {
			marsh := moex44.MOEX44LogonMarshaler
			marsh.Format = f0.Format{}
			return marsh
		}(),
		input: moex44Logon{
			MsgType35:      "A",
			SenderCompID49: "Foo",
			TargetCompID56: "Bar",
			MsgSeqNum34:    1,
			SendingTime52:  time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			HeartBtInt108:  42 * time.Second,
			Password554:    "87654321",
		},
		expectedError: errors.New(`FIX marshal: marshaler format validation: missing fields, missing sorting order, unexpected sorting order length: 0, struct: &{BeginString8: BodyLength9:0 MsgType35:A SenderCompID49:Foo TargetCompID56:Bar MsgSeqNum34:1 PossDupFlag43:false PossResend97:false SendingTime52:2021-03-12 12:35:12 +0000 UTC OrigSendingTime122:0001-01-01 00:00:00 +0000 UTC EncryptMethod98:0 HeartBtInt108:42s ResetSeqNumFlag141:false Password554:87654321 NewPassword925: SessionStatus1409:0 CancelOnDisconnect6867: LanguageID6936: CheckSum10:}`),
	},
}

func TestMarshal(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range MarshalTestCases {
		tc := tc
		t.Run(fmt.Sprintf("marshal %d", tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			p, unknown, warns, err := tc.marshaler.Marshal(&tc.input)
			if !equal4.ErrorEqual(err, tc.expectedError) {
				t.Errorf("unexpected marshal error, expected: %s, received: %s %s", tc.expectedError, err, linkToExample)
			}

			received := string(bytes.ReplaceAll(p, []byte{0x01}, []byte{'|'}))
			if received != tc.expected {
				t.Errorf("unexpected FIX message, expected: %q, received: %q, input: %+v  %s", tc.expected, received, tc.input, linkToExample)
			}

			if !reflect.DeepEqual(unknown, tc.expectedUnknown) {
				t.Errorf("unexpected unknown fields , expected: %v, received: %v, input: %+v  %s", tc.expectedUnknown, unknown, tc.input, linkToExample)
			}

			if len(warns) != len(tc.expectedWarns) {
				t.Fatalf("unexpected validation result, expected: len=%d %+v, received: len=%d %+v %s", len(tc.expectedWarns), tc.expectedWarns, len(warns), warns, linkToExample)
			}

			for i := 0; i < len(warns); i++ {
				if !equal4.ErrorEqual(warns[i], tc.expectedWarns[i]) {
					t.Errorf("unexpected validation error, expected: %s, received: %s %s", tc.expectedWarns[i], warns[i], linkToExample)
				}
			}
		})
	}
}

func BenchmarkMarshal(b *testing.B) {
	for _, tc := range MarshalTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("marshal %s %d", tc.name, tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _, err := moex44.MOEX44LogonMarshaler.Marshal(&tc.input)
				if err != nil {
					b.Fatalf("unexpected marshal error: %s", err)
				}
			}
		})
	}
}

type moex44Logon struct {
	BeginString8           string        `MOEX44:"8"`
	BodyLength9            int           `MOEX44:"9"`
	MsgType35              string        `MOEX44:"35"`
	SenderCompID49         string        `MOEX44:"49"`
	TargetCompID56         string        `MOEX44:"56"`
	MsgSeqNum34            int           `MOEX44:"34"`
	PossDupFlag43          bool          `MOEX44:"43"`
	PossResend97           bool          `MOEX44:"97"`
	SendingTime52          time.Time     `MOEX44:"52"`
	OrigSendingTime122     time.Time     `MOEX44:"122"`
	EncryptMethod98        int           `MOEX44:"98"`
	HeartBtInt108          time.Duration `MOEX44:"108"`
	ResetSeqNumFlag141     bool          `MOEX44:"141"`
	Password554            string        `MOEX44:"554"`
	NewPassword925         string        `MOEX44:"925"`
	SessionStatus1409      int           `MOEX44:"1409"`
	CancelOnDisconnect6867 string        `MOEX44:"6867"`
	LanguageID6936         string        `MOEX44:"6936"`
	CheckSum10             string        `MOEX44:"10"`
}

var MarshalUnknownTestCases = []struct {
	line            int
	name            string
	input           logonWithUnknown
	expected        string
	expectedUnknown map[int]interface{}
	expectedWarns   []error
	expectedError   error
	benchmark       bool
}{
	{
		line: line(),
		name: "two unknown fields",
		input: logonWithUnknown{
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
			Unknown100500:   "Hello, World!",
			Unknown100501:   "Hello, Work!",
		},
		expected:        "8=FIX.4.4|9=144|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|100500=Hello, World!|100501=Hello, Work!|10=119|",
		expectedUnknown: map[int]interface{}{100500: "Hello, World!", 100501: "Hello, Work!"},
		expectedWarns: []error{
			errors.New(`FIX marshal: missing codec of the field tag 100500 ""`),
			errors.New(`FIX marshal: missing codec of the field tag 100501 ""`),
		},
	},
	{
		line: line(),
		name: "two unknown fields but second with zero value",
		input: logonWithUnknown{
			MsgType35:       "A",
			SenderCompID49:  "Foo",
			TargetCompID56:  "Bar",
			MsgSeqNum34:     1,
			SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
			EncryptMethod98: 0,
			HeartBtInt108:   42 * time.Second,
			Password554:     "87654321",
			Unknown100500:   "Hello, World!",
			Unknown100501:   "",
		},
		expected:        "8=FIX.4.4|9=132|35=A|49=Foo|56=Bar|34=1|43=N|97=N|52=20210312-12:35:12.000000000|98=0|108=42|141=N|554=87654321|1409=0|100500=Hello, World!|100501=|10=112|",
		expectedUnknown: map[int]interface{}{100500: "Hello, World!", 100501: ""},
		expectedWarns: []error{
			errors.New(`FIX marshal: missing codec of the field tag 100500 ""`),
			errors.New(`FIX marshal: missing codec of the field tag 100501 ""`),
		},
	},
}

func TestMarshalUnknown(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range MarshalUnknownTestCases {
		tc := tc
		t.Run(fmt.Sprintf("marshal %d", tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			p, unknown, warns, err := moex44.MOEX44LogonMarshaler.Marshal(&tc.input)
			if !equal4.ErrorEqual(err, tc.expectedError) {
				t.Errorf("unexpected marshal error, expected: %s, received: %s %s", tc.expectedError, err, linkToExample)
			}

			received := string(bytes.ReplaceAll(p, []byte{0x01}, []byte{'|'}))
			if received != tc.expected {
				t.Errorf("unexpected FIX message, expected: %q, received: %q, input: %+v  %s", tc.expected, received, tc.input, linkToExample)
			}

			if fmt.Sprint(unknown) != fmt.Sprint(tc.expectedUnknown) {
				t.Errorf("unexpected unknown field, expected: %q, received: %q, input: %+v  %s", tc.expectedUnknown, unknown, tc.input, linkToExample)
			}

			if len(warns) != len(tc.expectedWarns) {
				t.Fatalf("unexpected validation result, expected: %+v, received: %+v %s", tc.expectedWarns, warns, linkToExample)
			}

			for i := 0; i < len(warns); i++ {
				if !equal4.ErrorEqual(warns[i], tc.expectedWarns[i]) {
					t.Errorf("unexpected validation error, expected: %s, received: %s %s", tc.expectedWarns[i], warns[i], linkToExample)
				}
			}
		})
	}
}

type logonWithUnknown struct {
	BeginString8       string        `MOEX44:"8"`
	BodyLength9        int           `MOEX44:"9"`
	MsgType35          string        `MOEX44:"35"`
	SenderCompID49     string        `MOEX44:"49"`
	TargetCompID56     string        `MOEX44:"56"`
	MsgSeqNum34        int           `MOEX44:"34"`
	PossDupFlag43      bool          `MOEX44:"43"`
	PossResend97       bool          `MOEX44:"97"`
	SendingTime52      time.Time     `MOEX44:"52"`
	EncryptMethod98    int           `MOEX44:"98"`
	HeartBtInt108      time.Duration `MOEX44:"108"`
	ResetSeqNumFlag141 bool          `MOEX44:"141"`
	Password554        string        `MOEX44:"554"`
	SessionStatus1409  int           `MOEX44:"1409"`
	Unknown100500      string        `MOEX44:"100500"`
	Unknown100501      string        `MOEX44:"100501"`
	CheckSum10         string        `MOEX44:"10"`
}

func line() int { _, _, l, _ := runtime.Caller(1); return l }
