// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix44securitytyperequest is a format of the FIX.4.4 SecurityTypeRequest message.
package fix44securitytyperequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX44SecurityTypeRequestMarshaler   = marshfix.Marshal{Tag: "FIX44", Format: FIX44SecurityTypeRequest}
	FIX44SecurityTypeRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX44", Format: FIX44SecurityTypeRequest}
)

// FIX44SecurityTypeRequest is a FIX.4.4 format of the SecurityTypeRequest message which maps the codecs into individual fields.
var FIX44SecurityTypeRequest = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:              f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:                 f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:               f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
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
		MessageEncoding347:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369: f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		HopCompID628:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:         f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:               f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SecurityReqID320:          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		TradingSessionID336:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionSubID625:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Product460:                f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		SecurityType167:           f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSETBACKEDSECURITIES */, "AMENDED" /* AMENDEDRESTATED */, "AN" /* OTHERANTICIPATIONNOTESBANGANETC */, "BA" /* BANKERSACCEPTANCE */, "BN" /* BANKNOTES */, "BOX" /* BILLOFEXCHANGES */, "BRADY" /* BRADYBOND */, "BRIDGE" /* BRIDGELOAN */, "BUYSELL" /* BUYSELLBACK */, "CB" /* CONVERTIBLEBOND */, "CD" /* CERTIFICATEOFDEPOSIT */, "CL" /* CALLLOANS */, "CMBS" /* CORPMORTGAGEBACKEDSECURITIES */, "CMO" /* COLLATERALIZEDMORTGAGEOBLIGATION */, "COFO" /* CERTIFICATEOFOBLIGATION */, "COFP" /* CERTIFICATEOFPARTICIPATION */, "CORP" /* CORPORATEBOND */, "CP" /* COMMERCIALPAPER */, "CPP" /* CORPORATEPRIVATEPLACEMENT */, "CS" /* COMMONSTOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTORINPOSSESSION */, "DN" /* DEPOSITNOTES */, "DUAL" /* DUALCURRENCY */, "EUCD" /* EUROCERTIFICATEOFDEPOSIT */, "EUCORP" /* EUROCORPORATEBOND */, "EUCP" /* EUROCOMMERCIALPAPER */, "EUSOV" /* EUROSOVEREIGNS */, "EUSUPRA" /* EUROSUPRANATIONALCOUPONS */, "FAC" /* FEDERALAGENCYCOUPON */, "FADN" /* FEDERALAGENCYDISCOUNTNOTE */, "FOR" /* FOREIGNEXCHANGECONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERALOBLIGATIONBONDS */, "IET" /* IOETTEMORTGAGE */, "LOFC" /* LETTEROFCREDIT */, "LQN" /* LIQUIDITYNOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGEBACKEDSECURITIES */, "MF" /* MUTUALFUND */, "MIO" /* MORTGAGEINTERESTONLY */, "MLEG" /* MULTILEGINSTRUMENT */, "MPO" /* MORTGAGEPRINCIPALONLY */, "MPP" /* MORTGAGEPRIVATEPLACEMENT */, "MPT" /* MISCELLANEOUSPASSTHROUGH */, "MT" /* MANDATORYTENDER */, "MTN" /* MEDIUMTERMNOTES */, "NONE" /* NOSECURITYTYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATEEXPORTFUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORYNOTE */, "PS" /* PREFERREDSTOCK */, "PZFJ" /* PLAZOSFIJOS */, "RAN" /* REVENUEANTICIPATIONNOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUEBONDS */, "RVLV" /* REVOLVERLOAN */, "RVLVTRM" /* REVOLVERTERMLOAN */, "SECLOAN" /* SECURITIESLOAN */, "SECPLEDGE" /* SECURITIESPLEDGE */, "SPCLA" /* SPECIALASSESSMENT */, "SPCLO" /* SPECIALOBLIGATION */, "SPCLT" /* SPECIALTAX */, "STN" /* SHORTTERMLOANNOTE */, "STRUCT" /* STRUCTUREDNOTES */, "SUPRA" /* USDSUPRANATIONALCOUPONS */, "SWING" /* SWINGLINEFACILITY */, "TAN" /* TAXANTICIPATIONNOTE */, "TAXA" /* TAXALLOCATION */, "TBA" /* TOBEANNOUNCED */, "TBILL" /* USTREASURYBILL */, "TBOND" /* USTREASURYBOND */, "TCAL" /* PRINCIPALSTRIPOFACALLABLEBONDORNOTE */, "TD" /* TIMEDEPOSIT */, "TECP" /* TAXEXEMPTCOMMERCIALPAPER */, "TERM" /* TERMLOAN */, "TINT" /* INTERESTSTRIPFROMANYBONDORNOTE */, "TIPS" /* TREASURYINFLATIONPROTECTEDSECURITIES */, "TNOTE" /* USTREASURYNOTE */, "TPRN" /* PRINCIPALSTRIPFROMANONCALLABLEBONDORNOTE */, "TRAN" /* TAXREVENUEANTICIPATIONNOTE */, "UST" /* USTREASURYNOTEDEPRECATEDVALUEUSETNOTE */, "USTB" /* USTREASURYBILLDEPRECATEDVALUEUSETBILL */, "VRDN" /* VARIABLERATEDEMANDNOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "WLD" /* WILDCARDENTRY */, "XCN" /* EXTENDEDCOMMNOTE */, "XLINKD" /* INDEXEDLINKED */, "YANK" /* YANKEECORPORATEBOND */, "YCD" /* YANKEECERTIFICATEOFDEPOSIT */), f0.Var{1, 65536}},
		SecuritySubType762:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SignatureLength93:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:               f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,              // STRING
		BodyLength9,               // LENGTH
		MsgType35,                 // STRING
		SenderCompID49,            // STRING
		TargetCompID56,            // STRING
		OnBehalfOfCompID115,       // STRING
		DeliverToCompID128,        // STRING
		SecureDataLen90,           // LENGTH
		SecureData91,              // DATA
		MsgSeqNum34,               // SEQNUM
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
		LastMsgSeqNumProcessed369, // SEQNUM
		HopCompID628,              // STRING
		HopSendingTime629,         // UTCTIMESTAMP
		HopRefID630,               // SEQNUM
		SecurityReqID320,          // STRING
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		TradingSessionID336,       // STRING
		TradingSessionSubID625,    // STRING
		Product460,                // INT
		SecurityType167,           // STRING
		SecuritySubType762,        // STRING
		SignatureLength93,         // LENGTH
		Signature89,               // DATA
		CheckSum10,                // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8              = 8   // STRING
	BodyLength9               = 9   // LENGTH
	MsgType35                 = 35  // STRING
	SenderCompID49            = 49  // STRING
	TargetCompID56            = 56  // STRING
	OnBehalfOfCompID115       = 115 // STRING
	DeliverToCompID128        = 128 // STRING
	SecureDataLen90           = 90  // LENGTH
	SecureData91              = 91  // DATA
	MsgSeqNum34               = 34  // SEQNUM
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
	LastMsgSeqNumProcessed369 = 369 // SEQNUM
	HopCompID628              = 628 // STRING
	HopSendingTime629         = 629 // UTCTIMESTAMP
	HopRefID630               = 630 // SEQNUM
	SecurityReqID320          = 320 // STRING
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	TradingSessionID336       = 336 // STRING
	TradingSessionSubID625    = 625 // STRING
	Product460                = 460 // INT
	SecurityType167           = 167 // STRING
	SecuritySubType762        = 762 // STRING
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
