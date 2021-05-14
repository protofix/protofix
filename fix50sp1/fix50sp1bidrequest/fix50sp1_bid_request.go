// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1bidrequest is a format of the FIX.5.0 servicepack 1 BidRequest message.
package fix50sp1bidrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1BidRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1BidRequest}
	FIX50SP1BidRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1BidRequest}
)

// FIX50SP1BidRequest is a FIX.5.0 servicepack 1 format of the BidRequest message which maps the codecs into individual fields.
var FIX50SP1BidRequest = f0.Format{
	Fields: map[int]f0.Codec{
		BidID390:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClientBidID391:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		BidRequestTransType374:    f0.Fld{Req, f0.ASCII, f0.String("C" /* CANCEL */, "N" /* NO */), f0.Con{1}},
		ListName392:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TotNoRelatedSym393:        f0.Fld{Req, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		BidType394:                f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* NON_DISCLOSED_STYLE */, 2 /* DISCLOSED_SYTLE */, 3 /* NO_BIDDING_PROCESS */), f0.Var{1, 19}},
		NumTickets395:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		Currency15:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		SideValue1396:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SideValue2397:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LiquidityIndType409:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* 5_DAY_MOVING_AVERAGE */, 2 /* 20_DAY_MOVING_AVERAGE */, 3 /* NORMAL_MARKET_SIZE */, 4 /* OTHER */), f0.Var{1, 19}},
		WtAverageLiquidity410:     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ExchangeForPhysical411:    f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		OutMainCntryUIndex412:     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CrossPercent413:           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ProgRptReqs414:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BUY_SIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUE_REQUEST */, 2 /* SELL_SIDE_PERIODICALLY_SENDS_STATUS_USING_LIST_STATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD */, 3 /* REAL_TIME_EXECUTION_REPORTS */), f0.Var{1, 19}},
		ProgPeriodInterval415:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		IncTaxInd416:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* NET */, 2 /* GROSS */), f0.Var{1, 19}},
		ForexReq121:               f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		NumBidders417:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		TradeDate75:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		BidTradeType418:           f0.Fld{Req, f0.ASCII, f0.String("A" /* AGENCY */, "G" /* VWAP_GUARANTEE */, "J" /* GUARANTEED_CLOSE */, "R" /* RISK_TRADE */), f0.Con{1}},
		BasisPxType419:            f0.Fld{Req, f0.ASCII, f0.String("2" /* CLOSING_PRICE_AT_MORNINGN_SESSION */, "3" /* CLOSING_PRICE */, "4" /* CURRENT_PRICE */, "5" /* SQ */, "6" /* VWAP_THROUGH_A_DAY */, "7" /* VWAP_THROUGH_A_MORNING_SESSION */, "8" /* VWAP_THROUGH_AN_AFTERNOON_SESSION */, "9" /* VWAP_THROUGH_A_DAY_EXCEPT_YORI */, "A" /* VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI */, "B" /* VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI */, "C" /* STRIKE */, "D" /* OPEN */, "Z" /* OTHERS */), f0.Con{1}},
		StrikeTime443:             f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		BidDescriptorType399:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* SECTOR */, 2 /* COUNTRY */, 3 /* INDEX */), f0.Var{1, 19}},
		BidDescriptor400:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SideValueInd401:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* SIDE_VALUE_1 */, 2 /* SIDE_VALUE_2 */), f0.Var{1, 19}},
		LiquidityValue404:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LiquidityNumSecurities441: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		LiquidityPctLow402:        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LiquidityPctHigh403:       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		EFPTrackingError405:       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		FairValue406:              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OutsideIndexPct407:        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ValueOfFutures408:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ListID66:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Side54:                    f0.Fld{Opt, f0.ASCII, f0.String("1" /* BUY */, "2" /* SELL */, "3" /* BUY_MINUS */, "4" /* SELL_PLUS */, "5" /* SELL_SHORT */, "6" /* SELL_SHORT_EXEMPT */, "7" /* UNDISCLOSED */, "8" /* CROSS */, "9" /* CROSS_SHORT */, "A" /* CROSS_SHORT_EXXMPT */, "B" /* AS_DEFINED */, "C" /* OPPOSITE */, "D" /* SUBSCRIBE */, "E" /* REDEEM */, "F" /* LEND */, "G" /* BORROW */), f0.Con{1}},
		TradingSessionID336:       f0.Fld{Opt, f0.ASCII, f0.String("1" /* DAY */, "2" /* HALFDAY */, "3" /* MORNING */, "4" /* AFTERNOON */, "5" /* EVENING */, "6" /* AFTER_HOURS */), f0.Var{1, 65536}},
		TradingSessionSubID625:    f0.Fld{Opt, f0.ASCII, f0.String("1" /* PRE_TRADING */, "2" /* OPENING_OR_OPENING_AUCTION */, "3" /* 3 */, "4" /* CLOSING_OR_CLOSING_AUCTION */, "5" /* POST_TRADING */, "6" /* INTRADAY_AUCTION */, "7" /* QUIESCENT */), f0.Var{1, 65536}},
		NetGrossInd430:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* NET */, 2 /* GROSS */), f0.Var{1, 19}},
		SettlType63:               f0.Fld{Opt, f0.ASCII, f0.String("0" /* REGULAR */, "1" /* CASH */, "2" /* NEXT_DAY */, "3" /* T_PLUS_2 */, "4" /* T_PLUS_3 */, "5" /* T_PLUS_4 */, "6" /* FUTURE */, "7" /* WHEN_AND_IF_ISSUED */, "8" /* SELLERS_OPTION */, "9" /* T_PLUS_5 */, "C" /* FX_SPOT_NEXT_SETTLEMENT */, "B" /* BROKEN_DATE */), f0.Var{1, 65536}},
		SettlDate64:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		Account1:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AcctIDSource660:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BIC */, 2 /* SID_CODE */, 3 /* TFM */, 4 /* OMGEO */, 5 /* DTCC_CODE */, 99 /* OTHER */), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BidID390,                  // STRING
		ClientBidID391,            // STRING
		BidRequestTransType374,    // CHAR
		ListName392,               // STRING
		TotNoRelatedSym393,        // INT
		BidType394,                // INT
		NumTickets395,             // INT
		Currency15,                // CURRENCY
		SideValue1396,             // AMT
		SideValue2397,             // AMT
		LiquidityIndType409,       // INT
		WtAverageLiquidity410,     // PERCENTAGE
		ExchangeForPhysical411,    // BOOLEAN
		OutMainCntryUIndex412,     // AMT
		CrossPercent413,           // PERCENTAGE
		ProgRptReqs414,            // INT
		ProgPeriodInterval415,     // INT
		IncTaxInd416,              // INT
		ForexReq121,               // BOOLEAN
		NumBidders417,             // INT
		TradeDate75,               // LOCALMKTDATE
		BidTradeType418,           // CHAR
		BasisPxType419,            // CHAR
		StrikeTime443,             // UTCTIMESTAMP
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		BidDescriptorType399,      // INT
		BidDescriptor400,          // STRING
		SideValueInd401,           // INT
		LiquidityValue404,         // AMT
		LiquidityNumSecurities441, // INT
		LiquidityPctLow402,        // PERCENTAGE
		LiquidityPctHigh403,       // PERCENTAGE
		EFPTrackingError405,       // PERCENTAGE
		FairValue406,              // AMT
		OutsideIndexPct407,        // PERCENTAGE
		ValueOfFutures408,         // AMT
		ListID66,                  // STRING
		Side54,                    // CHAR
		TradingSessionID336,       // STRING
		TradingSessionSubID625,    // STRING
		NetGrossInd430,            // INT
		SettlType63,               // STRING
		SettlDate64,               // LOCALMKTDATE
		Account1,                  // STRING
		AcctIDSource660,           // INT
	},
}

const Req, Opt = true, false

const (
	BidID390                  = 390 // STRING
	ClientBidID391            = 391 // STRING
	BidRequestTransType374    = 374 // CHAR
	ListName392               = 392 // STRING
	TotNoRelatedSym393        = 393 // INT
	BidType394                = 394 // INT
	NumTickets395             = 395 // INT
	Currency15                = 15  // CURRENCY
	SideValue1396             = 396 // AMT
	SideValue2397             = 397 // AMT
	LiquidityIndType409       = 409 // INT
	WtAverageLiquidity410     = 410 // PERCENTAGE
	ExchangeForPhysical411    = 411 // BOOLEAN
	OutMainCntryUIndex412     = 412 // AMT
	CrossPercent413           = 413 // PERCENTAGE
	ProgRptReqs414            = 414 // INT
	ProgPeriodInterval415     = 415 // INT
	IncTaxInd416              = 416 // INT
	ForexReq121               = 121 // BOOLEAN
	NumBidders417             = 417 // INT
	TradeDate75               = 75  // LOCALMKTDATE
	BidTradeType418           = 418 // CHAR
	BasisPxType419            = 419 // CHAR
	StrikeTime443             = 443 // UTCTIMESTAMP
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	BidDescriptorType399      = 399 // INT
	BidDescriptor400          = 400 // STRING
	SideValueInd401           = 401 // INT
	LiquidityValue404         = 404 // AMT
	LiquidityNumSecurities441 = 441 // INT
	LiquidityPctLow402        = 402 // PERCENTAGE
	LiquidityPctHigh403       = 403 // PERCENTAGE
	EFPTrackingError405       = 405 // PERCENTAGE
	FairValue406              = 406 // AMT
	OutsideIndexPct407        = 407 // PERCENTAGE
	ValueOfFutures408         = 408 // AMT
	ListID66                  = 66  // STRING
	Side54                    = 54  // CHAR
	TradingSessionID336       = 336 // STRING
	TradingSessionSubID625    = 625 // STRING
	NetGrossInd430            = 430 // INT
	SettlType63               = 63  // STRING
	SettlDate64               = 64  // LOCALMKTDATE
	Account1                  = 1   // STRING
	AcctIDSource660           = 660 // INT
)
