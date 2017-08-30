// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_Authors struct {
	// Author
	Author []string
}

func NewCT_Authors() *CT_Authors {
	ret := &CT_Authors{}
	return ret
}
func (m *CT_Authors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Author != nil {
		seauthor := xml.StartElement{Name: xml.Name{Local: "x:author"}}
		e.EncodeElement(m.Author, seauthor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Authors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Authors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "author":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Author = append(m.Author, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Authors
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Authors) Validate() error {
	return m.ValidateWithPath("CT_Authors")
}
func (m *CT_Authors) ValidateWithPath(path string) error {
	return nil
}