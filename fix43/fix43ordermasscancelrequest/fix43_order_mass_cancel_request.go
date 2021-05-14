// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix43ordermasscancelrequest is a format of the FIX.4.3 OrderMassCancelRequest message.
package fix43ordermasscancelrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX43OrderMassCancelRequestMarshaler   = marshfix.Marshal{Tag: "FIX43", Format: FIX43OrderMassCancelRequest}
	FIX43OrderMassCancelRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX43", Format: FIX43OrderMassCancelRequest}
)

// FIX43OrderMassCancelRequest is a FIX.4.3 format of the OrderMassCancelRequest message which maps the codecs into individual fields.
var FIX43OrderMassCancelRequest = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:                            f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.3"), f0.Con{7}},
		MsgType35:                               f0.Fld{Req, f0.ASCII, f0.String("m" /* LIST_STRIKE_PRICE */, "l" /* BID_RESPONSE */, "k" /* BID_REQUEST */, "j" /* BUSINESS_MESSAGE_REJECT */, "i" /* MASS_QUOTE */, "h" /* TRADING_SESSION_STATUS */, "q" /* ORDER_MASS_CANCEL_REQUEST */, "AE" /* TRADE_CAPTURE_REPORT */, "o" /* REGISTRATION_INSTRUCTIONS */, "z" /* DERIVATIVE_SECURITY_LIST_REQUEST */, "AI" /* QUOTE_STATUS_REPORT */, "AH" /* RFQ_REQUEST */, "AG" /* QUOTE_REQUEST_REJECT */, "AF" /* ORDER_MASS_STATUS_REQUEST */, "n" /* XML_MESSAGE */, "AD" /* TRADE_CAPTURE_REPORT_REQUEST */, "g" /* TRADING_SESSION_STATUS_REQUEST */, "AC" /* MULTILEG_ORDER_CANCEL_REPLACE */, "AA" /* DERIVATIVE_SECURITY_LIST */, "r" /* ORDER_MASS_CANCEL_REPORT */, "y" /* SECURITY_LIST */, "x" /* SECURITY_LIST_REQUEST */, "w" /* SECURITY_TYPES */, "v" /* SECURITY_TYPE_REQUEST */, "u" /* CROSS_ORDER_CANCEL_REQUEST */, "t" /* CROSS_ORDER_CANCEL_REPLACE_REQUEST */, "s" /* NEW_ORDER_s */, "AB" /* NEW_ORDER_AB */, "9" /* ORDER_CANCEL_REJECT */, "7" /* ADVERTISEMENT */, "M" /* LIST_STATUS_REQUEST */, "E" /* ORDER_LIST */, "D" /* ORDER_SINGLE */, "C" /* EMAIL */, "B" /* NEWS */, "A" /* LOGON */, "N" /* LIST_STATUS */, "f" /* SECURITY_STATUS */, "P" /* ALLOCATION_ACK */, "8" /* EXECUTION_REPORT */, "0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "K" /* LIST_CANCEL_REQUEST */, "e" /* SECURITY_STATUS_REQUEST */, "d" /* SECURITY_DEFINITION */, "c" /* SECURITY_DEFINITION_REQUEST */, "b" /* MASS_QUOTE_ACKNOWLEDGEMENT */, "a" /* QUOTE_STATUS_REQUEST */, "Z" /* QUOTE_CANCEL */, "Y" /* MARKET_DATA_REQUEST_REJECT */, "F" /* ORDER_CANCEL_REQUEST */, "J" /* ALLOCATION */, "p" /* REGISTRATION_INSTRUCTIONS_RESPONSE */, "L" /* LIST_EXECUTE */, "X" /* MARKET_DATA_INCREMENTAL_REFRESH */, "W" /* MARKET_DATA_SNAPSHOT_FULL_REFRESH */, "V" /* MARKET_DATA_REQUEST */, "T" /* SETTLEMENT_INSTRUCTIONS */, "S" /* QUOTE */, "R" /* QUOTE_REQUEST */, "Q" /* DONT_KNOW_TRADE */, "H" /* ORDER_STATUS_REQUEST */), f0.Var{1, 2}},
		SenderCompID49:                          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:                          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:                         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:                            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:                             f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SenderSubID50:                           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SenderLocationID142:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetSubID57:                           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetLocationID143:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfSubID116:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfLocationID144:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToSubID129:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToLocationID145:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PossDupFlag43:                           f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:                            f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:                           f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:                      f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		XmlDataLen212:                           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		XmlData213:                              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MessageEncoding347:                      f0.Fld{Opt, f0.ASCII, f0.String("UTF-8" /* UTF_8 */, "ISO-2022-JP" /* ISO_2022_JP */, "EUC-JP" /* EUC_JP */, "SHIFT_JIS" /* SHIFT_JIS */), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369:               f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		OnBehalfOfSendingTime370:                f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopCompID628:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:                       f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:                             f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		ClOrdID11:                               f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryClOrdID526:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MassCancelRequestType530:                f0.Fld{Req, f0.ASCII, f0.String("1" /* CANCEL_ORDERS_FOR_A_SECURITY */, "7" /* CANCEL_ALL_ORDERS */, "6" /* CANCEL_ORDERS_FOR_A_TRADING_SESSION */, "5" /* CANCEL_ORDERS_FOR_A_SECURITYTYPE */, "4" /* CANCEL_ORDERS_FOR_A_CFICODE */, "2" /* CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY */, "3" /* CANCEL_ORDERS_FOR_A_PRODUCT */), f0.Con{1}},
		TradingSessionID336:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionSubID625:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Side54:                                  f0.Fld{Opt, f0.ASCII, f0.String("6" /* SELL_SHORT_EXEMPT */, "B" /* AS_DEFINED */, "C" /* OPPOSITE */, "8" /* CROSS */, "9" /* CROSS_SHORT */, "1" /* BUY */, "2" /* SELL */, "3" /* BUY_MINUS */, "4" /* SELL_PLUS */, "A" /* CROSS_SHORT_EXEMPT */, "5" /* SELL_SHORT */, "7" /* UNDISCLOSED */), f0.Con{1}},
		TransactTime60:                          f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Text58:                                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:                       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                          f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		Symbol55:                                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SymbolSfx65:                             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityID48:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityIDSource22:                      f0.Fld{Opt, f0.ASCII, f0.String("E" /* SICOVAM */, "2" /* SEDOL */, "1" /* CUSIP */, "3" /* QUIK */, "F" /* BELGIAN */, "D" /* VALOREN */, "C" /* DUTCH */, "B" /* WERTPAPIER */, "A" /* BLOOMBERG_SYMBOL */, "9" /* CONSOLIDATED_TAPE_ASSOCIATION */, "8" /* EXCHANGE_SYMBOL */, "7" /* ISO_COUNTRY_CODE */, "6" /* ISO_CURRENCY_CODE */, "5" /* RIC_CODE */, "4" /* ISIN_NUMBER */, "G" /* COMMON */), f0.Var{1, 65536}},
		Product460:                              f0.Fld{Opt, f0.ASCII, f0.IntDefault(8 /* LOAN */, 12 /* OTHER */, 11 /* MUNICIPAL */, 1 /* AGENCY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 2 /* COMMODITY */, 6 /* GOVERNMENT */, 10 /* MORTGAGE */, 7 /* INDEX */, 9 /* MONEYMARKET */, 5 /* EQUITY */), f0.Var{1, 19}},
		CFICode461:                              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityType167:                         f0.Fld{Opt, f0.ASCII, f0.String("CP" /* COMMERCIAL_PAPER */, "VRDN" /* VARIABLE_RATE_DEMAND_NOTE */, "PZFJ" /* PLAZOS_FIJOS */, "PN" /* PROMISSORY_NOTE */, "ONITE" /* OVERNIGHT */, "MTN" /* MEDIUM_TERM_NOTES */, "TECP" /* TAX_EXEMPT_COMMERCIAL_PAPER */, "AMENDED" /* AMENDED_RESTATED */, "BRIDGE" /* BRIDGE_LOAN */, "LOFC" /* LETTER_OF_CREDIT */, "SWING" /* SWING_LINE_FACILITY */, "DINP" /* DEBTOR_IN_POSSESSION */, "DEFLTED" /* DEFAULTED */, "WITHDRN" /* WITHDRAWN */, "LQN" /* LIQUIDITY_NOTE */, "MATURED" /* MATURED */, "DN" /* DEPOSIT_NOTES */, "RETIRED" /* RETIRED */, "BA" /* BANKERS_ACCEPTANCE */, "BN" /* BANK_NOTES */, "BOX" /* BILL_OF_EXCHANGES */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CL" /* CALL_LOANS */, "REPLACD" /* REPLACED */, "MT" /* MANDATORY_TENDER */, "RVLVTRM" /* REVOLVER_TERM_LOAN */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "STN" /* SHORT_TERM_LOAN_NOTE */, "MPT" /* MISCELLANEOUS_PASS_THROUGH */, "TBA" /* TO_BE_ANNOUNCED */, "AN" /* OTHER_ANTICIPATION_NOTES_BAN_GAN_ETC */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "COFP" /* CERTIFICATE_OF_PARTICIPATION */, "MBS" /* MORTGAGE_BACKED_SECURITIES */, "REV" /* REVENUE_BONDS */, "SPCLA" /* SPECIAL_ASSESSMENT */, "SPCLO" /* SPECIAL_OBLIGATION */, "SPCLT" /* SPECIAL_TAX */, "TAN" /* TAX_ANTICIPATION_NOTE */, "TAXA" /* TAX_ALLOCATION */, "COFO" /* CERTIFICATE_OF_OBLIGATION */, "TD" /* TIME_DEPOSIT */, "GO" /* GENERAL_OBLIGATION_BONDS */, "?" /* WILDCARD_ENTRY */, "WAR" /* WARRANT */, "MF" /* MUTUAL_FUND */, "MLEG" /* MULTI_LEG_INSTRUMENT */, "TRAN" /* TAX_REVENUE_ANTICIPATION_NOTE */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "RP" /* REPURCHASE_AGREEMENT */, "NONE" /* NO_SECURITY_TYPE */, "XCN" /* EXTENDED_COMM_NOTE */, "POOL" /* AGENCY_POOLS */, "ABS" /* ASSET_BACKED_SECURITIES */, "CMBS" /* CORP_MORTGAGE_BACKED_SECURITIES */, "CMO" /* COLLATERALIZED_MORTGAGE_OBLIGATION */, "IET" /* IOETTE_MORTGAGE */, "RVRP" /* REVERSE_REPURCHASE_AGREEMENT */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "RAN" /* REVENUE_ANTICIPATION_NOTE */, "RVLV" /* REVOLVER_LOAN */, "FAC" /* FEDERAL_AGENCY_COUPON */, "FADN" /* FEDERAL_AGENCY_DISCOUNT_NOTE */, "PEF" /* PRIVATE_EXPORT_FUNDING */, "CORP" /* CORPORATE_BOND */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CB" /* CONVERTIBLE_BOND */, "DUAL" /* DUAL_CURRENCY */, "XLINKD" /* INDEXED_LINKED */, "YANK" /* YANKEE_CORPORATE_BOND */, "CS" /* COMMON_STOCK */, "PS" /* PREFERRED_STOCK */, "BRADY" /* BRADY_BOND */, "TBOND" /* US_TREASURY_BOND */, "TINT" /* INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE */, "TIPS" /* TREASURY_INFLATION_PROTECTED_SECURITIES */, "TCAL" /* PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE */, "TPRN" /* PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE */, "UST" /* US_TREASURY_NOTE_BOND */, "USTB" /* US_TREASURY_BILL */, "TERM" /* TERM_LOAN */, "STRUCT" /* STRUCTURED_NOTES */), f0.Var{1, 65536}},
		MaturityMonthYear200:                    f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		MaturityDate541:                         f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		CouponPaymentDate224:                    f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		IssueDate225:                            f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		RepoCollateralSecurityType239:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseTerm226:                       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseRate227:                       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Factor228:                               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CreditRating255:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrRegistry543:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CountryOfIssue470:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		StateOrProvinceOfIssue471:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LocaleOfIssue472:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RedemptionDate240:                       f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		StrikePrice202:                          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OptAttribute206:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContractMultiplier231:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CouponRate223:                           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SecurityExchange207:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Issuer106:                               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedIssuerLen348:                     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedIssuer349:                        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityDesc107:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedSecurityDescLen350:               f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedSecurityDesc351:                  f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityAltID455:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityAltIDSource456:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSymbol311:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSymbolSfx312:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityID309:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityIDSource305:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingProduct462:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingCFICode463:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityType310:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingMaturityMonthYear313:          f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		UnderlyingMaturityDate542:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingCouponPaymentDate241:          f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		UnderlyingIssueDate242:                  f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		UnderlyingRepoCollateralSecurityType243: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseTerm244:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseRate245:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFactor246:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCreditRating256:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingInstrRegistry595:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCountryOfIssue592:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		UnderlyingStateOrProvinceOfIssue593:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingLocaleOfIssue594:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingRedemptionDate247:             f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		UnderlyingStrikePrice316:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
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
		UnderlyingSecurityAltID458:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityAltIDSource459:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SignatureLength93:                       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:                             f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,                            // STRING
		BodyLength9,                             // LENGTH
		MsgType35,                               // STRING
		SenderCompID49,                          // STRING
		TargetCompID56,                          // STRING
		OnBehalfOfCompID115,                     // STRING
		DeliverToCompID128,                      // STRING
		SecureDataLen90,                         // LENGTH
		SecureData91,                            // DATA
		MsgSeqNum34,                             // SEQNUM
		SenderSubID50,                           // STRING
		SenderLocationID142,                     // STRING
		TargetSubID57,                           // STRING
		TargetLocationID143,                     // STRING
		OnBehalfOfSubID116,                      // STRING
		OnBehalfOfLocationID144,                 // STRING
		DeliverToSubID129,                       // STRING
		DeliverToLocationID145,                  // STRING
		PossDupFlag43,                           // BOOLEAN
		PossResend97,                            // BOOLEAN
		SendingTime52,                           // UTCTIMESTAMP
		OrigSendingTime122,                      // UTCTIMESTAMP
		XmlDataLen212,                           // LENGTH
		XmlData213,                              // DATA
		MessageEncoding347,                      // STRING
		LastMsgSeqNumProcessed369,               // SEQNUM
		OnBehalfOfSendingTime370,                // UTCTIMESTAMP
		HopCompID628,                            // STRING
		HopSendingTime629,                       // UTCTIMESTAMP
		HopRefID630,                             // SEQNUM
		ClOrdID11,                               // STRING
		SecondaryClOrdID526,                     // STRING
		MassCancelRequestType530,                // CHAR
		TradingSessionID336,                     // STRING
		TradingSessionSubID625,                  // STRING
		Side54,                                  // CHAR
		TransactTime60,                          // UTCTIMESTAMP
		Text58,                                  // STRING
		EncodedTextLen354,                       // LENGTH
		EncodedText355,                          // DATA
		Symbol55,                                // STRING
		SymbolSfx65,                             // STRING
		SecurityID48,                            // STRING
		SecurityIDSource22,                      // STRING
		Product460,                              // INT
		CFICode461,                              // STRING
		SecurityType167,                         // STRING
		MaturityMonthYear200,                    // MONTHYEAR
		MaturityDate541,                         // LOCALMKTDATE
		CouponPaymentDate224,                    // UTCDATE
		IssueDate225,                            // UTCDATE
		RepoCollateralSecurityType239,           // INT
		RepurchaseTerm226,                       // INT
		RepurchaseRate227,                       // PERCENTAGE
		Factor228,                               // FLOAT
		CreditRating255,                         // STRING
		InstrRegistry543,                        // STRING
		CountryOfIssue470,                       // COUNTRY
		StateOrProvinceOfIssue471,               // STRING
		LocaleOfIssue472,                        // STRING
		RedemptionDate240,                       // UTCDATE
		StrikePrice202,                          // PRICE
		OptAttribute206,                         // CHAR
		ContractMultiplier231,                   // FLOAT
		CouponRate223,                           // PERCENTAGE
		SecurityExchange207,                     // EXCHANGE
		Issuer106,                               // STRING
		EncodedIssuerLen348,                     // LENGTH
		EncodedIssuer349,                        // DATA
		SecurityDesc107,                         // STRING
		EncodedSecurityDescLen350,               // LENGTH
		EncodedSecurityDesc351,                  // DATA
		SecurityAltID455,                        // STRING
		SecurityAltIDSource456,                  // STRING
		UnderlyingSymbol311,                     // STRING
		UnderlyingSymbolSfx312,                  // STRING
		UnderlyingSecurityID309,                 // STRING
		UnderlyingSecurityIDSource305,           // STRING
		UnderlyingProduct462,                    // INT
		UnderlyingCFICode463,                    // STRING
		UnderlyingSecurityType310,               // STRING
		UnderlyingMaturityMonthYear313,          // MONTHYEAR
		UnderlyingMaturityDate542,               // LOCALMKTDATE
		UnderlyingCouponPaymentDate241,          // UTCDATE
		UnderlyingIssueDate242,                  // UTCDATE
		UnderlyingRepoCollateralSecurityType243, // INT
		UnderlyingRepurchaseTerm244,             // INT
		UnderlyingRepurchaseRate245,             // PERCENTAGE
		UnderlyingFactor246,                     // FLOAT
		UnderlyingCreditRating256,               // STRING
		UnderlyingInstrRegistry595,              // STRING
		UnderlyingCountryOfIssue592,             // COUNTRY
		UnderlyingStateOrProvinceOfIssue593,     // STRING
		UnderlyingLocaleOfIssue594,              // STRING
		UnderlyingRedemptionDate247,             // UTCDATE
		UnderlyingStrikePrice316,                // PRICE
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
		UnderlyingSecurityAltID458,              // STRING
		UnderlyingSecurityAltIDSource459,        // STRING
		SignatureLength93,                       // LENGTH
		Signature89,                             // DATA
		CheckSum10,                              // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8                            = 8   // STRING
	BodyLength9                             = 9   // LENGTH
	MsgType35                               = 35  // STRING
	SenderCompID49                          = 49  // STRING
	TargetCompID56                          = 56  // STRING
	OnBehalfOfCompID115                     = 115 // STRING
	DeliverToCompID128                      = 128 // STRING
	SecureDataLen90                         = 90  // LENGTH
	SecureData91                            = 91  // DATA
	MsgSeqNum34                             = 34  // SEQNUM
	SenderSubID50                           = 50  // STRING
	SenderLocationID142                     = 142 // STRING
	TargetSubID57                           = 57  // STRING
	TargetLocationID143                     = 143 // STRING
	OnBehalfOfSubID116                      = 116 // STRING
	OnBehalfOfLocationID144                 = 144 // STRING
	DeliverToSubID129                       = 129 // STRING
	DeliverToLocationID145                  = 145 // STRING
	PossDupFlag43                           = 43  // BOOLEAN
	PossResend97                            = 97  // BOOLEAN
	SendingTime52                           = 52  // UTCTIMESTAMP
	OrigSendingTime122                      = 122 // UTCTIMESTAMP
	XmlDataLen212                           = 212 // LENGTH
	XmlData213                              = 213 // DATA
	MessageEncoding347                      = 347 // STRING
	LastMsgSeqNumProcessed369               = 369 // SEQNUM
	OnBehalfOfSendingTime370                = 370 // UTCTIMESTAMP
	HopCompID628                            = 628 // STRING
	HopSendingTime629                       = 629 // UTCTIMESTAMP
	HopRefID630                             = 630 // SEQNUM
	ClOrdID11                               = 11  // STRING
	SecondaryClOrdID526                     = 526 // STRING
	MassCancelRequestType530                = 530 // CHAR
	TradingSessionID336                     = 336 // STRING
	TradingSessionSubID625                  = 625 // STRING
	Side54                                  = 54  // CHAR
	TransactTime60                          = 60  // UTCTIMESTAMP
	Text58                                  = 58  // STRING
	EncodedTextLen354                       = 354 // LENGTH
	EncodedText355                          = 355 // DATA
	Symbol55                                = 55  // STRING
	SymbolSfx65                             = 65  // STRING
	SecurityID48                            = 48  // STRING
	SecurityIDSource22                      = 22  // STRING
	Product460                              = 460 // INT
	CFICode461                              = 461 // STRING
	SecurityType167                         = 167 // STRING
	MaturityMonthYear200                    = 200 // MONTHYEAR
	MaturityDate541                         = 541 // LOCALMKTDATE
	CouponPaymentDate224                    = 224 // UTCDATE
	IssueDate225                            = 225 // UTCDATE
	RepoCollateralSecurityType239           = 239 // INT
	RepurchaseTerm226                       = 226 // INT
	RepurchaseRate227                       = 227 // PERCENTAGE
	Factor228                               = 228 // FLOAT
	CreditRating255                         = 255 // STRING
	InstrRegistry543                        = 543 // STRING
	CountryOfIssue470                       = 470 // COUNTRY
	StateOrProvinceOfIssue471               = 471 // STRING
	LocaleOfIssue472                        = 472 // STRING
	RedemptionDate240                       = 240 // UTCDATE
	StrikePrice202                          = 202 // PRICE
	OptAttribute206                         = 206 // CHAR
	ContractMultiplier231                   = 231 // FLOAT
	CouponRate223                           = 223 // PERCENTAGE
	SecurityExchange207                     = 207 // EXCHANGE
	Issuer106                               = 106 // STRING
	EncodedIssuerLen348                     = 348 // LENGTH
	EncodedIssuer349                        = 349 // DATA
	SecurityDesc107                         = 107 // STRING
	EncodedSecurityDescLen350               = 350 // LENGTH
	EncodedSecurityDesc351                  = 351 // DATA
	SecurityAltID455                        = 455 // STRING
	SecurityAltIDSource456                  = 456 // STRING
	UnderlyingSymbol311                     = 311 // STRING
	UnderlyingSymbolSfx312                  = 312 // STRING
	UnderlyingSecurityID309                 = 309 // STRING
	UnderlyingSecurityIDSource305           = 305 // STRING
	UnderlyingProduct462                    = 462 // INT
	UnderlyingCFICode463                    = 463 // STRING
	UnderlyingSecurityType310               = 310 // STRING
	UnderlyingMaturityMonthYear313          = 313 // MONTHYEAR
	UnderlyingMaturityDate542               = 542 // LOCALMKTDATE
	UnderlyingCouponPaymentDate241          = 241 // UTCDATE
	UnderlyingIssueDate242                  = 242 // UTCDATE
	UnderlyingRepoCollateralSecurityType243 = 243 // INT
	UnderlyingRepurchaseTerm244             = 244 // INT
	UnderlyingRepurchaseRate245             = 245 // PERCENTAGE
	UnderlyingFactor246                     = 246 // FLOAT
	UnderlyingCreditRating256               = 256 // STRING
	UnderlyingInstrRegistry595              = 595 // STRING
	UnderlyingCountryOfIssue592             = 592 // COUNTRY
	UnderlyingStateOrProvinceOfIssue593     = 593 // STRING
	UnderlyingLocaleOfIssue594              = 594 // STRING
	UnderlyingRedemptionDate247             = 247 // UTCDATE
	UnderlyingStrikePrice316                = 316 // PRICE
	UnderlyingOptAttribute317               = 317 // CHAR
	UnderlyingContractMultiplier436         = 436 // FLOAT
	UnderlyingCouponRate435                 = 435 // PERCENTAGE
	UnderlyingSecurityExchange308           = 308 // EXCHANGE
	UnderlyingIssuer306                     = 306 // STRING
	EncodedUnderlyingIssuerLen362           = 362 // LENGTH
	EncodedUnderlyingIssuer363              = 363 // DATA
	UnderlyingSecurityDesc307               = 307 // STRING
	EncodedUnderlyingSecurityDescLen364     = 364 // LENGTH
	EncodedUnderlyingSecurityDesc365        = 365 // DATA
	UnderlyingSecurityAltID458              = 458 // STRING
	UnderlyingSecurityAltIDSource459        = 459 // STRING
	SignatureLength93                       = 93  // LENGTH
	Signature89                             = 89  // DATA
	CheckSum10                              = 10  // STRING
)
