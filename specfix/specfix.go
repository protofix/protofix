// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package specfix represents XML specification of the FIX protocol.
package specfix

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"io"
)

//go:embed MOEX44.xml
var MOEX44 []byte

// Spec represents a FIX messages based on a XML specification of the FIX protocol.
type Spec struct {
	XMLName       xml.Name       `xml:"fix"`
	Type          string         `xml:"type,attr"`
	Major         string         `xml:"major,attr"`
	Minor         string         `xml:"minor,attr"`
	Servicepack   string         `xml:"servicepack,attr"`
	Header        Header         `xml:"header"`
	Messages      []Message      `xml:"messages>message"`
	Trailer       []Field        `xml:"trailer>field"`
	ComponentDefs []ComponentDef `xml:"components>component"`
	FieldDefs     []FieldDef     `xml:"fields>field"`
}

type Header struct {
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

type Message struct {
	Name       string      `xml:"name,attr"`
	MsgCat     string      `xml:"msgcat,attr"`
	MsgType    string      `xml:"msgtype,attr"`
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

type Field struct {
	Name     string
	Required bool
}

func (fld *Field) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

type FieldDef struct {
	Number int
	Name   string
	Type   string
	MinLen int
	MaxLen int
	Values []Value
}

func (fld *FieldDef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type field struct {
		XMLName xml.Name `xml:"field"`
		Number  int      `xml:"number,attr"`
		Name    string   `xml:"name,attr"`
		Type    string   `xml:"type,attr"`
		Len     int      `xml:"length,attr"`
		MinLen  int      `xml:"minlength,attr"`
		MaxLen  int      `xml:"maxlength,attr"`
		Values  []Value  `xml:"value"`
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

type Value struct {
	Enum        string `xml:"enum,attr"`
	Description string `xml:"description,attr"`
}

type Component struct {
	Name     string
	Required bool
}

func (fld *Component) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type component struct {
		XMLName  xml.Name `xml:"component"`
		Name     string   `xml:"name,attr"`
		Required string   `xml:"required,attr"`
	}
	var comp component
	err := d.DecodeElement(&comp, &start)
	if err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("xml decoder: decode: %w", err)
	}
	fld.Name = comp.Name
	if comp.Required == "Y" {
		fld.Required = true
	}
	return nil
}

type ComponentDef struct {
	Name       string      `xml:"name,attr"`
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

type Group struct {
	Name       string      `xml:"name,attr"`
	Required   string      `xml:"required,attr"`
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}
