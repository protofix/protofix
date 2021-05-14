// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix43bidresponse is a format of the FIX.4.3 BidResponse message.
package fix43bidresponse

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX43BidResponseMarshaler   = marshfix.Marshal{Tag: "FIX43", Format: FIX43BidResponse}
	FIX43BidResponseUnmarshaler = marshfix.Unmarshal{Tag: "FIX43", Format: FIX43BidResponse}
)

// FIX43BidResponse is a FIX.4.3 format of the BidResponse message which maps the codecs into individual fields.
var FIX43BidResponse = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:              f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.3"), f0.Con{7}},
		MsgType35:                 f0.Fld{Req, f0.ASCII, f0.String("m" /* LIST_STRIKE_PRICE */, "l" /* BID_RESPONSE */, "k" /* BID_REQUEST */, "j" /* BUSINESS_MESSAGE_REJECT */, "i" /* MASS_QUOTE */, "h" /* TRADING_SESSION_STATUS */, "q" /* ORDER_MASS_CANCEL_REQUEST */, "AE" /* TRADE_CAPTURE_REPORT */, "o" /* REGISTRATION_INSTRUCTIONS */, "z" /* DERIVATIVE_SECURITY_LIST_REQUEST */, "AI" /* QUOTE_STATUS_REPORT */, "AH" /* RFQ_REQUEST */, "AG" /* QUOTE_REQUEST_REJECT */, "AF" /* ORDER_MASS_STATUS_REQUEST */, "n" /* XML_MESSAGE */, "AD" /* TRADE_CAPTURE_REPORT_REQUEST */, "g" /* TRADING_SESSION_STATUS_REQUEST */, "AC" /* MULTILEG_ORDER_CANCEL_REPLACE */, "AA" /* DERIVATIVE_SECURITY_LIST */, "r" /* ORDER_MASS_CANCEL_REPORT */, "y" /* SECURITY_LIST */, "x" /* SECURITY_LIST_REQUEST */, "w" /* SECURITY_TYPES */, "v" /* SECURITY_TYPE_REQUEST */, "u" /* CROSS_ORDER_CANCEL_REQUEST */, "t" /* CROSS_ORDER_CANCEL_REPLACE_REQUEST */, "s" /* NEW_ORDER_s */, "AB" /* NEW_ORDER_AB */, "9" /* ORDER_CANCEL_REJECT */, "7" /* ADVERTISEMENT */, "M" /* LIST_STATUS_REQUEST */, "E" /* ORDER_LIST */, "D" /* ORDER_SINGLE */, "C" /* EMAIL */, "B" /* NEWS */, "A" /* LOGON */, "N" /* LIST_STATUS */, "f" /* SECURITY_STATUS */, "P" /* ALLOCATION_ACK */, "8" /* EXECUTION_REPORT */, "0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "K" /* LIST_CANCEL_REQUEST */, "e" /* SECURITY_STATUS_REQUEST */, "d" /* SECURITY_DEFINITION */, "c" /* SECURITY_DEFINITION_REQUEST */, "b" /* MASS_QUOTE_ACKNOWLEDGEMENT */, "a" /* QUOTE_STATUS_REQUEST */, "Z" /* QUOTE_CANCEL */, "Y" /* MARKET_DATA_REQUEST_REJECT */, "F" /* ORDER_CANCEL_REQUEST */, "J" /* ALLOCATION */, "p" /* REGISTRATION_INSTRUCTIONS_RESPONSE */, "L" /* LIST_EXECUTE */, "X" /* MARKET_DATA_INCREMENTAL_REFRESH */, "W" /* MARKET_DATA_SNAPSHOT_FULL_REFRESH */, "V" /* MARKET_DATA_REQUEST */, "T" /* SETTLEMENT_INSTRUCTIONS */, "S" /* QUOTE */, "R" /* QUOTE_REQUEST */, "Q" /* DONT_KNOW_TRADE */, "H" /* ORDER_STATUS_REQUEST */), f0.Var{1, 2}},
		SenderCompID49:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:               f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SenderSubID50:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SenderLocationID142:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetSubID57:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetLocationID143:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfSubID116:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfLocationID144:   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToSubID129:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToLocationID145:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PossDupFlag43:             f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:              f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:             f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:        f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		XmlDataLen212:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		XmlData213:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MessageEncoding347:        f0.Fld{Opt, f0.ASCII, f0.String("UTF-8" /* UTF_8 */, "ISO-2022-JP" /* ISO_2022_JP */, "EUC-JP" /* EUC_JP */, "SHIFT_JIS" /* SHIFT_JIS */), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369: f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		OnBehalfOfSendingTime370:  f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopCompID628:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:         f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:               f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		BidID390:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClientBidID391:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SignatureLength93:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:               f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,              // STRING
		BodyLength9,               // LENGTH
		MsgType35,                 // STRING
		SenderCompID49,            // STRING
		TargetCompID56,            // STRING
		OnBehalfOfCompID115,       // STRING
		DeliverToCompID128,        // STRING
		SecureDataLen90,           // LENGTH
		SecureData91,              // DATA
		MsgSeqNum34,               // SEQNUM
		SenderSubID50,             // STRING
		SenderLocationID142,       // STRING
		TargetSubID57,             // STRING
		TargetLocationID143,       // STRING
		OnBehalfOfSubID116,        // STRING
		OnBehalfOfLocationID144,   // STRING
		DeliverToSubID129,         // STRING
		DeliverToLocationID145,    // STRING
		PossDupFlag43,             // BOOLEAN
		PossResend97,              // BOOLEAN
		SendingTime52,             // UTCTIMESTAMP
		OrigSendingTime122,        // UTCTIMESTAMP
		XmlDataLen212,             // LENGTH
		XmlData213,                // DATA
		MessageEncoding347,        // STRING
		LastMsgSeqNumProcessed369, // SEQNUM
		OnBehalfOfSendingTime370,  // UTCTIMESTAMP
		HopCompID628,              // STRING
		HopSendingTime629,         // UTCTIMESTAMP
		HopRefID630,               // SEQNUM
		BidID390,                  // STRING
		ClientBidID391,            // STRING
		SignatureLength93,         // LENGTH
		Signature89,               // DATA
		CheckSum10,                // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8              = 8   // STRING
	BodyLength9               = 9   // LENGTH
	MsgType35                 = 35  // STRING
	SenderCompID49            = 49  // STRING
	TargetCompID56            = 56  // STRING
	OnBehalfOfCompID115       = 115 // STRING
	DeliverToCompID128        = 128 // STRING
	SecureDataLen90           = 90  // LENGTH
	SecureData91              = 91  // DATA
	MsgSeqNum34               = 34  // SEQNUM
	SenderSubID50             = 50  // STRING
	SenderLocationID142       = 142 // STRING
	TargetSubID57             = 57  // STRING
	TargetLocationID143       = 143 // STRING
	OnBehalfOfSubID116        = 116 // STRING
	OnBehalfOfLocationID144   = 144 // STRING
	DeliverToSubID129         = 129 // STRING
	DeliverToLocationID145    = 145 // STRING
	PossDupFlag43             = 43  // BOOLEAN
	PossResend97              = 97  // BOOLEAN
	SendingTime52             = 52  // UTCTIMESTAMP
	OrigSendingTime122        = 122 // UTCTIMESTAMP
	XmlDataLen212             = 212 // LENGTH
	XmlData213                = 213 // DATA
	MessageEncoding347        = 347 // STRING
	LastMsgSeqNumProcessed369 = 369 // SEQNUM
	OnBehalfOfSendingTime370  = 370 // UTCTIMESTAMP
	HopCompID628              = 628 // STRING
	HopSendingTime629         = 629 // UTCTIMESTAMP
	HopRefID630               = 630 // SEQNUM
	BidID390                  = 390 // STRING
	ClientBidID391            = 391 // STRING
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
