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
)

type CT_TextBulletSizePercent struct {
	ValAttr string
}

func NewCT_TextBulletSizePercent() *CT_TextBulletSizePercent {
	ret := &CT_TextBulletSizePercent{}
	ret.ValAttr = "100%"
	return ret
}
func (m *CT_TextBulletSizePercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextBulletSizePercent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = "100%"
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletSizePercent: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TextBulletSizePercent) Validate() error {
	return m.ValidateWithPath("CT_TextBulletSizePercent")
}
func (m *CT_TextBulletSizePercent) ValidateWithPath(path string) error {
	if !ST_TextBulletSizePercentPatternRe.MatchString(m.ValAttr) {
		return fmt.Errorf(`%s/m.ValAttr must match '%s' (have %v)`, path, ST_TextBulletSizePercentPatternRe, m.ValAttr)
	}
	return nil
}