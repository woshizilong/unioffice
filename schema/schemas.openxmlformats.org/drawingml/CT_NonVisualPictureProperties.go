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

type CT_NonVisualPictureProperties struct {
	PreferRelativeResizeAttr *bool
	PicLocks                 *CT_PictureLocking
	ExtLst                   *CT_OfficeArtExtensionList
}

func NewCT_NonVisualPictureProperties() *CT_NonVisualPictureProperties {
	ret := &CT_NonVisualPictureProperties{}
	return ret
}

func (m *CT_NonVisualPictureProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PreferRelativeResizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferRelativeResize"},
			Value: fmt.Sprintf("%d", b2i(*m.PreferRelativeResizeAttr))})
	}
	e.EncodeToken(start)
	if m.PicLocks != nil {
		sepicLocks := xml.StartElement{Name: xml.Name{Local: "a:picLocks"}}
		e.EncodeElement(m.PicLocks, sepicLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NonVisualPictureProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "preferRelativeResize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreferRelativeResizeAttr = &parsed
		}
	}
lCT_NonVisualPictureProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "picLocks":
				m.PicLocks = NewCT_PictureLocking()
				if err := d.DecodeElement(m.PicLocks, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NonVisualPictureProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualPictureProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NonVisualPictureProperties and its children
func (m *CT_NonVisualPictureProperties) Validate() error {
	return m.ValidateWithPath("CT_NonVisualPictureProperties")
}

// ValidateWithPath validates the CT_NonVisualPictureProperties and its children, prefixing error messages with path
func (m *CT_NonVisualPictureProperties) ValidateWithPath(path string) error {
	if m.PicLocks != nil {
		if err := m.PicLocks.ValidateWithPath(path + "/PicLocks"); err != nil {
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
