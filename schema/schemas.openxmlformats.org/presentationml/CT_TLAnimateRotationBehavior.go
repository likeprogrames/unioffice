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
	"strconv"
)

type CT_TLAnimateRotationBehavior struct {
	// By
	ByAttr *int32
	// From
	FromAttr *int32
	// To
	ToAttr *int32
	CBhvr  *CT_TLCommonBehaviorData
}

func NewCT_TLAnimateRotationBehavior() *CT_TLAnimateRotationBehavior {
	ret := &CT_TLAnimateRotationBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}
func (m *CT_TLAnimateRotationBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "by"},
			Value: fmt.Sprintf("%v", *m.ByAttr)})
	}
	if m.FromAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "from"},
			Value: fmt.Sprintf("%v", *m.FromAttr)})
	}
	if m.ToAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "to"},
			Value: fmt.Sprintf("%v", *m.ToAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLAnimateRotationBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "by" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ByAttr = &pt
		}
		if attr.Name.Local == "from" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.FromAttr = &pt
		}
		if attr.Name.Local == "to" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ToAttr = &pt
		}
	}
lCT_TLAnimateRotationBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cBhvr":
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateRotationBehavior
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLAnimateRotationBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateRotationBehavior")
}
func (m *CT_TLAnimateRotationBehavior) ValidateWithPath(path string) error {
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	return nil
}