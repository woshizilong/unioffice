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

type CT_PrintProperties struct {
	// Print Output
	PrnWhatAttr ST_PrintWhat
	// Print Color Mode
	ClrModeAttr ST_PrintColorMode
	// Print Hidden Slides
	HiddenSlidesAttr *bool
	// Scale to Fit Paper when printing
	ScaleToFitPaperAttr *bool
	// Frame slides when printing
	FrameSlidesAttr *bool
	ExtLst          *CT_ExtensionList
}

func NewCT_PrintProperties() *CT_PrintProperties {
	ret := &CT_PrintProperties{}
	return ret
}

func (m *CT_PrintProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PrnWhatAttr != ST_PrintWhatUnset {
		attr, err := m.PrnWhatAttr.MarshalXMLAttr(xml.Name{Local: "prnWhat"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ClrModeAttr != ST_PrintColorModeUnset {
		attr, err := m.ClrModeAttr.MarshalXMLAttr(xml.Name{Local: "clrMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HiddenSlidesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenSlides"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenSlidesAttr))})
	}
	if m.ScaleToFitPaperAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scaleToFitPaper"},
			Value: fmt.Sprintf("%d", b2i(*m.ScaleToFitPaperAttr))})
	}
	if m.FrameSlidesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "frameSlides"},
			Value: fmt.Sprintf("%d", b2i(*m.FrameSlidesAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PrintProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "prnWhat" {
			m.PrnWhatAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "clrMode" {
			m.ClrModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hiddenSlides" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenSlidesAttr = &parsed
		}
		if attr.Name.Local == "scaleToFitPaper" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ScaleToFitPaperAttr = &parsed
		}
		if attr.Name.Local == "frameSlides" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FrameSlidesAttr = &parsed
		}
	}
lCT_PrintProperties:
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
				log.Printf("skipping unsupported element on CT_PrintProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PrintProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PrintProperties and its children
func (m *CT_PrintProperties) Validate() error {
	return m.ValidateWithPath("CT_PrintProperties")
}

// ValidateWithPath validates the CT_PrintProperties and its children, prefixing error messages with path
func (m *CT_PrintProperties) ValidateWithPath(path string) error {
	if err := m.PrnWhatAttr.ValidateWithPath(path + "/PrnWhatAttr"); err != nil {
		return err
	}
	if err := m.ClrModeAttr.ValidateWithPath(path + "/ClrModeAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
