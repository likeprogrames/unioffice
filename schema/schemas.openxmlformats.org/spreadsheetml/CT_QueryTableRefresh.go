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

type CT_QueryTableRefresh struct {
	// Preserve Sort & Filter Layout
	PreserveSortFilterLayoutAttr *bool
	// Next Field Id Wrapped
	FieldIdWrappedAttr *bool
	// Headers In Last Refresh
	HeadersInLastRefreshAttr *bool
	// Minimum Refresh Version
	MinimumVersionAttr *uint8
	// Next field id
	NextIdAttr *uint32
	// Columns Left
	UnboundColumnsLeftAttr *uint32
	// Columns Right
	UnboundColumnsRightAttr *uint32
	// Query table fields
	QueryTableFields *CT_QueryTableFields
	// Deleted Fields
	QueryTableDeletedFields *CT_QueryTableDeletedFields
	// Sort State
	SortState *CT_SortState
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_QueryTableRefresh() *CT_QueryTableRefresh {
	ret := &CT_QueryTableRefresh{}
	ret.QueryTableFields = NewCT_QueryTableFields()
	return ret
}
func (m *CT_QueryTableRefresh) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PreserveSortFilterLayoutAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveSortFilterLayout"},
			Value: fmt.Sprintf("%v", *m.PreserveSortFilterLayoutAttr)})
	}
	if m.FieldIdWrappedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldIdWrapped"},
			Value: fmt.Sprintf("%v", *m.FieldIdWrappedAttr)})
	}
	if m.HeadersInLastRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headersInLastRefresh"},
			Value: fmt.Sprintf("%v", *m.HeadersInLastRefreshAttr)})
	}
	if m.MinimumVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minimumVersion"},
			Value: fmt.Sprintf("%v", *m.MinimumVersionAttr)})
	}
	if m.NextIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nextId"},
			Value: fmt.Sprintf("%v", *m.NextIdAttr)})
	}
	if m.UnboundColumnsLeftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unboundColumnsLeft"},
			Value: fmt.Sprintf("%v", *m.UnboundColumnsLeftAttr)})
	}
	if m.UnboundColumnsRightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unboundColumnsRight"},
			Value: fmt.Sprintf("%v", *m.UnboundColumnsRightAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sequeryTableFields := xml.StartElement{Name: xml.Name{Local: "x:queryTableFields"}}
	e.EncodeElement(m.QueryTableFields, sequeryTableFields)
	if m.QueryTableDeletedFields != nil {
		sequeryTableDeletedFields := xml.StartElement{Name: xml.Name{Local: "x:queryTableDeletedFields"}}
		e.EncodeElement(m.QueryTableDeletedFields, sequeryTableDeletedFields)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "x:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_QueryTableRefresh) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.QueryTableFields = NewCT_QueryTableFields()
	for _, attr := range start.Attr {
		if attr.Name.Local == "preserveSortFilterLayout" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveSortFilterLayoutAttr = &parsed
		}
		if attr.Name.Local == "fieldIdWrapped" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldIdWrappedAttr = &parsed
		}
		if attr.Name.Local == "headersInLastRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadersInLastRefreshAttr = &parsed
		}
		if attr.Name.Local == "minimumVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinimumVersionAttr = &pt
		}
		if attr.Name.Local == "nextId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NextIdAttr = &pt
		}
		if attr.Name.Local == "unboundColumnsLeft" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.UnboundColumnsLeftAttr = &pt
		}
		if attr.Name.Local == "unboundColumnsRight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.UnboundColumnsRightAttr = &pt
		}
	}
lCT_QueryTableRefresh:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "queryTableFields":
				if err := d.DecodeElement(m.QueryTableFields, &el); err != nil {
					return err
				}
			case "queryTableDeletedFields":
				m.QueryTableDeletedFields = NewCT_QueryTableDeletedFields()
				if err := d.DecodeElement(m.QueryTableDeletedFields, &el); err != nil {
					return err
				}
			case "sortState":
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
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
			break lCT_QueryTableRefresh
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_QueryTableRefresh) Validate() error {
	return m.ValidateWithPath("CT_QueryTableRefresh")
}
func (m *CT_QueryTableRefresh) ValidateWithPath(path string) error {
	if err := m.QueryTableFields.ValidateWithPath(path + "/QueryTableFields"); err != nil {
		return err
	}
	if m.QueryTableDeletedFields != nil {
		if err := m.QueryTableDeletedFields.ValidateWithPath(path + "/QueryTableDeletedFields"); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}