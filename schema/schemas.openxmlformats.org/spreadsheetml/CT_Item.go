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
	"strconv"
)

type CT_Item struct {
	// Item User Caption
	NAttr *string
	// Item Type
	TAttr ST_ItemType
	// Hidden
	HAttr *bool
	// Character
	SAttr *bool
	// Hide Details
	SdAttr *bool
	// Calculated Member
	FAttr *bool
	// Missing
	MAttr *bool
	// Child Items
	CAttr *bool
	// Item Index
	XAttr *uint32
	// Expanded
	DAttr *bool
	// Drill Across Attributes
	EAttr *bool
}

func NewCT_Item() *CT_Item {
	ret := &CT_Item{}
	return ret
}

func (m *CT_Item) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
			Value: fmt.Sprintf("%v", *m.NAttr)})
	}
	if m.TAttr != ST_ItemTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "h"},
			Value: fmt.Sprintf("%d", b2i(*m.HAttr))})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%d", b2i(*m.SAttr))})
	}
	if m.SdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sd"},
			Value: fmt.Sprintf("%d", b2i(*m.SdAttr))})
	}
	if m.FAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "f"},
			Value: fmt.Sprintf("%d", b2i(*m.FAttr))})
	}
	if m.MAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "m"},
			Value: fmt.Sprintf("%d", b2i(*m.MAttr))})
	}
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%d", b2i(*m.CAttr))})
	}
	if m.XAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
			Value: fmt.Sprintf("%v", *m.XAttr)})
	}
	if m.DAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "d"},
			Value: fmt.Sprintf("%d", b2i(*m.DAttr))})
	}
	if m.EAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "e"},
			Value: fmt.Sprintf("%d", b2i(*m.EAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NAttr = &parsed
		}
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "h" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HAttr = &parsed
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = &parsed
		}
		if attr.Name.Local == "sd" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SdAttr = &parsed
		}
		if attr.Name.Local == "f" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FAttr = &parsed
		}
		if attr.Name.Local == "m" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MAttr = &parsed
		}
		if attr.Name.Local == "c" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CAttr = &parsed
		}
		if attr.Name.Local == "x" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.XAttr = &pt
		}
		if attr.Name.Local == "d" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DAttr = &parsed
		}
		if attr.Name.Local == "e" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Item: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Item and its children
func (m *CT_Item) Validate() error {
	return m.ValidateWithPath("CT_Item")
}

// ValidateWithPath validates the CT_Item and its children, prefixing error messages with path
func (m *CT_Item) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	return nil
}
