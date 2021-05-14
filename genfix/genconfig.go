// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a FIX message format generator.
package genfix

import (
	"fmt"
	"strconv"
	"time"

	f0 "github.com/protofix/protofix/codecfix"
)

// FIX40 is a FIX.4.0 configuration.
func FIX40() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX41 is a FIX.4.1 configuration.
func FIX41() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX42 is a FIX.4.2 configuration.
func FIX42() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX43 is a FIX.4.3 configuration.
func FIX43() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX44 is a FIX.4.4 configuration.
func FIX44() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX50 is a FIX.5.0 configuration.
func FIX50() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX50SP1 is a FIX.5.0 service pack 1 configuration.
func FIX50SP1() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIX50SP2 is a FIX.5.0 service pack 2 configuration.
func FIX50SP2() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// FIXT11 is a FIXT.1.0 (transport) configuration.
func FIXT11() []Option {
	c := Config{
		Serialize: map[string]Serialize{
			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"UTCTIMESTAMP":        Serialize{Func: "UTCTimestampMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			f0.BeginString8: LengthTag{Min: 8, Max: 8},

			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

// MOEX44 is a FIX.4.4 (MOEX.4.4)  configuration.
func MOEX44() []Option {
	const (
		SessionStatus1409      = 1409
		CancelOnDisconnect6867 = 6867
		LanguageID6936         = 6936
	)

	c := Config{
		Serialize: map[string]Serialize{
			"UTCTIMESTAMP": Serialize{Func: "UTCTimestampNanosecondTime"},

			"AMT":                 Serialize{Func: "Float"},
			"BOOLEAN":             Serialize{Func: "BoolDefault"},
			"BYTES":               Serialize{Func: "Bytes"},
			"CHAR":                Serialize{Func: "String"},
			"CHECKSUM":            Serialize{Func: "CheckSum"},
			"COUNTRY":             Serialize{Func: "String"},
			"CURRENCY":            Serialize{Func: "String"},
			"DATA":                Serialize{Func: "Bytes"},
			"DATE":                Serialize{Func: "UTCDateOnlyTime"},
			"DAYOFMONTH":          Serialize{Func: "UTCDateOnlyTime"},
			"EXCHANGE":            Serialize{Func: "String"},
			"FLOAT":               Serialize{Func: "Float"},
			"INT":                 Serialize{Func: "IntDefault"},
			"LANGUAGE":            Serialize{Func: "String"},
			"LENGTH":              Serialize{Func: "Length"},
			"LOCALMKTDATE":        Serialize{Func: "LocalMktDateTime"},
			"MONTHYEAR":           Serialize{Func: "MonthYearTime"},
			"MULTIPLECHARVALUE":   Serialize{Func: "String"},
			"MULTIPLESTRINGVALUE": Serialize{Func: "String"},
			"MULTIPLEVALUESTRING": Serialize{Func: "String"},
			"NUMINGROUP":          Serialize{Func: "IntDefault"},
			"PERCENTAGE":          Serialize{Func: "Float"},
			"PRICE":               Serialize{Func: "Float"},
			"PRICEOFFSET":         Serialize{Func: "Float"},
			"QTY":                 Serialize{Func: "Float"},
			"SEQNUM":              Serialize{Func: "SeqNum"},
			"STRING":              Serialize{Func: "String"},
			"TIME":                Serialize{Func: "TZTime"},
			"TZTIMEONLY":          Serialize{Func: "TZTime"},
			"TZTIMESTAMP":         Serialize{Func: "TZTimestampMillisecondTime"},
			"UNKNOWN":             Serialize{Func: "Unknown"},
			"UTCDATE":             Serialize{Func: "UTCDateOnlyTime"},
			"UTCDATEONLY":         Serialize{Func: "UTCDateOnlyTime"},
			"UTCTIMEONLY":         Serialize{Func: "UTCTimeOnlyMillisecondTime"},
			"XMLDATA":             Serialize{Func: "Bytes"},
		},
		SerializeTag: map[int]SerializeTag{
			f0.BeginString8:  SerializeTag{Func: "StringDefault"},
			f0.BodyLength9:   SerializeTag{Func: "BodyLength"},
			f0.CheckSum10:    SerializeTag{Func: "CheckSum"},
			f0.HeartBtInt108: SerializeTag{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"AMT":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"BOOLEAN":             EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"CHAR":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CHECKSUM":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"COUNTRY":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"CURRENCY":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DATA":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"DATE":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"DAYOFMONTH":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"EXCHANGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"FLOAT":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"INT":                 EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LANGUAGE":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"LENGTH":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LOCALMKTDATE":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MONTHYEAR":           EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLECHARVALUE":   EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLESTRINGVALUE": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"MULTIPLEVALUESTRING": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"NUMINGROUP":          EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"PERCENTAGE":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICE":               EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"PRICEOFFSET":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"QTY":                 EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%v", s), nil }},
			"SEQNUM":              EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":              EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TIME":                EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMEONLY":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"TZTIMESTAMP":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
			"UTCDATE":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCDATEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMEONLY":         EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMP":        EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"XMLDATA":             EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("[]byte(%q)", s), nil }},
		},
		EnumFormatTag: map[int]EnumFormatTag{
			f0.HeartBtInt108: EnumFormatTag{Func: func(s string) (string, error) {
				t, err := time.ParseDuration(s + "s")
				return strconv.Itoa(int(t.Nanoseconds())), err
			}},
		},
		Length: map[string]Length{
			"AMT":                 Length{Min: 1, Max: 19},
			"BOOLEAN":             Length{Min: 1, Max: 1},
			"BYTES":               Length{Min: 1, Max: 65536},
			"CHAR":                Length{Min: 1, Max: 1},
			"CHECKSUM":            Length{Min: 3, Max: 3},
			"COUNTRY":             Length{Min: 2, Max: 2},
			"CURRENCY":            Length{Min: 3, Max: 3},
			"DATA":                Length{Min: 1, Max: 65536},
			"DATE":                Length{Min: 8, Max: 8},
			"DAYOFMONTH":          Length{Min: 8, Max: 8},
			"EXCHANGE":            Length{Min: 1, Max: 65536},
			"FLOAT":               Length{Min: 1, Max: 19},
			"INT":                 Length{Min: 1, Max: 19},
			"LANGUAGE":            Length{Min: 2, Max: 2},
			"LENGTH":              Length{Min: 1, Max: 19},
			"LOCALMKTDATE":        Length{Min: 8, Max: 8},
			"MONTHYEAR":           Length{Min: 6, Max: 8},
			"MULTIPLECHARVALUE":   Length{Min: 1, Max: 65536},
			"MULTIPLESTRINGVALUE": Length{Min: 1, Max: 65536},
			"MULTIPLEVALUESTRING": Length{Min: 1, Max: 65536},
			"NUMINGROUP":          Length{Min: 1, Max: 19},
			"PERCENTAGE":          Length{Min: 1, Max: 19},
			"PRICE":               Length{Min: 1, Max: 19},
			"PRICEOFFSET":         Length{Min: 1, Max: 19},
			"QTY":                 Length{Min: 1, Max: 19},
			"SEQNUM":              Length{Min: 1, Max: 19},
			"STRING":              Length{Min: 1, Max: 65536},
			"TIME":                Length{Min: 8, Max: 12},
			"TZTIMEONLY":          Length{Min: 8, Max: 12},
			"TZTIMESTAMP":         Length{Min: 1, Max: 27},
			"UNKNOWN":             Length{Min: 1, Max: 65536},
			"UTCDATE":             Length{Min: 8, Max: 8},
			"UTCDATEONLY":         Length{Min: 8, Max: 8},
			"UTCTIMEONLY":         Length{Min: 8, Max: 12},
			"UTCTIMESTAMP":        Length{Min: 1, Max: 27},
			"XMLDATA":             Length{Min: 1, Max: 65536},
		},
		LengthTag: map[int]LengthTag{
			SessionStatus1409:      LengthTag{Min: 1, Max: 1},
			CancelOnDisconnect6867: LengthTag{Min: 1, Max: 1},
			LanguageID6936:         LengthTag{Min: 1, Max: 1},

			f0.BeginString8:    LengthTag{Min: 7, Max: 7},
			f0.BodyLength9:     LengthTag{Min: 1, Max: 18},
			f0.CheckSum10:      LengthTag{Min: 3, Max: 3},
			f0.MsgType35:       LengthTag{Min: 1, Max: 2},
			f0.EncryptMethod98: LengthTag{Min: 1, Max: 1},
			f0.HeartBtInt108:   LengthTag{Min: 1, Max: 18},
		},
	}
	return c.Options()
}

func NewConfig() Config {
	return Config{
		Serialize:     make(map[string]Serialize),
		SerializeTag:  make(map[int]SerializeTag),
		EnumFormat:    make(map[string]EnumFormat),
		EnumFormatTag: make(map[int]EnumFormatTag),
		Length:        make(map[string]Length),
		LengthTag:     make(map[int]LengthTag),
	}
}

type Config struct {
	Serialize     map[string]Serialize
	SerializeTag  map[int]SerializeTag
	EnumFormat    map[string]EnumFormat
	EnumFormatTag map[int]EnumFormatTag
	Length        map[string]Length
	LengthTag     map[int]LengthTag
}

func (cfg Config) Options() []Option {
	var opts []Option

	for key, val := range cfg.Serialize {
		k, v := key, val
		o := func(c *Config) {
			v.Type = k
			c.Serialize[k] = v
		}
		opts = append(opts, o)
	}

	for key, val := range cfg.SerializeTag {
		k, v := key, val
		o := func(c *Config) {
			v.Tag = k
			c.SerializeTag[k] = v
		}
		opts = append(opts, o)
	}

	for key, val := range cfg.EnumFormat {
		k, v := key, val
		o := func(c *Config) {
			v.Type = k
			c.EnumFormat[k] = v
		}
		opts = append(opts, o)
	}

	for key, val := range cfg.EnumFormatTag {
		k, v := key, val
		o := func(c *Config) {
			v.Tag = k
			c.EnumFormatTag[k] = v
		}
		opts = append(opts, o)
	}

	for key, val := range cfg.Length {
		k, v := key, val
		o := func(c *Config) {
			v.Type = k
			c.Length[k] = v
		}
		opts = append(opts, o)
	}

	for key, val := range cfg.LengthTag {
		k, v := key, val
		o := func(c *Config) {
			v.Tag = k
			c.LengthTag[k] = v
		}
		opts = append(opts, o)
	}

	return opts
}

type Serialize struct {
	Type string
	Func string
}

type SerializeTag struct {
	Tag  int
	Func string
}

type EnumFormat struct {
	Type string
	Func func(string) (string, error)
}

type EnumFormatTag struct {
	Tag  int
	Func func(string) (string, error)
}

type Length struct {
	Type string
	Min  int
	Max  int
}

type LengthTag struct {
	Tag int
	Min int
	Max int
}

type Option func(*Config)

func WithSerialize(serializes ...Serialize) Option {
	return func(c *Config) {
		for _, s := range serializes {
			c.Serialize[s.Type] = s
		}
	}
}

func WithSerializeTag(serializes ...SerializeTag) Option {
	return func(c *Config) {
		for _, s := range serializes {
			c.SerializeTag[s.Tag] = s
		}
	}
}

func WithEnumFormat(formats ...EnumFormat) Option {
	return func(c *Config) {
		for _, f := range formats {
			c.EnumFormat[f.Type] = f
		}
	}
}

func WithEnumFormatTag(formats ...EnumFormatTag) Option {
	return func(c *Config) {
		for _, f := range formats {
			c.EnumFormatTag[f.Tag] = f
		}
	}
}

func WithLength(lengths ...Length) Option {
	return func(c *Config) {
		for _, l := range lengths {
			c.Length[l.Type] = l
		}
	}
}

func WithLengthTag(lengths ...LengthTag) Option {
	return func(c *Config) {
		for _, l := range lengths {
			c.LengthTag[l.Tag] = l
		}
	}
}
