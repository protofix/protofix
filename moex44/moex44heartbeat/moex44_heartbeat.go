// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package moex44heartbeat is a format of the MOEX.4.4 (FIX.4.4) Heartbeat message.
package moex44heartbeat

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	MOEX44HeartbeatMarshaler   = marshfix.Marshal{Tag: "MOEX44", Format: MOEX44Heartbeat}
	MOEX44HeartbeatUnmarshaler = marshfix.Unmarshal{Tag: "MOEX44", Format: MOEX44Heartbeat}
)

// MOEX44Heartbeat is a MOEX.4.4 (FIX.4.4) format of the Heartbeat message which maps the codecs into individual fields.
var MOEX44Heartbeat = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:       f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:          f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MsgSeqNum34:        f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		PossDupFlag43:      f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:       f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:      f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		OrigSendingTime122: f0.Fld{Opt, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		TestReqID112:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,       // STRING
		BodyLength9,        // LENGTH
		MsgType35,          // STRING
		SenderCompID49,     // STRING
		TargetCompID56,     // STRING
		MsgSeqNum34,        // SEQNUM
		PossDupFlag43,      // BOOLEAN
		PossResend97,       // BOOLEAN
		SendingTime52,      // UTCTIMESTAMP
		OrigSendingTime122, // UTCTIMESTAMP
		TestReqID112,       // STRING
		CheckSum10,         // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8       = 8   // STRING
	BodyLength9        = 9   // LENGTH
	MsgType35          = 35  // STRING
	SenderCompID49     = 49  // STRING
	TargetCompID56     = 56  // STRING
	MsgSeqNum34        = 34  // SEQNUM
	PossDupFlag43      = 43  // BOOLEAN
	PossResend97       = 97  // BOOLEAN
	SendingTime52      = 52  // UTCTIMESTAMP
	OrigSendingTime122 = 122 // UTCTIMESTAMP
	TestReqID112       = 112 // STRING
	CheckSum10         = 10  // STRING
)
