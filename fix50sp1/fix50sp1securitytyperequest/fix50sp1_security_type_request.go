// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1securitytyperequest is a format of the FIX.5.0 servicepack 1 SecurityTypeRequest message.
package fix50sp1securitytyperequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1SecurityTypeRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1SecurityTypeRequest}
	FIX50SP1SecurityTypeRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1SecurityTypeRequest}
)

// FIX50SP1SecurityTypeRequest is a FIX.5.0 servicepack 1 format of the SecurityTypeRequest message which maps the codecs into individual fields.
var FIX50SP1SecurityTypeRequest = f0.Format{
	Fields: map[int]f0.Codec{
		SecurityReqID320:       f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Text58:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:      f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:         f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		TradingSessionID336:    f0.Fld{Opt, f0.ASCII, f0.String("1" /* DAY */, "2" /* HALFDAY */, "3" /* MORNING */, "4" /* AFTERNOON */, "5" /* EVENING */, "6" /* AFTER_HOURS */), f0.Var{1, 65536}},
		TradingSessionSubID625: f0.Fld{Opt, f0.ASCII, f0.String("1" /* PRE_TRADING */, "2" /* OPENING_OR_OPENING_AUCTION */, "3" /* 3 */, "4" /* CLOSING_OR_CLOSING_AUCTION */, "5" /* POST_TRADING */, "6" /* INTRADAY_AUCTION */, "7" /* QUIESCENT */), f0.Var{1, 65536}},
		Product460:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		SecurityType167:        f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSET_BACKED_SECURITIES */, "AMENDED" /* AMENDED_RESTATED */, "AN" /* OTHER_ANTICIPATION_NOTES */, "BA" /* BANKERS_ACCEPTANCE */, "BN" /* BANK_NOTES */, "BOX" /* BILL_OF_EXCHANGES */, "BRADY" /* BRADY_BOND */, "BRIDGE" /* BRIDGE_LOAN */, "BUYSELL" /* BUY_SELLBACK */, "CB" /* CONVERTIBLE_BOND */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CL" /* CALL_LOANS */, "CMBS" /* CORP_MORTGAGE_BACKED_SECURITIES */, "CMO" /* COLLATERALIZED_MORTGAGE_OBLIGATION */, "COFO" /* CERTIFICATE_OF_OBLIGATION */, "COFP" /* CERTIFICATE_OF_PARTICIPATION */, "CORP" /* CORPORATE_BOND */, "CP" /* COMMERCIAL_PAPER */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CS" /* COMMON_STOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTOR_IN_POSSESSION */, "DN" /* DEPOSIT_NOTES */, "DUAL" /* DUAL_CURRENCY */, "EUCD" /* EURO_CERTIFICATE_OF_DEPOSIT */, "EUCORP" /* EURO_CORPORATE_BOND */, "EUCP" /* EURO_COMMERCIAL_PAPER */, "EUSOV" /* EURO_SOVEREIGNS */, "EUSUPRA" /* EURO_SUPRANATIONAL_COUPONS */, "FAC" /* FEDERAL_AGENCY_COUPON */, "FADN" /* FEDERAL_AGENCY_DISCOUNT_NOTE */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERAL_OBLIGATION_BONDS */, "IET" /* IOETTE_MORTGAGE */, "LOFC" /* LETTER_OF_CREDIT */, "LQN" /* LIQUIDITY_NOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGE_BACKED_SECURITIES */, "MF" /* MUTUAL_FUND */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "MLEG" /* MULTILEG_INSTRUMENT */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "MPT" /* MISCELLANEOUS_PASS_THROUGH */, "MT" /* MANDATORY_TENDER */, "MTN" /* MEDIUM_TERM_NOTES */, "NONE" /* NO_SECURITY_TYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATE_EXPORT_FUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORY_NOTE */, "PS" /* PREFERRED_STOCK */, "PZFJ" /* PLAZOS_FIJOS */, "RAN" /* REVENUE_ANTICIPATION_NOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUE_BONDS */, "RVLV" /* REVOLVER_LOAN */, "RVLVTRM" /* REVOLVER_TERM_LOAN */, "SECLOAN" /* SECURITIES_LOAN */, "SECPLEDGE" /* SECURITIES_PLEDGE */, "SPCLA" /* SPECIAL_ASSESSMENT */, "SPCLO" /* SPECIAL_OBLIGATION */, "SPCLT" /* SPECIAL_TAX */, "STN" /* SHORT_TERM_LOAN_NOTE */, "STRUCT" /* STRUCTURED_NOTES */, "SUPRA" /* USD_SUPRANATIONAL_COUPONS */, "SWING" /* SWING_LINE_FACILITY */, "TAN" /* TAX_ANTICIPATION_NOTE */, "TAXA" /* TAX_ALLOCATION */, "TBA" /* TO_BE_ANNOUNCED */, "TBILL" /* US_TREASURY_BILL_TBILL */, "TBOND" /* US_TREASURY_BOND */, "TCAL" /* PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE */, "TD" /* TIME_DEPOSIT */, "TECP" /* TAX_EXEMPT_COMMERCIAL_PAPER */, "TERM" /* TERM_LOAN */, "TINT" /* INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE */, "TIPS" /* TREASURY_INFLATION_PROTECTED_SECURITIES */, "TNOTE" /* US_TREASURY_NOTE_TNOTE */, "TPRN" /* PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE */, "TRAN" /* TAX_REVENUE_ANTICIPATION_NOTE */, "UST" /* US_TREASURY_NOTE_UST */, "USTB" /* US_TREASURY_BILL_USTB */, "VRDN" /* VARIABLE_RATE_DEMAND_NOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "?" /* WILDCARD_ENTRY_FOR_USE_ON_SECURITY_DEFINITION_REQUEST */, "XCN" /* EXTENDED_COMM_NOTE */, "XLINKD" /* INDEXED_LINKED */, "YANK" /* YANKEE_CORPORATE_BOND */, "YCD" /* YANKEE_CERTIFICATE_OF_DEPOSIT */, "OOP" /* OPTIONS_ON_PHYSICAL */, "OOF" /* OPTIONS_ON_FUTURES */, "CASH" /* CASH */, "OOC" /* OPTIONS_ON_COMBO */, "IRS" /* INTEREST_RATE_SWAP */, "BDN" /* BANK_DEPOSITORY_NOTE */, "CAMM" /* CANADIAN_MONEY_MARKETS */, "CAN" /* CANADIAN_TREASURY_NOTES */, "CTB" /* CANADIAN_TREASURY_BILLS */, "CDS" /* CREDIT_DEFAULT_SWAP */, "CMB" /* CANADIAN_MORTGAGE_BONDS */, "EUFRN" /* EURO_CORPORATE_FLOATING_RATE_NOTES */, "FRN" /* US_CORPORATE_FLOATING_RATE_NOTES */, "PROV" /* CANADIAN_PROVINCIAL_BONDS */, "SLQN" /* SECURED_LIQUIDITY_NOTE */, "TB" /* TREASURY_BILL */, "TLQN" /* TERM_LIQUIDITY_NOTE */, "TMCP" /* TAXABLE_MUNICIPAL_CP */), f0.Var{1, 65536}},
		SecuritySubType762:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketID1301:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketSegmentID1300:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		SecurityReqID320,       // STRING
		Text58,                 // STRING
		EncodedTextLen354,      // LENGTH
		EncodedText355,         // DATA
		TradingSessionID336,    // STRING
		TradingSessionSubID625, // STRING
		Product460,             // INT
		SecurityType167,        // STRING
		SecuritySubType762,     // STRING
		MarketID1301,           // EXCHANGE
		MarketSegmentID1300,    // STRING
	},
}

const Req, Opt = true, false

const (
	SecurityReqID320       = 320  // STRING
	Text58                 = 58   // STRING
	EncodedTextLen354      = 354  // LENGTH
	EncodedText355         = 355  // DATA
	TradingSessionID336    = 336  // STRING
	TradingSessionSubID625 = 625  // STRING
	Product460             = 460  // INT
	SecurityType167        = 167  // STRING
	SecuritySubType762     = 762  // STRING
	MarketID1301           = 1301 // EXCHANGE
	MarketSegmentID1300    = 1300 // STRING
)
