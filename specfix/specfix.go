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
	"log"
)

//go:embed MOEX44.xml
var MOEX44 []byte

// Spec represents a FIX messages based on a XML specification of the FIX protocol.
type Spec struct {
	XMLName     xml.Name          `xml:"fix"`
	Type        string            `xml:"type,attr"`
	Major       string            `xml:"major,attr"`
	Minor       string            `xml:"minor,attr"`
	Servicepack string            `xml:"servicepack,attr"`
	Header      Header            `xml:"header"`
	Messages    []Message         `xml:"messages>message"`
	Components  []ComponentDefine `xml:"components>component"`
	Fields      []FieldDefine     `xml:"fields>field"`
	Trailer     Trailer           `xml:"trailer"`
}

func (spec Spec) MemorizeComponents() (map[string]ComponentMemo, error) {
	fldByName := make(map[string]FieldDefine, len(spec.Fields))

	for _, fld := range spec.Fields {
		fldByName[fld.Name] = fld
	}

	compByName := make(map[string]ComponentMemo, len(spec.Components))

	// Populating the fields of the components until satisfying all
	// component-to-component references due to the ability
	// to follow components in an unsorted order.
	for i, incomplete := 0, true; incomplete; i++ {
		if i == 1000 {
			log.Printf("Too many attempts %d to fill in the fields of the components.", i)
			break
		}

		incomplete = false
		for _, comp := range spec.Components {
			if _, ok := compByName[comp.Name]; ok {
				continue
			}

			flds, ok := memorizeSpecComponent(comp, fldByName, compByName)
			if !ok {
				incomplete = true
				continue
			}

			compByName[comp.Name] = ComponentMemo{Name: comp.Name, Fields: flds}
		}
	}

	return compByName, nil
}

// memorizeSpecComponent transforms a component definition containing nesting to the flat fields.
func memorizeSpecComponent(src ComponentDefine, fldByName map[string]FieldDefine, compByName map[string]ComponentMemo) ([]FieldMemo, bool) {
	var flds []FieldMemo

	for _, fld := range src.Fields {
		fDef, ok := fldByName[fld.Name]
		if !ok {
			return nil, false
		}
		f := FieldMemo{Name: fld.Name, Required: fld.Required, Define: fDef}
		flds = append(flds, f)
	}

	for _, comp := range src.Components {
		cDef, ok := compByName[comp.Name]
		if !ok {
			return nil, false
		}
		for _, fld := range cDef.Fields {
			fDef, ok := fldByName[fld.Name]
			if !ok {
				return nil, false
			}
			f := FieldMemo{Name: fld.Name, Required: fld.Required, Define: fDef}
			flds = append(flds, f)
		}
	}

	for _, grp := range src.Groups {
		f, ok := memorizeSpecGroup(grp, fldByName, compByName)
		if !ok {
			return nil, false
		}
		flds = append(flds, f...)
	}

	return flds, true
}

// memorizeSpecGroup transforms a group containing nesting to the flat fields.
func memorizeSpecGroup(src Group, fldByName map[string]FieldDefine, compByName map[string]ComponentMemo) ([]FieldMemo, bool) {
	var flds []FieldMemo

	for _, fld := range src.Fields {
		fDef, ok := fldByName[fld.Name]
		if !ok {
			return nil, false
		}
		f := FieldMemo{Name: fld.Name, Required: fld.Required, Define: fDef}
		flds = append(flds, f)
	}

	for _, comp := range src.Components {
		cDef, ok := compByName[comp.Name]
		if !ok {
			return nil, false
		}
		for _, fld := range cDef.Fields {
			fDef, ok := fldByName[fld.Name]
			if !ok {
				return nil, false
			}
			f := FieldMemo{Name: fld.Name, Required: fld.Required, Define: fDef}
			flds = append(flds, f)
		}
	}

	for _, grp := range src.Groups {
		f, ok := memorizeSpecGroup(grp, fldByName, compByName)
		if !ok {
			return nil, false
		}
		flds = append(flds, f...)
	}

	return flds, true
}

type Header struct {
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

type Trailer struct {
	Fields []Field `xml:"field"`
}

type Message struct {
	Name       string      `xml:"name,attr"`
	MsgCat     string      `xml:"msgcat,attr"`
	MsgType    string      `xml:"msgtype,attr"`
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

// MemorizeFields all memorized fields of the message.
func (msg Message) MemorizeFields(hdr Header, trl Trailer, fldByName map[string]FieldDefine, compByName map[string]ComponentMemo) ([]FieldMemo, error) {
	// Transform a fields/components/groups specs to a fields.
	var flds, f []FieldMemo
	var err error

	f, err = memorizeMessageFields(hdr.Fields, fldByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	f, err = memorizeMessageComponents(hdr.Components, compByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	f, err = memorizeMessageGroups(hdr.Groups, fldByName, compByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	f, err = memorizeMessageFields(msg.Fields, fldByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	f, err = memorizeMessageComponents(msg.Components, compByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	f, err = memorizeMessageFields(trl.Fields, fldByName)
	if err != nil {
		return nil, err
	}
	flds = append(flds, f...)

	return flds, nil
}

// memorizeMessageFields transforms xml fields to template fields.
func memorizeMessageFields(src []Field, fldByName map[string]FieldDefine) ([]FieldMemo, error) {
	var flds []FieldMemo

	for _, fld := range src {
		d, ok := fldByName[fld.Name]
		if !ok {
			log.Printf("Missing field %q specification/definition.", fld.Name)
		}

		f := FieldMemo{Name: fld.Name, Required: fld.Required, Define: d}
		flds = append(flds, f)
	}

	return flds, nil
}

// memorizeMessageComponents transforms xml components to template fields.
func memorizeMessageComponents(src []Component, compByName map[string]ComponentMemo) ([]FieldMemo, error) {
	var flds []FieldMemo

	for _, comp := range src {
		c, ok := compByName[comp.Name]
		if !ok {
			log.Printf("Missing component %q specification/definition.", comp.Name)
			continue
		}
		flds = append(flds, c.Fields...)
	}

	return flds, nil
}

// memorizeMessageGroups transforms xml groups to template fields.
func memorizeMessageGroups(src []Group, fldByName map[string]FieldDefine, compByName map[string]ComponentMemo) ([]FieldMemo, error) {
	var flds, f []FieldMemo
	var err error

	for _, group := range src {
		f, err = memorizeMessageFields(group.Fields, fldByName)
		if err != nil {
			return nil, err
		}
		flds = append(flds, f...)

		f, err = memorizeMessageComponents(group.Components, compByName)
		if err != nil {
			return nil, err
		}
		flds = append(flds, f...)

		f, err = memorizeMessageGroups(group.Groups, fldByName, compByName)
		if err != nil {
			return nil, err
		}
		flds = append(flds, f...)
	}

	return flds, nil
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

type FieldDefine struct {
	Number int
	Name   string
	Type   string
	MinLen int
	MaxLen int
	Values []Value
}

func (fld *FieldDefine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type field struct {
		XMLName xml.Name `xml:"field"`
		Number  int      `xml:"number,attr"`
		Name    string   `xml:"name,attr"`
		Type    string   `xml:"type,attr"`
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
	fld.MinLen = f.MinLen
	fld.MaxLen = f.MaxLen
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

type ComponentDefine struct {
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

// FieldMemo is a denormalized/unnormalized form of the field.
type FieldMemo struct {
	Name     string
	Required bool
	Define   FieldDefine
}

// ComponentMemo is a denormalized/unnormalized form of the component.
type ComponentMemo struct {
	Name     string
	Required bool
	Fields   []FieldMemo
}
