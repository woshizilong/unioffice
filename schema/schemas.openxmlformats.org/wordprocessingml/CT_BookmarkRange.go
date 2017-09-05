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

type CT_BookmarkRange struct {
	ColFirstAttr             *int64
	ColLastAttr              *int64
	DisplacedByCustomXmlAttr ST_DisplacedByCustomXml
	// Annotation Identifier
	IdAttr int64
}

func NewCT_BookmarkRange() *CT_BookmarkRange {
	ret := &CT_BookmarkRange{}
	return ret
}

func (m *CT_BookmarkRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ColFirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colFirst"},
			Value: fmt.Sprintf("%v", *m.ColFirstAttr)})
	}
	if m.ColLastAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colLast"},
			Value: fmt.Sprintf("%v", *m.ColLastAttr)})
	}
	if m.DisplacedByCustomXmlAttr != ST_DisplacedByCustomXmlUnset {
		attr, err := m.DisplacedByCustomXmlAttr.MarshalXMLAttr(xml.Name{Local: "w:displacedByCustomXml"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BookmarkRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "colFirst" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ColFirstAttr = &parsed
		}
		if attr.Name.Local == "colLast" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ColLastAttr = &parsed
		}
		if attr.Name.Local == "displacedByCustomXml" {
			m.DisplacedByCustomXmlAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BookmarkRange: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_BookmarkRange and its children
func (m *CT_BookmarkRange) Validate() error {
	return m.ValidateWithPath("CT_BookmarkRange")
}

// ValidateWithPath validates the CT_BookmarkRange and its children, prefixing error messages with path
func (m *CT_BookmarkRange) ValidateWithPath(path string) error {
	if err := m.DisplacedByCustomXmlAttr.ValidateWithPath(path + "/DisplacedByCustomXmlAttr"); err != nil {
		return err
	}
	return nil
}
