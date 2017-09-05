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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Slide struct {
	// Show Slide in Slide Show
	ShowAttr *bool
	// Common slide data for slides
	CSld *CT_CommonSlideData
	// Color Scheme Map Override
	ClrMapOvr *drawingml.CT_ColorMappingOverride
	// Slide Transition
	Transition *CT_SlideTransition
	// Slide Timing Information for a Slide
	Timing               *CT_SlideTiming
	ExtLst               *CT_ExtensionListModify
	ShowMasterSpAttr     *bool
	ShowMasterPhAnimAttr *bool
}

func NewCT_Slide() *CT_Slide {
	ret := &CT_Slide{}
	ret.CSld = NewCT_CommonSlideData()
	return ret
}

func (m *CT_Slide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "show"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowAttr))})
	}
	if m.ShowMasterSpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMasterSp"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowMasterSpAttr))})
	}
	if m.ShowMasterPhAnimAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMasterPhAnim"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowMasterPhAnimAttr))})
	}
	e.EncodeToken(start)
	secSld := xml.StartElement{Name: xml.Name{Local: "p:cSld"}}
	e.EncodeElement(m.CSld, secSld)
	if m.ClrMapOvr != nil {
		seclrMapOvr := xml.StartElement{Name: xml.Name{Local: "p:clrMapOvr"}}
		e.EncodeElement(m.ClrMapOvr, seclrMapOvr)
	}
	if m.Transition != nil {
		setransition := xml.StartElement{Name: xml.Name{Local: "p:transition"}}
		e.EncodeElement(m.Transition, setransition)
	}
	if m.Timing != nil {
		setiming := xml.StartElement{Name: xml.Name{Local: "p:timing"}}
		e.EncodeElement(m.Timing, setiming)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Slide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSld = NewCT_CommonSlideData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "show" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAttr = &parsed
		}
		if attr.Name.Local == "showMasterSp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterSpAttr = &parsed
		}
		if attr.Name.Local == "showMasterPhAnim" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterPhAnimAttr = &parsed
		}
	}
lCT_Slide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cSld":
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case "clrMapOvr":
				m.ClrMapOvr = drawingml.NewCT_ColorMappingOverride()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			case "transition":
				m.Transition = NewCT_SlideTransition()
				if err := d.DecodeElement(m.Transition, &el); err != nil {
					return err
				}
			case "timing":
				m.Timing = NewCT_SlideTiming()
				if err := d.DecodeElement(m.Timing, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Slide %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Slide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Slide and its children
func (m *CT_Slide) Validate() error {
	return m.ValidateWithPath("CT_Slide")
}

// ValidateWithPath validates the CT_Slide and its children, prefixing error messages with path
func (m *CT_Slide) ValidateWithPath(path string) error {
	if err := m.CSld.ValidateWithPath(path + "/CSld"); err != nil {
		return err
	}
	if m.ClrMapOvr != nil {
		if err := m.ClrMapOvr.ValidateWithPath(path + "/ClrMapOvr"); err != nil {
			return err
		}
	}
	if m.Transition != nil {
		if err := m.Transition.ValidateWithPath(path + "/Transition"); err != nil {
			return err
		}
	}
	if m.Timing != nil {
		if err := m.Timing.ValidateWithPath(path + "/Timing"); err != nil {
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
