// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1marketdefinitionupdatereport is a format of the FIX.5.0 servicepack 1 MarketDefinitionUpdateReport message.
package fix50sp1marketdefinitionupdatereport

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1MarketDefinitionUpdateReportMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1MarketDefinitionUpdateReport}
	FIX50SP1MarketDefinitionUpdateReportUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1MarketDefinitionUpdateReport}
)

// FIX50SP1MarketDefinitionUpdateReport is a FIX.5.0 servicepack 1 format of the MarketDefinitionUpdateReport message which maps the codecs into individual fields.
var FIX50SP1MarketDefinitionUpdateReport = f0.Format{
	Fields: map[int]f0.Codec{
		MarketReportID1394:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketReqID1393:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketUpdateAction1395:     f0.Fld{Opt, f0.ASCII, f0.String("A" /* ADD */, "D" /* DELETE */, "M" /* MODIFY */), f0.Con{1}},
		MarketID1301:               f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketSegmentID1300:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketSegmentDesc1396:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedMktSegmDescLen1397:  f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedMktSegmDesc1398:     f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		ParentMktSegmID1325:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Currency15:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		TransactTime60:             f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Text58:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:          f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:             f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		ExpirationCycle827:         f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* EXPIRE_ON_TRADING_SESSION_CLOSE */, 1 /* EXPIRE_ON_TRADING_SESSION_OPEN */, 2 /* TRADING_ELIGIBILITY_EXPIRATION_SPECIFIED_IN_THE_DATE_AND_TIME_FIELDS_EVENTDATE */), f0.Var{1, 19}},
		MinTradeVol562:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		MaxTradeVol1140:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		MaxPriceVariation1143:      f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ImpliedMarketIndicator1144: f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* NOT_IMPLIED */, 1 /* IMPLIED_IN */, 2 /* IMPLIED_OUT */, 3 /* BOTH_IMPLIED_IN_AND_IMPLIED_OUT */), f0.Var{1, 19}},
		TradingCurrency1245:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		RoundLot561:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		MultilegModel1377:          f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* PREDEFINED_MULTILEG_SECURITY */, 1 /* USER_DEFINED_MULTLEG_SECURITY */, 2 /* USER_DEFINED_NON_SECURITIZED_MULTILEG */), f0.Var{1, 19}},
		MultilegPriceMethod1378:    f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* NET_PRICE */, 1 /* REVERSED_NET_PRICE */, 2 /* YIELD_DIFFERENCE */, 3 /* INDIVIDUAL */, 4 /* CONTRACT_WEIGHTED_AVERAGE_PRICE */, 5 /* MULTIPLIED_PRICE */), f0.Var{1, 19}},
		PriceType423:               f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PERCENTAGE */, 10 /* FIXED_CABINET_TRADE_PRICE */, 11 /* VARIABLE_CABINET_TRADE_PRICE */, 2 /* PER_UNIT */, 3 /* FIXED_AMOUNT */, 4 /* DISCOUNT */, 5 /* PREMIUM */, 6 /* SPREAD */, 7 /* TED_PRICE */, 8 /* TED_YIELD */, 9 /* YIELD */, 13 /* PRODUCT_TICKS_IN_HALFS */, 14 /* PRODUCT_TICKS_IN_FOURTHS */, 15 /* PRODUCT_TICKS_IN_EIGHTS */, 16 /* PRODUCT_TICKS_IN_SIXTEENTHS */, 17 /* PRODUCT_TICKS_IN_THIRTY_SECONDS */, 18 /* PRODUCT_TICKS_IN_SIXTY_FORTHS */, 19 /* PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTS */), f0.Var{1, 19}},
		StartTickPriceRange1206:    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		EndTickPriceRange1207:      f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		TickIncrement1208:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		TickRuleType1209:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* REGULAR */, 1 /* VARIABLE */, 2 /* FIXED */, 3 /* TRADED_AS_A_SPREAD_LEG */, 4 /* SETTLED_AS_A_SPREAD_LEG */), f0.Var{1, 19}},
		LotType1093:                f0.Fld{Opt, f0.ASCII, f0.String("1" /* ODD_LOT */, "2" /* ROUND_LOT */, "3" /* BLOCK_LOT */), f0.Con{1}},
		MinLotSize1231:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		PriceLimitType1306:         f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* PRICE */, 1 /* TICKS */, 2 /* PERCENTAGE */), f0.Var{1, 19}},
		LowLimitPrice1148:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		HighLimitPrice1149:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		TradingReferencePrice1150:  f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrdType40:                  f0.Fld{Opt, f0.ASCII, f0.String("1" /* MARKET */, "2" /* LIMIT */, "3" /* STOP */, "4" /* STOP_LIMIT */, "5" /* MARKET_ON_CLOSE */, "6" /* WITH_OR_WITHOUT */, "7" /* LIMIT_OR_BETTER */, "8" /* LIMIT_WITH_OR_WITHOUT */, "9" /* ON_BASIS */, "A" /* ON_CLOSE */, "B" /* LIMIT_ON_CLOSE */, "C" /* FOREX_MARKET */, "D" /* PREVIOUSLY_QUOTED */, "E" /* PREVIOUSLY_INDICATED */, "F" /* FOREX_LIMIT */, "G" /* FOREX_SWAP */, "H" /* FOREX_PREVIOUSLY_QUOTED */, "I" /* FUNARI */, "J" /* MARKET_IF_TOUCHED */, "K" /* MARKET_WITH_LEFT_OVER_AS_LIMIT */, "L" /* PREVIOUS_FUND_VALUATION_POINT */, "M" /* NEXT_FUND_VALUATION_POINT */, "P" /* PEGGED */, "Q" /* COUNTER_ORDER_SELECTION */), f0.Con{1}},
		TimeInForce59:              f0.Fld{Opt, f0.ASCII, f0.String("0" /* DAY */, "1" /* GOOD_TILL_CANCEL */, "2" /* AT_THE_OPENING */, "3" /* IMMEDIATE_OR_CANCEL */, "4" /* FILL_OR_KILL */, "5" /* GOOD_TILL_CROSSING */, "6" /* GOOD_TILL_DATE */, "7" /* AT_THE_CLOSE */, "8" /* GOOD_THROUGH_CROSSING */, "9" /* AT_CROSSING */), f0.Con{1}},
		ExecInstValue1308:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ApplID1180:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ApplSeqNum1181:             f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		ApplLastSeqNum1350:         f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		ApplResendFlag1352:         f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		MarketReportID1394,         // STRING
		MarketReqID1393,            // STRING
		MarketUpdateAction1395,     // CHAR
		MarketID1301,               // EXCHANGE
		MarketSegmentID1300,        // STRING
		MarketSegmentDesc1396,      // STRING
		EncodedMktSegmDescLen1397,  // LENGTH
		EncodedMktSegmDesc1398,     // DATA
		ParentMktSegmID1325,        // STRING
		Currency15,                 // CURRENCY
		TransactTime60,             // UTCTIMESTAMP
		Text58,                     // STRING
		EncodedTextLen354,          // LENGTH
		EncodedText355,             // DATA
		ExpirationCycle827,         // INT
		MinTradeVol562,             // QTY
		MaxTradeVol1140,            // QTY
		MaxPriceVariation1143,      // FLOAT
		ImpliedMarketIndicator1144, // INT
		TradingCurrency1245,        // CURRENCY
		RoundLot561,                // QTY
		MultilegModel1377,          // INT
		MultilegPriceMethod1378,    // INT
		PriceType423,               // INT
		StartTickPriceRange1206,    // PRICE
		EndTickPriceRange1207,      // PRICE
		TickIncrement1208,          // PRICE
		TickRuleType1209,           // INT
		LotType1093,                // CHAR
		MinLotSize1231,             // QTY
		PriceLimitType1306,         // INT
		LowLimitPrice1148,          // PRICE
		HighLimitPrice1149,         // PRICE
		TradingReferencePrice1150,  // PRICE
		OrdType40,                  // CHAR
		TimeInForce59,              // CHAR
		ExecInstValue1308,          // CHAR
		ApplID1180,                 // STRING
		ApplSeqNum1181,             // SEQNUM
		ApplLastSeqNum1350,         // SEQNUM
		ApplResendFlag1352,         // BOOLEAN
	},
}

const Req, Opt = true, false

const (
	MarketReportID1394         = 1394 // STRING
	MarketReqID1393            = 1393 // STRING
	MarketUpdateAction1395     = 1395 // CHAR
	MarketID1301               = 1301 // EXCHANGE
	MarketSegmentID1300        = 1300 // STRING
	MarketSegmentDesc1396      = 1396 // STRING
	EncodedMktSegmDescLen1397  = 1397 // LENGTH
	EncodedMktSegmDesc1398     = 1398 // DATA
	ParentMktSegmID1325        = 1325 // STRING
	Currency15                 = 15   // CURRENCY
	TransactTime60             = 60   // UTCTIMESTAMP
	Text58                     = 58   // STRING
	EncodedTextLen354          = 354  // LENGTH
	EncodedText355             = 355  // DATA
	ExpirationCycle827         = 827  // INT
	MinTradeVol562             = 562  // QTY
	MaxTradeVol1140            = 1140 // QTY
	MaxPriceVariation1143      = 1143 // FLOAT
	ImpliedMarketIndicator1144 = 1144 // INT
	TradingCurrency1245        = 1245 // CURRENCY
	RoundLot561                = 561  // QTY
	MultilegModel1377          = 1377 // INT
	MultilegPriceMethod1378    = 1378 // INT
	PriceType423               = 423  // INT
	StartTickPriceRange1206    = 1206 // PRICE
	EndTickPriceRange1207      = 1207 // PRICE
	TickIncrement1208          = 1208 // PRICE
	TickRuleType1209           = 1209 // INT
	LotType1093                = 1093 // CHAR
	MinLotSize1231             = 1231 // QTY
	PriceLimitType1306         = 1306 // INT
	LowLimitPrice1148          = 1148 // PRICE
	HighLimitPrice1149         = 1149 // PRICE
	TradingReferencePrice1150  = 1150 // PRICE
	OrdType40                  = 40   // CHAR
	TimeInForce59              = 59   // CHAR
	ExecInstValue1308          = 1308 // CHAR
	ApplID1180                 = 1180 // STRING
	ApplSeqNum1181             = 1181 // SEQNUM
	ApplLastSeqNum1350         = 1350 // SEQNUM
	ApplResendFlag1352         = 1352 // BOOLEAN
)
