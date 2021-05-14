// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1registrationinstructionsresponse is a format of the FIX.5.0 servicepack 1 RegistrationInstructionsResponse message.
package fix50sp1registrationinstructionsresponse

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1RegistrationInstructionsResponseMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1RegistrationInstructionsResponse}
	FIX50SP1RegistrationInstructionsResponseUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1RegistrationInstructionsResponse}
)

// FIX50SP1RegistrationInstructionsResponse is a FIX.5.0 servicepack 1 format of the RegistrationInstructionsResponse message which maps the codecs into individual fields.
var FIX50SP1RegistrationInstructionsResponse = f0.Format{
	Fields: map[int]f0.Codec{
		RegistID513:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RegistTransType514:     f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* REPLACE */, "2" /* CANCEL */), f0.Con{1}},
		RegistRefID508:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClOrdID11:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Account1:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AcctIDSource660:        f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BIC */, 2 /* SID_CODE */, 3 /* TFM */, 4 /* OMGEO */, 5 /* DTCC_CODE */, 99 /* OTHER */), f0.Var{1, 19}},
		RegistStatus506:        f0.Fld{Req, f0.ASCII, f0.String("A" /* ACCEPTED */, "H" /* HELD */, "N" /* REMINDER */, "R" /* REJECTED */), f0.Con{1}},
		RegistRejReasonCode507: f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* INVALID_UNACCEPTABLE_ACCOUNT_TYPE */, 10 /* INVALID_UNACEEPTABLE_INVESTOR_ID_SOURCE */, 11 /* INVALID_UNACCEPTABLE_DATE_OF_BIRTH */, 12 /* INVALID_UNACCEPTABLE_INVESTOR_COUNTRY_OF_RESIDENCE */, 13 /* INVALID_UNACCEPTABLE_NO_DISTRIB_INSTNS */, 14 /* INVALID_UNACCEPTABLE_DISTRIB_PERCENTAGE */, 15 /* INVALID_UNACCEPTABLE_DISTRIB_PAYMENT_METHOD */, 16 /* INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NAME */, 17 /* INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_CODE */, 18 /* INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NUM */, 2 /* INVALID_UNACCEPTABLE_TAX_EXEMPT_TYPE */, 3 /* INVALID_UNACCEPTABLE_OWNERSHIP_TYPE */, 4 /* INVALID_UNACCEPTABLE_NO_REG_DETAILS */, 5 /* INVALID_UNACCEPTABLE_REG_SEQ_NO */, 6 /* INVALID_UNACCEPTABLE_REG_DETAILS */, 7 /* INVALID_UNACCEPTABLE_MAILING_DETAILS */, 8 /* INVALID_UNACCEPTABLE_MAILING_INSTRUCTIONS */, 9 /* INVALID_UNACCEPTABLE_INVESTOR_ID */, 99 /* OTHER */), f0.Var{1, 19}},
		RegistRejReasonText496: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyID448:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:       f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREAN_INVESTOR_ID */, "2" /* TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID */, "3" /* TAIWANESE_TRADING_ACCT */, "4" /* MALAYSIAN_CENTRAL_DEPOSITORY */, "5" /* CHINESE_INVESTOR_ID */, "6" /* UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER */, "7" /* US_SOCIAL_SECURITY_NUMBER */, "8" /* US_EMPLOYER_OR_TAX_ID_NUMBER */, "9" /* AUSTRALIAN_BUSINESS_NUMBER */, "A" /* AUSTRALIAN_TAX_FILE_NUMBER */, "B" /* BIC */, "C" /* GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER */, "D" /* PROPRIETARY */, "E" /* ISO_COUNTRY_CODE */, "F" /* SETTLEMENT_ENTITY_LOCATION */, "G" /* MIC */, "H" /* CSD_PARTICIPANT_MEMBER_CODE */, "I" /* DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT */), f0.Con{1}},
		PartyRole452:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTING_FIRM */, 10 /* SETTLEMENT_LOCATION */, 11 /* ORDER_ORIGINATION_TRADER */, 12 /* EXECUTING_TRADER */, 13 /* ORDER_ORIGINATION_FIRM */, 14 /* GIVEUP_CLEARING_FIRM */, 15 /* CORRESPONDANT_CLEARING_FIRM */, 16 /* EXECUTING_SYSTEM */, 17 /* CONTRA_FIRM */, 18 /* CONTRA_CLEARING_FIRM */, 19 /* SPONSORING_FIRM */, 2 /* BROKER_OF_CREDIT */, 20 /* UNDERLYING_CONTRA_FIRM */, 21 /* CLEARING_ORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMER_ACCOUNT */, 25 /* CORRESPONDENT_CLEARING_ORGANIZATION */, 26 /* CORRESPONDENT_BROKER */, 27 /* BUYER_SELLER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENT_ID */, 30 /* AGENT */, 31 /* SUB_CUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTED_PARTY */, 34 /* REGULATORY_BODY */, 35 /* LIQUIDITY_PROVIDER */, 36 /* ENTERING_TRADER */, 37 /* CONTRA_TRADER */, 38 /* POSITION_ACCOUNT */, 4 /* CLEARING_FIRM */, 5 /* INVESTOR_ID */, 6 /* INTRODUCING_FIRM */, 7 /* ENTERING_FIRM */, 8 /* LOCATE */, 9 /* FUND_MANAGER_CLIENT_ID */, 60 /* INTRODUCING_BROKER */, 41 /* CONTRA_POSITION_ACCOUNT */, 42 /* CONTRA_EXCHANGE */, 43 /* INTERNAL_CARRY_ACCOUNT */, 44 /* ORDER_ENTRY_OPERATOR_ID */, 45 /* SECONDARY_ACCOUNT_NUMBER */, 46 /* FORIEGN_FIRM */, 47 /* THIRD_PARTY_ALLOCATION_FIRM */, 48 /* CLAIMING_ACCOUNT */, 49 /* ASSET_MANAGER */, 50 /* PLEDGOR_ACCOUNT */, 51 /* PLEDGEE_ACCOUNT */, 52 /* LARGE_TRADER_REPORTABLE_ACCOUNT */, 53 /* TRADER_MNEMONIC */, 54 /* SENDER_LOCATION */, 55 /* SESSION_ID */, 56 /* ACCEPTABLE_COUNTERPARTY */, 57 /* UNACCEPTABLE_COUNTERPARTY */, 58 /* ENTERING_UNIT */, 59 /* EXECUTING_UNIT */, 39 /* CONTRA_INVESTOR_ID */, 40 /* TRANSFER_TO_FIRM */, 61 /* QUOTE_ORIGINATOR */, 62 /* REPORT_ORIGINATOR */, 63 /* SYSTEMATIC_INTERNALISER */, 64 /* MULTILATERAL_TRADING_FACILITY */, 65 /* REGULATED_MARKET */, 66 /* MARKET_MAKER */, 67 /* INVESTMENT_FIRM */, 68 /* HOST_COMPETENT_AUTHORITY */, 69 /* HOME_COMPETENT_AUTHORITY */, 70 /* COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY */, 71 /* COMPETENT_AUTHORITY_OF_THE_TRANSACTION */, 72 /* REPORTING_INTERMEDIARY */, 73 /* EXECUTION_VENUE */, 74 /* MARKET_DATA_ENTRY_ORIGINATOR */, 75 /* LOCATION_ID */, 76 /* DESK_ID */, 77 /* MARKET_DATA_MARKET */, 78 /* ALLOCATION_ENTITY */, 79 /* PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES */, 80 /* STEP_OUT_FIRM */, 81 /* BROKERCLEARINGID */), f0.Var{1, 19}},
		PartySubID523:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIES_ACCOUNT_NUMBER */, 11 /* REGISTRATION_NUMBER */, 12 /* REGISTERED_ADDRESS_12 */, 13 /* REGULATORY_STATUS */, 14 /* REGISTRATION_NAME */, 15 /* CASH_ACCOUNT_NUMBER */, 16 /* BIC */, 17 /* CSD_PARTICIPANT_MEMBER_CODE */, 18 /* REGISTERED_ADDRESS_18 */, 19 /* FUND_ACCOUNT_NAME */, 2 /* PERSON */, 20 /* TELEX_NUMBER */, 21 /* FAX_NUMBER */, 22 /* SECURITIES_ACCOUNT_NAME */, 23 /* CASH_ACCOUNT_NAME */, 24 /* DEPARTMENT */, 25 /* LOCATION_DESK */, 26 /* POSITION_ACCOUNT_TYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 5 /* FULL_LEGAL_NAME_OF_FIRM */, 6 /* POSTAL_ADDRESS */, 7 /* PHONE_NUMBER */, 8 /* EMAIL_ADDRESS */, 9 /* CONTACT_NAME */, 27 /* SECURITY_LOCATE_ID */, 28 /* MARKET_MAKER */, 29 /* ELIGIBLE_COUNTERPARTY */, 30 /* PROFESSIONAL_CLIENT */, 31 /* LOCATION */, 32 /* EXECUTION_VENUE */, 33 /* CURRENCY_DELIVERY_IDENTIFIER */), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		RegistID513,            // STRING
		RegistTransType514,     // CHAR
		RegistRefID508,         // STRING
		ClOrdID11,              // STRING
		Account1,               // STRING
		AcctIDSource660,        // INT
		RegistStatus506,        // CHAR
		RegistRejReasonCode507, // INT
		RegistRejReasonText496, // STRING
		PartyID448,             // STRING
		PartyIDSource447,       // CHAR
		PartyRole452,           // INT
		PartySubID523,          // STRING
		PartySubIDType803,      // INT
	},
}

const Req, Opt = true, false

const (
	RegistID513            = 513 // STRING
	RegistTransType514     = 514 // CHAR
	RegistRefID508         = 508 // STRING
	ClOrdID11              = 11  // STRING
	Account1               = 1   // STRING
	AcctIDSource660        = 660 // INT
	RegistStatus506        = 506 // CHAR
	RegistRejReasonCode507 = 507 // INT
	RegistRejReasonText496 = 496 // STRING
	PartyID448             = 448 // STRING
	PartyIDSource447       = 447 // CHAR
	PartyRole452           = 452 // INT
	PartySubID523          = 523 // STRING
	PartySubIDType803      = 803 // INT
)
