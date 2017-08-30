// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
)

type RFC3066 struct {
}

func NewRFC3066() *RFC3066 {
	ret := &RFC3066{}
	return ret
}
func (m *RFC3066) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "RFC3066"
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *RFC3066) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing RFC3066: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *RFC3066) Validate() error {
	return m.ValidateWithPath("RFC3066")
}
func (m *RFC3066) ValidateWithPath(path string) error {
	return nil
}