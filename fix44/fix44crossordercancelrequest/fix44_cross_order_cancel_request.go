// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix44crossordercancelrequest is a format of the FIX.4.4 CrossOrderCancelRequest message.
package fix44crossordercancelrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX44CrossOrderCancelRequestMarshaler   = marshfix.Marshal{Tag: "FIX44", Format: FIX44CrossOrderCancelRequest}
	FIX44CrossOrderCancelRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX44", Format: FIX44CrossOrderCancelRequest}
)

// FIX44CrossOrderCancelRequest is a FIX.4.4 format of the CrossOrderCancelRequest message which maps the codecs into individual fields.
var FIX44CrossOrderCancelRequest = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:                  f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:                     f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:               f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:                  f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:                   f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SenderSubID50:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SenderLocationID142:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetSubID57:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetLocationID143:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfSubID116:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfLocationID144:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToSubID129:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToLocationID145:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PossDupFlag43:                 f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:                  f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:                 f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:            f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		XmlDataLen212:                 f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		XmlData213:                    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MessageEncoding347:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369:     f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		HopCompID628:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:             f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:                   f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		OrderID37:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CrossID548:                    f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OrigCrossID551:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CrossType549:                  f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* CROSSAON */, 2 /* CROSSIOC */, 3 /* CROSSONESIDE */, 4 /* CROSSSAMEPRICE */), f0.Var{1, 19}},
		CrossPrioritization550:        f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* NONE */, 1 /* BUY_SIDE_PRIORITIZED */, 2 /* SELL_SIDE_PRIORITIZED */), f0.Var{1, 19}},
		TransactTime60:                f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Symbol55:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SymbolSfx65:                   f0.Fld{Opt, f0.ASCII, f0.String("CD" /* EUCPLUMPSUMINTEREST */, "WI" /* WHENISSUED */), f0.Var{1, 65536}},
		SecurityID48:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityIDSource22:            f0.Fld{Opt, f0.ASCII, f0.String("1" /* CUSIP */, "2" /* SEDOL */, "3" /* QUIK */, "4" /* ISIN */, "5" /* RIC */, "6" /* ISOCURR */, "7" /* ISOCOUNTRY */, "8" /* EXCHSYMB */, "9" /* CTA */, "A" /* BLMBRG */, "B" /* WERTPAPIER */, "C" /* DUTCH */, "D" /* VALOREN */, "E" /* SICOVAM */, "F" /* BELGIAN */, "G" /* COMMON */, "H" /* CLEARINGHOUSE */, "I" /* FPML */, "J" /* OPTIONPRICEREPORTINGAUTHORITY */), f0.Var{1, 65536}},
		Product460:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		CFICode461:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityType167:               f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSETBACKEDSECURITIES */, "AMENDED" /* AMENDEDRESTATED */, "AN" /* OTHERANTICIPATIONNOTESBANGANETC */, "BA" /* BANKERSACCEPTANCE */, "BN" /* BANKNOTES */, "BOX" /* BILLOFEXCHANGES */, "BRADY" /* BRADYBOND */, "BRIDGE" /* BRIDGELOAN */, "BUYSELL" /* BUYSELLBACK */, "CB" /* CONVERTIBLEBOND */, "CD" /* CERTIFICATEOFDEPOSIT */, "CL" /* CALLLOANS */, "CMBS" /* CORPMORTGAGEBACKEDSECURITIES */, "CMO" /* COLLATERALIZEDMORTGAGEOBLIGATION */, "COFO" /* CERTIFICATEOFOBLIGATION */, "COFP" /* CERTIFICATEOFPARTICIPATION */, "CORP" /* CORPORATEBOND */, "CP" /* COMMERCIALPAPER */, "CPP" /* CORPORATEPRIVATEPLACEMENT */, "CS" /* COMMONSTOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTORINPOSSESSION */, "DN" /* DEPOSITNOTES */, "DUAL" /* DUALCURRENCY */, "EUCD" /* EUROCERTIFICATEOFDEPOSIT */, "EUCORP" /* EUROCORPORATEBOND */, "EUCP" /* EUROCOMMERCIALPAPER */, "EUSOV" /* EUROSOVEREIGNS */, "EUSUPRA" /* EUROSUPRANATIONALCOUPONS */, "FAC" /* FEDERALAGENCYCOUPON */, "FADN" /* FEDERALAGENCYDISCOUNTNOTE */, "FOR" /* FOREIGNEXCHANGECONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERALOBLIGATIONBONDS */, "IET" /* IOETTEMORTGAGE */, "LOFC" /* LETTEROFCREDIT */, "LQN" /* LIQUIDITYNOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGEBACKEDSECURITIES */, "MF" /* MUTUALFUND */, "MIO" /* MORTGAGEINTERESTONLY */, "MLEG" /* MULTILEGINSTRUMENT */, "MPO" /* MORTGAGEPRINCIPALONLY */, "MPP" /* MORTGAGEPRIVATEPLACEMENT */, "MPT" /* MISCELLANEOUSPASSTHROUGH */, "MT" /* MANDATORYTENDER */, "MTN" /* MEDIUMTERMNOTES */, "NONE" /* NOSECURITYTYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATEEXPORTFUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORYNOTE */, "PS" /* PREFERREDSTOCK */, "PZFJ" /* PLAZOSFIJOS */, "RAN" /* REVENUEANTICIPATIONNOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUEBONDS */, "RVLV" /* REVOLVERLOAN */, "RVLVTRM" /* REVOLVERTERMLOAN */, "SECLOAN" /* SECURITIESLOAN */, "SECPLEDGE" /* SECURITIESPLEDGE */, "SPCLA" /* SPECIALASSESSMENT */, "SPCLO" /* SPECIALOBLIGATION */, "SPCLT" /* SPECIALTAX */, "STN" /* SHORTTERMLOANNOTE */, "STRUCT" /* STRUCTUREDNOTES */, "SUPRA" /* USDSUPRANATIONALCOUPONS */, "SWING" /* SWINGLINEFACILITY */, "TAN" /* TAXANTICIPATIONNOTE */, "TAXA" /* TAXALLOCATION */, "TBA" /* TOBEANNOUNCED */, "TBILL" /* USTREASURYBILL */, "TBOND" /* USTREASURYBOND */, "TCAL" /* PRINCIPALSTRIPOFACALLABLEBONDORNOTE */, "TD" /* TIMEDEPOSIT */, "TECP" /* TAXEXEMPTCOMMERCIALPAPER */, "TERM" /* TERMLOAN */, "TINT" /* INTERESTSTRIPFROMANYBONDORNOTE */, "TIPS" /* TREASURYINFLATIONPROTECTEDSECURITIES */, "TNOTE" /* USTREASURYNOTE */, "TPRN" /* PRINCIPALSTRIPFROMANONCALLABLEBONDORNOTE */, "TRAN" /* TAXREVENUEANTICIPATIONNOTE */, "UST" /* USTREASURYNOTEDEPRECATEDVALUEUSETNOTE */, "USTB" /* USTREASURYBILLDEPRECATEDVALUEUSETBILL */, "VRDN" /* VARIABLERATEDEMANDNOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "WLD" /* WILDCARDENTRY */, "XCN" /* EXTENDEDCOMMNOTE */, "XLINKD" /* INDEXEDLINKED */, "YANK" /* YANKEECORPORATEBOND */, "YCD" /* YANKEECERTIFICATEOFDEPOSIT */), f0.Var{1, 65536}},
		SecuritySubType762:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MaturityMonthYear200:          f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		MaturityDate541:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		CouponPaymentDate224:          f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		IssueDate225:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		RepoCollateralSecurityType239: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseTerm226:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseRate227:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Factor228:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CreditRating255:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrRegistry543:              f0.Fld{Opt, f0.ASCII, f0.String("BIC" /* CUSTODIAN */, "ISO" /* COUNTRY */, "ZZ" /* PHYSICAL */), f0.Var{1, 65536}},
		CountryOfIssue470:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		StateOrProvinceOfIssue471:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LocaleOfIssue472:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RedemptionDate240:             f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		StrikePrice202:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		StrikeCurrency947:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		OptAttribute206:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContractMultiplier231:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CouponRate223:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SecurityExchange207:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Issuer106:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedIssuerLen348:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedIssuer349:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityDesc107:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedSecurityDescLen350:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedSecurityDesc351:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		Pool691:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ContractSettlMonth667:         f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		CPProgram875:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* 3A3 */, 2 /* 42 */, 99 /* OTHER */), f0.Var{1, 19}},
		CPRegType876:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DatedDate873:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		InterestAccrualDate874:        f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		SecurityAltID455:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityAltIDSource456:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EventType865:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PUT */, 2 /* CALL */, 3 /* TENDER */, 4 /* SINKINGFUNDCALL */, 99 /* OTHER */), f0.Var{1, 19}},
		EventDate866:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		EventPx867:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		EventText868:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SignatureLength93:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:                   f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,                  // STRING
		BodyLength9,                   // LENGTH
		MsgType35,                     // STRING
		SenderCompID49,                // STRING
		TargetCompID56,                // STRING
		OnBehalfOfCompID115,           // STRING
		DeliverToCompID128,            // STRING
		SecureDataLen90,               // LENGTH
		SecureData91,                  // DATA
		MsgSeqNum34,                   // SEQNUM
		SenderSubID50,                 // STRING
		SenderLocationID142,           // STRING
		TargetSubID57,                 // STRING
		TargetLocationID143,           // STRING
		OnBehalfOfSubID116,            // STRING
		OnBehalfOfLocationID144,       // STRING
		DeliverToSubID129,             // STRING
		DeliverToLocationID145,        // STRING
		PossDupFlag43,                 // BOOLEAN
		PossResend97,                  // BOOLEAN
		SendingTime52,                 // UTCTIMESTAMP
		OrigSendingTime122,            // UTCTIMESTAMP
		XmlDataLen212,                 // LENGTH
		XmlData213,                    // DATA
		MessageEncoding347,            // STRING
		LastMsgSeqNumProcessed369,     // SEQNUM
		HopCompID628,                  // STRING
		HopSendingTime629,             // UTCTIMESTAMP
		HopRefID630,                   // SEQNUM
		OrderID37,                     // STRING
		CrossID548,                    // STRING
		OrigCrossID551,                // STRING
		CrossType549,                  // INT
		CrossPrioritization550,        // INT
		TransactTime60,                // UTCTIMESTAMP
		Symbol55,                      // STRING
		SymbolSfx65,                   // STRING
		SecurityID48,                  // STRING
		SecurityIDSource22,            // STRING
		Product460,                    // INT
		CFICode461,                    // STRING
		SecurityType167,               // STRING
		SecuritySubType762,            // STRING
		MaturityMonthYear200,          // MONTHYEAR
		MaturityDate541,               // LOCALMKTDATE
		CouponPaymentDate224,          // LOCALMKTDATE
		IssueDate225,                  // LOCALMKTDATE
		RepoCollateralSecurityType239, // INT
		RepurchaseTerm226,             // INT
		RepurchaseRate227,             // PERCENTAGE
		Factor228,                     // FLOAT
		CreditRating255,               // STRING
		InstrRegistry543,              // STRING
		CountryOfIssue470,             // COUNTRY
		StateOrProvinceOfIssue471,     // STRING
		LocaleOfIssue472,              // STRING
		RedemptionDate240,             // LOCALMKTDATE
		StrikePrice202,                // PRICE
		StrikeCurrency947,             // CURRENCY
		OptAttribute206,               // CHAR
		ContractMultiplier231,         // FLOAT
		CouponRate223,                 // PERCENTAGE
		SecurityExchange207,           // EXCHANGE
		Issuer106,                     // STRING
		EncodedIssuerLen348,           // LENGTH
		EncodedIssuer349,              // DATA
		SecurityDesc107,               // STRING
		EncodedSecurityDescLen350,     // LENGTH
		EncodedSecurityDesc351,        // DATA
		Pool691,                       // STRING
		ContractSettlMonth667,         // MONTHYEAR
		CPProgram875,                  // INT
		CPRegType876,                  // STRING
		DatedDate873,                  // LOCALMKTDATE
		InterestAccrualDate874,        // LOCALMKTDATE
		SecurityAltID455,              // STRING
		SecurityAltIDSource456,        // STRING
		EventType865,                  // INT
		EventDate866,                  // LOCALMKTDATE
		EventPx867,                    // PRICE
		EventText868,                  // STRING
		SignatureLength93,             // LENGTH
		Signature89,                   // DATA
		CheckSum10,                    // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8                  = 8   // STRING
	BodyLength9                   = 9   // LENGTH
	MsgType35                     = 35  // STRING
	SenderCompID49                = 49  // STRING
	TargetCompID56                = 56  // STRING
	OnBehalfOfCompID115           = 115 // STRING
	DeliverToCompID128            = 128 // STRING
	SecureDataLen90               = 90  // LENGTH
	SecureData91                  = 91  // DATA
	MsgSeqNum34                   = 34  // SEQNUM
	SenderSubID50                 = 50  // STRING
	SenderLocationID142           = 142 // STRING
	TargetSubID57                 = 57  // STRING
	TargetLocationID143           = 143 // STRING
	OnBehalfOfSubID116            = 116 // STRING
	OnBehalfOfLocationID144       = 144 // STRING
	DeliverToSubID129             = 129 // STRING
	DeliverToLocationID145        = 145 // STRING
	PossDupFlag43                 = 43  // BOOLEAN
	PossResend97                  = 97  // BOOLEAN
	SendingTime52                 = 52  // UTCTIMESTAMP
	OrigSendingTime122            = 122 // UTCTIMESTAMP
	XmlDataLen212                 = 212 // LENGTH
	XmlData213                    = 213 // DATA
	MessageEncoding347            = 347 // STRING
	LastMsgSeqNumProcessed369     = 369 // SEQNUM
	HopCompID628                  = 628 // STRING
	HopSendingTime629             = 629 // UTCTIMESTAMP
	HopRefID630                   = 630 // SEQNUM
	OrderID37                     = 37  // STRING
	CrossID548                    = 548 // STRING
	OrigCrossID551                = 551 // STRING
	CrossType549                  = 549 // INT
	CrossPrioritization550        = 550 // INT
	TransactTime60                = 60  // UTCTIMESTAMP
	Symbol55                      = 55  // STRING
	SymbolSfx65                   = 65  // STRING
	SecurityID48                  = 48  // STRING
	SecurityIDSource22            = 22  // STRING
	Product460                    = 460 // INT
	CFICode461                    = 461 // STRING
	SecurityType167               = 167 // STRING
	SecuritySubType762            = 762 // STRING
	MaturityMonthYear200          = 200 // MONTHYEAR
	MaturityDate541               = 541 // LOCALMKTDATE
	CouponPaymentDate224          = 224 // LOCALMKTDATE
	IssueDate225                  = 225 // LOCALMKTDATE
	RepoCollateralSecurityType239 = 239 // INT
	RepurchaseTerm226             = 226 // INT
	RepurchaseRate227             = 227 // PERCENTAGE
	Factor228                     = 228 // FLOAT
	CreditRating255               = 255 // STRING
	InstrRegistry543              = 543 // STRING
	CountryOfIssue470             = 470 // COUNTRY
	StateOrProvinceOfIssue471     = 471 // STRING
	LocaleOfIssue472              = 472 // STRING
	RedemptionDate240             = 240 // LOCALMKTDATE
	StrikePrice202                = 202 // PRICE
	StrikeCurrency947             = 947 // CURRENCY
	OptAttribute206               = 206 // CHAR
	ContractMultiplier231         = 231 // FLOAT
	CouponRate223                 = 223 // PERCENTAGE
	SecurityExchange207           = 207 // EXCHANGE
	Issuer106                     = 106 // STRING
	EncodedIssuerLen348           = 348 // LENGTH
	EncodedIssuer349              = 349 // DATA
	SecurityDesc107               = 107 // STRING
	EncodedSecurityDescLen350     = 350 // LENGTH
	EncodedSecurityDesc351        = 351 // DATA
	Pool691                       = 691 // STRING
	ContractSettlMonth667         = 667 // MONTHYEAR
	CPProgram875                  = 875 // INT
	CPRegType876                  = 876 // STRING
	DatedDate873                  = 873 // LOCALMKTDATE
	InterestAccrualDate874        = 874 // LOCALMKTDATE
	SecurityAltID455              = 455 // STRING
	SecurityAltIDSource456        = 456 // STRING
	EventType865                  = 865 // INT
	EventDate866                  = 866 // LOCALMKTDATE
	EventPx867                    = 867 // PRICE
	EventText868                  = 868 // STRING
	SignatureLength93             = 93  // LENGTH
	Signature89                   = 89  // DATA
	CheckSum10                    = 10  // STRING
)
