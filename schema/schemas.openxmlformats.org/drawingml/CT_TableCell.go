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

type CT_TableCell struct {
	RowSpanAttr  *int32
	GridSpanAttr *int32
	HMergeAttr   *bool
	VMergeAttr   *bool
	IdAttr       *string
	TxBody       *CT_TextBody
	TcPr         *CT_TableCellProperties
	ExtLst       *CT_OfficeArtExtensionList
}

func NewCT_TableCell() *CT_TableCell {
	ret := &CT_TableCell{}
	return ret
}

func (m *CT_TableCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RowSpanAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowSpan"},
			Value: fmt.Sprintf("%v", *m.RowSpanAttr)})
	}
	if m.GridSpanAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gridSpan"},
			Value: fmt.Sprintf("%v", *m.GridSpanAttr)})
	}
	if m.HMergeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hMerge"},
			Value: fmt.Sprintf("%d", b2i(*m.HMergeAttr))})
	}
	if m.VMergeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vMerge"},
			Value: fmt.Sprintf("%d", b2i(*m.VMergeAttr))})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	if m.TxBody != nil {
		setxBody := xml.StartElement{Name: xml.Name{Local: "a:txBody"}}
		e.EncodeElement(m.TxBody, setxBody)
	}
	if m.TcPr != nil {
		setcPr := xml.StartElement{Name: xml.Name{Local: "a:tcPr"}}
		e.EncodeElement(m.TcPr, setcPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rowSpan" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RowSpanAttr = &pt
		}
		if attr.Name.Local == "gridSpan" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.GridSpanAttr = &pt
		}
		if attr.Name.Local == "hMerge" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HMergeAttr = &parsed
		}
		if attr.Name.Local == "vMerge" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VMergeAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
lCT_TableCell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "txBody":
				m.TxBody = NewCT_TextBody()
				if err := d.DecodeElement(m.TxBody, &el); err != nil {
					return err
				}
			case "tcPr":
				m.TcPr = NewCT_TableCellProperties()
				if err := d.DecodeElement(m.TcPr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TableCell %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableCell
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableCell and its children
func (m *CT_TableCell) Validate() error {
	return m.ValidateWithPath("CT_TableCell")
}

// ValidateWithPath validates the CT_TableCell and its children, prefixing error messages with path
func (m *CT_TableCell) ValidateWithPath(path string) error {
	if m.TxBody != nil {
		if err := m.TxBody.ValidateWithPath(path + "/TxBody"); err != nil {
			return err
		}
	}
	if m.TcPr != nil {
		if err := m.TcPr.ValidateWithPath(path + "/TcPr"); err != nil {
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
