// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50allocationreportack is a format of the FIX.5.0 AllocationReportAck message.
package fix50allocationreportack

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50AllocationReportAckMarshaler   = marshfix.Marshal{Tag: "FIX50", Format: FIX50AllocationReportAck}
	FIX50AllocationReportAckUnmarshaler = marshfix.Unmarshal{Tag: "FIX50", Format: FIX50AllocationReportAck}
)

// FIX50AllocationReportAck is a FIX.5.0 format of the AllocationReportAck message which maps the codecs into individual fields.
var FIX50AllocationReportAck = f0.Format{
	Fields: map[int]f0.Codec{
		AllocReportID755:              f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AllocID70:                     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryAllocID793:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradeDate75:                   f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		TransactTime60:                f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		AllocStatus87:                 f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* ACCEPTED */, 1 /* BLOCK_LEVEL_REJECT */, 2 /* ACCOUNT_LEVEL_REJECT */, 3 /* RECEIVED */, 4 /* INCOMPLETE */, 5 /* REJECTED_BY_INTERMEDIARY */, 6 /* ALLOCATION_PENDING */, 7 /* REVERSED */), f0.Var{1, 19}},
		AllocRejCode88:                f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* UNKNOWN_ACCOUNT */, 1 /* INCORRECT_QUANTITY */, 10 /* UNKNOWN_OR_STALE_EXECID */, 11 /* MISMATCHED_DATA */, 12 /* UNKNOWN_CLORDID */, 13 /* WAREHOUSE_REQUEST_REJECTED */, 2 /* INCORRECT_AVERAGEG_PRICE */, 3 /* UNKNOWN_EXECUTING_BROKER_MNEMONIC */, 4 /* COMMISSION_DIFFERENCE */, 5 /* UNKNOWN_ORDERID */, 6 /* UNKNOWN_LISTID */, 7 /* OTHER */, 8 /* INCORRECT_ALLOCATED_QUANTITY */, 9 /* CALCULATION_DIFFERENCE */), f0.Var{1, 19}},
		AllocReportType794:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(3 /* SELLSIDE_CALCULATED_USING_PRELIMINARY */, 4 /* SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY */, 5 /* WAREHOUSE_RECAP */, 8 /* REQUEST_TO_INTERMEDIARY */, 2 /* PRELIMINARY_REQUEST_TO_INTERMEDIARY */, 9 /* ACCEPT */, 10 /* REJECT */, 11 /* ACCEPT_PENDING */, 12 /* COMPLETE */, 14 /* REVERSE_PENDING */), f0.Var{1, 19}},
		AllocIntermedReqType808:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PENDING_ACCEPT */, 2 /* PENDING_RELEASE */, 3 /* PENDING_REVERSAL */, 4 /* ACCEPT */, 5 /* BLOCK_LEVEL_REJECT */, 6 /* ACCOUNT_LEVEL_REJECT */), f0.Var{1, 19}},
		MatchStatus573:                f0.Fld{Opt, f0.ASCII, f0.String("0" /* COMPARED_MATCHED_OR_AFFIRMED */, "1" /* UNCOMPARED_UNMATCHED_OR_UNAFFIRED */, "2" /* ADVISORY_OR_ALERT */), f0.Con{1}},
		Product460:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		SecurityType167:               f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSET_BACKED_SECURITIES */, "AMENDED" /* AMENDED_RESTATED */, "AN" /* OTHER_ANTICIPATION_NOTES */, "BA" /* BANKERS_ACCEPTANCE */, "BN" /* BANK_NOTES */, "BOX" /* BILL_OF_EXCHANGES */, "BRADY" /* BRADY_BOND */, "BRIDGE" /* BRIDGE_LOAN */, "BUYSELL" /* BUY_SELLBACK */, "CB" /* CONVERTIBLE_BOND */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CL" /* CALL_LOANS */, "CMBS" /* CORP_MORTGAGE_BACKED_SECURITIES */, "CMO" /* COLLATERALIZED_MORTGAGE_OBLIGATION */, "COFO" /* CERTIFICATE_OF_OBLIGATION */, "COFP" /* CERTIFICATE_OF_PARTICIPATION */, "CORP" /* CORPORATE_BOND */, "CP" /* COMMERCIAL_PAPER */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CS" /* COMMON_STOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTOR_IN_POSSESSION */, "DN" /* DEPOSIT_NOTES */, "DUAL" /* DUAL_CURRENCY */, "EUCD" /* EURO_CERTIFICATE_OF_DEPOSIT */, "EUCORP" /* EURO_CORPORATE_BOND */, "EUCP" /* EURO_COMMERCIAL_PAPER */, "EUSOV" /* EURO_SOVEREIGNS */, "EUSUPRA" /* EURO_SUPRANATIONAL_COUPONS */, "FAC" /* FEDERAL_AGENCY_COUPON */, "FADN" /* FEDERAL_AGENCY_DISCOUNT_NOTE */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERAL_OBLIGATION_BONDS */, "IET" /* IOETTE_MORTGAGE */, "LOFC" /* LETTER_OF_CREDIT */, "LQN" /* LIQUIDITY_NOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGE_BACKED_SECURITIES */, "MF" /* MUTUAL_FUND */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "MLEG" /* MULTILEG_INSTRUMENT */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "MPT" /* MISCELLANEOUS_PASS_THROUGH */, "MT" /* MANDATORY_TENDER */, "MTN" /* MEDIUM_TERM_NOTES */, "NONE" /* NO_SECURITY_TYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATE_EXPORT_FUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORY_NOTE */, "PS" /* PREFERRED_STOCK */, "PZFJ" /* PLAZOS_FIJOS */, "RAN" /* REVENUE_ANTICIPATION_NOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUE_BONDS */, "RVLV" /* REVOLVER_LOAN */, "RVLVTRM" /* REVOLVER_TERM_LOAN */, "SECLOAN" /* SECURITIES_LOAN */, "SECPLEDGE" /* SECURITIES_PLEDGE */, "SPCLA" /* SPECIAL_ASSESSMENT */, "SPCLO" /* SPECIAL_OBLIGATION */, "SPCLT" /* SPECIAL_TAX */, "STN" /* SHORT_TERM_LOAN_NOTE */, "STRUCT" /* STRUCTURED_NOTES */, "SUPRA" /* USD_SUPRANATIONAL_COUPONS */, "SWING" /* SWING_LINE_FACILITY */, "TAN" /* TAX_ANTICIPATION_NOTE */, "TAXA" /* TAX_ALLOCATION */, "TBA" /* TO_BE_ANNOUNCED */, "TBILL" /* US_TREASURY_BILL_TBILL */, "TBOND" /* US_TREASURY_BOND */, "TCAL" /* PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE */, "TD" /* TIME_DEPOSIT */, "TECP" /* TAX_EXEMPT_COMMERCIAL_PAPER */, "TERM" /* TERM_LOAN */, "TINT" /* INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE */, "TIPS" /* TREASURY_INFLATION_PROTECTED_SECURITIES */, "TNOTE" /* US_TREASURY_NOTE_TNOTE */, "TPRN" /* PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE */, "TRAN" /* TAX_REVENUE_ANTICIPATION_NOTE */, "UST" /* US_TREASURY_NOTE_UST */, "USTB" /* US_TREASURY_BILL_USTB */, "VRDN" /* VARIABLE_RATE_DEMAND_NOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "WLD" /* WILDCARD_ENTRY */, "XCN" /* EXTENDED_COMM_NOTE */, "XLINKD" /* INDEXED_LINKED */, "YANK" /* YANKEE_CORPORATE_BOND */, "YCD" /* YANKEE_CERTIFICATE_OF_DEPOSIT */, "OOP" /* OPTIONS_ON_PHYSICAL */, "OOF" /* OPTIONS_ON_FUTURES */, "CASH" /* CASH */), f0.Var{1, 65536}},
		Text58:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		ClearingBusinessDate715:       f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		AvgPxIndicator819:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* NO_AVERAGE_PRICING */, 1 /* TRADE_IS_PART_OF_AN_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID */, 2 /* LAST_TRADE_IS_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID */), f0.Var{1, 19}},
		Quantity53:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		AllocTransType71:              f0.Fld{Opt, f0.ASCII, f0.String("0" /* NEW */, "1" /* REPLACE */, "2" /* CANCEL */, "3" /* PRELIMINARY */, "4" /* CALCULATED */, "5" /* CALCULATED_WITHOUT_PRELIMINARY */, "6" /* REVERSAL */), f0.Con{1}},
		PartyID448:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:              f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREAN_INVESTOR_ID */, "2" /* TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID */, "3" /* TAIWANESE_TRADING_ACCT */, "4" /* MALAYSIAN_CENTRAL_DEPOSITORY */, "5" /* CHINESE_INVESTOR_ID */, "6" /* UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER */, "7" /* US_SOCIAL_SECURITY_NUMBER */, "8" /* US_EMPLOYER_OR_TAX_ID_NUMBER */, "9" /* AUSTRALIAN_BUSINESS_NUMBER */, "A" /* AUSTRALIAN_TAX_FILE_NUMBER */, "B" /* BIC */, "C" /* GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER */, "D" /* PROPRIETARY */, "E" /* ISO_COUNTRY_CODE */, "F" /* SETTLEMENT_ENTITY_LOCATION */, "G" /* MIC */, "H" /* CSD_PARTICIPANT_MEMBER_CODE */, "I" /* DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT */), f0.Con{1}},
		PartyRole452:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTING_FIRM */, 10 /* SETTLEMENT_LOCATION */, 11 /* ORDER_ORIGINATION_TRADER */, 12 /* EXECUTING_TRADER */, 13 /* ORDER_ORIGINATION_FIRM */, 14 /* GIVEUP_CLEARING_FIRM */, 15 /* CORRESPONDANT_CLEARING_FIRM */, 16 /* EXECUTING_SYSTEM */, 17 /* CONTRA_FIRM */, 18 /* CONTRA_CLEARING_FIRM */, 19 /* SPONSORING_FIRM */, 2 /* BROKER_OF_CREDIT */, 20 /* UNDERLYING_CONTRA_FIRM */, 21 /* CLEARING_ORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMER_ACCOUNT */, 25 /* CORRESPONDENT_CLEARING_ORGANIZATION */, 26 /* CORRESPONDENT_BROKER */, 27 /* BUYER_SELLER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENT_ID */, 30 /* AGENT */, 31 /* SUB_CUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTED_PARTY */, 34 /* REGULATORY_BODY */, 35 /* LIQUIDITY_PROVIDER */, 36 /* ENTERING_TRADER */, 37 /* CONTRA_TRADER */, 38 /* POSITION_ACCOUNT */, 4 /* CLEARING_FIRM */, 5 /* INVESTOR_ID */, 6 /* INTRODUCING_FIRM */, 7 /* ENTERING_FIRM */, 8 /* LOCATE */, 9 /* FUND_MANAGER_CLIENT_ID */, 60 /* INTRODUCING_BROKER */, 41 /* CONTRA_POSITION_ACCOUNT */, 42 /* CONTRA_EXCHANGE */, 43 /* INTERNAL_CARRY_ACCOUNT */, 44 /* ORDER_ENTRY_OPERATOR_ID */, 45 /* SECONDARY_ACCOUNT_NUMBER */, 46 /* FORIEGN_FIRM */, 47 /* THIRD_PARTY_ALLOCATION_FIRM */, 48 /* CLAIMING_ACCOUNT */, 49 /* ASSET_MANAGER */, 50 /* PLEDGOR_ACCOUNT */, 51 /* PLEDGEE_ACCOUNT */, 52 /* LARGE_TRADER_REPORTABLE_ACCOUNT */, 53 /* TRADER_MNEMONIC */, 54 /* SENDER_LOCATION */, 55 /* SESSION_ID */, 56 /* ACCEPTABLE_COUNTERPARTY */, 57 /* UNACCEPTABLE_COUNTERPARTY */, 58 /* ENTERING_UNIT */, 59 /* EXECUTING_UNIT */, 39 /* CONTRA_INVESTOR_ID */, 40 /* TRANSFER_TO_FIRM */, 61 /* QUOTE_ORIGINATOR */, 62 /* REPORT_ORIGINATOR */, 63 /* SYSTEMATIC_INTERNALISER */, 64 /* MULTILATERAL_TRADING_FACILITY */, 65 /* REGULATED_MARKET */, 66 /* MARKET_MAKER */, 67 /* INVESTMENT_FIRM */, 68 /* HOST_COMPETENT_AUTHORITY */, 69 /* HOME_COMPETENT_AUTHORITY */, 70 /* COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY */, 71 /* COMPETENT_AUTHORITY_OF_THE_TRANSACTION */, 72 /* REPORTING_INTERMEDIARY */, 73 /* EXECUTION_VENUE */, 74 /* MARKET_DATA_ENTRY_ORIGINATOR */, 75 /* LOCATION_ID */, 76 /* DESK_ID */, 77 /* MARKET_DATA_MARKET */, 78 /* ALLOCATION_ENTITY */), f0.Var{1, 19}},
		PartySubID523:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIES_ACCOUNT_NUMBER */, 11 /* REGISTRATION_NUMBER */, 12 /* REGISTERED_ADDRESS_12 */, 13 /* REGULATORY_STATUS */, 14 /* REGISTRATION_NAME */, 15 /* CASH_ACCOUNT_NUMBER */, 16 /* BIC */, 17 /* CSD_PARTICIPANT_MEMBER_CODE */, 18 /* REGISTERED_ADDRESS_18 */, 19 /* FUND_ACCOUNT_NAME */, 2 /* PERSON */, 20 /* TELEX_NUMBER */, 21 /* FAX_NUMBER */, 22 /* SECURITIES_ACCOUNT_NAME */, 23 /* CASH_ACCOUNT_NAME */, 24 /* DEPARTMENT */, 25 /* LOCATION_DESK */, 26 /* POSITION_ACCOUNT_TYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 5 /* FULL_LEGAL_NAME_OF_FIRM */, 6 /* POSTAL_ADDRESS */, 7 /* PHONE_NUMBER */, 8 /* EMAIL_ADDRESS */, 9 /* CONTACT_NAME */, 27 /* SECURITY_LOCATE_ID */, 28 /* MARKET_MAKER */, 29 /* ELIGIBLE_COUNTERPARTY */, 30 /* PROFESSIONAL_CLIENT */, 31 /* LOCATION */, 32 /* EXECUTION_VENUE */), f0.Var{1, 19}},
		AllocAccount79:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AllocAcctIDSource661:          f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		AllocPrice366:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		IndividualAllocID467:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IndividualAllocRejCode776:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		AllocText161:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedAllocTextLen360:        f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedAllocText361:           f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecondaryIndividualAllocID989: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AllocCustomerCapacity993:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IndividualAllocType992:        f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* SUB_ALLOCATE */, 2 /* THIRD_PARTY_ALLOCATION */), f0.Var{1, 19}},
		AllocQty80:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		AllocPositionEffect1047:       f0.Fld{Opt, f0.ASCII, f0.String("O" /* OPEN */, "C" /* CLOSE */, "R" /* ROLLED */, "F" /* FIFO */), f0.Con{1}},
		NestedPartyID524:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NestedPartyIDSource525:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		NestedPartyRole538:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		NestedPartySubID545:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NestedPartySubIDType805:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		AllocReportID755,              // STRING
		AllocID70,                     // STRING
		SecondaryAllocID793,           // STRING
		TradeDate75,                   // LOCALMKTDATE
		TransactTime60,                // UTCTIMESTAMP
		AllocStatus87,                 // INT
		AllocRejCode88,                // INT
		AllocReportType794,            // INT
		AllocIntermedReqType808,       // INT
		MatchStatus573,                // CHAR
		Product460,                    // INT
		SecurityType167,               // STRING
		Text58,                        // STRING
		EncodedTextLen354,             // LENGTH
		EncodedText355,                // DATA
		ClearingBusinessDate715,       // LOCALMKTDATE
		AvgPxIndicator819,             // INT
		Quantity53,                    // QTY
		AllocTransType71,              // CHAR
		PartyID448,                    // STRING
		PartyIDSource447,              // CHAR
		PartyRole452,                  // INT
		PartySubID523,                 // STRING
		PartySubIDType803,             // INT
		AllocAccount79,                // STRING
		AllocAcctIDSource661,          // INT
		AllocPrice366,                 // PRICE
		IndividualAllocID467,          // STRING
		IndividualAllocRejCode776,     // INT
		AllocText161,                  // STRING
		EncodedAllocTextLen360,        // LENGTH
		EncodedAllocText361,           // DATA
		SecondaryIndividualAllocID989, // STRING
		AllocCustomerCapacity993,      // STRING
		IndividualAllocType992,        // INT
		AllocQty80,                    // QTY
		AllocPositionEffect1047,       // CHAR
		NestedPartyID524,              // STRING
		NestedPartyIDSource525,        // CHAR
		NestedPartyRole538,            // INT
		NestedPartySubID545,           // STRING
		NestedPartySubIDType805,       // INT
	},
}

const Req, Opt = true, false

const (
	AllocReportID755              = 755  // STRING
	AllocID70                     = 70   // STRING
	SecondaryAllocID793           = 793  // STRING
	TradeDate75                   = 75   // LOCALMKTDATE
	TransactTime60                = 60   // UTCTIMESTAMP
	AllocStatus87                 = 87   // INT
	AllocRejCode88                = 88   // INT
	AllocReportType794            = 794  // INT
	AllocIntermedReqType808       = 808  // INT
	MatchStatus573                = 573  // CHAR
	Product460                    = 460  // INT
	SecurityType167               = 167  // STRING
	Text58                        = 58   // STRING
	EncodedTextLen354             = 354  // LENGTH
	EncodedText355                = 355  // DATA
	ClearingBusinessDate715       = 715  // LOCALMKTDATE
	AvgPxIndicator819             = 819  // INT
	Quantity53                    = 53   // QTY
	AllocTransType71              = 71   // CHAR
	PartyID448                    = 448  // STRING
	PartyIDSource447              = 447  // CHAR
	PartyRole452                  = 452  // INT
	PartySubID523                 = 523  // STRING
	PartySubIDType803             = 803  // INT
	AllocAccount79                = 79   // STRING
	AllocAcctIDSource661          = 661  // INT
	AllocPrice366                 = 366  // PRICE
	IndividualAllocID467          = 467  // STRING
	IndividualAllocRejCode776     = 776  // INT
	AllocText161                  = 161  // STRING
	EncodedAllocTextLen360        = 360  // LENGTH
	EncodedAllocText361           = 361  // DATA
	SecondaryIndividualAllocID989 = 989  // STRING
	AllocCustomerCapacity993      = 993  // STRING
	IndividualAllocType992        = 992  // INT
	AllocQty80                    = 80   // QTY
	AllocPositionEffect1047       = 1047 // CHAR
	NestedPartyID524              = 524  // STRING
	NestedPartyIDSource525        = 525  // CHAR
	NestedPartyRole538            = 538  // INT
	NestedPartySubID545           = 545  // STRING
	NestedPartySubIDType805       = 805  // INT
)