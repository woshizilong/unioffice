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

type CT_Placeholder struct {
	// Placeholder Type
	TypeAttr ST_PlaceholderType
	// Placeholder Orientation
	OrientAttr ST_Direction
	// Placeholder Size
	SzAttr ST_PlaceholderSize
	// Placeholder Index
	IdxAttr *uint32
	// Placeholder has custom prompt
	HasCustomPromptAttr *bool
	ExtLst              *CT_ExtensionListModify
}

func NewCT_Placeholder() *CT_Placeholder {
	ret := &CT_Placeholder{}
	return ret
}

func (m *CT_Placeholder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_PlaceholderTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OrientAttr != ST_DirectionUnset {
		attr, err := m.OrientAttr.MarshalXMLAttr(xml.Name{Local: "orient"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SzAttr != ST_PlaceholderSizeUnset {
		attr, err := m.SzAttr.MarshalXMLAttr(xml.Name{Local: "sz"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
			Value: fmt.Sprintf("%v", *m.IdxAttr)})
	}
	if m.HasCustomPromptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hasCustomPrompt"},
			Value: fmt.Sprintf("%d", b2i(*m.HasCustomPromptAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Placeholder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "orient" {
			m.OrientAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "sz" {
			m.SzAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IdxAttr = &pt
		}
		if attr.Name.Local == "hasCustomPrompt" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HasCustomPromptAttr = &parsed
		}
	}
lCT_Placeholder:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Placeholder %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Placeholder
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Placeholder and its children
func (m *CT_Placeholder) Validate() error {
	return m.ValidateWithPath("CT_Placeholder")
}

// ValidateWithPath validates the CT_Placeholder and its children, prefixing error messages with path
func (m *CT_Placeholder) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.OrientAttr.ValidateWithPath(path + "/OrientAttr"); err != nil {
		return err
	}
	if err := m.SzAttr.ValidateWithPath(path + "/SzAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
