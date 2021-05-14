// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix44ioi is a format of the FIX.4.4 IOI message.
package fix44ioi

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX44IOIMarshaler   = marshfix.Marshal{Tag: "FIX44", Format: FIX44IOI}
	FIX44IOIUnmarshaler = marshfix.Unmarshal{Tag: "FIX44", Format: FIX44IOI}
)

// FIX44IOI is a FIX.4.4 format of the IOI message which maps the codecs into individual fields.
var FIX44IOI = f0.Format{
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
		IOIID23:                       f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IOITransType28:                f0.Fld{Req, f0.ASCII, f0.String("C" /* CANCEL */, "N" /* NEW */, "R" /* REPLACE */), f0.Con{1}},
		IOIRefID26:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Side54:                        f0.Fld{Req, f0.ASCII, f0.String("1" /* BUY */, "2" /* SELL */, "3" /* BUYMIN */, "4" /* SELLPLUS */, "5" /* SELLSHT */, "6" /* SELLSHTEX */, "7" /* UNDISC */, "8" /* CROSS */, "9" /* CROSSSHORT */, "A" /* CROSSSHORTEX */, "B" /* ASDEFINED */, "C" /* OPPOSITE */, "D" /* SUBSCRIBE */, "E" /* REDEEM */, "F" /* LENDFINANCING */, "G" /* BORROWFINANCING */), f0.Con{1}},
		QtyType854:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* UNITS */, 1 /* CONTRACTS */), f0.Var{1, 19}},
		IOIQty27:                      f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Currency15:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		PriceType423:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PCT */, 10 /* FIXEDCABINETTRADEPRICE */, 11 /* VARIABLECABINETTRADEPRICE */, 2 /* CPS */, 3 /* ABS */, 4 /* DISCOUNT */, 5 /* PREMIUM */, 6 /* SPREAD */, 7 /* TEDPRICE */, 8 /* TEDYIELD */, 9 /* YIELD */), f0.Var{1, 19}},
		Price44:                       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ValidUntilTime62:              f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		IOIQltyInd25:                  f0.Fld{Opt, f0.ASCII, f0.String("H" /* HIGH */, "L" /* LOW */, "M" /* MEDIUM */), f0.Con{1}},
		IOINaturalFlag130:             f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		Text58:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		TransactTime60:                f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		URLLink149:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
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
		AgreementDesc913:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AgreementID914:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AgreementDate915:              f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		AgreementCurrency918:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		TerminationType788:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* OVERNIGHT */, 2 /* TERM */, 3 /* FLEXIBLE */, 4 /* OPEN */), f0.Var{1, 19}},
		StartDate916:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		EndDate917:                    f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		DeliveryType919:               f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* VERSUSPAYMENT */, 1 /* FREE */, 2 /* TRIPARTY */, 3 /* HOLDINCUSTODY */), f0.Var{1, 19}},
		MarginRatio898:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrderQty38:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CashOrderQty152:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrderPercent516:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RoundingDirection468:          f0.Fld{Opt, f0.ASCII, f0.String("0" /* ROUNDNEAREST */, "1" /* ROUNDDOWN */, "2" /* ROUNDUP */), f0.Con{1}},
		RoundingModulus469:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		StipulationType233:            f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ABSOLUTEPREPAYMENTSPEED */, "AMT" /* AMT */, "AUTOREINV" /* AUTOREINVESTMENTATRATEORBETTER */, "BANKQUAL" /* BANKQUALIFIED */, "BGNCON" /* BARGAINCONDITIONS */, "COUPON" /* COUPONRANGE */, "CPP" /* CONSTANTPREPAYMENTPENALTY */, "CPR" /* CONSTANTPREPAYMENTRATE */, "CPY" /* CONSTANTPREPAYMENTYIELD */, "CURRENCY" /* ISOCURRENCYCODE */, "CUSTOMDATE" /* CUSTOMSTARTENDDATE */, "GEOG" /* GEOGRAPHICSANDRANGE */, "HAIRCUT" /* VALUATIONDISCOUNT */, "HEP" /* FINALCPROFHOMEEQUITYPREPAYMENTCURVE */, "INSURED" /* INSURED */, "ISSUE" /* YEARORYEARMONTHOFISSUE */, "ISSUER" /* ISSUERSTICKER */, "ISSUESIZE" /* ISSUESIZERANGE */, "LOOKBACK" /* LOOKBACKDAYS */, "LOT" /* EXPLICITLOTIDENTIFIER */, "LOTVAR" /* LOTVARIANCEVALUEINPERCENTMAXIMUMOVERORUNDERALLOCATIONALLOWED */, "MAT" /* MATURITYYEARANDMONTH */, "MATURITY" /* MATURITYRANGE */, "MAXDNOM" /* MAXIMUMDENOMINATION */, "MAXSUBS" /* MAXIMUMSUBSTITUTIONSREPO */, "MHP" /* PERCENTOFMANUFACTUREDHOUSINGPREPAYMENTCURVE */, "MINDNOM" /* MINIMUMDENOMINATION */, "MININCR" /* MINIMUMINCREMENT */, "MINQTY" /* MINIMUMQUANTITY */, "MPR" /* MONTHLYPREPAYMENTRATE */, "PAYFREQ" /* PAYMENTFREQUENCYCALENDAR */, "PIECES" /* NUMBEROFPIECES */, "PMAX" /* POOLSMAXIMUM */, "PMIN" /* POOLSMINIMUM */, "PPC" /* PERCENTOFPROSPECTUSPREPAYMENTCURVE */, "PPL" /* POOLSPERLOT */, "PPM" /* POOLSPERMILLION */, "PPT" /* POOLSPERTRADE */, "PRICE" /* PRICERANGE */, "PRICEFREQ" /* PRICINGFREQUENCY */, "PROD" /* PRODUCTIONYEAR */, "PROTECT" /* CALLPROTECTION */, "PSA" /* PERCENTOFBMAPREPAYMENTCURVE */, "PURPOSE" /* PURPOSE */, "PXSOURCE" /* BENCHMARKPRICESOURCE */, "RATING" /* RATINGSOURCEANDRANGE */, "REDEMPTION" /* TYPEOFREDEMPTIONVALUESARE */, "RESTRICTED" /* RESTRICTED */, "SECTOR" /* MARKETSECTOR */, "SECTYPE" /* SECURITYTYPEINCLUDEDOREXCLUDED */, "SMM" /* SINGLEMONTHLYMORTALITY */, "STRUCT" /* STRUCTURE */, "SUBSFREQ" /* SUBSTITUTIONSFREQUENCYREPO */, "SUBSLEFT" /* SUBSTITUTIONSLEFTREPO */, "TEXT" /* FREEFORMTEXT */, "TRDVAR" /* TRADEVARIANCEVALUEINPERCENTMAXIMUMOVERORUNDERALLOCATIONALLOWED */, "WAC" /* WEIGHTEDAVERAGECOUPON */, "WAL" /* WEIGHTEDAVERAGELIFECOUPON */, "WALA" /* WEIGHTEDAVERAGELOANAGE */, "WAM" /* WEIGHTEDAVERAGEMATURITY */, "WHOLE" /* WHOLEPOOL */, "YIELD" /* YIELDRANGE */), f0.Var{1, 65536}},
		StipulationValue234:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Spread218:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		BenchmarkCurveCurrency220:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		BenchmarkCurveName221:         f0.Fld{Opt, f0.ASCII, f0.String("EONIA" /* EONIA */, "EUREPO" /* EUREPO */, "Euribor" /* EURIBOR */, "FutureSWAP" /* FUTURESWAP */, "LIBID" /* LIBID */, "LIBOR" /* LIBOR */, "MuniAAA" /* MUNIAAA */, "OTHER" /* OTHER */, "Pfandbriefe" /* PFANDBRIEFE */, "SONIA" /* SONIA */, "SWAP" /* SWAP */, "Treasury" /* TREASURY */), f0.Var{1, 65536}},
		BenchmarkCurvePoint222:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		BenchmarkPrice662:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		BenchmarkPriceType663:         f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		BenchmarkSecurityID699:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		BenchmarkSecurityIDSource761:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		YieldType235:                  f0.Fld{Opt, f0.ASCII, f0.String("AFTERTAX" /* AFTERTAXYIELD */, "ANNUAL" /* ANNUALYIELD */, "ATISSUE" /* YIELDATISSUE */, "AVGMATURITY" /* YIELDTOAVGMATURITY */, "BOOK" /* BOOKYIELD */, "CALL" /* YIELDTONEXTCALL */, "CHANGE" /* YIELDCHANGESINCECLOSE */, "CLOSE" /* CLOSINGYIELD */, "COMPOUND" /* COMPOUNDYIELD */, "CURRENT" /* CURRENTYIELD */, "GOVTEQUIV" /* GVNTEQUIVALENTYIELD */, "GROSS" /* TRUEGROSSYIELD */, "INFLATION" /* YIELDINFLATIONASSUMPTION */, "INVERSEFLOATER" /* INVFLOATERBONDYIELD */, "LASTCLOSE" /* MOSTRECENTCLOSINGYIELD */, "LASTMONTH" /* CLOSINGYIELDMOSTRECENTMONTH */, "LASTQUARTER" /* CLOSINGYIELDMOSTRECENTQUARTER */, "LASTYEAR" /* CLOSINGYIELDMOSTRECENTYEAR */, "LONGAVGLIFE" /* YIELDTOLONGESTAVERAGELIFE */, "MARK" /* MARKTOMARKETYIELD */, "MATURITY" /* YIELDTOMATURITY */, "NEXTREFUND" /* YIELDTONEXTREFUNDSINKING */, "OPENAVG" /* OPENAVERAGEYIELD */, "PREVCLOSE" /* PREVIOUSCLOSEYIELD */, "PROCEEDS" /* PROCEEDSYIELD */, "PUT" /* YIELDTONEXTPUT */, "SEMIANNUAL" /* SEMI */, "SHORTAVGLIFE" /* YIELDTOSHORTESTAVERAGELIFE */, "SIMPLE" /* SIMPLEYIELD */, "TAXEQUIV" /* TAXEQUIVALENTYIELD */, "TENDER" /* YIELDTOTENDERDATE */, "TRUE" /* TRUEYIELD */, "VALUE1_32" /* YIELDVALUEOF132 */, "WORST" /* YIELDTOWORSTCONVENTION */), f0.Var{1, 65536}},
		Yield236:                      f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		YieldCalcDate701:              f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		YieldRedemptionDate696:        f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		YieldRedemptionPrice697:       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		YieldRedemptionPriceType698:   f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
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
		IOIID23,                       // STRING
		IOITransType28,                // CHAR
		IOIRefID26,                    // STRING
		Side54,                        // CHAR
		QtyType854,                    // INT
		IOIQty27,                      // STRING
		Currency15,                    // CURRENCY
		PriceType423,                  // INT
		Price44,                       // PRICE
		ValidUntilTime62,              // UTCTIMESTAMP
		IOIQltyInd25,                  // CHAR
		IOINaturalFlag130,             // BOOLEAN
		Text58,                        // STRING
		EncodedTextLen354,             // LENGTH
		EncodedText355,                // DATA
		TransactTime60,                // UTCTIMESTAMP
		URLLink149,                    // STRING
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
		AgreementDesc913,              // STRING
		AgreementID914,                // STRING
		AgreementDate915,              // LOCALMKTDATE
		AgreementCurrency918,          // CURRENCY
		TerminationType788,            // INT
		StartDate916,                  // LOCALMKTDATE
		EndDate917,                    // LOCALMKTDATE
		DeliveryType919,               // INT
		MarginRatio898,                // PERCENTAGE
		OrderQty38,                    // QTY
		CashOrderQty152,               // QTY
		OrderPercent516,               // PERCENTAGE
		RoundingDirection468,          // CHAR
		RoundingModulus469,            // FLOAT
		StipulationType233,            // STRING
		StipulationValue234,           // STRING
		Spread218,                     // PRICEOFFSET
		BenchmarkCurveCurrency220,     // CURRENCY
		BenchmarkCurveName221,         // STRING
		BenchmarkCurvePoint222,        // STRING
		BenchmarkPrice662,             // PRICE
		BenchmarkPriceType663,         // INT
		BenchmarkSecurityID699,        // STRING
		BenchmarkSecurityIDSource761,  // STRING
		YieldType235,                  // STRING
		Yield236,                      // PERCENTAGE
		YieldCalcDate701,              // LOCALMKTDATE
		YieldRedemptionDate696,        // LOCALMKTDATE
		YieldRedemptionPrice697,       // PRICE
		YieldRedemptionPriceType698,   // INT
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
	IOIID23                       = 23  // STRING
	IOITransType28                = 28  // CHAR
	IOIRefID26                    = 26  // STRING
	Side54                        = 54  // CHAR
	QtyType854                    = 854 // INT
	IOIQty27                      = 27  // STRING
	Currency15                    = 15  // CURRENCY
	PriceType423                  = 423 // INT
	Price44                       = 44  // PRICE
	ValidUntilTime62              = 62  // UTCTIMESTAMP
	IOIQltyInd25                  = 25  // CHAR
	IOINaturalFlag130             = 130 // BOOLEAN
	Text58                        = 58  // STRING
	EncodedTextLen354             = 354 // LENGTH
	EncodedText355                = 355 // DATA
	TransactTime60                = 60  // UTCTIMESTAMP
	URLLink149                    = 149 // STRING
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
	AgreementDesc913              = 913 // STRING
	AgreementID914                = 914 // STRING
	AgreementDate915              = 915 // LOCALMKTDATE
	AgreementCurrency918          = 918 // CURRENCY
	TerminationType788            = 788 // INT
	StartDate916                  = 916 // LOCALMKTDATE
	EndDate917                    = 917 // LOCALMKTDATE
	DeliveryType919               = 919 // INT
	MarginRatio898                = 898 // PERCENTAGE
	OrderQty38                    = 38  // QTY
	CashOrderQty152               = 152 // QTY
	OrderPercent516               = 516 // PERCENTAGE
	RoundingDirection468          = 468 // CHAR
	RoundingModulus469            = 469 // FLOAT
	StipulationType233            = 233 // STRING
	StipulationValue234           = 234 // STRING
	Spread218                     = 218 // PRICEOFFSET
	BenchmarkCurveCurrency220     = 220 // CURRENCY
	BenchmarkCurveName221         = 221 // STRING
	BenchmarkCurvePoint222        = 222 // STRING
	BenchmarkPrice662             = 662 // PRICE
	BenchmarkPriceType663         = 663 // INT
	BenchmarkSecurityID699        = 699 // STRING
	BenchmarkSecurityIDSource761  = 761 // STRING
	YieldType235                  = 235 // STRING
	Yield236                      = 236 // PERCENTAGE
	YieldCalcDate701              = 701 // LOCALMKTDATE
	YieldRedemptionDate696        = 696 // LOCALMKTDATE
	YieldRedemptionPrice697       = 697 // PRICE
	YieldRedemptionPriceType698   = 698 // INT
	SignatureLength93             = 93  // LENGTH
	Signature89                   = 89  // DATA
	CheckSum10                    = 10  // STRING
)
