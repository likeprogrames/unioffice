// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_Path2D struct {
	WAttr           *int64
	HAttr           *int64
	FillAttr        ST_PathFillMode
	StrokeAttr      *bool
	ExtrusionOkAttr *bool
	Close           []*CT_Path2DClose
	MoveTo          []*CT_Path2DMoveTo
	LnTo            []*CT_Path2DLineTo
	ArcTo           []*CT_Path2DArcTo
	QuadBezTo       []*CT_Path2DQuadBezierTo
	CubicBezTo      []*CT_Path2DCubicBezierTo
}

func NewCT_Path2D() *CT_Path2D {
	ret := &CT_Path2D{}
	return ret
}
func (m *CT_Path2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.HAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "h"},
			Value: fmt.Sprintf("%v", *m.HAttr)})
	}
	if m.FillAttr != ST_PathFillModeUnset {
		attr, err := m.FillAttr.MarshalXMLAttr(xml.Name{Local: "fill"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stroke"},
			Value: fmt.Sprintf("%v", *m.StrokeAttr)})
	}
	if m.ExtrusionOkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "extrusionOk"},
			Value: fmt.Sprintf("%v", *m.ExtrusionOkAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Close != nil {
		seclose := xml.StartElement{Name: xml.Name{Local: "a:close"}}
		e.EncodeElement(m.Close, seclose)
	}
	if m.MoveTo != nil {
		semoveTo := xml.StartElement{Name: xml.Name{Local: "a:moveTo"}}
		e.EncodeElement(m.MoveTo, semoveTo)
	}
	if m.LnTo != nil {
		selnTo := xml.StartElement{Name: xml.Name{Local: "a:lnTo"}}
		e.EncodeElement(m.LnTo, selnTo)
	}
	if m.ArcTo != nil {
		searcTo := xml.StartElement{Name: xml.Name{Local: "a:arcTo"}}
		e.EncodeElement(m.ArcTo, searcTo)
	}
	if m.QuadBezTo != nil {
		sequadBezTo := xml.StartElement{Name: xml.Name{Local: "a:quadBezTo"}}
		e.EncodeElement(m.QuadBezTo, sequadBezTo)
	}
	if m.CubicBezTo != nil {
		secubicBezTo := xml.StartElement{Name: xml.Name{Local: "a:cubicBezTo"}}
		e.EncodeElement(m.CubicBezTo, secubicBezTo)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Path2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
		if attr.Name.Local == "h" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.HAttr = &parsed
		}
		if attr.Name.Local == "fill" {
			m.FillAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "stroke" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StrokeAttr = &parsed
		}
		if attr.Name.Local == "extrusionOk" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ExtrusionOkAttr = &parsed
		}
	}
lCT_Path2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "close":
				tmp := NewCT_Path2DClose()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Close = append(m.Close, tmp)
			case "moveTo":
				tmp := NewCT_Path2DMoveTo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.MoveTo = append(m.MoveTo, tmp)
			case "lnTo":
				tmp := NewCT_Path2DLineTo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LnTo = append(m.LnTo, tmp)
			case "arcTo":
				tmp := NewCT_Path2DArcTo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ArcTo = append(m.ArcTo, tmp)
			case "quadBezTo":
				tmp := NewCT_Path2DQuadBezierTo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.QuadBezTo = append(m.QuadBezTo, tmp)
			case "cubicBezTo":
				tmp := NewCT_Path2DCubicBezierTo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CubicBezTo = append(m.CubicBezTo, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2D
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Path2D) Validate() error {
	return m.ValidateWithPath("CT_Path2D")
}
func (m *CT_Path2D) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if *m.WAttr < 0 {
			return fmt.Errorf("%s/m.WAttr must be >= 0 (have %v)", path, *m.WAttr)
		}
		if *m.WAttr > 27273042316900 {
			return fmt.Errorf("%s/m.WAttr must be <= 27273042316900 (have %v)", path, *m.WAttr)
		}
	}
	if m.HAttr != nil {
		if *m.HAttr < 0 {
			return fmt.Errorf("%s/m.HAttr must be >= 0 (have %v)", path, *m.HAttr)
		}
		if *m.HAttr > 27273042316900 {
			return fmt.Errorf("%s/m.HAttr must be <= 27273042316900 (have %v)", path, *m.HAttr)
		}
	}
	if err := m.FillAttr.ValidateWithPath(path + "/FillAttr"); err != nil {
		return err
	}
	for i, v := range m.Close {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Close[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.MoveTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/MoveTo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.LnTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LnTo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ArcTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ArcTo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.QuadBezTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/QuadBezTo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CubicBezTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CubicBezTo[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}