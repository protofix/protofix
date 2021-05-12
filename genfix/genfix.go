// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a FIX message format generator.
package genfix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/specfix"
	"github.com/protofix/protofix/stringfix"
)

//go:generate go run main.go

type Package struct {
	Name   string
	Format string
	Info   string
	Spec   []byte
}

// Generate generates a Go FIX codecs by a XML specifications.
func Generate(dir string, specs []Package, opts ...Option) error {
	conf := Configuration{
		Serialize: map[string]Serialize{
			"BOOLEAN":           Serialize{Func: "BoolDefault"},
			"BYTES":             Serialize{Func: "Bytes"},
			"CHECKSUM":          Serialize{Func: "ChecksumString"},
			"INT":               Serialize{Func: "IntDefault"},
			"SEQNUM":            Serialize{Func: "SeqNum"},
			"LENGTH":            Serialize{Func: "BodyLength"},
			"STRING":            Serialize{Func: "String"},
			"UNKNOWN":           Serialize{Func: "Unknown"},
			"UTCTIMESTAMP":      Serialize{Func: "UTCTimestampNanosecondTime"},
			"UTCTIMESTAMPMILLI": Serialize{Func: "UTCTimestampNanosecondTime"},
			"UTCTIMESTAMPNANO":  Serialize{Func: "UTCTimestampNanosecondTime"},
		},
		CustomSerialize: map[int]CustomSerialize{
			8:   CustomSerialize{Func: "StringDefault"},
			10:  CustomSerialize{Func: "ChecksumString"},
			108: CustomSerialize{Func: "SecondsDuration"},
		},
		EnumFormat: map[string]EnumFormat{
			"BOOLEAN":           EnumFormat{Func: func(s string) (string, error) { m := map[string]string{"Y": "true", "N": "false"}; return m[s], nil }},
			"BYTES":             EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"CHECKSUM":          EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"INT":               EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"SEQNUM":            EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"LENGTH":            EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"STRING":            EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UNKNOWN":           EnumFormat{Func: func(s string) (string, error) { return s, nil }},
			"UTCTIMESTAMP":      EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMPMILLI": EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
			"UTCTIMESTAMPNANO":  EnumFormat{Func: func(s string) (string, error) { return fmt.Sprintf("%q", s), nil }},
		},
		CustomEnumFormat: map[int]EnumFormat{
			108: EnumFormat{Func: func(s string) (string, error) { return s + "*time.Second", nil }},
		},
		Length: map[string]Length{
			"BOOLEAN":           Length{Min: 1, Max: 1},
			"BYTES":             Length{Min: 1, Max: 65536},
			"CHECKSUM":          Length{Min: 3, Max: 3},
			"INT":               Length{Min: 1, Max: 19},
			"SEQNUM":            Length{Min: 1, Max: 19},
			"LENGTH":            Length{Min: 1, Max: 19},
			"STRING":            Length{Min: 1, Max: 65536},
			"UNKNOWN":           Length{Min: 1, Max: 65536},
			"UTCTIMESTAMP":      Length{Min: 1, Max: 27},
			"UTCTIMESTAMPMILLI": Length{Min: 1, Max: 27},
			"UTCTIMESTAMPNANO":  Length{Min: 1, Max: 27},
		},
	}

	for _, opt := range opts {
		opt(&conf)
	}

	for _, pkg := range specs {
		pth := path.Join(dir, pkg.Name)
		fmt.Printf("Generating %s ...", pth)

		err := packageGenerate{pkg: pkg, dir: pth, conf: conf}.generate(pkg.Spec)
		if err != nil {
			return fmt.Errorf("generate package: %q, directory: %q, %w", pkg.Name, pth, err)
		}

		fmt.Println("ok")
	}

	return nil
}

type packageGenerate struct {
	pkg  Package
	dir  string
	conf Configuration
}

type constantData struct {
	Number int
	Name   string
	Type   string
}

type codecData struct {
	Number     int
	Name       string
	Type       string
	Required   bool
	Serializer string
	Enums      []fieldEnum
	MinLen     int
	MaxLen     int
}

type fieldEnum struct {
	Value       string
	Description string
}

// type tmplField struct {
// 	Name     string
// 	Required bool
// 	Define   fieldDefine
// }

// generate generates a Go FIX codec by a XML specification.
func (gen packageGenerate) generate(src []byte) error {
	var spec specfix.Spec
	r := bytes.NewReader(src)
	d := xml.NewDecoder(r)

	err := d.Decode(&spec)
	if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}

	consts := make([]constantData, 0, len(spec.Fields))

	for _, fld := range spec.Fields {
		f := constantData{
			Number: fld.Number,
			Name:   fld.Name,
			Type:   fld.Type,
		}
		consts = append(consts, f)
	}

	tmpl, err := template.New("Main").Parse(mainTmpl)
	if err != nil {
		return fmt.Errorf("text template parse: %w", err)
	}

	var b bytes.Buffer

	err = tmpl.ExecuteTemplate(&b, "Main", struct {
		Package    string
		FormatDesc string
		Year       int
		Constants  []constantData
	}{
		Package:    gen.pkg.Name,
		FormatDesc: gen.pkg.Info,
		Year:       time.Now().Year(),
		Constants:  consts,
	})
	if err != nil {
		return fmt.Errorf("text template execute: %w", err)
	}

	if _, err := os.Stat(gen.dir); os.IsNotExist(err) {
		err = os.MkdirAll(gen.dir, os.ModeDir|0755)
		if err != nil {
			return fmt.Errorf("os: make directory: %w", err)
		}
	}

	fname := fmt.Sprintf("%s_format.go", gen.pkg.Name)
	pth := path.Join(gen.dir, fname)
	f, err := os.Create(pth)
	if err != nil {
		return fmt.Errorf("os: create file %q: %w", pth, err)
	}
	defer f.Close()

	p, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("go format source: %w", err)
	}

	_, err = f.Write(p)
	if err != nil {
		return fmt.Errorf("os file write %q: %w", pth, err)
	}

	fldByName := make(map[string]specfix.FieldDefine, len(spec.Fields))

	for _, f := range spec.Fields {
		fldByName[f.Name] = f
	}

	compByName, err := spec.MemorizeComponents()
	if err != nil {
		return fmt.Errorf("xml spec: memorize components: %w", err)
	}

	for _, msg := range spec.Messages {
		m := messageGenerate{
			pkg:        gen.pkg,
			dir:        gen.dir,
			header:     spec.Header,
			message:    msg,
			trailer:    spec.Trailer,
			fldByName:  fldByName,
			compByName: compByName,
			conf:       gen.conf,
		}
		err = m.generate(src)
		if err != nil {
			return fmt.Errorf("generate message: %q package: %q, directory: %q, %w", msg.Name, gen.pkg.Name, gen.dir, err)
		}
	}

	return nil
}

type messageGenerate struct {
	pkg        Package
	dir        string
	header     specfix.Header
	message    specfix.Message
	trailer    specfix.Trailer
	fldByName  map[string]specfix.FieldDefine
	compByName map[string]specfix.ComponentMemo
	conf       Configuration
}

// generate generates a messages for the Go FIX codec by the XML specification.
func (gen messageGenerate) generate(src []byte) error {
	fldByName, err := gen.message.MemorizeFields(gen.header, gen.trailer, gen.fldByName, gen.compByName)
	if err != nil {
		return err
	}

	var (
		bodyLengthCodec codecData
		checkSumCodec   codecData
		unknownCodec    = codecData{
			Type:       "UNKNOWN",
			Serializer: gen.conf.Serialize["UNKNOWN"].Func,
			MinLen:     gen.conf.Length["UNKNOWN"].Min,
			MaxLen:     gen.conf.Length["UNKNOWN"].Max,
		}
	)

	codecs := make([]codecData, 0, len(fldByName))
	consts := make([]constantData, 0, len(fldByName))

	for _, fld := range fldByName {
		min := fld.Define.MinLen
		if min == 0 {
			min = gen.conf.Length[fld.Define.Type].Min
		}

		max := fld.Define.MaxLen
		if max == 0 {
			max = gen.conf.Length[fld.Define.Type].Max
		}

		enums := make([]fieldEnum, 0, len(fld.Define.Values))

		if fld.Define.Number == f0.BeginString8 {
			e := fieldEnum{Value: gen.pkg.Format}
			enums = append(enums, e)
		}

		for _, val := range fld.Define.Values {
			e := fieldEnum{Value: val.Enum, Description: val.Description}
			enums = append(enums, e)
		}

		var serializer string

		if s, ok := gen.conf.CustomSerialize[fld.Define.Number]; ok {
			serializer = s.Func
		}

		if serializer == "" {
			if s, ok := gen.conf.Serialize[fld.Define.Type]; ok {
				serializer = s.Func
			}
		}

		f := codecData{
			Number:     fld.Define.Number,
			Name:       fld.Define.Name,
			Type:       fld.Define.Type,
			Required:   fld.Required,
			Serializer: serializer,
			Enums:      enums,
			MinLen:     min,
			MaxLen:     max,
		}

		c := constantData{
			Number: f.Number,
			Name:   f.Name,
			Type:   f.Type,
		}

		consts = append(consts, c)

		if fld.Define.Number == f0.BodyLength9 {
			bodyLengthCodec = f
			continue
		}

		if fld.Define.Number == f0.CheckSum10 {
			checkSumCodec = f
			continue
		}

		codecs = append(codecs, f)
	}

	funcs := template.FuncMap{
		"EnumFormat": func(tag int, typ, constraint string) string {
			var enum string
			var err error

			if f, ok := gen.conf.CustomEnumFormat[tag]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Custom format FIX enum: %s.", err)
				}
			}

			if enum != "" {
				return enum
			}

			if f, ok := gen.conf.EnumFormat[typ]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Format FIX enum: %s.", err)
				}
			}

			return enum
		},
	}

	tmpl, err := template.New("Message").Funcs(funcs).Parse(messageTmpl)
	if err != nil {
		return fmt.Errorf("text template parse: %w", err)
	}

	var b bytes.Buffer

	err = tmpl.ExecuteTemplate(&b, "Message", struct {
		Package         string
		Format          string
		FormatDesc      string
		Year            int
		Message         string
		Codecs          []codecData
		BodyLengthCodec codecData
		CheckSumCodec   codecData
		UnknownCodec    codecData
		Sort            []constantData
	}{
		Package:         gen.pkg.Name,
		Format:          strings.ToUpper(gen.pkg.Name),
		FormatDesc:      gen.pkg.Info,
		Year:            time.Now().Year(),
		Message:         gen.message.Name,
		Codecs:          codecs,
		BodyLengthCodec: bodyLengthCodec,
		CheckSumCodec:   checkSumCodec,
		UnknownCodec:    unknownCodec,
		Sort:            consts,
	})
	if err != nil {
		return fmt.Errorf("text template execute: %w", err)
	}

	fname := fmt.Sprintf("%s_%s_format.go", gen.pkg.Name, stringfix.ToSnakeCase(gen.message.Name))
	pth := path.Join(gen.dir, fname)
	f, err := os.Create(pth)
	if err != nil {
		return fmt.Errorf("os: create file %q: %w", pth, err)
	}
	defer f.Close()

	p, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("go format source: %w", err)
	}

	_, err = f.Write(p)
	if err != nil {
		return fmt.Errorf("os file write %q: %w", pth, err)
	}
	return nil
}

var mainTmpl = `
// Code generated by protofix; DO NOT EDIT.
// Copyright {{ .Year }} The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package {{ .Package }} provides {{ .FormatDesc }} protocol format.
package {{ .Package }}

const Req, Opt = true, false

const (
{{- range .Constants }}
	{{ .Name }}{{ .Number }} = {{ .Number }} // {{ .Type }}
{{- end }}
)
`

type Configuration struct {
	Serialize        map[string]Serialize
	CustomSerialize  map[int]CustomSerialize
	EnumFormat       map[string]EnumFormat
	CustomEnumFormat map[int]EnumFormat
	Length           map[string]Length
}

type Serialize struct {
	Type string
	Func string
}

type CustomSerialize struct {
	Tag  int
	Func string
}

type EnumFormat struct {
	Type string
	Func func(string) (string, error)
}

type Length struct {
	Type string
	Min  int
	Max  int
}

type Option func(*Configuration)

func WithSerialize(serializes ...Serialize) Option {
	return func(c *Configuration) {
		for _, s := range serializes {
			c.Serialize[s.Type] = s
		}
	}
}

func WithCustomSerialize(serializes ...CustomSerialize) Option {
	return func(c *Configuration) {
		for _, s := range serializes {
			c.CustomSerialize[s.Tag] = s
		}
	}
}

func WithEnumFormat(formats ...EnumFormat) Option {
	return func(c *Configuration) {
		for _, f := range formats {
			c.EnumFormat[f.Type] = f
		}
	}
}

func WithLength(lengths ...Length) Option {
	return func(c *Configuration) {
		for _, l := range lengths {
			c.Length[l.Type] = l
		}
	}
}

var messageTmpl = `
// Code generated by protofix; DO NOT EDIT.
// Copyright {{ .Year }} The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package {{ .Package }} provides {{ .FormatDesc }} protocol format.
package {{ .Package }}

import (
	"time"

	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	{{ .Format }}{{ .Message }}Marshaler   = marshfix.Marshal{Tag: "{{ .Format }}", Format: {{ .Format }}{{ .Message }}}
	{{ .Format }}{{ .Message }}Unmarshaler = marshfix.Unmarshal{Tag: "{{ .Format }}", Format: {{ .Format }}{{ .Message }}}
)

// {{ .Format }}{{ .Message }} is a {{ .FormatDesc }} format of the {{ .Message }} message which maps the codecs into individual fields.
var {{ .Format }}{{ .Message }} = f0.Format{
	Fields: map[int]f0.Codec{
		{{- range $index, $codec := .Codecs }}
			{{ $codec.Name }}{{ $codec.Number }}: f0.Fld{
				{{- if $codec.Required -}} Req {{- else -}} Opt {{- end -}}, {{- /**/ -}}
				f0.ASCII, {{- /**/ -}}
				f0.{{ $codec.Serializer }}(
					{{- range $codec.Enums -}}
						{{- EnumFormat $codec.Number $codec.Type .Value -}}
						{{- if ne .Description "" -}}
							/* {{ .Description }} */
						{{- end -}}
						,
					{{- end -}}
				),
				{{- if and (ne $codec.MinLen 0) (eq $codec.MinLen $codec.MaxLen) -}}
					f0.Con{ {{- $codec.MinLen -}} }
				{{- else -}}
					f0.Var{ {{- $codec.MinLen -}}, {{- $codec.MaxLen -}} }
				{{- end -}} },
		{{- end }}
	},
	BodyLength9:	f0.BodyLengthFld{ {{- /**/ -}}
		f0.ASCII, {{- /**/ -}}
		f0.{{ .BodyLengthCodec.Serializer }}(
			{{- range .BodyLengthCodec.Enums -}}
				{{- EnumFormat .BodyLengthCodec.Number .BodyLengthCodec.Type .Value -}}
				{{- if ne .Description "" -}}
					/* {{ .Description }} */
				{{- end -}}
				,
			{{- end -}}
		),
		{{- if and (ne .BodyLengthCodec.MinLen 0) (eq .BodyLengthCodec.MinLen .BodyLengthCodec.MaxLen) -}}
			f0.Con{ {{- .BodyLengthCodec.MinLen -}} }
		{{- else -}}
			f0.Var{ {{- .BodyLengthCodec.MinLen -}}, {{- .BodyLengthCodec.MaxLen -}} }
		{{- end -}}
	},
	CheckSum10:		f0.ChecksumStringFld{ {{- /**/ -}}
		f0.ASCII, {{- /**/ -}}
		f0.{{ .CheckSumCodec.Serializer }}(
			{{- range .CheckSumCodec.Enums -}}
				{{- EnumFormat .CheckSumCodec.Number .CheckSumCodec.Type .Value -}}
				{{- if ne .Description "" -}}
					/* {{ .Description }} */
				{{- end -}}
				,
			{{- end -}}
		),
		{{- if and (ne .CheckSumCodec.MinLen 0) (eq .CheckSumCodec.MinLen .CheckSumCodec.MaxLen) -}}
			f0.Con{ {{- .CheckSumCodec.MinLen -}} }
		{{- else -}}
			f0.Var{ {{- .CheckSumCodec.MinLen -}}, {{- .CheckSumCodec.MaxLen -}} }
		{{- end -}}
	},
	Unknown:			f0.UnknownFld{ {{- /**/ -}}
		f0.ASCII, {{- /**/ -}}
		f0.{{ .UnknownCodec.Serializer }}(
			{{- range .UnknownCodec.Enums -}}
				{{- EnumFormat .UnknownCodec.Number .UnknownCodec.Type .Value -}}
				{{- if ne .Description "" -}}
					/* {{ .Description }} */
				{{- end -}}
				,
			{{- end -}}
		),
		{{- if and (ne .UnknownCodec.MinLen 0) (eq .UnknownCodec.MinLen .UnknownCodec.MaxLen) -}}
			f0.Con{ {{- .UnknownCodec.MinLen -}} }
		{{- else -}}
			f0.Var{ {{- .UnknownCodec.MinLen -}}, {{- .UnknownCodec.MaxLen -}} }
		{{- end -}}
	},
	Sort: []int{
	{{- range .Sort }}
		{{ .Name }}{{ .Number }}, // {{ .Type }}
	{{- end }}
	},
}
`
