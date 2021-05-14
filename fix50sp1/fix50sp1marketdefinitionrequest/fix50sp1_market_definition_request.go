// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1marketdefinitionrequest is a format of the FIX.5.0 servicepack 1 MarketDefinitionRequest message.
package fix50sp1marketdefinitionrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1MarketDefinitionRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1MarketDefinitionRequest}
	FIX50SP1MarketDefinitionRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1MarketDefinitionRequest}
)

// FIX50SP1MarketDefinitionRequest is a FIX.5.0 servicepack 1 format of the MarketDefinitionRequest message which maps the codecs into individual fields.
var FIX50SP1MarketDefinitionRequest = f0.Format{
	Fields: map[int]f0.Codec{
		MarketReqID1393:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SubscriptionRequestType263: f0.Fld{Req, f0.ASCII, f0.String("0" /* SNAPSHOT */, "1" /* SNAPSHOT_PLUS_UPDATES */, "2" /* DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST */), f0.Con{1}},
		MarketID1301:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MarketSegmentID1300:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ParentMktSegmID1325:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		MarketReqID1393,            // STRING
		SubscriptionRequestType263, // CHAR
		MarketID1301,               // EXCHANGE
		MarketSegmentID1300,        // STRING
		ParentMktSegmID1325,        // STRING
	},
}

const Req, Opt = true, false

const (
	MarketReqID1393            = 1393 // STRING
	SubscriptionRequestType263 = 263  // CHAR
	MarketID1301               = 1301 // EXCHANGE
	MarketSegmentID1300        = 1300 // STRING
	ParentMktSegmID1325        = 1325 // STRING
)
