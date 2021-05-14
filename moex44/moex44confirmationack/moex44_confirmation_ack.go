// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package moex44confirmationack is a format of the MOEX.4.4 (FIX.4.4) ConfirmationAck message.
package moex44confirmationack

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	MOEX44ConfirmationAckMarshaler   = marshfix.Marshal{Tag: "MOEX44", Format: MOEX44ConfirmationAck}
	MOEX44ConfirmationAckUnmarshaler = marshfix.Unmarshal{Tag: "MOEX44", Format: MOEX44ConfirmationAck}
)

// MOEX44ConfirmationAck is a MOEX.4.4 (FIX.4.4) format of the ConfirmationAck message which maps the codecs into individual fields.
var MOEX44ConfirmationAck = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:        f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:           f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MsgSeqNum34:         f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		PossDupFlag43:       f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:        f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:       f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:  f0.Fld{Opt, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		ConfirmID664:        f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradeDate75:         f0.Fld{Req, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		TransactTime60:      f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		AffirmStatus940:     f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* RECEIVED */, 2 /* CONFIRMREJECTED */, 3 /* AFFIRMED */), f0.Var{1, 19}},
		ConfirmRejReason774: f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* MISMATCHEDACCOUNT */, 2 /* MISSINGSETTLEMENTINSTRUCTIONS */, 99 /* OTHER */), f0.Var{1, 19}},
		MatchStatus573:      f0.Fld{Opt, f0.ASCII, f0.String("0" /* COMPMATAFF */, "1" /* UNCOMPUNMATUNAFF */, "2" /* ADVALERT */), f0.Con{1}},
		Text58:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:   f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:      f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,        // STRING
		BodyLength9,         // LENGTH
		MsgType35,           // STRING
		SenderCompID49,      // STRING
		TargetCompID56,      // STRING
		MsgSeqNum34,         // SEQNUM
		PossDupFlag43,       // BOOLEAN
		PossResend97,        // BOOLEAN
		SendingTime52,       // UTCTIMESTAMP
		OrigSendingTime122,  // UTCTIMESTAMP
		ConfirmID664,        // STRING
		TradeDate75,         // LOCALMKTDATE
		TransactTime60,      // UTCTIMESTAMP
		AffirmStatus940,     // INT
		ConfirmRejReason774, // INT
		MatchStatus573,      // CHAR
		Text58,              // STRING
		EncodedTextLen354,   // LENGTH
		EncodedText355,      // DATA
		CheckSum10,          // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8        = 8   // STRING
	BodyLength9         = 9   // LENGTH
	MsgType35           = 35  // STRING
	SenderCompID49      = 49  // STRING
	TargetCompID56      = 56  // STRING
	MsgSeqNum34         = 34  // SEQNUM
	PossDupFlag43       = 43  // BOOLEAN
	PossResend97        = 97  // BOOLEAN
	SendingTime52       = 52  // UTCTIMESTAMP
	OrigSendingTime122  = 122 // UTCTIMESTAMP
	ConfirmID664        = 664 // STRING
	TradeDate75         = 75  // LOCALMKTDATE
	TransactTime60      = 60  // UTCTIMESTAMP
	AffirmStatus940     = 940 // INT
	ConfirmRejReason774 = 774 // INT
	MatchStatus573      = 573 // CHAR
	Text58              = 58  // STRING
	EncodedTextLen354   = 354 // LENGTH
	EncodedText355      = 355 // DATA
	CheckSum10          = 10  // STRING
)