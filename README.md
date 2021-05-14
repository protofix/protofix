# protofix

[![Build Status](https://cloud.drone.io/api/badges/protofix/protofix/status.svg)](https://cloud.drone.io/protofix/protofix)
[![Go Reference](https://pkg.go.dev/badge/github.com/protofix/protofix.svg)](https://pkg.go.dev/github.com/protofix/protofix)

FIX protocol codec for Go.  
Source files are distributed under the BSD-style license
found in the [LICENSE](./LICENSE) file.

## About

The software is considered to be at an alpha level of readiness -
its incomplete and extremely slow and allocates a lots of memory)

## Install

    go get github.com/protofix/protofix@v0.0.52

## Usage

```go
package main

import (
    "bytes"
    "time"

    "github.com/protofix/protofix/fix44/fix44logon"
)

func main() {
    type Logon struct {
        BeginString8       string    `FIX44:"8"`
        BodyLength9        int       `FIX44:"9"`
        MsgType35          string    `FIX44:"35"`
        SenderCompID49     string    `FIX44:"49"`
        TargetCompID56     string    `FIX44:"56"`
        MsgSeqNum34        int       `FIX44:"34"`
        SendingTime52      time.Time `FIX44:"52"`
        EncryptMethod98    int       `FIX44:"98"`
        HeartBtInt108      int       `FIX44:"108"`
        ResetSeqNumFlag141 bool      `FIX44:"141"`
        Username553        string    `FIX44:"553"`
        Password554        string    `FIX44:"554"`
        CheckSum10         string    `FIX44:"10"`
    }

    input := []byte("8=FIX.4.4|9=102|35=A|49=BuySide|56=SellSide|34=1|52=20190605-11:40:30.392|98=0|108=30|141=Y|553=Username|554=Password|10=104|")
    input = bytes.ReplaceAll(input, []byte{'|'}, []byte{0x01})

    logon := Logon{}

    _, _, err := fix44logon.FIX44LogonUnmarshaler.Unmarshal(input, &logon)
    if err != nil {
        fmt.Printf("unexpected unmarshal error: %s", err)
    }

    output, _, _, err := fix44logon.FIX44LogonMarshaler.Marshal(&logon)
    if err != nil {
        fmt.Printf("unexpected marshal error: %s", err)
    }

    input = bytes.ReplaceAll(input, []byte{0x01}, []byte{'|'})
    output = bytes.ReplaceAll(output, []byte{0x01}, []byte{'|'})

    if bytes.Equal(input, output) {
        fmt.Printf("Messages %q are equal %q.\n", input, output)
    } else {
        fmt.Printf("Messages %q are not equal %q!\n", input, output)
    }
```

Output:

```
Messages equal
```

See this example in [fix44_logon_readme_test.go][]

[fix44_logon_readme_test.go]: ./fix44/fix44logon/fix44_logon_readme_test.go#L15

## Benchmark

```
go test -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/protofix/protofix/fix44/fix44logon
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkLogonUnmarshalMarshal-8          46309         25296 ns/op
```
