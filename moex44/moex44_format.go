// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package moex44 provides FIX.4.4 (MOEX.4.4) protocol format.
package moex44

import (
	"time"

	f0 "github.com/danil/protofix/codecfix"
	"github.com/danil/protofix/marshfix"
)

var (
	MOEX44LogonMarshaler   = marshfix.Marshal{Tag: "MOEX44"}
	MOEX44LogonUnmarshaler = marshfix.Unmarshal{Tag: "MOEX44"}
)

// MOEX44 is a FIX.4.4 (MOEX.4.4) format which maps the codecs into individual fields.
var (
	// MOEX44Head is a FIX.4.4 (MOEX.4.4) format of the head shuld contain each message
	// The head should contain message type but each message have individual type.
	// Also trail/tail shuld contain each message.
	MOEX44Base = f0.Format{
		Head: map[int]f0.Codec{
			f0.BeginString8:   f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
			f0.SenderCompID49: f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 12}},
			f0.TargetCompID56: f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, f0.MaxString}},
		},
		BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, f0.MaxInt}},
		Body: map[int]f0.Codec{
			f0.MsgSeqNum34:        f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, f0.MaxInt}},
			f0.PossDupFlag43:      f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
			f0.PossResend97:       f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
			f0.SendingTime52:      f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Con{27}},
			f0.OrigSendingTime122: f0.Fld{Opt, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Con{27}},
		},
		Checksum10: f0.ChecksumStringFld{f0.ASCII, f0.ChecksumString(), f0.Con{3}},
		Unknown:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, f0.MaxBytes}},
	}

	// MOEX44Logon is a FIX.4.4 (MOEX.4.4) format of the logon message.
	MOEX44Logon = f0.Format{
		Head: map[int]f0.Codec{
			f0.MsgType35: f0.Fld{Req, f0.ASCII, f0.StringDefault(f0.LogonA), f0.Var{1, 2}},
		},
		Body: map[int]f0.Codec{
			f0.EncryptMethod98:        f0.Fld{Req, f0.ASCII, f0.IntDefault(0), f0.Con{1}},
			f0.HeartBtInt108:          f0.Fld{Req, f0.ASCII, f0.SecondsDuration(HeartBtIntCstr...), f0.Var{1, f0.MaxInt}},
			f0.ResetSeqNumFlag141:     f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
			f0.Password554:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 8}},
			f0.NewPassword925:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 8}},
			f0.SessionStatus1409:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(0, 3), f0.Con{1}},
			f0.CancelOnDisconnect6867: f0.Fld{Opt, f0.ASCII, f0.String("A"), f0.Con{1}},
			f0.LanguageID6936:         f0.Fld{Opt, f0.ASCII, f0.String("R", "E"), f0.Con{1}},
		},
	}
)

const Req, Opt = true, false

func init() {
	// MOEX44Logon initialization.
	for k, v := range MOEX44Base.Head {
		MOEX44Logon.Head[k] = v
	}
	MOEX44Logon.BodyLength9 = MOEX44Base.BodyLength9
	for k, v := range MOEX44Base.Body {
		MOEX44Logon.Body[k] = v
	}
	MOEX44Logon.Checksum10 = MOEX44Base.Checksum10
	MOEX44Logon.Unknown = MOEX44Base.Unknown
	MOEX44LogonMarshaler.Format = MOEX44Logon
	MOEX44LogonUnmarshaler.Format = MOEX44Logon
}

var HeartBtIntCstr = func() []time.Duration {
	d := make([]time.Duration, 60)
	for i := 0; i < 60; i++ {
		d[i] = time.Duration(i+1) * time.Second
	}
	return d
}()
