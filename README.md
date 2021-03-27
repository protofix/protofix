protofix
========

[![Build Status](https://cloud.drone.io/api/badges/danil/protofix/status.svg)](https://cloud.drone.io/danil/protofix)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/protofix.svg)](https://pkg.go.dev/github.com/danil/protofix)

FIX.4.4 format and FIX protocol codec for Go.  
Source files are distributed under the BSD-style license
found in the [LICENSE](./LICENSE) file.

About
-----

The software is considered to be at an alpha level of readiness -
its incomplete and extremely slow and allocates a lots of memory)

Install
-------

    go get github.com/danil/protofix@v0.0.41

Usage
-----

```go
package main

import (
    "fmt"

    "github.com/danil/protofix/moex44"
)

type MOEX44Logon struct {
    BeginString8           string    `MOEX44:"8"`
    BodyLength9            int       `MOEX44:"9"`
    MsgType35              string    `MOEX44:"35"`
    SenderCompID49         string    `MOEX44:"49"`
    TargetCompID56         string    `MOEX44:"56"`
    MsgSeqNum34            int       `MOEX44:"34"`
    PossDupFlag43          bool      `MOEX44:"43"`
    PossResend97           bool      `MOEX44:"97"`
    SendingTime52          time.Time `MOEX44:"52"`
    OrigSendingTime122     time.Time `MOEX44:"122"`
    EncryptMethod98        int       `MOEX44:"98"`
    HeartBtInt108          int       `MOEX44:"108"`
    ResetSeqNumFlag141     bool      `MOEX44:"141"`
    Password554            string    `MOEX44:"554"`
    NewPassword925         string    `MOEX44:"925"`
    SessionStatus1409      int       `MOEX44:"1409"`
    CancelOnDisconnect6867 string    `MOEX44:"6867"`
    LanguageID6936         string    `MOEX44:"6936"`
    Checksum10             string    `MOEX44:"10"`
}

func main() {
    input := MOEX44Logon{
        MsgType35:       "A",
        SenderCompID49:  "Foo",
        TargetCompID56:  "Bar",
        MsgSeqNum34:     1,
        SendingTime52:   time.Date(2021, time.March, 12, 12, 35, 12, 0, time.UTC),
        EncryptMethod98: 0,
        HeartBtInt108:   42,
        Password554:     "12345678",
    }
    p, _ := moex44.MOEX44LogonMarshaler.Marshal(&input)
    fmt.Println(string(bytes.ReplaceAll(p, []byte{0x01}, []byte{'|'})))
}
```

Output:

```
8=FIX.4.4|9=103|35=A|34=1|43=N|49=Foo|52=20210312-12:35:12.000000000|56=Bar|97=N|98=0|108=42|141=N|554=12345678|1409=0|10=060|
```

Benchmark
---------

```
go test -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/danil/protofix/moex44
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkMarshal/marshal_many_fields_275-8         	  103887	     11953 ns/op
```
