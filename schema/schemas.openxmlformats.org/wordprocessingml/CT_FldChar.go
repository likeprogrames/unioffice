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
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_FldChar struct {
	// Field Character Type
	FldCharTypeAttr ST_FldCharType
	// Field Should Not Be Recalculated
	FldLockAttr *sharedTypes.ST_OnOff
	// Field Result Invalidated
	DirtyAttr *sharedTypes.ST_OnOff
	// Custom Field Data
	FldData *CT_Text
	// Form Field Properties
	FfData *CT_FFData
	// Previous Numbering Field Properties
	NumberingChange *CT_TrackChangeNumbering
}

func NewCT_FldChar() *CT_FldChar {
	ret := &CT_FldChar{}
	ret.FldCharTypeAttr = ST_FldCharType(1)
	return ret
}
func (m *CT_FldChar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.FldCharTypeAttr.MarshalXMLAttr(xml.Name{Local: "w:fldCharType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.FldLockAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fldLock"},
			Value: fmt.Sprintf("%v", *m.FldLockAttr)})
	}
	if m.DirtyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dirty"},
			Value: fmt.Sprintf("%v", *m.DirtyAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.FldData != nil {
		sefldData := xml.StartElement{Name: xml.Name{Local: "w:fldData"}}
		e.EncodeElement(m.FldData, sefldData)
	}
	if m.FfData != nil {
		seffData := xml.StartElement{Name: xml.Name{Local: "w:ffData"}}
		e.EncodeElement(m.FfData, seffData)
	}
	if m.NumberingChange != nil {
		senumberingChange := xml.StartElement{Name: xml.Name{Local: "w:numberingChange"}}
		e.EncodeElement(m.NumberingChange, senumberingChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FldChar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FldCharTypeAttr = ST_FldCharType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "fldCharType" {
			m.FldCharTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "fldLock" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FldLockAttr = &parsed
		}
		if attr.Name.Local == "dirty" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DirtyAttr = &parsed
		}
	}
lCT_FldChar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fldData":
				m.FldData = NewCT_Text()
				if err := d.DecodeElement(m.FldData, &el); err != nil {
					return err
				}
			case "ffData":
				m.FfData = NewCT_FFData()
				if err := d.DecodeElement(m.FfData, &el); err != nil {
					return err
				}
			case "numberingChange":
				m.NumberingChange = NewCT_TrackChangeNumbering()
				if err := d.DecodeElement(m.NumberingChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FldChar
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FldChar) Validate() error {
	return m.ValidateWithPath("CT_FldChar")
}
func (m *CT_FldChar) ValidateWithPath(path string) error {
	if m.FldCharTypeAttr == ST_FldCharTypeUnset {
		return fmt.Errorf("%s/FldCharTypeAttr is a mandatory field", path)
	}
	if err := m.FldCharTypeAttr.ValidateWithPath(path + "/FldCharTypeAttr"); err != nil {
		return err
	}
	if m.FldLockAttr != nil {
		if err := m.FldLockAttr.ValidateWithPath(path + "/FldLockAttr"); err != nil {
			return err
		}
	}
	if m.DirtyAttr != nil {
		if err := m.DirtyAttr.ValidateWithPath(path + "/DirtyAttr"); err != nil {
			return err
		}
	}
	if m.FldData != nil {
		if err := m.FldData.ValidateWithPath(path + "/FldData"); err != nil {
			return err
		}
	}
	if m.FfData != nil {
		if err := m.FfData.ValidateWithPath(path + "/FfData"); err != nil {
			return err
		}
	}
	if m.NumberingChange != nil {
		if err := m.NumberingChange.ValidateWithPath(path + "/NumberingChange"); err != nil {
			return err
		}
	}
	return nil
}