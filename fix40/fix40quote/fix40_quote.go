// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix40quote is a format of the FIX.4.0 Quote message.
package fix40quote

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX40QuoteMarshaler   = marshfix.Marshal{Tag: "FIX40", Format: FIX40Quote}
	FIX40QuoteUnmarshaler = marshfix.Unmarshal{Tag: "FIX40", Format: FIX40Quote}
)

// FIX40Quote is a FIX.4.0 format of the Quote message which maps the codecs into individual fields.
var FIX40Quote = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:        f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.0"), f0.Con{7}},
		MsgType35:           f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "7" /* ADVERTISEMENT */, "8" /* EXECUTION_REPORT */, "9" /* ORDER_CANCEL_REJECT */, "A" /* LOGON */, "B" /* NEWS */, "C" /* EMAIL */, "D" /* ORDER_D */, "E" /* ORDER_E */, "F" /* ORDER_CANCEL_REQUEST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "H" /* ORDER_STATUS_REQUEST */, "J" /* ALLOCATION */, "K" /* LIST_CANCEL_REQUEST */, "L" /* LIST_EXECUTE */, "M" /* LIST_STATUS_REQUEST */, "N" /* LIST_STATUS */, "P" /* ALLOCATION_ACK */, "Q" /* DONT_KNOW_TRADE */, "R" /* QUOTE_REQUEST */, "S" /* QUOTE */), f0.Var{1, 2}},
		SenderCompID49:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		TargetCompID56:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		OnBehalfOfCompID115: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		DeliverToCompID128:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SecureDataLen90:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:         f0.Fld{Req, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		SenderSubID50:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		TargetSubID57:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		OnBehalfOfSubID116:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		DeliverToSubID129:   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		PossDupFlag43:       f0.Fld{Opt, f0.ASCII, f0.String("N" /* NO */, "Y" /* YES */), f0.Con{1}},
		PossResend97:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SendingTime52:       f0.Fld{Req, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		OrigSendingTime122:  f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		QuoteReqID131:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		QuoteID117:          f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		Symbol55:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		SymbolSfx65:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SecurityID48:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		IDSource22:          f0.Fld{Opt, f0.ASCII, f0.String("1" /* CUSIP */, "2" /* SEDOL */, "3" /* QUIK */, "4" /* ISIN_NUMBER */, "5" /* RIC_CODE */), f0.Con{1}},
		Issuer106:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SecurityDesc107:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		BidPx132:            f0.Fld{Req, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OfferPx133:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		BidSize134:          f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		OfferSize135:        f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		ValidUntilTime62:    f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		SignatureLength93:   f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:         f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,        // CHAR
		BodyLength9,         // INT
		MsgType35,           // CHAR
		SenderCompID49,      // CHAR
		TargetCompID56,      // CHAR
		OnBehalfOfCompID115, // CHAR
		DeliverToCompID128,  // CHAR
		SecureDataLen90,     // LENGTH
		SecureData91,        // DATA
		MsgSeqNum34,         // INT
		SenderSubID50,       // CHAR
		TargetSubID57,       // CHAR
		OnBehalfOfSubID116,  // CHAR
		DeliverToSubID129,   // CHAR
		PossDupFlag43,       // CHAR
		PossResend97,        // CHAR
		SendingTime52,       // TIME
		OrigSendingTime122,  // TIME
		QuoteReqID131,       // CHAR
		QuoteID117,          // CHAR
		Symbol55,            // CHAR
		SymbolSfx65,         // CHAR
		SecurityID48,        // CHAR
		IDSource22,          // CHAR
		Issuer106,           // CHAR
		SecurityDesc107,     // CHAR
		BidPx132,            // FLOAT
		OfferPx133,          // FLOAT
		BidSize134,          // INT
		OfferSize135,        // INT
		ValidUntilTime62,    // TIME
		SignatureLength93,   // LENGTH
		Signature89,         // DATA
		CheckSum10,          // CHAR
	},
}

const Req, Opt = true, false

const (
	BeginString8        = 8   // CHAR
	BodyLength9         = 9   // INT
	MsgType35           = 35  // CHAR
	SenderCompID49      = 49  // CHAR
	TargetCompID56      = 56  // CHAR
	OnBehalfOfCompID115 = 115 // CHAR
	DeliverToCompID128  = 128 // CHAR
	SecureDataLen90     = 90  // LENGTH
	SecureData91        = 91  // DATA
	MsgSeqNum34         = 34  // INT
	SenderSubID50       = 50  // CHAR
	TargetSubID57       = 57  // CHAR
	OnBehalfOfSubID116  = 116 // CHAR
	DeliverToSubID129   = 129 // CHAR
	PossDupFlag43       = 43  // CHAR
	PossResend97        = 97  // CHAR
	SendingTime52       = 52  // TIME
	OrigSendingTime122  = 122 // TIME
	QuoteReqID131       = 131 // CHAR
	QuoteID117          = 117 // CHAR
	Symbol55            = 55  // CHAR
	SymbolSfx65         = 65  // CHAR
	SecurityID48        = 48  // CHAR
	IDSource22          = 22  // CHAR
	Issuer106           = 106 // CHAR
	SecurityDesc107     = 107 // CHAR
	BidPx132            = 132 // FLOAT
	OfferPx133          = 133 // FLOAT
	BidSize134          = 134 // INT
	OfferSize135        = 135 // INT
	ValidUntilTime62    = 62  // TIME
	SignatureLength93   = 93  // LENGTH
	Signature89         = 89  // DATA
	CheckSum10          = 10  // CHAR
)
