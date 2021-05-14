// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp2streamassignmentreportack is a format of the FIX.5.0 servicepack 2 StreamAssignmentReportACK message.
package fix50sp2streamassignmentreportack

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP2StreamAssignmentReportACKMarshaler   = marshfix.Marshal{Tag: "FIX50SP2", Format: FIX50SP2StreamAssignmentReportACK}
	FIX50SP2StreamAssignmentReportACKUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP2", Format: FIX50SP2StreamAssignmentReportACK}
)

// FIX50SP2StreamAssignmentReportACK is a FIX.5.0 servicepack 2 format of the StreamAssignmentReportACK message which maps the codecs into individual fields.
var FIX50SP2StreamAssignmentReportACK = f0.Format{
	Fields: map[int]f0.Codec{
		StreamAsgnAckType1503:   f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* ASSIGNMENT_ACCEPTED */, 1 /* ASSIGNMENT_REJECTED */), f0.Var{1, 19}},
		StreamAsgnRptID1501:     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		StreamAsgnRejReason1502: f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* UNKNOWN_CLIENT */, 1 /* EXCEEDS_MAXIMUM_SIZE */, 2 /* UNKNOWN_OR_INVALID_CURRENCY_PAIR */, 3 /* NO_AVAILABLE_STREAM */, 99 /* OTHER */), f0.Var{1, 19}},
		Text58:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:          f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		StreamAsgnAckType1503,   // INT
		StreamAsgnRptID1501,     // STRING
		StreamAsgnRejReason1502, // INT
		Text58,                  // STRING
		EncodedTextLen354,       // LENGTH
		EncodedText355,          // DATA
	},
}

const Req, Opt = true, false

const (
	StreamAsgnAckType1503   = 1503 // INT
	StreamAsgnRptID1501     = 1501 // STRING
	StreamAsgnRejReason1502 = 1502 // INT
	Text58                  = 58   // STRING
	EncodedTextLen354       = 354  // LENGTH
	EncodedText355          = 355  // DATA
)
