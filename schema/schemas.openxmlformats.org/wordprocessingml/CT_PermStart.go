// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PermStart struct {
	EdGrpAttr    ST_EdGrp
	EdAttr       *string
	ColFirstAttr *int32
	ColLastAttr  *int32
	// Annotation ID
	IdAttr string
	// Annotation Displaced By Custom XML Markup
	DisplacedByCustomXmlAttr ST_DisplacedByCustomXml
}

func NewCT_PermStart() *CT_PermStart {
	ret := &CT_PermStart{}
	return ret
}
func (m *CT_PermStart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.EdGrpAttr != ST_EdGrpUnset {
		attr, err := m.EdGrpAttr.MarshalXMLAttr(xml.Name{Local: "w:edGrp"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:ed"},
			Value: fmt.Sprintf("%v", *m.EdAttr)})
	}
	if m.ColFirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colFirst"},
			Value: fmt.Sprintf("%v", *m.ColFirstAttr)})
	}
	if m.ColLastAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colLast"},
			Value: fmt.Sprintf("%v", *m.ColLastAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.DisplacedByCustomXmlAttr != ST_DisplacedByCustomXmlUnset {
		attr, err := m.DisplacedByCustomXmlAttr.MarshalXMLAttr(xml.Name{Local: "w:displacedByCustomXml"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PermStart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "edGrp" {
			m.EdGrpAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ed" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EdAttr = &parsed
		}
		if attr.Name.Local == "colFirst" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ColFirstAttr = &pt
		}
		if attr.Name.Local == "colLast" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ColLastAttr = &pt
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
		if attr.Name.Local == "displacedByCustomXml" {
			m.DisplacedByCustomXmlAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PermStart: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_PermStart) Validate() error {
	return m.ValidateWithPath("CT_PermStart")
}
func (m *CT_PermStart) ValidateWithPath(path string) error {
	if err := m.EdGrpAttr.ValidateWithPath(path + "/EdGrpAttr"); err != nil {
		return err
	}
	if err := m.DisplacedByCustomXmlAttr.ValidateWithPath(path + "/DisplacedByCustomXmlAttr"); err != nil {
		return err
	}
	return nil
}