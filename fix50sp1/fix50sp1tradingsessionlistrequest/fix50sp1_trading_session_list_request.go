// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1tradingsessionlistrequest is a format of the FIX.5.0 servicepack 1 TradingSessionListRequest message.
package fix50sp1tradingsessionlistrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1TradingSessionListRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1TradingSessionListRequest}
	FIX50SP1TradingSessionListRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1TradingSessionListRequest}
)

// FIX50SP1TradingSessionListRequest is a FIX.5.0 servicepack 1 format of the TradingSessionListRequest message which maps the codecs into individual fields.
var FIX50SP1TradingSessionListRequest = f0.Format{
	Fields: map[int]f0.Codec{
		TradSesReqID335:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionID336:        f0.Fld{Opt, f0.ASCII, f0.String("1" /* DAY */, "2" /* HALFDAY */, "3" /* MORNING */, "4" /* AFTERNOON */, "5" /* EVENING */, "6" /* AFTER_HOURS */), f0.Var{1, 65536}},
		TradingSessionSubID625:     f0.Fld{Opt, f0.ASCII, f0.String("1" /* PRE_TRADING */, "2" /* OPENING_OR_OPENING_AUCTION */, "3" /* 3 */, "4" /* CLOSING_OR_CLOSING_AUCTION */, "5" /* POST_TRADING */, "6" /* INTRADAY_AUCTION */, "7" /* QUIESCENT */), f0.Var{1, 65536}},
		SecurityExchange207:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradSesMethod338:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* ELECTRONIC */, 2 /* OPEN_OUTCRY */, 3 /* TWO_PARTY */), f0.Var{1, 19}},
		TradSesMode339:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* TESTING */, 2 /* SIMULATED */, 3 /* PRODUCTION */), f0.Var{1, 19}},
		SubscriptionRequestType263: f0.Fld{Req, f0.ASCII, f0.String("0" /* SNAPSHOT */, "1" /* SNAPSHOT_PLUS_UPDATES */, "2" /* DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST */), f0.Con{1}},
		MarketID1301:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketSegmentID1300:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		TradSesReqID335,            // STRING
		TradingSessionID336,        // STRING
		TradingSessionSubID625,     // STRING
		SecurityExchange207,        // EXCHANGE
		TradSesMethod338,           // INT
		TradSesMode339,             // INT
		SubscriptionRequestType263, // CHAR
		MarketID1301,               // EXCHANGE
		MarketSegmentID1300,        // STRING
	},
}

const Req, Opt = true, false

const (
	TradSesReqID335            = 335  // STRING
	TradingSessionID336        = 336  // STRING
	TradingSessionSubID625     = 625  // STRING
	SecurityExchange207        = 207  // EXCHANGE
	TradSesMethod338           = 338  // INT
	TradSesMode339             = 339  // INT
	SubscriptionRequestType263 = 263  // CHAR
	MarketID1301               = 1301 // EXCHANGE
	MarketSegmentID1300        = 1300 // STRING
)
