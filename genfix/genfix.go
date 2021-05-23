// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a FIX message format generator.
package genfix

import (
	"bytes"
	_ "embed"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/specfix"
	"github.com/protofix/protofix/strcase"
	"golang.org/x/tools/imports"
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

	err := fixPackage{pkg: pkg, dir: pth, conf: conf}.generate(pkg.Spec)
	if err != nil {
		return fmt.Errorf("generate package: %q, directory: %q, %w", pkg.Name, pth, err)
	}

	fmt.Println("ok")

	return nil
}

type fixPackage struct {
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
func (gen fixPackage) generate(src []byte) error {
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
		m := fixMessage{
			pkg:        gen.pkg,
			dir:        gen.dir,
			header:     spec.Header,
			message:    msg,
			trailer:    spec.Trailer,
			fldByName:  fldByName,
			compByName: compByName,
			conf:       gen.conf,
		}
		p, err := m.generate(src)
		if err != nil {
			return fmt.Errorf("generate message: %q package: %q, directory: %q, %w", msg.Name, gen.pkg.Name, gen.dir, err)
		}

		err = m.write(p)
		if err != nil {
			return fmt.Errorf("write message: %q package: %q, directory: %q, %w", msg.Name, gen.pkg.Name, gen.dir, err)
		}
	}

	return nil
}

type fixMessage struct {
	pkg        Package
	dir        string
	header     specfix.Header
	message    specfix.Message
	trailer    specfix.Trailer
	fldByName  map[string]specfix.FieldDefine
	compByName map[string]specfix.ComponentMemo
	conf       Config
}

//go:embed message.tmpl
var messageTmpl []byte

// generate generates a messages for the Go FIX codec by the XML specification.
func (fix fixMessage) generate(src []byte) ([]byte, error) {
	fldByName, err := fix.message.MemorizeFields(fix.header, fix.trailer, fix.fldByName, fix.compByName)
	if err != nil {
		return nil, err
	}

	codecs := make([]codecData, 0, len(fldByName))
	var specialCodecs []codecData
	consts := make([]constantData, 0, len(fldByName))

	for _, fld := range fldByName {
		var min, max int

		if l, ok := fix.conf.LengthTag[fld.Define.Number]; ok {
			min = l.Min
			max = l.Max
		}

		if min == 0 || max == 0 {
			if l, ok := fix.conf.Length[fld.Define.Type]; ok {
				if min == 0 {
					min = l.Min
				}
				if max == 0 {
					max = l.Max
				}
			}
		}

		if min == 0 || max == 0 {
			if l, ok := fix.conf.Length["UNKNOWN"]; ok {
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
			e := fieldEnum{Value: fix.pkg.Format}
			enums = append(enums, e)
		}

		for _, val := range fld.Define.Values {
			e := fieldEnum{Value: val.Enum, Description: val.Description}
			enums = append(enums, e)
		}

		var serializer string

		if s, ok := fix.conf.SerializeTag[fld.Define.Number]; ok {
			serializer = s.Func
		}

		if serializer == "" {
			if s, ok := fix.conf.Serialize[fld.Define.Type]; ok {
				serializer = s.Func
			}
		}

		if serializer == "" {
			if s, ok := fix.conf.Serialize["UNKNOWN"]; ok {
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
		Serializer: fix.conf.Serialize["UNKNOWN"].Func,
		MinLen:     fix.conf.Length["UNKNOWN"].Min,
		MaxLen:     fix.conf.Length["UNKNOWN"].Max,
	}

	specialCodecs = append(specialCodecs, unknown)

	funcs := template.FuncMap{
		"EnumFormat": func(tag int, typ, constraint string) string {
			var enum string
			var err error

			if f, ok := fix.conf.EnumFormatTag[tag]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Custom format FIX enum: %s.", err)
				}
			}

			if enum != "" {
				return enum
			}

			if f, ok := fix.conf.EnumFormat[typ]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Format FIX enum: %s.", err)
				}
			}

			if enum != "" {
				return enum
			}

			if f, ok := fix.conf.EnumFormat["UNKNOWN"]; ok {
				enum, err = f.Func(constraint)
				if err != nil {
					log.Printf("Format FIX enum: %s.", err)
				}
			}

			return enum
		},
	}

	tmpl, err := template.New("message").Funcs(funcs).Parse(string(messageTmpl))
	if err != nil {
		return nil, fmt.Errorf("text template parse: %w", err)
	}

	pkg := fix.pkg.Name + strings.ToLower(fix.message.Name)
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
		Format:        strings.ToUpper(fix.pkg.Name),
		FormatDesc:    fix.pkg.Info,
		Year:          time.Now().Year(),
		Message:       fix.message.Name,
		Codecs:        codecs,
		SpecialCodecs: specialCodecs,
		Sort:          consts,
		Constants:     consts,
	})
	if err != nil {
		return nil, fmt.Errorf("text template execute: %w", err)
	}

	dir := path.Join(fix.dir, pkg)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir|0755)
		if err != nil {
			return nil, fmt.Errorf("os: make directory: %w", err)
		}
	}

	fname := fmt.Sprintf("%s_%s.go", fix.pkg.Name, strcase.CamelToSnake(fix.message.Name))
	pth := path.Join(dir, fname)
	f, err := os.Create(pth)
	if err != nil {
		return nil, fmt.Errorf("os: create file %q: %w", pth, err)
	}
	defer f.Close()

	p, err := imports.Process("", b.Bytes(), &imports.Options{
		Fragment:  true,
		Comments:  true,
		TabIndent: true,
		TabWidth:  8,
	})
	if err != nil {
		return nil, fmt.Errorf("imports process: %w", err)
	}

	return p, nil
}

func (fix fixMessage) write(src []byte) error {
	pkg := fix.pkg.Name + strings.ToLower(fix.message.Name)
	dir := path.Join(fix.dir, pkg)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir|0755)
		if err != nil {
			return fmt.Errorf("os: make directory: %w", err)
		}
	}

	fname := fmt.Sprintf("%s_%s.go", fix.pkg.Name, strcase.CamelToSnake(fix.message.Name))
	pth := path.Join(dir, fname)
	f, err := os.Create(pth)
	if err != nil {
		return fmt.Errorf("os: create file %q: %w", pth, err)
	}
	defer f.Close()

	_, err = f.Write(src)
	if err != nil {
		return fmt.Errorf("os file write %q: %w", pth, err)
	}
	return nil
}
