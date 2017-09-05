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

type CT_MetadataStringIndex struct {
	// Index Value
	XAttr uint32
	// String is a Set
	SAttr *bool
}

func NewCT_MetadataStringIndex() *CT_MetadataStringIndex {
	ret := &CT_MetadataStringIndex{}
	return ret
}

func (m *CT_MetadataStringIndex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%d", b2i(*m.SAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MetadataStringIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.XAttr = uint32(parsed)
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MetadataStringIndex: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MetadataStringIndex and its children
func (m *CT_MetadataStringIndex) Validate() error {
	return m.ValidateWithPath("CT_MetadataStringIndex")
}

// ValidateWithPath validates the CT_MetadataStringIndex and its children, prefixing error messages with path
func (m *CT_MetadataStringIndex) ValidateWithPath(path string) error {
	return nil
}
