// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a FIX message format generator.
package genfix_test

import (
	"reflect"
	"testing"

	"github.com/protofix/protofix/genfix/internal/testspec"
)

func TestReq(t *testing.T) {
	if !testspec.Req {
		t.Errorf("unexpected Req, expected: %t, received: %t", true, testspec.Req)
	}
}

func TestOpt(t *testing.T) {
	if testspec.Opt {
		t.Errorf("unexpected Opt, expected: %t, received: %t", false, testspec.Opt)
	}
}

func TestTESTSPECLogonSort(t *testing.T) {
	recieved := testspec.TESTSPECLogon.Sort
	expected := []int{
		testspec.BeginString8,
		testspec.BodyLength9,
		testspec.MsgType35,
		testspec.SenderCompID49,
		testspec.TargetCompID56,
		testspec.MsgSeqNum34,
		testspec.PossDupFlag43,
		testspec.PossResend97,
		testspec.SendingTime52,
		testspec.OrigSendingTime122,
		testspec.AgreementDesc913,
		testspec.YieldType235,
		testspec.UnderlyingStipType888,
		testspec.HopCompID628,
		testspec.EncryptMethod98,
		testspec.HeartBtInt108,
		testspec.ResetSeqNumFlag141,
		testspec.Password554,
		testspec.NewPassword925,
		testspec.SessionStatus1409,
		testspec.CancelOnDisconnect6867,
		testspec.LanguageID6936,
		testspec.PosType703,
		testspec.NestedPartyID524,
		testspec.NestedPartySubID545,
		testspec.CheckSum10,
	}
	if !reflect.DeepEqual(recieved, expected) {
		t.Errorf("unexpected test spec logon sort order, expected: %v, received: %v", expected, recieved)
	}
}
