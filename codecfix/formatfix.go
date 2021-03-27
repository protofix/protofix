// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codecfix

import (
	"errors"
	"strings"
)

// Format maps a codes to a each individual field of the FIX.
type Format struct {
	Head        map[int]Codec
	BodyLength9 BodyLengthFld
	Body        map[int]Codec
	Checksum10  ChecksumStringFld
	Unknown     UnknownFld
}

func (f Format) Validate() error {
	var errs []string

	if f.Head == nil {
		errs = append(errs, "missing header")
	}
	if f.Body == nil {
		errs = append(errs, "missing body")
	}

	var err error

	if errs != nil {
		err = errors.New(strings.Join(errs, ", "))
	}
	return err
}

const Req, Opt = true, false

// Field types.
const (
	BeginString8           = 8
	BodyLength9            = 9
	MsgType35              = 35
	SenderCompID49         = 49
	TargetCompID56         = 56
	ApplVerID1128          = 1128
	MsgSeqNum34            = 34
	PossDupFlag43          = 43
	PossResend97           = 97
	SendingTime52          = 52
	OrigSendingTime122     = 122
	EncryptMethod98        = 98
	HeartBtInt108          = 108
	ResetSeqNumFlag141     = 141
	Password554            = 554
	NewPassword925         = 925
	SessionStatus1409      = 1409
	CancelOnDisconnect6867 = 6867
	LanguageID6936         = 6936
	Checksum10             = 10
)

var FldText = map[int]string{
	BeginString8:           "BeginString",
	BodyLength9:            "BodyLength",
	MsgType35:              "MsgType",
	SenderCompID49:         "SenderCompID",
	TargetCompID56:         "TargetCompID",
	ApplVerID1128:          "ApplVerID",
	MsgSeqNum34:            "MsgSeqNum",
	PossDupFlag43:          "PossDupFlag",
	PossResend97:           "PossResend",
	SendingTime52:          "SendingTime",
	OrigSendingTime122:     "OrigSendingTime",
	EncryptMethod98:        "EncryptMethod",
	HeartBtInt108:          "HeartBtInt",
	ResetSeqNumFlag141:     "ResetSeqNumFlag",
	Password554:            "Password",
	NewPassword925:         "NewPassword",
	SessionStatus1409:      "SessionStatus",
	CancelOnDisconnect6867: "CancelOnDisconnect",
	LanguageID6936:         "LanguageID",
	Checksum10:             "Checksum",
}

// Message types.
const (
	Heartbeat0                          = "0"
	TestRequest1                        = "1"
	ResendRequest2                      = "2"
	Reject3                             = "3"
	SequenceReset4                      = "4"
	Logout5                             = "5"
	IndicationOfInterest6               = "6"
	Advertisement7                      = "7"
	ExecutionReport8                    = "8"
	OrderCancelReject9                  = "9"
	LogonA                              = "A"
	NewsB                               = "B"
	EmailC                              = "C"
	NewOrderSingleD                     = "D"
	NewOrderListE                       = "E"
	OrderCancelRequestF                 = "F"
	OrderCancelReplaceRequestG          = "G"
	OrderStatusRequestH                 = "H"
	AllocationJ                         = "J"
	ListCancelRequestK                  = "K"
	ListExecuteL                        = "L"
	ListStatusRequestM                  = "M"
	ListStatusN                         = "N"
	AllocationACKP                      = "P"
	DontKnowTradeQ                      = "Q"
	QuoteRequestR                       = "R"
	QuoteS                              = "S"
	SettlementInstructionsT             = "T"
	MarketDataRequestV                  = "V"
	MarketDataSnapshotFullRefreshW      = "W"
	MarketDataIncrementalRefreshX       = "X"
	MarketDataRequestRejectY            = "Y"
	QuoteCancelZ                        = "Z"
	QuoteStatusRequesta                 = "a"
	QuoteAcknowledgementb               = "b"
	SecurityDefinitionRequestc          = "c"
	SecurityDefinitiond                 = "d"
	SecurityStatusRequeste              = "e"
	SecurityStatusf                     = "f"
	TradingSessionStatusRequestg        = "g"
	TradingSessionStatush               = "h"
	MassQuotei                          = "i"
	BusinessMessageRejectj              = "j"
	BidRequestk                         = "k"
	BidResponsel                        = "l"
	ListStrikePricem                    = "m"
	XMLMessagen                         = "n"
	RegistrationInstructionso           = "o"
	RegistrationInstructionsResponsep   = "p"
	OrderMassCancelRequestq             = "q"
	OrderMassCancelReportr              = "r"
	NewOrderCrosss                      = "s"
	CrossOrderCancelReplaceRequestt     = "t"
	CrossOrderCancelRequestu            = "u"
	SecurityTypeRequestv                = "v"
	SecurityTypesw                      = "w"
	SecurityListRequestx                = "x"
	SecurityListy                       = "y"
	DerivativeSecurityListRequestz      = "z"
	DerivativeSecurityListAA            = "AA"
	NewOrderMultilegAB                  = "AB"
	MultilegOrderCancelReplaceRequestAC = "AC"
	TradeCaptureReportRequestAD         = "AD"
	TradeCaptureReportAE                = "AE"
	OrderMassStatusRequestAF            = "AF"
	QuoteRequestRejectAG                = "AG"
	RFQRequestAH                        = "AH"
	QuoteStatusReportAI                 = "AI"
	QuoteResponseAJ                     = "AJ"
	ConfirmationAK                      = "AK"
	PositionMaintenanceRequestAL        = "AL"
	PositionMaintenanceReportAM         = "AM"
	RequestForPositionsAN               = "AN"
	RequestForPositionsAckAO            = "AO"
	PositionReportAP                    = "AP"
	TradeCaptureReportRequestAckAQ      = "AQ"
	TradeCaptureReportAckAR             = "AR"
	AllocationReportAS                  = "AS"
	AllocationReportAckAT               = "AT"
	ConfirmationAckAU                   = "AU"
	SettlementInstructionRequestAV      = "AV"
	AssignmentReportAW                  = "AW"
	CollateralRequestAX                 = "AX"
	CollateralAssignmentAY              = "AY"
	CollateralResponseAZ                = "AZ"
	CollateralReportBA                  = "BA"
	CollateralInquiryBB                 = "BB"
	NetworkStatusRequestBC              = "BC"
	NetworkStatusResponseBD             = "BD"
	UserRequestBE                       = "BE"
	UserResponseBF                      = "BF"
	CollateralInquiryAckBG              = "BG"
	ConfirmationRequestBH               = "BH"
)

var MsgTypes = []string{
	Heartbeat0,
	TestRequest1,
	ResendRequest2,
	Reject3,
	SequenceReset4,
	Logout5,
	IndicationOfInterest6,
	Advertisement7,
	ExecutionReport8,
	OrderCancelReject9,
	LogonA,
	NewsB,
	EmailC,
	NewOrderSingleD,
	NewOrderListE,
	OrderCancelRequestF,
	OrderCancelReplaceRequestG,
	OrderStatusRequestH,
	AllocationJ,
	ListCancelRequestK,
	ListExecuteL,
	ListStatusRequestM,
	ListStatusN,
	AllocationACKP,
	DontKnowTradeQ,
	QuoteRequestR,
	QuoteS,
	SettlementInstructionsT,
	MarketDataRequestV,
	MarketDataSnapshotFullRefreshW,
	MarketDataIncrementalRefreshX,
	MarketDataRequestRejectY,
	QuoteCancelZ,
	QuoteStatusRequesta,
	QuoteAcknowledgementb,
	SecurityDefinitionRequestc,
	SecurityDefinitiond,
	SecurityStatusRequeste,
	SecurityStatusf,
	TradingSessionStatusRequestg,
	TradingSessionStatush,
	MassQuotei,
	BusinessMessageRejectj,
	BidRequestk,
	BidResponsel,
	ListStrikePricem,
	XMLMessagen,
	RegistrationInstructionso,
	RegistrationInstructionsResponsep,
	OrderMassCancelRequestq,
	OrderMassCancelReportr,
	NewOrderCrosss,
	CrossOrderCancelReplaceRequestt,
	CrossOrderCancelRequestu,
	SecurityTypeRequestv,
	SecurityTypesw,
	SecurityListRequestx,
	SecurityListy,
	DerivativeSecurityListRequestz,
	DerivativeSecurityListAA,
	NewOrderMultilegAB,
	MultilegOrderCancelReplaceRequestAC,
	TradeCaptureReportRequestAD,
	TradeCaptureReportAE,
	OrderMassStatusRequestAF,
	QuoteRequestRejectAG,
	RFQRequestAH,
	QuoteStatusReportAI,
	QuoteResponseAJ,
	ConfirmationAK,
	PositionMaintenanceRequestAL,
	PositionMaintenanceReportAM,
	RequestForPositionsAN,
	RequestForPositionsAckAO,
	PositionReportAP,
	TradeCaptureReportRequestAckAQ,
	TradeCaptureReportAckAR,
	AllocationReportAS,
	AllocationReportAckAT,
	ConfirmationAckAU,
	SettlementInstructionRequestAV,
	AssignmentReportAW,
	CollateralRequestAX,
	CollateralAssignmentAY,
	CollateralResponseAZ,
	CollateralReportBA,
	CollateralInquiryBB,
	NetworkStatusRequestBC,
	NetworkStatusResponseBD,
	UserRequestBE,
	UserResponseBF,
	CollateralInquiryAckBG,
	ConfirmationRequestBH,
}
