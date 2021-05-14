// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1userresponse is a format of the FIX.5.0 servicepack 1 UserResponse message.
package fix50sp1userresponse

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1UserResponseMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1UserResponse}
	FIX50SP1UserResponseUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1UserResponse}
)

// FIX50SP1UserResponse is a FIX.5.0 servicepack 1 format of the UserResponse message which maps the codecs into individual fields.
var FIX50SP1UserResponse = f0.Format{
	Fields: map[int]f0.Codec{
		UserRequestID923:  f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Username553:       f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UserStatus926:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* LOGGED_IN */, 2 /* NOT_LOGGED_IN */, 3 /* USER_NOT_RECOGNISED */, 4 /* PASSWORD_INCORRECT */, 5 /* PASSWORD_CHANGED */, 6 /* OTHER */, 7 /* FORCED_USER_LOGOUT_BY_EXCHANGE */, 8 /* SESSION_SHUTDOWN_WARNING */), f0.Var{1, 19}},
		UserStatusText927: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		UserRequestID923,  // STRING
		Username553,       // STRING
		UserStatus926,     // INT
		UserStatusText927, // STRING
	},
}

const Req, Opt = true, false

const (
	UserRequestID923  = 923 // STRING
	Username553       = 553 // STRING
	UserStatus926     = 926 // INT
	UserStatusText927 = 927 // STRING
)