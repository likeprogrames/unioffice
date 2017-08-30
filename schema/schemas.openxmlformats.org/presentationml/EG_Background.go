// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type EG_Background struct {
	// Background Properties
	BgPr *CT_BackgroundProperties
	// Background Style Reference
	BgRef *drawingml.CT_StyleMatrixReference
}

func NewEG_Background() *EG_Background {
	ret := &EG_Background{}
	return ret
}
func (m *EG_Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BgPr != nil {
		sebgPr := xml.StartElement{Name: xml.Name{Local: "p:bgPr"}}
		e.EncodeElement(m.BgPr, sebgPr)
	}
	if m.BgRef != nil {
		sebgRef := xml.StartElement{Name: xml.Name{Local: "p:bgRef"}}
		e.EncodeElement(m.BgRef, sebgRef)
	}
	return nil
}
func (m *EG_Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Background:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bgPr":
				m.BgPr = NewCT_BackgroundProperties()
				if err := d.DecodeElement(m.BgPr, &el); err != nil {
					return err
				}
			case "bgRef":
				m.BgRef = drawingml.NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.BgRef, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Background
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_Background) Validate() error {
	return m.ValidateWithPath("EG_Background")
}
func (m *EG_Background) ValidateWithPath(path string) error {
	if m.BgPr != nil {
		if err := m.BgPr.ValidateWithPath(path + "/BgPr"); err != nil {
			return err
		}
	}
	if m.BgRef != nil {
		if err := m.BgRef.ValidateWithPath(path + "/BgRef"); err != nil {
			return err
		}
	}
	return nil
}