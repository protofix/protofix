// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix42marketdatarequestreject is a format of the FIX.4.2 MarketDataRequestReject message.
package fix42marketdatarequestreject

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX42MarketDataRequestRejectMarshaler   = marshfix.Marshal{Tag: "FIX42", Format: FIX42MarketDataRequestReject}
	FIX42MarketDataRequestRejectUnmarshaler = marshfix.Unmarshal{Tag: "FIX42", Format: FIX42MarketDataRequestReject}
)

// FIX42MarketDataRequestReject is a FIX.4.2 format of the MarketDataRequestReject message which maps the codecs into individual fields.
var FIX42MarketDataRequestReject = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:              f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.2"), f0.Con{7}},
		MsgType35:                 f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "7" /* ADVERTISEMENT */, "8" /* EXECUTION_REPORT */, "9" /* ORDER_CANCEL_REJECT */, "a" /* QUOTE_STATUS_REQUEST */, "A" /* LOGON */, "B" /* NEWS */, "b" /* QUOTE_ACKNOWLEDGEMENT */, "C" /* EMAIL */, "c" /* SECURITY_DEFINITION_REQUEST */, "D" /* ORDER_SINGLE */, "d" /* SECURITY_DEFINITION */, "E" /* ORDER_LIST */, "e" /* SECURITY_STATUS_REQUEST */, "f" /* SECURITY_STATUS */, "F" /* ORDER_CANCEL_REQUEST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "g" /* TRADING_SESSION_STATUS_REQUEST */, "H" /* ORDER_STATUS_REQUEST */, "h" /* TRADING_SESSION_STATUS */, "i" /* MASS_QUOTE */, "j" /* BUSINESS_MESSAGE_REJECT */, "J" /* ALLOCATION */, "K" /* LIST_CANCEL_REQUEST */, "k" /* BID_REQUEST */, "l" /* BID_RESPONSE */, "L" /* LIST_EXECUTE */, "m" /* LIST_STRIKE_PRICE */, "M" /* LIST_STATUS_REQUEST */, "N" /* LIST_STATUS */, "P" /* ALLOCATION_ACK */, "Q" /* DONT_KNOW_TRADE */, "R" /* QUOTE_REQUEST */, "S" /* QUOTE */, "T" /* SETTLEMENT_INSTRUCTIONS */, "V" /* MARKET_DATA_REQUEST */, "W" /* MARKET_DATA_SNAPSHOT_FULL_REFRESH */, "X" /* MARKET_DATA_INCREMENTAL_REFRESH */, "Y" /* MARKET_DATA_REQUEST_REJECT */, "Z" /* QUOTE_CANCEL */), f0.Var{1, 2}},
		SenderCompID49:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:               f0.Fld{Req, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
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
		MessageEncoding347:        f0.Fld{Opt, f0.ASCII, f0.String("EUC-JP" /* EUC_JP */, "ISO-2022-JP" /* ISO_2022_JP */, "SHIFT_JIS" /* SHIFT_JIS */, "UTF-8" /* UTF_8 */), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		OnBehalfOfSendingTime370:  f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		MDReqID262:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MDReqRejReason281:         f0.Fld{Opt, f0.ASCII, f0.String("0" /* UNKNOWN_SYMBOL */, "1" /* DUPLICATE_MDREQID */, "2" /* INSUFFICIENT_BANDWIDTH */, "3" /* INSUFFICIENT_PERMISSIONS */, "4" /* UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE */, "5" /* UNSUPPORTED_MARKETDEPTH */, "6" /* UNSUPPORTED_MDUPDATETYPE */, "7" /* UNSUPPORTED_AGGREGATEDBOOK */, "8" /* UNSUPPORTED_MDENTRYTYPE */), f0.Con{1}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SignatureLength93:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:               f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,              // STRING
		BodyLength9,               // INT
		MsgType35,                 // STRING
		SenderCompID49,            // STRING
		TargetCompID56,            // STRING
		OnBehalfOfCompID115,       // STRING
		DeliverToCompID128,        // STRING
		SecureDataLen90,           // LENGTH
		SecureData91,              // DATA
		MsgSeqNum34,               // INT
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
		LastMsgSeqNumProcessed369, // INT
		OnBehalfOfSendingTime370,  // UTCTIMESTAMP
		MDReqID262,                // STRING
		MDReqRejReason281,         // CHAR
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		SignatureLength93,         // LENGTH
		Signature89,               // DATA
		CheckSum10,                // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8              = 8   // STRING
	BodyLength9               = 9   // INT
	MsgType35                 = 35  // STRING
	SenderCompID49            = 49  // STRING
	TargetCompID56            = 56  // STRING
	OnBehalfOfCompID115       = 115 // STRING
	DeliverToCompID128        = 128 // STRING
	SecureDataLen90           = 90  // LENGTH
	SecureData91              = 91  // DATA
	MsgSeqNum34               = 34  // INT
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
	LastMsgSeqNumProcessed369 = 369 // INT
	OnBehalfOfSendingTime370  = 370 // UTCTIMESTAMP
	MDReqID262                = 262 // STRING
	MDReqRejReason281         = 281 // CHAR
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
