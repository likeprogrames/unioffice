// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_QueryTableFields struct {
	// Column Count
	CountAttr *uint32
	// QueryTable Field
	QueryTableField []*CT_QueryTableField
}

func NewCT_QueryTableFields() *CT_QueryTableFields {
	ret := &CT_QueryTableFields{}
	return ret
}
func (m *CT_QueryTableFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.QueryTableField != nil {
		sequeryTableField := xml.StartElement{Name: xml.Name{Local: "x:queryTableField"}}
		e.EncodeElement(m.QueryTableField, sequeryTableField)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_QueryTableFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_QueryTableFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "queryTableField":
				tmp := NewCT_QueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.QueryTableField = append(m.QueryTableField, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryTableFields
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_QueryTableFields) Validate() error {
	return m.ValidateWithPath("CT_QueryTableFields")
}
func (m *CT_QueryTableFields) ValidateWithPath(path string) error {
	for i, v := range m.QueryTableField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/QueryTableField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}