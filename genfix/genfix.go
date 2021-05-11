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

	f0 "github.com/danil/protofix/codecfix"
	"github.com/danil/protofix/specfix"
	"github.com/danil/protofix/stringfix"
)

//go:generate go run main.go

type Package struct {
	Name string
	Info string
	Spec []byte
}

// Generate generates a Go FIX codecs by a XML specifications.
func Generate(dir string, specs []Package) error {
	for _, pkg := range specs {
		pth := path.Join(dir, pkg.Name)
		fmt.Printf("Generating %s ...", pth)

		err := packageGen{pkg: pkg, dir: pth}.generate(pkg.Spec)
		if err != nil {
			return fmt.Errorf("generate package: %q, directory: %q, %w", pkg.Name, pth, err)
		}

		fmt.Println("ok")
	}

	return nil
}

type packageGen struct {
	pkg Package
	dir string
}

type fieldDef struct {
	Number int
	Name   string
	Type   string
	MinLen int
	MaxLen int
}

type componentDef struct {
	Name   string
	Fields []tplField
}

type tplMessage struct {
	Name   string
	Fields []tplField
}

type tplField struct {
	Name     string
	Required bool
	Def      fieldDef
}

// generate generates a Go FIX codec by a XML specification.
func (g packageGen) generate(src []byte) error {
	var spec specfix.Spec
	r := bytes.NewReader(src)
	d := xml.NewDecoder(r)

	err := d.Decode(&spec)
	if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}

	fldDefs := make([]fieldDef, 0, len(spec.FieldDefs))
	fldDefByName := make(map[string]fieldDef, len(spec.FieldDefs))

	for _, fld := range spec.FieldDefs {
		minLen := fld.MinLen
		maxLen := fld.MaxLen
		if minLen == 0 && maxLen == 0 {
			typeLen := typeLengths[tmplTypeByText[fld.Type]]
			minLen = typeLen.Min
			maxLen = typeLen.Max
		}

		fldDef := fieldDef{
			Number: fld.Number,
			Name:   fld.Name,
			Type:   fld.Type,
			MinLen: minLen,
			MaxLen: maxLen,
			// Values []*xmlValue // FIXME: enums
		}
		fldDefs = append(fldDefs, fldDef)
		fldDefByName[fld.Name] = fldDef
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
		Fields     []fieldDef
	}{
		Package:    g.pkg.Name,
		FormatDesc: g.pkg.Info,
		Year:       time.Now().Year(),
		Fields:     fldDefs,
	})
	if err != nil {
		return fmt.Errorf("text template execute: %w", err)
	}

	if _, err := os.Stat(g.dir); os.IsNotExist(err) {
		err = os.MkdirAll(g.dir, os.ModeDir|0755)
		if err != nil {
			return fmt.Errorf("os: make directory: %w", err)
		}
	}

	fname := fmt.Sprintf("%s_format.go", g.pkg.Name)
	pth := path.Join(g.dir, fname)
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

	compDefByName := make(map[string]componentDef, len(spec.ComponentDefs))

	// Populating the fields of the components until satisfying all
	// component-to-component references due to the ability
	// to follow components in an unsorted order.
	for i, incomplete := 0, true; incomplete; i++ {
		if i == 1000 {
			log.Printf("Too many attempts %d to fill in the fields of the components.", i)
			break
		}

		incomplete = false
		for _, comp := range spec.ComponentDefs {
			if _, ok := compDefByName[comp.Name]; ok {
				continue
			}

			flds, ok := flattenComponent(comp, fldDefByName, compDefByName)
			if !ok {
				incomplete = true
				continue
			}

			compDefByName[comp.Name] = componentDef{Name: comp.Name, Fields: flds}
		}
	}

	for _, msg := range spec.Messages {
		m := messageGen{
			pkg:           g.pkg,
			dir:           g.dir,
			message:       tplMessage{Name: msg.Name},
			header:        spec.Header,
			fields:        msg.Fields,
			components:    msg.Components,
			groups:        msg.Groups,
			trailer:       spec.Trailer,
			fieldDefs:     fldDefByName,
			componentDefs: compDefByName,
		}
		err = m.generate(src)
		if err != nil {
			return fmt.Errorf("generate message: %q package: %q, directory: %q, %w", msg.Name, g.pkg.Name, g.dir, err)
		}
	}

	return nil
}

// flattenComponent transforms the component definition containing nesting to the flat fields.
func flattenComponent(src specfix.ComponentDef, fldDefs map[string]fieldDef, compDefs map[string]componentDef) ([]tplField, bool) {
	var flds []tplField

	for _, fld := range src.Fields {
		fDef, ok := fldDefs[fld.Name]
		if !ok {
			return nil, false
		}
		f := tplField{Name: fld.Name, Required: fld.Required, Def: fDef}
		flds = append(flds, f)
	}

	for _, comp := range src.Components {
		cDef, ok := compDefs[comp.Name]
		if !ok {
			return nil, false
		}
		for _, fld := range cDef.Fields {
			fDef, ok := fldDefs[fld.Name]
			if !ok {
				return nil, false
			}
			f := tplField{Name: fld.Name, Required: fld.Required, Def: fDef}
			flds = append(flds, f)
		}
	}

	for _, grp := range src.Groups {
		f, ok := flattenGroup(grp, fldDefs, compDefs)
		if !ok {
			return nil, false
		}
		flds = append(flds, f...)
	}

	return flds, true
}

// flattenGroup transforms the group containing nesting to the flat fields.
func flattenGroup(src specfix.Group, fldDefs map[string]fieldDef, compDefs map[string]componentDef) ([]tplField, bool) {
	var flds []tplField

	for _, fld := range src.Fields {
		fDef, ok := fldDefs[fld.Name]
		if !ok {
			return nil, false
		}
		f := tplField{Name: fld.Name, Required: fld.Required, Def: fDef}
		flds = append(flds, f)
	}

	for _, comp := range src.Components {
		cDef, ok := compDefs[comp.Name]
		if !ok {
			return nil, false
		}
		for _, fld := range cDef.Fields {
			fDef, ok := fldDefs[fld.Name]
			if !ok {
				return nil, false
			}
			f := tplField{Name: fld.Name, Required: fld.Required, Def: fDef}
			flds = append(flds, f)
		}
	}

	for _, grp := range src.Groups {
		f, ok := flattenGroup(grp, fldDefs, compDefs)
		if !ok {
			return nil, false
		}
		flds = append(flds, f...)
	}

	return flds, true
}

type messageGen struct {
	pkg           Package
	dir           string
	message       tplMessage
	header        specfix.Header
	fields        []specfix.Field
	components    []specfix.Component
	groups        []specfix.Group
	trailer       []specfix.Field
	fieldDefs     map[string]fieldDef
	componentDefs map[string]componentDef
}

// generate generates a messages for the Go FIX codec by the XML specification.
func (g messageGen) generate(src []byte) error {
	// Transform a fields/components/groups specs to a fields.
	var flds []tplField
	xfmr := xfmrSpec{fieldDefs: g.fieldDefs, componentDefs: g.componentDefs}
	flds = append(flds, xfmr.fields(g.header.Fields...)...)
	flds = append(flds, xfmr.components(g.header.Components...)...)
	flds = append(flds, xfmr.groups(g.header.Groups...)...)
	flds = append(flds, xfmr.fields(g.fields...)...)
	flds = append(flds, xfmr.components(g.components...)...)
	flds = append(flds, xfmr.fields(g.trailer...)...)

	g.message.Fields = flds

	tmpl, err := template.New("Message").Parse(messageTmpl)
	if err != nil {
		return fmt.Errorf("text template parse: %w", err)
	}

	defs := make([]tplField, 0, len(g.message.Fields))
	for _, f := range g.message.Fields {
		if f.Def.Number == f0.BodyLength9 || f.Def.Number == f0.CheckSum10 {
			continue
		}
		defs = append(defs, f)
	}

	var b bytes.Buffer

	err = tmpl.ExecuteTemplate(&b, "Message", struct {
		Package     string
		Format      string
		FormatDesc  string
		Year        int
		Message     string
		Definitions []tplField
		Sort        []tplField
	}{
		Package:     g.pkg.Name,
		Format:      strings.ToUpper(g.pkg.Name),
		FormatDesc:  g.pkg.Info,
		Year:        time.Now().Year(),
		Message:     g.message.Name,
		Definitions: defs,
		Sort:        g.message.Fields,
	})
	if err != nil {
		return fmt.Errorf("text template execute: %w", err)
	}

	fname := fmt.Sprintf("%s_%s_format.go", g.pkg.Name, stringfix.ToSnakeCase(g.message.Name))
	pth := path.Join(g.dir, fname)
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

// xfmrSpec transforms XML fields/components/groups to template fields.
type xfmrSpec struct {
	fieldDefs     map[string]fieldDef
	componentDefs map[string]componentDef
}

// fields transforms xml fields to template fields.
func (xfmr xfmrSpec) fields(src ...specfix.Field) []tplField {
	var flds []tplField

	for _, fld := range src {
		fDef, ok := xfmr.fieldDefs[fld.Name]
		if !ok {
			log.Printf("Missing field %q specification/definition.", fld.Name)
		}

		f := tplField{Name: fld.Name, Required: fld.Required, Def: fDef}
		flds = append(flds, f)
	}

	return flds
}

// components transforms xml components to template fields.
func (xfmr xfmrSpec) components(src ...specfix.Component) []tplField {
	var flds []tplField

	for _, comp := range src {
		cDef, ok := xfmr.componentDefs[comp.Name]
		if !ok {
			log.Printf("Missing component %q specification/definition.", comp.Name)
			continue
		}
		flds = append(flds, cDef.Fields...)
	}

	return flds
}

// groups transforms xml groups to template fields.
func (xfmr xfmrSpec) groups(src ...specfix.Group) []tplField {
	var flds []tplField

	for _, group := range src {
		flds = append(flds, xfmr.fields(group.Fields...)...)
		flds = append(flds, xfmr.components(group.Components...)...)
		flds = append(flds, xfmr.groups(group.Groups...)...)
	}

	return flds
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
{{- range .Fields }}
	{{ .Name }}{{ .Number }} = {{ .Number }} // {{ .Type }}
{{- end }}
)
`

type fieldType int

const (
	tmplBool fieldType = 1 + iota
)

var tmplTypeByText = map[string]fieldType{
	"BOOLEAN": tmplBool,
}

type typeLength struct{ Min, Max int }

var typeLengths = map[fieldType]typeLength{
	tmplBool: typeLength{1, 1},
}

var messageTmpl = `
// Code generated by protofix; DO NOT EDIT.
// Copyright {{ .Year }} The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package {{ .Package }} provides {{ .FormatDesc }} protocol format.
package {{ .Package }}

import (
	f0 "github.com/danil/protofix/codecfix"
	"github.com/danil/protofix/marshfix"
)

var (
	{{ .Format }}{{ .Message }}Marshaler   = marshfix.Marshal{Tag: "{{ .Format }}", Format: {{ .Format }}{{ .Message }}}
	{{ .Format }}{{ .Message }}Unmarshaler = marshfix.Unmarshal{Tag: "{{ .Format }}", Format: {{ .Format }}{{ .Message }}}
)

// {{ .Format }}{{ .Message }} is a {{ .FormatDesc }} format of the {{ .Message }} message which maps the codecs into individual fields.
var {{ .Format }}{{ .Message }} = f0.Format{
	Fields: map[int]f0.Codec{
		{{- range .Definitions }}
			{{ .Name }}{{ .Def.Number }}: f0.Fld{
				{{- if .Required -}}
					Req
				{{- else -}}
					Opt
				{{- end -}}, f0.ASCII, f0.StringDefault(),
				{{- if and (ne .Def.MinLen 0) (eq .Def.MinLen .Def.MaxLen) -}}
					f0.Con{ {{- .Def.MinLen -}} }
				{{- else -}}
					f0.Var{ {{- .Def.MinLen -}}, {{- .Def.MaxLen -}} }
				{{- end -}} }, // {{ .Def.Type -}}
		{{- end }}
	},
	BodyLength9:	f0.BodyLengthFld{},
	CheckSum10:		f0.ChecksumStringFld{},
	Unknown:			f0.UnknownFld{},
	Sort: []int{
	{{- range .Sort }}
		{{ .Name }}{{ .Def.Number }}, // {{ .Def.Type }}
	{{- end }}
	},
}
`
