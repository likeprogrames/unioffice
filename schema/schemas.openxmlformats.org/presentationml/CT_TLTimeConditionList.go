// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TLTimeConditionList struct {
	// Condition
	Cond []*CT_TLTimeCondition
}

func NewCT_TLTimeConditionList() *CT_TLTimeConditionList {
	ret := &CT_TLTimeConditionList{}
	return ret
}
func (m *CT_TLTimeConditionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	second := xml.StartElement{Name: xml.Name{Local: "p:cond"}}
	e.EncodeElement(m.Cond, second)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLTimeConditionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTimeConditionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cond":
				tmp := NewCT_TLTimeCondition()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cond = append(m.Cond, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeConditionList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLTimeConditionList) Validate() error {
	return m.ValidateWithPath("CT_TLTimeConditionList")
}
func (m *CT_TLTimeConditionList) ValidateWithPath(path string) error {
	for i, v := range m.Cond {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cond[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}