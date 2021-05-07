// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package genfix is a XML specification parser and FIX message format generator.
package genfix

import (
	"encoding/xml"
	"fmt"
	"io"
)

// xmlSpec represents a FIX messages based on a XML specification.
type xmlSpec struct {
	XMLName     xml.Name     `xml:"fix"`
	Type        string       `xml:"type,attr"`
	Major       string       `xml:"major,attr"`
	Minor       string       `xml:"minor,attr"`
	Servicepack string       `xml:"servicepack,attr"`
	Header      []*xmlFld    `xml:"header>field"`
	Messages    []*xmlMsg    `xml:"messages>message"`
	Trailer     []*xmlFld    `xml:"trailer>field"`
	Components  []*xmlComp   `xml:"components>component"`
	Fields      []*xmlFldDef `xml:"fields>field"`
}

type xmlMsg struct {
	Name    string      `xml:"name,attr"`
	MsgCat  string      `xml:"msgcat,attr"`
	MsgType string      `xml:"msgtype,attr"`
	Comps   []*xmlComp  `xml:"component"`
	Groups  []*xmlGroup `xml:"group"`
	Fields  []*xmlFld   `xml:"field"`
}

type xmlFld struct {
	Name     string
	Required bool
}

func (fld *xmlFld) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type field struct {
		XMLName  xml.Name `xml:"field"`
		Name     string   `xml:"name,attr"`
		Required string   `xml:"required,attr"`
	}
	var f field
	err := d.DecodeElement(&f, &start)
	if err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}
	fld.Name = f.Name
	if f.Required == "Y" {
		fld.Required = true
	}
	return nil
}

type xmlFldDef struct {
	Number int
	Name   string
	Type   string
	MinLen int
	MaxLen int
	Values []*xmlValue
}

func (fld *xmlFldDef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type field struct {
		XMLName xml.Name    `xml:"field"`
		Number  int         `xml:"number,attr"`
		Name    string      `xml:"name,attr"`
		Type    string      `xml:"type,attr"`
		Len     int         `xml:"length,attr"`
		MinLen  int         `xml:"minlength,attr"`
		MaxLen  int         `xml:"maxlength,attr"`
		Values  []*xmlValue `xml:"value"`
	}
	var f field
	err := d.DecodeElement(&f, &start)
	if err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}
	fld.Number = f.Number
	fld.Name = f.Name
	fld.Type = f.Type
	if f.Len != 0 {
		fld.MinLen = f.Len
		fld.MaxLen = f.Len
	} else {
		fld.MinLen = f.MinLen
		fld.MaxLen = f.MaxLen
	}
	fld.Values = f.Values
	return nil
}

type xmlValue struct {
	Enum        string `xml:"enum,attr"`
	Description string `xml:"description,attr"`
}

type xmlComp struct {
	Name     string      `xml:"name,attr"`
	Required string      `xml:"required,attr"`
	Comps    []*xmlComp  `xml:"component"`
	Groups   []*xmlGroup `xml:"group"`
	Fields   []*xmlFld   `xml:"field"`
}

type xmlGroup struct {
	Name     string      `xml:"name,attr"`
	Required string      `xml:"required,attr"`
	Comps    []*xmlComp  `xml:"component"`
	Groups   []*xmlGroup `xml:"group"`
	Fields   []*xmlFld   `xml:"field"`
}
