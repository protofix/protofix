// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a FIX message format generator.
package genfix

import (
	"bytes"
	"encoding/xml"
	"errors"
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
	"github.com/protofix/protofix/strcase"
)

//go:generate go run main.go

type Package struct {
	Name   string
	Format string
	Info   string
	Spec   []byte
}

// Generate generates a Go FIX codecs by a XML specifications.
func Generate(dir string, pkg Package, opts ...Option) error {
	if len(opts) == 0 {
		return errors.New("missing configuration")
	}

	conf := NewConfig()

	for _, opt := range opts {
		opt(&conf)
	}

	pth := path.Join(dir, pkg.Name)
	fmt.Printf("Generating %s ...", pth)

	err := packageGenerate{pkg: pkg, dir: pth, conf: conf}.generate(pkg.Spec)
	if err != nil {
		return fmt.Errorf("generate package: %q, directory: %q, %w", pkg.Name, pth, err)
	}

	fmt.Println("ok")

	return nil
}

type packageGenerate struct {
	pkg  Package
	dir  string
	conf Config
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

// generate generates a Go FIX codec by a XML specification.
func (gen packageGenerate) generate(src []byte) error {
	var spec specfix.Spec
	r := bytes.NewReader(src)
	d := xml.NewDecoder(r)

	err := d.Decode(&spec)
	if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}

	if _, err := os.Stat(gen.dir); os.IsNotExist(err) {
		err = os.MkdirAll(gen.dir, os.ModeDir|0755)
		if err != nil {
			return fmt.Errorf("os: make directory: %w", err)
		}
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
	conf       Config
}

// generate generates a messages for the Go FIX codec by the XML specification.
func (gen messageGenerate) generate(src []byte) error {
	fldByName, err := gen.message.MemorizeFields(gen.header, gen.trailer, gen.fldByName, gen.compByName)
	if err != nil {
		return err
	}

	codecs := make([]codecData, 0, len(fldByName))
	var specialCodecs []codecData
	consts := make([]constantData, 0, len(fldByName))

	for _, fld := range fldByName {
		var min, max int

		if l, ok := gen.conf.LengthTag[fld.Define.Number]; ok {
			min = l.Min
			max = l.Max
		}

		if min == 0 || max == 0 {
			if l, ok := gen.conf.Length[fld.Define.Type]; ok {
				if min == 0 {
					min = l.Min
				}
				if max == 0 {
					max = l.Max
				}
			}
		}

		if min == 0 || max == 0 {
			if l, ok := gen.conf.Length["UNKNOWN"]; ok {
				if min == 0 {
					min = l.Min
				}
				if max == 0 {
					max = l.Max
				}
			}
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

		if s, ok := gen.conf.SerializeTag[fld.Define.Number]; ok {
			serializer = s.Func
		}

		if serializer == "" {
			if s, ok := gen.conf.Serialize[fld.Define.Type]; ok {
				serializer = s.Func
			}
		}

		if serializer == "" {
			if s, ok := gen.conf.Serialize["UNKNOWN"]; ok {
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

		if fld.Define.Number == f0.BodyLength9 || fld.Define.Number == f0.CheckSum10 {
			specialCodecs = append(specialCodecs, f)
			continue
		}

		codecs = append(codecs, f)
	}

	unknown := codecData{
		Name:       "Unknown",
		Type:       "UNKNOWN",
		Serializer: gen.conf.Serialize["UNKNOWN"].Func,
		MinLen:     gen.conf.Length["UNKNOWN"].Min,
		MaxLen:     gen.conf.Length["UNKNOWN"].Max,
	}

	specialCodecs = append(specialCodecs, unknown)

	funcs := template.FuncMap{
		"EnumFormat": func(tag int, typ, constraint string) string {
			var enum string
			var err error

			if f, ok := gen.conf.EnumFormatTag[tag]; ok {
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

			if enum != "" {
				return enum
			}

			if f, ok := gen.conf.EnumFormat["UNKNOWN"]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Format FIX enum: %s.", err)
				}
			}

			return enum
		},
	}

	tmpl, err := template.New("message").Funcs(funcs).Parse(messageTmpl)
	if err != nil {
		return fmt.Errorf("text template parse: %w", err)
	}

	_, err = tmpl.New("specialCodec").Parse(customCodecTmpl)
	if err != nil {
		return fmt.Errorf("text template parse: %w", err)
	}

	pkg := gen.pkg.Name + strings.ToLower(gen.message.Name)
	var b bytes.Buffer

	err = tmpl.ExecuteTemplate(&b, "message", struct {
		Package       string
		Format        string
		FormatDesc    string
		Year          int
		Message       string
		Codecs        []codecData
		SpecialCodecs []codecData
		Sort          []constantData
		Constants     []constantData
	}{
		Package:       pkg,
		Format:        strings.ToUpper(gen.pkg.Name),
		FormatDesc:    gen.pkg.Info,
		Year:          time.Now().Year(),
		Message:       gen.message.Name,
		Codecs:        codecs,
		SpecialCodecs: specialCodecs,
		Sort:          consts,
		Constants:     consts,
	})
	if err != nil {
		return fmt.Errorf("text template execute: %w", err)
	}

	dir := path.Join(gen.dir, pkg)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir|0755)
		if err != nil {
			return fmt.Errorf("os: make directory: %w", err)
		}
	}

	fname := fmt.Sprintf("%s_%s.go", gen.pkg.Name, strcase.CamelToSnake(gen.message.Name))
	pth := path.Join(dir, fname)
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

var messageTmpl = `
// Code generated by protofix; DO NOT EDIT.
// Copyright {{ .Year }} The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package {{ .Package }} is a format of the {{ .FormatDesc }} {{ .Message }} message.
package {{ .Package }}

import (
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
		{{- range .Codecs }}
			{{ .Name }}{{ .Number }}: f0.Fld{
				{{- if .Required -}} Req {{- else -}} Opt {{- end -}}, {{- /**/ -}}
				f0.ASCII, {{- /**/ -}}
				f0.{{ .Serializer }}(
					{{- $number := .Number -}}
					{{- $type := .Type -}}
					{{- range .Enums -}}
						{{- EnumFormat $number $type .Value -}}
						{{- if ne .Description "" -}}
							/* {{ .Description }} */
						{{- end -}}
						,
					{{- end -}}
				),
				{{- if and (ne .MinLen 0) (eq .MinLen .MaxLen) -}}
					f0.Con{ {{- .MinLen -}} }
				{{- else -}}
					f0.Var{ {{- .MinLen -}}, {{- .MaxLen -}} }
				{{- end -}} },
		{{- end }}
	},
	{{- range .SpecialCodecs -}}
		{{- template "specialCodec" . -}}
	{{- end }}
	Sort: []int{
	{{- range .Sort }}
		{{ .Name }}{{ .Number }}, // {{ .Type }}
	{{- end }}
	},
}

const Req, Opt = true, false

const (
{{- range .Constants }}
	{{ .Name }}{{ .Number }} = {{ .Number }} // {{ .Type }}
{{- end }}
)
`

var customCodecTmpl = `
{{ .Name }}{{ .Number }}: f0.{{ .Name }}Fld{  {{- /**/ -}}
	f0.ASCII, {{- /**/ -}}
	f0.{{ .Serializer }}(
		{{- $number := .Number -}}
		{{- $type := .Type -}}
		{{- range .Enums -}}
			{{- EnumFormat $number $type .Value -}}
			{{- if ne .Description "" -}}
				/* {{ .Description }} */
			{{- end -}}
			,
		{{- end -}}
	),
	{{- if and (ne .MinLen 0) (eq .MinLen .MaxLen) -}}
		f0.Con{ {{- .MinLen -}} },
	{{- else -}}
		f0.Var{ {{- .MinLen -}}, {{- .MaxLen -}} },
	{{- end -}}

}, `
