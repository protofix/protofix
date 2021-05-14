// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1confirmationrequest is a format of the FIX.5.0 servicepack 1 ConfirmationRequest message.
package fix50sp1confirmationrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1ConfirmationRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1ConfirmationRequest}
	FIX50SP1ConfirmationRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1ConfirmationRequest}
)

// FIX50SP1ConfirmationRequest is a FIX.5.0 servicepack 1 format of the ConfirmationRequest message which maps the codecs into individual fields.
var FIX50SP1ConfirmationRequest = f0.Format{
	Fields: map[int]f0.Codec{
		ConfirmReqID859:          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ConfirmType773:           f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* STATUS */, 2 /* CONFIRMATION */, 3 /* CONFIRMATION_REQUEST_REJECTED */), f0.Var{1, 19}},
		AllocID70:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryAllocID793:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IndividualAllocID467:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TransactTime60:           f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		AllocAccount79:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AllocAcctIDSource661:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		AllocAccountType798:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* ACCOUNT_IS_CARRIED_PN_CUSTOMER_SIDE_OF_BOOKS */, 2 /* ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS */, 3 /* HOUSE_TRADER */, 4 /* FLOOR_TRADER */, 6 /* ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED */, 7 /* ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED */, 8 /* JOINT_BACK_OFFICE_ACCOUNT */), f0.Var{1, 19}},
		Text58:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:        f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:           f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		ClOrdID11:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OrderID37:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryOrderID198:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryClOrdID526:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ListID66:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OrderQty38:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrderAvgPx799:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrderBookingQty800:       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Nested2PartyID757:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Nested2PartyIDSource758:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		Nested2PartyRole759:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		Nested2PartySubID760:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Nested2PartySubIDType807: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		ConfirmReqID859,          // STRING
		ConfirmType773,           // INT
		AllocID70,                // STRING
		SecondaryAllocID793,      // STRING
		IndividualAllocID467,     // STRING
		TransactTime60,           // UTCTIMESTAMP
		AllocAccount79,           // STRING
		AllocAcctIDSource661,     // INT
		AllocAccountType798,      // INT
		Text58,                   // STRING
		EncodedTextLen354,        // LENGTH
		EncodedText355,           // DATA
		ClOrdID11,                // STRING
		OrderID37,                // STRING
		SecondaryOrderID198,      // STRING
		SecondaryClOrdID526,      // STRING
		ListID66,                 // STRING
		OrderQty38,               // QTY
		OrderAvgPx799,            // PRICE
		OrderBookingQty800,       // QTY
		Nested2PartyID757,        // STRING
		Nested2PartyIDSource758,  // CHAR
		Nested2PartyRole759,      // INT
		Nested2PartySubID760,     // STRING
		Nested2PartySubIDType807, // INT
	},
}

const Req, Opt = true, false

const (
	ConfirmReqID859          = 859 // STRING
	ConfirmType773           = 773 // INT
	AllocID70                = 70  // STRING
	SecondaryAllocID793      = 793 // STRING
	IndividualAllocID467     = 467 // STRING
	TransactTime60           = 60  // UTCTIMESTAMP
	AllocAccount79           = 79  // STRING
	AllocAcctIDSource661     = 661 // INT
	AllocAccountType798      = 798 // INT
	Text58                   = 58  // STRING
	EncodedTextLen354        = 354 // LENGTH
	EncodedText355           = 355 // DATA
	ClOrdID11                = 11  // STRING
	OrderID37                = 37  // STRING
	SecondaryOrderID198      = 198 // STRING
	SecondaryClOrdID526      = 526 // STRING
	ListID66                 = 66  // STRING
	OrderQty38               = 38  // QTY
	OrderAvgPx799            = 799 // PRICE
	OrderBookingQty800       = 800 // QTY
	Nested2PartyID757        = 757 // STRING
	Nested2PartyIDSource758  = 758 // CHAR
	Nested2PartyRole759      = 759 // INT
	Nested2PartySubID760     = 760 // STRING
	Nested2PartySubIDType807 = 807 // INT
)
