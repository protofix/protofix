// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fix44logon_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/protofix/protofix/fix44/fix44logon"
)

func TestFIX44LogonUnmarshalMarshal(t *testing.T) {
	type Logon struct {
		BeginString8       string    `FIX44:"8"`
		BodyLength9        int       `FIX44:"9"`
		MsgType35          string    `FIX44:"35"`
		SenderCompID49     string    `FIX44:"49"`
		TargetCompID56     string    `FIX44:"56"`
		MsgSeqNum34        int       `FIX44:"34"`
		SendingTime52      time.Time `FIX44:"52"`
		EncryptMethod98    int       `FIX44:"98"`
		HeartBtInt108      int       `FIX44:"108"`
		ResetSeqNumFlag141 bool      `FIX44:"141"`
		Username553        string    `FIX44:"553"`
		Password554        string    `FIX44:"554"`
		CheckSum10         string    `FIX44:"10"`
	}

	input := []byte("8=FIX.4.4|9=102|35=A|49=BuySide|56=SellSide|34=1|52=20190605-11:40:30.392|98=0|108=30|141=Y|553=Username|554=Password|10=104|")
	input = bytes.ReplaceAll(input, []byte{'|'}, []byte{0x01})

	logon := Logon{}

	_, _, err := fix44logon.FIX44LogonUnmarshaler.Unmarshal(input, &logon)
	if err != nil {
		t.Errorf("unexpected unmarshal error: %s", err)
	}

	output, _, _, err := fix44logon.FIX44LogonMarshaler.Marshal(&logon)
	if err != nil {
		t.Errorf("unexpected marshal error: %s", err)
	}

	t.Logf("%+v\n", logon)

	input = bytes.ReplaceAll(input, []byte{0x01}, []byte{'|'})
	output = bytes.ReplaceAll(output, []byte{0x01}, []byte{'|'})

	if bytes.Equal(input, output) {
		t.Logf("Message %q are equal to %q.\n", input, output)
	} else {
		t.Errorf("Message %q are not equal to %q!\n", input, output)
	}
}

func BenchmarkLogonUnmarshalMarshal(b *testing.B) {
	type Logon struct {
		BeginString8       string    `FIX44:"8"`
		BodyLength9        int       `FIX44:"9"`
		MsgType35          string    `FIX44:"35"`
		SenderCompID49     string    `FIX44:"49"`
		TargetCompID56     string    `FIX44:"56"`
		MsgSeqNum34        int       `FIX44:"34"`
		SendingTime52      time.Time `FIX44:"52"`
		EncryptMethod98    int       `FIX44:"98"`
		HeartBtInt108      int       `FIX44:"108"`
		ResetSeqNumFlag141 bool      `FIX44:"141"`
		Username553        string    `FIX44:"553"`
		Password554        string    `FIX44:"554"`
		CheckSum10         string    `FIX44:"10"`
	}

	logon := Logon{}

	input := []byte("8=FIX.4.4|9=102|35=A|49=BuySide|56=SellSide|34=1|52=20190605-11:40:30.392|98=0|108=30|141=Y|553=Username|554=Password|10=104|")
	input = bytes.ReplaceAll(input, []byte{'|'}, []byte{0x01})

	for i := 0; i < b.N; i++ {
		_, _, err := fix44logon.FIX44LogonUnmarshaler.Unmarshal(input, &logon)
		if err != nil {
			b.Fatalf("unexpected unmarshal error: %s", err)
		}

		output, _, _, err := fix44logon.FIX44LogonMarshaler.Marshal(&logon)
		if err != nil {
			b.Fatalf("unexpected marshal error: %s", err)
		}

		if !bytes.Equal(input, output) {
			b.Fatalf("messages %q not equal %q", input, output)
		}
	}
}
