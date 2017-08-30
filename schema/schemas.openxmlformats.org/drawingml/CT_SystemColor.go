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
)

type CT_SystemColor struct {
	ValAttr           ST_SystemColorVal
	LastClrAttr       *string
	EG_ColorTransform []*EG_ColorTransform
}

func NewCT_SystemColor() *CT_SystemColor {
	ret := &CT_SystemColor{}
	ret.ValAttr = ST_SystemColorVal(1)
	return ret
}
func (m *CT_SystemColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LastClrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lastClr"},
			Value: fmt.Sprintf("%v", *m.LastClrAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.EG_ColorTransform != nil {
		for _, c := range m.EG_ColorTransform {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SystemColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_SystemColorVal(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "lastClr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastClrAttr = &parsed
		}
	}
lCT_SystemColor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tint":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Tint = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Tint, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "shade":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Shade = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Shade, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "comp":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Comp = NewCT_ComplementTransform()
				if err := d.DecodeElement(tmpcolortransform.Comp, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "inv":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Inv = NewCT_InverseTransform()
				if err := d.DecodeElement(tmpcolortransform.Inv, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "gray":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Gray = NewCT_GrayscaleTransform()
				if err := d.DecodeElement(tmpcolortransform.Gray, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "alpha":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Alpha = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Alpha, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "alphaOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.AlphaOff = NewCT_FixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.AlphaOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "alphaMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.AlphaMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(tmpcolortransform.AlphaMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "hue":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Hue = NewCT_PositiveFixedAngle()
				if err := d.DecodeElement(tmpcolortransform.Hue, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "hueOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.HueOff = NewCT_Angle()
				if err := d.DecodeElement(tmpcolortransform.HueOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "hueMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.HueMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(tmpcolortransform.HueMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "sat":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Sat = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Sat, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "satOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.SatOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.SatOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "satMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.SatMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.SatMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "lum":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Lum = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Lum, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "lumOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.LumOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.LumOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "lumMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.LumMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.LumMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "red":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Red = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Red, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "redOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.RedOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.RedOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "redMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.RedMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.RedMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "green":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Green = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Green, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "greenOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.GreenOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.GreenOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "greenMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.GreenMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.GreenMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "blue":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Blue = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Blue, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "blueOff":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.BlueOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.BlueOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "blueMod":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.BlueMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.BlueMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "gamma":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Gamma = NewCT_GammaTransform()
				if err := d.DecodeElement(tmpcolortransform.Gamma, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case "invGamma":
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.InvGamma = NewCT_InverseGammaTransform()
				if err := d.DecodeElement(tmpcolortransform.InvGamma, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SystemColor
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SystemColor) Validate() error {
	return m.ValidateWithPath("CT_SystemColor")
}
func (m *CT_SystemColor) ValidateWithPath(path string) error {
	if m.ValAttr == ST_SystemColorValUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	for i, v := range m.EG_ColorTransform {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ColorTransform[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}