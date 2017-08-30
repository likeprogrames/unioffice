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

type CT_ChartsheetView struct {
	// Sheet Tab Selected
	TabSelectedAttr *bool
	// Window Zoom Scale
	ZoomScaleAttr *uint32
	// Workbook View Id
	WorkbookViewIdAttr uint32
	// Zoom To Fit
	ZoomToFitAttr *bool
	ExtLst        *CT_ExtensionList
}

func NewCT_ChartsheetView() *CT_ChartsheetView {
	ret := &CT_ChartsheetView{}
	return ret
}
func (m *CT_ChartsheetView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TabSelectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tabSelected"},
			Value: fmt.Sprintf("%v", *m.TabSelectedAttr)})
	}
	if m.ZoomScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScale"},
			Value: fmt.Sprintf("%v", *m.ZoomScaleAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookViewId"},
		Value: fmt.Sprintf("%v", m.WorkbookViewIdAttr)})
	if m.ZoomToFitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomToFit"},
			Value: fmt.Sprintf("%v", *m.ZoomToFitAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ChartsheetView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tabSelected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TabSelectedAttr = &parsed
		}
		if attr.Name.Local == "zoomScale" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScaleAttr = &pt
		}
		if attr.Name.Local == "workbookViewId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.WorkbookViewIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "zoomToFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZoomToFitAttr = &parsed
		}
	}
lCT_ChartsheetView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetView
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ChartsheetView) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetView")
}
func (m *CT_ChartsheetView) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}