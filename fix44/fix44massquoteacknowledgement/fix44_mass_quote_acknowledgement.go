// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix44massquoteacknowledgement is a format of the FIX.4.4 MassQuoteAcknowledgement message.
package fix44massquoteacknowledgement

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX44MassQuoteAcknowledgementMarshaler   = marshfix.Marshal{Tag: "FIX44", Format: FIX44MassQuoteAcknowledgement}
	FIX44MassQuoteAcknowledgementUnmarshaler = marshfix.Unmarshal{Tag: "FIX44", Format: FIX44MassQuoteAcknowledgement}
)

// FIX44MassQuoteAcknowledgement is a FIX.4.4 format of the MassQuoteAcknowledgement message which maps the codecs into individual fields.
var FIX44MassQuoteAcknowledgement = f0.Format{
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
		QuoteReqID131:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		QuoteID117:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		QuoteStatus297:            f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* ACCPT */, 1 /* CXLSYM */, 10 /* PENDING */, 11 /* PASS */, 12 /* LOCKEDMARKETWARNING */, 13 /* CROSSMARKETWARNING */, 14 /* CANCELEDDUETOLOCKMARKET */, 15 /* CANCELEDDUETOCROSSMARKET */, 2 /* CXLSECTYPE */, 3 /* CXLUNDER */, 4 /* CXLALL */, 5 /* REJ */, 6 /* REMOVED */, 7 /* EXPIRED */, 8 /* QUERY */, 9 /* QUOTENOTFOUND */), f0.Var{1, 19}},
		QuoteRejectReason300:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* UNKNSYM */, 2 /* EXCHCLSD */, 3 /* ORDEXLIM */, 4 /* TOOLATE */, 5 /* UNKNORD */, 6 /* DUPORD */, 7 /* INVSPREAD */, 8 /* INVPX */, 9 /* NOTAUTH */, 99 /* OTHER */), f0.Var{1, 19}},
		QuoteResponseLevel301:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* NOACK */, 1 /* ACKNEG */, 2 /* ACKEACH */), f0.Var{1, 19}},
		QuoteType537:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* INDICATIVE */, 1 /* TRADEABLE */, 2 /* RESTRICTEDTRADEABLE */, 3 /* COUNTER */), f0.Var{1, 19}},
		Account1:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AcctIDSource660:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BIC */, 2 /* SIDCODE */, 3 /* TFMGSPTA */, 4 /* OMGEOALERTID */, 5 /* DTCCCODE */, 99 /* OTHER */), f0.Var{1, 19}},
		AccountType581:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* ACCOUNTCUSTOMER */, 2 /* ACCOUNTNONCUSTOMER */, 3 /* HOUSETRADER */, 4 /* FLOORTRADER */, 6 /* ACCOUNTNONCUSTOMERCROSS */, 7 /* HOUSETRADERCROSS */, 8 /* JOINTBOACCT */), f0.Var{1, 19}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		PartyID448:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:          f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREANINVESTORID */, "2" /* TAIWANESEQUALIFIED */, "3" /* TAIWANESETRADINGACCT */, "4" /* MCDNUMBER */, "5" /* CHINESEBSHARE */, "6" /* UKNATIONALINSPENNUMBER */, "7" /* USSOCIALSECURITY */, "8" /* USEMPLOYERIDNUMBER */, "9" /* AUSTRALIANBUSINESSNUMBER */, "A" /* AUSTRALIANTAXFILENUMBER */, "B" /* BIC */, "C" /* ACCPTMARKETPART */, "D" /* PROPCODE */, "E" /* ISOCODE */, "F" /* SETTLENTLOC */, "G" /* MIC */, "H" /* CSDPARTCODE */, "I" /* DIRECTEDDEFINEDISITC */), f0.Con{1}},
		PartyRole452:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTINGFIRM */, 10 /* SETTLEMENTLOCATION */, 11 /* INITIATINGTRADER */, 12 /* EXECUTINGTRADER */, 13 /* ORDERORIGINATOR */, 14 /* GIVEUPCLEARINGFIRM */, 15 /* CORRESPONDANTCLEARINGFIRM */, 16 /* EXECUTINGSYSTEM */, 17 /* CONTRAFIRM */, 18 /* CONTRACLEARINGFIRM */, 19 /* SPONSORINGFIRM */, 2 /* BROKEROFCREDIT */, 20 /* UNDRCONTRAFIRM */, 21 /* CLEARINGORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMERACCOUNT */, 25 /* CORRESPONDENTCLEARINGORGANIZATION */, 26 /* CORRESPONDENTBROKER */, 27 /* BUYERSELLERRECEIVERDELIVERER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENTID */, 30 /* AGENT */, 31 /* SUBCUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTEDPARTY */, 34 /* REGULATORYBODY */, 35 /* LIQUIDITYPROVIDER */, 36 /* ENTERINGTRADER */, 37 /* CONTRATRADER */, 38 /* POSITIONACCOUNT */, 39 /* ALLOCENTITY */, 4 /* CLEARINGFIRM */, 5 /* INVESTORID */, 6 /* INTRODUCINGFIRM */, 7 /* ENTERINGFIRM */, 8 /* LOCATELENDINGFIRM */, 9 /* FUNDMANAGER */), f0.Var{1, 19}},
		PartySubID523:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:         f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIESACCOUNTNUMBER */, 11 /* REGISTRATIONNUMBER */, 12 /* REGISTEREDADDRESS_12 */, 13 /* REGULATORYSTATUS */, 14 /* REGISTRATIONNAME */, 15 /* CASHACCOUNT */, 16 /* BIC */, 17 /* CSDPARTICIPANTMEMBERCODE */, 18 /* REGISTEREDADDRESS_18 */, 19 /* FUNDACCOUNTNAME */, 2 /* PERSON */, 20 /* TELEXNUMBER */, 21 /* FAXNUMBER */, 22 /* SECURITIESACCOUNTNAME */, 23 /* CASHACCOUNTNAME */, 24 /* DEPARTMENT */, 25 /* LOCATIONDESK */, 26 /* POSITIONACCOUNTTYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 4000 /* RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES */, 5 /* FULLLEGALNAMEOFFIRM */, 6 /* POSTALADDRESS */, 7 /* PHONENUMBER */, 8 /* EMAILADDRESS */, 9 /* CONTACTNAME */), f0.Var{1, 19}},
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
		QuoteReqID131,             // STRING
		QuoteID117,                // STRING
		QuoteStatus297,            // INT
		QuoteRejectReason300,      // INT
		QuoteResponseLevel301,     // INT
		QuoteType537,              // INT
		Account1,                  // STRING
		AcctIDSource660,           // INT
		AccountType581,            // INT
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		PartyID448,                // STRING
		PartyIDSource447,          // CHAR
		PartyRole452,              // INT
		PartySubID523,             // STRING
		PartySubIDType803,         // INT
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
	QuoteReqID131             = 131 // STRING
	QuoteID117                = 117 // STRING
	QuoteStatus297            = 297 // INT
	QuoteRejectReason300      = 300 // INT
	QuoteResponseLevel301     = 301 // INT
	QuoteType537              = 537 // INT
	Account1                  = 1   // STRING
	AcctIDSource660           = 660 // INT
	AccountType581            = 581 // INT
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	PartyID448                = 448 // STRING
	PartyIDSource447          = 447 // CHAR
	PartyRole452              = 452 // INT
	PartySubID523             = 523 // STRING
	PartySubIDType803         = 803 // INT
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
