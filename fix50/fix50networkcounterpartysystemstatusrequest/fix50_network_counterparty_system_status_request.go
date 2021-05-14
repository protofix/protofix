// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50networkcounterpartysystemstatusrequest is a format of the FIX.5.0 NetworkCounterpartySystemStatusRequest message.
package fix50networkcounterpartysystemstatusrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50NetworkCounterpartySystemStatusRequestMarshaler   = marshfix.Marshal{Tag: "FIX50", Format: FIX50NetworkCounterpartySystemStatusRequest}
	FIX50NetworkCounterpartySystemStatusRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50", Format: FIX50NetworkCounterpartySystemStatusRequest}
)

// FIX50NetworkCounterpartySystemStatusRequest is a FIX.5.0 format of the NetworkCounterpartySystemStatusRequest message which maps the codecs into individual fields.
var FIX50NetworkCounterpartySystemStatusRequest = f0.Format{
	Fields: map[int]f0.Codec{
		NetworkRequestType935: f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* SNAPSHOT */, 2 /* SUBSCRIBE */, 4 /* STOP_SUBSCRIBING */, 8 /* LEVEL_OF_DETAIL_THEN_NOCOMPIDS_BECOMES_REQUIRED */), f0.Var{1, 19}},
		NetworkRequestID933:   f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RefCompID930:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RefSubID931:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LocationID283:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeskID284:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		NetworkRequestType935, // INT
		NetworkRequestID933,   // STRING
		RefCompID930,          // STRING
		RefSubID931,           // STRING
		LocationID283,         // STRING
		DeskID284,             // STRING
	},
}

const Req, Opt = true, false

const (
	NetworkRequestType935 = 935 // INT
	NetworkRequestID933   = 933 // STRING
	RefCompID930          = 930 // STRING
	RefSubID931           = 931 // STRING
	LocationID283         = 283 // STRING
	DeskID284             = 284 // STRING
)