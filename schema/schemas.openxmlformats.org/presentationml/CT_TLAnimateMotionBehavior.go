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

type CT_TLAnimateMotionBehavior struct {
	// Origin
	OriginAttr ST_TLAnimateMotionBehaviorOrigin
	// Path
	PathAttr *string
	// Path Edit Mode
	PathEditModeAttr ST_TLAnimateMotionPathEditMode
	// Relative Angle
	RAngAttr *int32
	// Points Types
	PtsTypesAttr *string
	CBhvr        *CT_TLCommonBehaviorData
	By           *CT_TLPoint
	// From
	From *CT_TLPoint
	To   *CT_TLPoint
	// Rotation Center
	RCtr *CT_TLPoint
}

func NewCT_TLAnimateMotionBehavior() *CT_TLAnimateMotionBehavior {
	ret := &CT_TLAnimateMotionBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLAnimateMotionBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OriginAttr != ST_TLAnimateMotionBehaviorOriginUnset {
		attr, err := m.OriginAttr.MarshalXMLAttr(xml.Name{Local: "origin"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PathAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "path"},
			Value: fmt.Sprintf("%v", *m.PathAttr)})
	}
	if m.PathEditModeAttr != ST_TLAnimateMotionPathEditModeUnset {
		attr, err := m.PathEditModeAttr.MarshalXMLAttr(xml.Name{Local: "pathEditMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RAngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rAng"},
			Value: fmt.Sprintf("%v", *m.RAngAttr)})
	}
	if m.PtsTypesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ptsTypes"},
			Value: fmt.Sprintf("%v", *m.PtsTypesAttr)})
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.By != nil {
		seby := xml.StartElement{Name: xml.Name{Local: "p:by"}}
		e.EncodeElement(m.By, seby)
	}
	if m.From != nil {
		sefrom := xml.StartElement{Name: xml.Name{Local: "p:from"}}
		e.EncodeElement(m.From, sefrom)
	}
	if m.To != nil {
		seto := xml.StartElement{Name: xml.Name{Local: "p:to"}}
		e.EncodeElement(m.To, seto)
	}
	if m.RCtr != nil {
		serCtr := xml.StartElement{Name: xml.Name{Local: "p:rCtr"}}
		e.EncodeElement(m.RCtr, serCtr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimateMotionBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "origin" {
			m.OriginAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "path" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PathAttr = &parsed
		}
		if attr.Name.Local == "pathEditMode" {
			m.PathEditModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "rAng" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RAngAttr = &pt
		}
		if attr.Name.Local == "ptsTypes" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PtsTypesAttr = &parsed
		}
	}
lCT_TLAnimateMotionBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cBhvr":
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case "by":
				m.By = NewCT_TLPoint()
				if err := d.DecodeElement(m.By, &el); err != nil {
					return err
				}
			case "from":
				m.From = NewCT_TLPoint()
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case "to":
				m.To = NewCT_TLPoint()
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			case "rCtr":
				m.RCtr = NewCT_TLPoint()
				if err := d.DecodeElement(m.RCtr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLAnimateMotionBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateMotionBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimateMotionBehavior and its children
func (m *CT_TLAnimateMotionBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateMotionBehavior")
}

// ValidateWithPath validates the CT_TLAnimateMotionBehavior and its children, prefixing error messages with path
func (m *CT_TLAnimateMotionBehavior) ValidateWithPath(path string) error {
	if err := m.OriginAttr.ValidateWithPath(path + "/OriginAttr"); err != nil {
		return err
	}
	if err := m.PathEditModeAttr.ValidateWithPath(path + "/PathEditModeAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.By != nil {
		if err := m.By.ValidateWithPath(path + "/By"); err != nil {
			return err
		}
	}
	if m.From != nil {
		if err := m.From.ValidateWithPath(path + "/From"); err != nil {
			return err
		}
	}
	if m.To != nil {
		if err := m.To.ValidateWithPath(path + "/To"); err != nil {
			return err
		}
	}
	if m.RCtr != nil {
		if err := m.RCtr.ValidateWithPath(path + "/RCtr"); err != nil {
			return err
		}
	}
	return nil
}
