// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50derivativesecuritylistrequest is a format of the FIX.5.0 DerivativeSecurityListRequest message.
package fix50derivativesecuritylistrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50DerivativeSecurityListRequestMarshaler   = marshfix.Marshal{Tag: "FIX50", Format: FIX50DerivativeSecurityListRequest}
	FIX50DerivativeSecurityListRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50", Format: FIX50DerivativeSecurityListRequest}
)

// FIX50DerivativeSecurityListRequest is a FIX.5.0 format of the DerivativeSecurityListRequest message which maps the codecs into individual fields.
var FIX50DerivativeSecurityListRequest = f0.Format{
	Fields: map[int]f0.Codec{
		SecurityReqID320:                        f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityListRequestType559:              f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* SYMBOL */, 1 /* SECURITYTYPE_AND_OR_CFICODE */, 2 /* PRODUCT */, 3 /* TRADINGSESSIONID */, 4 /* ALL_SECURITIES */), f0.Var{1, 19}},
		SecuritySubType762:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Currency15:                              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		Text58:                                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:                       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                          f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		TradingSessionID336:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionSubID625:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SubscriptionRequestType263:              f0.Fld{Opt, f0.ASCII, f0.String("0" /* SNAPSHOT */, "1" /* SNAPSHOT_PLUS_UPDATES */, "2" /* DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST */), f0.Con{1}},
		UnderlyingSymbol311:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSymbolSfx312:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityID309:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityIDSource305:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingProduct462:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingCFICode463:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityType310:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecuritySubType763:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingMaturityMonthYear313:          f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		UnderlyingMaturityDate542:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingCouponPaymentDate241:          f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingIssueDate242:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingRepoCollateralSecurityType243: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseTerm244:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseRate245:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFactor246:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCreditRating256:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingInstrRegistry595:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCountryOfIssue592:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		UnderlyingStateOrProvinceOfIssue593:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingLocaleOfIssue594:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingRedemptionDate247:             f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingStrikePrice316:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingStrikeCurrency941:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		UnderlyingOptAttribute317:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		UnderlyingContractMultiplier436:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCouponRate435:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSecurityExchange308:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingIssuer306:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedUnderlyingIssuerLen362:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedUnderlyingIssuer363:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		UnderlyingSecurityDesc307:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedUnderlyingSecurityDescLen364:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedUnderlyingSecurityDesc365:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		UnderlyingCPProgram877:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCPRegType878:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCurrency318:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		UnderlyingQty879:                        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingPx810:                         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingDirtyPrice882:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingEndPrice883:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingStartValue884:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCurrentValue885:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingEndValue886:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingAllocationPercent972:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSettlementType975:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(2 /* T_PLUS_1 */, 4 /* T_PLUS_3 */, 5 /* T_PLUS_4 */), f0.Var{1, 19}},
		UnderlyingCashAmount973:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCashType974:                   f0.Fld{Opt, f0.ASCII, f0.String("FIXED" /* FIXED */, "DIFF" /* DIFF */), f0.Var{1, 65536}},
		UnderlyingUnitOfMeasure998:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingTimeUnit1000:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCapValue1038:                  f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSettlMethod1039:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingAdjustedQuantity1044:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFXRate1045:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFXRateCalc1046:                f0.Fld{Opt, f0.ASCII, f0.String("M" /* MULTIPLY */, "D" /* DIVIDE */), f0.Con{1}},
		UnderlyingSecurityAltID458:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityAltIDSource459:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingStipType888:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingStipValue889:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartyID1059:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartyIDSource1060:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		UndlyInstrumentPartyRole1061:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UndlyInstrumentPartySubID1063:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartySubIDType1064:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		SecurityReqID320,                        // STRING
		SecurityListRequestType559,              // INT
		SecuritySubType762,                      // STRING
		Currency15,                              // CURRENCY
		Text58,                                  // STRING
		EncodedTextLen354,                       // LENGTH
		EncodedText355,                          // DATA
		TradingSessionID336,                     // STRING
		TradingSessionSubID625,                  // STRING
		SubscriptionRequestType263,              // CHAR
		UnderlyingSymbol311,                     // STRING
		UnderlyingSymbolSfx312,                  // STRING
		UnderlyingSecurityID309,                 // STRING
		UnderlyingSecurityIDSource305,           // STRING
		UnderlyingProduct462,                    // INT
		UnderlyingCFICode463,                    // STRING
		UnderlyingSecurityType310,               // STRING
		UnderlyingSecuritySubType763,            // STRING
		UnderlyingMaturityMonthYear313,          // MONTHYEAR
		UnderlyingMaturityDate542,               // LOCALMKTDATE
		UnderlyingCouponPaymentDate241,          // LOCALMKTDATE
		UnderlyingIssueDate242,                  // LOCALMKTDATE
		UnderlyingRepoCollateralSecurityType243, // INT
		UnderlyingRepurchaseTerm244,             // INT
		UnderlyingRepurchaseRate245,             // PERCENTAGE
		UnderlyingFactor246,                     // FLOAT
		UnderlyingCreditRating256,               // STRING
		UnderlyingInstrRegistry595,              // STRING
		UnderlyingCountryOfIssue592,             // COUNTRY
		UnderlyingStateOrProvinceOfIssue593,     // STRING
		UnderlyingLocaleOfIssue594,              // STRING
		UnderlyingRedemptionDate247,             // LOCALMKTDATE
		UnderlyingStrikePrice316,                // PRICE
		UnderlyingStrikeCurrency941,             // CURRENCY
		UnderlyingOptAttribute317,               // CHAR
		UnderlyingContractMultiplier436,         // FLOAT
		UnderlyingCouponRate435,                 // PERCENTAGE
		UnderlyingSecurityExchange308,           // EXCHANGE
		UnderlyingIssuer306,                     // STRING
		EncodedUnderlyingIssuerLen362,           // LENGTH
		EncodedUnderlyingIssuer363,              // DATA
		UnderlyingSecurityDesc307,               // STRING
		EncodedUnderlyingSecurityDescLen364,     // LENGTH
		EncodedUnderlyingSecurityDesc365,        // DATA
		UnderlyingCPProgram877,                  // STRING
		UnderlyingCPRegType878,                  // STRING
		UnderlyingCurrency318,                   // CURRENCY
		UnderlyingQty879,                        // QTY
		UnderlyingPx810,                         // PRICE
		UnderlyingDirtyPrice882,                 // PRICE
		UnderlyingEndPrice883,                   // PRICE
		UnderlyingStartValue884,                 // AMT
		UnderlyingCurrentValue885,               // AMT
		UnderlyingEndValue886,                   // AMT
		UnderlyingAllocationPercent972,          // PERCENTAGE
		UnderlyingSettlementType975,             // INT
		UnderlyingCashAmount973,                 // AMT
		UnderlyingCashType974,                   // STRING
		UnderlyingUnitOfMeasure998,              // STRING
		UnderlyingTimeUnit1000,                  // STRING
		UnderlyingCapValue1038,                  // AMT
		UnderlyingSettlMethod1039,               // STRING
		UnderlyingAdjustedQuantity1044,          // QTY
		UnderlyingFXRate1045,                    // FLOAT
		UnderlyingFXRateCalc1046,                // CHAR
		UnderlyingSecurityAltID458,              // STRING
		UnderlyingSecurityAltIDSource459,        // STRING
		UnderlyingStipType888,                   // STRING
		UnderlyingStipValue889,                  // STRING
		UndlyInstrumentPartyID1059,              // STRING
		UndlyInstrumentPartyIDSource1060,        // CHAR
		UndlyInstrumentPartyRole1061,            // INT
		UndlyInstrumentPartySubID1063,           // STRING
		UndlyInstrumentPartySubIDType1064,       // INT
	},
}

const Req, Opt = true, false

const (
	SecurityReqID320                        = 320  // STRING
	SecurityListRequestType559              = 559  // INT
	SecuritySubType762                      = 762  // STRING
	Currency15                              = 15   // CURRENCY
	Text58                                  = 58   // STRING
	EncodedTextLen354                       = 354  // LENGTH
	EncodedText355                          = 355  // DATA
	TradingSessionID336                     = 336  // STRING
	TradingSessionSubID625                  = 625  // STRING
	SubscriptionRequestType263              = 263  // CHAR
	UnderlyingSymbol311                     = 311  // STRING
	UnderlyingSymbolSfx312                  = 312  // STRING
	UnderlyingSecurityID309                 = 309  // STRING
	UnderlyingSecurityIDSource305           = 305  // STRING
	UnderlyingProduct462                    = 462  // INT
	UnderlyingCFICode463                    = 463  // STRING
	UnderlyingSecurityType310               = 310  // STRING
	UnderlyingSecuritySubType763            = 763  // STRING
	UnderlyingMaturityMonthYear313          = 313  // MONTHYEAR
	UnderlyingMaturityDate542               = 542  // LOCALMKTDATE
	UnderlyingCouponPaymentDate241          = 241  // LOCALMKTDATE
	UnderlyingIssueDate242                  = 242  // LOCALMKTDATE
	UnderlyingRepoCollateralSecurityType243 = 243  // INT
	UnderlyingRepurchaseTerm244             = 244  // INT
	UnderlyingRepurchaseRate245             = 245  // PERCENTAGE
	UnderlyingFactor246                     = 246  // FLOAT
	UnderlyingCreditRating256               = 256  // STRING
	UnderlyingInstrRegistry595              = 595  // STRING
	UnderlyingCountryOfIssue592             = 592  // COUNTRY
	UnderlyingStateOrProvinceOfIssue593     = 593  // STRING
	UnderlyingLocaleOfIssue594              = 594  // STRING
	UnderlyingRedemptionDate247             = 247  // LOCALMKTDATE
	UnderlyingStrikePrice316                = 316  // PRICE
	UnderlyingStrikeCurrency941             = 941  // CURRENCY
	UnderlyingOptAttribute317               = 317  // CHAR
	UnderlyingContractMultiplier436         = 436  // FLOAT
	UnderlyingCouponRate435                 = 435  // PERCENTAGE
	UnderlyingSecurityExchange308           = 308  // EXCHANGE
	UnderlyingIssuer306                     = 306  // STRING
	EncodedUnderlyingIssuerLen362           = 362  // LENGTH
	EncodedUnderlyingIssuer363              = 363  // DATA
	UnderlyingSecurityDesc307               = 307  // STRING
	EncodedUnderlyingSecurityDescLen364     = 364  // LENGTH
	EncodedUnderlyingSecurityDesc365        = 365  // DATA
	UnderlyingCPProgram877                  = 877  // STRING
	UnderlyingCPRegType878                  = 878  // STRING
	UnderlyingCurrency318                   = 318  // CURRENCY
	UnderlyingQty879                        = 879  // QTY
	UnderlyingPx810                         = 810  // PRICE
	UnderlyingDirtyPrice882                 = 882  // PRICE
	UnderlyingEndPrice883                   = 883  // PRICE
	UnderlyingStartValue884                 = 884  // AMT
	UnderlyingCurrentValue885               = 885  // AMT
	UnderlyingEndValue886                   = 886  // AMT
	UnderlyingAllocationPercent972          = 972  // PERCENTAGE
	UnderlyingSettlementType975             = 975  // INT
	UnderlyingCashAmount973                 = 973  // AMT
	UnderlyingCashType974                   = 974  // STRING
	UnderlyingUnitOfMeasure998              = 998  // STRING
	UnderlyingTimeUnit1000                  = 1000 // STRING
	UnderlyingCapValue1038                  = 1038 // AMT
	UnderlyingSettlMethod1039               = 1039 // STRING
	UnderlyingAdjustedQuantity1044          = 1044 // QTY
	UnderlyingFXRate1045                    = 1045 // FLOAT
	UnderlyingFXRateCalc1046                = 1046 // CHAR
	UnderlyingSecurityAltID458              = 458  // STRING
	UnderlyingSecurityAltIDSource459        = 459  // STRING
	UnderlyingStipType888                   = 888  // STRING
	UnderlyingStipValue889                  = 889  // STRING
	UndlyInstrumentPartyID1059              = 1059 // STRING
	UndlyInstrumentPartyIDSource1060        = 1060 // CHAR
	UndlyInstrumentPartyRole1061            = 1061 // INT
	UndlyInstrumentPartySubID1063           = 1063 // STRING
	UndlyInstrumentPartySubIDType1064       = 1064 // INT
)
